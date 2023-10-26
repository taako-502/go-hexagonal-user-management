package user_primary_adapter

import (
	"bytes"
	user_service "go-sample-api/application/services/user"
	primary_adapter "go-sample-api/primary/adapter"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T){
	// Setup
	e := primary_adapter.SetupEchoForTest()
	u := user_service.NewUserService(e)
	fake := user_secondary_adapter.NewFakeUserRepository()
	e = Create(u, fake)
	reqJSON := `{"username": "test", "email": "test@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBufferString(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) 
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")
	
	t.Run("Success", func(t *testing.T){
		// ハンドラーを呼び出す
		e.Router().Find(http.MethodPost, "/user", c)
		handler := c.Handler()
		assert.NoError(t, handler(c))
		// レスポンスを検証する
		assert.Equal(t, http.StatusOK, rec.Code)
	})	
}