package api

import (
	"bytes"
	"database/sql"
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetItem(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	item, err := a.itemRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, item, http.StatusOK)
	return
}

const tooltipTemplate = `
<div class="eqitem">
	
    <span class="slot"><span class="item icon-{{.Item.Icon}}"></span></span><br>{{.Item.Name}}<br>
    {{if .Item.Magic}}Magic{{end}}{{if .Item.Notransfer}}No Trade{{end}}<br>
    Class: {{.Item.ClassList}}<br>
    Race: {{.Item.RaceList}}<br>
    {{.Item.SlotList}}<br>
    <br>
    <table cellpadding="0" cellspacing="0" border="0">
    <tbody>
        <tr><td colspan="2">Size:<span style="float:right; padding-left:8px; white-space:nowrap">{{.Item.SizeName}}</span></td><td style="padding-right:8px;"></td>{{if .Item.Hp}}<td style="padding-right:4px" nowrap="">HP:</td><td align="right">{{.Item.Hp}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr><td style="padding-right:4px" nowrap="">Weight:</td><td align="right">{{.Item.Weight}}</td><td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Mana}}<td style="padding-right:4px" nowrap="">Mana:</td><td align="right">{{.Item.Mana}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.Reclevel}}<td style="padding-right:4px" nowrap="">Rec Level:</td><td align="right">{{.Item.Reclevel}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Endur}}<td style="padding-right:4px" nowrap="">Endur:</td><td align="right">{{.Item.Endur}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.Reqlevel}}<td style="padding-right:4px" nowrap="">Req Level:</td><td align="right">{{.Item.Reqlevel}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr><td style="height:4px;font-size:1px">&nbsp;</td></tr>
        <tr>{{if .Item.Astr}}<td style="padding-right:4px" nowrap="">Strength:</td><td align="right">{{.Item.Astr}}</td>{{end}}{{if .Item.Heroic_str}}<td class="hvalue" style="padding-right:8px;">{{.Item.Heroic_str}}</td>{{end}}{{if .Item.Mr}}<td style="padding-right:4px" nowrap="">Magic:</td><td align="right">{{.Item.Mr}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Attack}}<td style="padding-right:4px" nowrap="">Attack:</td><td align="right">{{.Item.Attack}}</td><td class="hvalue" style="padding-right:8px;"></td>{{end}}</tr>
        <tr>{{if .Item.Asta}}<td style="padding-right:4px" nowrap="">Stamina:</td><td align="right">{{.Item.Asta}}</td>{{end}}{{if .Item.Heroic_sta}}<td class="hvalue" style="padding-right:8px;">+{{.Item.Heroic_sta}}</td>{{end}}{{if .Item.Fr}}<td style="padding-right:4px" nowrap="">Fire:</td><td align="right">{{.Item.Fr}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Regen}}<td style="padding-right:4px" nowrap="">HP Regen:</td><td align="right">{{.Item.Regen}}</td><td class="hvalue" style="padding-right:8px;"></td>{{end}}</tr>
        <tr>{{if .Item.Aint}}<td style="padding-right:4px" nowrap="">Intelligence:</td><td align="right">{{.Item.Aint}}</td>{{end}}{{if .Item.Heroic_int}}<td class="hvalue" style="padding-right:8px;">+{{.Item.Heroic_int}}</td>{{end}}{{if .Item.Cr}}<td style="padding-right:4px" nowrap="">Cold:</td><td align="right">{{.Item.Cr}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td>{{if .Item.Manaregen}}<td style="padding-right:4px" nowrap="">Mana Regen:</td><td align="right">{{.Item.Manaregen}}</td><td class="hvalue" style="padding-right:8px;"></td>{{end}}</tr>
        <tr>{{if .Item.Awis}}<td style="padding-right:4px" nowrap="">Wisdom:</td><td align="right">{{.Item.Awis}}</td>{{end}}{{if .Item.Heroic_wis}}<td class="hvalue" style="padding-right:8px;">+{{.Item.Heroic_wis}}</td>{{end}}{{if .Item.Dr}}<td style="padding-right:4px" nowrap="">Disease:</td><td align="right">{{.Item.Dr}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.Aagi}}<td style="padding-right:4px" nowrap="">Agility:</td><td align="right">{{.Item.Aagi}}</td>{{end}}{{if .Item.Heroic_agi}}<td class="hvalue" style="padding-right:8px;">+{{.Item.Heroic_agi}}</td>{{end}}{{if .Item.Pr}}<td style="padding-right:4px" nowrap="">Poison:</td><td align="right">{{.Item.Pr}}</td>{{end}}<td class="hvalue" style="padding-right:8px;"></td></tr>
        <tr>{{if .Item.Adex}}<td style="padding-right:4px" nowrap="">Dexterity:</td><td align="right">{{.Item.Adex}}</td>{{end}}{{if .Item.Heroic_dex}}<td class="hvalue" style="padding-right:8px;">+{{.Item.Heroic_dex}}</td>{{end}}</tr>
        <tr>{{if .Item.Acha}}<td style="padding-right:4px" nowrap="">Charisma:</td><td align="right">{{.Item.Acha}}</td>{{end}}{{if .Item.Heroic_cha}}<td class="hvalue" style="padding-right:8px;">+{{.Item.Heroic_cha}}</td>{{end}}</tr>
    </tbody>
    </table>
    {{/*<div class="augments">
        Slot 1, type 3 (General: Spell Effect): empty<br>
        Slot 2, type 5 (Weapon: Elem Damage): empty<br>
        Slot 3, type 7 (General: Group): empty<br>
        Slot 4, type 9 (General: Dragons Points): empty<br>
    </div>*/}}
    {{/*<div class="effects">
        Effect: <a rel="eq:spell:9616" href="/spell/9616" target="_blank">Sharpshooting VII</a> (Worn)<br>
        Focus Effect: <a rel="eq:spell:42971" href="/spell/42971" target="_blank">Detrimental Duration 26 L110</a><br>
    </div>*/}}

</div>`

func (a *Api) GetItemTooltip(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type ItemTooltip struct {
		Name    string `json:"name"`
		Id      int64  `json:"id"`
		Content string `json:"content"`
	}
	item, err := a.itemRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	tmp, err := template.New("tooltip").Parse(tooltipTemplate)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	var tpl bytes.Buffer
	type TemplateData struct {
		Item *model.Item
	}

	templateData := &TemplateData{
		Item: item,
	}
	err = tmp.Execute(&tpl, templateData)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	itemTooltip := &ItemTooltip{
		Name:    item.Name,
		Id:      item.Id,
		Content: tpl.String(),
	}

	writeData(w, r, itemTooltip, http.StatusOK)
	return
}

func (a *Api) CreateItem(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	item := &model.Item{}
	err = decodeBody(r, item)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.itemRepo.Create(item)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, item, http.StatusCreated)
	return
}

func (a *Api) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.itemRepo.Delete(id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			writeData(w, r, nil, http.StatusNotModified)
			return
		default:
			err = errors.Wrap(err, "Request failed")
			writeError(w, r, err, http.StatusInternalServerError)
		}
		return
	}
	writeData(w, r, nil, http.StatusNoContent)
	return
}

func (a *Api) EditItem(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	item := &model.Item{}
	err = decodeBody(r, item)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.itemRepo.Edit(id, item)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, item, http.StatusOK)
	return
}

func (a *Api) ListItem(w http.ResponseWriter, r *http.Request) {
	pageSize := getIntParam(r, "pageSize")
	pageNumber := getIntParam(r, "pageNumber")

	items, err := a.itemRepo.List(pageSize, pageNumber)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, items, http.StatusOK)
	return
}
