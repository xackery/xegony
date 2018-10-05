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

// TemplateRead reads a provided template
func (s *Server) TemplateRead(ctx context.Context, path string) (resp *template.Template, err error) {
	respMsg, err := s.runQuery(ctx, "TemplateRead", path)
	if err != nil {
		return
	}
	resp, ok := respMsg.(*template.Template)
	if !ok {
		err = fmt.Errorf("failed to convert response to template")
		return
	}
	return
}

// do not call onTemplateRequest. use getTemplate instead. (used by pump)
func (s *Server) onTemplateRead(path string) (resp *template.Template, err error) {
	if len(path) < 3 {
		err = fmt.Errorf("template path too short: %s", path)
		return
	}
	resp, ok := s.templates[path]
	if ok {
		//return
	}
	resp, err = s.onLoadTemplate(nil, "body", path)
	if err != nil {
		err = errors.Wrapf(err, "failed to load template %s", path)
		return
	}
	resp, err = s.onLoadDefaultTemplates(resp)
	if err != nil {
		err = errors.Wrap(err, "failed to load default templates")
		return
	}
	s.templates[path] = resp
	return
}

// called exclusively by onTemplateRead
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
