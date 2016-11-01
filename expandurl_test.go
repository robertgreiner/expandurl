package expandurl

import (
	"testing"
)

func TestCanExpandValidUrl(t *testing.T) {
	result, err := Expand("http://bit.ly/2eXwN8A")

	if err != nil {
		t.Error("Should have successfully expanded URL")
	}

	if result != "http://robertgreiner.com/" {
		t.Error("Returned incorrect URL")
	}
}

func TestShouldReturnErrorOnInvalidCharacter(t *testing.T) {
	_, err := Expand("http://b it.ly/2eXwN8A")

	if err == nil {
		t.Error("Should have returned error on invalid character in URL")
	}
}

func TestShouldNotReturnErrorOnValidUrl(t *testing.T) {
	_, err := Expand("http://bit.ly/2eXwN8A")

	if err != nil {
		t.Error("Should not have thrown error on valid URL")
	}
}

func TestShouldDecodeEscapedCharacters(t *testing.T) {
	result, err := Expand("http:%2F%2Fbit.ly%2F2eXwN8A")

	if err != nil {
		t.Error("Should have successfully expanded URL with escaped characters")
	}

	if result != "http://robertgreiner.com/" {
		t.Error("Returned incorrect URL")
	}
}
