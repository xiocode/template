/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          template_test.go
 * Description:   testing
 */

package template

import (
	"github.com/going/toolkit/log"
	"testing"
)

func TestTemplate(t *testing.T) {
	tpl := NewTemplate("name", "<html>{#A A A #}{% block title %}A bolder title{% end %}{{ myvalue }}</html>")
	tpl.parse()
	log.Println(tpl.chunks)
}
