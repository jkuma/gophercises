package database

import "testing"

func TestDB(t *testing.T) {
	db := DB()

	if err := db.Ping(); err != nil {
		t.Errorf("The database it not up: %v", err)
	}
}
