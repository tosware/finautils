package finautils

import "testing"

func TestFinRefLen(t *testing.T) {
	_, err := CreateFinnishReference("10")
	if err == nil {
		t.Error("Accepted shorter than 3 chars base")
	}
	_, err = CreateFinnishReference("12345678901234567890")
	if err == nil {
		t.Error("Accepted loger than 19 chars base")
	}
}

func TestFinRefCreation(t *testing.T) {
	ref, err := CreateFinnishReference("100")
	if err != nil {
		t.Error("Error creating ref", err.Error())
	}
	if ref != "1009" {
		t.Error("Invalid Ref Number:", ref)
	}
}
