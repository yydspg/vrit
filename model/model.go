package model

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Meta struct {
	group string
	mh    *MetaHeader
}

type MetaHeader struct {
	accept          string
	accept_encoding string
	accept_language string
	connection      string
	user_agent      string
	cookie          string
	host            string
}
type VritRequest struct {
	mh      *MetaHeader
	uri     string
	host    string
	method  string
	payload string
}

type RequestGroup struct {
	reqs []*VritRequest
}
type Vrit struct {
	m *Meta
	r *RequestGroup
}

func NewVrit(json map[string]interface{}) (*Vrit, error) {
	t := &Vrit{}
	if json == nil {
		return nil, errors.New("[ERROR]:build vrit error,because json empty")
	}
	if v, ok := json["meta"]; ok {
		p, err := NewMeta(v.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		t.m = p
	} else {
		return nil, errors.New("[ERROR]:meta empty")
	}
	if v, ok := json["data"]; ok  {
		p, err := NewRequestGroup(t.m.mh, v.([]interface{}))
		if err != nil {
			return nil, err
		}
		t.r = p
	} else {
		return nil, errors.New("[ERROR]:data empty")
	}
	return t, nil
}

func NewRequestGroup(mh *MetaHeader, json []interface{}) (*RequestGroup, error) {
	if json == nil {
		return nil, errors.New("[ERROR]:build request group error,because json empty")
	}
	t := []*VritRequest{}
	for _, v := range json {
		tmp, err := NewVritReuqest(mh, v.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		t = append(t, tmp)
	}
	a := &RequestGroup{}
	a.reqs = t
	return a, nil
}

func (h *VritRequest) NewHttpReuqest() (*http.Request, error) {
	req, err := http.NewRequest(strings.ToUpper(h.method), buildURL(h.host, h.uri), nil)
	if err != nil {
		fmt.Println("[ERROR]:fail to build http request")
		return nil, errors.New("[ERROR]:fail to build http request")
	}
	return req, nil
}

func buildURL(host, uri string) string {
	host = strings.TrimSuffix(host, "/")
	uri = strings.TrimPrefix(uri, "/")
	return host + "/" + uri
}
func NewVritReuqest(mh *MetaHeader, json map[string]interface{}) (*VritRequest, error) {
	h := &VritRequest{}
	if mh == nil {
		return nil, errors.New("[ERROR]:meta header empty")
	}
	if json == nil {
		return nil, errors.New("[ERROR]:build http header error,because json empty")
	}
	method := ""
	if v, ok := json["method"]; ok && len(json["method"].(string)) != 0 {
		h.method = v.(string)
		method = h.method
	} else {
		return nil, errors.New("[ERROR]:http method invalid")
	}
	if method == "post" || method == "put" {
		if v, ok := json["payload"]; ok && len(json["payload"].(string)) != 0 {
			h.payload = v.(string)
		} else {
			fmt.Println("[WARNING]:post or put method payload empty")
		}
	}
	if v, ok := json["uri"]; ok && len(json["uri"].(string)) != 0 {
		h.uri = v.(string)
	} else {
		return nil, errors.New("[ERROR]:http uri invalid")
	}
	h.host = mh.host
	return h, nil
}

func NewMeta(json map[string]interface{}) (*Meta, error) {
	m := &Meta{}
	if v, ok := json["group"]; ok {
		m.group = v.(string)
	} else {
		m.group = "vrit"
		fmt.Println("[HINT]:use default group name vrit")
	}
	if v, ok := json["header"]; ok {
		mh, err := NewMetaHeader(v.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		m.mh = mh
	} else {
		return nil, errors.New("[ERROR]:build meta header error,because json data empty")
	}
	return m, nil
}

func NewMetaHeader(json map[string]interface{}) (*MetaHeader, error) {
	m := &MetaHeader{}
	if json == nil {
		return nil, errors.New("[ERROR]:build meta header error,because json data empty")
	}
	if v, ok := json["host"]; ok &&len(json["host"].(string)) != 0 {
		m.host = v.(string)
	} else {
		return nil, errors.New("[ERROR]:host empty")
	}
	if v, ok := json["accept"]; ok && len(json["accept"].(string)) != 0 {
		m.accept = v.(string)
	} else {
		m.accept = "application/json"
		fmt.Println("[HINT]:use default request head accept type application/json")
	}
	if v, ok := json["accept_encoding"]; ok && len(json["accept_encoding"].(string)) != 0 {
		m.accept_encoding = v.(string)
	} else {
		m.accept_encoding = "gzip, deflate, br, zstd"
		fmt.Println("[HINT]:use default request head accept encoding type gzip, deflate, br, zstd")
	}
	if v, ok := json["connection"]; ok && len(json["connection"].(string)) != 0 {
		m.connection = v.(string)
	} else {
		m.connection = "keep-alive"
		fmt.Println("[HINT]:use default request head connection type keep-alive")
	}
	if v, ok := json["user_agent"]; ok && len(json["user_agent"].(string)) != 0 {
		m.user_agent = v.(string)
	} else {
		m.user_agent = "Mozilla/5.0"
		fmt.Println("[HINT]:use default request head user_agent type Mozilla/5.0")
	}
	if v, ok := json["cookie"]; ok && len(json["cookie"].(string)) != 0 {
		m.cookie = v.(string)
	} else {
		m.cookie = ""
		fmt.Println("[HINT]:use default request head cookie type empty")
	}
	return m, nil
}
