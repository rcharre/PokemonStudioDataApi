package pagination

type Page[T any] struct {
	Page    int `json:"page"`
	Size    int `json:"size"`
	Content []T `json:"content"`
	Total   int `json:"total"`
}

func NewPage[T any](page int, size int, content []T, total int) Page[T] {

	return Page[T]{
		Page:    page,
		Size:    size,
		Content: content,
		Total:   total,
	}
}

func ApplyPageRequest[T any](pageRequest PageRequest, all []T) Page[T] {
	total := len(all)
	start := pageRequest.Page * pageRequest.Size
	end := start + pageRequest.Size

	if start > total {
		start = total - 1
	}

	if end > total {
		end = total
	}

	content := all[start:end]
	return NewPage(pageRequest.Page, pageRequest.Size, content, total)
}
