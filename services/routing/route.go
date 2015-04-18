package routing

import (
	"net/http"
)

type Route struct {
	Pattern   string
	HttpVerbs []string
	Handler   func(w http.ResponseWriter, r *http.Request, pathVars []string)
}
