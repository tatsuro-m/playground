package user

import (
	"bytes"
	"encoding/json"
	"gin-gorm-tutorial/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

type req struct {
	body map[string]interface{}
}
type expected struct {
	code int
	body map[string]interface{}
}

func TestController_Create(t *testing.T) {
	tests := []struct {
		name     string
		req      req
		expected expected
	}{
		{
			name: "正常に作成されること",
			req: req{map[string]interface{}{
				"first_name": "test_first1",
				"last_name":  "test_last1",
			}},
			expected: expected{
				code: http.StatusCreated,
				body: map[string]interface{}{
					"first_name": "test_first1",
					"last_name":  "test_last1",
				},
			},
		},
		{
			name: "余計なフィールドがあっても作成されること",
			req: req{map[string]interface{}{
				"first_name":       "test_first1",
				"last_name":        "test_last1",
				"not_exists_field": "hogehoge",
			}},
			expected: expected{
				code: http.StatusCreated,
				body: map[string]interface{}{
					"first_name": "test_first1",
					"last_name":  "test_last1",
				},
			},
		},
		{
			name: "first_name が欠けていたらエラーになること",
			req: req{map[string]interface{}{
				"last_name": "test_last1",
			}},
			expected: expected{
				code: http.StatusBadRequest,
				body: map[string]interface{}{},
			},
		},
		{
			name: "last_name が欠けていたらエラーになること",
			req: req{map[string]interface{}{
				"first_name": "test_first1",
			}},
			expected: expected{
				code: http.StatusBadRequest,
				body: map[string]interface{}{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange ---
			db.Init()

			reqBody, _ := json.Marshal(tt.req.body)
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
			assert.Equal(t, tt.expected.code, res.Code)

			var resMap map[string]interface{}
			_ = json.Unmarshal(res.Body.Bytes(), &resMap)

			if resMap == nil {
				assert.Nil(t, resMap)
			} else {
				for k, v := range tt.expected.body {
					assert.Equal(t, v, resMap[k])
				}
				assert.Contains(t, resMap, "id")
				assert.Contains(t, resMap, "created_at")
				assert.Contains(t, resMap, "updated_at")
			}
		})
	}
}
