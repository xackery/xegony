package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/box"
	"github.com/xackery/xegony/cases"
)

var (
	templateLock = sync.RWMutex{}
	templates    = make(map[string]*template.Template)
)

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

func loadTemplate(oldTemplate *template.Template, key string, path string) (tmp *template.Template, err error) {
	templateLock.Lock()
	defer templateLock.Unlock()

	useCache := cases.GetConfigValue("webCacheTemplate")
	var ok bool
	if useCache == "1" {
		tmp, ok = templates[key]
		if ok {
			return
		}
	}
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
		"unescapeJS":  unescapeJS,
	}

	if oldTemplate == nil {
		tmp, err = template.New(key).Funcs(funcMap).Parse(string(bData))
		if err != nil {
			return
		}
	} else {
		tmp, err = oldTemplate.New(key).Funcs(funcMap).Parse(string(bData))
		if err != nil {
			return
		}
	}
	if useCache == "1" {
		templates[key] = tmp
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
