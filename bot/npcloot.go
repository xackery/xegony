package bot

/*
import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)


func (a *Bot) npcLootStatus(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	type Content struct {
		Message     string
		Status      string
		Runtime     string
		LastStarted time.Time
	}

	var bot *Status
	if bot, err = a.getStatus("npcloot"); err != nil {
				return
	}

	content := &Content{
		Message:     fmt.Sprintf("Bot is %s, last started %s", bot.State, bot.StartTime),
		Status:      bot.State,
		Runtime:     fmt.Sprintf("%.2f minutes", bot.getRuntime().Minutes()),
		LastStarted: bot.StartTime,
	}

		return
}

func (a *Bot) npcLootCreate(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	type Content struct {
		Message string
	}

	content := &Content{
		Message: "Starting bot to process npc loot",
	}

	if err = a.startBot("npcloot"); err != nil {
				return
	}

	go a.allNpcLootCreate()

		return
}

func (a *Bot) allNpcLootCreate() {

	for i := int64(0); i < 999; i++ {
		err = a.CreateNpcDropsCache(i)
		if err != nil {
			fmt.Println("Failed to create npc entry:", err.Error())
			continue
		}
	}
	a.endBot("npcloot")
}

// CreateNpcDropsCache is a shortcut function to prepare cache
func (a *Bot) CreateNpcDropsCache(zoneID int64) (err error) {
	start := time.Now()
	//err = a.npcLootRepo.Truncate()
	if err != nil {
		err = errors.Wrap(err, "Failed to truncate cache")
		return
	}
	npcs, err := a.npcRepo.ListByZone(zoneID, nil)
	if err != nil {
		return
	}

	entryCount := 0
	var loottable *model.LootTable
	var lootdrop *model.LootDrop
	var item *model.Item
	npcLoot := &model.NpcLoot{}
	for _, npc := range npcs {
		if npc.LoottableID == 0 {
			continue
		}
		loottable, err = a.lootTableRepo.Get(npc.LoottableID, nil)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Failed to get lootTable %d", npc.LoottableID))
			fmt.Println(err)
			continue
		}
		for _, loottableEntry := range loottable.Entries {
			lootdrop, err = a.lootDropRepo.Get(loottableEntry.LootdropID, nil)
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("Failed to get lootDrop %d", loottableEntry.LootdropID))
				fmt.Println(err)
				continue
			}
			for _, lootdropEntry := range lootdrop.Entries {
				item, err = a.itemRepo.Get(lootdropEntry.ItemID, nil)
				if err != nil {
					err = errors.Wrap(err, fmt.Sprintf("Failed to get item %d", lootdropEntry.ItemID))
					fmt.Println(err)
					continue
				}
				//fmt.Println(len(npcs), "npcs fetched, doing ", npc, "and got loottable", loottable, "and lootdrop", lootdrop, "with item", item)
				npcLoot.NpcID = npc.ID
				npcLoot.ItemID = item.ID
				err = a.npcLootRepo.Create(npcLoot, nil)
				if err != nil {
					err = errors.Wrap(err, fmt.Sprintf("Failed to create cache entry npc_id: %d item_id: %d", npc.ID, item.ID))
					fmt.Println(err)
					continue
				}
				entryCount++
				if entryCount%50 == 0 {
					fmt.Println(npc.ID, time.Since(start))
				}
			}
		}
	}
	fmt.Println("Created", entryCount, "entries for ", len(npcs), "npcs in", time.Since(start))
	return
}
*/
