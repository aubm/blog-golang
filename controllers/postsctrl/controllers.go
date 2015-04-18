package postsctrl

import (
	"encoding/json"
	"fmt"
	"github.com/aubm/blog/models"
	"github.com/aubm/blog/services/postsservice"
	"io"
	"net/http"
	"strconv"
)

func IndexController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	posts := postsservice.GetPosts()
	b, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b[:]))
}

func DetailsController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	postId, _ := strconv.ParseInt(pathVars[0], 10, 64)
	post := postsservice.GetOnePost(postId)
	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b[:]))
}

func AddController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	newPost := models.Post{
		Title:   title,
		Content: content,
	}
	postsservice.SavePost(&newPost)
	b, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b[:]))
}

func UpdateController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	// TODO
}

func DeleteController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	// TODO
}
