package post_test

import (
	"bytes"
	"encoding/json"
	"gin-gorm-tutorial/server"
	test_helper "gin-gorm-tutorial/test-helper"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_Index(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	u := test_helper.InsertUser(t, 1)[0]
	test_helper.InsertPost(t, 5, u)

	router := server.Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	router.ServeHTTP(w, req)

	var reqBody []map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &reqBody)
	assert.Len(t, reqBody, 5)
	for i, b := range reqBody {
		assert.Equal(t, "title"+strconv.Itoa(i), b["title"])
		assert.Equal(t, "content"+strconv.Itoa(i), b["content"])
	}
}

func TestController_Create(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	router := server.Router()
	w := httptest.NewRecorder()

	u := test_helper.InsertUser(t, 1)[0]

	body := map[string]interface{}{"title": "title1", "content": "content1", "user_id": u.ID, "User": u}
	reqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/posts", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resBody map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &resBody)
	assert.Equal(t, "title1", resBody["title"])
	assert.Equal(t, "content1", resBody["content"])
}
