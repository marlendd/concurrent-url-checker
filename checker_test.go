package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Запускаем проверку нашего локального сервера
	urls := []string{server.URL}
	results := checkAll(urls, 1)

	if len(results) != 1 {
		t.Fatal("Должен быть 1 результат")
	}

	if results[0].Status != 200 {
		t.Errorf("Ожидался статус 200, получили %d", results[0].Status)
	}
	
	if results[0].Err != nil {
		t.Errorf("Ожидалась чистая проверка, получили ошибку: %v", results[0].Err)
	}
}