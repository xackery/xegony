//Package model provides data and shared structures
package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	//WAR Warrior
	WAR = 1 << 0
	//CLR Cleric
	CLR = 1 << 1
	//PAL Paladin
	PAL = 1 << 2
	//RNG Ranger
	RNG = 1 << 3
	//SHD Shadow Knight
	SHD = 1 << 4
	//DRU Druid
	DRU = 1 << 5
	//MNK Monk
	MNK = 1 << 6
	//BRD Bard
	BRD = 1 << 7
	//ROG Rogue
	ROG = 1 << 8
	//SHM Shaman
	SHM = 1 << 9
	//NEC Necromancer
	NEC = 1 << 10
	//WIZ Wizard
	WIZ = 1 << 11
	//MAG Magician
	MAG = 1 << 12
	//ENC Enchanter
	ENC = 1 << 13
	//BST Beastlord
	BST = 1 << 14
	//BER Berserker
	BER = 1 << 15
	//ALL All
	ALL = 0xFFFF
)

var (
	isRulesLoaded bool

	deities = map[int64]string{
		0:  "Agnostic",
		1:  "Bertoxxulous",
		2:  "Brell Serilis",
		3:  "Cazic Thule",
		4:  "Erollisi Marr",
		5:  "Bristlebane",
		6:  "Innoruuk",
		7:  "Karana",
		8:  "Mithaniel Marr",
		9:  "Prexus",
		10: "Quellious",
		11: "Rallos Zek",
		12: "Rodcet Nife",
		13: "Solusek Ro",
		14: "The Tribunal",
		15: "Tunare",
		16: "Veeshan",
	}
	skills = map[int64]string{
		0:  "1H Blunt",
		1:  "1H Slashing",
		2:  "2H Blunt",
		3:  "2H Slashing",
		4:  "Abjuration",
		5:  "Alteration",
		6:  "Apply Poison",
		7:  "Archery",
		8:  "Backstab",
		9:  "Bind Wound",
		10: "Bash",
		11: "Block",
		12: "Brass Instruments",
		13: "Channeling",
		14: "Conjuration",
		15: "Defense",
		16: "Disarm",
		17: "Disarm Traps",
		18: "Divination",
		19: "Dodge",
		20: "Double Attack",
		21: "Dragon Punch",
		22: "Duel Wield",
		23: "Eagle Strike",
		24: "Evocation",
		25: "Feign Death",
		26: "Flying Kick",
		27: "Forage",
		28: "Hand To Hand",
		29: "Hide",
		30: "Kick",
		31: "Meditate",
		32: "Mend",
		33: "Offense",
		34: "Parry",
		35: "Pick Lock",
		36: "Piercing",
		37: "Riposte",
		38: "Round Kick",
		39: "Safe Fall",
		40: "Sense Heading",
		41: "Sing",
		42: "Sneak",
		43: "Specialize Abjure",
		44: "Specialize Alteration",
		45: "Specialize Conjuration",
		46: "Specialize Divinatation",
		47: "Specialize Evocation",
		48: "Pick Pockets",
		49: "Stringed Instruments",
		50: "Swimming",
		51: "Throwing",
		52: "Tiger Claw",
		53: "Tracking",
		54: "Wind Instruments",
		55: "Fishing",
		56: "Make Poison",
		57: "Tinkering",
		58: "Research",
		59: "Alchemy",
		60: "Baking",
		61: "Tailoring",
		62: "Sense Traps",
		63: "Blacksmithing",
		64: "Fletching",
		65: "Brewing",
		66: "Alcohol Tolerance",
		67: "Begging",
		68: "Jewelry Making",
		69: "Pottery",
		70: "Percussion Instruments",
		71: "Intimidation",
		72: "Berserking",
		73: "Taunt",
		74: "Frenzy",
		75: "Remove Traps",
		76: "Triple Attack",
		77: "2H Piercing",
	}
)

//CashName returns human readable cash
func CashName(money int64) string {
	amount := ""

	val := money / 1000
	if val > 0 {
		amount += fmt.Sprintf("%dpp ", val)
		money -= val * 1000
	}

	val = money / 100
	if val > 0 {
		amount += fmt.Sprintf("%dgp ", val)
		money -= val * 100
	}

	val = money / 10
	if val > 0 {
		amount += fmt.Sprintf("%dsp ", val)
		money -= val * 10
	}
	val = money / 1
	if val > 0 {
		amount += fmt.Sprintf("%dcp ", val)
	}

	if len(amount) > 0 {
		amount = amount[0 : len(amount)-1]
	}

	return amount
}

//SkillName returns the skill human readable name based on ID
func SkillName(id int64) string {
	skill, ok := skills[id]
	if ok {
		return skill
	}
	return "Unknown"
}

//CleanName returns sanitized names
func CleanName(name string) string {
	var re = regexp.MustCompile(`[^0-9A-Za-z_]+`)
	cleanName := strings.Replace(name, " ", "_", -1)
	cleanName = strings.Replace(cleanName, "#", "", -1)
	cleanName = strings.TrimSpace(re.ReplaceAllString(cleanName, ""))
	cleanName = strings.Replace(cleanName, "_", " ", -1)
	return cleanName
}

//RuleR grabs float64 version of rule table value
func RuleR(name string) float64 {

	val := GetRule(name)
	fVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Println("Rule", name, "was attempted to be parsed as float (RuleR) but failed")
	}
	return fVal
}

//GetRule grabs string version of rule table value
func GetRule(name string) string {
	switch name {
	case "Character:ExpMultiplier":
		return "2.0000000000000"
	case "Zone:HotZoneBonus":
		return "0.5000000000000"
	}
	return ""
}

//ZoneName returns human readable name
func ZoneName(zoneID int64) string {
	switch zoneID {
	case 1:
		return "qeynos"
	case 2:
		return "qeynos2"
	case 3:
		return "qrg"
	case 4:
		return "qeytoqrg"
	case 5:
		return "highpass"
	case 6:
		return "highkeep"
	case 8:
		return "freportn"
	case 9:
		return "freportw"
	case 10:
		return "freporte"
	case 11:
		return "runnyeye"
	case 12:
		return "qey2hh1"
	case 13:
		return "northkarana"
	case 14:
		return "southkarana"
	case 15:
		return "eastkarana"
	case 16:
		return "beholder"
	case 17:
		return "blackburrow"
	case 18:
		return "paw"
	case 19:
		return "rivervale"
	case 20:
		return "kithicor"
	case 21:
		return "commons"
	case 22:
		return "ecommons"
	case 23:
		return "erudnint"
	case 24:
		return "erudnext"
	case 25:
		return "nektulos"
	case 26:
		return "cshome"
	case 27:
		return "lavastorm"
	case 28:
		return "nektropos"
	case 29:
		return "halas"
	case 30:
		return "everfrost"
	case 31:
		return "soldunga"
	case 32:
		return "soldungb"
	case 33:
		return "misty"
	case 34:
		return "nro"
	case 35:
		return "sro"
	case 36:
		return "befallen"
	case 37:
		return "oasis"
	case 38:
		return "tox"
	case 39:
		return "hole"
	case 40:
		return "neriaka"
	case 41:
		return "neriakb"
	case 42:
		return "neriakc"
	case 43:
		return "neriakd"
	case 44:
		return "najena"
	case 45:
		return "qcat"
	case 46:
		return "innothule"
	case 47:
		return "feerrott"
	case 48:
		return "cazicthule"
	case 49:
		return "oggok"
	case 50:
		return "rathemtn"
	case 51:
		return "lakerathe"
	case 52:
		return "grobb"
	case 53:
		return "aviak"
	case 54:
		return "gfaydark"
	case 55:
		return "akanon"
	case 56:
		return "steamfont"
	case 57:
		return "lfaydark"
	case 58:
		return "crushbone"
	case 59:
		return "mistmoore"
	case 60:
		return "kaladima"
	case 61:
		return "felwithea"
	case 62:
		return "felwitheb"
	case 63:
		return "unrest"
	case 64:
		return "kedge"
	case 65:
		return "guktop"
	case 66:
		return "gukbottom"
	case 67:
		return "kaladimb"
	case 68:
		return "butcher"
	case 69:
		return "oot"
	case 70:
		return "cauldron"
	case 71:
		return "airplane"
	case 72:
		return "fearplane"
	case 73:
		return "permafrost"
	case 74:
		return "kerraridge"
	case 75:
		return "paineel"
	case 76:
		return "hateplane"
	case 77:
		return "arena"
	case 78:
		return "fieldofbone"
	case 79:
		return "warslikswood"
	case 80:
		return "soltemple"
	case 81:
		return "droga"
	case 82:
		return "cabwest"
	case 83:
		return "swampofnohope"
	case 84:
		return "firiona"
	case 85:
		return "lakeofillomen"
	case 86:
		return "dreadlands"
	case 87:
		return "burningwood"
	case 88:
		return "kaesora"
	case 89:
		return "sebilis"
	case 90:
		return "citymist"
	case 91:
		return "skyfire"
	case 92:
		return "frontiermtns"
	case 93:
		return "overthere"
	case 94:
		return "emeraldjungle"
	case 95:
		return "trakanon"
	case 96:
		return "timorous"
	case 97:
		return "kurn"
	case 98:
		return "erudsxing"
	case 100:
		return "stonebrunt"
	case 101:
		return "warrens"
	case 102:
		return "karnor"
	case 103:
		return "chardok"
	case 104:
		return "dalnir"
	case 105:
		return "charasis"
	case 106:
		return "cabeast"
	case 107:
		return "nurga"
	case 108:
		return "veeshan"
	case 109:
		return "veksar"
	case 110:
		return "iceclad"
	case 111:
		return "frozenshadow"
	case 112:
		return "velketor"
	case 113:
		return "kael"
	case 114:
		return "skyshrine"
	case 115:
		return "thurgadina"
	case 116:
		return "eastwastes"
	case 117:
		return "cobaltscar"
	case 118:
		return "greatdivide"
	case 119:
		return "wakening"
	case 120:
		return "westwastes"
	case 121:
		return "crystal"
	case 123:
		return "necropolis"
	case 124:
		return "templeveeshan"
	case 125:
		return "sirens"
	case 126:
		return "mischiefplane"
	case 127:
		return "growthplane"
	case 128:
		return "sleeper"
	case 129:
		return "thurgadinb"
	case 130:
		return "erudsxing2"
	case 150:
		return "shadowhaven"
	case 151:
		return "bazaar"
	case 152:
		return "nexus"
	case 153:
		return "echo"
	case 154:
		return "acrylia"
	case 155:
		return "sharvahl"
	case 156:
		return "paludal"
	case 157:
		return "fungusgrove"
	case 158:
		return "vexthal"
	case 159:
		return "sseru"
	case 160:
		return "katta"
	case 161:
		return "netherbian"
	case 162:
		return "ssratemple"
	case 163:
		return "griegsend"
	case 164:
		return "thedeep"
	case 165:
		return "shadeweaver"
	case 166:
		return "hollowshade"
	case 167:
		return "grimling"
	case 168:
		return "mseru"
	case 169:
		return "letalis"
	case 170:
		return "twilight"
	case 171:
		return "thegrey"
	case 172:
		return "tenebrous"
	case 173:
		return "maiden"
	case 174:
		return "dawnshroud"
	case 175:
		return "scarlet"
	case 176:
		return "umbral"
	case 179:
		return "akheva"
	case 180:
		return "arena2"
	case 181:
		return "jaggedpine"
	case 182:
		return "nedaria"
	case 183:
		return "tutorial"
	case 184:
		return "load"
	case 185:
		return "load2"
	case 186:
		return "hateplaneb"
	case 187:
		return "shadowrest"
	case 188:
		return "tutoriala"
	case 189:
		return "tutorialb"
	case 190:
		return "clz"
	case 200:
		return "codecay"
	case 201:
		return "pojustice"
	case 202:
		return "poknowledge"
	case 203:
		return "potranquility"
	case 204:
		return "ponightmare"
	case 205:
		return "podisease"
	case 206:
		return "poinnovation"
	case 207:
		return "potorment"
	case 208:
		return "povalor"
	case 209:
		return "bothunder"
	case 210:
		return "postorms"
	case 211:
		return "hohonora"
	case 212:
		return "solrotower"
	case 213:
		return "powar"
	case 214:
		return "potactics"
	case 215:
		return "poair"
	case 216:
		return "powater"
	case 217:
		return "pofire"
	case 218:
		return "poeartha"
	case 219:
		return "potimea"
	case 220:
		return "hohonorb"
	case 221:
		return "nightmareb"
	case 222:
		return "poearthb"
	case 223:
		return "potimeb"
	case 224:
		return "gunthak"
	case 225:
		return "dulak"
	case 226:
		return "torgiran"
	case 227:
		return "nadox"
	case 228:
		return "hatesfury"
	case 229:
		return "guka"
	case 230:
		return "ruja"
	case 231:
		return "taka"
	case 232:
		return "mira"
	case 233:
		return "mmca"
	case 234:
		return "gukb"
	case 235:
		return "rujb"
	case 236:
		return "takb"
	case 237:
		return "mirb"
	case 238:
		return "mmcb"
	case 239:
		return "gukc"
	case 240:
		return "rujc"
	case 241:
		return "takc"
	case 242:
		return "mirc"
	case 243:
		return "mmcc"
	case 244:
		return "gukd"
	case 245:
		return "rujd"
	case 246:
		return "takd"
	case 247:
		return "mird"
	case 248:
		return "mmcd"
	case 249:
		return "guke"
	case 250:
		return "ruje"
	case 251:
		return "take"
	case 252:
		return "mire"
	case 253:
		return "mmce"
	case 254:
		return "gukf"
	case 255:
		return "rujf"
	case 256:
		return "takf"
	case 257:
		return "mirf"
	case 258:
		return "mmcf"
	case 259:
		return "gukg"
	case 260:
		return "rujg"
	case 261:
		return "takg"
	case 262:
		return "mirg"
	case 263:
		return "mmcg"
	case 264:
		return "gukh"
	case 265:
		return "rujh"
	case 266:
		return "takh"
	case 267:
		return "mirh"
	case 268:
		return "mmch"
	case 269:
		return "ruji"
	case 270:
		return "taki"
	case 271:
		return "miri"
	case 272:
		return "mmci"
	case 273:
		return "rujj"
	case 274:
		return "takj"
	case 275:
		return "mirj"
	case 276:
		return "mmcj"
	case 277:
		return "chardokb"
	case 278:
		return "soldungc"
	case 279:
		return "abysmal"
	case 280:
		return "natimbi"
	case 281:
		return "qinimi"
	case 282:
		return "riwwi"
	case 283:
		return "barindu"
	case 284:
		return "ferubi"
	case 285:
		return "snpool"
	case 286:
		return "snlair"
	case 287:
		return "snplant"
	case 288:
		return "sncrematory"
	case 289:
		return "tipt"
	case 290:
		return "vxed"
	case 291:
		return "yxtta"
	case 292:
		return "uqua"
	case 293:
		return "kodtaz"
	case 294:
		return "ikkinz"
	case 295:
		return "qvic"
	case 296:
		return "inktuta"
	case 297:
		return "txevu"
	case 298:
		return "tacvi"
	case 299:
		return "qvicb"
	case 300:
		return "wallofslaughter"
	case 301:
		return "bloodfields"
	case 302:
		return "draniksscar"
	case 303:
		return "causeway"
	case 304:
		return "chambersa"
	case 305:
		return "chambersb"
	case 306:
		return "chambersc"
	case 307:
		return "chambersd"
	case 308:
		return "chamberse"
	case 309:
		return "chambersf"
	case 316:
		return "provinggrounds"
	case 317:
		return "anguish"
	case 318:
		return "dranikhollowsa"
	case 319:
		return "dranikhollowsb"
	case 320:
		return "dranikhollowsc"
	case 328:
		return "dranikcatacombsa"
	case 329:
		return "dranikcatacombsb"
	case 330:
		return "dranikcatacombsc"
	case 331:
		return "draniksewersa"
	case 332:
		return "draniksewersb"
	case 333:
		return "draniksewersc"
	case 334:
		return "riftseekers"
	case 335:
		return "harbingers"
	case 336:
		return "dranik"
	case 337:
		return "broodlands"
	case 338:
		return "stillmoona"
	case 339:
		return "stillmoonb"
	case 340:
		return "thundercrest"
	case 341:
		return "delvea"
	case 342:
		return "delveb"
	case 343:
		return "thenest"
	case 344:
		return "guildlobby"
	case 345:
		return "guildhall"
	case 346:
		return "barter"
	case 347:
		return "illsalin"
	case 348:
		return "illsalina"
	case 349:
		return "illsalinb"
	case 350:
		return "illsalinc"
	case 351:
		return "dreadspire"
	case 354:
		return "drachnidhive"
	case 355:
		return "drachnidhivea"
	case 356:
		return "drachnidhiveb"
	case 357:
		return "drachnidhivec"
	case 358:
		return "westkorlach"
	case 359:
		return "westkorlacha"
	case 360:
		return "westkorlachb"
	case 361:
		return "westkorlachc"
	case 362:
		return "eastkorlach"
	case 363:
		return "eastkorlacha"
	case 364:
		return "shadowspine"
	case 365:
		return "corathus"
	case 366:
		return "corathusa"
	case 367:
		return "corathusb"
	case 368:
		return "nektulosa"
	case 369:
		return "arcstone"
	case 370:
		return "relic"
	case 371:
		return "skylance"
	case 372:
		return "devastation"
	case 373:
		return "devastationa"
	case 374:
		return "rage"
	case 375:
		return "ragea"
	case 376:
		return "takishruins"
	case 377:
		return "takishruinsa"
	case 378:
		return "elddar"
	case 379:
		return "elddara"
	case 380:
		return "theater"
	case 381:
		return "theatera"
	case 382:
		return "freeporteast"
	case 383:
		return "freeportwest"
	case 384:
		return "freeportsewers"
	case 385:
		return "freeportacademy"
	case 386:
		return "freeporttemple"
	case 387:
		return "freeportmilitia"
	case 388:
		return "freeportarena"
	case 389:
		return "freeportcityhall"
	case 390:
		return "freeporttheater"
	case 391:
		return "freeporthall"
	case 392:
		return "northro"
	case 393:
		return "southro"
	case 394:
		return "crescent"
	case 395:
		return "moors"
	case 396:
		return "stonehive"
	case 397:
		return "mesa"
	case 398:
		return "roost"
	case 399:
		return "steppes"
	case 400:
		return "icefall"
	case 401:
		return "valdeholm"
	case 402:
		return "frostcrypt"
	case 403:
		return "sunderock"
	case 404:
		return "vergalid"
	case 405:
		return "direwind"
	case 406:
		return "ashengate"
	case 407:
		return "highpasshold"
	case 408:
		return "commonlands"
	case 409:
		return "oceanoftears"
	case 410:
		return "kithforest"
	case 411:
		return "befallenb"
	case 412:
		return "highpasskeep"
	case 413:
		return "innothuleb"
	case 414:
		return "toxxulia"
	case 415:
		return "mistythicket"
	case 416:
		return "kattacastrum"
	case 417:
		return "thalassius"
	case 418:
		return "atiiki"
	case 419:
		return "zhisza"
	case 420:
		return "silyssar"
	case 421:
		return "solteris"
	case 422:
		return "barren"
	case 423:
		return "buriedsea"
	case 424:
		return "jardelshook"
	case 425:
		return "monkeyrock"
	case 426:
		return "suncrest"
	case 427:
		return "deadbone"
	case 428:
		return "blacksail"
	case 429:
		return "maidensgrave"
	case 430:
		return "redfeather"
	case 431:
		return "shipmvp"
	case 432:
		return "shipmvu"
	case 433:
		return "shippvu"
	case 434:
		return "shipuvu"
	case 435:
		return "shipmvm"
	case 436:
		return "mechanotus"
	case 437:
		return "mansion"
	case 438:
		return "steamfactory"
	case 439:
		return "shipworkshop"
	case 440:
		return "gyrospireb"
	case 441:
		return "gyrospirez"
	case 442:
		return "dragonscale"
	case 443:
		return "lopingplains"
	case 444:
		return "hillsofshade"
	case 445:
		return "bloodmoon"
	case 446:
		return "crystallos"
	case 447:
		return "guardian"
	case 448:
		return "steamfontmts"
	case 449:
		return "cryptofshade"
	case 451:
		return "dragonscaleb"
	case 452:
		return "oldfieldofbone"
	case 453:
		return "oldkaesoraa"
	case 454:
		return "oldkaesorab"
	case 455:
		return "oldkurn"
	case 456:
		return "oldkithicor"
	case 457:
		return "oldcommons"
	case 458:
		return "oldhighpass"
	case 459:
		return "thevoida"
	case 460:
		return "thevoidb"
	case 461:
		return "thevoidc"
	case 462:
		return "thevoidd"
	case 463:
		return "thevoide"
	case 464:
		return "thevoidf"
	case 465:
		return "thevoidg"
	case 466:
		return "oceangreenhills"
	case 467:
		return "oceangreenvillage"
	case 468:
		return "oldblackburrow"
	case 469:
		return "bertoxtemple"
	case 470:
		return "discord"
	case 471:
		return "discordtower"
	case 472:
		return "oldbloodfield"
	case 473:
		return "precipiceofwar"
	case 474:
		return "olddranik"
	case 475:
		return "toskirakk"
	case 476:
		return "korascian"
	case 477:
		return "rathechamber"
	case 480:
		return "brellsrest"
	case 481:
		return "fungalforest"
	case 482:
		return "underquarry"
	case 483:
		return "coolingchamber"
	case 484:
		return "shiningcity"
	case 485:
		return "arthicrex"
	case 486:
		return "foundation"
	case 487:
		return "lichencreep"
	case 488:
		return "pellucid"
	case 489:
		return "stonesnake"
	case 490:
		return "brellstemple"
	case 491:
		return "convorteum"
	case 492:
		return "brellsarena"
	case 493:
		return "weddingchapel"
	case 494:
		return "weddingchapeldark"
	case 495:
		return "dragoncrypt"
	case 700:
		return "feerrott2"
	case 701:
		return "thulehouse1"
	case 702:
		return "thulehouse2"
	case 703:
		return "housegarden"
	case 704:
		return "thulelibrary"
	case 705:
		return "well"
	case 706:
		return "fallen"
	case 707:
		return "morellcastle"
	case 708:
		return "somnium"
	case 709:
		return "alkabormare"
	case 710:
		return "miragulmare"
	case 711:
		return "thuledream"
	case 712:
		return "neighborhood"
	case 724:
		return "argath"
	case 725:
		return "arelis"
	case 726:
		return "sarithcity"
	case 727:
		return "rubak"
	case 728:
		return "beastdomain"
	case 729:
		return "resplendent"
	case 730:
		return "pillarsalra"
	case 731:
		return "windsong"
	case 732:
		return "cityofbronze"
	case 733:
		return "sepulcher"
	case 734:
		return "eastsepulcher"
	case 735:
		return "westsepulcher"
	case 752:
		return "shardslanding"
	case 753:
		return "xorbb"
	case 754:
		return "kaelshard"
	case 755:
		return "eastwastesshard"
	case 756:
		return "crystalshard"
	case 757:
		return "breedinggrounds"
	case 758:
		return "eviltree"
	case 759:
		return "grelleth"
	case 760:
		return "chapterhouse"
	case 996:
		return "arttest"
	case 998:
		return "fhalls"
	case 999:
		return "apprentice"
	}
	return "unknown"
}

//ClassIcon returns xa-icon name
func ClassIcon(class int64) string {
	switch class {
	case 1:
		return "xa-shield" //warrior
	case 2:
		return "xa-ankh" //cleric
	case 3:
		return "xa-fireball-sword" //paladin
	case 4:
		return "xa-arrow-cluster" //ranger
	case 5:
		return "xa-bat-sword" //shd
	case 6:
		return "xa-leaf" //druid
	case 7:
		return "xa-hand-emblem" //Monk
	case 8:
		return "xa-ocarina" //Bard
	case 9:
		return "xa-hood" //rogue
	case 10:
		return "xa-incense" //shaman
	case 11:
		return "xa-skull" //necro
	case 12:
		return "xa-fire" //wiz
	case 13:
		return "xa-burning-book" //magician
	case 14:
		return "xa-crystal-ball" //enchanter
	case 15:
		return "xa-pawprint" //beastlord
	case 16:
		return "xa-axe" //ber
	}
	return "xa-help"
}

//ClassName returns human readable version of a class
func ClassName(class int64) string {
	switch class {
	case 1:
		return "Warrior"
	case 2:
		return "Cleric"
	case 3:
		return "Paladin"
	case 4:
		return "Ranger"
	case 5:
		return "Shadowknight"
	case 6:
		return "Druid"
	case 7:
		return "Monk"
	case 8:
		return "Bard"
	case 9:
		return "Rogue"
	case 10:
		return "Shaman"
	case 11:
		return "Necromancer"
	case 12:
		return "Wizard"
	case 13:
		return "Magician"
	case 14:
		return "Enchanter"
	case 15:
		return "Beastlord"
	case 16:
		return "Berserker"
	case 20:
		return "GM Warrior"
	case 21:
		return "GM Cleric"
	case 22:
		return "GM Paladin"
	case 23:
		return "GM Ranger"
	case 24:
		return "GM Shadow Knight"
	case 25:
		return "GM Druid"
	case 26:
		return "GM Monk"
	case 27:
		return "GM Bard"
	case 28:
		return "GM Rogue"
	case 29:
		return "GM Shaman"
	case 30:
		return "GM Necromancer"
	case 31:
		return "GM Wizard"
	case 32:
		return "GM Magician"
	case 33:
		return "GM Enchanter"
	case 34:
		return "GM Beastlord"
	case 35:
		return "GM Berserker"
	case 40:
		return "Banker"
	case 41:
		return "Shopkeeper"
	case 59:
		return "Discord Merchant"
	case 60:
		return "Adventure Recruiter"
	case 61:
		return "Adventure Merchant"
	case 63:
		return "Tribute Master"
	case 64:
		return "Guild Tribute Master?"
	case 66:
		return "Guild Bank"
	case 67:
		return "Radiant Crystal Merchant"
	case 68:
		return "Ebon Crystal Merchant"
	case 69:
		return "Fellowships"
	case 70:
		return "Alternate Currency Merchant"
	case 71:
		return "Mercenary Merchant"
	}
	return fmt.Sprintf("Unknown (%d)", class)
}

//RaceIcon returns the race icon
func RaceIcon(race int64) string {
	val, ok := raceIcons[race]
	if ok {
		return string(val)
	}
	return "xa-help"
}

//RaceName returns the race name
func RaceName(race int64) string {
	val, ok := raceNames[race]
	if ok {
		return string(val)
	}
	return fmt.Sprintf("Unknown (%d)", race)
}

var slots = map[int64]string{
	0:  "Charm",
	1:  "Ear",
	2:  "Head",
	3:  "Face",
	4:  "Ear",
	5:  "Neck",
	6:  "Shoulders",
	7:  "Arms",
	8:  "Back",
	9:  "Wrist",
	10: "Wrist",
	11: "Range",
	12: "Hands",
	13: "Primary",
	14: "Secondary",
	15: "Fingers",
	16: "Fingers",
	17: "Chest",
	18: "Legs",
	19: "Feet",
	20: "Waist",
	21: "Ammo",
}

var slotFlags = map[int64]int64{
	0:  1 << 0,        // Charm
	1:  1<<1 | 1<<4,   // Ear1 + Ear2
	2:  1 << 2,        // Head
	3:  1 << 3,        // Face
	4:  1 << 5,        // Neck
	5:  1 << 6,        // Shoulders
	6:  1 << 7,        // Arms
	7:  1 << 8,        // Back
	8:  1<<9 | 1<<10,  // Wrist1 + Wrist2
	9:  1 << 11,       // Range
	10: 1 << 12,       // Hands
	11: 1 << 13,       // Primary
	12: 1 << 14,       // Secondary
	13: 1<<15 | 1<<16, // Fingers1 + Fingers2
	14: 1 << 17,       // Chest
	15: 1 << 18,       // Legs
	16: 1 << 19,       // Feet
	17: 1 << 20,       // Waist
	18: 1 << 21,       // Ammo
}

var languages = map[int64]string{
	0:  "Common",
	1:  "Barbarian",
	2:  "Erudian",
	3:  "Elvish",
	4:  "Dark Elvish",
	5:  "Dwarvish",
	6:  "Troll",
	7:  "Ogre",
	8:  "Gnomish",
	9:  "Halfling",
	10: "Thieves Cant",
	11: "Old Erudian",
	12: "Elder Elvish",
	13: "Froglok",
	14: "Goblin",
	15: "Gnoll",
	16: "Combine",
	17: "Elder Teir`Dal",
	18: "Lizardman",
	19: "Orcish",
	20: "Faerie",
	21: "Dragon",
	22: "Elder Dragon",
	23: "Dark Speech",
	24: "Vah Shir",
	25: "Alaran",
	26: "Hadal",
}

var bagTypes = map[int64]string{
	0:  "None",
	1:  "General",
	2:  "Quiver",
	3:  "Belt Pouch",
	4:  "Pouch",
	5:  "Backpack",
	6:  "Small Box",
	7:  "Medium Box",
	8:  "Bandolier",
	9:  "Medicine Bag",
	10: "Toolbox",
	11: "Research Tome",
	12: "Mortar and Pestle",
	13: "Quest Container",
	14: "Mixing Bowl",
	15: "Spit",
	16: "Sewing Kit",
	17: "Unknown (17)",
	18: "Fletching Kit",
	19: "Distillery",
	20: "Jeweler's Kit",
	21: "Pottery Wheel",
	22: "Kiln",
	23: "Unknown (23)",
	24: "Lexicon",
	25: "Grimoire",
	26: "Binding",
	27: "Tome",
	28: "Unknown (28)",
	29: "Research Tome 2",
	30: "Quest Container 2",
	31: "Tackle Box",
	32: "Trader's Satchel",
	33: "Augmentation Sealer",
	34: "Ice Cream Churn",
	35: "Ornamentation",
	36: "Ornamentation Stripper",
	37: "Unattuner",
	38: "Tradeskill",
}

var augTypes = map[int64]string{
	0:  "0 (None)",
	1:  "1 (General: Single Stat)",
	2:  "2 (General: Multiple Stat)",
	3:  "3 (General: Spell Effect)",
	4:  "4 (Weapon: General)",
	5:  "5 (Weapon: Elem Damage)",
	6:  "6 (Weapon: Base Damage)",
	7:  "7 (General: Group)",
	8:  "8 (General: Raid)",
	9:  "9 (General: Dragons Points)",
	10: "10 (Crafted: Common)",
	11: "11 (Crafted: Group)",
	12: "12 (Crafted: Raid)",
	13: "13 (Energeiac: Group)",
	14: "14 (Energeiac: Raid)",
	15: "15 (Emblem)",
	16: "16 (Crafted: Group Symbol)",
	17: "17 (Crafted: Raid Foci)",
	18: "18 (Unknown)",
	19: "19 (Unknown)",
	20: "20 (Ornamentation)",
	21: "21 (Special Ornamentation)",
	22: "22 (Unknown)",
	23: "23 (Unknown)",
	24: "24 (Unknown)",
	25: "25 (Unknown)",
	26: "26 (Unknown)",
	27: "27 (Unknown)",
	28: "28 (Unknown)",
	29: "29 (Unknown)",
	30: "30 (Invisible: Epic Upgrade)",
}

var augRestrictions = map[int64]string{
	0:  "None",
	1:  "Armor Only",
	2:  "Weapons Only",
	3:  "One-Handed Weapons Only",
	4:  "Two-Handed Weapons Only",
	5:  "One-Handed Slashing Weapons Only",
	6:  "One-Handed Blunt Weapons Only",
	7:  "Piercing Only",
	8:  "Hand-to-Hand Weapons Only",
	9:  "Two-Handed Slashing Weapons Only",
	10: "Two-Handed Blunt Weapons Only",
	11: "Two-Handed Piercing Weapons Only",
	12: "Ranged Weapons Only",
	13: "Shields Only",
	14: "1H Slashing, 1H Blunt, or H2H Weapons Only",
	15: "1H Blunt or Hand-to-Hand Weapons Only",
	16: "Unknown (16)",
	17: "1H Blunt or 2H Blunt Weapons Only",
}

var lightTypes = map[int64]string{
	0:  "None",
	1:  "1 Lux - Candle ",
	2:  "2 Lux - Torch ",
	3:  "3 Lux - Blue Glow ",
	4:  "6 Lux - Small Lantern ",
	5:  "7 Lux - Blue Glow ",
	6:  "8 Lux - Large Lantern ",
	7:  "9 Lux - Flameless Lantern ",
	8:  "9 Lux - Globe of Stars ",
	9:  "3 Lux - Light Globe ",
	10: "7 Lux - Lightstone ",
	11: "9 Lux - Greater Lightstone ",
	12: "4 Lux - Red Glow ",
	13: "5 Lux - Blue Glow ",
	14: "4 Lux - Red Glow ",
	15: "5 Lux - Blue Glow ",
}

var materials = map[int64]string{
	0:  "Cloth / None",
	1:  "Leather",
	2:  "Chain",
	3:  "Plate",
	4:  "Monk",
	5:  "Unknown (5)",
	6:  "Unknown (6)",
	7:  "Kunark Chain",
	8:  "Unknown (8)",
	9:  "Unknown (9)",
	10: "Crimson Robe",
	11: "Flowing Black Robe",
	12: "Cryosilk Robe",
	13: "Robe of the Oracle",
	14: "Robe of the Kedge",
	15: "Shining Metallic Robes",
	16: "Plain Robe",
	17: "Velious Leather 1",
	18: "Velious Chain 1",
	19: "Velious Plate 1",
	20: "Velious Leather 2",
	21: "Velious Chain 2",
	22: "Velious Plate 2",
	23: "Velious Monk",
}

var bardTypes = map[int64]int64{
	0: 23,
	1: 24,
	2: 25,
	3: 26,
	4: 50,
	5: 51,
}

var itemTypes = map[int64]string{
	0:  "1H Slashing",
	1:  "2H Slashing",
	2:  "1H Piercing",
	3:  "1H Blunt",
	4:  "2H Blunt",
	5:  "Archery - Bow", // Bow
	6:  "Unknown (6)",
	7:  "Throwing - Large", // Large
	8:  "Shield",
	9:  "Scroll",
	10: "Armor",
	11: "Inventory",
	12: "Lockpicks",
	13: "Unknown (13)",
	14: "Food",
	15: "Drink",
	16: "Light",
	17: "Combinable",
	18: "Bandage",
	19: "Throwing - Small", // Small
	20: "Spell",            // spells and tomes
	21: "Potion",
	22: "Unknown (22)",
	23: "Wind",
	24: "Stringed",
	25: "Brass",
	26: "Percussion",
	27: "Archery - Arrow", // Arrow
	28: "Unknown (28)",
	29: "Jewelry",
	30: "Skull",
	31: "Book", // skill-up tomes/books? (would probably need a pp flag if true...)
	32: "Note",
	33: "Key",
	34: "Coin",
	35: "2H Piercing",
	36: "Fishing Pole",
	37: "Fishing Bait",
	38: "Alcohol",
	39: "Key", // keys and satchels?? (questable keys?)
	40: "Compass",
	41: "Unknown (41)",
	42: "Poison", // might be wrong, but includes poisons
	43: "Unknown (43)",
	44: "Unknown (44)",
	45: "Hand/Hand",
	46: "Unknown (46)",
	47: "Unknown (47)",
	48: "Unknown (48)",
	49: "Unknown (49)",
	50: "Singing",
	51: "All Instruments",
	52: "Charm",
	53: "Dye",
	54: "Augmentation",
	55: "Augmentation Solvent",
	56: "Augmentation Distiller",
	57: "Unknown (57)",
	58: "Banner/Fellowship Kit",
	59: "Unknown (59)",
	60: "Recipe Book",
	61: "Advanced Recipe",
	62: "Journal",      // only one(1) database entry
	63: "Alt Currency", // alt-currency (as opposed to coinage)
	64: "Perfected Augmentation Distiller",
	65: "Unknown (65)",
	66: "Unknown (66)",
	67: "Unknown (67)",
	68: "Mount",
	69: "Unknown (69)",
}

var itemTypeClasses = map[int64]int64{
	0:  WAR + PAL + RNG + SHD + DRU + BRD + ROG,                   // 1HS
	1:  WAR + PAL + RNG + SHD + BER,                               // 2HS
	2:  WAR + RNG + SHD + BRD + ROG + NEC + WIZ + MAG + ENC + BST, // Pierce
	3:  ALL - BER,                                                 // 1HB
	4:  ALL - BRD - ROG,                                           // 2HB
	5:  WAR + PAL + RNG + SHD + ROG,                               // Bow
	6:  0,                                                         // Unknown (6)
	7:  ALL - CLR - PAL - DRU - SHM,                               // Throwing
	8:  WAR + CLR + PAL + RNG + SHD + DRU + BRD + ROG + SHM,       // Shield
	9:  ALL,                                                       // Scroll
	10: ALL,                                                       // Armor
	11: ALL,                                                       // Inventory
	12: BRD + ROG,                                                 // Lockpicks
	13: 0,                                                         // Unknown (13)
	14: ALL,                                                       // Food
	15: ALL,                                                       // Drink
	16: ALL,                                                       // Light
	17: ALL,                                                       // Combinable
	18: ALL,                                                       // Bandage
	19: ALL - CLR - PAL - DRU - SHM,                               // Throwing - Small // Small
	20: ALL,                                                       // Spell			// spells and tomes
	21: ALL,                                                       // Potion
	22: ALL,                                                       // Unknown (22)
	23: BRD,                                                       // Wind
	24: BRD,                                                       // Stringed
	25: BRD,                                                       // Brass
	26: BRD,                                                       // Percussion
	27: WAR + PAL + RNG + SHD + ROG,                               // Archery - Arrow // Arrow
	28: ALL,                                                       // Unknown (28)
	29: ALL,                                                       // Jewelry
	30: ALL,                                                       // Skull
	31: ALL,                                                       // Book			// skill-up tomes/books? (would probably need a pp flag if true...)
	32: ALL,                                                       // Note
	33: ALL,                                                       // Key
	34: ALL,                                                       // Coin
	35: WAR + RNG + SHD + BRD + ROG + SHM + BST,                   // 2H Piercing
	36: ALL,                                                       // Fishing Pole
	37: ALL,                                                       // Fishing Bait
	38: ALL,                                                       // Alcohol
	39: ALL,                                                       // Key			// keys and satchels?? (questable keys?)
	40: ALL,                                                       // Compass
	41: ALL,                                                       // Unknown (41)
	42: ROG,                                                       // Poison			// might be wrong, but includes poisons
	43: ALL,                                                       // Unknown (43)
	44: ALL,                                                       // Unknown (44)
	45: MNK + BST,                                                 // Hand/Hand
	46: ALL,                                                       // Unknown (46)
	47: ALL,                                                       // Unknown (47)
	48: ALL,                                                       // Unknown (48)
	49: ALL,                                                       // Unknown (49)
	50: BRD,                                                       // Singing
	51: BRD,                                                       // All Instruments
	52: ALL,                                                       // Charm
	53: ALL,                                                       // Dye
	54: ALL,                                                       // Augmentation
	55: ALL,                                                       // Augmentation Solvent
	56: ALL,                                                       // Augmentation Distiller
	57: ALL,                                                       // Unknown (57)
	58: ALL,                                                       // Banner/Fellowship Kit
	59: ALL,                                                       // Unknown (59)
	60: ALL,                                                       // Recipe Book
	61: ALL,                                                       // Advanced Recipe
	62: ALL,                                                       // Journal (only one database entry)
	63: ALL,                                                       // Alt Currency (as opposed to coinage)
	64: ALL,                                                       // Perfected Augmentation Distiller
	65: ALL,                                                       // Unknown (65)
	66: ALL,                                                       // Unknown (66)
	67: ALL,                                                       // Unknown (67)
	68: ALL,                                                       // Mount
	69: ALL,                                                       // Unknown (69)
}

var iconModels = map[int64]int64{
	// Blunts
	567:  24,
	578:  37,
	581:  33,
	601:  10938,
	602:  126,
	686:  75,
	737:  18,
	738:  31,
	741:  19,
	809:  46,
	810:  47,
	811:  45,
	821:  35,
	822:  51,
	887:  29,
	889:  32,
	891:  52,
	903:  49,
	973:  72,
	1083: 75,
	1117: 178,
	1172: 96,
	1175: 111,
	1187: 103,
	1188: 102,
	1189: 101,
	1194: 112,
	1216: 130,
	1249: 134,
	1262: 185,
	1263: 177,
	1265: 169,
	1268: 10932,
	1274: 10406,
	1275: 10407,
	1277: 10410,
	1279: 10501,
	1280: 10502,
	1281: 10503,
	1282: 10663, // 10504 = blue sparkles. 10663 = purpl.
	1283: 10505,
	1288: 10517,
	1289: 10518,
	1290: 10519,
	1291: 10520,
	1292: 10521,
	1293: 10522,
	1294: 10523,
	1295: 10524,
	1296: 10525,
	1320: 10604,
	1321: 10605,
	1322: 10606,
	1323: 10607,
	1324: 10608,
	1325: 10609,
	1347: 10634,
	1348: 10635,
	1349: 10636,
	1350: 10637,
	1351: 10638,
	1352: 10639,
	1354: 10642,
	1355: 10643,
	1359: 10647,
	1376: 10506,
	1377: 10507,
	1378: 10508,
	1391: 10672,
	1393: 10675,
	1395: 10677, // Another standing torc!
	1399: 10681,
	1401: 10683,
	1405: 10688,
	1412: 10695,
	1421: 10706,
	1422: 10707,
	1428: 10713,
	1447: 10716,
	1449: 10718,
	1459: 10731,
	1460: 10732,
	1508: 10736, // 10752 = No Particls
	1511: 10739, // 10755 = No Particls
	1512: 10740, // 10756 = No Particls
	1513: 10741, // 10757 = No Particls
	1515: 10743, // 10759 = No Particls
	1520: 10748, // 10764 = No Particls
	1522: 10750, // 10766 = No Particls
	1523: 10767,
	1524: 10768,
	1530: 10774,
	1532: 10777,
	1540: 10783,
	1544: 10788,
	1550: 10794,
	1551: 10795,
	1552: 10796,
	1661: 10817,
	1662: 10818,
	1663: 10819,
	1664: 10820,
	1666: 10822,
	1667: 10823,
	1673: 10829,
	1674: 10830,
	1678: 10834,
	1679: 10835,
	1680: 10836,
	1681: 10837,
	1682: 10838,
	1683: 10839,
	1701: 10843,
	1703: 10845,
	1704: 10846,
	1705: 10847,
	1709: 10851,
	1710: 10852,
	1711: 10853,
	1712: 10854,
	1717: 10859,
	1782: 10924,
	1783: 10925,
	1784: 10926,
	1785: 10927,
	1786: 10928,
	1787: 10929,
	1788: 10930,
	1789: 10931,
	1790: 10932,
	1791: 10933,
	1792: 10934,
	1793: 10935,
	1794: 10936,
	1795: 10937,
	1796: 10938,
	1797: 10939,
	1798: 10940,
	1799: 10941,
	1800: 10942,
	1801: 10943,
	1802: 10944,
	1803: 10945,
	1804: 10946,
	1805: 10947,
	1806: 10948,
	1807: 10949,
	1858: 11031,
	1859: 11032,
	1860: 11033,
	1861: 11034,
	1862: 11035,
	1871: 11044,
	1872: 11045,
	1879: 11052,
	2154: 11087,
	2164: 11097,
	2165: 11098,
	2171: 11104,
	2172: 11105,
	2173: 11106,
	2174: 11107,
	2184: 11117,
	2186: 11119,
	2229: 11130,
	2232: 11133,

	// Books
	504:  28,
	682:  65,
	683:  28,
	777:  27,
	778:  27,
	789:  27,
	860:  28,
	861:  65,
	862:  28,
	863:  65,
	865:  27,
	1357: 10645,
	1358: 10646,
	1485: 400,
	1497: 27,
	2017: 11059,
	2031: 11073,

	// Bows & Arrows

	597:  10412,
	598:  10,
	681:  10725,
	725:  64,
	726:  64,
	1024: 10300,
	1039: 10,
	1104: 10300,
	1254: 199,
	1330: 10614,
	1379: 10641,
	1448: 10717,
	1453: 10725,
	1545: 10789,
	1628: 10641,
	1629: 10717,
	1630: 10789,
	1855: 10997,
	1856: 10998,

	// Containers

	557:  64,
	565:  64,
	608:  11054,
	609:  11054,
	689:  64,
	690:  64,
	691:  64,
	723:  11054,
	724:  11054,
	730:  11054,
	836:  11054,
	837:  11054,
	883:  11054,
	884:  64,
	892:  10802,
	979:  64,
	1016: 64,
	1017: 64,
	1112: 10800,
	1113: 10801,
	1114: 10803,
	1115: 10804,
	1116: 10805,
	1142: 64,
	1144: 64,
	1239: 64,
	1444: 64,
	1921: 64,
	1922: 64,
	1923: 64,
	1924: 64,
	1925: 64,
	1926: 64,
	1927: 64,
	1928: 64,
	1929: 64,
	1930: 64,
	1931: 64,
	1932: 64,
	1933: 64,
	1934: 64,
	1935: 64,
	1936: 64,
	1937: 64,
	1938: 64,
	2010: 11063,
	2011: 11062,
	2012: 11054,
	2020: 11062,
	2021: 11063,

	// Drink

	692:  10808,
	708:  56,
	709:  56,
	710:  10861,
	711:  10861,
	829:  10861,
	1557: 10806,
	1558: 10807,
	1559: 10808,
	1719: 10861,
	1939: 10861,

	// Food
	673:  11068,
	784:  11068,
	922:  11070,
	1000: 11068,
	1021: 11068,
	1086: 11068,
	1105: 11068,
	1108: 11068,
	1111: 11068,
	1688: 250,
	1691: 11071,
	1693: 11055,
	1696: 11071,
	2006: 11069,
	2008: 11069,
	2027: 11069,
	2028: 11070,
	2029: 11071,

	// Gear - Hands
	971:  68,
	1878: 11051,

	// Gear - Head
	511:  5003,
	523:  5003,
	536:  5808,
	550:  625,
	625:  5002,
	628:  5328,
	629:  5423,
	639:  5003,
	640:  5361,
	641:  5301,
	642:  5151,
	664:  5480,
	745:  5538,
	746:  5148,
	747:  5538,
	1608: 5838,
	1609: 5001,
	1610: 5361,
	1611: 5361,
	1612: 5032,
	1613: 5031,
	1614: 5032,
	1615: 5838,
	2113: 5031,
	2114: 5781,
	2115: 5151,
	2116: 5003,

	// Gear - Chest
	527:  6300,
	538:  6300,
	621:  6450,
	624:  6000,
	632:  64,
	678:  64,
	712:  4366,
	713:  4366,
	714:  64,
	838:  4363,
	928:  4365,
	929:  4365,
	930:  4360,
	931:  4360,
	940:  4364,
	941:  4362,
	942:  4362,
	1126: 4361,
	1149: 64,
	1568: 6660,
	1569: 6060,
	1570: 6210,
	1571: 6540,
	1572: 6180,
	1573: 6060,
	1574: 64,
	1575: 4363,
	2105: 4363,
	2106: 4366,
	2107: 4362,
	2108: 4362,
	2109: 6300,
	2110: 6420,
	2111: 6060,
	2112: 6300,

	// Gear - Waist
	843:  4562,
	1119: 4560,
	1120: 4372,
	1150: 4710,

	// Gear - Arms
	543:  8036,
	546:  8007,
	622:  8246,
	623:  8187,
	634:  7067,
	669:  8186,
	670:  8577,
	1592: 9666,
	1593: 9666,
	1594: 9666,
	1595: 9666,
	1596: 9666,
	1597: 9666,
	1598: 9666,
	1599: 9666,
	2117: 9547,
	2118: 9097,
	2119: 7067,
	2120: 9547,

	// Gear - Shoulders
	798: 7397,

	// Gear - Wrists
	516:  8097,
	520:  8097,
	521:  8247,
	620:  8607,
	637:  8757,
	638:  9367,
	671:  8037,
	1040: 8097,
	1044: 8157,
	1055: 8097,
	1057: 8037,
	1235: 8007,
	1584: 8067,
	1585: 8757,
	1586: 8727,
	1587: 8187,
	1588: 8157,
	1589: 8727,
	1590: 8667,
	1591: 8757,
	2121: 8187,
	2122: 9217,
	2123: 8067,
	2124: 8247,

	// Gemstones
	507:  10611,
	767:  11167,
	944:  11163,
	945:  11165,
	946:  11163,
	947:  11167,
	948:  11163,
	949:  11164,
	950:  11164,
	951:  11164,
	952:  11163,
	953:  11164,
	954:  11165,
	955:  11124,
	956:  11167,
	957:  11163,
	958:  11163,
	959:  11164,
	960:  11164,
	961:  11164,
	962:  11163,
	963:  11165,
	964:  11164,
	965:  11165,
	966:  11168,
	967:  11057,
	968:  11167,
	969:  11068,
	1253: 11168,
	1429: 11163,
	1430: 11165,
	1431: 10724,
	1432: 11165,
	1433: 11163,
	1434: 11162,
	1435: 11163,
	1436: 11164,
	1437: 11163,
	1438: 11068,
	1439: 10724,
	1440: 11163,
	1441: 11165,
	1442: 10724,
	1443: 11068,
	1452: 11068,
	1462: 11057,
	1463: 11057,
	1464: 11057,
	1465: 11057,
	1466: 11057,
	1467: 11057,
	1468: 11169,
	1469: 11057,
	1476: 11124,
	1477: 11165,
	1479: 10724,
	1480: 10724,
	1484: 11165,
	1486: 11165,
	1502: 11162,
	1503: 11163,
	1504: 11162,
	1505: 11123,
	1506: 11163,
	1535: 11165,
	1536: 11166,
	1940: 11164,
	1941: 11164,
	1942: 11164,
	1943: 11164,
	1944: 11164,
	1945: 11164,
	1946: 11124,
	1947: 11124,
	1948: 11124,
	1949: 11124,
	1950: 11124,
	1951: 11124,
	1952: 11166,
	1953: 11166,
	1954: 11166,
	1955: 11166,
	1956: 11166,
	1957: 11166,
	1958: 11163,
	1959: 11163,
	1960: 11163,
	1961: 11163,
	1962: 11163,
	1963: 11163,
	1964: 11165,
	1965: 11165,
	1966: 11165,
	1967: 11165,
	1968: 11165,
	1969: 11165,
	1970: 11165,
	1971: 11165,
	1972: 11165,
	1973: 11165,
	1974: 11165,
	1975: 11165,
	1976: 11162,
	1977: 11162,
	1978: 11162,
	1979: 11162,
	1980: 11162,
	1981: 11162,
	1982: 11168,
	1983: 11168,
	1984: 11168,
	1985: 11168,
	1986: 11168,
	1987: 11168,
	1988: 11167,
	1989: 11167,
	1990: 11167,
	1991: 11167,
	1992: 11167,
	1993: 11167,
	1994: 11164,
	1995: 11124,
	1996: 11166,
	1997: 11163,
	1998: 11165,
	1999: 11165,
	2000: 11162,
	2001: 11169,
	2002: 11168,
	2085: 11123,
	2086: 11124,
	2087: 11167,
	2088: 11165,
	2089: 11166,
	2090: 11162,
	2091: 11163,
	2092: 11165,
	2093: 11124,
	2094: 11166,
	2095: 11168,
	2096: 11167,
	2097: 11166,
	2098: 11167,
	2099: 11163,
	2100: 11163,
	2101: 11162,
	2102: 11168,
	2103: 11166,
	2104: 11169,
	2190: 11123,
	2191: 11124,
	2233: 11154,
	2234: 11155,
	2235: 11156,
	2236: 11157,
	2237: 11158,
	2238: 11159,
	2239: 11160,
	2240: 11161,
	2241: 11162,
	2242: 11163,
	2243: 11164,
	2244: 11165,
	2245: 11166,
	2246: 11167,
	2247: 11168,
	2248: 11169,

	// Instruments

	551:  10603,
	593:  15,
	594:  10600,
	751:  10601,
	1152: 10602,
	1317: 10601,
	1318: 11501,
	1319: 11502,

	// Parts
	552:  64,
	553:  64,
	554:  64,
	555:  64,
	556:  64,
	680:  64,
	743:  71,
	744:  64,
	774:  11118,
	775:  11118,
	786:  11118,
	787:  11118,
	794:  71,
	797:  11069,
	799:  11118,
	800:  15,
	801:  15,
	804:  11058,
	807:  15,
	814:  64,
	817:  64,
	818:  64,
	819:  11069,
	820:  64,
	823:  224,
	833:  64,
	834:  64,
	835:  64,
	853:  71,
	859:  11069,
	871:  10611,
	885:  402,
	905:  10695,
	906:  10695,
	907:  10695,
	917:  11070,
	918:  64,
	919:  64,
	920:  64,
	927:  220,
	943:  224,
	972:  10611,
	980:  64,
	981:  64,
	982:  64,
	983:  64,
	984:  64,
	985:  64,
	986:  64,
	987:  64,
	988:  64,
	989:  64,
	990:  64,
	991:  64,
	992:  64,
	993:  64,
	994:  64,
	995:  64,
	996:  64,
	997:  64,
	1001: 64,
	1003: 11069,
	1026: 10724,
	1067: 71,
	1069: 64,
	1070: 64,
	1088: 11069,
	1089: 71,
	1094: 11058,
	1125: 94,
	1130: 11069,
	1131: 11069,
	1136: 10611,
	1137: 10695,
	1139: 220,
	1203: 15,
	1204: 64,
	1205: 64,
	1207: 11118,
	1208: 64,
	1209: 64,
	1219: 15,
	1220: 217,
	1221: 64,
	1222: 64,
	1223: 220,
	1224: 220,
	1225: 224,
	1229: 15,
	1231: 64,
	1232: 11069,
	1233: 217,
	1234: 224,
	1236: 11069,
	1241: 15,
	1242: 220,
	1243: 217,
	1251: 64,
	1252: 64,
	1257: 64,
	1258: 64,
	1259: 64,
	1260: 64,
	1261: 64,
	1445: 64,
	1470: 64,
	1471: 64,
	1472: 64,
	1473: 64,
	1474: 64,
	1475: 64,
	1487: 64,
	1488: 64,
	1489: 64,
	1490: 64,
	1491: 64,
	1492: 64,
	1493: 64,
	1494: 64,
	1495: 64,
	1498: 10695,
	1499: 10724,
	1500: 10695,
	1501: 10695,
	1633: 64,
	1635: 71,
	1637: 10724,
	1723: 64,
	1904: 64,
	1905: 64,
	1906: 64,
	1907: 64,
	1908: 64,
	1909: 64,
	1910: 64,
	1911: 64,
	1912: 64,
	1913: 64,
	2016: 11058,
	2137: 11058,
	2138: 10695,
	2139: 10695,
	2150: 64,
	2151: 64,
	2185: 11118,

	// Pierce
	574:  10633,
	591:  10650,
	592:  10006,
	736:  23,
	740:  10100,
	742:  30,
	762:  10680,
	763:  10680,
	768:  10686,
	776:  10100,
	888:  10686,
	1163: 79,
	1179: 94,
	1182: 98,
	1183: 99,
	1214: 118,
	1266: 176,
	1267: 167,
	1270: 10028,
	1298: 10527,
	1299: 10528,
	1328: 10612,
	1329: 10613,
	1334: 10618,
	1337: 10621,
	1340: 10627,
	1343: 10630,
	1346: 10633,
	1356: 10644,
	1362: 10650,
	1365: 10650,
	1368: 10656,
	1371: 10659,
	1374: 10662,
	1382: 10624,
	1397: 10679,
	1398: 10680,
	1403: 10686,
	1404: 10687,
	1406: 10689,
	1407: 10690,
	1409: 10692,
	1410: 10693,
	1416: 10701,
	1418: 10703,
	1424: 10709,
	1426: 10711,
	1427: 10712,
	1451: 10722,
	1461: 10680,
	1507: 10735, // 10751 - Without particle effet
	1518: 10746, // 10762 - Without particle effet
	1525: 10769,
	1529: 10773,
	1539: 10782,
	1549: 10793,
	1554: 10798,
	1555: 10799,
	1668: 10824,
	1672: 10828,
	1675: 10831,
	1686: 10842,
	1702: 10844,
	1724: 10866,
	1726: 10868,
	1728: 10870,
	1740: 10882,
	1761: 10903,
	1768: 10910,
	1769: 10911,
	1770: 10912,
	1771: 10913,
	1772: 10914,
	1773: 10915,
	1774: 10916,
	1775: 10917,
	1776: 10918,
	1777: 10919,
	1808: 10950,
	1809: 10951,
	1810: 10952,
	1811: 10953,
	1812: 10954,
	1813: 10955,
	1814: 10956,
	1815: 10957,
	1816: 10958,
	1868: 11041,
	2155: 11088,
	2156: 11089,
	2158: 11091,
	2159: 11092,
	2160: 11093,
	2166: 11099,
	2168: 11101, // PIERCE STARS
	2179: 11112,
	2181: 11114,
	2183: 11116,
	2230: 11131,

	// Plants
	721:  11060,
	795:  11075,
	796:  403,
	854:  11081,
	911:  10000,
	1073: 11060,
	1133: 403,
	1196: 403,
	1197: 11060,
	1198: 11060,
	1199: 11079,
	1200: 11075,
	1201: 11056,
	1202: 403,
	1206: 11081,
	1636: 11060,
	2014: 11056,
	2018: 11060,
	2019: 11061,
	2024: 11066,
	2025: 11067,
	2032: 11075,
	2033: 11075,
	2034: 11076,
	2035: 11077,
	2036: 11078,
	2037: 11079,
	2038: 11080,
	2039: 11081,
	2040: 11082,
	2041: 11083,
	2042: 11084,
	2192: 11060,
	2193: 11060,
	2194: 11079,
	2195: 11079,
	2196: 11060,
	2197: 11060,
	2198: 11079,
	2199: 11079,
	2200: 11060,
	2201: 11060,
	2202: 11079,
	2203: 11079,
	2204: 11060,
	2205: 11060,
	2206: 11080,
	2207: 11079,
	2208: 11060,
	2209: 11060,
	2210: 11080,
	2211: 11079,
	2212: 11060,
	2213: 11060,
	2214: 11079,
	2215: 11079,
	2216: 11060,
	2217: 11060,
	2218: 11079,
	2219: 11079,
	2220: 11060,
	2221: 11060,
	2222: 11080,
	2223: 11080,
	2224: 11074,
	2225: 11074,
	2226: 11077,
	2227: 11077,

	// Rings
	505:  10512,
	508:  10512,
	509:  10510,
	512:  10512,
	513:  10511,
	515:  10512,
	607:  10511,
	612:  10509,
	613:  10512,
	614:  10511,
	615:  10512,
	616:  10512,
	617:  68,
	674:  10512,
	675:  10512,
	748:  68,
	765:  68,
	872:  10512,
	873:  10512,
	874:  10512,
	875:  10512,
	876:  10512,
	877:  10512,
	878:  68,
	879:  10512,
	880:  10511,
	1041: 10510,
	1045: 10510,
	1047: 10512,
	1051: 68,
	1052: 10509,
	1059: 10512,
	1060: 10512,
	1061: 10512,
	1064: 10512,
	1071: 10512,
	1072: 10510,
	1148: 10512,
	1255: 10511,
	1616: 10511,
	1617: 68,
	1618: 10512,
	1619: 10512,
	1620: 10512,
	1621: 68,
	1622: 10509,
	1623: 10511,
	1624: 10510,
	1645: 10512,
	1721: 68,
	1722: 10509,
	1886: 10509,
	1887: 10512,
	1888: 10512,
	1889: 10512,
	1890: 10512,
	1891: 10512,
	1892: 10512,
	1893: 10510,

	// Shields
	542:  203,
	606:  211,
	676:  11144,
	758:  204,
	759:  209,
	760:  208,
	805:  202,
	970:  214,
	974:  212,
	976:  213,
	1244: 218,
	1245: 217,
	1246: 220,
	1300: 10530,
	1301: 10531,
	1302: 10532,
	1303: 10533,
	1304: 10534,
	1305: 10535,
	1306: 10536,
	1307: 10537,
	1308: 10538,
	1309: 10539,
	1310: 10540,
	1311: 10541,
	1312: 10542,
	1313: 10543,
	1314: 10534,
	1327: 210,
	1375: 11002,
	1383: 10664,
	1384: 10665,
	1387: 10668,
	1388: 10669,
	1389: 10670,
	1390: 10671,
	1408: 10691,
	1457: 10729,
	1458: 10730,
	1496: 67367,
	1510: 10738,
	1528: 10772,
	1531: 10775,
	1538: 10781,
	1546: 10790,
	1670: 10826,
	1671: 10827,
	1676: 10832,
	1677: 10833,
	1707: 10849,
	1708: 10850,
	1715: 10857,
	1716: 10858,
	1818: 10960,
	1819: 10961,
	1820: 10962,
	1821: 10963,
	1822: 10964,
	1823: 10965,
	1824: 10966,
	1825: 10967,
	1826: 10968,
	1827: 10969,
	1828: 10970,
	1829: 10971,
	1830: 10972,
	1831: 10973,
	1832: 10974,
	1833: 10975,
	1834: 10976,
	1835: 10977,
	1836: 10978,
	1837: 10979,
	1838: 10980,
	1839: 10981,
	1840: 10982,
	1841: 10983,
	1842: 10984,
	1843: 10985,
	1844: 10986,
	1845: 10987,
	1846: 10988,
	1847: 10989,
	1848: 10990,
	1849: 10991,
	1850: 10992,
	1851: 10993,
	1852: 10994,
	1853: 10995,
	1854: 10996,
	1875: 11048,
	1876: 11049,
	2152: 11085,
	2153: 11086,
	2169: 11103,
	2170: 11102,
	2177: 11110,
	2178: 11111,

	// Slashing
	519:  161,
	568:  25,
	569:  10005,
	573:  60,
	575:  53,
	576:  11,
	577:  58,
	579:  39,
	580:  10026,
	588:  10023,
	589:  180,
	590:  168,
	596:  61,
	603:  10004,
	604:  41,
	605:  42,
	781:  40,
	847:  62, // Red blade?? Currently Fiery Avenger 2HS (oooool)
	882:  190,
	890:  175,
	902:  10010, // Use better if ca.
	975:  71,
	1164: 80,
	1165: 81,
	1166: 83,
	1167: 84,
	1168: 85,
	1169: 86,
	1170: 87,
	1171: 95,
	1173: 97,
	1174: 82,
	1176: 88,
	1177: 76,
	1178: 90,
	1180: 91,
	1181: 92,
	1184: 104,
	1185: 105,
	1186: 100,
	1190: 106,
	1191: 108,
	1192: 113,
	1193: 107,
	1195: 109,
	1211: 110,
	1212: 93,
	1215: 131,
	1247: 135,
	1264: 182,
	1269: 80,
	1271: 181,
	1272: 160, // Or 6.
	1273: 60,
	1284: 10513,
	1285: 10514,
	1286: 10515,
	1287: 10516,
	1297: 10526,
	1315: 10545, // 184 is greenish elvis.
	1326: 10610,
	1331: 10615,
	1332: 10616,
	1333: 10617,
	1335: 10619,
	1336: 10620,
	1338: 10625,
	1339: 10626,
	1341: 10628,
	1342: 10629,
	1344: 10631,
	1345: 10632,
	1353: 10640,
	1360: 10648,
	1361: 10649,
	1363: 10651,
	1364: 10652,
	1366: 10654,
	1367: 10655,
	1369: 10657,
	1370: 10658,
	1372: 10660,
	1373: 10661,
	1380: 10622,
	1381: 10623,
	1392: 10674,
	1394: 10676,
	1396: 10678,
	1400: 10682,
	1402: 10685,
	1411: 10694,
	1413: 10696,
	1415: 10700,
	1417: 10702,
	1419: 10704,
	1420: 10705,
	1423: 10708,
	1425: 10710,
	1446: 10715,
	1450: 10719,
	1454: 10726,
	1455: 10727,
	1456: 10728,
	1509: 10737, // 10753 w/o particls
	1516: 10744, // 10760 w/o particls
	1517: 10745, // 10761 w/o particls
	1519: 10747, // 10763 w/o particls
	1521: 10749, // 10765 w/o particls
	1526: 10770,
	1527: 10771,
	1533: 61, // Whip. May have been swapped to 10778 rapier w/ no ico.
	1534: 10776,
	1537: 10780,
	1541: 10784,
	1542: 10785,
	1543: 10786,
	1547: 10791,
	1548: 10792,
	1553: 10797,
	1631: 10904, // 1H Axe. Must have been scrappe.
	1632: 10905, // 2H Axe. Must have been scrappe.
	1654: 10810,
	1655: 10811,
	1656: 10812,
	1657: 10813,
	1658: 10814,
	1659: 10815,
	1660: 10816,
	1665: 10821,
	1669: 10825,
	1684: 10840,
	1685: 10841,
	1706: 10848,
	1713: 10855,
	1714: 10856,
	1718: 10860,
	1720: 10862,
	1725: 10867,
	1727: 10869,
	1729: 10871,
	1730: 10872,
	1731: 10873,
	1732: 10874,
	1733: 10875,
	1734: 10876,
	1735: 10877,
	1736: 10878,
	1737: 10879,
	1738: 10880,
	1739: 10881,
	1741: 10883,
	1742: 10884,
	1743: 10885,
	1744: 10886,
	1745: 10887,
	1746: 10888,
	1747: 10889,
	1748: 10890,
	1749: 10891,
	1750: 10892,
	1751: 10893,
	1752: 10894,
	1753: 10895,
	1754: 10896,
	1755: 10897,
	1756: 10898,
	1757: 10899,
	1758: 10900,
	1759: 10901,
	1760: 10902,
	1762: 10904,
	1763: 10905,
	1764: 10906,
	1765: 10907,
	1766: 10908,
	1767: 10909,
	1778: 10920,
	1779: 10921,
	1780: 10922,
	1781: 10923,
	1817: 10959,
	1869: 11042,
	1870: 11043,
	1873: 11046,
	1874: 11047,
	1880: 11053,
	2157: 11090,
	2161: 11094,
	2162: 11095,
	2167: 11100,
	2175: 11108,
	2176: 11109,
	2180: 11113,
	2182: 11115,
	2228: 11129,
	2231: 11132,

	// Tradeskills
	716:  65,
	732:  123,
	733:  11057,
	734:  10724,
	749:  38,
	790:  65,
	812:  65,
	857:  133,
	858:  133,
	978:  65,
	1008: 64,
	1012: 65,
	1013: 64,
	1031: 11072,
	1032: 11072,
	1033: 11072,
	1034: 11072,
	1035: 11072,
	1036: 64,
	1074: 11065,
	1075: 11064,
	1076: 11065,
	1082: 78,
	1090: 11068,
	1095: 10724,
	1096: 64,
	1103: 78,
	1110: 65,
	1135: 64,
	1138: 64,
	1140: 65,
	1143: 38,
	1151: 64,
	1213: 10944,
	1238: 64,
	1278: 122,
	1914: 11064,
	1915: 11064,
	1916: 11064,
	1917: 11065,
	2003: 11068,
	2013: 11065,
	2015: 11057,
	2022: 11064,
	2023: 11065,
	2026: 11068,
	2082: 64,
	2163: 11096,

	// Trinkets
	547:  36,
	559:  36,
	572:  61,
	644:  11122,
	645:  11121,
	646:  11122,
	647:  11121,
	648:  65,
	649:  65,
	650:  65,
	651:  65,
	652:  65,
	653:  65,
	654:  65,
	684:  48,
	685:  48,
	720:  10697,
	728:  48,
	729:  48,
	735:  199,
	893:  72,
	894:  72,
	895:  72,
	896:  65,
	897:  65,
	898:  65,
	899:  65,
	900:  65,
	901:  65,
	904:  55,
	1007: 67,
	1027: 133,
	1028: 199,
	1127: 133,
	1128: 36,
	1129: 11050,
	1248: 133,
	1276: 10409,
	1316: 10600,
	1385: 10666,
	1386: 10667,
	1414: 10697,
	1481: 11071,
	1514: 10742, // 10758 w/o particls
	1857: 10999,
	1863: 11036,
	1864: 11037,
	1865: 11038,
	1866: 11039,
	1867: 11040,
	1896: 11122,
	1897: 11122,
	1898: 11122,
	2030: 11072,
	2079: 11121,
	2084: 11122,
	2187: 11120,
	2188: 11121,
	2189: 11122,
	2249: 11139,
	2252: 11120,
	2253: 65,
	// 10750 - Lightning halo, rising stars
	// 1877's is 11050 (iceball w/ particles. Hand held item?
}

var itemClasses = map[int64]string{
	0:   "Common",
	1:   "Container",
	2:   "Readable",
	255: "Unknown",
}

var itemGroups = map[int64]string{
	0: "Unknown",
	1: "Weapon",
	2: "Armor",
	3: "Augmentation",
	4: "Container",
	5: "Instrument",
	6: "FoodDrink",
	7: "Benefit",
	8: "Potion",
}
