package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	ctx         context.Context
	httpRequest *http.Request
	Header      *http.Header
	Client      *http.Client
}

type Response struct {
	Resp       *http.Response
	content    []byte
	text       string
	req        *Request
	StatusCode int
}

type Header map[string]string
type Params map[string]string
type Datas map[string]string
type Jsons interface{}
type Auth []string
type Timeout time.Duration

func req() *Request {
	req := new(Request)

	req.httpRequest = &http.Request{
		Method:     "GET",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	req.Header = &req.httpRequest.Header
	req.Client = &http.Client{
		Timeout: time.Duration(15 * time.Second),
	}
	return req
}

func (req *Request) SetTimeout(n time.Duration) {
	req.Client.Timeout = time.Duration(n * time.Second)
}

func Requests(method string, origurl string, args ...interface{}) (*Response, error) {
	switch strings.ToUpper(method) {
	case "GET":
		return Get(origurl, args...)
	case "POST":
		return PostJson(origurl, args...)
	case "DELETE":
		return Delete(origurl, args...)
	case "PUT":
		return Put(origurl, args...)
	default:
		return Get(origurl, args...)
	}
}

func Get(origurl string, args ...interface{}) (*Response, error) {
	return req().Get(origurl, args...)
}
func (req *Request) Get(origurl string, args ...interface{}) (*Response, error) {
	req.httpRequest.Method = "GET"
	// 解析请求参数
	params := []map[string]string{}
	for _, arg := range args {
		switch a := arg.(type) {
		case Timeout:
			req.SetTimeout(time.Duration(a))
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case Params:
			params = append(params, a)
		case Auth:
			req.httpRequest.SetBasicAuth(a[0], a[1])
		}
	}
	// 将参数绑定到url
	disturl, _ := buildURLParams(origurl, params...)
	URL, err := url.Parse(disturl)
	if err != nil {
		return nil, err
	}

	// 发送请求
	req.httpRequest.URL = URL
	res, err := req.Client.Do(req.httpRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 获取结果
	resp := &Response{Resp: res, req: req}
	resp.Content()
	defer res.Body.Close()
	return resp, nil
}

func PostForm(origurl string, args ...interface{}) (*Response, error) {
	return req().PostForm(origurl, args...)
}
func (req *Request) PostForm(origurl string, args ...interface{}) (*Response, error) {
	req.httpRequest.Method = "POST"
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	params := []map[string]string{}
	datas := []map[string]string{}
	for _, arg := range args {
		switch a := arg.(type) {
		case Timeout:
			req.SetTimeout(time.Duration(a))
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case Params:
			params = append(params, a)
		case Datas:
			datas = append(datas, a)
		case Auth:
			req.httpRequest.SetBasicAuth(a[0], a[1])
		}
	}

	disturl, _ := buildURLParams(origurl, params...)
	Forms := url.Values{}
	for _, data := range datas {
		for key, value := range data {
			Forms.Add(key, value)
		}
	}
	data := Forms.Encode()
	req.httpRequest.Body = ioutil.NopCloser(strings.NewReader(data))
	req.httpRequest.ContentLength = int64(len(data))

	URL, err := url.Parse(disturl)
	if err != nil {
		return nil, err
	}
	req.httpRequest.URL = URL
	res, err := req.Client.Do(req.httpRequest)
	req.httpRequest.Body = nil
	req.httpRequest.GetBody = nil
	req.httpRequest.ContentLength = 0
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp := &Response{Resp: res, req: req}
	resp.Content()
	defer res.Body.Close()
	return resp, nil
}

func PostJson(origurl string, args ...interface{}) (*Response, error) {
	return req().PostJson(origurl, args...)
}
func (req *Request) PostJson(origurl string, args ...interface{}) (*Response, error) {
	req.httpRequest.Method = "POST"
	req.Header.Set("Content-Type", "application/json")
	for _, arg := range args {
		switch a := arg.(type) {
		case Timeout:
			req.SetTimeout(time.Duration(a))
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case string:
			req.httpRequest.Body = ioutil.NopCloser(strings.NewReader(arg.(string)))
		case Auth:
			req.httpRequest.SetBasicAuth(a[0], a[1])

		case Jsons:
			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(a)
			if err != nil {
				return nil, err
			}
			req.httpRequest.Body = ioutil.NopCloser(b)
		}
	}

	URL, err := url.Parse(origurl)
	if err != nil {
		return nil, err
	}

	req.httpRequest.URL = URL
	res, err := req.Client.Do(req.httpRequest)
	req.httpRequest.Body = nil
	req.httpRequest.GetBody = nil
	req.httpRequest.ContentLength = 0
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &Response{Resp: res, req: req}
	resp.Content()
	defer res.Body.Close()
	return resp, nil
}

func Delete(origurl string, args ...interface{}) (*Response, error) {
	return req().Delete(origurl, args...)
}
func (req *Request) Delete(origurl string, args ...interface{}) (*Response, error) {
	req.httpRequest.Method = "DELETE"
	req.Header.Set("Content-Type", "application/json")
	for _, arg := range args {
		switch a := arg.(type) {
		case Timeout:
			req.SetTimeout(time.Duration(a))
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case Auth:
			req.httpRequest.SetBasicAuth(a[0], a[1])
		}
	}
	// 发送请求
	URL, err := url.Parse(origurl)
	if err != nil {
		return nil, err
	}
	req.httpRequest.URL = URL
	res, err := req.Client.Do(req.httpRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 获取结果
	resp := &Response{Resp: res, req: req}
	resp.Content()
	defer res.Body.Close()
	return resp, nil
}

func Put(origurl string, args ...interface{}) (*Response, error) {
	request := req()
	request.httpRequest.Method = "PUT"
	return request.PutOrPatch(origurl, args...)
}
func Patch(origurl string, args ...interface{}) (*Response, error) {
	request := req()
	request.httpRequest.Method = "PATCH"
	return request.PutOrPatch(origurl, args...)
}
func (req *Request) PutOrPatch(origurl string, args ...interface{}) (*Response, error) {
	req.Header.Set("Content-Type", "application/json")
	for _, arg := range args {
		switch a := arg.(type) {
		case Timeout:
			req.SetTimeout(time.Duration(a))
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case Auth:
			req.httpRequest.SetBasicAuth(a[0], a[1])
		case Jsons:
			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(a)
			if err != nil {
				return nil, err
			}
			req.httpRequest.Body = ioutil.NopCloser(b)
		}
	}

	// 发送请求
	URL, err := url.Parse(origurl)
	if err != nil {
		return nil, err
	}
	req.httpRequest.URL = URL
	res, err := req.Client.Do(req.httpRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 获取结果
	resp := &Response{Resp: res, req: req}
	resp.Content()
	defer res.Body.Close()
	return resp, nil
}

// get请求将参数绑定到url中
func buildURLParams(userURL string, params ...map[string]string) (string, error) {
	parsedURL, err := url.Parse(userURL)
	if err != nil {
		return "", err
	}

	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return "", nil
	}

	for _, param := range params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}
	if len(parsedQuery) > 0 {
		return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?"), nil
	}
	return strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), nil
}

func (resp *Response) Content() []byte {
	var err error
	var Body = resp.Resp.Body
	if resp.Resp.Header.Get("Content-Encoding") == "gzip" && resp.req.Header.Get("Accept-Encoding") != "" {
		reader, err := gzip.NewReader(Body)
		if err != nil {
			return nil
		}
		Body = reader
	}

	resp.content, err = ioutil.ReadAll(Body)
	if err != nil {
		return nil
	}
	resp.StatusCode = resp.Resp.StatusCode
	return resp.content
}

func (resp *Response) Text() string {
	if resp.content == nil {
		resp.Content()
	}
	resp.text = string(resp.content)
	return resp.text
}

func (resp *Response) Json(v interface{}) error {
	if resp.content == nil {
		resp.Content()
	}
	return json.Unmarshal(resp.content, v)
}
