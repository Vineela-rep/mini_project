package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test(m *testing.M)  {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func testHTTPResponse(t *testing.T,r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool){
	w:=httptest.NewRecorder()

	r.ServeHTTP(w,req)
	if !f(w){
		t.Fail()
	}
}

func TestGetAllUsers(t *testing.T){
	response,_:=http.Get("http://localhost:8080/user-sales/user")
	fmt.Println(response)
	resp,_:=io.ReadAll(response.Body)
	fmt.Println(string(resp))
}


