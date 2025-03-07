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
	SELECT *
	FROM user_posts
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query user posts: %w", err)
	}
	defer rows.Close()

	var posts []Post = make([]Post, 0)
	for rows.Next() {
		var post Post

		err := rows.Scan(
			&post.ID,
			&post.CreatedAt,
			&post.ThreadId,
			&post.AuthordId,
			&post.Content,
		)
		if err != nil {
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

	if post.ThreadId == "" {
		post.ThreadId = hashUniqueID()
	}

	var createdAt = time.Now().UnixMilli()
	var postId int
	err = tx.QueryRow(
		`INSERT INTO user_posts (user_id, created_at, thread_id, content)
		 VALUES ($1, $2, $3, $4) RETURNING id
		`, post.AuthordId, createdAt, post.ThreadId, post.Content,
	).Scan(&postId)

	if err != nil {
		return nil, fmt.Errorf("error inserting posts into user_posts: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %v", err)
	}

	postResponse := Post{
		ID:        postId,
		CreatedAt: createdAt,
		AuthordId: post.AuthordId,
		ThreadId:  post.ThreadId,
		Content:   post.Content,
	}
	log.Println("REturning", postResponse)
	return &postResponse, nil
}

func hashUniqueID() string {
	data := fmt.Sprintf("%d", time.Now().UnixNano())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16] // Shorten if needed
}
