package user_primary_adapter

import (
	user_service "go-hexagonal-user-management/core/services/user"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T){
	// Setup
	e := echo.New()
	u := user_service.NewUserService(e)
	fake := user_secondary_adapter.NewFakeUserRepository()
	e = Delete(u, fake)
	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/1")
	
	t.Run("Success", func(t *testing.T){
		// ハンドラーを呼び出す
		e.Router().Find(http.MethodDelete, "/user/1", c) // 変数にすること
		handler := c.Handler()
		assert.NoError(t, handler(c))
		// レスポンスを検証する
		assert.Equal(t, http.StatusOK, rec.Code)
	})	
}