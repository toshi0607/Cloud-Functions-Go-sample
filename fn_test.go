package hello

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	q := r.URL.Query()
	q.Add("name", "toshi0607")
	r.URL.RawQuery = q.Encode()

	Hello(w, r)

	rw := w.Result()
	defer rw.Body.Close()
	wantStatus := http.StatusOK
	if s := rw.StatusCode; s != wantStatus {
		t.Fatalf("got: %d, want: %d", s, wantStatus)
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("failed to read res body")
	}
	wantString := "Hello, toshi0607!"
	if s := string(b); s != wantString {
		t.Fatalf("got: %s, want: %s", s, wantString)
	}
}
