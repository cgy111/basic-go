package web

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "hello#world123"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}

func TestNil(t *testing.T) {
	testTypeAssert(nil)

}

func testTypeAssert(c any) {
	claims := c.(*UserClaims)
	println(claims.Uid)
}

func TestUserHandler_Signup(t *testing.T) {
	testCases := []struct {
		name string
	}{}
	req, err := http.NewRequest(http.MethodPost,
		"/users/signup", bytes.NewBuffer([]byte(`
{
	"email":"123sty@qq.com",
	"password":"hello@world123"
}
`)))
	require.NoError(t, err)
	//可以继续使用req
	resp := httptest.NewRecorder()
	resp.Header()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//	这里怎么拿到这个响应
			handler := NewUserHandler(nil, nil)
			ctx := &gin.Context{}
			handler.Signup(ctx)
		})
	}
}
