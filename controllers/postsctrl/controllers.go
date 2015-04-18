package postsctrl

import (
	"encoding/json"
	"fmt"
	"github.com/aubm/blog-golang/models"
	"github.com/aubm/blog-golang/services/postsservice"
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
	post := findPostByStringId(pathVars[0])
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
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b[:]))
}

func UpdateController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	post := findPostByStringId(pathVars[0])
	title := r.PostFormValue("title")
	if title != "" {
		post.Title = title
	}
	content := r.PostFormValue("content")
	if content != "" {
		post.Content = content
	}
	postsservice.SavePost(&post)
	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b[:]))
}

func DeleteController(w http.ResponseWriter, r *http.Request, pathVars []string) {
	post := findPostByStringId(pathVars[0])
	postsservice.DeletePost(&post)
	w.WriteHeader(http.StatusNoContent)
}

func findPostByStringId(postIdString string) models.Post {
	postId, _ := strconv.ParseInt(postIdString, 10, 64)
	post := postsservice.GetOnePost(postId)
	return post
}
