package router

type Path struct {
	uri    string
	head   *Path
	router *Router
	paths  []*Path
}

func (p *Path) Add(u string) *Path {
	path := &Path{}
	path.uri = u
	path.head = p
	path.router = p.router
	p.paths = append(p.paths, path)
	return path
}

func (p *Path) Resolve() string {
	if p.head != nil {
		return p.head.Resolve() + "" + p.uri
	}
	return p.uri
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
