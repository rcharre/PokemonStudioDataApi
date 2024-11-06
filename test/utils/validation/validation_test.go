package validation_test

import (
	"github.com/rcharre/psapi/pkg/utils/validation"
	"testing"
)

func TestValidation(t *testing.T) {
	validation := validation.NewValidation("key", "value")
	if validation.Key != "key" {
		t.Error("Key should be \"key\", has", validation.Key)
	}

	if validation.Value != "value" {
		t.Error("Value should be \"value\", has", validation.Value)
	}

	str := validation.String()
	if str != "key=value" {
		t.Error("String should be \"key=value\", has", str)
	}
}
