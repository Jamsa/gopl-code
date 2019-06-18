package main

/*
模板的使用
go run ch4/issueshtml.go repo:golang/go is:open json decoder
*/
import (
	"html/template" /* html模板比text模板增加了字符串自动转义等特性 */
	"log"
	"os"

	"gopl.io/ch4/github"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style="text-align: left">
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func main() {
	report := template.Must(template.New("issuelist").
		Parse(templ))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
