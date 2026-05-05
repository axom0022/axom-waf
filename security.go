package main

import (
	"regexp"
	"strings"
)

var sqlres []*regexp.Regexp
var xssres []*regexp.Regexp

func init() {
	for _, p := range config.sqlpatterns {
		sqlres = append(sqlres, regexp.MustCompile("(?i)"+p))
	}
	for _, p := range config.xsspatterns {
		xssres = append(xssres, regexp.MustCompile("(?i)"+p))
	}
}

func ismalicious(data string) bool {
	lower := strings.ToLower(data)
	for _, re := range sqlres {
		if re.MatchString(lower) {
			return true
		}
	}
	for _, re := range xssres {
		if re.MatchString(lower) {
			return true
		}
	}
	return false
}
