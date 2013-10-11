/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          example.go
 * Description:   example
 */

package main

import (
	"github.com/going/toolkit/log"
	"text/template/parse"
)

func main() {
	trees, err := parse.Parse("name", "<html>Dear {{.Name}}</html>", "{{", "}}", nil, nil)
	log.Println(trees["name"], err)
}
