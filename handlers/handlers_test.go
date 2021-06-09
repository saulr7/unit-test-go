package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	Get(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, w.Code)
	}

	// t.Logf(w.Body.String())

	got := Person{}

	err := json.NewDecoder(w.Body).Decode(&got)

	if err != nil {
		t.Errorf("Couldn't read body")
	}

	want := Person{"Saul", 20}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}

}
