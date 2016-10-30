package expandurl

import (
	"testing"
)

func TestCanExpandValidUrl(t *testing.T) {
	result, err := Expand("http://bit.ly/2eXwN8A")

	if err != nil {
		t.Log("Should have successfully expanded URL")
		t.Fail()
	}

	if result != "http://robertgreiner.com/" {
		t.Log("Returned incorrect URL")
		t.Fail()
	}
}

func TestShouldReturnErrorOnInvalidCharacter(t *testing.T) {
	_, err := Expand("http://b it.ly/2eXwN8A")

	if err == nil {
		t.Log("Should have returned error on invalid character in URL")
		t.Fail()
	}
}

func TestShouldNotReturnErrorOnValidUrl(t *testing.T) {
	_, err := Expand("http://bit.ly/2eXwN8A")

	if err != nil {
		t.Log("Should not have thrown error on valid URL")
		t.Fail()
	}
}
