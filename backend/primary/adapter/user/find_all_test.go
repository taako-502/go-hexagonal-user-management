package user_primary_adapter

import (
	"encoding/json"
	"go-sample-api/application/domain"
	user_service "go-sample-api/application/services/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T){
	// Setup
	e := echo.New()
	u := user_service.NewUserService(e)
	fake := user_secondary_adapter.NewFakeUserRepository()
	e = FindAll(u, fake)
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")
	
	t.Run("Success", func(t *testing.T){
		// ハンドラーを呼び出す
		e.Router().Find(http.MethodGet, "/user", c)
		handler := c.Handler()
		assert.NoError(t, handler(c))
		// レスポンスを検証する
		assert.Equal(t, http.StatusOK, rec.Code)
		var users []domain.User
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &users))
		assert.EqualValues(t, 1, len(users))
		assert.Equal(t, "test", users[0].Username)
		assert.Equal(t, "test@test.com", users[0].Email)
	})	



}