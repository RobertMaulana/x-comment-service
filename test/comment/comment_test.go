package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	Host = "localhost:8080"
)


func TestGetComments(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, err := http.NewRequest("GET", Host + "/orgs/xendit/comments", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	assert.Equal(t, resp.Code, 200)
}

