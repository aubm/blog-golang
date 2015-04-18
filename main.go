package main

import (
	"github.com/aubm/blog-golang/controllers/postsctrl"
	"github.com/aubm/blog-golang/services/routing"
	"net/http"
)

func main() {
	var handler routing.Handler
	handler.AddPublicDir("./public")
	handler.RegisterRoute(routing.Route{`^/api/posts$`, []string{"GET"}, postsctrl.IndexController})
	handler.RegisterRoute(routing.Route{`^/api/posts/([0-9]+)$`, []string{"GET"}, postsctrl.DetailsController})
	handler.RegisterRoute(routing.Route{`^/api/posts$`, []string{"POST"}, postsctrl.AddController})
	handler.RegisterRoute(routing.Route{`^/api/posts/([0-9]+)$`, []string{"PATCH"}, postsctrl.UpdateController})
	handler.RegisterRoute(routing.Route{`^/api/posts/([0-9]+)$`, []string{"DELETE"}, postsctrl.DeleteController})
	http.ListenAndServe(":8080", handler)
}
