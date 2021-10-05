// package main

// import (
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// func TestHealth(t *testing.T) {
// 	app := fiber.New()

// 	app.Get("/", HealthCheck)

// 	req := httptest.NewRequest(http.MethodGet, "/", nil)

// 	resp, err := app.Test(req)
// 	body, _ := io.ReadAll(resp.Body)

// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, 200, resp.StatusCode)
// 	assert.Equal(t, "Alive - V1.0.0", string(body))
// }
