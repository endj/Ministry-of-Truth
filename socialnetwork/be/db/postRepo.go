package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type PostRepo struct {
	DB *sql.DB
}

func (u PostRepo) QueryPosts() ([]Post, error) {
	rows, err := u.DB.Query(`
	SELECT 
    	user_posts.id AS post_id,
    	user_posts.created_at as timestamp,
    	user_posts.thread_id,
    	user_posts.user_id,
    	user_profiles.name AS author_name,
		user_posts.op,
    	user_posts.content
    FROM user_posts
    JOIN user_profiles ON user_posts.user_id = user_profiles.id
	ORDER BY timestamp DESC;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query user posts: %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post

		if err := rows.Scan(
			&post.ID,
			&post.CreatedAt,
			&post.ThreadId,
			&post.AuthordId,
			&post.Author,
			&post.OP,
			&post.Content,
		); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}
	return posts, err
}

func (u PostRepo) CreatePost(post PostRequest) (*Post, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	log.Println("Creating post: ", post)

	var openingPost = post.ThreadId == ""
	if openingPost {
		post.ThreadId = hashUniqueID()
	}

	var createdAt = time.Now().UnixMilli()
	var postId int
	if err := tx.QueryRow(
		`INSERT INTO user_posts (user_id, created_at, thread_id, content, op)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id
		`, post.AuthordId, createdAt, post.ThreadId, post.Content, boolToInt(openingPost),
	).Scan(&postId); err != nil {
		return nil, fmt.Errorf("error inserting posts into user_posts: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %v", err)
	}

	return &Post{
		ID:        postId,
		CreatedAt: createdAt,
		AuthordId: post.AuthordId,
		ThreadId:  post.ThreadId,
		Content:   post.Content,
		OP:        boolToInt(openingPost),
	}, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func hashUniqueID() string {
	data := fmt.Sprintf("%d", time.Now().UnixNano())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16]
}
