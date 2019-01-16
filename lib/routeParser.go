package lib

import (
	"regexp"
	"strings"
)

func GenerateRegex(route string) *regexp.Regexp {
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
