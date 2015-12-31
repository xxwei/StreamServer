package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type LinkedList struct {
	value interface{}
	next  *LinkedList
}

// Slice map interface channel 按引用类型传递
func Test_Login(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/api/login", nil)
	if err != nil {
		t.Fatal(err)
	}
}
