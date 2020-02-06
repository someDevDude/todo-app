package models

//TodoFull object
type TodoFull struct {
	ID          int64  `schema:"id" db:"ID"`
	Title       string `schema:"title" db:"TITLE"`
	Description string `schema:"description" db:"DESCRIPTION"`
	Done        bool   `schema:"done" db:"DONE"`
}
