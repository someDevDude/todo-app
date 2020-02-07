package models

//TodoFull object
type TodoFull struct {
	ID          int64  `json:"id" db:"ID"`
	Title       string `json:"title" db:"TITLE"`
	Description string `json:"description" db:"DESCRIPTION"`
	Done        bool   `json:"done" db:"DONE"`
}
