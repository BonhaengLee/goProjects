package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIndexPathHandler(t *testing.T) { // Test 접두사와 인자는 필수 형식
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req) // >> mux로 받아와서 ServeHTTP에 주어야 타겟에 맞춰 핸들러에 분배해준다.

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body) // body의 버퍼값을 전부 읽어옴
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) { // Test 접두사와 인자는 필수 형식
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // Issue. mux를 제대로 사용하지 못해서 타겟 적용이 안되므로 /에 GET을 했는데도 PASS됨

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req) // >> mux로 받아와서 ServeHTTP에 주어야 타겟에 맞춰 핸들러에 분배해준다.

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body) // body의 버퍼값을 전부 읽어옴
	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) { // Test 접두사와 인자는 필수 형식
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=Coa", nil) // Issue. mux를 제대로 사용하지 못해서 타겟 적용이 안되므로 /에 GET을 했는데도 PASS됨

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req) // >> mux로 받아와서 ServeHTTP에 주어야 타겟에 맞춰 핸들러에 분배해준다.

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body) // body의 버퍼값을 전부 읽어옴
	assert.Equal("Hello Coa!", string(data))
}