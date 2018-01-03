package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	spawnGroupID, err := getIntVar(r, "spawnGroupID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
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
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawn, err := a.spawnRepo.Get(spawnGroupID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	fmt.Println("Spawn", spawn)

	content := Content{
		Site:        site,
		SpawnEntrys: spawnEntrys,
		Spawn:       spawn,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawnentry/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spawnentry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		SpawnEntry *model.SpawnEntry
		Npc        *model.Npc
	}

	spawnGroupID, err := getIntVar(r, "spawnGroupID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	if strings.ToLower(getVar(r, "npcID")) == "details" {
		a.getSpawn(w, r)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	spawnEntry, _, err := a.spawnEntryRepo.Get(spawnGroupID, npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npc, err := a.npcRepo.Get(npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error for npc")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	site := a.newSite(r)
	site.Page = "spawnentry"
	site.Title = "spawnentry"
	site.Section = "spawnentry"

	content := Content{
		Site:       site,
		SpawnEntry: spawnEntry,
		Npc:        npc,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawnentry/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spawnentry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
