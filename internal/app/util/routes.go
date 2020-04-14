package util

import (
	"runtime/debug"
	"strings"

	"github.com/gorilla/mux"
)

type RouteDescription struct {
	Name    string
	Methods string
	Path    string
}

func ExtractRoutes(r *mux.Router) []RouteDescription {
	ret := []RouteDescription{}
	var _ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		pathTemplate, _ := route.GetPathTemplate()
		name := route.GetName()
		m := strings.Join(methods, ", ")
		ret = append(ret, RouteDescription{name, m, pathTemplate})
		return nil
	})
	return ret
}

func ExtractModules() *debug.BuildInfo {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}
	return bi
}
