package pagination_test

import (
	"github.com/rcharre/psapi/pkg/utils/pagination"
	"testing"
)

func TestApplyPageRequest(t *testing.T) {
	content := []int{1, 2, 3, 4}
	contentLen := len(content)
	pr := pagination.NewPageRequest(1, 2)

	page := pagination.ApplyPageRequest(pr, content)

	resLen := len(page.Content)
	expectedLen := 2

	if resLen != expectedLen {
		t.Error("Page content len should be", expectedLen, ", has", resLen)
	}

	if page.Total != contentLen {
		t.Error("Page total should be", contentLen, ", has", page.Total)
	}

	if page.Page != pr.Page {
		t.Error("Page should be", pr.Page, ", has", page.Page)
	}

	if page.Size != pr.Size {
		t.Error("Page size should be", pr.Size, ", has", page.Size)
	}
}
