package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	testHeader := http.Header{
		"Authorization": {"yoo SOME_TOKEN"},
	}

	got, err := GetAPIKey(testHeader)
	if err != nil {
		t.Fatalf("GetAPIKey failed with error: %v", err)
	}
	want := "SOME_TOKEN"

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	testHeader := http.Header{
		"Authorization": {"HEllo ME_TOKEN"},
	}

	_, err := GetAPIKey(testHeader)
	if err == nil {
		t.Fatal("How???")
	}
}
