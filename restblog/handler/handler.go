package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"restblog/model"
)

var allPosts = make(map[string]model.Post) // Capitalized 'Post' for exported struct

// POST (Create)
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusInternalServerError)
		return
	}

	if _, ok := allPosts[post.Title]; ok {
		http.Error(w, "Post title already exists", http.StatusBadRequest) // fixed syntax: . to ,
		return
	}

	allPosts[post.Title] = post

	fmt.Fprintf(w, "%+v", post) // Use Fprintf to write formatted output
}

// GET
func ListPosts(w http.ResponseWriter, r *http.Request) {
	titles := []string{}

	for _, post := range allPosts {
		titles = append(titles, post.Title)
	}

	if len(titles) == 0 {
		http.Error(w, "No posts found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(titles)
}

// GET BY TITLE
func GetPostByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title") // Declare the variable
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	post, ok := allPosts[title]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(post); err != nil { // fixed syntax error: `! =` to `!=`
		http.Error(w, "Failed to encode post", http.StatusInternalServerError)
	}
}

// PUT
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatedPost model.Post

	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	title := updatedPost.Title
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// check if such post exists
	_, ok := allPosts[title]
	if !ok {
		http.Error(w, "BlogPost not found", http.StatusNotFound)
		return
	}

	// update post
	allPosts[title] = updatedPost

	// return ok status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPost)
}

// DELETE
func DeletePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	if _, ok := allPosts[title]; !ok { // Declare and check 'ok' in same line
		http.Error(w, "No post with such title", http.StatusNotFound)
		return
	}

	delete(allPosts, title)
	w.WriteHeader(http.StatusOK)
}
