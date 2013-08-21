package hello

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

func TestHandleIndexContainsAmsterdam(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	body := response.Body.String()
	if !strings.Contains(body, "Amsterdam") {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "Amsterdam", body)
	}
}

func TestHandleIndexContainsSF(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	body := response.Body.String()
	if !strings.Contains(body, "San Francisco") {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "San Francisco", body)
	}
}
