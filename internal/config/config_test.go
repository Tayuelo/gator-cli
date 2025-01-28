package config

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	config, err := Read()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if config.DBURL != "postgres://example" {
		t.Errorf("DB URL don't match: %s", config.DBURL)
	}
}
