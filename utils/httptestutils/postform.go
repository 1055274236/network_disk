package httptestutils

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func ParseToUrlVaule(params map[string]string) url.Values {
	content := url.Values{}
	for key, value := range params {
		content.Set(key, value)
	}
	return content
}

// PostForm 根据特定请求url和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(url string, params map[string]string) *httptest.ResponseRecorder {

	reqBody := ParseToUrlVaule(params)

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(reqBody.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	Engine.ServeHTTP(w, req)

	return w
}
