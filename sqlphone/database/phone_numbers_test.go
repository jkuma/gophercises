package database

import "testing"

func TestNumberExists(t *testing.T) {
	for _, number := range numbers() {
		if !NumberExists(number) {
			t.Errorf("Number %v should exist in phone_numbers table.", number)
		}
	}
}
