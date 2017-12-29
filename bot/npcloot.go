package bot

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Bot) NpcLootStatus(w http.ResponseWriter, r *http.Request) {
	type Content struct {
		Message string
	}
	content := &Content{
		Message: "Idle",
	}
	writeData(w, r, content, http.StatusOK)
	return
}

func (a *Bot) CreateNpcDropsCache() (err error) {
	start := time.Now()
	err = a.npcLootRepo.Truncate()
	if err != nil {
		err = errors.Wrap(err, "Failed to truncate cache")
		return
	}
	npcs, err := a.npcRepo.List()
	if err != nil {
		return
	}

	entryCount := 0
	var loottable *model.LootTable
	var lootdrop *model.LootDrop
	var item *model.Item
	npcLoot := &model.NpcLoot{}
	for _, npc := range npcs {
		if npc.LoottableId == 0 {
			continue
		}
		loottable, err = a.lootTableRepo.Get(npc.LoottableId)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Failed to get lootTable %d", npc.LoottableId))
			fmt.Println(err)
			continue
		}
		for _, loottableEntry := range loottable.Entries {
			lootdrop, err = a.lootDropRepo.Get(loottableEntry.LootdropId)
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("Failed to get lootDrop %d", loottableEntry.LootdropId))
				fmt.Println(err)
				continue
			}
			for _, lootdropEntry := range lootdrop.Entries {
				item, err = a.itemRepo.Get(lootdropEntry.ItemId)
				if err != nil {
					err = errors.Wrap(err, fmt.Sprintf("Failed to get item %d", lootdropEntry.ItemId))
					fmt.Println(err)
					continue
				}
				//fmt.Println(len(npcs), "npcs fetched, doing ", npc, "and got loottable", loottable, "and lootdrop", lootdrop, "with item", item)
				npcLoot.NpcId = npc.Id
				npcLoot.ItemId = item.Id
				err = a.npcLootRepo.Create(npcLoot)
				if err != nil {
					err = errors.Wrap(err, fmt.Sprintf("Failed to create cache entry npc_id: %d item_id: %d", npc.Id, item.Id))
					fmt.Println(err)
					continue
				}
				entryCount++
			}
		}
		fmt.Println(npc.Id, time.Since(start))
	}
	fmt.Println("Created", entryCount, "entries for ", len(npcs), "npcs in", time.Since(start))
	return
}