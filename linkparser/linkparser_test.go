package linkparser

import (
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	needle := "https://github.com/gophercises"
	html := `
<html>
<head>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
<h1>Social stuffs</h1>
<div>
    <a href="https://www.twitter.com/joncalhoun">
        Check me out on twitter
        <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
        Gophercises is on <strong>Github</strong>!
    </a>
</div>
</body>
</html>
`

	links := Parser{}.Parse(strings.NewReader(html))

	if len(links) != 2 {
		t.Fatalf("The number of links should be 2.")
	}

	if links[1].Href != needle {
		t.Errorf("The second link of %v should redirect to %v", html, needle)
	}
}
