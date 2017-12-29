package model

import (
	"fmt"
)

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

func ZoneName(zoneId int64) string {
	switch zoneId {
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
	return "Unknown"
}

func RaceName(race int64) string {
	switch race {
	case 1:
		return "Human"
	case 2:
		return "Barbarian"
	case 3:
		return "Erudite"
	case 4:
		return "Wood Elf"
	case 5:
		return "High Elf"
	case 6:
		return "Dark Elf"
	case 7:
		return "Half Elf"
	case 8:
		return "Dwarf"
	case 9:
		return "Troll"
	case 10:
		return "Ogre"
	case 11:
		return "Halfling"
	case 12:
		return "Gnome"
	case 13:
		return "Aviak"
	case 14:
		return "Werewolf"
	case 15:
		return "Brownie"
	case 16:
		return "Centaur"
	case 17:
		return "Golem"
	case 18:
		return "Giant"
	case 19:
		return "Trakanon"
	case 20:
		return "Venril Sathir"
	case 21:
		return "Evil Eye"
	case 22:
		return "Beetle"
	case 23:
		return "Kerran"
	case 24:
		return "Fish"
	case 25:
		return "Fairy"
	case 26:
		return "Froglok"
	case 27:
		return "Froglok"
	case 28:
		return "Fungusman"
	case 29:
		return "Gargoyle"
	case 30:
		return "Gasbag"
	case 31:
		return "Gelatinous Cube"
	case 32:
		return "Ghost"
	case 33:
		return "Ghoul"
	case 34:
		return "Bat"
	case 35:
		return "Eel"
	case 36:
		return "Rat"
	case 37:
		return "Snake"
	case 38:
		return "Spider"
	case 39:
		return "Gnoll"
	case 40:
		return "Goblin"
	case 41:
		return "Gorilla"
	case 42:
		return "Wolf"
	case 43:
		return "Bear"
	case 44:
		return "Guard"
	case 45:
		return "Demi Lich"
	case 46:
		return "Imp"
	case 47:
		return "Griffin"
	case 48:
		return "Kobold"
	case 49:
		return "Dragon"
	case 50:
		return "Lion"
	case 51:
		return "Lizard Man"
	case 52:
		return "Mimic"
	case 53:
		return "Minotaur"
	case 54:
		return "Orc"
	case 55:
		return "Beggar"
	case 56:
		return "Pixie"
	case 57:
		return "Drachnid"
	case 58:
		return "Solusek Ro"
	case 59:
		return "Goblin"
	case 60:
		return "Skeleton"
	case 61:
		return "Shark"
	case 62:
		return "Tunare"
	case 63:
		return "Tiger"
	case 64:
		return "Treant"
	case 65:
		return "Vampire"
	case 66:
		return "Rallos Zek"
	case 67:
		return "Human"
	case 68:
		return "Tentacle Terror"
	case 69:
		return "Will-O-Wisp"
	case 70:
		return "Zombie"
	case 71:
		return "Human"
	case 72:
		return "Ship"
	case 73:
		return "Launch"
	case 74:
		return "Piranha"
	case 75:
		return "Elemental"
	case 76:
		return "Puma"
	case 77:
		return "Dark Elf"
	case 78:
		return "Erudite"
	case 79:
		return "Bixie"
	case 80:
		return "Reanimated Hand"
	case 81:
		return "Halfling"
	case 82:
		return "Scarecrow"
	case 83:
		return "Skunk"
	case 84:
		return "Snake Elemental"
	case 85:
		return "Spectre"
	case 86:
		return "Sphinx"
	case 87:
		return "Armadillo"
	case 88:
		return "Clockwork Gnome"
	case 89:
		return "Drake"
	case 90:
		return "Barbarian"
	case 91:
		return "Alligator"
	case 92:
		return "Troll"
	case 93:
		return "Ogre"
	case 94:
		return "Dwarf"
	case 95:
		return "Cazic Thule"
	case 96:
		return "Cockatrice"
	case 97:
		return "Daisy Man"
	case 98:
		return "Vampire"
	case 99:
		return "Amygdalan"
	case 100:
		return "Dervish"
	case 101:
		return "Efreeti"
	case 102:
		return "Tadpole"
	case 103:
		return "Kedge"
	case 104:
		return "Leech"
	case 105:
		return "Swordfish"
	case 106:
		return "Guard"
	case 107:
		return "Mammoth"
	case 108:
		return "Eye"
	case 109:
		return "Wasp"
	case 110:
		return "Mermaid"
	case 111:
		return "Harpy"
	case 112:
		return "Guard"
	case 113:
		return "Drixie"
	case 114:
		return "Ghost Ship"
	case 115:
		return "Clam"
	case 116:
		return "Seahorse"
	case 117:
		return "Ghost"
	case 118:
		return "Ghost"
	case 119:
		return "Sabertooth"
	case 120:
		return "Wolf"
	case 121:
		return "Gorgon"
	case 122:
		return "Dragon"
	case 123:
		return "Innoruuk"
	case 124:
		return "Unicorn"
	case 125:
		return "Pegasus"
	case 126:
		return "Djinn"
	case 127:
		return "Invisible Man"
	case 128:
		return "Iksar"
	case 129:
		return "Scorpion"
	case 130:
		return "Vah Shir"
	case 131:
		return "Sarnak"
	case 132:
		return "Draglock"
	case 133:
		return "Drolvarg"
	case 134:
		return "Mosquito"
	case 135:
		return "Rhinoceros"
	case 136:
		return "Xalgoz"
	case 137:
		return "Goblin"
	case 138:
		return "Yeti"
	case 139:
		return "Iksar"
	case 140:
		return "Giant"
	case 141:
		return "Boat"
	case 142:
		return "Object"
	case 143:
		return "Tree"
	case 144:
		return "Burynai"
	case 145:
		return "Goo"
	case 146:
		return "Sarnak Spirit"
	case 147:
		return "Iksar Spirit"
	case 148:
		return "Fish"
	case 149:
		return "Scorpion"
	case 150:
		return "Erollisi"
	case 151:
		return "Tribunal"
	case 152:
		return "Bertoxxulous"
	case 153:
		return "Bristlebane"
	case 154:
		return "Fay Drake"
	case 155:
		return "Undead Sarnak"
	case 156:
		return "Ratman"
	case 157:
		return "Wyvern"
	case 158:
		return "Wurm"
	case 159:
		return "Devourer"
	case 160:
		return "Iksar Golem"
	case 161:
		return "Undead Iksar"
	case 162:
		return "ManEating Plant"
	case 163:
		return "Raptor"
	case 164:
		return "Sarnak Golem"
	case 165:
		return "Dragon"
	case 166:
		return "Animated Hand"
	case 167:
		return "Succulent"
	case 168:
		return "Holgresh"
	case 169:
		return "Brontotherium"
	case 170:
		return "Snow Dervish"
	case 171:
		return "Dire Wolf"
	case 172:
		return "Manticore"
	case 173:
		return "Totem"
	case 174:
		return "Ice Spectre"
	case 175:
		return "Enchanted Armor"
	case 176:
		return "Snow Rabbit"
	case 177:
		return "Walrus"
	case 178:
		return "Geonid"
	case 181:
		return "Yakkar"
	case 182:
		return "Faun"
	case 183:
		return "Coldain"
	case 184:
		return "Dragon"
	case 185:
		return "Hag"
	case 186:
		return "Hippogriff"
	case 187:
		return "Siren"
	case 188:
		return "Giant"
	case 189:
		return "Giant"
	case 190:
		return "Othmir"
	case 191:
		return "Ulthork"
	case 192:
		return "Dragon"
	case 193:
		return "Abhorrent"
	case 194:
		return "Sea Turtle"
	case 195:
		return "Dragon"
	case 196:
		return "Dragon"
	case 197:
		return "Ronnie Test"
	case 198:
		return "Dragon"
	case 199:
		return "Shik'Nar"
	case 200:
		return "Rockhopper"
	case 201:
		return "Underbulk"
	case 202:
		return "Grimling"
	case 203:
		return "Worm"
	case 204:
		return "Evan Test"
	case 205:
		return "Shadel"
	case 206:
		return "Owlbear"
	case 207:
		return "Rhino Beetle"
	case 208:
		return "Vampire"
	case 209:
		return "Earth Elemental"
	case 210:
		return "Air Elemental"
	case 211:
		return "Water Elemental"
	case 212:
		return "Fire Elemental"
	case 213:
		return "Wetfang Minnow"
	case 214:
		return "Thought Horror"
	case 215:
		return "Tegi"
	case 216:
		return "Horse"
	case 217:
		return "Shissar"
	case 218:
		return "Fungal Fiend"
	case 219:
		return "Vampire"
	case 220:
		return "Stonegrabber"
	case 221:
		return "Scarlet Cheetah"
	case 222:
		return "Zelniak"
	case 223:
		return "Lightcrawler"
	case 224:
		return "Shade"
	case 225:
		return "Sunflower"
	case 226:
		return "Sun Revenant"
	case 227:
		return "Shrieker"
	case 228:
		return "Galorian"
	case 229:
		return "Netherbian"
	case 230:
		return "Akheva"
	case 231:
		return "Grieg Veneficus"
	case 232:
		return "Sonic Wolf"
	case 233:
		return "Ground Shaker"
	case 234:
		return "Vah Shir Skeleton"
	case 235:
		return "Wretch"
	case 236:
		return "Seru"
	case 237:
		return "Recuso"
	case 238:
		return "Vah Shir"
	case 239:
		return "Guard"
	case 240:
		return "Teleport Man"
	case 241:
		return "Werewolf"
	case 242:
		return "Nymph"
	case 243:
		return "Dryad"
	case 244:
		return "Treant"
	case 245:
		return "Fly"
	case 246:
		return "Tarew Marr"
	case 247:
		return "Solusek Ro"
	case 248:
		return "Clockwork Golem"
	case 249:
		return "Clockwork Brain"
	case 250:
		return "Banshee"
	case 251:
		return "Guard of Justice"
	case 252:
		return "Mini POM"
	case 253:
		return "Diseased Fiend"
	case 254:
		return "Solusek Ro Guard"
	case 255:
		return "Bertoxxulous"
	case 256:
		return "The Tribunal"
	case 257:
		return "Terris Thule"
	case 258:
		return "Vegerog"
	case 259:
		return "Crocodile"
	case 260:
		return "Bat"
	case 261:
		return "Hraquis"
	case 262:
		return "Tranquilion"
	case 263:
		return "Tin Soldier"
	case 264:
		return "Nightmare Wraith"
	case 265:
		return "Malarian"
	case 266:
		return "Knight of Pestilence"
	case 267:
		return "Lepertoloth"
	case 268:
		return "Bubonian"
	case 269:
		return "Bubonian Underling"
	case 270:
		return "Pusling"
	case 271:
		return "Water Mephit"
	case 272:
		return "Stormrider"
	case 273:
		return "Junk Beast"
	case 274:
		return "Broken Clockwork"
	case 275:
		return "Giant Clockwork"
	case 276:
		return "Clockwork Beetle"
	case 277:
		return "Nightmare Goblin"
	case 278:
		return "Karana"
	case 279:
		return "Blood Raven"
	case 280:
		return "Nightmare Gargoyle"
	case 281:
		return "Mouth of Insanity"
	case 282:
		return "Skeletal Horse"
	case 283:
		return "Saryrn"
	case 284:
		return "Fennin Ro"
	case 285:
		return "Tormentor"
	case 286:
		return "Soul Devourer"
	case 287:
		return "Nightmare"
	case 288:
		return "Rallos Zek"
	case 289:
		return "Vallon Zek"
	case 290:
		return "Tallon Zek"
	case 291:
		return "Air Mephit"
	case 292:
		return "Earth Mephit"
	case 293:
		return "Fire Mephit"
	case 294:
		return "Nightmare Mephit"
	case 295:
		return "Zebuxoruk"
	case 296:
		return "Mithaniel Marr"
	case 297:
		return "Undead Knight"
	case 298:
		return "The Rathe"
	case 299:
		return "Xegony"
	case 300:
		return "Fiend"
	case 301:
		return "Test Object"
	case 302:
		return "Crab"
	case 303:
		return "Phoenix"
	case 304:
		return "Dragon"
	case 305:
		return "Bear"
	case 306:
		return "Giant"
	case 307:
		return "Giant"
	case 308:
		return "Giant"
	case 309:
		return "Giant"
	case 310:
		return "Giant"
	case 311:
		return "Giant"
	case 312:
		return "Giant"
	case 313:
		return "War Wraith"
	case 314:
		return "Wrulon"
	case 315:
		return "Kraken"
	case 316:
		return "Poison Frog"
	case 317:
		return "Nilborien"
	case 318:
		return "Valorian"
	case 319:
		return "War Boar"
	case 320:
		return "Efreeti"
	case 321:
		return "War Boar"
	case 322:
		return "Valorian"
	case 323:
		return "Animated Armor"
	case 324:
		return "Undead Footman"
	case 325:
		return "Rallos Zek Minion"
	case 326:
		return "Arachnid"
	case 327:
		return "Crystal Spider"
	case 328:
		return "Zebuxoruk's Cage"
	case 329:
		return "Bastion of Thunder Portal"
	case 330:
		return "Froglok"
	case 331:
		return "Troll"
	case 332:
		return "Troll"
	case 333:
		return "Troll"
	case 334:
		return "Ghost"
	case 335:
		return "Pirate"
	case 336:
		return "Pirate"
	case 337:
		return "Pirate"
	case 338:
		return "Pirate"
	case 339:
		return "Pirate"
	case 340:
		return "Pirate"
	case 341:
		return "Pirate"
	case 342:
		return "Pirate"
	case 343:
		return "Frog"
	case 344:
		return "Troll Zombie"
	case 345:
		return "Luggald"
	case 346:
		return "Luggald"
	case 347:
		return "Luggalds"
	case 348:
		return "Drogmore"
	case 349:
		return "Froglok Skeleton"
	case 350:
		return "Undead Froglok"
	case 351:
		return "Knight of Hate"
	case 352:
		return "Arcanist of Hate"
	case 353:
		return "Veksar"
	case 354:
		return "Veksar"
	case 355:
		return "Veksar"
	case 356:
		return "Chokidai"
	case 357:
		return "Undead Chokidai"
	case 358:
		return "Undead Veksar"
	case 359:
		return "Vampire"
	case 360:
		return "Vampire"
	case 361:
		return "Rujarkian Orc"
	case 362:
		return "Bone Golem"
	case 363:
		return "Synarcana"
	case 364:
		return "Sand Elf"
	case 365:
		return "Vampire"
	case 366:
		return "Rujarkian Orc"
	case 367:
		return "Skeleton"
	case 368:
		return "Mummy"
	case 369:
		return "Goblin"
	case 370:
		return "Insect"
	case 371:
		return "Froglok Ghost"
	case 372:
		return "Dervish"
	case 373:
		return "Shade"
	case 374:
		return "Golem"
	case 375:
		return "Evil Eye"
	case 376:
		return "Box"
	case 377:
		return "Barrel"
	case 378:
		return "Chest"
	case 379:
		return "Vase"
	case 380:
		return "Table"
	case 381:
		return "Weapon Rack"
	case 382:
		return "Coffin"
	case 383:
		return "Bones"
	case 384:
		return "Jokester"
	case 385:
		return "Nihil"
	case 386:
		return "Trusik"
	case 387:
		return "Stone Worker"
	case 388:
		return "Hynid"
	case 389:
		return "Turepta"
	case 390:
		return "Cragbeast"
	case 391:
		return "Stonemite"
	case 392:
		return "Ukun"
	case 393:
		return "Ixt"
	case 394:
		return "Ikaav"
	case 395:
		return "Aneuk"
	case 396:
		return "Kyv"
	case 397:
		return "Noc"
	case 398:
		return "Ra`tuk"
	case 399:
		return "Taneth"
	case 400:
		return "Huvul"
	case 401:
		return "Mutna"
	case 402:
		return "Mastruq"
	case 403:
		return "Taelosian"
	case 404:
		return "Discord Ship"
	case 405:
		return "Stone Worker"
	case 406:
		return "Mata Muram"
	case 407:
		return "Lightning Warrior"
	case 408:
		return "Succubus"
	case 409:
		return "Bazu"
	case 410:
		return "Feran"
	case 411:
		return "Pyrilen"
	case 412:
		return "Chimera"
	case 413:
		return "Dragorn"
	case 414:
		return "Murkglider"
	case 415:
		return "Rat"
	case 416:
		return "Bat"
	case 417:
		return "Gelidran"
	case 418:
		return "Discordling"
	case 419:
		return "Girplan"
	case 420:
		return "Minotaur"
	case 421:
		return "Dragorn Box"
	case 422:
		return "Runed Orb"
	case 423:
		return "Dragon Bones"
	case 424:
		return "Muramite Armor Pile"
	case 425:
		return "Crystal Shard"
	case 426:
		return "Portal"
	case 427:
		return "Coin Purse"
	case 428:
		return "Rock Pile"
	case 429:
		return "Murkglider Egg Sack"
	case 430:
		return "Drake"
	case 431:
		return "Dervish"
	case 432:
		return "Drake"
	case 433:
		return "Goblin"
	case 434:
		return "Kirin"
	case 435:
		return "Dragon"
	case 436:
		return "Basilisk"
	case 437:
		return "Dragon"
	case 438:
		return "Dragon"
	case 439:
		return "Puma"
	case 440:
		return "Spider"
	case 441:
		return "Spider Queen"
	case 442:
		return "Animated Statue"
	case 445:
		return "Dragon Egg"
	case 446:
		return "Dragon Statue"
	case 447:
		return "Lava Rock"
	case 448:
		return "Animated Statue"
	case 449:
		return "Spider Egg Sack"
	case 450:
		return "Lava Spider"
	case 451:
		return "Lava Spider Queen"
	case 452:
		return "Dragon"
	case 453:
		return "Giant"
	case 454:
		return "Werewolf"
	case 455:
		return "Kobold"
	case 456:
		return "Sporali"
	case 457:
		return "Gnomework"
	case 458:
		return "Orc"
	case 459:
		return "Corathus"
	case 460:
		return "Coral"
	case 461:
		return "Drachnid"
	case 462:
		return "Drachnid Cocoon"
	case 463:
		return "Fungus Patch"
	case 464:
		return "Gargoyle"
	case 465:
		return "Witheran"
	case 466:
		return "Dark Lord"
	case 467:
		return "Shiliskin"
	case 468:
		return "Snake"
	case 469:
		return "Evil Eye"
	case 470:
		return "Minotaur"
	case 471:
		return "Zombie"
	case 472:
		return "Clockwork Boar"
	case 473:
		return "Fairy"
	case 474:
		return "Witheran"
	case 475:
		return "Air Elemental"
	case 476:
		return "Earth Elemental"
	case 477:
		return "Fire Elemental"
	case 478:
		return "Water Elemental"
	case 479:
		return "Alligator"
	case 480:
		return "Bear"
	case 481:
		return "Scaled Wolf"
	case 482:
		return "Wolf"
	case 483:
		return "Spirit Wolf"
	case 484:
		return "Skeleton"
	case 485:
		return "Spectre"
	case 486:
		return "Bolvirk"
	case 487:
		return "Banshee"
	case 488:
		return "Banshee"
	case 489:
		return "Elddar"
	case 490:
		return "Forest Giant"
	case 491:
		return "Bone Golem"
	case 492:
		return "Horse"
	case 493:
		return "Pegasus"
	case 494:
		return "Shambling Mound"
	case 495:
		return "Scrykin"
	case 496:
		return "Treant"
	case 497:
		return "Vampire"
	case 498:
		return "Ayonae Ro"
	case 499:
		return "Sullon Zek"
	case 500:
		return "Banner"
	case 501:
		return "Flag"
	case 502:
		return "Rowboat"
	case 503:
		return "Bear Trap"
	case 504:
		return "Clockwork Bomb"
	case 505:
		return "Dynamite Keg"
	case 506:
		return "Pressure Plate"
	case 507:
		return "Puffer Spore"
	case 508:
		return "Stone Ring"
	case 509:
		return "Root Tentacle"
	case 510:
		return "Runic Symbol"
	case 511:
		return "Saltpetter Bomb"
	case 512:
		return "Floating Skull"
	case 513:
		return "Spike Trap"
	case 514:
		return "Totem"
	case 515:
		return "Web"
	case 516:
		return "Wicker Basket"
	case 517:
		return "Nightmare/Unicorn"
	case 518:
		return "Horse"
	case 519:
		return "Nightmare/Unicorn"
	case 520:
		return "Bixie"
	case 521:
		return "Centaur"
	case 522:
		return "Drakkin"
	case 523:
		return "Giant"
	case 524:
		return "Gnoll"
	case 525:
		return "Griffin"
	case 526:
		return "Giant Shade"
	case 527:
		return "Harpy"
	case 528:
		return "Mammoth"
	case 529:
		return "Satyr"
	case 530:
		return "Dragon"
	case 531:
		return "Dragon"
	case 532:
		return "Dyn'Leth"
	case 533:
		return "Boat"
	case 534:
		return "Weapon Rack"
	case 535:
		return "Armor Rack"
	case 536:
		return "Honey Pot"
	case 537:
		return "Jum Jum Bucket"
	case 538:
		return "Toolbox"
	case 539:
		return "Stone Jug"
	case 540:
		return "Small Plant"
	case 541:
		return "Medium Plant"
	case 542:
		return "Tall Plant"
	case 543:
		return "Wine Cask"
	case 544:
		return "Elven Boat"
	case 545:
		return "Gnomish Boat"
	case 546:
		return "Barrel Barge Ship"
	case 547:
		return "Goo"
	case 548:
		return "Goo"
	case 549:
		return "Goo"
	case 550:
		return "Merchant Ship"
	case 551:
		return "Pirate Ship"
	case 552:
		return "Ghost Ship"
	case 553:
		return "Banner"
	case 554:
		return "Banner"
	case 555:
		return "Banner"
	case 556:
		return "Banner"
	case 557:
		return "Banner"
	case 558:
		return "Aviak"
	case 559:
		return "Beetle"
	case 560:
		return "Gorilla"
	case 561:
		return "Kedge"
	case 562:
		return "Kerran"
	case 563:
		return "Shissar"
	case 564:
		return "Siren"
	case 565:
		return "Sphinx"
	case 566:
		return "Human"
	case 567:
		return "Campfire"
	case 568:
		return "Brownie"
	case 569:
		return "Dragon"
	case 570:
		return "Exoskeleton"
	case 571:
		return "Ghoul"
	case 572:
		return "Clockwork Guardian"
	case 573:
		return "Mantrap"
	case 574:
		return "Minotaur"
	case 575:
		return "Scarecrow"
	case 576:
		return "Shade"
	case 577:
		return "Rotocopter"
	case 578:
		return "Tentacle Terror"
	case 579:
		return "Wereorc"
	case 580:
		return "Worg"
	case 581:
		return "Wyvern"
	case 582:
		return "Chimera"
	case 583:
		return "Kirin"
	case 584:
		return "Puma"
	case 585:
		return "Boulder"
	case 586:
		return "Banner"
	case 587:
		return "Elven Ghost"
	case 588:
		return "Human Ghost"
	case 589:
		return "Chest"
	case 590:
		return "Chest"
	case 591:
		return "Crystal"
	case 592:
		return "Coffin"
	case 593:
		return "Guardian CPU"
	case 594:
		return "Worg"
	case 595:
		return "Mansion"
	case 596:
		return "Floating Island"
	case 597:
		return "Cragslither"
	case 598:
		return "Wrulon"
	case 599:
		return "Spell Particle 1"
	case 600:
		return "Invisible Man of Zomm"
	case 601:
		return "Robocopter of Zomm"
	case 602:
		return "Burynai"
	case 603:
		return "Frog"
	case 604:
		return "Dracolich"
	case 605:
		return "Iksar Ghost"
	case 606:
		return "Iksar Skeleton"
	case 607:
		return "Mephit"
	case 608:
		return "Muddite"
	case 609:
		return "Raptor"
	case 610:
		return "Sarnak"
	case 611:
		return "Scorpion"
	case 612:
		return "Tsetsian"
	case 613:
		return "Wurm"
	case 614:
		return "Nekhon"
	case 615:
		return "Hydra Crystal"
	case 616:
		return "Crystal Sphere"
	case 617:
		return "Gnoll"
	case 618:
		return "Sokokar"
	case 619:
		return "Stone Pylon"
	case 620:
		return "Demon Vulture"
	case 621:
		return "Wagon"
	case 622:
		return "God of Discord"
	case 623:
		return "Feran Mount"
	case 624:
		return "Ogre NPC Male"
	case 625:
		return "Sokokar Mount"
	case 626:
		return "Giant"
	case 627:
		return "Sokokar"
	case 628:
		return "10th Anniversary Banner"
	case 629:
		return "10th Anniversary Cake"
	case 630:
		return "Wine Cask"
	case 631:
		return "Hydra Mount"
	case 632:
		return "Hydra NPC"
	case 633:
		return "Wedding Flowers"
	case 634:
		return "Wedding Arbor"
	case 635:
		return "Wedding Altar"
	case 636:
		return "Powder Keg"
	case 637:
		return "Apexus"
	case 638:
		return "Bellikos"
	case 639:
		return "Brell's First Creation"
	case 640:
		return "Brell"
	case 641:
		return "Crystalskin Ambuloid"
	case 642:
		return "Cliknar Queen"
	case 643:
		return "Cliknar Soldier"
	case 644:
		return "Cliknar Worker"
	case 645:
		return "Coldain"
	case 646:
		return "Coldain"
	case 647:
		return "Crystalskin Sessiloid"
	case 648:
		return "Genari"
	case 649:
		return "Gigyn"
	case 650:
		return "Greken Young Adult"
	case 651:
		return "Greken Young"
	case 652:
		return "Cliknar Mount"
	case 653:
		return "Telmira"
	case 654:
		return "Spider Mount"
	case 655:
		return "Bear Mount"
	case 656:
		return "Rat Mount Mystery Race"
	case 657:
		return "Sessiloid Mount"
	case 658:
		return "Morell Thule"
	case 659:
		return "Marionette"
	case 660:
		return "Book Dervish"
	case 661:
		return "Topiary Lion"
	case 662:
		return "Rotdog"
	case 663:
		return "Amygdalan"
	case 664:
		return "Sandman"
	case 665:
		return "Grandfather Clock"
	case 666:
		return "Gingerbread Man"
	case 667:
		return "Royal Guard"
	case 668:
		return "Rabbit"
	case 669:
		return "Blind Dreamer"
	case 670:
		return "Cazic Thule"
	case 671:
		return "Topiary Lion Mount"
	case 672:
		return "Rot Dog Mount"
	case 673:
		return "Goral Mount"
	case 674:
		return "Selyrah Mount"
	case 675:
		return "Sclera Mount"
	case 676:
		return "Braxi Mount"
	case 677:
		return "Kangon Mount"
	case 678:
		return "Erudite"
	case 679:
		return "Wurm Mount"
	case 680:
		return "Raptor Mount"
	case 681:
		return "Invisible Man"
	case 682:
		return "Whirligig"
	case 683:
		return "Gnomish Balloon"
	case 684:
		return "Gnomish Rocket Pack"
	case 685:
		return "Gnomish Hovering Transport"
	case 686:
		return "Selyrah"
	case 687:
		return "Goral"
	case 688:
		return "Braxi"
	case 689:
		return "Kangon"
	case 690:
		return "Invisible Man"
	case 691:
		return "Floating Tower"
	case 692:
		return "Explosive Cart"
	case 693:
		return "Blimp Ship"
	case 694:
		return "Tumbleweed"
	case 695:
		return "Alaran"
	case 696:
		return "Swinetor"
	case 697:
		return "Triumvirate"
	case 698:
		return "Hadal"
	case 699:
		return "Hovering Platform"
	case 700:
		return "Parasitic Scavenger"
	case 701:
		return "Grendlaen"
	case 702:
		return "Ship in a Bottle"
	case 703:
		return "Alaran Sentry Stone"
	case 704:
		return "Dervish"
	case 705:
		return "Regeneration Pool"
	case 706:
		return "Teleportation Stand"
	case 707:
		return "Relic Case"
	case 708:
		return "Alaran Ghost"
	case 709:
		return "Skystrider"
	case 710:
		return "Water Spout"
	case 711:
		return "Aviak Pull Along"
	case 712:
		return "Gelatinous Cube"
	case 713:
		return "Cat"
	case 714:
		return "Elk Head"
	case 715:
		return "Holgresh"
	case 716:
		return "Beetle"
	case 717:
		return "Vine Maw"
	case 718:
		return "Ratman"
	case 719:
		return "Fallen Knight"
	case 720:
		return "Flying Carpet"
	case 721:
		return "Carrier Hand"
	case 722:
		return "Akheva"
	case 723:
		return "Servant of Shadow"
	case 724:
		return "Luclin"
	}

	return "Unknown"
}
