package endpoint_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/mohammad-quanit/go-tdd/Server"
)

func TestPostsEndpoint(t *testing.T) {
	//create a req to pass our handler, pass nil as third param
	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatalf("Error: %v\n", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetPosts)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong status code, got %v, want %v\n", status, http.StatusOK)
	}

	// check the length of data in response
	var data []api.Response
	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Errorf("Erro decoding response body %v\n", err)
		return
	}
	gotExpectedLength := len(data)
	wantExpectedLength := 101
	if gotExpectedLength != wantExpectedLength {
		t.Errorf("Unexpected length of data: got %v, want %v", gotExpectedLength, wantExpectedLength)
	}
}
