package models

// {id: 5, description: "2", priority: "medium", status: "pending", title: "1"}
type Tasks struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Title       string `json:"title"`
}
