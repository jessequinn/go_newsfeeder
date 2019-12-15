package newsfeed

import "database/sql"

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	DB *sql.DB
}

func New(db *sql.DB) *Repo {
	stmt, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "newsfeed" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(255),
	"post"	TEXT
	);`)
	stmt.Exec()

	return &Repo{
		DB: db,
	}
}

func (r *Repo) Add(item Item) {
	stmt, _ := r.DB.Prepare(`
	INSERT INTO newsfeed (title, post) VALUES (?,?);
	`)
	stmt.Exec(item.Title, item.Post)
}

func (r *Repo) GetAll() []Item {
	items := []Item{}
	rows, _ := r.DB.Query(`SELECT * FROM newsfeed;`)
	var id int
	var title string
	var post string

	for rows.Next() {
		rows.Scan(&id, &title, &post)
		item := Item{
			ID:    id,
			Title: title,
			Post:  post,
		}
		items = append(items, item)
	}

	return items
}
