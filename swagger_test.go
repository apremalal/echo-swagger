package echoswagger

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func TestWrapHandler(t *testing.T) {

	router := echo.New()

	router.GET("/*", WrapHandler)

	w5 := performRequest("GET", "/", router)
	assert.Equal(t, 200, w5.Code)

	w1 := performRequest("GET", "/index.html", router)
	assert.Equal(t, 200, w1.Code)

	w2 := performRequest("GET", "/doc.json", router)
	assert.Equal(t, 200, w2.Code)

	w3 := performRequest("GET", "/favicon-16x16.png", router)
	assert.Equal(t, 200, w3.Code)

	w4 := performRequest("GET", "/notfound", router)
	assert.Equal(t, 404, w4.Code)

}

func performRequest(method, target string, e *echo.Echo) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, r)
	return w
}
