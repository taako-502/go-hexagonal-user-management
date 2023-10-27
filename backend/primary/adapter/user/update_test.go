package user_primary_adapter

import (
	"bytes"
	user_service "go-hexagonal-user-management/core/services/user"
	primary_adapter "go-hexagonal-user-management/primary/adapter"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	// Setup
	e := primary_adapter.SetupEchoForTest()
	u := user_service.NewUserService(e)
	fake := user_secondary_adapter.NewFakeUserRepository()
	e = Update(u, fake)
	reqJSON := `{"username": "test", "email": "test@example.com"}`
	req := httptest.NewRequest(http.MethodPatch, "/user/1", bytes.NewBufferString(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/1")

	t.Run("Success", func(t *testing.T) {
		// ハンドラーを呼び出す
		e.Router().Find(http.MethodPatch, "/user/1", c)
		handler := c.Handler()
		assert.NoError(t, handler(c))
		// レスポンスを検証する
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
