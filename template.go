/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          template.go
 * Description:   Beego Template Engine
 */

package template

import (
	"errors"
	"github.com/going/toolkit/log"
	"strings"
)

type Template struct {
	name     string
	text     string
	line     int
	position int
	chunks   []Node
	runes    []rune
}

func NewTemplate(name, text string) *Template {
	return &Template{
		name:   name,
		text:   text,
		line:   1,
		chunks: make([]Node, 0, 32),
	}
}

func (t *Template) Render(params map[string]interface{}) {

}

func (t *Template) find(element string, start int) (int, error) {
	if start < 0 {
		return -1, errors.New("start must great than or equal 0!")
	}
	start += t.position
	index := start + strings.Index(t.text[start:], element)
	if index != -1 {
		index -= t.position
	}
	return index, nil
}

func (t *Template) consume(count int) string {
	if count == -1 {
		count = len(t.text) - t.position
	}
	newposition := t.position + count
	t.line += strings.Count(t.text[t.position:newposition], "\n")
	subtext := t.text[t.position:newposition]
	t.position = newposition
	return subtext
}

func (t *Template) remaining() int {
	return len(t.text) - t.position
}

func (t *Template) residue() string {
	return t.text[t.position:]
}

func (t *Template) next(offset int) string {
	return string(t.text[t.position+offset])
}

func parseTemplate(t *Template, block, loop string) []Node {
	chunks := make([]Node, 0, 5)
	for {
		char := ""
		curly := 0
		for {
			curly, _ = t.find("{", curly)
			if curly == -1 || curly+1 == t.remaining() {
				chunks = append(chunks, TextNode{value: t.consume(-1), line: t.line})
				// return
				return chunks
			}
			char = t.next(curly + 1)
			if !strings.Contains("{%#", char) {
				curly += 1
				continue
			}
			if curly+2 < t.remaining() && t.next(curly+1) == "{" && t.next(curly+2) == "{" {
				curly += 1
				continue
			}
			break
		}
		if curly > 0 {
			chunks = append(chunks, TextNode{value: t.consume(curly), line: t.line})
		}
		LeftDelimiter := t.consume(2) // delim left {{
		line := t.line

		if t.remaining() > 0 && t.next(0) == "!" {
			t.consume(1)
			chunks = append(chunks, TextNode{value: LeftDelimiter, line: line})
			continue
		}

		// Comment
		if LeftDelimiter == "{#" {
			end, err := t.find("#}", 0)
			if err != nil {
				log.Error(err)
				// TODO Error
			}
			if end == -1 {
				log.Error("-1")
				// TODO Error
			}
			t.consume(end) // comments
			t.consume(2)   // skip comment end delim)
			continue
		}

		// Expression
		if LeftDelimiter == "{{" {
			end, err := t.find("}}", 0)
			if err != nil {
				log.Error(err)
				// TODO Error
			}
			if end == -1 {
				log.Error("-1")
				// TODO Error
			}
			contents := strings.Trim(t.consume(end), " ")
			t.consume(2) // skip comment end delim
			if contents == "" {
				// TODO Error
			}
			chunks = append(chunks, ExpressionNode{expression: contents, line: line})
			continue
		}

		// Block
		if LeftDelimiter == "{%" {
			end, err := t.find("%}", 0)
			if err != nil {
				log.Error(err)
				// TODO Error
			}
			if end == -1 {
				log.Error("-1")
				// TODO Error
			}
			contents := strings.Trim(t.consume(end), " ")
			t.consume(2) // skip comment end delim
			if contents == "" {
				// TODO Error
			}
			spaceIndex := strings.Index(contents, " ")
			if spaceIndex == -1 {
				spaceIndex = len(contents)
			}
			operator := contents[:spaceIndex]
			suffix := strings.Trim(contents[spaceIndex:], " ")

			switch operator {
			case "end":
				return chunks
			case "extends":
			case "block":
				if suffix == "" {
					log.Error("block missing name!")
					break
				}
				chunks = append(chunks, ExpressionNode{expression: contents, line: line})
				block := parseTemplate(t, operator, "")
				chunks = append(chunks, NamedBlockNode{name: suffix, body: block, line: t.line})
			}
			continue
		}
	}
}

func (t *Template) parse() {
	t.chunks = parseTemplate(t, "", "")
}
