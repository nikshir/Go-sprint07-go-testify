package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandleWhenStatusOk(t *testing.T) {
	// запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверяем на StatusOk(код ответа 200)
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Неверный код статуса")
	//Проверяем что тело ответа не пустое.
	assert.NotEmpty(t, responseRecorder.Body, "Тело ответа не должно быть пустым")
}

func TestMainHandleWhenCityNotSupported(t *testing.T) {
	// запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=10&city=somethingStrange", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Ожидаемый ответ сервера.
	expected := "wrong city value"

	//Проверяем на статус BadRequest(код статуса 400)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest, "Неверный код статуса")
	//Проверяем тело ответа на наличие ожидаемого ответа сервера.
	assert.Contains(t, responseRecorder.Body.String(), expected)
}

func TestMainHandleWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Получаем количество кафе в ответе.
	outputCount := len(strings.Split(responseRecorder.Body.String(), ","))

	//Проверяем на StatusOk(код ответа 200)
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Неверный код статуса")
	//Сравниваем количество выведенных кафе с максимальным количеством
	assert.Equal(t, outputCount, totalCount, "Выведено неверное количество кафе")
}
