package routing

import (
	"fmt"
	"net/http"
	"regexp"
)

type Handler struct {
	routes []Route
}

func (h *Handler) RegisterRoute(r Route) {
	h.routes = append(h.routes, r)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if !h.isHttpMethodAllowed(route, *r) {
			continue
		}
		if !h.doRequestPathMatch(route, *r) {
			continue
		}
		pathVars := h.extractPathVariables(route, *r)
		route.Handler(w, r, pathVars)
		return
	}
	fmt.Println("No pattern found")
}

func (h Handler) isHttpMethodAllowed(route Route, r http.Request) bool {
	for _, m := range route.HttpVerbs {
		if m == r.Method {
			return true
		}
	}
	return false
}

func (h Handler) doRequestPathMatch(route Route, r http.Request) bool {
	validId := regexp.MustCompile(route.Pattern)
	if validId.MatchString(r.URL.Path) {
		return true
	}
	return false
}

func (h Handler) extractPathVariables(route Route, r http.Request) []string {
	validId := regexp.MustCompile(route.Pattern)
	matches := validId.FindStringSubmatch(r.URL.Path)
	return append(matches[1:])
}
