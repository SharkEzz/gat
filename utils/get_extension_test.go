package utils_test

import (
	"testing"

	"github.com/SharkEzz/gat/utils"
)

func TestGetExtension(t *testing.T) {
	fileName := "test.md"
	ext := utils.GetExtension(fileName)

	if ext != "md" {
		t.Errorf("expected extension to be md, got %s", ext)
	}

	fileName = "test.go"
	ext = utils.GetExtension(fileName)

	if ext != "go" {
		t.Errorf("expected extension to be go, got %s", ext)
	}

	fileName = "test.c"
	ext = utils.GetExtension(fileName)

	if ext != "c" {
		t.Errorf("expected extension to be c, got %s", ext)
	}

	fileName = "test."
	ext = utils.GetExtension(fileName)

	if ext != "" {
		t.Errorf("expected extension to be empty, got %s", ext)
	}

	fileName = "test"
	ext = utils.GetExtension(fileName)

	if ext != "" {
		t.Errorf("expected extension to be empty, got %s", ext)
	}
}
