package rest

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

// templateRequest wraps a template request
type templateRequest struct {
	Ctx      context.Context
	Path     string
	Method   string
	RespChan chan *templateResponse
}

// templateResponse wraps the response to a template
type templateResponse struct {
	Resp  *template.Template
	Error error
}

// gets a template
func (s *Server) getTemplate(req *templateRequest) (resp *template.Template, err error) {
	if req == nil {
		err = fmt.Errorf("invalid template request (nil)")
		return
	}
	if req.Ctx == nil {
		req.Ctx = context.Background()
	}
	if req.Method == "" {
		req.Method = "GET"
	}
	respChan := make(chan *templateResponse)
	req.RespChan = respChan
	select {
	case <-time.After(3 * time.Second):
	case <-req.Ctx.Done():
	case s.templateChan <- req:
		select {
		case <-time.After(3 * time.Second):
		case <-req.Ctx.Done():
		case respT := <-respChan:
			if respT.Error != nil {
				err = errors.Wrap(respT.Error, "failed to load template")
				return
			}
			resp = respT.Resp
		}
	}
	return
}

// do not call onTemplateRequest. use getTemplate instead. (used by pump)
func (s *Server) onTemplateRequest(req *templateRequest) (resp *template.Template, err error) {
	if req.Ctx == nil {
		req.Ctx = context.Background()
	}
	if len(req.Path) < 3 {
		err = fmt.Errorf("template path too short: %s", req.Path)
		return
	}
	var ok bool
	switch req.Method {
	case "GET":
		resp, ok = s.templates[req.Path]
		if ok {
			//return
		}
		resp, err = s.onLoadTemplate(nil, "body", req.Path)
		resp, err = s.onLoadDefaultTemplates(resp)
		s.templates[req.Path] = resp
	default:
		err = fmt.Errorf("invalid request method %s", req.Method)
		return
	}
	return
}

func (s *Server) onLoadDefaultTemplates(req *template.Template) (resp *template.Template, err error) {
	resp = req
	names := map[int]string{0: "sidebar", 1: "header", 2: "root"}
	for i := 0; i < len(names); i++ {
		resp, err = s.onLoadTemplate(resp, names[i], names[i]+".tpl")
		if err != nil {
			err = errors.Wrapf(err, "failed to load %s", names[i])
			return
		}
	}
	return
}

// called exclusively by onTemplateRequest
func (s *Server) onLoadTemplate(root *template.Template, name string, path string) (resp *template.Template, err error) {

	var data []byte
	data, err = ioutil.ReadFile("template/" + path)
	if err != nil {
		err = errors.Wrap(err, "failed to read template")
		//TODO: data, err = box.ReadFile("template/" + path)
		return
	}

	funcMap := template.FuncMap{
		"comma":       commaFormat,
		"iszonelevel": isZoneLevel,
		"time":        timeFormat,
		"unescapeJS":  unescapeJS,
	}
	if root != nil {
		fmt.Println("loading", name)
		resp, err = root.New(name).Funcs(funcMap).Parse(string(data))
		if err != nil {
			return
		}
		return
	}
	resp, err = template.New(name).Funcs(funcMap).Parse(string(data))
	fmt.Println("loading root", name)
	if err != nil {
		return
	}
	return
}

//comma will take an integer and put commas, e.g. 1234 => 1,234
func commaFormat(v int64) string {
	return humanize.Comma(v)
}

func timeFormat(v time.Time) string {
	return humanize.Time(v)
}

func isZoneLevel(level int64, levels int64) bool {
	return ((levels & level) == level)
}

func unescapeJS(v interface{}) template.JS {
	return template.JS(fmt.Sprint(v))
}
