package main
import (
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestSoma(t *testing.T) {
	assert.Equal(t, 3.8, soma(1.5, 2.3))
}
func TestSomaHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/1.5/2.3",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	dados, _ := io.ReadAll(resp.Body)
	resultado, _ := strconv.ParseFloat(string(dados), 64)
	assert.Equal(t, 3.8, resultado)
}
func TestSomaHandler_ErrorParam1(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/a/1",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
func TestSomaHandler_ErrorParam2(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/1/a",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
func TestSomaHandler_ErrorParam(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/soma",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
func TestSomaHandler_ErrorSParam(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/1.0",
		nil)
	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}