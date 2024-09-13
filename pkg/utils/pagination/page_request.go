package pagination

type PageRequest struct {
	Page int `json:"page" minimum:"0"`
	Size int `json:"size" minimum:"0"`
}

func NewPageRequest(page int, size int) *PageRequest {
	return &PageRequest{
		Page: page,
		Size: size,
	}
}
