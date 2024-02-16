package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateEventHandler(t *testing.T) {
	// Создание фейкового запроса
	body := []byte(`{"id": 1, "date": "2024-02-15T15:04:05Z"}`)
	req, err := http.NewRequest("POST", "/create_event", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Создание фейкового объекта ResponseWriter для записи ответа
	rr := httptest.NewRecorder()

	// Вызов обработчика
	handler := http.HandlerFunc(createEventHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ошибка: ожидается %v, получено %v", http.StatusOK, status)
	}

	// Проверка тела ответа
	expected := "Row success created"
	if rr.Body.String() != expected {
		t.Errorf("ошибка: ожидается %v, получено %v", expected, rr.Body.String())
	}
}

func TestUpdateEventHandler(t *testing.T) {
	// Создание фейкового запроса с JSON-телом
	body := []byte(`{"id": 1, "date": "2024-02-15T15:04:05Z"}`)
	req, err := http.NewRequest("PUT", "/update_event", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Создание фейкового объекта ResponseWriter для записи ответа
	rr := httptest.NewRecorder()

	// Загрузка фейковых событий в список
	events = []Event{{ID: 1, Date: time.Now()}}

	// Вызов обработчика
	handler := http.HandlerFunc(updateEventHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ошибка: ожидается %v, получено %v", http.StatusOK, status)
	}

	// Проверка тела ответа
	expected := "Row success update"
	if rr.Body.String() != expected {
		t.Errorf("ошибка: ожидается %v, получено %v", expected, rr.Body.String())
	}

	// Проверка обновленного события в списке events
	if len(events) != 1 {
		t.Errorf("ошибка: ожидается 1 событие в списке, получено %d", len(events))
	}
	if events[0].ID != 1 || events[0].Date != time.Date(2024, 2, 15, 15, 4, 5, 0, time.UTC) {
		t.Errorf("ошибка: неверно обновленное событие")
	}
}

func TestDeleteEventHandler(t *testing.T) {
	// Создание фейкового DELETE-запроса с параметром id
	req, err := http.NewRequest("DELETE", "/delete_event?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создание фейкового объекта ResponseWriter для записи ответа
	rr := httptest.NewRecorder()

	// Загрузка фейковых событий в список
	events = []Event{{ID: 1, Date: time.Now()}}

	// Вызов обработчика
	handler := http.HandlerFunc(deleteEventHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ошибка: ожидается %v, получено %v", http.StatusOK, status)
	}

	// Проверка тела ответа
	expected := "Row success delete"
	if rr.Body.String() != expected {
		t.Errorf("ошибка: ожидается %v, получено %v", expected, rr.Body.String())
	}
}

func TestEventsForDayHandler(t *testing.T) {
	// Создание фейкового GET-запроса
	req, err := http.NewRequest("GET", "/events_for_day", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создание фейкового объекта ResponseWriter для записи ответа
	rr := httptest.NewRecorder()

	// Загрузка фейковых событий в список
	events = []Event{{ID: 1, Date: time.Now()}}

	// Вызов обработчика
	handler := http.HandlerFunc(eventsForDayHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ошибка: ожидается %v, получено %v", http.StatusOK, status)
	}

	// Проверка формата ответа (должен быть JSON)
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("ошибка: ожидается Content-Type: application/json, получено %v", contentType)
	}

}

func TestEventsForWeekHandler(t *testing.T) {
	// Создание фейкового GET-запроса
	req, err := http.NewRequest("GET", "/events_for_week", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создание фейкового объекта ResponseWriter для записи ответа
	rr := httptest.NewRecorder()

	// Загрузка фейковых событий в список
	events = []Event{{ID: 1, Date: time.Now()}}

	// Вызов обработчика
	handler := http.HandlerFunc(eventsForWeekHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ошибка: ожидается %v, получено %v", http.StatusOK, status)
	}

	// Проверка формата ответа (должен быть JSON)
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("ошибка: ожидается Content-Type: application/json, получено %v", contentType)
	}

}

func TestEventsForMonthHandler(t *testing.T) {
	// Создание фейкового GET-запроса
	req, err := http.NewRequest("GET", "/events_for_month", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создание фейкового объекта ResponseWriter для записи ответа
	rr := httptest.NewRecorder()

	// Загрузка фейковых событий в список
	events = []Event{{ID: 1, Date: time.Now()}}

	// Вызов обработчика
	handler := http.HandlerFunc(eventsForMonthHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ошибка: ожидается %v, получено %v", http.StatusOK, status)
	}

	// Проверка формата ответа (должен быть JSON)
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("ошибка: ожидается Content-Type: application/json, получено %v", contentType)
	}
}
