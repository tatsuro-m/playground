package user

import (
	"bytes"
	"encoding/json"
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestController_Create(t *testing.T) {
	// Arrange ---
	db.Init()

	mapData := map[string]interface{}{
		"first_name": "first_name",
		"last_name":  "last_name",
	}

	reqBody, _ := json.Marshal(mapData)
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/api/v1/users",
		bytes.NewBuffer(reqBody),
	)

	// Act ---
	var ctrl Controller
	ctrl.Create(c)

	// Assert ---
	assert.Equal(t, http.StatusCreated, res.Code)
	var user entity.User
	_ = json.Unmarshal(res.Body.Bytes(), &user)
	assert.Equal(t, "last_name", user.LastName)
}
