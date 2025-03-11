package handlers

import (
	"app/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request, repo db.PostRepo) {
	posts, err := repo.QueryPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("error querying database: %v", err), http.StatusInternalServerError)
		return
	}

	var responsePosts []Post = make([]Post, 0)
	for _, post := range posts {
		responsePosts = append(responsePosts, toPostResponse(&post))
	}
	if err := json.NewEncoder(w).Encode(responsePosts); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request, repo db.PostRepo) {
	request, err := toPostRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postResponse, err := repo.CreatePost(*request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error processing request: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(toPostResponse(postResponse)); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func toPostRequest(r *http.Request) (*db.PostRequest, error) {
	var postRequest PostRequest

	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		log.Println("Failed to decode body", err.Error())
		return nil, fmt.Errorf("invalid JSON payload: " + err.Error())
	}
	return &db.PostRequest{
		AuthordId: postRequest.AuthordId,
		ThreadId:  postRequest.ThreadId,
		Content:   postRequest.Content,
	}, nil
}

func toPostResponse(post *db.Post) Post {
	return Post{
		Id:        post.ID,
		CreatedAt: post.CreatedAt,
		AuthordId: post.AuthordId,
		Author:    post.Author,
		ThreadId:  post.ThreadId,
		OP:        post.OP,
		Content:   post.Content,
	}
}
