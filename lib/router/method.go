package router

import (
	"fmt"
	"net/http"
)

func (p *Path) Post(u string, h http.Handler) {
	path := p.Add(u)
	uri := path.Resolve()
	fmt.Println(uri)
	path.router.mux.Handle(uri, h)
}

func (p *Path) Get(u string, h http.Handler) {
	path := p.Add(u)
	uri := path.Resolve()
	fmt.Println(uri)
	path.router.mux.Handle(uri, h)
}

func (p *Path) Put(u string, h http.Handler) {
	path := p.Add(u)
	uri := path.Resolve()
	fmt.Println(uri)
	path.router.mux.Handle(uri, h)
}

func (p *Path) Delete(u string, h http.Handler) {
	path := p.Add(u)
	uri := path.Resolve()
	fmt.Println(uri)
	path.router.mux.Handle(uri, h)
}

/*
type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers     []*controllerInfo
	Application *App
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {
	parts := strings.Split(pattern, "/")

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

	//recreate the url pattern, with parameters replaced
	//by regular expressions. Then compile the regex.

	pattern = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		//TODO add error handling here to avoid panic
		panic(regexErr)
		return
	}

	//now create the Route
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controllerType = t

	p.routers = append(p.routers, route)
}
*/
