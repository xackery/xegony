package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

var ()

//RaceRepository handles RaceRepository cases and is a gateway to storage
type RaceRepository struct {
	stor              storage.Storage
	raceCache         map[int64]*model.Race
	isRaceCacheLoaded bool
}

//Initialize handler
func (c *RaceRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}

	c.stor = stor
	c.isRaceCacheLoaded = false
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *RaceRepository) rebuildCache() (err error) {
	if c.isRaceCacheLoaded {
		return
	}
	c.isRaceCacheLoaded = true
	c.raceCache = make(map[int64]*model.Race)
	races, err := c.list()
	if err != nil {
		return
	}

	for _, race := range races {
		c.raceCache[race.ID] = race
	}
	return
}

//Get handler
func (c *RaceRepository) Get(race *model.Race, user *model.User) (err error) {
	race = c.raceCache[race.ID]
	//race, err = c.stor.GetRace(raceID)
	return
}

//GetByName gets a race by it's name
func (c *RaceRepository) GetByName(race *model.Race, user *model.User) (err error) {
	for _, raceC := range c.raceCache {
		if raceC.Name == race.Name {
			race = raceC
			return
		}
	}
	return
}

//Create handler
func (c *RaceRepository) Create(race *model.Race, user *model.User) (err error) {
	if race == nil {
		err = fmt.Errorf("Empty race")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	race.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(race))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	//err = c.stor.CreateRace(race)
	//if err != nil {
	//	return
	//}
	c.isRaceCacheLoaded = false
	c.rebuildCache()
	return
}

//Edit handler
func (c *RaceRepository) Edit(raceID int64, race *model.Race, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(race))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	//if err = c.stor.EditRace(raceID, race); err != nil {
	//	return
	//}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

//Delete handler
func (c *RaceRepository) Delete(raceID int64, user *model.User) (err error) {
	//err = c.stor.DeleteRace(raceID)
	//if err != nil {
	//	return
	//}
	//if err = c.rebuildCache(); err != nil {
	//	return
	//}
	return
}

func (c *RaceRepository) list() (races []*model.Race, err error) {
	for _, race := range racesList {
		races = append(races, race)
	}
	return
}

//List handler
func (c *RaceRepository) List(user *model.User) (races []*model.Race, err error) {
	for _, race := range c.raceCache {
		races = append(races, race)
	}
	return
}

func (c *RaceRepository) prepare(race *model.Race, user *model.User) (err error) {

	return
}

//newSchema handler
func (c *RaceRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

//getSchemaProperty handler
func (c *RaceRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	case "type":
		prop.Type = "integer"
		prop.Minimum = 0
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}
	return
}

var racesList = map[int64]*model.Race{
	1: {
		ID:     1,
		Name:   "Human",
		Male:   "HUM",
		Female: "HUF",
		Icon:   "xa-player",
	},
	2: {
		ID:     2,
		Name:   "Barbarian",
		Male:   "BAM",
		Female: "BAF",
		Icon:   "xa-fox",
	},

	3: {
		ID:     3,
		Name:   "Erudite",
		Male:   "ERM",
		Female: "ERF",
		Icon:   "xa-book",
	},
	4: {
		ID:     4,
		Name:   "Wood Elf",
		Male:   "ELM",
		Female: "ELF",
		Icon:   "xa-pine-tree",
	},
	5: {
		ID:     5,
		Name:   "High Elf",
		Male:   "HIM",
		Female: "HIF",
		Icon:   "xa-tesla",
	},
	6: {
		ID:     6,
		Name:   "Dark Elf",
		Male:   "DAM",
		Female: "DAF",
		Icon:   "xa-bleeding-eye",
	},
	7: {
		ID:     7,
		Name:   "Half Elf",
		Male:   "HAM",
		Female: "HAF",
		Icon:   "xa-aware",
	},
	8: {
		ID:     8,
		Name:   "Dwarf",
		Male:   "DWM",
		Female: "DWF",
		Icon:   "xa-beer",
	},
	9: {
		ID:     9,
		Name:   "Troll",
		Male:   "TRM",
		Female: "TRF",
		Icon:   "xa-bird-mask",
	},
	10: {
		ID:     10,
		Name:   "Ogre",
		Male:   "OGM",
		Female: "OGF",
		Icon:   "xa-muscle-fat",
	},
	11: {
		ID:     11,
		Name:   "Halfling",
		Male:   "HOM",
		Female: "HOF",
		Icon:   "xa-footprint",
	},
	12: {
		ID:     12,
		Name:   "Gnome",
		Male:   "GNM",
		Female: "GNF",
		Icon:   "xa-gears",
	},
	13: {
		ID:      13,
		Name:    "Aviak",
		Neutral: "AVI",
	},
	14: {
		ID:      14,
		Name:    "Werewolf",
		Neutral: "WER",
	},
	15: {
		ID:     15,
		Name:   "Brownie",
		Male:   "BRM",
		Female: "BRF",
	},
	16: {
		ID:      16,
		Name:    "Centaur",
		Male:    "CEM",
		Female:  "CEF",
		Neutral: "CEN",
	},
	17: {
		ID:      17,
		Name:    "Golem",
		Male:    "GOM",
		Female:  "GOF",
		Neutral: "GOL",
	},
	18: {
		ID:      18,
		Name:    "Giant",
		Neutral: "GIA",
	},
	19: {
		ID:      19,
		Name:    "Trakanon",
		Neutral: "TRK",
	},
	20: {
		ID:      20,
		Name:    "Venril Sathir",
		Neutral: "VST",
	},
	21: {
		ID:      21,
		Name:    "Evil Eye",
		Neutral: "BEH",
	},
	22: {
		ID:      22,
		Name:    "Beetle",
		Neutral: "BET",
	},
	23: {
		ID:     23,
		Name:   "Kerran",
		Male:   "CPM",
		Female: "CPF",
	},
	24: {
		ID:      24,
		Name:    "Fish",
		Neutral: "FIS",
	},

	25: {
		ID:     25,
		Name:   "Fairy",
		Male:   "FAM",
		Female: "FAF",
	},
	26: {
		ID:      26,
		Name:    "Froglok",
		Neutral: "FRO",
		Icon:    "xa-water-drop",
	},
	27: {
		ID:      27,
		Name:    "Froglok",
		Neutral: "FRG",
	},
	28: {
		ID:      28,
		Name:    "Fungusman",
		Neutral: "FUN",
	},
	29: {
		ID:      29,
		Name:    "Gargoyle",
		Male:    "GAM",
		Neutral: "GAR",
	},
	30: {
		ID:      30,
		Name:    "Gasbag",
		Neutral: "BEH",
	},
	31: {
		ID:      31,
		Name:    "Gelatinous Cube",
		Neutral: "CUB",
	},
	32: {
		ID:     32,
		Name:   "Ghost",
		Male:   "GHM",
		Female: "GHF",
	},
	33: {
		ID:      33,
		Name:    "Ghoul",
		Neutral: "GHU",
	},
	34: {
		ID:      34,
		Name:    "Bat",
		Neutral: "BAT",
	},
	36: {
		ID:      36,
		Name:    "Rat",
		Neutral: "RAT",
	},
	37: {
		ID:      37,
		Name:    "Snake",
		Neutral: "SNA",
	},
	38: {
		ID:      38,
		Name:    "Spider",
		Neutral: "SPI",
	},
	39: {
		ID:      39,
		Name:    "Gnoll",
		Neutral: "GNN",
	},
	40: {
		ID:      40,
		Name:    "Goblin",
		Neutral: "GOB",
	},
	41: {
		ID:      41,
		Name:    "Gorilla",
		Neutral: "GOR",
	},
	42: {
		ID:      42,
		Name:    "Wolf",
		Male:    "WOM",
		Female:  "WOF",
		Neutral: "WOL",
	},
	43: {
		ID:      43,
		Name:    "Bear",
		Neutral: "BEA",
	},
	44: {
		ID:     44,
		Name:   "Guard",
		Male:   "FPM",
		Female: "FPF",
	},
	45: {
		ID:      45,
		Name:    "Demi Lich",
		Neutral: "DML",
	},
	46: {
		ID:      46,
		Name:    "Imp",
		Neutral: "IMP",
	},
	47: {
		ID:      47,
		Name:    "Griffin",
		Neutral: "GRI",
	},
	48: {
		ID:      48,
		Name:    "Kobold",
		Neutral: "KOB",
	},
	49: {
		ID:      49,
		Name:    "Dragon",
		Neutral: "DRA",
		Icon:    "xa-wyvern",
	},
	50: {
		ID:     50,
		Name:   "Lion",
		Male:   "LIM",
		Female: "LIF",
	},
	51: {
		ID:      51,
		Name:    "Lizard Man",
		Neutral: "LIZ",
	},
	52: {
		ID:     52,
		Name:   "Mimic",
		Male:   "MIM",
		Female: "MIF",
	},
	53: {
		ID:      53,
		Name:    "Minotaur",
		Neutral: "MIN",
	},
	54: {
		ID:      54,
		Name:    "Orc",
		Neutral: "ORC",
	},
	55: {
		ID:     55,
		Name:   "Beggar",
		Male:   "BGM",
		Female: "BGF",
	},
	56: {
		ID:     56,
		Name:   "Pixie",
		Male:   "PIM",
		Female: "PIF",
	},
	57: {
		ID:     57,
		Name:   "Drachnid",
		Male:   "DRM",
		Female: "DRF",
	},
	58: {
		ID:      58,
		Name:    "Solusek Ro",
		Neutral: "SOL",
	},
	59: {
		ID:      59,
		Name:    "Goblin",
		Neutral: "BGG",
	},
	60: {
		ID:      60,
		Name:    "Skeleton",
		Neutral: "SKE",
	},
	61: {
		ID:      61,
		Name:    "Shark",
		Neutral: "SHA",
	},
	62: {
		ID:      62,
		Name:    "Tunare",
		Neutral: "TUN",
	},
	63: {
		ID:      63,
		Name:    "Tiger",
		Neutral: "TIG",
	},
	64: {
		ID:      64,
		Name:    "Treant",
		Neutral: "TRE",
	},
	65: {
		ID:     65,
		Name:   "Vampire",
		Male:   "DVM",
		Female: "DVF",
	},
	66: {
		ID:      66,
		Name:    "Rallos Zek",
		Neutral: "RAL",
	},
	67: {
		ID:     67,
		Name:   "Human",
		Male:   "HHM",
		Female: "HHF",
	},
	68: {
		ID:      68,
		Name:    "Tentacle Terror",
		Male:    "TEM",
		Female:  "TEF",
		Neutral: "TEN",
	},
	69: {
		ID:      69,
		Name:    "Will-O-Wisp",
		Neutral: "WIL",
	},
	70: {
		ID:     70,
		Name:   "Zombie",
		Male:   "ZOM",
		Female: "ZOF",
	},
	71: {
		ID:     71,
		Name:   "Human",
		Male:   "QCM",
		Female: "QCF",
	},
	72: {
		ID:      72,
		Name:    "Ship",
		Male:    "SHIP",
		Female:  "PRE",
		Neutral: "PRE",
	},
	73: {
		ID:      73,
		Name:    "Launch",
		Male:    "LAUNCHM",
		Neutral: "LAUNCH",
	},
	74: {
		ID:      74,
		Name:    "Piranha",
		Neutral: "PIR",
	},
	75: {
		ID:      75,
		Name:    "Elemental",
		Neutral: "ELE",
	},
	76: {
		ID:      76,
		Name:    "Puma",
		Neutral: "PUM",
	},
	77: {
		ID:     77,
		Name:   "Dark Elf",
		Male:   "NGM",
		Female: "NGF",
	},
	78: {
		ID:     78,
		Name:   "Erudite",
		Male:   "EGM",
		Female: "EGF",
	},
	79: {
		ID:      79,
		Name:    "Bixie",
		Neutral: "BIX",
	},
	80: {
		ID:      80,
		Name:    "Reanimated Hand",
		Neutral: "REA",
	},
	81: {
		ID:     81,
		Name:   "Halfling",
		Male:   "RIM",
		Female: "RIF",
	},
	82: {
		ID:      82,
		Name:    "Scarecrow",
		Neutral: "SCA",
	},
	83: {
		ID:      83,
		Name:    "Skunk",
		Neutral: "SKU",
	},
	84: {
		ID:      84,
		Name:    "Snake Elemental",
		Neutral: "SNE",
	},
	85: {
		ID:      85,
		Name:    "Spectre",
		Neutral: "SPE",
	},
	86: {
		ID:      86,
		Name:    "Sphinx",
		Neutral: "SPH",
	},
	87: {
		ID:      87,
		Name:    "Armadillo",
		Neutral: "ARM",
	},
	88: {
		ID:      88,
		Name:    "Clockwork Gnome",
		Male:    "CLM",
		Female:  "CLF",
		Neutral: "CLN",
	},
	89: {
		ID:      89,
		Name:    "Drake",
		Neutral: "DRK",
	},
	90: {
		ID:     90,
		Name:   "Barbarian",
		Male:   "HLM",
		Female: "HLF",
	},
	91: {
		ID:      91,
		Name:    "Alligator",
		Neutral: "ALL",
	},
	92: {
		ID:     92,
		Name:   "Troll",
		Male:   "GRM",
		Female: "GRF",
	},
	93: {
		ID:     93,
		Name:   "Ogre",
		Male:   "OKM",
		Female: "OKF",
	},
	94: {
		ID:      94,
		Name:    "Dwarf",
		Male:    "KAM",
		Female:  "KAF",
		Neutral: "KAN",
	},
	95: {
		ID:      95,
		Name:    "Cazic Thule",
		Neutral: "CAZ",
	},
	96: {
		ID:      96,
		Name:    "Cockatrice",
		Neutral: "COC",
	},
	97: {
		ID:      97,
		Name:    "Daisy Man",
		Neutral: "DIA",
	},
	98: {
		ID:     98,
		Name:   "Vampire",
		Male:   "VSM",
		Female: "VSF",
	},
	99: {
		ID:      99,
		Name:    "Amygdalan",
		Neutral: "DEN",
	},
	100: {
		ID:      100,
		Name:    "Dervish",
		Neutral: "DER",
	},
	101: {
		ID:      101,
		Name:    "Efreeti",
		Neutral: "EFR",
	},
	102: {
		ID:      102,
		Name:    "Tadpole",
		Neutral: "FRT",
	},
	103: {
		ID:      103,
		Name:    "Kedge",
		Neutral: "KED",
	},
	104: {
		ID:      104,
		Name:    "Leech",
		Neutral: "LEE",
	},
	105: {
		ID:      105,
		Name:    "Swordfish",
		Neutral: "SWO",
	},
	106: {
		ID:     106,
		Name:   "Guard",
		Male:   "FEM",
		Female: "FEF",
	},
	107: {
		ID:      107,
		Name:    "Mammoth",
		Neutral: "MAM",
	},
	108: {
		ID:      108,
		Name:    "Eye",
		Neutral: "EYE",
	},
	109: {
		ID:      109,
		Name:    "Wasp",
		Neutral: "WAS",
	},
	110: {
		ID:      110,
		Name:    "Mermaid",
		Neutral: "MER",
	},
	111: {
		ID:      111,
		Name:    "Harpy",
		Neutral: "HAR",
	},
	112: {
		ID:     112,
		Name:   "Guard",
		Male:   "GFM",
		Female: "GFF",
	},
	113: {
		ID:      113,
		Name:    "Drixie",
		Neutral: "DRI",
	},
	114: {
		ID:      114,
		Name:    "Ghost Ship",
		Neutral: "GSP",
	},
	115: {
		ID:      115,
		Name:    "Clam",
		Neutral: "CLA",
	},
	116: {
		ID:      116,
		Name:    "Seahorse",
		Neutral: "SEA",
	},
	117: {
		ID:     117,
		Name:   "Ghost",
		Male:   "GDM",
		Female: "GDF",
	},
	118: {
		ID:     118,
		Name:   "Ghost",
		Male:   "GEM",
		Female: "GEF",
	},
	119: {
		ID:      119,
		Name:    "Saber-toothed Cat",
		Neutral: "STC",
	},
	120: {
		ID:      120,
		Name:    "Wolf",
		Neutral: "WOE",
	},
	121: {
		ID:      121,
		Name:    "Gorgon",
		Neutral: "GRG",
	},
	122: {
		ID:      122,
		Name:    "Dragon",
		Neutral: "DRU",
	},
	123: {
		ID:      123,
		Name:    "Innoruuk",
		Neutral: "INN",
	},
	124: {
		ID:      124,
		Name:    "Unicorn",
		Neutral: "UNI",
	},
	125: {
		ID:      125,
		Name:    "Pegasus",
		Neutral: "PEG",
	},
	126: {
		ID:      126,
		Name:    "Djinn",
		Neutral: "DJI",
	},
	127: {
		ID:     127,
		Name:   "Invisible Man",
		Male:   "IVM",
		Female: "IVF",
	},
	128: {
		ID:     128,
		Name:   "Iksar",
		Male:   "IKM",
		Female: "IKF",
		Icon:   "xa-gecko",
	},
	129: {
		ID:      129,
		Name:    "Scorpion",
		Neutral: "SCR",
	},
	130: {
		ID:     130,
		Name:   "Vah Shir",
		Male:   "KEM",
		Female: "KEF",
		Icon:   "xa-lion",
	},
	131: {
		ID:      131,
		Name:    "Sarnak",
		Neutral: "SRW",
	},
	132: {
		ID:      132,
		Name:    "Draglock",
		Neutral: "DLK",
	},
	133: {
		ID:      133,
		Name:    "Drolvarg",
		Neutral: "LYC",
	},
	134: {
		ID:      134,
		Name:    "Mosquito",
		Neutral: "MOS",
	},
	135: {
		ID:      135,
		Name:    "Rhinoceros",
		Neutral: "RHI",
	},
	136: {
		ID:      136,
		Name:    "Xalgoz",
		Neutral: "XAL",
	},
	137: {
		ID:      137,
		Name:    "Goblin",
		Neutral: "KGO",
	},
	138: {
		ID:      138,
		Name:    "Yeti",
		Neutral: "YET",
	},
	139: {
		ID:      139,
		Name:    "Iksar",
		Male:    "ICM",
		Female:  "ICF",
		Neutral: "ICN",
	},
	140: {
		ID:      140,
		Name:    "Giant",
		Neutral: "FGI",
	},
	141: {
		ID:      141,
		Name:    "Boat",
		Male:    "BOAT",
		Female:  "BOAT",
		Neutral: "BOAT",
	},
	144: {
		ID:      144,
		Name:    "Burynai",
		Neutral: "BRN",
	},
	145: {
		ID:      145,
		Name:    "Goo",
		Neutral: "GOO",
	},
	146: {
		ID:      146,
		Name:    "Sarnak Spirit",
		Neutral: "SSN",
	},
	147: {
		ID:     147,
		Name:   "Iksar Spirit",
		Male:   "SIM",
		Female: "SIF",
	},
	148: {
		ID:      148,
		Name:    "Fish",
		Neutral: "BAC",
	},
	149: {
		ID:      149,
		Name:    "Scorpion",
		Neutral: "ISC",
	},
	150: {
		ID:      150,
		Name:    "Erollisi",
		Neutral: "ERO",
	},
	151: {
		ID:      151,
		Name:    "Tribunal",
		Neutral: "TRI",
	},
	152: {
		ID:      152,
		Name:    "Bertoxxulous",
		Neutral: "BER",
	},
	153: {
		ID:      153,
		Name:    "Bristlebane",
		Neutral: "BRI",
	},
	154: {
		ID:      154,
		Name:    "Fay Drake",
		Neutral: "FDR",
	},
	155: {
		ID:      155,
		Name:    "Undead Sarnak",
		Neutral: "SSK",
	},
	156: {
		ID:      156,
		Name:    "Ratman",
		Neutral: "VRM",
	},
	157: {
		ID:      157,
		Name:    "Wyvern",
		Neutral: "WYV",
	},
	158: {
		ID:      158,
		Name:    "Wurm",
		Neutral: "WUR",
	},
	159: {
		ID:      159,
		Name:    "Devourer",
		Neutral: "DEV",
	},
	160: {
		ID:      160,
		Name:    "Iksar Golem",
		Neutral: "IKG",
	},
	161: {
		ID:      161,
		Name:    "Undead Iksar",
		Neutral: "IKS",
	},
	162: {
		ID:      162,
		Name:    "Man-Eating Plant",
		Neutral: "MEP",
	},
	163: {
		ID:      163,
		Name:    "Raptor",
		Neutral: "RAP",
	},
	164: {
		ID:      164,
		Name:    "Sarnak Golem",
		Neutral: "SGO",
	},
	165: {
		ID:      165,
		Name:    "Dragon",
		Neutral: "SED",
	},
	166: {
		ID:      166,
		Name:    "Animated Hand",
		Neutral: "IKH",
	},
	167: {
		ID:      167,
		Name:    "Succulent",
		Neutral: "SUC",
	},
	168: {
		ID:      168,
		Name:    "Holgresh",
		Neutral: "FMO",
	},
	169: {
		ID:      169,
		Name:    "Brontotherium",
		Neutral: "BTM",
	},
	170: {
		ID:      170,
		Name:    "Snow Dervish",
		Neutral: "SDE",
	},
	171: {
		ID:      171,
		Name:    "Dire Wolf",
		Neutral: "DIW",
	},
	172: {
		ID:      172,
		Name:    "Manticore",
		Neutral: "MTC",
	},
	173: {
		ID:      173,
		Name:    "Totem",
		Neutral: "TOT",
	},
	174: {
		ID:      174,
		Name:    "Ice Spectre",
		Neutral: "SPC",
	},
	175: {
		ID:      175,
		Name:    "Enchanted Armor",
		Neutral: "ENA",
	},
	176: {
		ID:      176,
		Name:    "Snow Rabbit",
		Neutral: "SBU",
	},
	177: {
		ID:      177,
		Name:    "Walrus",
		Neutral: "WAL",
	},
	178: {
		ID:      178,
		Name:    "Geonid",
		Neutral: "RGM",
	},
	181: {
		ID:      181,
		Name:    "Yakkar",
		Neutral: "YAK",
	},
	182: {
		ID:      182,
		Name:    "Faun",
		Neutral: "FAN",
	},
	183: {
		ID:      183,
		Name:    "Coldain",
		Male:    "COM",
		Female:  "COF",
		Neutral: "COK",
	},
	184: {
		ID:      184,
		Name:    "Dragon",
		Neutral: "DR2",
	},
	185: {
		ID:      185,
		Name:    "Hag",
		Neutral: "HAG",
	},
	186: {
		ID:      186,
		Name:    "Hippogriff",
		Neutral: "HIP",
	},
	187: {
		ID:      187,
		Name:    "Siren",
		Neutral: "SIR",
	},
	188: {
		ID:      188,
		Name:    "Giant",
		Neutral: "FSG",
	},
	189: {
		ID:      189,
		Name:    "Giant",
		Male:    "STM",
		Neutral: "STG",
	},
	190: {
		ID:      190,
		Name:    "Othmir",
		Neutral: "OTM",
	},
	191: {
		ID:      191,
		Name:    "Ulthork",
		Neutral: "WLM",
	},
	192: {
		ID:      192,
		Name:    "Dragon",
		Neutral: "CCD",
	},
	193: {
		ID:      193,
		Name:    "Abhorrent",
		Neutral: "ABH",
	},
	194: {
		ID:      194,
		Name:    "Sea Turtle",
		Neutral: "STU",
	},
	195: {
		ID:      195,
		Name:    "Dragon",
		Neutral: "BWD",
	},
	196: {
		ID:      196,
		Name:    "Dragon",
		Neutral: "GDR",
	},
	197: {
		ID:      197,
		Name:    "Ronnie Test",
		Neutral: "RON",
	},
	198: {
		ID:      198,
		Name:    "Dragon",
		Neutral: "PRI",
	},
	199: {
		ID:      199,
		Name:    "Shik'Nar",
		Neutral: "SKN",
	},
	200: {
		ID:      200,
		Name:    "Rockhopper",
		Neutral: "RHP",
	},
	201: {
		ID:      201,
		Name:    "Underbulk",
		Neutral: "UNB",
	},
	202: {
		ID:      202,
		Name:    "Grimling",
		Male:    "GMM",
		Female:  "GMF",
		Neutral: "GMN",
	},
	203: {
		ID:      203,
		Name:    "Worm",
		Neutral: "VAC",
	},
	204: {
		ID:      204,
		Name:    "Evan Test",
		Neutral: "ECS",
	},
	205: {
		ID:      205,
		Name:    "Shadel",
		Neutral: "KHA",
	},
	206: {
		ID:      206,
		Name:    "Owlbear",
		Neutral: "OWB",
	},
	207: {
		ID:      207,
		Name:    "Rhino Beetle",
		Neutral: "RNB",
	},
	208: {
		ID:     208,
		Name:   "Vampire",
		Male:   "VPM",
		Female: "VPF",
	},
	209: {
		ID:      209,
		Name:    "Earth Elemental",
		Neutral: "EEL",
	},
	210: {
		ID:      210,
		Name:    "Air Elemental",
		Neutral: "AEL",
	},
	211: {
		ID:      211,
		Name:    "Water Elemental",
		Neutral: "WEL",
	},
	212: {
		ID:      212,
		Name:    "Fire Elemental",
		Neutral: "FEL",
	},
	213: {
		ID:      213,
		Name:    "Wetfang Minnow",
		Neutral: "WET",
	},
	214: {
		ID:      214,
		Name:    "Thought Horror",
		Neutral: "THO",
	},
	215: {
		ID:      215,
		Name:    "Tegi",
		Neutral: "TEG",
	},
	216: {
		ID:      216,
		Name:    "Horse",
		Male:    "HSM",
		Female:  "HSF",
		Neutral: "HSN",
	},
	217: {
		ID:      217,
		Name:    "Shissar",
		Male:    "SHM",
		Female:  "SHF",
		Neutral: "SHN",
	},
	218: {
		ID:      218,
		Name:    "Fungal Fiend",
		Neutral: "FUG",
	},
	219: {
		ID:      219,
		Name:    "Vampire",
		Neutral: "VOL",
	},
	220: {
		ID:      220,
		Name:    "Stonegrabber",
		Neutral: "SGR",
	},
	221: {
		ID:      221,
		Name:    "Scarlet Cheetah",
		Neutral: "SCH",
	},
	222: {
		ID:      222,
		Name:    "Zelniak",
		Neutral: "ZEL",
	},
	223: {
		ID:      223,
		Name:    "Lightcrawler",
		Neutral: "LCR",
	},
	224: {
		ID:     224,
		Name:   "Shade",
		Male:   "SDM",
		Female: "SDF",
	},
	225: {
		ID:      225,
		Name:    "Sunflower",
		Neutral: "SNN",
	},
	226: {
		ID:      226,
		Name:    "Sun Revenant",
		Neutral: "SRV",
	},
	227: {
		ID:      227,
		Name:    "Shrieker",
		Neutral: "SHR",
	},
	228: {
		ID:      228,
		Name:    "Galorian",
		Neutral: "GAL",
	},
	229: {
		ID:      229,
		Name:    "Netherbian",
		Neutral: "NET",
	},
	230: {
		ID:      230,
		Name:    "Akheva",
		Male:    "AKM",
		Female:  "AKF",
		Neutral: "AKN",
	},
	231: {
		ID:      231,
		Name:    "Grieg Veneficus",
		Neutral: "SPR",
	},
	232: {
		ID:      232,
		Name:    "Sonic Wolf",
		Neutral: "SOW",
	},
	233: {
		ID:      233,
		Name:    "Ground Shaker",
		Male:    "GSM",
		Female:  "GSF",
		Neutral: "GSN",
	},
	234: {
		ID:      234,
		Name:    "Vah Shir Skeleton",
		Neutral: "KES",
	},
	235: {
		ID:      235,
		Name:    "Wretch",
		Neutral: "MUH",
	},
	236: {
		ID:      236,
		Name:    "Seru",
		Neutral: "SER",
	},
	237: {
		ID:      237,
		Name:    "Recuso",
		Male:    "REM",
		Female:  "REF",
		Neutral: "REN",
	},
	238: {
		ID:      238,
		Name:    "Vah Shir",
		Neutral: "VSK",
	},
	239: {
		ID:      239,
		Name:    "Guard",
		Neutral: "VSG",
	},
	240: {
		ID:      240,
		Name:    "Teleport Man",
		Male:    "TPM",
		Female:  "TPF",
		Neutral: "TPN",
	},
	241: {
		ID:      241,
		Name:    "Werewolf",
		Neutral: "LUJ",
	},
	242: {
		ID:      242,
		Name:    "Nymph",
		Neutral: "NYD",
	},
	243: {
		ID:      243,
		Name:    "Dryad",
		Neutral: "NYM",
	},
	244: {
		ID:      244,
		Name:    "Treant",
		Neutral: "TRN",
	},
	245: {
		ID:      245,
		Name:    "Fly",
		Neutral: "WRF",
	},
	246: {
		ID:      246,
		Name:    "Tarew Marr",
		Neutral: "TMR",
	},
	247: {
		ID:      247,
		Name:    "Solusek Ro",
		Neutral: "SRO",
	},
	248: {
		ID:      248,
		Name:    "Clockwork Golem",
		Neutral: "CLG",
	},
	249: {
		ID:      249,
		Name:    "Clockwork Brain",
		Neutral: "CLB",
	},
	250: {
		ID:      250,
		Name:    "Banshee",
		Neutral: "SKB",
	},
	251: {
		ID:      251,
		Name:    "Guard of Justice",
		Neutral: "GOJ",
	},
	252: {
		ID:      252,
		Name:    "Mini POM",
		Male:    "MINIPOM200",
		Female:  "MINIPOM200",
		Neutral: "MINIPOM200",
	},
	253: {
		ID:      253,
		Name:    "Diseased Fiend",
		Neutral: "DSB",
	},
	254: {
		ID:      254,
		Name:    "Solusek Ro Guard",
		Neutral: "SRG",
	},
	255: {
		ID:      255,
		Name:    "Bertoxxulous",
		Neutral: "BTX",
	},
	256: {
		ID:      256,
		Name:    "The Tribunal",
		Neutral: "TBU",
	},
	257: {
		ID:      257,
		Name:    "Terris Thule",
		Neutral: "TRT",
	},
	258: {
		ID:      258,
		Name:    "Vegerog",
		Neutral: "VEG",
	},
	259: {
		ID:      259,
		Name:    "Crocodile",
		Neutral: "CRO",
	},
	260: {
		ID:      260,
		Name:    "Bat",
		Neutral: "NBT",
	},
	261: {
		ID:      261,
		Name:    "Hraquis",
		Neutral: "SLG",
	},
	262: {
		ID:      262,
		Name:    "Tranquilion",
		Neutral: "TRQ",
	},
	263: {
		ID:      263,
		Name:    "Tin Soldier",
		Neutral: "TIN",
	},
	264: {
		ID:      264,
		Name:    "Nightmare Wraith",
		Neutral: "NMW",
	},
	265: {
		ID:      265,
		Name:    "Malarian",
		Neutral: "MAL",
	},
	266: {
		ID:      266,
		Name:    "Knight of Pestilence",
		Neutral: "KOP",
	},
	267: {
		ID:      267,
		Name:    "Lepertoloth",
		Neutral: "LEP",
	},
	268: {
		ID:      268,
		Name:    "Bubonian",
		Neutral: "BUB",
	},
	269: {
		ID:      269,
		Name:    "Bubonian Underling",
		Neutral: "BUU",
	},
	270: {
		ID:      270,
		Name:    "Pusling",
		Neutral: "PUS",
	},
	271: {
		ID:      271,
		Name:    "Water Mephit",
		Neutral: "WMP",
	},
	272: {
		ID:      272,
		Name:    "Stormrider",
		Neutral: "STR",
	},
	273: {
		ID:      273,
		Name:    "Junk Beast",
		Neutral: "JUB",
	},
	274: {
		ID:      274,
		Name:    "Broken Clockwork",
		Neutral: "BRC",
	},
	275: {
		ID:      275,
		Name:    "Giant Clockwork",
		Neutral: "GLC",
	},
	276: {
		ID:      276,
		Name:    "Clockwork Beetle",
		Neutral: "CWB",
	},
	277: {
		ID:      277,
		Name:    "Nightmare Goblin",
		Neutral: "NMG",
	},
	278: {
		ID:      278,
		Name:    "Karana",
		Neutral: "KAR",
	},
	279: {
		ID:      279,
		Name:    "Blood Raven",
		Neutral: "BRV",
	},
	280: {
		ID:      280,
		Name:    "Nightmare Gargoyle",
		Neutral: "GGL",
	},
	281: {
		ID:      281,
		Name:    "Mouth of Insanity",
		Neutral: "MOI",
	},
	282: {
		ID:      282,
		Name:    "Skeletal Horse",
		Neutral: "HSS",
	},
	283: {
		ID:      283,
		Name:    "Saryrn",
		Neutral: "SAR",
	},
	284: {
		ID:      284,
		Name:    "Fennin Ro",
		Neutral: "FEN",
	},
	285: {
		ID:      285,
		Name:    "Tormentor",
		Neutral: "TRW",
	},
	286: {
		ID:      286,
		Name:    "Soul Devourer",
		Neutral: "NPT",
	},
	287: {
		ID:      287,
		Name:    "Nightmare",
		Neutral: "NMH",
	},
	288: {
		ID:      288,
		Name:    "Rallos Zek",
		Neutral: "RAZ",
	},
	289: {
		ID:      289,
		Name:    "Vallon Zek",
		Neutral: "VAZ",
	},
	290: {
		ID:      290,
		Name:    "Tallon Zek",
		Neutral: "TAZ",
	},
	291: {
		ID:      291,
		Name:    "Air Mephit",
		Neutral: "AMP",
	},
	292: {
		ID:      292,
		Name:    "Earth Mephit",
		Neutral: "EMP",
	},
	293: {
		ID:      293,
		Name:    "Fire Mephit",
		Neutral: "FMP",
	},
	294: {
		ID:      294,
		Name:    "Nightmare Mephit",
		Neutral: "NMP",
	},
	295: {
		ID:      295,
		Name:    "Zebuxoruk",
		Neutral: "ZEB",
	},
	296: {
		ID:      296,
		Name:    "Mithaniel Marr",
		Neutral: "MAR",
	},
	297: {
		ID:      297,
		Name:    "Undead Knight",
		Neutral: "UDK",
	},
	298: {
		ID:      298,
		Name:    "The Rathe",
		Neutral: "RTH",
	},
	299: {
		ID:      299,
		Name:    "Xegony",
		Neutral: "XEG",
	},
	300: {
		ID:      300,
		Name:    "Fiend",
		Neutral: "GTD",
	},
	301: {
		ID:      301,
		Name:    "Test Object",
		Neutral: "ONT",
	},
	302: {
		ID:      302,
		Name:    "Crab",
		Neutral: "CRB",
	},
	303: {
		ID:      303,
		Name:    "Phoenix",
		Neutral: "PHX",
	},
	304: {
		ID:      304,
		Name:    "Dragon",
		Neutral: "TMT",
	},
	305: {
		ID:      305,
		Name:    "Bear",
		Neutral: "PBR",
	},
	306: {
		ID:      306,
		Name:    "Giant",
		Neutral: "STA",
	},
	307: {
		ID:      307,
		Name:    "Giant",
		Neutral: "SSA",
	},
	308: {
		ID:      308,
		Name:    "Giant",
		Neutral: "SKR",
	},
	309: {
		ID:      309,
		Name:    "Giant",
		Neutral: "SVO",
	},
	310: {
		ID:      310,
		Name:    "Giant",
		Neutral: "SMA",
	},
	311: {
		ID:      311,
		Name:    "Giant",
		Neutral: "STF",
	},
	312: {
		ID:      312,
		Name:    "Giant",
		Neutral: "SCE",
	},
	313: {
		ID:      313,
		Name:    "War Wraith",
		Neutral: "WRW",
	},
	314: {
		ID:      314,
		Name:    "Wrulon",
		Neutral: "WRU",
	},
	315: {
		ID:      315,
		Name:    "Kraken",
		Neutral: "KRK",
	},
	316: {
		ID:      316,
		Name:    "Poison Frog",
		Neutral: "PAF",
	},
	317: {
		ID:      317,
		Name:    "Nilborien",
		Neutral: "QZT",
	},
	318: {
		ID:      318,
		Name:    "Valorian",
		Neutral: "VAL",
	},
	319: {
		ID:      319,
		Name:    "War Boar",
		Neutral: "WRB",
	},
	320: {
		ID:      320,
		Name:    "Efreeti",
		Neutral: "EFE",
	},
	321: {
		ID:      321,
		Name:    "War Boar",
		Neutral: "WBU",
	},
	322: {
		ID:      322,
		Name:    "Valorian",
		Neutral: "BKN",
	},
	323: {
		ID:      323,
		Name:    "Animated Armor",
		Neutral: "AAM",
	},
	324: {
		ID:      324,
		Name:    "Undead Footman",
		Neutral: "UDF",
	},
	325: {
		ID:      325,
		Name:    "Rallos Zek Minion",
		Neutral: "RZM",
	},
	326: {
		ID:      326,
		Name:    "Arachnid",
		Neutral: "SPD",
	},
	327: {
		ID:      327,
		Name:    "Crystal Spider",
		Neutral: "SPL",
	},
	328: {
		ID:      328,
		Name:    "Zebuxoruk's Cage",
		Neutral: "ZBC",
	},
	329: {
		ID:      329,
		Name:    "BoT Portal",
		Neutral: "BTP",
	},
	330: {
		ID:     330,
		Name:   "Froglok",
		Male:   "FRM",
		Female: "FRF",
	},
	331: {
		ID:     331,
		Name:   "Troll",
		Male:   "TBM",
		Female: "TBF",
	},
	332: {
		ID:     332,
		Name:   "Troll",
		Male:   "FBM",
		Female: "FBF",
	},
	333: {
		ID:     333,
		Name:   "Troll",
		Male:   "TSM",
		Female: "TSF",
	},
	334: {
		ID:      334,
		Name:    "Ghost",
		Neutral: "SPB",
	},
	335: {
		ID:      335,
		Name:    "Pirate",
		Neutral: "TPB",
	},
	336: {
		ID:      336,
		Name:    "Pirate",
		Neutral: "TVP",
	},
	337: {
		ID:      337,
		Name:    "Pirate",
		Neutral: "TPO",
	},
	338: {
		ID:     338,
		Name:   "Pirate",
		Male:   "GPM",
		Female: "GPF",
	},
	339: {
		ID:     339,
		Name:   "Pirate",
		Male:   "DPM",
		Female: "DPF",
	},
	340: {
		ID:     340,
		Name:   "Pirate",
		Male:   "OPM",
		Female: "OPF",
	},
	341: {
		ID:     341,
		Name:   "Pirate",
		Male:   "HPM",
		Female: "HPF",
	},
	342: {
		ID:     342,
		Name:   "Pirate",
		Male:   "EPM",
		Female: "EPF",
	},
	343: {
		ID:      343,
		Name:    "Frog",
		Neutral: "RPF",
	},
	344: {
		ID:     344,
		Name:   "Troll Zombie",
		Male:   "TZM",
		Female: "TZF",
	},
	345: {
		ID:      345,
		Name:    "Luggald",
		Neutral: "LUG",
	},
	346: {
		ID:      346,
		Name:    "Luggald",
		Neutral: "LGA",
	},
	347: {
		ID:      347,
		Name:    "Luggalds",
		Neutral: "LGR",
	},
	348: {
		ID:      348,
		Name:    "Drogmore",
		Neutral: "FMT",
	},
	349: {
		ID:      349,
		Name:    "Froglok Skeleton",
		Neutral: "FSK",
	},
	350: {
		ID:      350,
		Name:    "Undead Froglok",
		Neutral: "FUD",
	},
	351: {
		ID:     351,
		Name:   "Knight of Hate",
		Male:   "IWM",
		Female: "IWF",
	},
	352: {
		ID:     352,
		Name:   "Arcanist of Hate",
		Male:   "IZM",
		Female: "IZF",
	},
	353: {
		ID:      353,
		Name:    "Veksar",
		Neutral: "VEK",
	},
	354: {
		ID:      354,
		Name:    "Veksar",
		Neutral: "GVK",
	},
	355: {
		ID:      355,
		Name:    "Veksar",
		Neutral: "BVK",
	},
	356: {
		ID:      356,
		Name:    "Chokidai",
		Neutral: "WOF",
	},
	357: {
		ID:      357,
		Name:    "Undead Chokidai",
		Neutral: "WUF",
	},
	358: {
		ID:      358,
		Name:    "Undead Veksar",
		Neutral: "UVK",
	},
	359: {
		ID:     359,
		Name:   "Vampire",
		Male:   "LMM",
		Female: "LMF",
	},
	360: {
		ID:     360,
		Name:   "Vampire",
		Male:   "MMM",
		Female: "MMF",
	},
	361: {
		ID:      361,
		Name:    "Rujarkian Orc",
		Neutral: "ROM",
	},
	362: {
		ID:      362,
		Name:    "Bone Golem",
		Neutral: "BGB",
	},
	363: {
		ID:      363,
		Name:    "Synarcana",
		Neutral: "SYN",
	},
	364: {
		ID:     364,
		Name:   "Sand Elf",
		Male:   "SEM",
		Female: "SEF",
	},
	365: {
		ID:      365,
		Name:    "Vampire",
		Neutral: "MMV",
	},
	366: {
		ID:      366,
		Name:    "Rujarkian Orc",
		Neutral: "ROE",
	},
	367: {
		ID:      367,
		Name:    "Skeleton",
		Neutral: "SKT",
	},
	368: {
		ID:      368,
		Name:    "Mummy",
		Neutral: "MMY",
	},
	369: {
		ID:      369,
		Name:    "Goblin",
		Neutral: "GBL",
	},
	370: {
		ID:      370,
		Name:    "Insect",
		Neutral: "NIN",
	},
	371: {
		ID:      371,
		Name:    "Froglok Ghost",
		Neutral: "FGH",
	},
	372: {
		ID:      372,
		Name:    "Dervish",
		Neutral: "DRV",
	},
	373: {
		ID:      373,
		Name:    "Shade",
		Neutral: "SDC",
	},
	374: {
		ID:      374,
		Name:    "Golem",
		Neutral: "GLM",
	},
	375: {
		ID:      375,
		Name:    "Evil Eye",
		Neutral: "EEY",
	},
	376: {
		ID:      376,
		Name:    "Box",
		Neutral: "BOX",
	},
	377: {
		ID:      377,
		Name:    "Barrel",
		Neutral: "BRL",
	},
	378: {
		ID:      378,
		Name:    "Chest",
		Neutral: "CST",
	},
	379: {
		ID:      379,
		Name:    "Vase",
		Neutral: "VAS",
	},
	380: {
		ID:      380,
		Name:    "Table",
		Neutral: "TBL",
	},
	381: {
		ID:      381,
		Name:    "Weapon Rack",
		Neutral: "RAK",
	},
	382: {
		ID:      382,
		Name:    "Coffin",
		Neutral: "CPT",
	},
	383: {
		ID:      383,
		Name:    "Bones",
		Neutral: "BON",
	},
	384: {
		ID:      384,
		Name:    "Jokester",
		Neutral: "JKR",
	},
	385: {
		ID:     385,
		Name:   "Nihil",
		Male:   "TNM",
		Female: "TNF",
	},
	386: {
		ID:     386,
		Name:   "Trusik",
		Male:   "TEM",
		Female: "TEF",
	},
	387: {
		ID:      387,
		Name:    "Stone Worker",
		Neutral: "TGL",
	},
	388: {
		ID:      388,
		Name:    "Hynid",
		Neutral: "TWF",
	},
	389: {
		ID:      389,
		Name:    "Turepta",
		Neutral: "TAC",
	},
	390: {
		ID:      390,
		Name:    "Cragbeast",
		Neutral: "TMB",
	},
	391: {
		ID:      391,
		Name:    "Stonemite",
		Neutral: "TTB",
	},
	392: {
		ID:      392,
		Name:    "Ukun",
		Neutral: "IWH",
	},
	393: {
		ID:      393,
		Name:    "Ixt",
		Neutral: "IEC",
	},
	394: {
		ID:      394,
		Name:    "Ikaav",
		Neutral: "ILA",
	},
	395: {
		ID:      395,
		Name:    "Aneuk",
		Neutral: "ICY",
	},
	396: {
		ID:      396,
		Name:    "Kyv",
		Neutral: "IHU",
	},
	397: {
		ID:      397,
		Name:    "Noc",
		Neutral: "ISB",
	},
	398: {
		ID:      398,
		Name:    "Ra`tuk",
		Neutral: "IBR",
	},
	399: {
		ID:      399,
		Name:    "Taneth",
		Neutral: "IFC",
	},
	400: {
		ID:      400,
		Name:    "Huvul",
		Neutral: "ILB",
	},
	401: {
		ID:      401,
		Name:    "Mutna",
		Neutral: "IWB",
	},
	402: {
		ID:      402,
		Name:    "Mastruq",
		Neutral: "ISE",
	},
	403: {
		ID:      403,
		Name:    "Taelosian",
		Neutral: "TLN",
	},
	404: {
		ID:      404,
		Name:    "Discord Ship",
		Neutral: "SHP",
	},
	405: {
		ID:      405,
		Name:    "Stone Worker",
		Neutral: "TGO",
	},
	406: {
		ID:      406,
		Name:    "Mata Muram",
		Neutral: "CLV",
	},
	407: {
		ID:      407,
		Name:    "Lightning Warrior",
		Neutral: "MUR",
	},
	408: {
		ID:      408,
		Name:    "Succubus",
		Neutral: "SCU",
	},
	409: {
		ID:      409,
		Name:    "Bazu",
		Neutral: "DSG",
	},
	410: {
		ID:      410,
		Name:    "Feran",
		Neutral: "FRA",
	},
	411: {
		ID:      411,
		Name:    "Pyrilen",
		Neutral: "SCU",
	},
	412: {
		ID:      412,
		Name:    "Chimera",
		Neutral: "CHM",
	},
	413: {
		ID:      413,
		Name:    "Dragorn",
		Neutral: "DDM",
	},
	414: {
		ID:      414,
		Name:    "Murkglider",
		Neutral: "DMA",
	},
	415: {
		ID:      415,
		Name:    "Rat",
		Neutral: "RTN",
	},
	416: {
		ID:      416,
		Name:    "Bat",
		Neutral: "BTN",
	},
	417: {
		ID:      417,
		Name:    "Gelidran",
		Neutral: "FRD",
	},
	418: {
		ID:      418,
		Name:    "Discordling",
		Neutral: "DSF",
	},
	419: {
		ID:      419,
		Name:    "Girplan",
		Neutral: "DDV",
	},
	420: {
		ID:      420,
		Name:    "Minotaur",
		Neutral: "MNT",
	},
	421: {
		ID:      421,
		Name:    "Dragorn Box",
		Neutral: "DBX",
	},
	422: {
		ID:      422,
		Name:    "Runed Orb",
		Neutral: "ROB",
	},
	423: {
		ID:      423,
		Name:    "Dragon Bones",
		Neutral: "DBP",
	},
	424: {
		ID:      424,
		Name:    "Muramite Armor Pile",
		Neutral: "MAP",
	},
	425: {
		ID:      425,
		Name:    "Crystal Shard",
		Neutral: "CRS",
	},
	426: {
		ID:      426,
		Name:    "Portal",
		Neutral: "PRT",
	},
	427: {
		ID:      427,
		Name:    "Coin Purse",
		Neutral: "CNP",
	},
	428: {
		ID:      428,
		Name:    "Rock Pile",
		Neutral: "RKP",
	},
	429: {
		ID:      429,
		Name:    "Murkglider Egg Sac",
		Neutral: "MES",
	},
	430: {
		ID:      430,
		Name:    "Drake",
		Neutral: "CDR",
	},
	431: {
		ID:      431,
		Name:    "Dervish",
		Neutral: "DVS",
	},
	432: {
		ID:      432,
		Name:    "Drake",
		Neutral: "DKE",
	},
	433: {
		ID:      433,
		Name:    "Goblin",
		Neutral: "GBN",
	},
	434: {
		ID:      434,
		Name:    "Kirin",
		Neutral: "KRN",
	},
	435: {
		ID:      435,
		Name:    "Dragon",
		Neutral: "LDR",
	},
	436: {
		ID:      436,
		Name:    "Basilisk",
		Neutral: "BAS",
	},
	437: {
		ID:      437,
		Name:    "Dragon",
		Neutral: "MDR",
	},
	438: {
		ID:      438,
		Name:    "Dragon",
		Neutral: "SDR",
	},
	439: {
		ID:      439,
		Name:    "Puma",
		Neutral: "PMA",
	},
	440: {
		ID:      440,
		Name:    "Spider",
		Neutral: "TAR",
	},
	441: {
		ID:      441,
		Name:    "Spider Queen",
		Neutral: "SPQ",
	},
	442: {
		ID:      442,
		Name:    "Animated Statue",
		Neutral: "ANS",
	},
	445: {
		ID:      445,
		Name:    "Dragon Egg",
		Neutral: "DRE",
	},
	446: {
		ID:      446,
		Name:    "Dragon Statue",
		Neutral: "DRS",
	},
	447: {
		ID:      447,
		Name:    "Lava Rock",
		Neutral: "LVR",
	},
	448: {
		ID:      448,
		Name:    "Animated Statue",
		Neutral: "ASM",
	},
	449: {
		ID:      449,
		Name:    "Spider Egg Sack",
		Neutral: "SEG",
	},
	450: {
		ID:      450,
		Name:    "Lava Spider",
		Neutral: "LSP",
	},
	451: {
		ID:      451,
		Name:    "Lava Spider Queen",
		Neutral: "LSQ",
	},
	452: {
		ID:      452,
		Name:    "Dragon",
		Neutral: "SHD",
	},
	453: {
		ID:      453,
		Name:    "Giant",
		Neutral: "FGT",
	},
	454: {
		ID:      454,
		Name:    "Werewolf",
		Neutral: "WWF",
	},
	455: {
		ID:      455,
		Name:    "Kobold",
		Neutral: "KBD",
	},
	456: {
		ID:      456,
		Name:    "Sporali",
		Neutral: "FNG",
	},
	457: {
		ID:      457,
		Name:    "Gnomework",
		Neutral: "CWG",
	},
	458: {
		ID:      458,
		Name:    "Orc",
		Neutral: "ORK",
	},
	459: {
		ID:      459,
		Name:    "Corathus",
		Neutral: "CRH",
	},
	460: {
		ID:      460,
		Name:    "Coral",
		Neutral: "CRL",
	},
	461: {
		ID:     461,
		Name:   "Drachnid",
		Male:   "DCM",
		Female: "DCF",
	},
	462: {
		ID:      462,
		Name:    "Drachnid Cocoon",
		Neutral: "DRC",
	},
	463: {
		ID:      463,
		Name:    "Fungus Patch",
		Neutral: "FGP",
	},
	464: {
		ID:      464,
		Name:    "Gargoyle",
		Neutral: "GGY",
	},
	465: {
		ID:      465,
		Name:    "Witheran",
		Neutral: "KOR",
	},
	466: {
		ID:      466,
		Name:    "Dark Lord",
		Neutral: "MYG",
	},
	467: {
		ID:      467,
		Name:    "Shiliskin",
		Neutral: "SHL",
	},
	468: {
		ID:      468,
		Name:    "Snake",
		Neutral: "SNK",
	},
	469: {
		ID:      469,
		Name:    "Evil Eye",
		Neutral: "EVE",
	},
	470: {
		ID:      470,
		Name:    "Minotaur",
		Neutral: "MNR",
	},
	471: {
		ID:     471,
		Name:   "Zombie",
		Male:   "ZMM",
		Female: "ZMF",
	},
	472: {
		ID:      472,
		Name:    "Clockwork Boar",
		Neutral: "CWR",
	},
	473: {
		ID:      473,
		Name:    "Fairy",
		Neutral: "FRY",
	},
	474: {
		ID:      474,
		Name:    "Witheran",
		Neutral: "KRB",
	},
	475: {
		ID:      475,
		Name:    "Air Elemental",
		Neutral: "AIE",
	},
	476: {
		ID:      476,
		Name:    "Earth Elemental",
		Neutral: "EAE",
	},
	477: {
		ID:      477,
		Name:    "Fire Elemental",
		Neutral: "FIE",
	},
	478: {
		ID:      478,
		Name:    "Water Elemental",
		Neutral: "WAE",
	},
	479: {
		ID:      479,
		Name:    "Alligator",
		Neutral: "ALR",
	},
	480: {
		ID:      480,
		Name:    "Bear",
		Neutral: "BAR",
	},
	481: {
		ID:      481,
		Name:    "Scaled Wolf",
		Neutral: "SCW",
	},
	482: {
		ID:      482,
		Name:    "Wolf",
		Neutral: "WLF",
	},
	483: {
		ID:      483,
		Name:    "Spirit Wolf",
		Neutral: "SPW",
	},
	484: {
		ID:      484,
		Name:    "Skeleton",
		Neutral: "SKL",
	},
	485: {
		ID:      485,
		Name:    "Spectre",
		Neutral: "SPT",
	},
	486: {
		ID:      486,
		Name:    "Bolvirk",
		Neutral: "BLV",
	},
	487: {
		ID:     487,
		Name:   "Banshee",
		Female: "BSE",
	},
	488: {
		ID:     488,
		Name:   "Banshee",
		Female: "BSG",
	},
	489: {
		ID:     489,
		Name:   "Elddar",
		Male:   "EEM",
		Female: "EEF",
	},
	490: {
		ID:      490,
		Name:    "Forest Giant",
		Neutral: "GFO",
	},
	491: {
		ID:      491,
		Name:    "Bone Golem",
		Neutral: "GLB",
	},
	492: {
		ID:      492,
		Name:    "Horse",
		Neutral: "HRS",
	},
	493: {
		ID:      493,
		Name:    "Pegasus",
		Neutral: "PGS",
	},
	494: {
		ID:      494,
		Name:    "Shambling Mound",
		Neutral: "SMD",
	},
	495: {
		ID:      495,
		Name:    "Scrykin",
		Neutral: "SRN",
	},
	496: {
		ID:      496,
		Name:    "Treant",
		Neutral: "TRA",
	},
	497: {
		ID:     497,
		Name:   "Vampire",
		Male:   "VAM",
		Female: "VAF",
	},
	498: {
		ID:     498,
		Name:   "Ayonae Ro",
		Female: "ARO",
	},
	499: {
		ID:     499,
		Name:   "Sullon Zek",
		Female: "SZK",
	},
	500: {
		ID:      500,
		Name:    "Banner",
		Neutral: "BNR",
	},
	501: {
		ID:      501,
		Name:    "Flag",
		Neutral: "FLG",
	},
	502: {
		ID:      502,
		Name:    "Rowboat",
		Neutral: "ROW",
	},
	503: {
		ID:      503,
		Name:    "Bear Trap",
		Neutral: "T00",
	},
	504: {
		ID:      504,
		Name:    "Clockwork Bomb",
		Neutral: "T01",
	},
	505: {
		ID:      505,
		Name:    "Dynamite Keg",
		Neutral: "T02",
	},
	506: {
		ID:      506,
		Name:    "Pressure Plate",
		Neutral: "T03",
	},
	507: {
		ID:      507,
		Name:    "Puffer Spore",
		Neutral: "T04",
	},
	508: {
		ID:      508,
		Name:    "Stone Ring",
		Neutral: "T05",
	},
	509: {
		ID:      509,
		Name:    "Root Tentacle",
		Neutral: "T06",
	},
	510: {
		ID:      510,
		Name:    "Runic Symbol",
		Neutral: "T07",
	},
	511: {
		ID:      511,
		Name:    "Saltpetter Bomb",
		Neutral: "T08",
	},
	512: {
		ID:      512,
		Name:    "Floating Skull",
		Neutral: "T09",
	},
	513: {
		ID:      513,
		Name:    "Spike Trap",
		Neutral: "T10",
	},
	514: {
		ID:      514,
		Name:    "Totem",
		Neutral: "T11",
	},
	515: {
		ID:      515,
		Name:    "Web",
		Neutral: "T12",
	},
	516: {
		ID:      516,
		Name:    "Wicker Basket",
		Neutral: "T13",
	},
	517: {
		ID:      517,
		Name:    "Nightmare/Unicorn",
		Neutral: "UNM",
	},
	518: {
		ID:      518,
		Name:    "Horse",
		Neutral: "HRS",
	},
	519: {
		ID:      519,
		Name:    "Nightmare/Unicorn",
		Neutral: "UNM",
	},
	520: {
		ID:      520,
		Name:    "Bixie",
		Neutral: "BXI",
	},
	521: {
		ID:      521,
		Name:    "Centaur",
		Neutral: "CNT",
	},
	522: {
		ID:     522,
		Name:   "Drakkin",
		Male:   "DKM",
		Female: "DKF",
	},
	523: {
		ID:      523,
		Name:    "Giant",
		Neutral: "GFR",
	},
	524: {
		ID:      524,
		Name:    "Gnoll",
		Neutral: "GNL",
	},
	525: {
		ID:      525,
		Name:    "Griffin",
		Neutral: "GRN",
	},
	526: {
		ID:      526,
		Name:    "Giant Shade",
		Neutral: "GFS",
	},
	527: {
		ID:      527,
		Name:    "Harpy",
		Neutral: "HRP",
	},
	528: {
		ID:      528,
		Name:    "Mammoth",
		Neutral: "MTH",
	},
	529: {
		ID:      529,
		Name:    "Satyr",
		Neutral: "SAT",
	},
	530: {
		ID:      530,
		Name:    "Dragon",
		Neutral: "DRG",
	},
	531: {
		ID:      531,
		Name:    "Dragon",
		Neutral: "DRN",
	},
	532: {
		ID:      532,
		Name:    "Dyn'Leth",
		Neutral: "DYN",
	},
	533: {
		ID:      533,
		Name:    "Boat",
		Neutral: "SHI",
	},
	534: {
		ID:      534,
		Name:    "Weapon Rack",
		Neutral: "I00",
	},
	535: {
		ID:      535,
		Name:    "Armor Rack",
		Neutral: "I01",
	},
	536: {
		ID:      536,
		Name:    "Honey Pot",
		Neutral: "I02",
	},
	537: {
		ID:      537,
		Name:    "Jum Jum Bucket",
		Neutral: "I03",
	},
	538: {
		ID:      538,
		Name:    "Toolbox",
		Neutral: "I04",
	},
	539: {
		ID:      539,
		Name:    "Stone Jug",
		Neutral: "I05",
	},
	540: {
		ID:      540,
		Name:    "Small Plant",
		Neutral: "I06",
	},
	541: {
		ID:      541,
		Name:    "Medium Plant",
		Neutral: "I07",
	},
	542: {
		ID:      542,
		Name:    "Tall Plant",
		Neutral: "I08",
	},
	543: {
		ID:      543,
		Name:    "Wine Cask",
		Neutral: "I09",
	},
	544: {
		ID:      544,
		Name:    "Elven Boat",
		Neutral: "B01",
	},
	545: {
		ID:      545,
		Name:    "Gnomish Boat",
		Neutral: "B02",
	},
	546: {
		ID:      546,
		Name:    "Barrel Barge Ship",
		Neutral: "B03",
	},
	547: {
		ID:      547,
		Name:    "Goo",
		Neutral: "GUL",
	},
	548: {
		ID:      548,
		Name:    "Goo",
		Neutral: "GUM",
	},
	549: {
		ID:      549,
		Name:    "Goo",
		Neutral: "GUS",
	},
	550: {
		ID:      550,
		Name:    "Merchant Ship",
		Neutral: "B04",
	},
	551: {
		ID:      551,
		Name:    "Pirate Ship",
		Neutral: "B05",
	},
	552: {
		ID:      552,
		Name:    "Ghost Ship",
		Neutral: "B06",
	},
	553: {
		ID:      553,
		Name:    "Banner",
		Neutral: "G00",
	},
	554: {
		ID:      554,
		Name:    "Banner",
		Neutral: "G01",
	},
	555: {
		ID:      555,
		Name:    "Banner",
		Neutral: "G02",
	},
	556: {
		ID:      556,
		Name:    "Banner",
		Neutral: "G03",
	},
	557: {
		ID:      557,
		Name:    "Banner",
		Neutral: "G04",
	},
	558: {
		ID:      558,
		Name:    "Aviak",
		Neutral: "AVK",
	},
	559: {
		ID:      559,
		Name:    "Beetle",
		Neutral: "BTL",
	},
	560: {
		ID:      560,
		Name:    "Gorilla",
		Neutral: "GRL",
	},
	561: {
		ID:      561,
		Name:    "Kedge",
		Neutral: "KDG",
	},
	562: {
		ID:     562,
		Name:   "Kerran",
		Male:   "KRM",
		Female: "KRF",
	},
	563: {
		ID:      563,
		Name:    "Shissar",
		Neutral: "SHS",
	},
	564: {
		ID:      564,
		Name:    "Siren",
		Neutral: "SIN",
	},
	565: {
		ID:      565,
		Name:    "Sphinx",
		Neutral: "SPX",
	},
	566: {
		ID:     566,
		Name:   "Human",
		Male:   "HNM",
		Female: "HNF",
	},
	567: {
		ID:      567,
		Name:    "Campfire",
		Neutral: "I10",
	},
	568: {
		ID:     568,
		Name:   "Brownie",
		Male:   "BNM",
		Female: "BNF",
	},
	569: {
		ID:      569,
		Name:    "Dragon",
		Neutral: "DRP",
	},
	570: {
		ID:      570,
		Name:    "Exoskeleton",
		Neutral: "EXO",
	},
	571: {
		ID:      571,
		Name:    "Ghoul",
		Neutral: "GHO",
	},
	572: {
		ID:      572,
		Name:    "Clockwork Guardian",
		Neutral: "GUA",
	},
	573: {
		ID:      573,
		Name:    "Mantrap",
		Neutral: "MTP",
	},
	574: {
		ID:      574,
		Name:    "Minotaur",
		Neutral: "MTR",
	},
	575: {
		ID:      575,
		Name:    "Scarecrow",
		Neutral: "SCC",
	},
	576: {
		ID:      576,
		Name:    "Shade",
		Neutral: "SHE",
	},
	577: {
		ID:      577,
		Name:    "Rotocopter",
		Neutral: "SWC",
	},
	578: {
		ID:      578,
		Name:    "Tentacle Terror",
		Neutral: "TNT",
	},
	579: {
		ID:      579,
		Name:    "Wereorc",
		Neutral: "WOK",
	},
	580: {
		ID:      580,
		Name:    "Worg",
		Neutral: "WOR",
	},
	581: {
		ID:      581,
		Name:    "Wyvern",
		Neutral: "WYR",
	},
	582: {
		ID:      582,
		Name:    "Chimera",
		Neutral: "MCH",
	},
	583: {
		ID:      583,
		Name:    "Kirin",
		Neutral: "MKI",
	},
	584: {
		ID:      584,
		Name:    "Puma",
		Neutral: "MPU",
	},
	585: {
		ID:      585,
		Name:    "Boulder",
		Neutral: "I11",
	},
	586: {
		ID:      586,
		Name:    "Banner",
		Neutral: "G05",
	},
	587: {
		ID:     587,
		Name:   "Elven Ghost",
		Male:   "XEM",
		Female: "XEF",
	},
	588: {
		ID:     588,
		Name:   "Human Ghost",
		Male:   "XHM",
		Female: "XHF",
	},
	589: {
		ID:      589,
		Name:    "Chest",
		Neutral: "I12",
	},
	590: {
		ID:      590,
		Name:    "Chest",
		Neutral: "I13",
	},
	591: {
		ID:      591,
		Name:    "Crystal",
		Neutral: "I14",
	},
	592: {
		ID:      592,
		Name:    "Coffin",
		Neutral: "I15",
	},
	593: {
		ID:      593,
		Name:    "Guardian CPU",
		Neutral: "I16",
	},
	594: {
		ID:      594,
		Name:    "Worg",
		Neutral: "MWO",
	},
	595: {
		ID:      595,
		Name:    "Mansion",
		Neutral: "OBJ_BLIMP",
	},
	596: {
		ID:      596,
		Name:    "Floating Island",
		Neutral: "OBP__MELDRATH",
	},
	597: {
		ID:      597,
		Name:    "Cragslither",
		Neutral: "MCS",
	},
	598: {
		ID:      598,
		Name:    "Wrulon",
		Neutral: "MWR",
	},
	599: {
		ID:      599,
		Name:    "Spell Particle 1",
		Neutral: "S01",
	},
	600: {
		ID:     600,
		Name:   "Invisible Man of Zomm",
		Male:   "IVM",
		Female: "IVF",
	},
	601: {
		ID:      601,
		Name:    "Robocopter of Zomm",
		Neutral: "SWC",
	},
	602: {
		ID:      602,
		Name:    "Burynai",
		Neutral: "BUR",
	},
	603: {
		ID:      603,
		Name:    "Frog",
		Neutral: "FGG",
	},
	604: {
		ID:      604,
		Name:    "Dracolich",
		Neutral: "DRL",
	},
	605: {
		ID:      605,
		Name:    "Iksar Ghost",
		Neutral: "XIM",
	},
	606: {
		ID:      606,
		Name:    "Iksar Skeleton",
		Neutral: "SKI",
	},
	607: {
		ID:      607,
		Name:    "Mephit",
		Neutral: "MPH",
	},
	608: {
		ID:      608,
		Name:    "Muddite",
		Neutral: "MUD",
	},
	609: {
		ID:      609,
		Name:    "Raptor",
		Neutral: "RPT",
	},
	610: {
		ID:      610,
		Name:    "Sarnak",
		Neutral: "SRK",
	},
	611: {
		ID:      611,
		Name:    "Scorpion",
		Neutral: "SCO",
	},
	612: {
		ID:      612,
		Name:    "Tsetsian",
		Neutral: "TSE",
	},
	613: {
		ID:      613,
		Name:    "Wurm",
		Neutral: "WRM",
	},
	614: {
		ID:      614,
		Name:    "Nekhon",
		Neutral: "BAL",
	},
	615: {
		ID:      615,
		Name:    "Hydra Crystal",
		Neutral: "HYC",
	},
	616: {
		ID:      616,
		Name:    "Crystal Sphere",
		Neutral: "CRY",
	},
	617: {
		ID:      617,
		Name:    "Gnoll",
		Neutral: "GND",
	},
	618: {
		ID:      618,
		Name:    "Sokokar",
		Neutral: "SOK",
	},
	619: {
		ID:      619,
		Name:    "Stone Pylon",
		Neutral: "PYS",
	},
	620: {
		ID:      620,
		Name:    "Demon Vulture",
		Neutral: "DVL",
	},
	621: {
		ID:      621,
		Name:    "Wagon",
		Neutral: "I17",
	},
	622: {
		ID:      622,
		Name:    "God of Discord",
		Neutral: "GOD",
	},
	623: {
		ID:      623,
		Name:    "Feran Mount",
		Neutral: "MFR",
	},
	624: {
		ID:     624,
		Name:   "Ogre NPC - Male",
		Male:   "ONM",
		Female: "ONF",
	},
	625: {
		ID:      625,
		Name:    "Sokokar Mount",
		Neutral: "MSO",
	},
	626: {
		ID:      626,
		Name:    "Giant (Rallosian mats)",
		Neutral: "GRA",
	},
	627: {
		ID:      627,
		Name:    "Sokokar (w saddle)",
		Neutral: "MSO",
	},
	628: {
		ID:      628,
		Name:    "10th Anniversary Banner",
		Neutral: "BNX",
	},
	629: {
		ID:      629,
		Name:    "10th Anniversary Cake",
		Neutral: "CAK",
	},
	630: {
		ID:      630,
		Name:    "Wine Cask",
		Neutral: "I18",
	},
	631: {
		ID:      631,
		Name:    "Hydra Mount",
		Neutral: "MHY",
	},
	632: {
		ID:      632,
		Name:    "Hydra NPC",
		Neutral: "HYD",
	},
	633: {
		ID:      633,
		Name:    "Wedding Flowers",
		Neutral: "I19",
	},
	634: {
		ID:      634,
		Name:    "Wedding Arbor",
		Neutral: "I20",
	},
	635: {
		ID:      635,
		Name:    "Wedding Altar",
		Neutral: "I21",
	},
	636: {
		ID:      636,
		Name:    "Powder Keg",
		Neutral: "I22",
	},
	637: {
		ID:      637,
		Name:    "Apexus",
		Neutral: "APX",
	},
	638: {
		ID:      638,
		Name:    "Bellikos",
		Neutral: "BEL",
	},
	639: {
		ID:      639,
		Name:    "Brell's First Creation",
		Neutral: "BFC",
	},
	640: {
		ID:      640,
		Name:    "Brell",
		Neutral: "BRE",
	},
	641: {
		ID:      641,
		Name:    "Crystalskin Ambuloid",
		Neutral: "CAM",
	},
	642: {
		ID:      642,
		Name:    "Cliknar Queen",
		Neutral: "CLQ",
	},
	643: {
		ID:      643,
		Name:    "Cliknar Soldier",
		Neutral: "CLS",
	},
	644: {
		ID:      644,
		Name:    "Cliknar Worker",
		Neutral: "CLW",
	},
	645: {
		ID:     645,
		Name:   "Coldain",
		Male:   "CDM",
		Female: "CDF",
	},
	647: {
		ID:      647,
		Name:    "Crystalskin Sessiloid",
		Neutral: "CSE",
	},
	648: {
		ID:      648,
		Name:    "Genari",
		Neutral: "GEN",
	},
	649: {
		ID:      649,
		Name:    "Gigyn",
		Neutral: "GIG",
	},
	650: {
		ID:      650,
		Name:    "Greken - Young Adult",
		Neutral: "GYA",
	},
	651: {
		ID:      651,
		Name:    "Greken - Young",
		Neutral: "GYO",
	},
	652: {
		ID:      652,
		Name:    "Cliknar Mount",
		Neutral: "MCL",
	},
	653: {
		ID:      653,
		Name:    "Telmira",
		Neutral: "TEL",
	},
	654: {
		ID:      654,
		Name:    "Spider Mount",
		Neutral: "MTA",
	},
	655: {
		ID:      655,
		Name:    "Bear Mount",
		Neutral: "MBR",
	},
	656: {
		ID:      656,
		Name:    "Rat Mount",
		Neutral: "MCR",
	},
	657: {
		ID:      657,
		Name:    "Sessiloid Mount",
		Neutral: "MSD",
	},
	658: {
		ID:      658,
		Name:    "Morell Thule",
		Neutral: "LTH",
	},
	659: {
		ID:      659,
		Name:    "Marionette",
		Neutral: "MRT",
	},
	660: {
		ID:      660,
		Name:    "Book Dervish",
		Neutral: "BKD",
	},
	661: {
		ID:      661,
		Name:    "Topiary Lion",
		Neutral: "TPL",
	},
	662: {
		ID:      662,
		Name:    "Rotdog",
		Neutral: "RDG",
	},
	663: {
		ID:      663,
		Name:    "Amygdalan",
		Neutral: "AMY",
	},
	664: {
		ID:      664,
		Name:    "Sandman",
		Neutral: "SND",
	},
	665: {
		ID:      665,
		Name:    "Grandfather Clock",
		Neutral: "GFC",
	},
	666: {
		ID:      666,
		Name:    "Gingerbread Man",
		Neutral: "GBM",
	},
	667: {
		ID:     667,
		Name:   "Royal Guard",
		Male:   "BFR",
		Female: "BFF",
	},
	668: {
		ID:      668,
		Name:    "Rabbit",
		Neutral: "BNY",
	},
	669: {
		ID:      669,
		Name:    "Blind Dreamer",
		Neutral: "BDR",
	},
	670: {
		ID:      670,
		Name:    "Cazic Thule",
		Neutral: "CTH",
	},
	671: {
		ID:      671,
		Name:    "Topiary Lion Mount",
		Neutral: "MTL",
	},
	672: {
		ID:      672,
		Name:    "Rot Dog Mount",
		Neutral: "MRD",
	},
	673: {
		ID:      673,
		Name:    "Goral Mount",
		Neutral: "MGL",
	},
	674: {
		ID:      674,
		Name:    "Selyrah Mount",
		Neutral: "MSL",
	},
	675: {
		ID:      675,
		Name:    "Sclera Mount",
		Neutral: "MSC",
	},
	676: {
		ID:      676,
		Name:    "Braxi Mount",
		Neutral: "MBX",
	},
	677: {
		ID:      677,
		Name:    "Kangon Mount",
		Neutral: "MKG",
	},
	678: {
		ID:   678,
		Name: "Erudite",
		Male: "ERU",
	},
	679: {
		ID:      679,
		Name:    "Wurm Mount",
		Neutral: "MWM",
	},
	680: {
		ID:      680,
		Name:    "Raptor Mount",
		Neutral: "MRP",
	},
	681: {
		ID:      681,
		Name:    "Invisible Man",
		Neutral: "INV",
	},
	682: {
		ID:      682,
		Name:    "Whirligig",
		Neutral: "MWI",
	},
	683: {
		ID:      683,
		Name:    "Gnomish Balloon",
		Neutral: "MBL",
	},
	684: {
		ID:      684,
		Name:    "Gnomish Rocket Pack",
		Neutral: "MRK",
	},
	685: {
		ID:      685,
		Name:    "Gnomish Hovering Transport",
		Neutral: "MHB",
	},
	686: {
		ID:      686,
		Name:    "Selyrah",
		Neutral: "SEY",
	},
	687: {
		ID:      687,
		Name:    "Goral",
		Neutral: "G}1",
	},
	688: {
		ID:      688,
		Name:    "Braxi",
		Neutral: "BRX",
	},
	689: {
		ID:      689,
		Name:    "Kangon",
		Neutral: "KNG",
	},
	690: {
		ID:      690,
		Name:    "Invisible Man",
		Neutral: "I23",
	},
	691: {
		ID:      691,
		Name:    "Floating Tower",
		Neutral: "B07",
	},
	692: {
		ID:      692,
		Name:    "Explosive Cart",
		Neutral: "B08",
	},
	693: {
		ID:      693,
		Name:    "Blimp Ship",
		Neutral: "B09",
	},
	694: {
		ID:      694,
		Name:    "Tumbleweed",
		Neutral: "I24",
	},
	695: {
		ID:      695,
		Name:    "Alaran",
		Neutral: "ALA",
	},
	696: {
		ID:      696,
		Name:    "Swinetor",
		Neutral: "SWI",
	},
	697: {
		ID:      697,
		Name:    "Triumvirate",
		Neutral: "TRG",
	},
	698: {
		ID:      698,
		Name:    "Hadal",
		Neutral: "HDL",
	},
	699: {
		ID:      699,
		Name:    "Hovering Platform",
		Neutral: "B10",
	},
	700: {
		ID:      700,
		Name:    "Parasitic Scavenger",
		Neutral: "PSC",
	},
	701: {
		ID:      701,
		Name:    "Grendlaen",
		Neutral: "GRD",
	},
	702: {
		ID:      702,
		Name:    "Ship in a Bottle",
		Neutral: "I25",
	},
	703: {
		ID:      703,
		Name:    "Alaran Sentry Stone",
		Neutral: "I26",
	},
	704: {
		ID:      704,
		Name:    "Dervish",
		Neutral: "SDV",
	},
	705: {
		ID:      705,
		Name:    "Regeneration Pool",
		Neutral: "I27",
	},
	706: {
		ID:      706,
		Name:    "Teleportation Stand",
		Neutral: "I28",
	},
	707: {
		ID:      707,
		Name:    "Relic Case",
		Neutral: "I29",
	},
	708: {
		ID:      708,
		Name:    "Alaran Ghost",
		Neutral: "ALG",
	},
	709: {
		ID:      709,
		Name:    "Skystrider",
		Neutral: "MPG",
	},
	710: {
		ID:      710,
		Name:    "Water Spout",
		Neutral: "I30",
	},
	711: {
		ID:      711,
		Name:    "Aviak Pull Along",
		Neutral: "I31",
	},
	712: {
		ID:      712,
		Name:    "Gelatinous Cube",
		Neutral: "GCB",
	},
	713: {
		ID:      713,
		Name:    "Cat",
		Neutral: "CAT",
	},
	714: {
		ID:      714,
		Name:    "Elk Head",
		Neutral: "I32",
	},
	715: {
		ID:      715,
		Name:    "Holgresh",
		Neutral: "HLG",
	},
	716: {
		ID:      716,
		Name:    "Beetle",
		Neutral: "BTS",
	},
	717: {
		ID:      717,
		Name:    "Vine Maw",
		Neutral: "VNM",
	},
	718: {
		ID:      718,
		Name:    "Ratman",
		Neutral: "CHT",
	},
	719: {
		ID:      719,
		Name:    "Fallen Knight",
		Neutral: "FKN",
	},
	720: {
		ID:      720,
		Name:    "Flying Carpet",
		Neutral: "MCP",
	},
	721: {
		ID:      721,
		Name:    "Carrier Hand",
		Neutral: "MRH",
	},
	722: {
		ID:     722,
		Name:   "Akheva",
		Male:   "AHM",
		Female: "AHF",
	},
	723: {
		ID:   723,
		Name: "Servant of Shadow",
		Male: "SOS",
	},
	724: {
		ID:     724,
		Name:   "Luclin",
		Female: "LUC",
	},
	725: {
		ID:      725,
		Name:    "Xaric the Unspoken",
		Neutral: "AXA",
	},
	726: {
		ID:      726,
		Name:    "Dervish (Ver. 5)",
		Neutral: "HDV",
	},
	727: {
		ID:      727,
		Name:    "Dervish (Ver. 6)",
		Neutral: "DV6",
	},
	728: {
		ID:      728,
		Name:    "God - Luclin (Ver. 2)",
		Neutral: "LU2",
	},
	729: {
		ID:      729,
		Name:    "God - Luclin (Ver. 3)",
		Neutral: "LU3",
	},
	730: {
		ID:      730,
		Name:    "Orb",
		Neutral: "ORB",
	},
	731: {
		ID:      731,
		Name:    "God - Luclin (Ver. 4)",
		Neutral: "LU4",
	},
	732: {
		ID:      732,
		Name:    "Pegasus",
		Neutral: "PG3",
	},
	2253: {
		ID:      2253,
		Name:    "(null)",
		Male:    "PPOINT",
		Female:  "PPOINT",
		Neutral: "PPOINT",
	},
	2254: {
		ID:      2254,
		Name:    "(null)",
		Male:    "PPOINT",
		Female:  "PPOINT",
		Neutral: "PPOINT",
	},
	2255: {
		ID:      2255,
		Name:    "(null)",
		Male:    "PPOINT",
		Female:  "PPOINT",
		Neutral: "PPOINT",
	},
	2256: {
		ID:      2256,
		Name:    "(null)",
		Male:    "PPOINT",
		Female:  "PPOINT",
		Neutral: "PPOINT",
	},
	2257: {
		ID:      2257,
		Name:    "(null)",
		Male:    "PPOINT",
		Female:  "PPOINT",
		Neutral: "PPOINT",
	},
	2258: {
		ID:      2258,
		Name:    "(null)",
		Male:    "PPOINT",
		Female:  "PPOINT",
		Neutral: "PPOINT",
	},
	2259: {
		ID:      2259,
		Name:    "(null)",
		Male:    "ARROW",
		Female:  "ARROW",
		Neutral: "ARROW",
	},
}
