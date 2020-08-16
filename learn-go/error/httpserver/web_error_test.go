package main

import (
	"os"
	"fmt"
	"errors"
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
	"net/http/httptest"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}


func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("Internal Server Error")
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct{
	han Handler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, test := range tests {
		fun := errWrapper(test.han)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		fun(response, request)

		verifyResponse(response.Result(), test.code, test.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, test := range tests {
		fun := errWrapper(test.han)
		server := httptest.NewServer(http.HandlerFunc(fun))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, test.code, test.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("测试未通过, 期望值: (%d, %s), 实际值: (%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	} else {
		t.Logf("测试通过, 期望值: (%d, %s), 实际值: (%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}