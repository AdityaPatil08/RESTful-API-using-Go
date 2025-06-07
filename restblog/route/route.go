package route

import (
	middleware "restblog/auth"
	"restblog/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/create", handler.CreatePost).Methods("POST")

	r.HandleFunc("/list", handler.ListPosts).Methods("GET")

	r.HandleFunc("/post", handler.GetPostByTitle).Methods("GET")

	r.HandleFunc("/edit", handler.UpdatePost).Methods("PUT")

	r.HandleFunc("/delete", middleware.RequireAuth(handler.DeletePost)).Methods("DELETE")
	return r
}
