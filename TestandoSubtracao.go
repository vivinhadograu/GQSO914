package main
import (
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestSub(t *testing.T) {
	assert.Equal(t, 1., sub(2, 1))
}
func TestSubHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/sub/2/1",
		nil)
	w := httptest.NewRecorder()
	subHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	dados, _ := io.ReadAll(resp.Body)
	resultado, _ := strconv.ParseFloat(string(dados), 64)
	assert.Equal(t, 1., resultado)
}
func TestSubHandler_ErrorParam1(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/sub/a/1",
		nil)
	w := httptest.NewRecorder()
	subHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
func TestSubHandler_ErrorParam2(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/sub/1/a",
		nil)
	w := httptest.NewRecorder()
	subHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
func TestSubHandler_ErrorSParam(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/sub/1",
		nil)
	w := httptest.NewRecorder()
	subHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}
func TestSubHandler_ErrorParam(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/sub/",
		nil)
	w := httptest.NewRecorder()
	subHandler(w, req)
	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode)
}