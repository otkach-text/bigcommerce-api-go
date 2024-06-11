package bigcommerce

// Video is entry for BC product videos
type Video struct {
	ID          int64  `json:"id"`
	ProductID   int64  `json:"product_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   int64  `json:"sort_order"`
	Type        string `json:"type"`
	VideoID     string `json:"video_id"`
	Length      string `json:"length"`
}
