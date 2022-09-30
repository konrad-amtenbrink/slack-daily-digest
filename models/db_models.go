package models

type User struct {
	ID string `db:"id"`
}

type Thread struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Url   string `db:"url"`
}
