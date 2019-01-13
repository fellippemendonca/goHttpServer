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

func (p *Path) GetPath() (string, map[int]string) {
	uri := p.GetUri()
	parts := strings.Split(uri, "/")
	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"
			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}
	uri = strings.Join(parts, "/")
	return uri, params
	/*
		regex, regexErr := regexp.Compile(uri)
		if regexErr != nil {
			panic(regexErr)
			return nil
		}
	*/
}
