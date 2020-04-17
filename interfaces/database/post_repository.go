package database

import (
	"first_docker_project/domain"
	"github.com/jmoiron/sqlx"
)

func FindAll(db *sqlx.DB) (domain.Posts, error) {
	var posts []domain.Post
	if err := db.Select(&posts, "SELECT id, title, content FROM posts"); err != nil {
		return nil, err
	}

	return posts, nil
}

func Store(db *sqlx.DB, p domain.Post) (int, error) {
	res, err := db.Exec("INSERT INTO posts (title, content) VALUES (?,?)", p.Title, p.Content)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

func FindByID(db *sqlx.DB, id int) (*domain.Post, error) {
	var post domain.Post

	if err := db.Get(&post, "SELECT id ,title, content FROM posts WHERE id = ?", id); err != nil {
		return nil, err
	}

	return &post, nil
}

func Update(db *sqlx.DB, p domain.Post) error {
	_, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", p.Title, p.Content, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func Delete(db *sqlx.DB, id int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
