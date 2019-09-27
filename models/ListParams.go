package models

//ListParams struct for tables
type ListParams struct {
	Q          string `schema:"q"`
	StartIndex int    `schema:"startIndex"`
	MaxResults int    `schema:"maxResults"`
	Sort       string `schema:"sort"`
}
