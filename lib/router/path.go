package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Path struct {
	uri   string
	head  *Path
	mux   *http.ServeMux
	route string
	regex *regexp.Regexp
}

func (p *Path) Add(u string) *Path {
	path := &Path{}
	path.uri = u
	path.head = p
	path.mux = p.mux
	return path
}

func (p *Path) resolvePath() string {
	if p.head != nil {
		return p.head.resolvePath() + "" + p.uri
	}
	return p.uri
}

func generateRegex(route string) *regexp.Regexp {
	parts := strings.Split(route, "/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			//a user may choose to override the default expression
			// similar to expressjs: ‘/user/:id([0-9]+)’

			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}

	//recreate the url uri, with parameters replaced
	//by regular expressions. Then compile the regex.

	route = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(route)
	if regexErr != nil {

		//TODO add error handling here to avoid panic
		panic(regexErr)
	}
	return regex
}

func (p *Path) Get(s string, next http.HandlerFunc) {
	path := p.resolvePath()
	route := generateRegex(path + "" + s)
	fmt.Println(path)
	fmt.Println(route)
	p.mux.HandleFunc(path+""+s, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case route.MatchString(r.URL.Path):
			next(w, r)
		default:
			w.Write([]byte("Unknown Pattern"))
		}
	})
}
