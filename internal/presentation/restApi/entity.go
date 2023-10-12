package restApi

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type responseList struct {
	List         []any `json:"list"`
	Page         int   `json:"page"`
	NextPage     int   `json:"next_page"`
	PrevPage     int   `json:"prev_page"`
	TotalPage    int   `json:"total_page"`
	TotalItems   int   `json:"total_items"`
	ItemsPerPage int   `json:"items_per_page"`
}
