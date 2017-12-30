package model

var raceNames = map[int64]string{
	1:   "Human",
	2:   "Barbarian",
	3:   "Erudite",
	4:   "Wood Elf",
	5:   "High Elf",
	6:   "Dark Elf",
	7:   "Half Elf",
	8:   "Dwarf",
	9:   "Troll",
	10:  "Ogre",
	11:  "Halfling",
	12:  "Gnome",
	13:  "Aviak",
	14:  "Werewolf",
	15:  "Brownie",
	16:  "Centaur",
	17:  "Golem",
	18:  "Giant",
	19:  "Trakanon",
	20:  "Venril Sathir",
	21:  "Evil Eye",
	22:  "Beetle",
	23:  "Kerran",
	24:  "Fish",
	25:  "Fairy",
	26:  "Froglok",
	27:  "Froglok",
	28:  "Fungusman",
	29:  "Gargoyle",
	30:  "Gasbag",
	31:  "Gelatinous Cube",
	32:  "Ghost",
	33:  "Ghoul",
	34:  "Bat",
	35:  "Eel",
	36:  "Rat",
	37:  "Snake",
	38:  "Spider",
	39:  "Gnoll",
	40:  "Goblin",
	41:  "Gorilla",
	42:  "Wolf",
	43:  "Bear",
	44:  "Guard",
	45:  "Demi Lich",
	46:  "Imp",
	47:  "Griffin",
	48:  "Kobold",
	49:  "Dragon",
	50:  "Lion",
	51:  "Lizard Man",
	52:  "Mimic",
	53:  "Minotaur",
	54:  "Orc",
	55:  "Beggar",
	56:  "Pixie",
	57:  "Drachnid",
	58:  "Solusek Ro",
	59:  "Goblin",
	60:  "Skeleton",
	61:  "Shark",
	62:  "Tunare",
	63:  "Tiger",
	64:  "Treant",
	65:  "Vampire",
	66:  "Rallos Zek",
	67:  "Human",
	68:  "Tentacle Terror",
	69:  "Will-O-Wisp",
	70:  "Zombie",
	71:  "Human",
	72:  "Ship",
	73:  "Launch",
	74:  "Piranha",
	75:  "Elemental",
	76:  "Puma",
	77:  "Dark Elf",
	78:  "Erudite",
	79:  "Bixie",
	80:  "Reanimated Hand",
	81:  "Halfling",
	82:  "Scarecrow",
	83:  "Skunk",
	84:  "Snake Elemental",
	85:  "Spectre",
	86:  "Sphinx",
	87:  "Armadillo",
	88:  "Clockwork Gnome",
	89:  "Drake",
	90:  "Barbarian",
	91:  "Alligator",
	92:  "Troll",
	93:  "Ogre",
	94:  "Dwarf",
	95:  "Cazic Thule",
	96:  "Cockatrice",
	97:  "Daisy Man",
	98:  "Vampire",
	99:  "Amygdalan",
	100: "Dervish",
	101: "Efreeti",
	102: "Tadpole",
	103: "Kedge",
	104: "Leech",
	105: "Swordfish",
	106: "Guard",
	107: "Mammoth",
	108: "Eye",
	109: "Wasp",
	110: "Mermaid",
	111: "Harpy",
	112: "Guard",
	113: "Drixie",
	114: "Ghost Ship",
	115: "Clam",
	116: "Seahorse",
	117: "Ghost",
	118: "Ghost",
	119: "Sabertooth",
	120: "Wolf",
	121: "Gorgon",
	122: "Dragon",
	123: "Innoruuk",
	124: "Unicorn",
	125: "Pegasus",
	126: "Djinn",
	127: "Invisible Man",
	128: "Iksar",
	129: "Scorpion",
	130: "Vah Shir",
	131: "Sarnak",
	132: "Draglock",
	133: "Drolvarg",
	134: "Mosquito",
	135: "Rhinoceros",
	136: "Xalgoz",
	137: "Goblin",
	138: "Yeti",
	139: "Iksar",
	140: "Giant",
	141: "Boat",
	142: "Object",
	143: "Tree",
	144: "Burynai",
	145: "Goo",
	146: "Sarnak Spirit",
	147: "Iksar Spirit",
	148: "Fish",
	149: "Scorpion",
	150: "Erollisi",
	151: "Tribunal",
	152: "Bertoxxulous",
	153: "Bristlebane",
	154: "Fay Drake",
	155: "Undead Sarnak",
	156: "Ratman",
	157: "Wyvern",
	158: "Wurm",
	159: "Devourer",
	160: "Iksar Golem",
	161: "Undead Iksar",
	162: "ManEating Plant",
	163: "Raptor",
	164: "Sarnak Golem",
	165: "Dragon",
	166: "Animated Hand",
	167: "Succulent",
	168: "Holgresh",
	169: "Brontotherium",
	170: "Snow Dervish",
	171: "Dire Wolf",
	172: "Manticore",
	173: "Totem",
	174: "Ice Spectre",
	175: "Enchanted Armor",
	176: "Snow Rabbit",
	177: "Walrus",
	178: "Geonid",
	181: "Yakkar",
	182: "Faun",
	183: "Coldain",
	184: "Dragon",
	185: "Hag",
	186: "Hippogriff",
	187: "Siren",
	188: "Giant",
	189: "Giant",
	190: "Othmir",
	191: "Ulthork",
	192: "Dragon",
	193: "Abhorrent",
	194: "Sea Turtle",
	195: "Dragon",
	196: "Dragon",
	197: "Ronnie Test",
	198: "Dragon",
	199: "Shik'Nar",
	200: "Rockhopper",
	201: "Underbulk",
	202: "Grimling",
	203: "Worm",
	204: "Evan Test",
	205: "Shadel",
	206: "Owlbear",
	207: "Rhino Beetle",
	208: "Vampire",
	209: "Earth Elemental",
	210: "Air Elemental",
	211: "Water Elemental",
	212: "Fire Elemental",
	213: "Wetfang Minnow",
	214: "Thought Horror",
	215: "Tegi",
	216: "Horse",
	217: "Shissar",
	218: "Fungal Fiend",
	219: "Vampire",
	220: "Stonegrabber",
	221: "Scarlet Cheetah",
	222: "Zelniak",
	223: "Lightcrawler",
	224: "Shade",
	225: "Sunflower",
	226: "Sun Revenant",
	227: "Shrieker",
	228: "Galorian",
	229: "Netherbian",
	230: "Akheva",
	231: "Grieg Veneficus",
	232: "Sonic Wolf",
	233: "Ground Shaker",
	234: "Vah Shir Skeleton",
	235: "Wretch",
	236: "Seru",
	237: "Recuso",
	238: "Vah Shir",
	239: "Guard",
	240: "Teleport Man",
	241: "Werewolf",
	242: "Nymph",
	243: "Dryad",
	244: "Treant",
	245: "Fly",
	246: "Tarew Marr",
	247: "Solusek Ro",
	248: "Clockwork Golem",
	249: "Clockwork Brain",
	250: "Banshee",
	251: "Guard of Justice",
	252: "Mini POM",
	253: "Diseased Fiend",
	254: "Solusek Ro Guard",
	255: "Bertoxxulous",
	256: "The Tribunal",
	257: "Terris Thule",
	258: "Vegerog",
	259: "Crocodile",
	260: "Bat",
	261: "Hraquis",
	262: "Tranquilion",
	263: "Tin Soldier",
	264: "Nightmare Wraith",
	265: "Malarian",
	266: "Knight of Pestilence",
	267: "Lepertoloth",
	268: "Bubonian",
	269: "Bubonian Underling",
	270: "Pusling",
	271: "Water Mephit",
	272: "Stormrider",
	273: "Junk Beast",
	274: "Broken Clockwork",
	275: "Giant Clockwork",
	276: "Clockwork Beetle",
	277: "Nightmare Goblin",
	278: "Karana",
	279: "Blood Raven",
	280: "Nightmare Gargoyle",
	281: "Mouth of Insanity",
	282: "Skeletal Horse",
	283: "Saryrn",
	284: "Fennin Ro",
	285: "Tormentor",
	286: "Soul Devourer",
	287: "Nightmare",
	288: "Rallos Zek",
	289: "Vallon Zek",
	290: "Tallon Zek",
	291: "Air Mephit",
	292: "Earth Mephit",
	293: "Fire Mephit",
	294: "Nightmare Mephit",
	295: "Zebuxoruk",
	296: "Mithaniel Marr",
	297: "Undead Knight",
	298: "The Rathe",
	299: "Xegony",
	300: "Fiend",
	301: "Test Object",
	302: "Crab",
	303: "Phoenix",
	304: "Dragon",
	305: "Bear",
	306: "Giant",
	307: "Giant",
	308: "Giant",
	309: "Giant",
	310: "Giant",
	311: "Giant",
	312: "Giant",
	313: "War Wraith",
	314: "Wrulon",
	315: "Kraken",
	316: "Poison Frog",
	317: "Nilborien",
	318: "Valorian",
	319: "War Boar",
	320: "Efreeti",
	321: "War Boar",
	322: "Valorian",
	323: "Animated Armor",
	324: "Undead Footman",
	325: "Rallos Zek Minion",
	326: "Arachnid",
	327: "Crystal Spider",
	328: "Zebuxoruk's Cage",
	329: "Bastion of Thunder Portal",
	330: "Froglok",
	331: "Troll",
	332: "Troll",
	333: "Troll",
	334: "Ghost",
	335: "Pirate",
	336: "Pirate",
	337: "Pirate",
	338: "Pirate",
	339: "Pirate",
	340: "Pirate",
	341: "Pirate",
	342: "Pirate",
	343: "Frog",
	344: "Troll Zombie",
	345: "Luggald",
	346: "Luggald",
	347: "Luggalds",
	348: "Drogmore",
	349: "Froglok Skeleton",
	350: "Undead Froglok",
	351: "Knight of Hate",
	352: "Arcanist of Hate",
	353: "Veksar",
	354: "Veksar",
	355: "Veksar",
	356: "Chokidai",
	357: "Undead Chokidai",
	358: "Undead Veksar",
	359: "Vampire",
	360: "Vampire",
	361: "Rujarkian Orc",
	362: "Bone Golem",
	363: "Synarcana",
	364: "Sand Elf",
	365: "Vampire",
	366: "Rujarkian Orc",
	367: "Skeleton",
	368: "Mummy",
	369: "Goblin",
	370: "Insect",
	371: "Froglok Ghost",
	372: "Dervish",
	373: "Shade",
	374: "Golem",
	375: "Evil Eye",
	376: "Box",
	377: "Barrel",
	378: "Chest",
	379: "Vase",
	380: "Table",
	381: "Weapon Rack",
	382: "Coffin",
	383: "Bones",
	384: "Jokester",
	385: "Nihil",
	386: "Trusik",
	387: "Stone Worker",
	388: "Hynid",
	389: "Turepta",
	390: "Cragbeast",
	391: "Stonemite",
	392: "Ukun",
	393: "Ixt",
	394: "Ikaav",
	395: "Aneuk",
	396: "Kyv",
	397: "Noc",
	398: "Ra`tuk",
	399: "Taneth",
	400: "Huvul",
	401: "Mutna",
	402: "Mastruq",
	403: "Taelosian",
	404: "Discord Ship",
	405: "Stone Worker",
	406: "Mata Muram",
	407: "Lightning Warrior",
	408: "Succubus",
	409: "Bazu",
	410: "Feran",
	411: "Pyrilen",
	412: "Chimera",
	413: "Dragorn",
	414: "Murkglider",
	415: "Rat",
	416: "Bat",
	417: "Gelidran",
	418: "Discordling",
	419: "Girplan",
	420: "Minotaur",
	421: "Dragorn Box",
	422: "Runed Orb",
	423: "Dragon Bones",
	424: "Muramite Armor Pile",
	425: "Crystal Shard",
	426: "Portal",
	427: "Coin Purse",
	428: "Rock Pile",
	429: "Murkglider Egg Sack",
	430: "Drake",
	431: "Dervish",
	432: "Drake",
	433: "Goblin",
	434: "Kirin",
	435: "Dragon",
	436: "Basilisk",
	437: "Dragon",
	438: "Dragon",
	439: "Puma",
	440: "Spider",
	441: "Spider Queen",
	442: "Animated Statue",
	445: "Dragon Egg",
	446: "Dragon Statue",
	447: "Lava Rock",
	448: "Animated Statue",
	449: "Spider Egg Sack",
	450: "Lava Spider",
	451: "Lava Spider Queen",
	452: "Dragon",
	453: "Giant",
	454: "Werewolf",
	455: "Kobold",
	456: "Sporali",
	457: "Gnomework",
	458: "Orc",
	459: "Corathus",
	460: "Coral",
	461: "Drachnid",
	462: "Drachnid Cocoon",
	463: "Fungus Patch",
	464: "Gargoyle",
	465: "Witheran",
	466: "Dark Lord",
	467: "Shiliskin",
	468: "Snake",
	469: "Evil Eye",
	470: "Minotaur",
	471: "Zombie",
	472: "Clockwork Boar",
	473: "Fairy",
	474: "Witheran",
	475: "Air Elemental",
	476: "Earth Elemental",
	477: "Fire Elemental",
	478: "Water Elemental",
	479: "Alligator",
	480: "Bear",
	481: "Scaled Wolf",
	482: "Wolf",
	483: "Spirit Wolf",
	484: "Skeleton",
	485: "Spectre",
	486: "Bolvirk",
	487: "Banshee",
	488: "Banshee",
	489: "Elddar",
	490: "Forest Giant",
	491: "Bone Golem",
	492: "Horse",
	493: "Pegasus",
	494: "Shambling Mound",
	495: "Scrykin",
	496: "Treant",
	497: "Vampire",
	498: "Ayonae Ro",
	499: "Sullon Zek",
	500: "Banner",
	501: "Flag",
	502: "Rowboat",
	503: "Bear Trap",
	504: "Clockwork Bomb",
	505: "Dynamite Keg",
	506: "Pressure Plate",
	507: "Puffer Spore",
	508: "Stone Ring",
	509: "Root Tentacle",
	510: "Runic Symbol",
	511: "Saltpetter Bomb",
	512: "Floating Skull",
	513: "Spike Trap",
	514: "Totem",
	515: "Web",
	516: "Wicker Basket",
	517: "Nightmare/Unicorn",
	518: "Horse",
	519: "Nightmare/Unicorn",
	520: "Bixie",
	521: "Centaur",
	522: "Drakkin",
	523: "Giant",
	524: "Gnoll",
	525: "Griffin",
	526: "Giant Shade",
	527: "Harpy",
	528: "Mammoth",
	529: "Satyr",
	530: "Dragon",
	531: "Dragon",
	532: "Dyn'Leth",
	533: "Boat",
	534: "Weapon Rack",
	535: "Armor Rack",
	536: "Honey Pot",
	537: "Jum Jum Bucket",
	538: "Toolbox",
	539: "Stone Jug",
	540: "Small Plant",
	541: "Medium Plant",
	542: "Tall Plant",
	543: "Wine Cask",
	544: "Elven Boat",
	545: "Gnomish Boat",
	546: "Barrel Barge Ship",
	547: "Goo",
	548: "Goo",
	549: "Goo",
	550: "Merchant Ship",
	551: "Pirate Ship",
	552: "Ghost Ship",
	553: "Banner",
	554: "Banner",
	555: "Banner",
	556: "Banner",
	557: "Banner",
	558: "Aviak",
	559: "Beetle",
	560: "Gorilla",
	561: "Kedge",
	562: "Kerran",
	563: "Shissar",
	564: "Siren",
	565: "Sphinx",
	566: "Human",
	567: "Campfire",
	568: "Brownie",
	569: "Dragon",
	570: "Exoskeleton",
	571: "Ghoul",
	572: "Clockwork Guardian",
	573: "Mantrap",
	574: "Minotaur",
	575: "Scarecrow",
	576: "Shade",
	577: "Rotocopter",
	578: "Tentacle Terror",
	579: "Wereorc",
	580: "Worg",
	581: "Wyvern",
	582: "Chimera",
	583: "Kirin",
	584: "Puma",
	585: "Boulder",
	586: "Banner",
	587: "Elven Ghost",
	588: "Human Ghost",
	589: "Chest",
	590: "Chest",
	591: "Crystal",
	592: "Coffin",
	593: "Guardian CPU",
	594: "Worg",
	595: "Mansion",
	596: "Floating Island",
	597: "Cragslither",
	598: "Wrulon",
	599: "Spell Particle 1",
	600: "Invisible Man of Zomm",
	601: "Robocopter of Zomm",
	602: "Burynai",
	603: "Frog",
	604: "Dracolich",
	605: "Iksar Ghost",
	606: "Iksar Skeleton",
	607: "Mephit",
	608: "Muddite",
	609: "Raptor",
	610: "Sarnak",
	611: "Scorpion",
	612: "Tsetsian",
	613: "Wurm",
	614: "Nekhon",
	615: "Hydra Crystal",
	616: "Crystal Sphere",
	617: "Gnoll",
	618: "Sokokar",
	619: "Stone Pylon",
	620: "Demon Vulture",
	621: "Wagon",
	622: "God of Discord",
	623: "Feran Mount",
	624: "Ogre NPC Male",
	625: "Sokokar Mount",
	626: "Giant",
	627: "Sokokar",
	628: "10th Anniversary Banner",
	629: "10th Anniversary Cake",
	630: "Wine Cask",
	631: "Hydra Mount",
	632: "Hydra NPC",
	633: "Wedding Flowers",
	634: "Wedding Arbor",
	635: "Wedding Altar",
	636: "Powder Keg",
	637: "Apexus",
	638: "Bellikos",
	639: "Brell's First Creation",
	640: "Brell",
	641: "Crystalskin Ambuloid",
	642: "Cliknar Queen",
	643: "Cliknar Soldier",
	644: "Cliknar Worker",
	645: "Coldain",
	646: "Coldain",
	647: "Crystalskin Sessiloid",
	648: "Genari",
	649: "Gigyn",
	650: "Greken Young Adult",
	651: "Greken Young",
	652: "Cliknar Mount",
	653: "Telmira",
	654: "Spider Mount",
	655: "Bear Mount",
	656: "Rat Mount Mystery Race",
	657: "Sessiloid Mount",
	658: "Morell Thule",
	659: "Marionette",
	660: "Book Dervish",
	661: "Topiary Lion",
	662: "Rotdog",
	663: "Amygdalan",
	664: "Sandman",
	665: "Grandfather Clock",
	666: "Gingerbread Man",
	667: "Royal Guard",
	668: "Rabbit",
	669: "Blind Dreamer",
	670: "Cazic Thule",
	671: "Topiary Lion Mount",
	672: "Rot Dog Mount",
	673: "Goral Mount",
	674: "Selyrah Mount",
	675: "Sclera Mount",
	676: "Braxi Mount",
	677: "Kangon Mount",
	678: "Erudite",
	679: "Wurm Mount",
	680: "Raptor Mount",
	681: "Invisible Man",
	682: "Whirligig",
	683: "Gnomish Balloon",
	684: "Gnomish Rocket Pack",
	685: "Gnomish Hovering Transport",
	686: "Selyrah",
	687: "Goral",
	688: "Braxi",
	689: "Kangon",
	690: "Invisible Man",
	691: "Floating Tower",
	692: "Explosive Cart",
	693: "Blimp Ship",
	694: "Tumbleweed",
	695: "Alaran",
	696: "Swinetor",
	697: "Triumvirate",
	698: "Hadal",
	699: "Hovering Platform",
	700: "Parasitic Scavenger",
	701: "Grendlaen",
	702: "Ship in a Bottle",
	703: "Alaran Sentry Stone",
	704: "Dervish",
	705: "Regeneration Pool",
	706: "Teleportation Stand",
	707: "Unknown (707)",
	708: "Alaran Ghost",
	709: "Skystrider",
	710: "Water Spout",
	711: "Aviak Pull Along",
	712: "Gelatinous Cube",
	713: "Cat",
	714: "Elk Head",
	715: "Holgresh",
	716: "Beetle",
	717: "Vine Maw",
	718: "Ratman",
	719: "Fallen Knight",
	720: "Flying Carpet",
	721: "Carrier Hand",
	722: "Akheva",
	723: "Servant of Shadow",
	724: "Luclin",
}

var raceIcons = map[int64]string{
	1:   "xa-player",       //human
	2:   "xa-fox",          //barbarian
	3:   "xa-book",         //erudite
	4:   "xa-pine-tree",    //woodelf
	5:   "xa-tesla",        //helf
	6:   "xa-bleeding-eye", //delf
	7:   "xa-aware",        //halfelf
	8:   "xa-beer",         //dwarf
	9:   "xa-bird-mask",    //troll
	10:  "xa-muscle-fat",   //ogre
	11:  "xa-footprint",    //halfling
	12:  "xa-gears",        //gnome
	13:  "xa-octopus",      //"Aviak"
	14:  "xa-octopus",      //"Werewolf"
	15:  "xa-octopus",      //"Brownie"
	16:  "xa-octopus",      //"Centaur"
	17:  "xa-octopus",      //"Golem"
	18:  "xa-octopus",      //"Giant"
	19:  "xa-octopus",      //"Trakanon"
	20:  "xa-octopus",      //"Venril Sathir"
	21:  "xa-octopus",      //"Evil Eye"
	22:  "xa-octopus",      //"Beetle"
	23:  "xa-octopus",      //"Kerran"
	24:  "xa-octopus",      //"Fish"
	25:  "xa-octopus",      //"Fairy"
	26:  "xa-water-drop",   //froglok
	27:  "xa-octopus",      //"Froglok"
	28:  "xa-octopus",      //"Fungusman"
	29:  "xa-octopus",      //"Gargoyle"
	30:  "xa-octopus",      //"Gasbag"
	31:  "xa-octopus",      //"Gelatinous Cube"
	32:  "xa-octopus",      //"Ghost"
	33:  "xa-octopus",      //"Ghoul"
	34:  "xa-octopus",      //"Bat"
	35:  "xa-octopus",      //"Eel"
	36:  "xa-octopus",      //"Rat"
	37:  "xa-octopus",      //"Snake"
	38:  "xa-octopus",      //"Spider"
	39:  "xa-octopus",      //"Gnoll"
	40:  "xa-octopus",      //"Goblin"
	41:  "xa-octopus",      //"Gorilla"
	42:  "xa-octopus",      //"Wolf"
	43:  "xa-octopus",      //"Bear"
	44:  "xa-octopus",      //"Guard"
	45:  "xa-octopus",      //"Demi Lich"
	46:  "xa-octopus",      //"Imp"
	47:  "xa-octopus",      //"Griffin"
	48:  "xa-octopus",      //"Kobold"
	49:  "xa-wyvern",       //"Dragon"
	50:  "xa-octopus",      //"Lion"
	51:  "xa-octopus",      //"Lizard Man"
	52:  "xa-octopus",      //"Mimic"
	53:  "xa-octopus",      //"Minotaur"
	54:  "xa-octopus",      //"Orc"
	55:  "xa-octopus",      //"Beggar"
	56:  "xa-octopus",      //"Pixie"
	57:  "xa-octopus",      //"Drachnid"
	58:  "xa-octopus",      //"Solusek Ro"
	59:  "xa-octopus",      //"Goblin"
	60:  "xa-octopus",      //"Skeleton"
	61:  "xa-octopus",      //"Shark"
	62:  "xa-octopus",      //"Tunare"
	63:  "xa-octopus",      //"Tiger"
	64:  "xa-octopus",      //"Treant"
	65:  "xa-octopus",      //"Vampire"
	66:  "xa-octopus",      //"Rallos Zek"
	67:  "xa-octopus",      //"Human"
	68:  "xa-octopus",      //"Tentacle Terror"
	69:  "xa-octopus",      //"Will-O-Wisp"
	70:  "xa-octopus",      //"Zombie"
	71:  "xa-octopus",      //"Human"
	72:  "xa-octopus",      //"Ship"
	73:  "xa-octopus",      //"Launch"
	74:  "xa-octopus",      //"Piranha"
	75:  "xa-octopus",      //"Elemental"
	76:  "xa-octopus",      //"Puma"
	77:  "xa-octopus",      //"Dark Elf"
	78:  "xa-octopus",      //"Erudite"
	79:  "xa-octopus",      //"Bixie"
	80:  "xa-octopus",      //"Reanimated Hand"
	81:  "xa-octopus",      //"Halfling"
	82:  "xa-octopus",      //"Scarecrow"
	83:  "xa-octopus",      //"Skunk"
	84:  "xa-octopus",      //"Snake Elemental"
	85:  "xa-octopus",      //"Spectre"
	86:  "xa-octopus",      //"Sphinx"
	87:  "xa-octopus",      //"Armadillo"
	88:  "xa-octopus",      //"Clockwork Gnome"
	89:  "xa-octopus",      //"Drake"
	90:  "xa-octopus",      //"Barbarian"
	91:  "xa-octopus",      //"Alligator"
	92:  "xa-octopus",      //"Troll"
	93:  "xa-octopus",      //"Ogre"
	94:  "xa-octopus",      //"Dwarf"
	95:  "xa-octopus",      //"Cazic Thule"
	96:  "xa-octopus",      //"Cockatrice"
	97:  "xa-octopus",      //"Daisy Man"
	98:  "xa-octopus",      //"Vampire"
	99:  "xa-octopus",      //"Amygdalan"
	100: "xa-octopus",      //"Dervish"
	101: "xa-octopus",      //"Efreeti"
	102: "xa-octopus",      //"Tadpole"
	103: "xa-octopus",      //"Kedge"
	104: "xa-octopus",      //"Leech"
	105: "xa-octopus",      //"Swordfish"
	106: "xa-octopus",      //"Guard"
	107: "xa-octopus",      //"Mammoth"
	108: "xa-octopus",      //"Eye"
	109: "xa-octopus",      //"Wasp"
	110: "xa-octopus",      //"Mermaid"
	111: "xa-octopus",      //"Harpy"
	112: "xa-octopus",      //"Guard"
	113: "xa-octopus",      //"Drixie"
	114: "xa-octopus",      //"Ghost Ship"
	115: "xa-octopus",      //"Clam"
	116: "xa-octopus",      //"Seahorse"
	117: "xa-octopus",      //"Ghost"
	118: "xa-octopus",      //"Ghost"
	119: "xa-octopus",      //"Sabertooth"
	120: "xa-octopus",      //"Wolf"
	121: "xa-octopus",      //"Gorgon"
	122: "xa-wyvern",       //"Dragon"
	123: "xa-octopus",      //"Innoruuk"
	124: "xa-octopus",      //"Unicorn"
	125: "xa-octopus",      //"Pegasus"
	126: "xa-octopus",      //"Djinn"
	127: "xa-octopus",      //"Invisible Man"
	128: "xa-gecko",        //iksar
	129: "xa-octopus",      //"Scorpion"
	130: "xa-lion",         //vahshir
	131: "xa-octopus",      //"Sarnak"
	132: "xa-octopus",      //"Draglock"
	133: "xa-octopus",      //"Drolvarg"
	134: "xa-octopus",      //"Mosquito"
	135: "xa-octopus",      //"Rhinoceros"
	136: "xa-octopus",      //"Xalgoz"
	137: "xa-octopus",      //"Goblin"
	138: "xa-octopus",      //"Yeti"
	139: "xa-octopus",      //"Iksar"
	140: "xa-octopus",      //"Giant"
	141: "xa-octopus",      //"Boat"
	142: "xa-octopus",      //"Object"
	143: "xa-octopus",      //"Tree"
	144: "xa-octopus",      //"Burynai"
	145: "xa-octopus",      //"Goo"
	146: "xa-octopus",      //"Sarnak Spirit"
	147: "xa-octopus",      //"Iksar Spirit"
	148: "xa-octopus",      //"Fish"
	149: "xa-octopus",      //"Scorpion"
	150: "xa-octopus",      //"Erollisi"
	151: "xa-octopus",      //"Tribunal"
	152: "xa-octopus",      //"Bertoxxulous"
	153: "xa-octopus",      //"Bristlebane"
	154: "xa-octopus",      //"Fay Drake"
	155: "xa-octopus",      //"Undead Sarnak"
	156: "xa-octopus",      //"Ratman"
	157: "xa-octopus",      //"Wyvern"
	158: "xa-octopus",      //"Wurm"
	159: "xa-octopus",      //"Devourer"
	160: "xa-octopus",      //"Iksar Golem"
	161: "xa-octopus",      //"Undead Iksar"
	162: "xa-octopus",      //"ManEating Plant"
	163: "xa-octopus",      //"Raptor"
	164: "xa-octopus",      //"Sarnak Golem"
	165: "xa-wyvern",       //"Dragon"
	166: "xa-octopus",      //"Animated Hand"
	167: "xa-octopus",      //"Succulent"
	168: "xa-octopus",      //"Holgresh"
	169: "xa-octopus",      //"Brontotherium"
	170: "xa-octopus",      //"Snow Dervish"
	171: "xa-octopus",      //"Dire Wolf"
	172: "xa-octopus",      //"Manticore"
	173: "xa-octopus",      //"Totem"
	174: "xa-octopus",      //"Ice Spectre"
	175: "xa-octopus",      //"Enchanted Armor"
	176: "xa-octopus",      //"Snow Rabbit"
	177: "xa-octopus",      //"Walrus"
	178: "xa-octopus",      //"Geonid"
	181: "xa-octopus",      //"Yakkar"
	182: "xa-octopus",      //"Faun"
	183: "xa-octopus",      //"Coldain"
	184: "xa-wyvern",       //"Dragon"
	185: "xa-octopus",      //"Hag"
	186: "xa-octopus",      //"Hippogriff"
	187: "xa-octopus",      //"Siren"
	188: "xa-octopus",      //"Giant"
	189: "xa-octopus",      //"Giant"
	190: "xa-octopus",      //"Othmir"
	191: "xa-octopus",      //"Ulthork"
	192: "xa-wyvern",       //"Dragon"
	193: "xa-octopus",      //"Abhorrent"
	194: "xa-octopus",      //"Sea Turtle"
	195: "xa-wyvern",       //"Dragon"
	196: "xa-wyvern",       //"Dragon"
	197: "xa-octopus",      //"Ronnie Test"
	198: "xa-wyvern",       //"Dragon"
	199: "xa-octopus",      //"Shik'Nar"
	200: "xa-octopus",      //"Rockhopper"
	201: "xa-octopus",      //"Underbulk"
	202: "xa-octopus",      //"Grimling"
	203: "xa-octopus",      //"Worm"
	204: "xa-octopus",      //"Evan Test"
	205: "xa-octopus",      //"Shadel"
	206: "xa-octopus",      //"Owlbear"
	207: "xa-octopus",      //"Rhino Beetle"
	208: "xa-octopus",      //"Vampire"
	209: "xa-octopus",      //"Earth Elemental"
	210: "xa-octopus",      //"Air Elemental"
	211: "xa-octopus",      //"Water Elemental"
	212: "xa-octopus",      //"Fire Elemental"
	213: "xa-octopus",      //"Wetfang Minnow"
	214: "xa-octopus",      //"Thought Horror"
	215: "xa-octopus",      //"Tegi"
	216: "xa-octopus",      //"Horse"
	217: "xa-octopus",      //"Shissar"
	218: "xa-octopus",      //"Fungal Fiend"
	219: "xa-octopus",      //"Vampire"
	220: "xa-octopus",      //"Stonegrabber"
	221: "xa-octopus",      //"Scarlet Cheetah"
	222: "xa-octopus",      //"Zelniak"
	223: "xa-octopus",      //"Lightcrawler"
	224: "xa-octopus",      //"Shade"
	225: "xa-octopus",      //"Sunflower"
	226: "xa-octopus",      //"Sun Revenant"
	227: "xa-octopus",      //"Shrieker"
	228: "xa-octopus",      //"Galorian"
	229: "xa-octopus",      //"Netherbian"
	230: "xa-octopus",      //"Akheva"
	231: "xa-octopus",      //"Grieg Veneficus"
	232: "xa-octopus",      //"Sonic Wolf"
	233: "xa-octopus",      //"Ground Shaker"
	234: "xa-octopus",      //"Vah Shir Skeleton"
	235: "xa-octopus",      //"Wretch"
	236: "xa-octopus",      //"Seru"
	237: "xa-octopus",      //"Recuso"
	238: "xa-octopus",      //"Vah Shir"
	239: "xa-octopus",      //"Guard"
	240: "xa-octopus",      //"Teleport Man"
	241: "xa-octopus",      //"Werewolf"
	242: "xa-octopus",      //"Nymph"
	243: "xa-octopus",      //"Dryad"
	244: "xa-octopus",      //"Treant"
	245: "xa-octopus",      //"Fly"
	246: "xa-octopus",      //"Tarew Marr"
	247: "xa-octopus",      //"Solusek Ro"
	248: "xa-octopus",      //"Clockwork Golem"
	249: "xa-octopus",      //"Clockwork Brain"
	250: "xa-octopus",      //"Banshee"
	251: "xa-octopus",      //"Guard of Justice"
	252: "xa-octopus",      //"Mini POM"
	253: "xa-octopus",      //"Diseased Fiend"
	254: "xa-octopus",      //"Solusek Ro Guard"
	255: "xa-octopus",      //"Bertoxxulous"
	256: "xa-octopus",      //"The Tribunal"
	257: "xa-octopus",      //"Terris Thule"
	258: "xa-octopus",      //"Vegerog"
	259: "xa-octopus",      //"Crocodile"
	260: "xa-octopus",      //"Bat"
	261: "xa-octopus",      //"Hraquis"
	262: "xa-octopus",      //"Tranquilion"
	263: "xa-octopus",      //"Tin Soldier"
	264: "xa-octopus",      //"Nightmare Wraith"
	265: "xa-octopus",      //"Malarian"
	266: "xa-octopus",      //"Knight of Pestilence"
	267: "xa-octopus",      //"Lepertoloth"
	268: "xa-octopus",      //"Bubonian"
	269: "xa-octopus",      //"Bubonian Underling"
	270: "xa-octopus",      //"Pusling"
	271: "xa-octopus",      //"Water Mephit"
	272: "xa-octopus",      //"Stormrider"
	273: "xa-octopus",      //"Junk Beast"
	274: "xa-octopus",      //"Broken Clockwork"
	275: "xa-octopus",      //"Giant Clockwork"
	276: "xa-octopus",      //"Clockwork Beetle"
	277: "xa-octopus",      //"Nightmare Goblin"
	278: "xa-octopus",      //"Karana"
	279: "xa-octopus",      //"Blood Raven"
	280: "xa-octopus",      //"Nightmare Gargoyle"
	281: "xa-octopus",      //"Mouth of Insanity"
	282: "xa-octopus",      //"Skeletal Horse"
	283: "xa-octopus",      //"Saryrn"
	284: "xa-octopus",      //"Fennin Ro"
	285: "xa-octopus",      //"Tormentor"
	286: "xa-octopus",      //"Soul Devourer"
	287: "xa-octopus",      //"Nightmare"
	288: "xa-octopus",      //"Rallos Zek"
	289: "xa-octopus",      //"Vallon Zek"
	290: "xa-octopus",      //"Tallon Zek"
	291: "xa-octopus",      //"Air Mephit"
	292: "xa-octopus",      //"Earth Mephit"
	293: "xa-octopus",      //"Fire Mephit"
	294: "xa-octopus",      //"Nightmare Mephit"
	295: "xa-octopus",      //"Zebuxoruk"
	296: "xa-octopus",      //"Mithaniel Marr"
	297: "xa-octopus",      //"Undead Knight"
	298: "xa-octopus",      //"The Rathe"
	299: "xa-octopus",      //"Xegony"
	300: "xa-octopus",      //"Fiend"
	301: "xa-octopus",      //"Test Object"
	302: "xa-octopus",      //"Crab"
	303: "xa-octopus",      //"Phoenix"
	304: "xa-wyvern",       //"Dragon"
	305: "xa-octopus",      //"Bear"
	306: "xa-octopus",      //"Giant"
	307: "xa-octopus",      //"Giant"
	308: "xa-octopus",      //"Giant"
	309: "xa-octopus",      //"Giant"
	310: "xa-octopus",      //"Giant"
	311: "xa-octopus",      //"Giant"
	312: "xa-octopus",      //"Giant"
	313: "xa-octopus",      //"War Wraith"
	314: "xa-octopus",      //"Wrulon"
	315: "xa-octopus",      //"Kraken"
	316: "xa-octopus",      //"Poison Frog"
	317: "xa-octopus",      //"Nilborien"
	318: "xa-octopus",      //"Valorian"
	319: "xa-octopus",      //"War Boar"
	320: "xa-octopus",      //"Efreeti"
	321: "xa-octopus",      //"War Boar"
	322: "xa-octopus",      //"Valorian"
	323: "xa-octopus",      //"Animated Armor"
	324: "xa-octopus",      //"Undead Footman"
	325: "xa-octopus",      //"Rallos Zek Minion"
	326: "xa-octopus",      //"Arachnid"
	327: "xa-octopus",      //"Crystal Spider"
	328: "xa-octopus",      //"Zebuxoruk's Cage"
	329: "xa-octopus",      //"Bastion of Thunder Portal"
	330: "xa-octopus",      //"Froglok"
	331: "xa-octopus",      //"Troll"
	332: "xa-octopus",      //"Troll"
	333: "xa-octopus",      //"Troll"
	334: "xa-octopus",      //"Ghost"
	335: "xa-octopus",      //"Pirate"
	336: "xa-octopus",      //"Pirate"
	337: "xa-octopus",      //"Pirate"
	338: "xa-octopus",      //"Pirate"
	339: "xa-octopus",      //"Pirate"
	340: "xa-octopus",      //"Pirate"
	341: "xa-octopus",      //"Pirate"
	342: "xa-octopus",      //"Pirate"
	343: "xa-octopus",      //"Frog"
	344: "xa-octopus",      //"Troll Zombie"
	345: "xa-octopus",      //"Luggald"
	346: "xa-octopus",      //"Luggald"
	347: "xa-octopus",      //"Luggalds"
	348: "xa-octopus",      //"Drogmore"
	349: "xa-octopus",      //"Froglok Skeleton"
	350: "xa-octopus",      //"Undead Froglok"
	351: "xa-octopus",      //"Knight of Hate"
	352: "xa-octopus",      //"Arcanist of Hate"
	353: "xa-octopus",      //"Veksar"
	354: "xa-octopus",      //"Veksar"
	355: "xa-octopus",      //"Veksar"
	356: "xa-octopus",      //"Chokidai"
	357: "xa-octopus",      //"Undead Chokidai"
	358: "xa-octopus",      //"Undead Veksar"
	359: "xa-octopus",      //"Vampire"
	360: "xa-octopus",      //"Vampire"
	361: "xa-octopus",      //"Rujarkian Orc"
	362: "xa-octopus",      //"Bone Golem"
	363: "xa-octopus",      //"Synarcana"
	364: "xa-octopus",      //"Sand Elf"
	365: "xa-octopus",      //"Vampire"
	366: "xa-octopus",      //"Rujarkian Orc"
	367: "xa-octopus",      //"Skeleton"
	368: "xa-octopus",      //"Mummy"
	369: "xa-octopus",      //"Goblin"
	370: "xa-octopus",      //"Insect"
	371: "xa-octopus",      //"Froglok Ghost"
	372: "xa-octopus",      //"Dervish"
	373: "xa-octopus",      //"Shade"
	374: "xa-octopus",      //"Golem"
	375: "xa-octopus",      //"Evil Eye"
	376: "xa-octopus",      //"Box"
	377: "xa-octopus",      //"Barrel"
	378: "xa-octopus",      //"Chest"
	379: "xa-octopus",      //"Vase"
	380: "xa-octopus",      //"Table"
	381: "xa-octopus",      //"Weapon Rack"
	382: "xa-octopus",      //"Coffin"
	383: "xa-octopus",      //"Bones"
	384: "xa-octopus",      //"Jokester"
	385: "xa-octopus",      //"Nihil"
	386: "xa-octopus",      //"Trusik"
	387: "xa-octopus",      //"Stone Worker"
	388: "xa-octopus",      //"Hynid"
	389: "xa-octopus",      //"Turepta"
	390: "xa-octopus",      //"Cragbeast"
	391: "xa-octopus",      //"Stonemite"
	392: "xa-octopus",      //"Ukun"
	393: "xa-octopus",      //"Ixt"
	394: "xa-octopus",      //"Ikaav"
	395: "xa-octopus",      //"Aneuk"
	396: "xa-octopus",      //"Kyv"
	397: "xa-octopus",      //"Noc"
	398: "xa-octopus",      //"Ra`tuk"
	399: "xa-octopus",      //"Taneth"
	400: "xa-octopus",      //"Huvul"
	401: "xa-octopus",      //"Mutna"
	402: "xa-octopus",      //"Mastruq"
	403: "xa-octopus",      //"Taelosian"
	404: "xa-octopus",      //"Discord Ship"
	405: "xa-octopus",      //"Stone Worker"
	406: "xa-octopus",      //"Mata Muram"
	407: "xa-octopus",      //"Lightning Warrior"
	408: "xa-octopus",      //"Succubus"
	409: "xa-octopus",      //"Bazu"
	410: "xa-octopus",      //"Feran"
	411: "xa-octopus",      //"Pyrilen"
	412: "xa-octopus",      //"Chimera"
	413: "xa-octopus",      //"Dragorn"
	414: "xa-octopus",      //"Murkglider"
	415: "xa-octopus",      //"Rat"
	416: "xa-octopus",      //"Bat"
	417: "xa-octopus",      //"Gelidran"
	418: "xa-octopus",      //"Discordling"
	419: "xa-octopus",      //"Girplan"
	420: "xa-octopus",      //"Minotaur"
	421: "xa-octopus",      //"Dragorn Box"
	422: "xa-octopus",      //"Runed Orb"
	423: "xa-wyvern",       //"Dragon Bones"
	424: "xa-octopus",      //"Muramite Armor Pile"
	425: "xa-octopus",      //"Crystal Shard"
	426: "xa-octopus",      //"Portal"
	427: "xa-octopus",      //"Coin Purse"
	428: "xa-octopus",      //"Rock Pile"
	429: "xa-octopus",      //"Murkglider Egg Sack"
	430: "xa-octopus",      //"Drake"
	431: "xa-octopus",      //"Dervish"
	432: "xa-octopus",      //"Drake"
	433: "xa-octopus",      //"Goblin"
	434: "xa-octopus",      //"Kirin"
	435: "xa-wyvern",       //"Dragon"
	436: "xa-octopus",      //"Basilisk"
	437: "xa-wyvern",       //"Dragon"
	438: "xa-wyvern",       //"Dragon"
	439: "xa-octopus",      //"Puma"
	440: "xa-octopus",      //"Spider"
	441: "xa-octopus",      //"Spider Queen"
	442: "xa-octopus",      //"Animated Statue"
	445: "xa-egg",          //"Dragon Egg"
	446: "xa-wyvern",       //"Dragon Statue"
	447: "xa-octopus",      //"Lava Rock"
	448: "xa-octopus",      //"Animated Statue"
	449: "xa-octopus",      //"Spider Egg Sack"
	450: "xa-octopus",      //"Lava Spider"
	451: "xa-octopus",      //"Lava Spider Queen"
	452: "xa-wyvern",       //"Dragon"
	453: "xa-octopus",      //"Giant"
	454: "xa-octopus",      //"Werewolf"
	455: "xa-octopus",      //"Kobold"
	456: "xa-octopus",      //"Sporali"
	457: "xa-octopus",      //"Gnomework"
	458: "xa-octopus",      //"Orc"
	459: "xa-octopus",      //"Corathus"
	460: "xa-octopus",      //"Coral"
	461: "xa-octopus",      //"Drachnid"
	462: "xa-octopus",      //"Drachnid Cocoon"
	463: "xa-octopus",      //"Fungus Patch"
	464: "xa-octopus",      //"Gargoyle"
	465: "xa-octopus",      //"Witheran"
	466: "xa-octopus",      //"Dark Lord"
	467: "xa-octopus",      //"Shiliskin"
	468: "xa-octopus",      //"Snake"
	469: "xa-octopus",      //"Evil Eye"
	470: "xa-octopus",      //"Minotaur"
	471: "xa-octopus",      //"Zombie"
	472: "xa-octopus",      //"Clockwork Boar"
	473: "xa-octopus",      //"Fairy"
	474: "xa-octopus",      //"Witheran"
	475: "xa-octopus",      //"Air Elemental"
	476: "xa-octopus",      //"Earth Elemental"
	477: "xa-octopus",      //"Fire Elemental"
	478: "xa-octopus",      //"Water Elemental"
	479: "xa-octopus",      //"Alligator"
	480: "xa-octopus",      //"Bear"
	481: "xa-octopus",      //"Scaled Wolf"
	482: "xa-octopus",      //"Wolf"
	483: "xa-octopus",      //"Spirit Wolf"
	484: "xa-octopus",      //"Skeleton"
	485: "xa-octopus",      //"Spectre"
	486: "xa-octopus",      //"Bolvirk"
	487: "xa-octopus",      //"Banshee"
	488: "xa-octopus",      //"Banshee"
	489: "xa-octopus",      //"Elddar"
	490: "xa-octopus",      //"Forest Giant"
	491: "xa-octopus",      //"Bone Golem"
	492: "xa-octopus",      //"Horse"
	493: "xa-octopus",      //"Pegasus"
	494: "xa-octopus",      //"Shambling Mound"
	495: "xa-octopus",      //"Scrykin"
	496: "xa-octopus",      //"Treant"
	497: "xa-octopus",      //"Vampire"
	498: "xa-octopus",      //"Ayonae Ro"
	499: "xa-octopus",      //"Sullon Zek"
	500: "xa-octopus",      //"Banner"
	501: "xa-octopus",      //"Flag"
	502: "xa-octopus",      //"Rowboat"
	503: "xa-octopus",      //"Bear Trap"
	504: "xa-octopus",      //"Clockwork Bomb"
	505: "xa-octopus",      //"Dynamite Keg"
	506: "xa-octopus",      //"Pressure Plate"
	507: "xa-octopus",      //"Puffer Spore"
	508: "xa-octopus",      //"Stone Ring"
	509: "xa-octopus",      //"Root Tentacle"
	510: "xa-octopus",      //"Runic Symbol"
	511: "xa-octopus",      //"Saltpetter Bomb"
	512: "xa-octopus",      //"Floating Skull"
	513: "xa-octopus",      //"Spike Trap"
	514: "xa-octopus",      //"Totem"
	515: "xa-octopus",      //"Web"
	516: "xa-octopus",      //"Wicker Basket"
	517: "xa-octopus",      //"Nightmare/Unicorn"
	518: "xa-octopus",      //"Horse"
	519: "xa-octopus",      //"Nightmare/Unicorn"
	520: "xa-octopus",      //"Bixie"
	521: "xa-octopus",      //"Centaur"
	522: "xa-wyvern",       //drakkin
	523: "xa-octopus",      //"Giant"
	524: "xa-octopus",      //"Gnoll"
	525: "xa-octopus",      //"Griffin"
	526: "xa-octopus",      //"Giant Shade"
	527: "xa-octopus",      //"Harpy"
	528: "xa-octopus",      //"Mammoth"
	529: "xa-octopus",      //"Satyr"
	530: "xa-wyvern",       //"Dragon"
	531: "xa-wyvern",       //"Dragon"
	532: "xa-octopus",      //"Dyn'Leth"
	533: "xa-octopus",      //"Boat"
	534: "xa-octopus",      //"Weapon Rack"
	535: "xa-octopus",      //"Armor Rack"
	536: "xa-octopus",      //"Honey Pot"
	537: "xa-octopus",      //"Jum Jum Bucket"
	538: "xa-octopus",      //"Toolbox"
	539: "xa-octopus",      //"Stone Jug"
	540: "xa-octopus",      //"Small Plant"
	541: "xa-octopus",      //"Medium Plant"
	542: "xa-octopus",      //"Tall Plant"
	543: "xa-octopus",      //"Wine Cask"
	544: "xa-octopus",      //"Elven Boat"
	545: "xa-octopus",      //"Gnomish Boat"
	546: "xa-octopus",      //"Barrel Barge Ship"
	547: "xa-octopus",      //"Goo"
	548: "xa-octopus",      //"Goo"
	549: "xa-octopus",      //"Goo"
	550: "xa-octopus",      //"Merchant Ship"
	551: "xa-octopus",      //"Pirate Ship"
	552: "xa-octopus",      //"Ghost Ship"
	553: "xa-octopus",      //"Banner"
	554: "xa-octopus",      //"Banner"
	555: "xa-octopus",      //"Banner"
	556: "xa-octopus",      //"Banner"
	557: "xa-octopus",      //"Banner"
	558: "xa-octopus",      //"Aviak"
	559: "xa-octopus",      //"Beetle"
	560: "xa-octopus",      //"Gorilla"
	561: "xa-octopus",      //"Kedge"
	562: "xa-octopus",      //"Kerran"
	563: "xa-octopus",      //"Shissar"
	564: "xa-octopus",      //"Siren"
	565: "xa-octopus",      //"Sphinx"
	566: "xa-octopus",      //"Human"
	567: "xa-octopus",      //"Campfire"
	568: "xa-octopus",      //"Brownie"
	569: "xa-wyvern",       //"Dragon"
	570: "xa-octopus",      //"Exoskeleton"
	571: "xa-octopus",      //"Ghoul"
	572: "xa-octopus",      //"Clockwork Guardian"
	573: "xa-octopus",      //"Mantrap"
	574: "xa-octopus",      //"Minotaur"
	575: "xa-octopus",      //"Scarecrow"
	576: "xa-octopus",      //"Shade"
	577: "xa-octopus",      //"Rotocopter"
	578: "xa-octopus",      //"Tentacle Terror"
	579: "xa-octopus",      //"Wereorc"
	580: "xa-octopus",      //"Worg"
	581: "xa-octopus",      //"Wyvern"
	582: "xa-octopus",      //"Chimera"
	583: "xa-octopus",      //"Kirin"
	584: "xa-octopus",      //"Puma"
	585: "xa-octopus",      //"Boulder"
	586: "xa-octopus",      //"Banner"
	587: "xa-octopus",      //"Elven Ghost"
	588: "xa-octopus",      //"Human Ghost"
	589: "xa-octopus",      //"Chest"
	590: "xa-octopus",      //"Chest"
	591: "xa-octopus",      //"Crystal"
	592: "xa-octopus",      //"Coffin"
	593: "xa-octopus",      //"Guardian CPU"
	594: "xa-octopus",      //"Worg"
	595: "xa-octopus",      //"Mansion"
	596: "xa-octopus",      //"Floating Island"
	597: "xa-octopus",      //"Cragslither"
	598: "xa-octopus",      //"Wrulon"
	599: "xa-octopus",      //"Spell Particle 1"
	600: "xa-octopus",      //"Invisible Man of Zomm"
	601: "xa-octopus",      //"Robocopter of Zomm"
	602: "xa-octopus",      //"Burynai"
	603: "xa-octopus",      //"Frog"
	604: "xa-octopus",      //"Dracolich"
	605: "xa-octopus",      //"Iksar Ghost"
	606: "xa-octopus",      //"Iksar Skeleton"
	607: "xa-octopus",      //"Mephit"
	608: "xa-octopus",      //"Muddite"
	609: "xa-octopus",      //"Raptor"
	610: "xa-octopus",      //"Sarnak"
	611: "xa-octopus",      //"Scorpion"
	612: "xa-octopus",      //"Tsetsian"
	613: "xa-octopus",      //"Wurm"
	614: "xa-octopus",      //"Nekhon"
	615: "xa-octopus",      //"Hydra Crystal"
	616: "xa-octopus",      //"Crystal Sphere"
	617: "xa-octopus",      //"Gnoll"
	618: "xa-octopus",      //"Sokokar"
	619: "xa-octopus",      //"Stone Pylon"
	620: "xa-octopus",      //"Demon Vulture"
	621: "xa-octopus",      //"Wagon"
	622: "xa-octopus",      //"God of Discord"
	623: "xa-octopus",      //"Feran Mount"
	624: "xa-octopus",      //"Ogre NPC Male"
	625: "xa-octopus",      //"Sokokar Mount"
	626: "xa-octopus",      //"Giant"
	627: "xa-octopus",      //"Sokokar"
	628: "xa-octopus",      //"10th Anniversary Banner"
	629: "xa-octopus",      //"10th Anniversary Cake"
	630: "xa-octopus",      //"Wine Cask"
	631: "xa-octopus",      //"Hydra Mount"
	632: "xa-octopus",      //"Hydra NPC"
	633: "xa-octopus",      //"Wedding Flowers"
	634: "xa-octopus",      //"Wedding Arbor"
	635: "xa-octopus",      //"Wedding Altar"
	636: "xa-octopus",      //"Powder Keg"
	637: "xa-octopus",      //"Apexus"
	638: "xa-octopus",      //"Bellikos"
	639: "xa-octopus",      //"Brell's First Creation"
	640: "xa-octopus",      //"Brell"
	641: "xa-octopus",      //"Crystalskin Ambuloid"
	642: "xa-octopus",      //"Cliknar Queen"
	643: "xa-octopus",      //"Cliknar Soldier"
	644: "xa-octopus",      //"Cliknar Worker"
	645: "xa-octopus",      //"Coldain"
	646: "xa-octopus",      //"Coldain"
	647: "xa-octopus",      //"Crystalskin Sessiloid"
	648: "xa-octopus",      //"Genari"
	649: "xa-octopus",      //"Gigyn"
	650: "xa-octopus",      //"Greken Young Adult"
	651: "xa-octopus",      //"Greken Young"
	652: "xa-octopus",      //"Cliknar Mount"
	653: "xa-octopus",      //"Telmira"
	654: "xa-octopus",      //"Spider Mount"
	655: "xa-octopus",      //"Bear Mount"
	656: "xa-octopus",      //"Rat Mount Mystery Race"
	657: "xa-octopus",      //"Sessiloid Mount"
	658: "xa-octopus",      //"Morell Thule"
	659: "xa-octopus",      //"Marionette"
	660: "xa-octopus",      //"Book Dervish"
	661: "xa-octopus",      //"Topiary Lion"
	662: "xa-octopus",      //"Rotdog"
	663: "xa-octopus",      //"Amygdalan"
	664: "xa-octopus",      //"Sandman"
	665: "xa-octopus",      //"Grandfather Clock"
	666: "xa-octopus",      //"Gingerbread Man"
	667: "xa-octopus",      //"Royal Guard"
	668: "xa-octopus",      //"Rabbit"
	669: "xa-octopus",      //"Blind Dreamer"
	670: "xa-octopus",      //"Cazic Thule"
	671: "xa-octopus",      //"Topiary Lion Mount"
	672: "xa-octopus",      //"Rot Dog Mount"
	673: "xa-octopus",      //"Goral Mount"
	674: "xa-octopus",      //"Selyrah Mount"
	675: "xa-octopus",      //"Sclera Mount"
	676: "xa-octopus",      //"Braxi Mount"
	677: "xa-octopus",      //"Kangon Mount"
	678: "xa-octopus",      //"Erudite"
	679: "xa-octopus",      //"Wurm Mount"
	680: "xa-octopus",      //"Raptor Mount"
	681: "xa-octopus",      //"Invisible Man"
	682: "xa-octopus",      //"Whirligig"
	683: "xa-octopus",      //"Gnomish Balloon"
	684: "xa-octopus",      //"Gnomish Rocket Pack"
	685: "xa-octopus",      //"Gnomish Hovering Transport"
	686: "xa-octopus",      //"Selyrah"
	687: "xa-octopus",      //"Goral"
	688: "xa-octopus",      //"Braxi"
	689: "xa-octopus",      //"Kangon"
	690: "xa-octopus",      //"Invisible Man"
	691: "xa-octopus",      //"Floating Tower"
	692: "xa-octopus",      //"Explosive Cart"
	693: "xa-octopus",      //"Blimp Ship"
	694: "xa-octopus",      //"Tumbleweed"
	695: "xa-octopus",      //"Alaran"
	696: "xa-octopus",      //"Swinetor"
	697: "xa-octopus",      //"Triumvirate"
	698: "xa-octopus",      //"Hadal"
	699: "xa-octopus",      //"Hovering Platform"
	700: "xa-octopus",      //"Parasitic Scavenger"
	701: "xa-octopus",      //"Grendlaen"
	702: "xa-octopus",      //"Ship in a Bottle"
	703: "xa-octopus",      //"Alaran Sentry Stone"
	704: "xa-octopus",      //"Dervish"
	705: "xa-octopus",      //"Regeneration Pool"
	706: "xa-octopus",      //"Teleportation Stand"
	707: "xa-octopus",      //"Relic Case"
	708: "xa-octopus",      //"Alaran Ghost"
	709: "xa-octopus",      //"Skystrider"
	710: "xa-octopus",      //"Water Spout"
	711: "xa-octopus",      //"Aviak Pull Along"
	712: "xa-octopus",      //"Gelatinous Cube"
	713: "xa-octopus",      //"Cat"
	714: "xa-octopus",      //"Elk Head"
	715: "xa-octopus",      //"Holgresh"
	716: "xa-octopus",      //"Beetle"
	717: "xa-octopus",      //"Vine Maw"
	718: "xa-octopus",      //"Ratman"
	719: "xa-octopus",      //"Fallen Knight"
	720: "xa-octopus",      //"Flying Carpet"
	721: "xa-octopus",      //"Carrier Hand"
	722: "xa-octopus",      //"Akheva"
	723: "xa-octopus",      //"Servant of Shadow"
	724: "xa-octopus",      //"Luclin"
}
