package web

import (
	"html/template"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/box"
)

type Template struct {
	template *template.Template
}

func (w *Web) loadStandardTemplate(oldTemplate *template.Template) (tmp *template.Template, err error) {
	tmp = oldTemplate
	tmp, err = w.loadTemplate(tmp, "navigation", "navigation.tpl")
	if err != nil {
		return
	}
	tmp, err = w.loadTemplate(tmp, "header", "header.tpl")
	if err != nil {
		return
	}
	tmp, err = w.loadTemplate(tmp, "root", "root.tpl")
	if err != nil {
		return
	}
	return
}

func (w *Web) getTemplate(key string) (value *template.Template) {

	if tmp, ok := w.templates[key]; ok {
		value = tmp.template
		return
	}
	return
}

func (w *Web) setTemplate(key string, value *template.Template) {
	tmp := &Template{
		template: value,
	}
	w.templates[key] = tmp
}

func (w *Web) loadTemplate(oldTemplate *template.Template, key string, path string) (tmp *template.Template, err error) {
	var bData []byte

	//First, we try to use local file
	if bData, err = ioutil.ReadFile("template/" + path); err != nil {
		//fallback to opening local directory
		if bData, err = box.ReadFile("template/" + path); err != nil {
			err = errors.Wrap(err, "Could not find template "+path)
			return
		}
	}

	if oldTemplate == nil {
		tmp, err = template.New(key).Parse(string(bData))
	} else {
		tmp, err = oldTemplate.New(key).Parse(string(bData))
	}
	return
}
