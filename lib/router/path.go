package router

import (
	"strings"
)

type Path struct {
	uri    string
	head   *Path
	router *Router
	paths  []*Path
	params map[int]string
}

func (p *Path) Add(u string) *Path {
	path := &Path{}
	path.uri = u
	path.head = p
	path.router = p.router
	p.paths = append(p.paths, path)
	return path
}

func (p *Path) GetUri() string {
	if p.head != nil {
		return p.head.GetUri() + "" + p.uri
	}
	return p.uri
}

func (p *Path) Resolve() (string, map[int]string) {
	uri := p.GetUri()
	parts := strings.Split(uri, "/")

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

    uri = strings.Join(parts, "/")
    //regex, regexErr := regexp.Compile(uri)
    //if regexErr != nil {

        //TODO add error handling here to avoid panic
        //panic(regexErr)
        //return uri, params
    //}
	return uri, params
}
