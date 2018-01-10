package cases

import (
	"fmt"
)

func className(classID int64) string {
	val, ok := classes[classID]
	if !ok {
		return fmt.Sprintf("Unknown (%d)", classID)
	}
	return val
}

var classes = map[int64]string{
	1:  "Warrior",
	2:  "Cleric",
	3:  "Paladin",
	4:  "Ranger",
	5:  "Shadowknight",
	6:  "Druid",
	7:  "Monk",
	8:  "Bard",
	9:  "Rogue",
	10: "Shaman",
	11: "Necromancer",
	12: "Wizard",
	13: "Magician",
	14: "Enchanter",
	15: "Beastlord",
	16: "Berserker",
	20: "GM Warrior",
	21: "GM Cleric",
	22: "GM Paladin",
	23: "GM Ranger",
	24: "GM Shadow Knight",
	25: "GM Druid",
	26: "GM Monk",
	27: "GM Bard",
	28: "GM Rogue",
	29: "GM Shaman",
	30: "GM Necromancer",
	31: "GM Wizard",
	32: "GM Magician",
	33: "GM Enchanter",
	34: "GM Beastlord",
	35: "GM Berserker",
	40: "Banker",
	41: "Shopkeeper",
	59: "Discord Merchant",
	60: "Adventure Recruiter",
	61: "Adventure Merchant",
	63: "Tribute Master",
	64: "Guild Tribute Master?",
	66: "Guild Bank",
	67: "Radiant Crystal Merchant",
	68: "Ebon Crystal Merchant",
	69: "Fellowships",
	70: "Alternate Currency Merchant",
	71: "Mercenary Merchant",
}
