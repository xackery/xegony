package bot

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Bot) CreateNpcDropsCache() (err error) {
	npcs, err := a.npcRepo.List()
	if err != nil {
		return
	}
	var loottable *model.LootTable
	var lootdrop *model.LootDrop
	var item *model.Item

	for _, npc := range npcs {
		loottable, err = a.lootTableRepo.Get(npc.LoottableId)
		if err != nil {
			err = errors.Wrap(err, "Failed to get lootTable")
			return
		}
		for _, loottableEntry := range loottable.Entries {
			lootdrop, err = a.lootDropRepo.Get(loottableEntry.LootdropId)
			if err != nil {
				err = errors.Wrap(err, "Failed to get lootDrop")
				return
			}
			for _, lootdropEntry := range lootdrop.Entries {
				item, err = a.itemRepo.Get(lootdropEntry.ItemId)
				if err != nil {
					err = errors.Wrap(err, "Failed to get item")
					return
				}
				fmt.Println(len(npcs), "npcs fetched, doing ", npc, "and got loottable", loottable, "and lootdrop", lootdrop, "with item", item)

				break
			}
			break
		}
		break
	}
	return
}
