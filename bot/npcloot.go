package bot

import (
	"fmt"
)

func (a *Bot) CreateNpcDropsCache() (err error) {
	npcs, err := a.npcRepo.List()
	if err != nil {
		return
	}
	for _, npc := range npcs {
		fmt.Println(npc)
	}
	return
}
