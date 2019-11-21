package stkhtml

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

type StackTrace struct {
	Traces []Trace
}

type Trace struct {
	Command  string
	Url      string
	Filename string
}

var defaultTemplate = `
<ol>
	{{range .Traces}}
		<li>{{.Command}}</li>
		<ul>
			<li><a href="{{.Url}}">{{.Filename}}</a></li>
		</ul>
	{{end}}
</ol>
`

func HtmlOutput(b []byte) (string, error) {
	var buf bytes.Buffer
	st := parseStackTrace(b)
	err := template.Must(template.New("").Parse(defaultTemplate)).Execute(&buf, st)

	return buf.String(), err
}

func parseStackTrace(b []byte) StackTrace {
	re := regexp.MustCompile(`(\/.+):(\d+)`)
	html := strings.Split(string(b), "\n")
	html = append(html[:0], html[1:]...)
	st := StackTrace{Traces: make([]Trace, len(html)/2)}

	for i, t := range st.Traces {
		// k is the current position in html slice.
		k := len(html)%len(html)/2 + i*2

		if k+1 < len(html) {
			res := re.FindAllSubmatch([]byte(html[k+1]), -1)

			if len(res) > 0 {
				t.Command, t.Url, t.Filename = html[k], fmt.Sprintf("%s/%s", res[0][1], res[0][2]), html[k+1]
				st.Traces[i] = t
			}
		}
	}

	return st
}
