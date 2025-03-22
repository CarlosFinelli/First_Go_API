package classes

type Album struct {
	Id     int      `json:"id"`
	Title  *string  `json:"title"`
	Artist *string  `json:"artist"`
	Price  *float32 `json:"price"`
}
