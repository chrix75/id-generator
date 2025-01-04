package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetID(t *testing.T) {
	router := setupRouter(0)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/id", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var value IDValue
	err := json.Unmarshal(w.Body.Bytes(), &value)
	assert.NoError(t, err)

	assert.Greater(t, value.ID, uint64(0))
}
