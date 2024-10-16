package pagination

import (
	"iter"
	"slices"
)

type Page[T any] struct {
	Page    int `json:"page"`
	Size    int `json:"size"`
	Content []T `json:"content"`
	Total   int `json:"total"`
}

func NewPage[T any](page int, size int, content []T, total int) *Page[T] {

	return &Page[T]{
		Page:    page,
		Size:    size,
		Content: content,
		Total:   total,
	}
}

func NewPageFromPageRequest[T any](pageRequest *PageRequest, content []T, total int) *Page[T] {
	return NewPage(pageRequest.Page, pageRequest.Size, content, total)
}

func ApplyPageRequestToIter[T any](pageRequest *PageRequest, iter iter.Seq[T]) *Page[T] {
	all := slices.Collect(iter)
	return ApplyPageRequest(pageRequest, all)
}

func ApplyPageRequest[T any](pageRequest *PageRequest, all []T) *Page[T] {
	total := len(all)
	start := pageRequest.Page * pageRequest.Size
	end := start + pageRequest.Size

	if start > total {
		start = total
	}

	if end > total {
		end = total
	}

	content := all[start:end]
	return NewPageFromPageRequest(pageRequest, content, total)
}
