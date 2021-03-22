package controller

type AddForumRequest struct {
	Title string `json:"title"`
	Topic string `json:"topic"`
	Text  string `json:"text"`
	Token string `json:"token"`
}

type AddForumResponse struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Topic  string `json:"topic"`
}

type GetForumResponse struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Topic  string `json:"topic"`
	Text   string `json:"text"`
}
