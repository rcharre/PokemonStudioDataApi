package validation_test

import (
	"psapi/pkg/utils/validation"
	"testing"
)

func TestValidation(t *testing.T) {
	validation := validation.NewValidation("key", "value")
	if validation.Key != "key" {
		t.Error("Key should be \"key\", get", validation.Key)
	}

	if validation.Value != "value" {
		t.Error("Value should be \"value\", get", validation.Value)
	}

	str := validation.String()
	if str != "key=value" {
		t.Error("String should be \"key=value\", get", str)
	}
}
