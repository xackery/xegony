package web

/*
import (
	"net/http"
	"html/template"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) spawnEntryRoutes() (routes []*route) {
	routes = []*route{

		//SpawnEntry
		{
			"ListSpawnEntry",
			"GET",
			"/spawn/{spawnGroupID}",
			a.listSpawnEntry,
		},
		{
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnGroupID}/{npcID}",
			a.getSpawnEntry,
		},
	}
	return
}

func (a *Web) listSpawnEntry(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {


	spawnGroupID, err := getIntVar(r, "spawnGroupID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
				return
	}

	type Content struct {
		Site        site
		Spawn       *model.Spawn
		SpawnEntrys []*model.SpawnEntry
	}

	site := a.newSite(r)
	site.Page = "spawnentry"
	site.Title = "spawnentry"
	site.Section = "spawnentry"

	spawnEntrys, _, err := a.spawnEntryRepo.List(spawnGroupID)
	if err != nil {
				return
	}

	spawn, err := a.spawnRepo.Get(spawnGroupID)
	if err != nil {
				return
	}

	content = Content{
		Site:        site,
		SpawnEntrys: spawnEntrys,
		Spawn:       spawn,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawnentry/list.tpl")
		if err != nil {
						return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
						return
		}

		a.setTemplate("spawnentry", tmp)
	}

		return
}

func (a *Web) getSpawnEntry(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {


	type Content struct {
		Site       site
		SpawnEntry *model.SpawnEntry
		Npc        *model.Npc
	}

	spawnGroupID, err := getIntVar(r, "spawnGroupID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
				return
	}

	if strings.ToLower(getVar(r, "npcID")) == "details" {
		a.getSpawn(w, r)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
				return
	}
	spawnEntry, _, err := a.spawnEntryRepo.Get(spawnGroupID, npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
				return
	}

	npc, err := a.npcRepo.Get(npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error for npc")
				return
	}
	site := a.newSite(r)
	site.Page = "spawnentry"
	site.Title = "spawnentry"
	site.Section = "spawnentry"

	content = Content{
		Site:       site,
		SpawnEntry: spawnEntry,
		Npc:        npc,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawnentry/get.tpl")
		if err != nil {
						return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
						return
		}

		a.setTemplate("spawnentry", tmp)
	}

		return
}
*/
