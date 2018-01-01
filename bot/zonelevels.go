package bot

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Bot) zoneLevelsStatus(w http.ResponseWriter, r *http.Request) {
	type Content struct {
		Message string
	}
	content := &Content{
		Message: "Idle",
	}
	writeData(w, r, content, http.StatusOK)
	return
}

// CreateZoneLevelCache is a shortcut function to prepare cache
func (a *Bot) CreateZoneLevelCache() (err error) {
	minCount := 20 //minimum number of mobs to be considered levelable there
	start := time.Now()
	var query string
	err = a.zoneLevelRepo.Truncate()
	if err != nil {
		err = errors.Wrap(err, "Failed to truncate cache")
		return
	}
	zones, err := a.zoneRepo.List()
	if err != nil {
		return
	}

	var spawns []*model.SpawnEntry
	var npc *model.Npc
	zoneLevel := &model.ZoneLevel{}
	fmt.Println("Processing", len(zones), "zones")
	for _, zone := range zones {

		if zone.ZoneIDNumber == 0 {
			fmt.Println("skipping zone with 0")
			continue
		}

		spawns, query, err = a.spawnEntryRepo.ListByZone(zone.ZoneIDNumber)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Failed to get spawnentry %s", query))
			return
		}
		mobCounter := map[int]int{}

		//fmt.Println("Zone", zone.ZoneIDNumber, "has", len(spawns), "entries")
		for _, spawn := range spawns {

			npc, err = a.npcRepo.Get(spawn.NpcID)
			if err != nil {
				fmt.Println("Invalid NPC provided", spawn.NpcID)
				continue
			}
			//can't kill non-class mobs
			if npc.Class < 1 && npc.Class > 17 {
				continue
			}

			if strings.Contains(npc.SpecialAbilities.String, "1^35") ||
				strings.Contains(npc.SpecialAbilities.String, "1^25") ||
				strings.Contains(npc.SpecialAbilities.String, "1^24") { //invul
				continue
			}

			if npc.Level == 1 { //1
				mobCounter[1]++
			}
			if npc.Level < 6 && npc.Level > 2 { //5
				mobCounter[2]++
			}
			if npc.Level < 11 && npc.Level > 7 { //10
				mobCounter[4]++
			}
			if npc.Level < 16 && npc.Level > 11 { //15
				mobCounter[8]++
			}
			if npc.Level < 21 && npc.Level > 15 { //20
				mobCounter[16]++
			}
			if npc.Level < 26 && npc.Level > 20 { //25
				mobCounter[32]++
			}
			if npc.Level < 31 && npc.Level > 25 { //30
				mobCounter[64]++
			}
			if npc.Level < 36 && npc.Level > 25 { //35
				mobCounter[128]++
			}
			if npc.Level < 41 && npc.Level > 35 { //40
				mobCounter[256]++
			}
			if npc.Level < 46 && npc.Level > 40 { //45
				mobCounter[512]++
			}
			if npc.Level < 51 && npc.Level > 45 { //50
				mobCounter[1024]++
			}
			if npc.Level < 56 && npc.Level > 50 { //55
				mobCounter[2048]++
			}
			if npc.Level < 61 && npc.Level > 55 { //60
				mobCounter[4096]++
			}
			if npc.Level < 66 && npc.Level > 60 { //65
				mobCounter[8192]++
			}
			if npc.Level < 71 && npc.Level > 65 { //70
				mobCounter[16384]++
			}
			if npc.Level < 76 && npc.Level > 70 { //75
				mobCounter[32768]++
			}
			if npc.Level < 81 && npc.Level > 75 { //80
				mobCounter[65536]++
			}
		}

		zoneLevel.ZoneID = zone.ZoneIDNumber
		zoneLevel.Levels = 0

		if mobCounter[1] >= minCount {
			zoneLevel.Levels |= 1
		}
		if mobCounter[2] >= minCount {
			zoneLevel.Levels |= 2
		}
		if mobCounter[4] >= minCount {
			zoneLevel.Levels |= 4
		}
		if mobCounter[8] >= minCount {
			zoneLevel.Levels |= 8
		}
		if mobCounter[16] >= minCount {
			zoneLevel.Levels |= 16
		}
		if mobCounter[32] >= minCount {
			zoneLevel.Levels |= 32
		}
		if mobCounter[64] >= minCount {
			zoneLevel.Levels |= 64
		}
		if mobCounter[128] >= minCount {
			zoneLevel.Levels |= 128
		}
		if mobCounter[256] >= minCount {
			zoneLevel.Levels |= 256
		}
		if mobCounter[512] >= minCount {
			zoneLevel.Levels |= 512
		}
		if mobCounter[1024] >= minCount {
			zoneLevel.Levels |= 1024
		}
		if mobCounter[2048] >= minCount {
			zoneLevel.Levels |= 2048
		}
		if mobCounter[4096] >= minCount {
			zoneLevel.Levels |= 4096
			if zoneLevel.Levels&1 == 1 {
				zoneLevel.Levels--
			}
			if zoneLevel.Levels&2 == 2 {
				zoneLevel.Levels -= 2
			}
			if zoneLevel.Levels&4 == 4 {
				zoneLevel.Levels -= 4
			}
		}
		if mobCounter[8192] >= minCount {
			zoneLevel.Levels |= 8192
			if zoneLevel.Levels&1 == 1 {
				zoneLevel.Levels--
			}
			if zoneLevel.Levels&2 == 2 {
				zoneLevel.Levels -= 2
			}
			if zoneLevel.Levels&4 == 4 {
				zoneLevel.Levels -= 4
			}
		}
		if mobCounter[16384] >= minCount {
			zoneLevel.Levels |= 16384
			if zoneLevel.Levels&1 == 1 {
				zoneLevel.Levels--
			}
			if zoneLevel.Levels&2 == 2 {
				zoneLevel.Levels -= 2
			}
			if zoneLevel.Levels&4 == 4 {
				zoneLevel.Levels -= 4
			}
		}

		err = a.zoneLevelRepo.Create(zoneLevel)

		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("Failed to create cache entry zone_id: %d levels: %d", zoneLevel.ZoneID, zoneLevel.Levels))
			fmt.Println(err)
			continue
		}
		fmt.Println(zone.ZoneIDNumber, time.Since(start))
	}

	fmt.Println("Finished in", time.Since(start))
	return
}
