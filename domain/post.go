package domain

type Post struct {
	ID      int    `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}

type Posts []Post
