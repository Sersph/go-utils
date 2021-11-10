package util

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var httpUtilInstance *httpUtil

type httpUtil struct {
	CookieJar  *cookiejar.Jar
	HttpClient *http.Client
}

var DefaultHeaders = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4350.7 Safari/537.36",
	"Accept":          "*/*",
	"Accept-Language": "zh-CN,zh;q=0.9",
	"Accept-Encoding": "gzip, deflate, br",
	"Connection":      "keep-alive",
}

func GetHttpUtil() *httpUtil {
	if httpUtilInstance == nil {
		cookieJar, _ := cookiejar.New(nil)
		httpClient := &http.Client{
			Timeout: time.Minute,
			Jar:     cookieJar,
		}
		httpUtilInstance = &httpUtil{
			CookieJar:  cookieJar,
			HttpClient: httpClient,
		}
	}
	return httpUtilInstance
}

func (*httpUtil) SetDefaultHeaders(req *http.Request) {
	if req != nil {
		for key, value := range DefaultHeaders {
			req.Header.Set(key, value)
		}
	}
}

func (*httpUtil) GetRespBody(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}

	var reader io.Reader
	var err error
	contentEncoding := resp.Header.Get("Content-Encoding")
	if "gzip" == contentEncoding {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = resp.Body
	}

	respbytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return respbytes, nil
}
