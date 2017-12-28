package model

import (
	"database/sql"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Zone struct {
	ShortName         sql.NullString `json:"shortName" db:"short_name"`                  //short_name` varchar(32) DEFAULT NULL,
	Id                int64          `json:"id" db:"id"`                                 //id` int(10) NOT NULL AUTO_INCREMENT,
	FileName          sql.NullString `json:"fileName" db:"file_name"`                    //file_name` varchar(16) DEFAULT NULL,
	LongName          string         `json:"longName" db:"long_name"`                    //long_name` text NOT NULL,
	MapFileName       sql.NullString `json:"mapFileName" db:"map_file_name"`             //map_file_name` varchar(100) DEFAULT NULL,
	SafeX             float64        `json:"safeX" db:"safe_x"`                          //safe_x` float NOT NULL DEFAULT '0',
	SafeT             float64        `json:"safeY" db:"safe_y"`                          //safe_y` float NOT NULL DEFAULT '0',
	SafeZ             float64        `json:"safeZ" db:"safe_z"`                          //safe_z` float NOT NULL DEFAULT '0',
	GraveyardId       float64        `json:"graveyardId" db:"graveyard_id"`              //graveyard_id` float NOT NULL DEFAULT '0',
	MinLevel          int64          `json:"minLevel" db:"min_level"`                    //min_level` tinyint(3) unsigned NOT NULL DEFAULT '0',
	MinStatus         int64          `json:"minStatus" db:"min_status"`                  //min_status` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ZoneIdNumber      int64          `json:"zoneidnumber" db:"zoneidnumber"`             //zoneidnumber` int(4) NOT NULL DEFAULT '0',
	Version           int64          `json:"version" db:"version"`                       //version` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Timezone          int64          `json:"timezone" db:"timezone"`                     //timezone` int(5) NOT NULL DEFAULT '0',
	MaxClients        int64          `json:"maxclients" db:"maxclients"`                 //maxclients` int(5) NOT NULL DEFAULT '0',
	Ruleset           int64          `json:"ruleset" db:"ruleset"`                       //ruleset` int(10) unsigned NOT NULL DEFAULT '0',
	Note              sql.NullString `json:"note" db:"note"`                             //note` varchar(80) DEFAULT NULL,
	Underworld        float64        `json:"underworld" db:"underworld"`                 //underworld` float NOT NULL DEFAULT '0',
	Minclip           float64        `json:"minclip" db:"minclip"`                       //minclip` float NOT NULL DEFAULT '450',
	Maxclip           float64        `json:"maxclip" db:"maxclip"`                       //maxclip` float NOT NULL DEFAULT '450',
	FogMinclip        float64        `json:"fogMinclip" db:"fog_minclip"`                //fog_minclip` float NOT NULL DEFAULT '450',
	FogMaxclip        float64        `json:"fogMaxclip" db:"fog_maxclip"`                //fog_maxclip` float NOT NULL DEFAULT '450',
	FogBlue           int64          `json:"fogBlue" db:"fog_blue"`                      //fog_blue` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogRed            int64          `json:"fogRed" db:"fog_red"`                        //fog_red` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen          int64          `json:"fogGreen" db:"fog_green"`                    //fog_green` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Sky               int64          `json:"sky" db:"sky"`                               //sky` tinyint(3) unsigned NOT NULL DEFAULT '1',
	Ztype             int64          `json:"ztype" db:"ztype"`                           //ztype` tinyint(3) unsigned NOT NULL DEFAULT '1',
	ZoneExpMultiplier float64        `json:"zoneExpMultiplier" db:"zone_exp_multiplier"` //zone_exp_multiplier` decimal(6,2) NOT NULL DEFAULT '0.00',
	Walkspeed         float64        `json:"walkspeed" db:"walkspeed"`                   //walkspeed` float NOT NULL DEFAULT '0.4',
	TimeType          int64          `json:"timeType" db:"time_type"`                    //time_type` tinyint(3) unsigned NOT NULL DEFAULT '2',
	FogRed1           int64          `json:"fogRed1" db:"fog_red1"`                      //fog_red1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen1         int64          `json:"fogGreen1" db:"fog_green1"`                  //fog_green1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue1          int64          `json:"fogBlue1" db:"fog_blue1"`                    //fog_blue1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip1       float64        `json:"fogMinclip1" db:"fog_minclip1"`              //fog_minclip1` float NOT NULL DEFAULT '450',
	FogMaxclip1       float64        `json:"fogMaxclip1" db:"fog_maxclip1"`              //fog_maxclip1` float NOT NULL DEFAULT '450',
	FogRed2           int64          `json:"fogRed2" db:"fog_red2"`                      //fog_red2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen2         int64          `json:"fogGreen2" db:"fog_green2"`                  //fog_green2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue2          int64          `json:"fogBlue2" db:"fog_blue2"`                    //fog_blue2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip2       float64        `json:"fogMinclip2" db:"fog_minclip2"`              //fog_minclip2` float NOT NULL DEFAULT '450',
	FogMaxclip2       float64        `json:"fogMaxclip2" db:"fog_maxclip2"`              //fog_maxclip2` float NOT NULL DEFAULT '450',
	FogRed3           int64          `json:"fogRed3" db:"fog_red3"`                      //fog_red3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen3         int64          `json:"fogGreen3" db:"fog_green3"`                  //fog_green3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue3          int64          `json:"fogBlue3" db:"fog_blue3"`                    //fog_blue3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip3       float64        `json:"fogMinclip3" db:"fog_minclip3"`              //fog_minclip3` float NOT NULL DEFAULT '450',
	FogMaxclip3       float64        `json:"fogMaxclip3" db:"fog_maxclip3"`              //fog_maxclip3` float NOT NULL DEFAULT '450',
	FogRed4           int64          `json:"fogRed4" db:"fog_red4"`                      //fog_red4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen4         int64          `json:"fogGreen4" db:"fog_green4"`                  //fog_green4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue4          int64          `json:"fogBlue4" db:"fog_blue4"`                    //fog_blue4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip4       float64        `json:"fogMinclip4" db:"fog_minclip4"`              //fog_minclip4` float NOT NULL DEFAULT '450',
	FogMaxclip4       float64        `json:"fogMaxclip4" db:"fog_maxclip4"`              //fog_maxclip4` float NOT NULL DEFAULT '450',
	FogDensity        float64        `json:"fogDensity" db:"fog_density"`                //fog_density` float NOT NULL DEFAULT '0',
	FlagNeeded        string         `json:"flagNeeded" db:"flag_needed"`                //flag_needed` varchar(128) NOT NULL DEFAULT '',
	Canbind           int64          `json:"canbind" db:"canbind"`                       //canbind` tinyint(4) NOT NULL DEFAULT '1',
	Cancombat         int64          `json:"cancombat" db:"cancombat"`                   //cancombat` tinyint(4) NOT NULL DEFAULT '1',
	Canlevitate       int64          `json:"canlevitate" db:"canlevitate"`               //canlevitate` tinyint(4) NOT NULL DEFAULT '1',
	Castoutdoor       int64          `json:"castoutdoor" db:"castoutdoor"`               //castoutdoor` tinyint(4) NOT NULL DEFAULT '1',
	Hotzone           int64          `json:"hotzone" db:"hotzone"`                       //hotzone` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Insttype          int64          `json:"insttype" db:"insttype"`                     //insttype` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
	Shutdowndelay     int64          `json:"shutdowndelay" db:"shutdowndelay"`           //shutdowndelay` bigint(16) unsigned NOT NULL DEFAULT '5000',
	Peqzone           int64          `json:"peqzone" db:"peqzone"`                       //peqzone` tinyint(4) NOT NULL DEFAULT '1',
	Expansion         int64          `json:"expansion" db:"expansion"`                   //expansion` tinyint(3) NOT NULL DEFAULT '0',
	Suspendbuffs      int64          `json:"suspendbuffs" db:"suspendbuffs"`             //suspendbuffs` tinyint(1) unsigned NOT NULL DEFAULT '0',
	RainChance1       int64          `json:"rainChance1" db:"rain_chance1"`              //rain_chance1` int(4) NOT NULL DEFAULT '0',
	RainChance2       int64          `json:"rainChance2" db:"rain_chance2"`              //rain_chance2` int(4) NOT NULL DEFAULT '0',
	RainChance3       int64          `json:"rainChance3" db:"rain_chance3"`              //rain_chance3` int(4) NOT NULL DEFAULT '0',
	RainChance4       int64          `json:"rainChance4" db:"rain_chance4"`              //rain_chance4` int(4) NOT NULL DEFAULT '0',
	RainDuration1     int64          `json:"rainDuration1" db:"rain_duration1"`          //rain_duration1` int(4) NOT NULL DEFAULT '0',
	RainDuration2     int64          `json:"rainDuration2" db:"rain_duration2"`          //rain_duration2` int(4) NOT NULL DEFAULT '0',
	RainDuration3     int64          `json:"rainDuration3" db:"rain_duration3"`          //rain_duration3` int(4) NOT NULL DEFAULT '0',
	RainDuration4     int64          `json:"rainDuration4" db:"rain_duration4"`          //rain_duration4` int(4) NOT NULL DEFAULT '0',
	SnowChance1       int64          `json:"snowChance1" db:"snow_chance1"`              //snow_chance1` int(4) NOT NULL DEFAULT '0',
	SnowChance2       int64          `json:"snowChance2" db:"snow_chance2"`              //snow_chance2` int(4) NOT NULL DEFAULT '0',
	SnowChance3       int64          `json:"snowChance3" db:"snow_chance3"`              //snow_chance3` int(4) NOT NULL DEFAULT '0',
	SnowChance4       int64          `json:"snowChance4" db:"snow_chance4"`              //snow_chance4` int(4) NOT NULL DEFAULT '0',
	SnowDuration1     int64          `json:"snowDuration1" db:"snow_duration1"`          //snow_duration1` int(4) NOT NULL DEFAULT '0',
	SnowDuration2     int64          `json:"snowDuration2" db:"snow_duration2"`          //snow_duration2` int(4) NOT NULL DEFAULT '0',
	SnowDuration3     int64          `json:"snowDuration3" db:"snow_duration3"`          //snow_duration3` int(4) NOT NULL DEFAULT '0',
	SnowDuration4     int64          `json:"snowDuration4" db:"snow_duration4"`          //snow_duration4` int(4) NOT NULL DEFAULT '0',
	Gravity           float64        `json:"gravity" db:"gravity"`                       //gravity` float NOT NULL DEFAULT '0.4',
	Type              int64          `json:"type" db:"type"`                             //type` int(3) NOT NULL DEFAULT '0',
	Skylock           int64          `json:"skylock" db:"skylock"`                       //skylock` tinyint(4) NOT NULL DEFAULT '0',
}

func (c *Zone) GetMinStatusName() string {
	switch {
	case c.MinStatus >= 200:
		return "Admin"
	case c.MinStatus >= 100:
		return "Guide"
	}
	return ""
}

func (c *Zone) ExpansionId() int64 {
	switch c.ExpansionBit() {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 4:
		return 3
	case 8:
		return 4
	case 16:
		return 5
	case 32:
		return 6
	case 64:
		return 7
	case 127:
		return 8
	case 256:
		return 9
	case 512:
		return 10
	case 1024:
		return 11
	case 2048:
		return 12
	case 4096:
		return 13
	case 8192:
		return 14
	case 16384:
		return 15 // - seeds of destruction
	case 32768: //underfoot
		return 16
	case 65536: //house of thule
		return 17
	case 131072: //veil of alaris
		return 18
	case 262144: //rain of fear
		return 19
	case 524288: //call of the forsaken
		return 20
	case 1048576: //the darkened sea
		return 21
	case 2097152: //the broken mirror
		return 22
	case 4194304: //empires of kunark
		return 23
	case 8388608: //ring of scale
		return 24
	default:
		return -1
	}
	return -1
}

/*
//https://articles.eqresource.com/zonelist.php
"abysmal"
"acrylia"
"airplane"
"akanon"
"akheva"
"anguish"
"apprentice"
"arcstone"
"arena"
"arena2"
"arttest"
"ashengate"
"atiiki"
"aviak"
"barindu"
"barren"
"barter"
"bazaar"
"befallen"
"befallenb"
"beholder"
"bertoxtemple"
"blackburrow"
"blacksail"
"bloodfields"
"bloodmoon"
"bothunder"
"broodlands"
"buriedsea"
"burningwood"
"butcher"
"cabeast"
"cabwest"
"cauldron"
"causeway"
"cazicthule"
"chambersa"
"chambersb"
"chambersc"
"chambersd"
"chamberse"
"chambersf"
"charasis"
"chardok"
"chardokb"
"citymist"
"clz"
"cobaltscar"
"codecay"
"commonlands"
"commons"
"corathus"
"corathusa"
"corathusb"
"crescent"
"crushbone"
"cryptofshade"
"crystal"
"crystallos"
"cshome"
"dalnir"
"dawnshroud"
"deadbone"
"delvea"
"delveb"
"devastationa"
"devastation"
"direwind"
"discord"
"discordtower"
"drachnidhive"
"drachnidhivea"
"drachnidhiveb"
"drachnidhivec"
"dragonscale"
"dragonscaleb"
"dranik"
"dranikcatacombsa"
"dranikcatacombsb"
"dranikcatacombsc"
"dranikhollowsa"
"dranikhollowsb"
"dranikhollowsc"
"draniksewersa"
"draniksewersb"
"draniksewersc"
"draniksscar"
"dreadlands"
"dreadspire"
"droga"
"dulak"
"eastkarana"
"eastkorlach"
"eastkorlacha"
"eastwastes"
"echo"
"ecommons"
"elddar"
"elddara"
"emeraldjungle"
"erudnext"
"erudnint"
"erudsxing"
"erudsxing2"
"everfrost"
"fearplane"
"feerrott"
"felwithea"
"felwitheb"
"ferubi"
"fhalls"
"fieldofbone"
"firiona"
"freeportacademy"
"freeportarena"
"freeportcityhall"
"freeporteast"
"freeporthall"
"freeportmilitia"
"freeportsewers"
"freeporttemple"
"freeporttheater"
"freeportwest"
"freporte"
"freportn"
"freportw"
"frontiermtns"
"frostcrypt"
"frozenshadow"
"fungusgrove"
"gfaydark"
"greatdivide"
"griegsend"
"grimling"
"grobb"
"growthplane"
"guardian"
"guildhall"
"guildlobby"
"guka"
"gukb"
"gukbottom"
"gukc"
"gukd"
"guke"
"gukf"
"gukg"
"gukh"
"guktop"
"gunthak"
"gyrospireb"
"gyrospirez"
"halas"
"harbingers"
"hateplane"
"hateplaneb"
"hatesfury"
"highkeep"
"highpass"
"highpasshold"
"highpasskeep"
"hillsofshade"
"hohonora"
"hohonorb"
"hole"
"hollowshade"
"iceclad"
"icefall"
"ikkinz"
"illsalin"
"illsalina"
"illsalinb"
"illsalinc"
"inktuta"
"innothule"
"innothuleb"
"jaggedpine"
"jardelshook"
"kael"
"kaesora"
"kaladima"
"kaladimb"
"karnor"
"katta"
"kattacastrum"
"kedge"
"kerraridge"
"kithforest"
"kithicor"
"kodtaz"
"korascian"
"kurn"
"lakeofillomen"
"lakerathe"
"lavastorm"
"letalis"
"lfaydark"
"load"
"load2"
"lopingplains"
"maiden"
"maidensgrave"
"mansion"
"mechanotus"
"mesa"
"mira"
"mirb"
"mirc"
"mird"
"mire"
"mirf"
"mirg"
"mirh"
"miri"
"mirj"
"mischiefplane"
"mistmoore"
"misty"
"mistythicket"
"mmca"
"mmcb"
"mmcc"
"mmcd"
"mmce"
"mmcf"
"mmcg"
"mmch"
"mmci"
"mmcj"
"monkeyrock"
"moors"
"mseru"
"nadox"
"najena"
"natimbi"
"necropolis"
"nedaria"
"nektropos"
"nektulos"
"nektulosa"
"neriaka"
"neriakb"
"neriakc"
"neriakd"
"netherbian"
"nexus"
"nightmareb"
"northkarana"
"northro"
"nro"
"nurga"
"oasis"
"oceangreenhills"
"oceangreenvillage"
"oceanoftears"
"oggok"
"oldblackburrow"
"oldbloodfield"
"oldcommons"
"olddranik"
"oldfieldofbone"
"oldhighpass"
"oldkaesoraa"
"oldkaesorab"
"oldkithicor"
"oldkurn"
"oot"
"overthere"
"paineel"
"paludal"
"paw"
"permafrost"
"poair"
"podisease"
"poeartha"
"poearthb"
"pofire"
"poinnovation"
"pojustice"
"poknowledge"
"ponightmare"
"postorms"
"potactics"
"potimea"
"potimeb"
"potorment"
"potranquility"
"povalor"
"powar"
"powater"
"precipiceofwar"
"provinggrounds"
"qcat"
"qey2hh1"
"qeynos"
"qeynos2"
"qeytoqrg"
"qinimi"
"qrg"
"qvic"
"qvicb"
"rage"
"ragea"
"rathechamber"
"rathemtn"
"redfeather"
"relic"
"riftseekers"
"rivervale"
"riwwi"
"roost"
"ruja"
"rujb"
"rujc"
"rujd"
"ruje"
"rujf"
"rujg"
"rujh"
"ruji"
"rujj"
"runnyeye"
"scarlet"
"sebilis"
"shadeweaver"
"shadowhaven"
"shadowrest"
"shadowspine"
"sharvahl"
"shipmvm"
"shipmvp"
"shipmvu"
"shippvu"
"shipuvu"
"shipworkshop"
"silyssar"
"sirens"
"skyfire"
"skylance"
"skyshrine"
"sleeper"
"sncrematory"
"snlair"
"snplant"
"snpool"
"soldunga"
"soldungb"
"soldungc"
"solrotower"
"soltemple"
"solteris"
"southkarana"
"southro"
"sro"
"sseru"
"ssratemple"
"steamfactory"
"steamfont"
"steamfontmts"
"steppes"
"stillmoona"
"stillmoonb"
"stonebrunt"
"stonehive"
"suncrest"
"sunderock"
"swampofnohope"
"tacvi"
"taka"
"takb"
"takc"
"takd"
"take"
"takf"
"takg"
"takh"
"taki"
"takishruins"
"takishruinsa"
"takj"
"templeveeshan"
"tenebrous"
"thalassius"
"theater"
"theatera"
"thedeep"
"thegrey"
"thenest"
"thevoida"
"thevoidb"
"thevoidc"
"thevoidd"
"thevoide"
"thevoidf"
"thevoidg"
"thundercrest"
"thurgadina"
"thurgadinb"
"timorous"
"tipt"
"torgiran"
"toskirakk"
"tox"
"toxxulia"
"trakanon"
"tutorial"
"tutoriala"
"tutorialb"
"twilight"
"txevu"
"umbral"
"unrest"
"uqua"
"valdeholm"
"veeshan"
"veksar"
"velketor"
"vergalid"
"vexthal"
*/

func (c *Zone) ExpansionBit() int64 {
	switch c.ShortName.String {
	case "xorbb",
		"weddingchapeldark",
		"weddingchapel",
		"nektulos",
		"warrens":
		return 0 // - classic
	case "overthere",
		"warslikswood":
		return 1 // - ruins of kunark
	case "thurgadina",
		"westwastes",
		"wakening":
		return 2 // - scars of velious
	case "acrylia", "akheva", "bazaar":
		return 4 // - shadows of luclin
	case "poknowledge":
		return 8 // - planes of power
	case "nadox":
		return 16 // - legacy of ykesha
	case "mira":
		return 32 // - lost dungeons of norrath
	case "wallofslaughter", "yxtta", "vxed":
		return 64 // - gates of discord
	case "abysmal", "anguish":
		return 128 // - omens of war
	case "stillmoona":
		return 256 // - dragons of norrath
	case
		"westkorlach",
		"westkorlacha",
		"westkorlachb",
		"westkorlachc":
		return 512 // - depths of darkhallow
	case "asdf":
		return 1024 // - prophecy of ro
	case "asdfg":
		return 2048 // - serpent's spine
	case "zhisza":
		return 4096 // - the buried sea
	case "hdshd":
		return 8192 // - secrets of faydwer
	case "fdo":
		return 16384 // - seeds of destruction
	case "shiningcity",
		"lichencreep",
		"fungalforest",
		"brellstemple",
		"dragoncrypt",
		"brellsrest",
		"coolingchamber",
		"pellucid",
		"arthicrex",
		"foundation",
		"underquarry",
		"brellsarena",
		"stonesnake",
		"convorteum":
		return 32768 //underfoot
	case "thulelibrary",
		"morellcastle",
		"alkabormare",
		"fallen",
		"feerrott2",
		"housegarden",
		"miragulmare",
		"somnium",
		"thuledream",
		"thulehouse1",
		"thulehouse2",
		"well",
		"neighborhood":
		return 65536 //house of thule
	case "westsepulcher",
		"eastsepulcher",
		"sepulcher",
		"cityofbronze",
		"windsong",
		"argath",
		"arelis",
		"sarithcity",
		"rubak",
		"beastdomain",
		"resplendent",
		"pillarsalra":
		return 131072 //veil of alaris
	case "shardslanding",
		"kaelshard",
		"grelleth",
		"eviltree",
		"eastwastesshard",
		"crystalshard",
		"chapterhouse",
		"breedinggrounds":
		return 262144 //rain of fear
	case "1222":
		return 524288 //call of the forsaken
	case "fff":
		return 1048576 //the darkened sea
	case "sdff":
		return 2097152 //the broken mirror
	case "111":
		return 4194304 //empires of kunark
	case "asdssf":
		return 8388608 //ring of scale
	default:
		return -1
	}
	return -1
}

func (c *Zone) ExpansionName() string {
	switch c.ExpansionId() {
	case -1:
		return "Unknown"
	case 0:
		return "Classic"
	case 1:
		return "Ruins of Kunark"
	case 2:
		return "Scars of Velious"
	case 3:
		return "Shadows of Luclin"
	case 4:
		return "Planes of Power"
	case 5:
		return "Legacy of Ykesha"
	case 6:
		return "Lost Dungeons of Norrath"
	case 7:
		return "Gates of Discord"
	case 8:
		return "Omens of War"
	case 9:
		return "Dragons of Norrath"
	case 10:
		return "Depths of Darkhallow"
	case 11:
		return "Prophecy of Ro"
	case 12:
		return "Serpent's Spine"
	case 13:
		return "The Buried Sea"
	case 14:
		return "Secrets of Faydwer"
	}
	return "Unknown"
}

func (c *Zone) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
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

func (c *Zone) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
