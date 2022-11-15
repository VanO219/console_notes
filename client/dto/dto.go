package dto

type Note struct {
	ID   int64  `json:"id,omitempty"`
	User User   `json:"user,omitempty"`
	Text string `json:"text,omitempty"`
}

type User struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
}
