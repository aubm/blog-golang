package routing

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"
)

type Handler struct {
	routes     []Route
	publicDirs []string
}

func (h *Handler) RegisterRoute(r Route) {
	h.routes = append(h.routes, r)
}

func (h *Handler) AddPublicDir(newPublicDir string) {
	h.publicDirs = append(h.publicDirs, newPublicDir)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// First check if we can serve a static file
	for _, dir := range h.publicDirs {
		filePathName := path.Join(dir, r.URL.Path)
		if fileInfo, err := os.Stat(filePathName); err == nil && fileInfo.IsDir() == false {
			http.ServeFile(w, r, filePathName)
			return
		}
	}

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
