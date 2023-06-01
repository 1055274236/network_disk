package httptestutils

import (
	"net/http"
	"net/http/httptest"
)

func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	return values[1:]
}

func Get(url string, query map[string]string) *httptest.ResponseRecorder {
	url += "?" + ParseToStr(query)
	// 构造get请求
	req := httptest.NewRequest(http.MethodGet, url, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	Engine.ServeHTTP(w, req)
	return w
}
