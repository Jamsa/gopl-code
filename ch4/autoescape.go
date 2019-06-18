package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>" /* 这里的特殊字符会被转义 */
	data.B = "<b>Hello!</b>" /* 这里能正常输出html源码 */
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
