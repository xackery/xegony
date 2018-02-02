package web

import (
	"html/template"
	"io/ioutil"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/box"
)

//Template wraps template objects
type Template struct {
	template *template.Template
}

func loadStandardTemplate(oldTemplate *template.Template) (tmp *template.Template, err error) {
	tmp = oldTemplate
	tmp, err = loadTemplate(tmp, "navmenu", "navmenu.tpl")
	if err != nil {
		return
	}
	tmp, err = loadTemplate(tmp, "navigation", "navigation.tpl")
	if err != nil {
		return
	}
	tmp, err = loadTemplate(tmp, "header", "header.tpl")
	if err != nil {
		return
	}
	tmp, err = loadTemplate(tmp, "root", "root.tpl")
	if err != nil {
		return
	}
	return
}

func getTemplate(key string) (value *template.Template) {

	if tmp, ok := templates[key]; ok {
		value = tmp.template
		return
	}
	return
}

func setTemplate(key string, value *template.Template) {
	tmp := &Template{
		template: value,
	}
	templates[key] = tmp
}

func loadTemplate(oldTemplate *template.Template, key string, path string) (tmp *template.Template, err error) {
	var bData []byte

	//First, we try to use local file
	if bData, err = ioutil.ReadFile("template/" + path); err != nil {
		//fallback to opening local directory
		if bData, err = box.ReadFile("template/" + path); err != nil {
			err = errors.Wrap(err, "Could not find template "+path)
			return
		}
	}

	funcMap := template.FuncMap{
		"comma":       commaFormat,
		"iszonelevel": isZoneLevel,
		"time":        timeFormat,
	}

	if oldTemplate == nil {
		tmp, err = template.New(key).Funcs(funcMap).Parse(string(bData))
	} else {
		tmp, err = oldTemplate.New(key).Funcs(funcMap).Parse(string(bData))
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
