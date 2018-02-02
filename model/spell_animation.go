package model

import ()

// Spell Animation represents animations on spells
// http://www.eqemulator.org/forums/showthread.php?t=30731
// swagger:model
type SpellAnimation struct {
	ID   int64
	Type string
	Name string
}

/*
var spellAnimationTypes = map[int64]*SpellAnimationType{
	-1: &SpellAnimationType{
		ID:   -1,
		Type: "-",
		Name: `None`,
	},
	0: &SpellAnimationType{
		ID:   0,
		Type: "-",
		Name: `None`,
	},
	2: &SpellAnimationType{
		ID:   2,
		Type: "D",
		Name: `Black smoke at hands and feet`,
	},
	3: &SpellAnimationType{
		ID:   3,
		Type: "B",
		Name: `Ice Dragon Breath`,
	},
	5: &SpellAnimationType{
		ID:   5,
		Type: "N",
		Name: `Stormcloud and falling rain`,
	},
	10: &SpellAnimationType{
		ID:   10,
		Type: "I",
		Name: `Yellow runes and sparkles`,
	},
	11: &SpellAnimationType{
		ID:   11,
		Type: "P",
		Name: `Cloud of bacterial spores`,
	},
	12: &SpellAnimationType{
		ID:   12,
		Type: "I",
		Name: `Spikes protrude from body`,
	},
	14: &SpellAnimationType{
		ID:   14,
		Type: "N",
		Name: `Delayed bubble of insects`,
	},
	15: &SpellAnimationType{
		ID:   15,
		Type: "S",
		Name: `Blue and white vertical pillars`,
	},
	16: &SpellAnimationType{
		ID:   16,
		Type: "D",
		Name: `Ghost faces rise from ground`,
	},
	17: &SpellAnimationType{
		ID:   17,
		Type: "F",
		Name: `Cascade of ice crystals`,
	},
	18: &SpellAnimationType{
		ID:   18,
		Type: "N",
		Name: `Molten metal drops penetrate`,
	},
	19: &SpellAnimationType{
		ID:   19,
		Type: "F",
		Name: `Cascade of snowballs`,
	},
	20: &SpellAnimationType{
		ID:   20,
		Type: "I",
		Name: `Shield and yellow sparkles`,
	},
	21: &SpellAnimationType{
		ID:   21,
		Type: "N",
		Name: `"Yellow sparkle tornados rise, leaves fall"`,
	},
	22: &SpellAnimationType{
		ID:   22,
		Type: "X",
		Name: `Abjuration Sound Effect 1`,
	},
	23: &SpellAnimationType{
		ID:   23,
		Type: "O",
		Name: `Green eyes hovering over`,
	},
	24: &SpellAnimationType{
		ID:   24,
		Type: "D",
		Name: `Disc of choking black spots`,
	},
	25: &SpellAnimationType{
		ID:   25,
		Type: "D",
		Name: `"Black smoke bubble, blue skulls"`,
	},
	26: &SpellAnimationType{
		ID:   26,
		Type: "D",
		Name: `"Black spots flow out in area, green skulls flow in"`,
	},
	27: &SpellAnimationType{
		ID:   27,
		Type: "X",
		Name: `Alteration Sound Effect 1`,
	},
	28: &SpellAnimationType{
		ID:   28,
		Type: "B",
		Name: `Fire dragon breath`,
	},
	29: &SpellAnimationType{
		ID:   29,
		Type: "F",
		Name: `Bonfire with black smoke`,
	},
	30: &SpellAnimationType{
		ID:   30,
		Type: "F",
		Name: `"Bonfire, ice storm, rocks shoot up"`,
	},
	32: &SpellAnimationType{
		ID:   32,
		Type: "X",
		Name: `Alteration Sound Effect 1`,
	},
	33: &SpellAnimationType{
		ID:   33,
		Type: "U",
		Name: `"Light blue pillar, blue face/hands from head"`,
	},
	34: &SpellAnimationType{
		ID:   34,
		Type: "O",
		Name: `One green eye hovers`,
	},
	35: &SpellAnimationType{
		ID:   35,
		Type: "B",
		Name: `Poison dragon breath`,
	},
	36: &SpellAnimationType{
		ID:   36,
		Type: "P",
		Name: `"Molten drops fall, poison cloud rises"`,
	},
	37: &SpellAnimationType{
		ID:   37,
		Type: "D",
		Name: `"Green fire and ghosts, with blue pillar on caster"`,
	},
	38: &SpellAnimationType{
		ID:   38,
		Type: "B",
		Name: `Delayed fire dragon breath with sound`,
	},
	39: &SpellAnimationType{
		ID:   39,
		Type: "N",
		Name: `"Leaves, rocks, and cloud of dirt"`,
	},
	40: &SpellAnimationType{
		ID:   40,
		Type: "D",
		Name: `"Pink/purple sparkles, red skull laughs"`,
	},
	41: &SpellAnimationType{
		ID:   41,
		Type: "F",
		Name: `Exploding ice crystal`,
	},
	42: &SpellAnimationType{
		ID:   42,
		Type: "P",
		Name: `Rising molten metal bubbles and green/black smoke`,
	},
	43: &SpellAnimationType{
		ID:   43,
		Type: "C",
		Name: `Big pink ball of light on caster`,
	},
	44: &SpellAnimationType{
		ID:   44,
		Type: "S",
		Name: `"Blue/pink sparkles, purple eye hovers"`,
	},
	45: &SpellAnimationType{
		ID:   45,
		Type: "M",
		Name: `Cyan musical notes rise`,
	},
	46: &SpellAnimationType{
		ID:   46,
		Type: "D",
		Name: `"Blue skulls around head, ghosts, ground green cloud"`,
	},
	47: &SpellAnimationType{
		ID:   47,
		Type: "I",
		Name: `"Armor, shield, pink sparkles"`,
	},
	48: &SpellAnimationType{
		ID:   48,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	49: &SpellAnimationType{
		ID:   49,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	50: &SpellAnimationType{
		ID:   50,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	51: &SpellAnimationType{
		ID:   51,
		Type: "S",
		Name: `"Pinks flow in, blues swirl, yellow pillar rises"`,
	},
	52: &SpellAnimationType{
		ID:   52,
		Type: "P",
		Name: `"Falling disease, rising poison, purple skulls"`,
	},
	53: &SpellAnimationType{
		ID:   53,
		Type: "D",
		Name: `Red/black smoke and sparkle bubble`,
	},
	54: &SpellAnimationType{
		ID:   54,
		Type: "C",
		Name: `Steam and yellow sparkle explosion at caster`,
	},
	55: &SpellAnimationType{
		ID:   55,
		Type: "D",
		Name: `"Pink cloud at feet, purple/green skulls"`,
	},
	58: &SpellAnimationType{
		ID:   58,
		Type: "D",
		Name: `"Purple fire, pink bubble, zombie head"`,
	},
	59: &SpellAnimationType{
		ID:   59,
		Type: "D",
		Name: `"Sparkle blast, red/green smoke"`,
	},
	60: &SpellAnimationType{
		ID:   60,
		Type: "S",
		Name: `"Blues/yellows rise, big pink ball"`,
	},
	61: &SpellAnimationType{
		ID:   61,
		Type: "D",
		Name: `"Ground aflame, red/black smoke/sparkles"`,
	},
	62: &SpellAnimationType{
		ID:   62,
		Type: "N",
		Name: `"Leafy tornado, blue sparkle explosion"`,
	},
	64: &SpellAnimationType{
		ID:   64,
		Type: "D",
		Name: `"Red/black fire/smoke, red arches flow in"`,
	},
	65: &SpellAnimationType{
		ID:   65,
		Type: "O",
		Name: `"Steam cloud, werewolf head"`,
	},
	67: &SpellAnimationType{
		ID:   67,
		Type: "B",
		Name: `Ice dragon breath from target`,
	},
	68: &SpellAnimationType{
		ID:   68,
		Type: "F",
		Name: `Huge pillar of flame`,
	},
	69: &SpellAnimationType{
		ID:   69,
		Type: "S",
		Name: `Subtle blue/pink swirling (15s)`,
	},
	71: &SpellAnimationType{
		ID:   71,
		Type: "F",
		Name: `Crashing ice meteor`,
	},
	72: &SpellAnimationType{
		ID:   72,
		Type: "L",
		Name: `Choking fiery flares`,
	},
	73: &SpellAnimationType{
		ID:   73,
		Type: "U",
		Name: `Blue ground and pillar`,
	},
	74: &SpellAnimationType{
		ID:   74,
		Type: "S",
		Name: `Golden sparkles from eyes`,
	},
	75: &SpellAnimationType{
		ID:   75,
		Type: "D",
		Name: `Large green skulls`,
	},
	76: &SpellAnimationType{
		ID:   76,
		Type: "D",
		Name: `Red static field flows in`,
	},
	77: &SpellAnimationType{
		ID:   77,
		Type: "X",
		Name: `Evocation Sound Effect 2`,
	},
	78: &SpellAnimationType{
		ID:   78,
		Type: "T",
		Name: `Blue/cyan portal`,
	},
	79: &SpellAnimationType{
		ID:   79,
		Type: "X",
		Name: `Conjuration Sound Effect 1`,
	},
	80: &SpellAnimationType{
		ID:   80,
		Type: "F",
		Name: `Embers explode outward`,
	},
	81: &SpellAnimationType{
		ID:   81,
		Type: "C",
		Name: `Golden flash at caster`,
	},
	82: &SpellAnimationType{
		ID:   82,
		Type: "E",
		Name: `Electric bubble`,
	},
	83: &SpellAnimationType{
		ID:   83,
		Type: "X",
		Name: `Conjuration Sound Effect 1`,
	},
	84: &SpellAnimationType{
		ID:   84,
		Type: "F",
		Name: `"Big fire sparkle, flames flicker at feet"`,
	},
	85: &SpellAnimationType{
		ID:   85,
		Type: "F",
		Name: `"Pillar of fire sparkles, flames at feet"`,
	},
	86: &SpellAnimationType{
		ID:   86,
		Type: "S",
		Name: `Blue-green sparkles flow inward`,
	},
	87: &SpellAnimationType{
		ID:   87,
		Type: "L",
		Name: `Purple light bubble cascades`,
	},
	88: &SpellAnimationType{
		ID:   88,
		Type: "S",
		Name: `"Blues flow in, purples flow out"`,
	},
	89: &SpellAnimationType{
		ID:   89,
		Type: "P",
		Name: `Delayed bubble and cone of disease`,
	},
	90: &SpellAnimationType{
		ID:   90,
		Type: "F",
		Name: `Flame aurora cascades`,
	},
	91: &SpellAnimationType{
		ID:   91,
		Type: "L",
		Name: `Purple ring of light at ground and throat`,
	},
	92: &SpellAnimationType{
		ID:   92,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	93: &SpellAnimationType{
		ID:   93,
		Type: "G",
		Name: `Cyan sparkles flow into ground`,
	},
	94: &SpellAnimationType{
		ID:   94,
		Type: "F",
		Name: `Embers surround caster and target`,
	},
	95: &SpellAnimationType{
		ID:   95,
		Type: "N",
		Name: `Blue/white bubbles flow inward`,
	},
	96: &SpellAnimationType{
		ID:   96,
		Type: "P",
		Name: `Disease cloud permeates ground`,
	},
	98: &SpellAnimationType{
		ID:   98,
		Type: "L",
		Name: `Small purple ball at waist`,
	},
	100: &SpellAnimationType{
		ID:   100,
		Type: "L",
		Name: `"Large flare, yellow flares from ground"`,
	},
	101: &SpellAnimationType{
		ID:   101,
		Type: "D",
		Name: `Green skulls permeate ground`,
	},
	102: &SpellAnimationType{
		ID:   102,
		Type: "L",
		Name: `"Purple ring, purples around the ground"`,
	},
	103: &SpellAnimationType{
		ID:   103,
		Type: "L",
		Name: `"Red-orange arch, central burst"`,
	},
	104: &SpellAnimationType{
		ID:   104,
		Type: "U",
		Name: `Blues blow around caster and target`,
	},
	105: &SpellAnimationType{
		ID:   105,
		Type: "P",
		Name: `Green acid bursts`,
	},
	106: &SpellAnimationType{
		ID:   106,
		Type: "U",
		Name: `"Blue smoke pillar, star sparklies"`,
	},
	107: &SpellAnimationType{
		ID:   107,
		Type: "U",
		Name: `"Blue bubble, smoke emanation"`,
	},
	108: &SpellAnimationType{
		ID:   108,
		Type: "P",
		Name: `Green sparkles flow into feet`,
	},
	109: &SpellAnimationType{
		ID:   109,
		Type: "P",
		Name: `Green ripples flow outward`,
	},
	110: &SpellAnimationType{
		ID:   110,
		Type: "S",
		Name: `"Purple ball, blue sparkle bubble"`,
	},
	111: &SpellAnimationType{
		ID:   111,
		Type: "N",
		Name: `Cyan air puffs flow inward and on ground`,
	},
	112: &SpellAnimationType{
		ID:   112,
		Type: "L",
		Name: `Flickering balls of light at waist level`,
	},
	113: &SpellAnimationType{
		ID:   113,
		Type: "U",
		Name: `Delayed rising blues in area`,
	},
	114: &SpellAnimationType{
		ID:   114,
		Type: "L",
		Name: `Purple ball at target and ground at caster`,
	},
	115: &SpellAnimationType{
		ID:   115,
		Type: "S",
		Name: `Red sparkle bubble`,
	},
	116: &SpellAnimationType{
		ID:   116,
		Type: "P",
		Name: `Disease bubble`,
	},
	117: &SpellAnimationType{
		ID:   117,
		Type: "F",
		Name: `"Blinding white/red flare-up, fireball"`,
	},
	118: &SpellAnimationType{
		ID:   118,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	119: &SpellAnimationType{
		ID:   119,
		Type: "L",
		Name: `"Green arch, blue pillar, white lights"`,
	},
	120: &SpellAnimationType{
		ID:   120,
		Type: "F",
		Name: `Embers completely fill area`,
	},
	121: &SpellAnimationType{
		ID:   121,
		Type: "L",
		Name: `"Blue/purple steam, white lights surround head"`,
	},
	122: &SpellAnimationType{
		ID:   122,
		Type: "F",
		Name: `Fire consumes torso`,
	},
	123: &SpellAnimationType{
		ID:   123,
		Type: "W",
		Name: `Tenth Anniversary Fireworks`,
	},
	124: &SpellAnimationType{
		ID:   124,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	125: &SpellAnimationType{
		ID:   125,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	126: &SpellAnimationType{
		ID:   126,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	127: &SpellAnimationType{
		ID:   127,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	128: &SpellAnimationType{
		ID:   128,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	129: &SpellAnimationType{
		ID:   129,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	130: &SpellAnimationType{
		ID:   130,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	131: &SpellAnimationType{
		ID:   131,
		Type: "X",
		Name: `Evocation Sound Effect 1`,
	},
	132: &SpellAnimationType{
		ID:   132,
		Type: "S",
		Name: `Blue/green pillars rise and fall`,
	},
	133: &SpellAnimationType{
		ID:   133,
		Type: "H",
		Name: `Black smoke emanates from hands`,
	},
	134: &SpellAnimationType{
		ID:   134,
		Type: "U",
		Name: `Blue/green pillar and long blue light ring (32s)`,
	},
	135: &SpellAnimationType{
		ID:   135,
		Type: "U",
		Name: `"Blue light, cyan cloud/puffs"`,
	},
	136: &SpellAnimationType{
		ID:   136,
		Type: "U",
		Name: `Delayed blue light at head potion`,
	},
	137: &SpellAnimationType{
		ID:   137,
		Type: "U",
		Name: `Blue glow flows from target's head to caster's`,
	},
	139: &SpellAnimationType{
		ID:   139,
		Type: "M",
		Name: `"Blue/green musical notes, ghosts, black smoke"`,
	},
	140: &SpellAnimationType{
		ID:   140,
		Type: "M",
		Name: `"Cyan musical notes, disease, virus"`,
	},
	141: &SpellAnimationType{
		ID:   141,
		Type: "U",
		Name: `"Huge blue pillar, blue/yellow sparkles (21s)"`,
	},
	142: &SpellAnimationType{
		ID:   142,
		Type: "U",
		Name: `"Blue pillar waterfall, white steam, blue sparkles"`,
	},
	144: &SpellAnimationType{
		ID:   144,
		Type: "D",
		Name: `Red/black smoke blows into face`,
	},
	145: &SpellAnimationType{
		ID:   145,
		Type: "N",
		Name: `Black/yellow insect bubble rises`,
	},
	146: &SpellAnimationType{
		ID:   146,
		Type: "D",
		Name: `Red aura and swirling sparkles (30s)`,
	},
	147: &SpellAnimationType{
		ID:   147,
		Type: "F",
		Name: `"Flames at head and feet, rising steam/smoke"`,
	},
	148: &SpellAnimationType{
		ID:   148,
		Type: "F",
		Name: `Ice bubble explosion and floating crystals`,
	},
	149: &SpellAnimationType{
		ID:   149,
		Type: "N",
		Name: `Dirt and smoke choke and rise from ground`,
	},
	150: &SpellAnimationType{
		ID:   150,
		Type: "F",
		Name: `Blue/yellow bonfire rises`,
	},
	151: &SpellAnimationType{
		ID:   151,
		Type: "D",
		Name: `"Smoke at feet, rising skulls"`,
	},
	152: &SpellAnimationType{
		ID:   152,
		Type: "D",
		Name: `"Red/black smoke, ember bubble around head"`,
	},
	153: &SpellAnimationType{
		ID:   153,
		Type: "F",
		Name: `Falling ice crystal shatters and explodes`,
	},
	154: &SpellAnimationType{
		ID:   154,
		Type: "H",
		Name: `Hands glow with flame`,
	},
	155: &SpellAnimationType{
		ID:   155,
		Type: "U",
		Name: `"Blue glow, blue sparkle bubble"`,
	},
	156: &SpellAnimationType{
		ID:   156,
		Type: "X",
		Name: `Conjuration Sound Effect 2`,
	},
	157: &SpellAnimationType{
		ID:   157,
		Type: "L",
		Name: `Yellow stars pour down`,
	},
	158: &SpellAnimationType{
		ID:   158,
		Type: "N",
		Name: `"Yellow aura, then butterflies and sparkles"`,
	},
	159: &SpellAnimationType{
		ID:   159,
		Type: "L",
		Name: `"Green/Purple aura, red steam swirls"`,
	},
	160: &SpellAnimationType{
		ID:   160,
		Type: "E",
		Name: `Lightning bolts assault torso`,
	},
	161: &SpellAnimationType{
		ID:   161,
		Type: "G",
		Name: `Bones and flames at feet`,
	},
	162: &SpellAnimationType{
		ID:   162,
		Type: "F",
		Name: `Pillar of falling snowflakes`,
	},
	163: &SpellAnimationType{
		ID:   163,
		Type: "E",
		Name: `Long lightning storm (12s)`,
	},
	164: &SpellAnimationType{
		ID:   164,
		Type: "E",
		Name: `Focused neural lightning storm`,
	},
	165: &SpellAnimationType{
		ID:   165,
		Type: "P",
		Name: `"Pink/yellow lights around head, yellow acid, black bubbles"`,
	},
	166: &SpellAnimationType{
		ID:   166,
		Type: "D",
		Name: `Bone ring around torso`,
	},
	167: &SpellAnimationType{
		ID:   167,
		Type: "W",
		Name: `Eleventh Anniversary Fireworks`,
	},
	168: &SpellAnimationType{
		ID:   168,
		Type: "O",
		Name: `Mistmoore's face and teleportation sound effect`,
	},
	169: &SpellAnimationType{
		ID:   169,
		Type: "N",
		Name: `"Rock tornado, pink light, purple smoke"`,
	},
	170: &SpellAnimationType{
		ID:   170,
		Type: "G",
		Name: `Pink/green smoke and black insects at feet`,
	},
	171: &SpellAnimationType{
		ID:   171,
		Type: "T",
		Name: `Dark portal`,
	},
	172: &SpellAnimationType{
		ID:   172,
		Type: "G",
		Name: `"Green smoke, black insects on ground"`,
	},
	173: &SpellAnimationType{
		ID:   173,
		Type: "O",
		Name: `"Orc head appears, blows bad breath"`,
	},
	174: &SpellAnimationType{
		ID:   174,
		Type: "N",
		Name: `Falling summoned large stones`,
	},
	175: &SpellAnimationType{
		ID:   175,
		Type: "N",
		Name: `Falling summoned small rocks`,
	},
	176: &SpellAnimationType{
		ID:   176,
		Type: "D",
		Name: `Ghostly chains prevent dark portal`,
	},
	177: &SpellAnimationType{
		ID:   177,
		Type: "F",
		Name: `Fire blasts and spinning swords`,
	},
	178: &SpellAnimationType{
		ID:   178,
		Type: "S",
		Name: `"Blue/pink ground, rising yellow-green pillar"`,
	},
	179: &SpellAnimationType{
		ID:   179,
		Type: "F",
		Name: `"Cold torso, snowflake ring"`,
	},
	180: &SpellAnimationType{
		ID:   180,
		Type: "F",
		Name: `"Wall of ice, pillar of snowflakes (15s)"`,
	},
	181: &SpellAnimationType{
		ID:   181,
		Type: "U",
		Name: `"Blue bubble, spinning ghostly swords"`,
	},
	182: &SpellAnimationType{
		ID:   182,
		Type: "D",
		Name: `"Red/black smoke, grasping sickly claws"`,
	},
	183: &SpellAnimationType{
		ID:   183,
		Type: "D",
		Name: `"Black smoke, red pillar, grasping black claws"`,
	},
	184: &SpellAnimationType{
		ID:   184,
		Type: "O",
		Name: `"Laughing orc faces, yellow-green aura, red smoke"`,
	},
	185: &SpellAnimationType{
		ID:   185,
		Type: "D",
		Name: `Red pumping heart/veins and blood cloud`,
	},
	186: &SpellAnimationType{
		ID:   186,
		Type: "F",
		Name: `Flames rise and light up steam/smoke orange`,
	},
	187: &SpellAnimationType{
		ID:   187,
		Type: "N",
		Name: `"Black wings/smoke materialize, blue lights from waist"`,
	},
	188: &SpellAnimationType{
		ID:   188,
		Type: "D",
		Name: `"Red pumping heart/veins, dark blood cloud"`,
	},
	189: &SpellAnimationType{
		ID:   189,
		Type: "D",
		Name: `"Black pumping heart/veins, dark blood cloud"`,
	},
	190: &SpellAnimationType{
		ID:   190,
		Type: "D",
		Name: `"Bright red pumping heart/veins, blood cloud"`,
	},
	191: &SpellAnimationType{
		ID:   191,
		Type: "P",
		Name: `"Dark green cloud on ground, black snakes attack"`,
	},
	192: &SpellAnimationType{
		ID:   192,
		Type: "P",
		Name: `"Green cloud on ground, green snakes attack"`,
	},
	193: &SpellAnimationType{
		ID:   193,
		Type: "N",
		Name: `"Black insects, black smoke"`,
	},
	194: &SpellAnimationType{
		ID:   194,
		Type: "P",
		Name: `"Black insects, poisonous cloud"`,
	},
	195: &SpellAnimationType{
		ID:   195,
		Type: "F",
		Name: `Solar flare storm`,
	},
	196: &SpellAnimationType{
		ID:   196,
		Type: "F",
		Name: `Solar flare explosion`,
	},
	197: &SpellAnimationType{
		ID:   197,
		Type: "F",
		Name: `Blue/yellow flames rise`,
	},
	198: &SpellAnimationType{
		ID:   198,
		Type: "F",
		Name: `"Flare out, swirling flare embers"`,
	},
	199: &SpellAnimationType{
		ID:   199,
		Type: "P",
		Name: `"Acid drops, acid steam, poison cloud"`,
	},
	200: &SpellAnimationType{
		ID:   200,
		Type: "P",
		Name: `Poison clouds and debris`,
	},
	201: &SpellAnimationType{
		ID:   201,
		Type: "D",
		Name: `Rising cone of red sparks and black smoke`,
	},
	202: &SpellAnimationType{
		ID:   202,
		Type: "E",
		Name: `Large lightning bolt strikes`,
	},
	203: &SpellAnimationType{
		ID:   203,
		Type: "U",
		Name: `"Blue glow and aura, white sparkles (15s)"`,
	},
	204: &SpellAnimationType{
		ID:   204,
		Type: "F",
		Name: `"Flames/smoke flow out from caster, consume target"`,
	},
	205: &SpellAnimationType{
		ID:   205,
		Type: "F",
		Name: `"Snowflakes flow out from caster, surround target"`,
	},
	206: &SpellAnimationType{
		ID:   206,
		Type: "E",
		Name: `Short focused lightning storm`,
	},
	207: &SpellAnimationType{
		ID:   207,
		Type: "P",
		Name: `"Acid sputters and falls, green smoke rises"`,
	},
	208: &SpellAnimationType{
		ID:   208,
		Type: "U",
		Name: `Rising balls of blue light`,
	},
	209: &SpellAnimationType{
		ID:   209,
		Type: "U",
		Name: `"Balls of blue light, ribbon, and sparkles"`,
	},
	210: &SpellAnimationType{
		ID:   210,
		Type: "G",
		Name: `Purple/cyan flames rise from feet`,
	},
	211: &SpellAnimationType{
		ID:   211,
		Type: "H",
		Name: `Blue-green sparkles emanate from hands`,
	},
	212: &SpellAnimationType{
		ID:   212,
		Type: "L",
		Name: `"Blue light rings, purple bubble"`,
	},
	213: &SpellAnimationType{
		ID:   213,
		Type: "H",
		Name: `Acid/smoke spurts from hands`,
	},
	214: &SpellAnimationType{
		ID:   214,
		Type: "N",
		Name: `Brown roots rise from ground and grab hold`,
	},
	215: &SpellAnimationType{
		ID:   215,
		Type: "N",
		Name: `Green-orange roots rise from ground and grab hold`,
	},
	216: &SpellAnimationType{
		ID:   216,
		Type: "N",
		Name: `"Pink/green sparkles, rising pillar of leaves"`,
	},
	217: &SpellAnimationType{
		ID:   217,
		Type: "N",
		Name: `"Pink/green butterflies, green/yellow ribbon"`,
	},
	218: &SpellAnimationType{
		ID:   218,
		Type: "T",
		Name: `Yellow rippling teleport`,
	},
	219: &SpellAnimationType{
		ID:   219,
		Type: "E",
		Name: `Electric field covers torso`,
	},
	220: &SpellAnimationType{
		ID:   220,
		Type: "L",
		Name: `"Yellow-green aura, orange-pink ribbons"`,
	},
	221: &SpellAnimationType{
		ID:   221,
		Type: "O",
		Name: `"World and moons hover, purple rains down"`,
	},
	222: &SpellAnimationType{
		ID:   222,
		Type: "S",
		Name: `"Green sparkles, red field"`,
	},
	223: &SpellAnimationType{
		ID:   223,
		Type: "M",
		Name: `"White sparkles, black notes"`,
	},
	224: &SpellAnimationType{
		ID:   224,
		Type: "D",
		Name: `"Red/black smoke cone, rising skulls"`,
	},
	225: &SpellAnimationType{
		ID:   225,
		Type: "L",
		Name: `Rising purple light/sparkles`,
	},
	226: &SpellAnimationType{
		ID:   226,
		Type: "D",
		Name: `"Blood cloud, black cloud, black insects"`,
	},
	227: &SpellAnimationType{
		ID:   227,
		Type: "T",
		Name: `Green field teleport`,
	},
	228: &SpellAnimationType{
		ID:   228,
		Type: "L",
		Name: `"Purple aura around ground and head, like porting in"`,
	},
	229: &SpellAnimationType{
		ID:   229,
		Type: "U",
		Name: `Blue pillar of auras and sparkles`,
	},
	230: &SpellAnimationType{
		ID:   230,
		Type: "G",
		Name: `Ember glow on ground`,
	},
	231: &SpellAnimationType{
		ID:   231,
		Type: "F",
		Name: `Lava erupts from ground beneath`,
	},
	232: &SpellAnimationType{
		ID:   232,
		Type: "I",
		Name: `Spike coat and bubble`,
	},
	233: &SpellAnimationType{
		ID:   233,
		Type: "L",
		Name: `Rising purple steam pillar`,
	},
	234: &SpellAnimationType{
		ID:   234,
		Type: "N",
		Name: `Green roots rise from ground and grab hold`,
	},
	235: &SpellAnimationType{
		ID:   235,
		Type: "L",
		Name: `White/purple starry glows flow outward`,
	},
	236: &SpellAnimationType{
		ID:   236,
		Type: "S",
		Name: `Cyan sparkles around throat and forms a pillar`,
	},
	237: &SpellAnimationType{
		ID:   237,
		Type: "D",
		Name: `Rising black/grey smoke and skulls`,
	},
	238: &SpellAnimationType{
		ID:   238,
		Type: "S",
		Name: `Green bubble and aura on ground`,
	},
	239: &SpellAnimationType{
		ID:   239,
		Type: "L",
		Name: `"Green aura, ember ring, ember bubble"`,
	},
	240: &SpellAnimationType{
		ID:   240,
		Type: "H",
		Name: `"Subtle red bubble and green ring, purple from hands"`,
	},
	241: &SpellAnimationType{
		ID:   241,
		Type: "F",
		Name: `"Flames inside, ice outside"`,
	},
	242: &SpellAnimationType{
		ID:   242,
		Type: "F",
		Name: `"Ice inside, flames outside"`,
	},
	243: &SpellAnimationType{
		ID:   243,
		Type: "F",
		Name: `Flames emanating from waist`,
	},
	244: &SpellAnimationType{
		ID:   244,
		Type: "D",
		Name: `"Black clouds, red glow, rising pillar of bones"`,
	},
	245: &SpellAnimationType{
		ID:   245,
		Type: "D",
		Name: `"Black ribbon, grey smoke"`,
	},
	246: &SpellAnimationType{
		ID:   246,
		Type: "D",
		Name: `Pillars of black smoke`,
	},
	247: &SpellAnimationType{
		ID:   247,
		Type: "P",
		Name: `"Acid smoke, darts fly all over area"`,
	},
	248: &SpellAnimationType{
		ID:   248,
		Type: "P",
		Name: `"Acid spurts up, sprays onto ground"`,
	},
	249: &SpellAnimationType{
		ID:   249,
		Type: "D",
		Name: `"Red glow and smoke, bones"`,
	},
	250: &SpellAnimationType{
		ID:   250,
		Type: "U",
		Name: `"Blue ribbon, glow at waist, sparkles"`,
	},
	251: &SpellAnimationType{
		ID:   251,
		Type: "U",
		Name: `"Blue glow rises, sparkles fall"`,
	},
	252: &SpellAnimationType{
		ID:   252,
		Type: "T",
		Name: `Orange-red flame teleport`,
	},
	253: &SpellAnimationType{
		ID:   253,
		Type: "L",
		Name: `"Blue aura/lights, black smoke (Mana Burn?)"`,
	},
	254: &SpellAnimationType{
		ID:   254,
		Type: "L",
		Name: `Ribbon of rising stars`,
	},
	255: &SpellAnimationType{
		ID:   255,
		Type: "L",
		Name: `"Subtle ring, large glowing white star"`,
	},
	256: &SpellAnimationType{
		ID:   256,
		Type: "B",
		Name: `"Blue ring/sparkles around head, then white bubble (13s)"`,
	},
	257: &SpellAnimationType{
		ID:   257,
		Type: "H",
		Name: `"Hands and arms aflame, then black smoke"`,
	},
	258: &SpellAnimationType{
		ID:   258,
		Type: "L",
		Name: `"Blue glow, green stars and ribbons"`,
	},
	259: &SpellAnimationType{
		ID:   259,
		Type: "L",
		Name: `Orange flare-bubbles emanate`,
	},
	260: &SpellAnimationType{
		ID:   260,
		Type: "F",
		Name: `Spouts of flame swirl and surround`,
	},
	261: &SpellAnimationType{
		ID:   261,
		Type: "F",
		Name: `"Ice formations permeate, snowflakes rise"`,
	},
	262: &SpellAnimationType{
		ID:   262,
		Type: "F",
		Name: `"Ice formations on ground, rising cyan sparkles"`,
	},
	263: &SpellAnimationType{
		ID:   263,
		Type: "P",
		Name: `Acid smoke rains down`,
	},
	264: &SpellAnimationType{
		ID:   264,
		Type: "T",
		Name: `"Blue bubbles permeate, star glow envelopes"`,
	},
	265: &SpellAnimationType{
		ID:   265,
		Type: "L",
		Name: `"Star falls and hits head, shower of stars and sparkles shoots out"`,
	},
	266: &SpellAnimationType{
		ID:   266,
		Type: "L",
		Name: `"Blue field, sparkles permeate, green sparkles leave"`,
	},
	267: &SpellAnimationType{
		ID:   267,
		Type: "H",
		Name: `"Flame aura at hands, black smoke at feet"`,
	},
	268: &SpellAnimationType{
		ID:   268,
		Type: "D",
		Name: `"Red smoke on ground, black smoke on body"`,
	},
	269: &SpellAnimationType{
		ID:   269,
		Type: "D",
		Name: `"Black smoke falls, green-orange-red smoke rises"`,
	},
	270: &SpellAnimationType{
		ID:   270,
		Type: "H",
		Name: `Green-purple viruses swarm from hands`,
	},
	271: &SpellAnimationType{
		ID:   271,
		Type: "P",
		Name: `"Orange flares and viruses, black smoke swirls"`,
	},
	272: &SpellAnimationType{
		ID:   272,
		Type: "D",
		Name: `Ring of black smoke rises`,
	},
	273: &SpellAnimationType{
		ID:   273,
		Type: "S",
		Name: `"Pink lights, cyan stars emanate, long sound (9s)"`,
	},
	274: &SpellAnimationType{
		ID:   274,
		Type: "F",
		Name: `"Cold envelopes, stars shine"`,
	},
	275: &SpellAnimationType{
		ID:   275,
		Type: "F",
		Name: `Fast fire envelopes and rages`,
	},
	276: &SpellAnimationType{
		ID:   276,
		Type: "F",
		Name: `Bonfire with black smoke`,
	},
	277: &SpellAnimationType{
		ID:   277,
		Type: "C",
		Name: `"Blue lights flow out, then pink/white glow"`,
	},
	278: &SpellAnimationType{
		ID:   278,
		Type: "U",
		Name: `Rising blue stars and aura`,
	},
	279: &SpellAnimationType{
		ID:   279,
		Type: "F",
		Name: `Three thin pillars of flame`,
	},
	280: &SpellAnimationType{
		ID:   280,
		Type: "F",
		Name: `Dozens of fireballs land and burn`,
	},
	281: &SpellAnimationType{
		ID:   281,
		Type: "L",
		Name: `"Pink glow falls, white-blue stars spiral up"`,
	},
	283: &SpellAnimationType{
		ID:   283,
		Type: "F",
		Name: `"Ice formations flow in, envelope, and explode"`,
	},
	284: &SpellAnimationType{
		ID:   284,
		Type: "F",
		Name: `Ice crystals rain down from above`,
	},
	285: &SpellAnimationType{
		ID:   285,
		Type: "C",
		Name: `Triple green star glow flows out from caster`,
	},
	286: &SpellAnimationType{
		ID:   286,
		Type: "L",
		Name: `Blue-green glow circles ground and head`,
	},
	287: &SpellAnimationType{
		ID:   287,
		Type: "S",
		Name: `"Yellow-green sparkles envelope, then pillar rises"`,
	},
	288: &SpellAnimationType{
		ID:   288,
		Type: "U",
		Name: `"White-blue ribbon, then blue flows out along ground"`,
	},
	289: &SpellAnimationType{
		ID:   289,
		Type: "U",
		Name: `Blue ground and full bubble`,
	},
	290: &SpellAnimationType{
		ID:   290,
		Type: "D",
		Name: `Red star and rising swirling bones`,
	},
	291: &SpellAnimationType{
		ID:   291,
		Type: "L",
		Name: `"White starglow around neck, ribbon above head"`,
	},
	292: &SpellAnimationType{
		ID:   292,
		Type: "L",
		Name: `White-blue ribbon`,
	},
	293: &SpellAnimationType{
		ID:   293,
		Type: "P",
		Name: `Acid clouds fall and rise`,
	},
	294: &SpellAnimationType{
		ID:   294,
		Type: "P",
		Name: `"Poison clouds rise, skulls fall"`,
	},
	295: &SpellAnimationType{
		ID:   295,
		Type: "L",
		Name: `"Starglow and emanating blue, then black smoke"`,
	},
	296: &SpellAnimationType{
		ID:   296,
		Type: "B",
		Name: `"Blue lights flow in, black smoke envelopes"`,
	},
	297: &SpellAnimationType{
		ID:   297,
		Type: "L",
		Name: `Blue bubble fades up to green`,
	},
	298: &SpellAnimationType{
		ID:   298,
		Type: "L",
		Name: `"Ring of blue-white stars circles, then white bubble"`,
	},
	299: &SpellAnimationType{
		ID:   299,
		Type: "O",
		Name: `"Rune tablet appears, then shatters with blue sparkles"`,
	},
	300: &SpellAnimationType{
		ID:   300,
		Type: "H",
		Name: `"Disease cloud emanates, black smoke rises from hands"`,
	},
	301: &SpellAnimationType{
		ID:   301,
		Type: "H",
		Name: `Hands and feet glow blue-green`,
	},
	302: &SpellAnimationType{
		ID:   302,
		Type: "L",
		Name: `Blue aura gets consumed by fire`,
	},
	303: &SpellAnimationType{
		ID:   303,
		Type: "O",
		Name: `"Green ribbon, jumping man, sparkles"`,
	},
	304: &SpellAnimationType{
		ID:   304,
		Type: "O",
		Name: `"Smoke, embers, blue glow, ox skull"`,
	},
	305: &SpellAnimationType{
		ID:   305,
		Type: "X",
		Name: `Abjuration Sound Effect 1`,
	},
	307: &SpellAnimationType{
		ID:   307,
		Type: "O",
		Name: `Silhouette of floating man in burning sun`,
	},
	309: &SpellAnimationType{
		ID:   309,
		Type: "B",
		Name: `Poison dragon breath that envelopes targets`,
	},
	310: &SpellAnimationType{
		ID:   310,
		Type: "N",
		Name: `"Cone of smoke, then green pillar, then falling leaves"`,
	},
	311: &SpellAnimationType{
		ID:   311,
		Type: "N",
		Name: `Cone of insects and black dust rises`,
	},
	312: &SpellAnimationType{
		ID:   312,
		Type: "L",
		Name: `"Starglow turns to sun, ghost heads rise"`,
	},
	313: &SpellAnimationType{
		ID:   313,
		Type: "I",
		Name: `"Golden starglow, breastplate, sparkles"`,
	},
	314: &SpellAnimationType{
		ID:   314,
		Type: "D",
		Name: `"Red lights flow in, black smoke rises"`,
	},
	315: &SpellAnimationType{
		ID:   315,
		Type: "F",
		Name: `Balls of flame surround`,
	},
	316: &SpellAnimationType{
		ID:   316,
		Type: "F",
		Name: `"Flame envelopes, pillars, and blows out with bubble"`,
	},
	317: &SpellAnimationType{
		ID:   317,
		Type: "L",
		Name: `Swirling orange aura ring and bubble`,
	},
	318: &SpellAnimationType{
		ID:   318,
		Type: "U",
		Name: `Swirling blue aura ring and bubble`,
	},
	319: &SpellAnimationType{
		ID:   319,
		Type: "C",
		Name: `Silent yellow teleport out for caster`,
	},
	320: &SpellAnimationType{
		ID:   320,
		Type: "U",
		Name: `"Blue glow, ribbon, and sparkles rise"`,
	},
	321: &SpellAnimationType{
		ID:   321,
		Type: "H",
		Name: `Pink glow on hands`,
	},
	322: &SpellAnimationType{
		ID:   322,
		Type: "L",
		Name: `Orange smoke lights rise`,
	},
	323: &SpellAnimationType{
		ID:   323,
		Type: "L",
		Name: `Pink-blue smoke lights rise`,
	},
	324: &SpellAnimationType{
		ID:   324,
		Type: "M",
		Name: `Blue notes and sparkles around caster's and target's heads`,
	},
	332: &SpellAnimationType{
		ID:   332,
		Type: "M",
		Name: `Red notes swirling`,
	},
	333: &SpellAnimationType{
		ID:   333,
		Type: "M",
		Name: `White notes swirling`,
	},
	334: &SpellAnimationType{
		ID:   334,
		Type: "M",
		Name: `Blue notes swirling`,
	},
	335: &SpellAnimationType{
		ID:   335,
		Type: "M",
		Name: `Green notes and skulls swirling`,
	},
	336: &SpellAnimationType{
		ID:   336,
		Type: "M",
		Name: `"Red and white notes, blue stars"`,
	},
	337: &SpellAnimationType{
		ID:   337,
		Type: "M",
		Name: `Red/green notes and white skulls rise`,
	},
	338: &SpellAnimationType{
		ID:   338,
		Type: "M",
		Name: `Red notes rising`,
	},
	339: &SpellAnimationType{
		ID:   339,
		Type: "M",
		Name: `White notes rising`,
	},
	340: &SpellAnimationType{
		ID:   340,
		Type: "M",
		Name: `Blue notes rising`,
	},
	341: &SpellAnimationType{
		ID:   341,
		Type: "M",
		Name: `Green notes and skulls rising`,
	},
	342: &SpellAnimationType{
		ID:   342,
		Type: "M",
		Name: `Red/White/Blue notes with blue sparkles`,
	},
	343: &SpellAnimationType{
		ID:   343,
		Type: "S",
		Name: `Pink and green sparkles surround head`,
	},
	344: &SpellAnimationType{
		ID:   344,
		Type: "M",
		Name: `"Black notes, white skulls"`,
	},
	345: &SpellAnimationType{
		ID:   345,
		Type: "M",
		Name: `Blue swirling notes`,
	},
	346: &SpellAnimationType{
		ID:   346,
		Type: "-",
		Name: `None`,
	},
	347: &SpellAnimationType{
		ID:   347,
		Type: "M",
		Name: `Green notes swirl around feet`,
	},
	348: &SpellAnimationType{
		ID:   348,
		Type: "M",
		Name: `Blue notes surround feet`,
	},
	349: &SpellAnimationType{
		ID:   349,
		Type: "D",
		Name: `Black smoky haze`,
	},
	350: &SpellAnimationType{
		ID:   350,
		Type: "U",
		Name: `Blue smoky haze`,
	},
	351: &SpellAnimationType{
		ID:   351,
		Type: "N",
		Name: `Brown smoky haze`,
	},
	352: &SpellAnimationType{
		ID:   352,
		Type: "P",
		Name: `Dark green smoky haze`,
	},
	353: &SpellAnimationType{
		ID:   353,
		Type: "L",
		Name: `Light green smoky haze`,
	},
	354: &SpellAnimationType{
		ID:   354,
		Type: "D",
		Name: `Black smoky haze with white skulls`,
	},
	355: &SpellAnimationType{
		ID:   355,
		Type: "L",
		Name: `Bright red smoky haze`,
	},
	356: &SpellAnimationType{
		ID:   356,
		Type: "L",
		Name: `White smoky haze`,
	},
	357: &SpellAnimationType{
		ID:   357,
		Type: "W",
		Name: `Super long delayed blue sparkles fall from sky`,
	},
	358: &SpellAnimationType{
		ID:   358,
		Type: "W",
		Name: `Long delayed falling confetti lights`,
	},
	359: &SpellAnimationType{
		ID:   359,
		Type: "W",
		Name: `Delayed falling lights and aura`,
	},
	360: &SpellAnimationType{
		ID:   360,
		Type: "W",
		Name: `Long delayed falling blue sparkles`,
	},
	361: &SpellAnimationType{
		ID:   361,
		Type: "W",
		Name: `Delayed falling confetti lights`,
	},
	362: &SpellAnimationType{
		ID:   362,
		Type: "W",
		Name: `Delayed starry firework blossoms`,
	},
	363: &SpellAnimationType{
		ID:   363,
		Type: "S",
		Name: `"Green steam on caster, stars falling from sky"`,
	},
	364: &SpellAnimationType{
		ID:   364,
		Type: "W",
		Name: `Starry fireworks`,
	},
	365: &SpellAnimationType{
		ID:   365,
		Type: "W",
		Name: `Falling stars`,
	},
	366: &SpellAnimationType{
		ID:   366,
		Type: "W",
		Name: `Starry fireworks`,
	},
	367: &SpellAnimationType{
		ID:   367,
		Type: "W",
		Name: `"Starry fireworks, blue swirling stars"`,
	},
	368: &SpellAnimationType{
		ID:   368,
		Type: "R",
		Name: `Icy mist around caster`,
	},
	369: &SpellAnimationType{
		ID:   369,
		Type: "R",
		Name: `Sparkling flames around caster`,
	},
	370: &SpellAnimationType{
		ID:   370,
		Type: "T",
		Name: `Portal to Akanon`,
	},
	371: &SpellAnimationType{
		ID:   371,
		Type: "T",
		Name: `Portal to Erudin`,
	},
	372: &SpellAnimationType{
		ID:   372,
		Type: "T",
		Name: `Portal to Felwithe`,
	},
	373: &SpellAnimationType{
		ID:   373,
		Type: "T",
		Name: `Portal to Freeport`,
	},
	374: &SpellAnimationType{
		ID:   374,
		Type: "T",
		Name: `Portal to Kelethin`,
	},
	375: &SpellAnimationType{
		ID:   375,
		Type: "T",
		Name: `Portal to Halas`,
	},
	376: &SpellAnimationType{
		ID:   376,
		Type: "T",
		Name: `Portal to Kaladim`,
	},
	377: &SpellAnimationType{
		ID:   377,
		Type: "T",
		Name: `Portal to Neriak`,
	},
	378: &SpellAnimationType{
		ID:   378,
		Type: "T",
		Name: `Portal to Oggok`,
	},
	379: &SpellAnimationType{
		ID:   379,
		Type: "T",
		Name: `Portal to the Plane of Knowledge`,
	},
	380: &SpellAnimationType{
		ID:   380,
		Type: "T",
		Name: `Portal to Qeynos`,
	},
	381: &SpellAnimationType{
		ID:   381,
		Type: "T",
		Name: `Portal to Rivervale`,
	},
	382: &SpellAnimationType{
		ID:   382,
		Type: "T",
		Name: `Portal to Shar Vahl`,
	},
	383: &SpellAnimationType{
		ID:   383,
		Type: "T",
		Name: `Portal to Grobb`,
	},
	384: &SpellAnimationType{
		ID:   384,
		Type: "T",
		Name: `Portal to Cabilis`,
	},
	385: &SpellAnimationType{
		ID:   385,
		Type: "T",
		Name: `Portal to Dranik's Scar`,
	},
	386: &SpellAnimationType{
		ID:   386,
		Type: "R",
		Name: `Dizzy stars swirling over head`,
	},
	387: &SpellAnimationType{
		ID:   387,
		Type: "R",
		Name: `Sleepy Z's swirling over head`,
	},
	388: &SpellAnimationType{
		ID:   388,
		Type: "-",
		Name: `None`,
	},
	389: &SpellAnimationType{
		ID:   389,
		Type: "R",
		Name: `Green bubbles up around head`,
	},
	390: &SpellAnimationType{
		ID:   390,
		Type: "R",
		Name: `Little black smoke clouds fall from torso`,
	},
	391: &SpellAnimationType{
		ID:   391,
		Type: "R",
		Name: `Black spots and blue skulls around head`,
	},
	392: &SpellAnimationType{
		ID:   392,
		Type: "R",
		Name: `Green and brown dust cloud at feet`,
	},
	393: &SpellAnimationType{
		ID:   393,
		Type: "R",
		Name: `Laughing blue skull and crossbones above head`,
	},
	394: &SpellAnimationType{
		ID:   394,
		Type: "R",
		Name: `Bright red smoke falling from torso`,
	},
	395: &SpellAnimationType{
		ID:   395,
		Type: "R",
		Name: `Red sparkles flash around head`,
	},
	396: &SpellAnimationType{
		ID:   396,
		Type: "R",
		Name: `Putrid green filth rises around feet`,
	},
	397: &SpellAnimationType{
		ID:   397,
		Type: "R",
		Name: `Purple sparkles emanate from head`,
	},
	398: &SpellAnimationType{
		ID:   398,
		Type: "-",
		Name: `None`,
	},
	399: &SpellAnimationType{
		ID:   399,
		Type: "-",
		Name: `None`,
	},
	400: &SpellAnimationType{
		ID:   400,
		Type: "R",
		Name: `Yellow-purple gaseous flame from torso`,
	},
	401: &SpellAnimationType{
		ID:   401,
		Type: "R",
		Name: `White bubbles floating around head`,
	},
	402: &SpellAnimationType{
		ID:   402,
		Type: "-",
		Name: `None`,
	},
	403: &SpellAnimationType{
		ID:   403,
		Type: "R",
		Name: `Yellow gaseous flame in torso`,
	},
	404: &SpellAnimationType{
		ID:   404,
		Type: "-",
		Name: `None`,
	},
	405: &SpellAnimationType{
		ID:   405,
		Type: "R",
		Name: `Bubbles rising up around feet`,
	},
	406: &SpellAnimationType{
		ID:   406,
		Type: "R",
		Name: `Little white sparkles`,
	},
	407: &SpellAnimationType{
		ID:   407,
		Type: "-",
		Name: `None`,
	},
	408: &SpellAnimationType{
		ID:   408,
		Type: "R",
		Name: `Star glowing above head`,
	},
	409: &SpellAnimationType{
		ID:   409,
		Type: "-",
		Name: `None`,
	},
	410: &SpellAnimationType{
		ID:   410,
		Type: "-",
		Name: `Hand sparkles, rising blue bubbles`,
	},
	411: &SpellAnimationType{
		ID:   411,
		Type: "R",
		Name: `Rising white bubbles`,
	},
	412: &SpellAnimationType{
		ID:   412,
		Type: "R",
		Name: `Light blue aura, white sparkles`,
	},
	413: &SpellAnimationType{
		ID:   413,
		Type: "R",
		Name: `Blue aura inside caster`,
	},
	414: &SpellAnimationType{
		ID:   414,
		Type: "R",
		Name: `Subtle star glowing inside caster's torso`,
	},
	415: &SpellAnimationType{
		ID:   415,
		Type: "R",
		Name: `Orange clouds emanating from caster's hands`,
	},
	416: &SpellAnimationType{
		ID:   416,
		Type: "-",
		Name: `None`,
	},
	417: &SpellAnimationType{
		ID:   417,
		Type: "-",
		Name: `None`,
	},
	418: &SpellAnimationType{
		ID:   418,
		Type: "-",
		Name: `None`,
	},
	419: &SpellAnimationType{
		ID:   419,
		Type: "-",
		Name: `None`,
	},
	420: &SpellAnimationType{
		ID:   420,
		Type: "-",
		Name: `None`,
	},
	421: &SpellAnimationType{
		ID:   421,
		Type: "-",
		Name: `None`,
	},
	422: &SpellAnimationType{
		ID:   422,
		Type: "-",
		Name: `None`,
	},
	423: &SpellAnimationType{
		ID:   423,
		Type: "-",
		Name: `None`,
	},
	424: &SpellAnimationType{
		ID:   424,
		Type: "-",
		Name: `None`,
	},
	425: &SpellAnimationType{
		ID:   425,
		Type: "-",
		Name: `None`,
	},
	426: &SpellAnimationType{
		ID:   426,
		Type: "T",
		Name: `Silent portal gate`,
	},
	427: &SpellAnimationType{
		ID:   427,
		Type: "-",
		Name: `None`,
	},
	428: &SpellAnimationType{
		ID:   428,
		Type: "-",
		Name: `None`,
	},
	429: &SpellAnimationType{
		ID:   429,
		Type: "-",
		Name: `None`,
	},
	430: &SpellAnimationType{
		ID:   430,
		Type: "-",
		Name: `None`,
	},
	431: &SpellAnimationType{
		ID:   431,
		Type: "-",
		Name: `None`,
	},
	432: &SpellAnimationType{
		ID:   432,
		Type: "-",
		Name: `None`,
	},
	433: &SpellAnimationType{
		ID:   433,
		Type: "-",
		Name: `None`,
	},
	434: &SpellAnimationType{
		ID:   434,
		Type: "D",
		Name: `Black smoke on ground, white around caster`,
	},
	435: &SpellAnimationType{
		ID:   435,
		Type: "-",
		Name: `None`,
	},
	436: &SpellAnimationType{
		ID:   436,
		Type: "C",
		Name: `Pink glow and green sparkles around caster's hands`,
	},
	437: &SpellAnimationType{
		ID:   437,
		Type: "X",
		Name: `Alteration Sound Effect 3`,
	},
	438: &SpellAnimationType{
		ID:   438,
		Type: "X",
		Name: `Alteration Sound Effect 3`,
	},
	439: &SpellAnimationType{
		ID:   439,
		Type: "C",
		Name: `Blue summoning circle under caster`,
	},
	440: &SpellAnimationType{
		ID:   440,
		Type: "C",
		Name: `Brown summoning circle under caster`,
	},
	441: &SpellAnimationType{
		ID:   441,
		Type: "C",
		Name: `Green summoning circle under caster`,
	},
	442: &SpellAnimationType{
		ID:   442,
		Type: "C",
		Name: `Red summoning circle under caster`,
	},
	443: &SpellAnimationType{
		ID:   443,
		Type: "C",
		Name: `White summoning circle under caster`,
	},
	444: &SpellAnimationType{
		ID:   444,
		Type: "C",
		Name: `Black smoke emanating from caster's hands`,
	},
	445: &SpellAnimationType{
		ID:   445,
		Type: "R",
		Name: `Gnomish balloon appears above caster and ropes lower down`,
	},
	446: &SpellAnimationType{
		ID:   446,
		Type: "P",
		Name: `Putrid green cloud with black insects`,
	},
	447: &SpellAnimationType{
		ID:   447,
		Type: "C",
		Name: `Pillar of flames around caster`,
	},
	448: &SpellAnimationType{
		ID:   448,
		Type: "C",
		Name: `Blue cloud and lightning around caster`,
	},
	449: &SpellAnimationType{
		ID:   449,
		Type: "-",
		Name: `None`,
	},
	450: &SpellAnimationType{
		ID:   450,
		Type: "R",
		Name: `Hearts rise up and swirl around caster`,
	},
	451: &SpellAnimationType{
		ID:   451,
		Type: "R",
		Name: `Black cloud around caster's torso and feet`,
	},
	452: &SpellAnimationType{
		ID:   452,
		Type: "R",
		Name: `Raincloud and rain on caster`,
	},
	453: &SpellAnimationType{
		ID:   453,
		Type: "R",
		Name: `Cloak of leaves surrounds caster`,
	},
	454: &SpellAnimationType{
		ID:   454,
		Type: "R",
		Name: `Light blue wisps bubble around caster`,
	},
	455: &SpellAnimationType{
		ID:   455,
		Type: "R",
		Name: `Blue and white cross glows from caster`,
	},
	456: &SpellAnimationType{
		ID:   456,
		Type: "R",
		Name: `Jester hats and spinning coins surround caster`,
	},
	457: &SpellAnimationType{
		ID:   457,
		Type: "R",
		Name: `Air tornado envelops caster with white cloud beneath`,
	},
	458: &SpellAnimationType{
		ID:   458,
		Type: "R",
		Name: `Circle of clear blue water surrounds caster`,
	},
	459: &SpellAnimationType{
		ID:   459,
		Type: "R",
		Name: `Blue crystals surround caster's hands and feet`,
	},
	460: &SpellAnimationType{
		ID:   460,
		Type: "R",
		Name: `Red crystals surround caster's hands and feet`,
	},
	461: &SpellAnimationType{
		ID:   461,
		Type: "R",
		Name: `Ice clouds billow from caster's hands and feet`,
	},
	462: &SpellAnimationType{
		ID:   462,
		Type: "R",
		Name: `Flames rise from caster's torso and hands`,
	},
	463: &SpellAnimationType{
		ID:   463,
		Type: "R",
		Name: `Red and blue crystals and smoke emanate from caster`,
	},
	464: &SpellAnimationType{
		ID:   464,
		Type: "R",
		Name: `Black and red veins flash around caster`,
	},
	465: &SpellAnimationType{
		ID:   465,
		Type: "R",
		Name: `Black clock/compass spins around caster`,
	},
	466: &SpellAnimationType{
		ID:   466,
		Type: "R",
		Name: `Blood spikes jut out from caster and drip down`,
	},
	467: &SpellAnimationType{
		ID:   467,
		Type: "R",
		Name: `Pink glowing symbol on caster`,
	},
	468: &SpellAnimationType{
		ID:   468,
		Type: "R",
		Name: `Bladed orbs spin and swirl around caster`,
	},
	469: &SpellAnimationType{
		ID:   469,
		Type: "R",
		Name: `Caster's hands glow with fire and white smoke rises`,
	},
	470: &SpellAnimationType{
		ID:   470,
		Type: "R",
		Name: `Black crystals and smoke at caster's feet`,
	},
	471: &SpellAnimationType{
		ID:   471,
		Type: "R",
		Name: `Column of swirling air rises up around caster`,
	},
	472: &SpellAnimationType{
		ID:   472,
		Type: "R",
		Name: `Caster's hands glow with black and blue smoke`,
	},
	473: &SpellAnimationType{
		ID:   473,
		Type: "R",
		Name: `Caster's feet burn with flickering flames`,
	},
	474: &SpellAnimationType{
		ID:   474,
		Type: "R",
		Name: `Column of water around caster with black particles`,
	},
	475: &SpellAnimationType{
		ID:   475,
		Type: "R",
		Name: `Clovers and dandelions sprout up around caster's feet`,
	},
	476: &SpellAnimationType{
		ID:   476,
		Type: "R",
		Name: `Stormcloud forms over caster and rain falls`,
	},
	477: &SpellAnimationType{
		ID:   477,
		Type: "R",
		Name: `Stormcloud forms over caster and snow falls`,
	},
	478: &SpellAnimationType{
		ID:   478,
		Type: "R",
		Name: `Stormcloud forms over caster and lightning strikes`,
	},
	479: &SpellAnimationType{
		ID:   479,
		Type: "C",
		Name: `Caster flips a gold coin`,
	},
	480: &SpellAnimationType{
		ID:   480,
		Type: "W",
		Name: `Yellow/red/green sparkling firework`,
	},
	481: &SpellAnimationType{
		ID:   481,
		Type: "W",
		Name: `Yellow/white sparkling firework`,
	},
	482: &SpellAnimationType{
		ID:   482,
		Type: "W",
		Name: `Yellow streaming firework`,
	},
	483: &SpellAnimationType{
		ID:   483,
		Type: "W",
		Name: `White and blue sparkling firework`,
	},
	484: &SpellAnimationType{
		ID:   484,
		Type: "W",
		Name: `Firework of Bertoxxulous`,
	},
	485: &SpellAnimationType{
		ID:   485,
		Type: "W",
		Name: `Firework of Cazic-Thule`,
	},
	486: &SpellAnimationType{
		ID:   486,
		Type: "W",
		Name: `Firework of Karana`,
	},
	487: &SpellAnimationType{
		ID:   487,
		Type: "W",
		Name: `Firework of Mithaniel Marr`,
	},
	488: &SpellAnimationType{
		ID:   488,
		Type: "W",
		Name: `Firework of Rallos Zek`,
	},
	489: &SpellAnimationType{
		ID:   489,
		Type: "W",
		Name: `Firework of Veeshan`,
	},
	490: &SpellAnimationType{
		ID:   490,
		Type: "W",
		Name: `Firework of Brell Serilis`,
	},
	491: &SpellAnimationType{
		ID:   491,
		Type: "W",
		Name: `Firework of Prexus`,
	},
	492: &SpellAnimationType{
		ID:   492,
		Type: "W",
		Name: `Firework of the Tribunal`,
	},
	493: &SpellAnimationType{
		ID:   493,
		Type: "W",
		Name: `Firework of Quellious`,
	},
	494: &SpellAnimationType{
		ID:   494,
		Type: "W",
		Name: `Firework of Rodcet Nife`,
	},
	495: &SpellAnimationType{
		ID:   495,
		Type: "W",
		Name: `Firework of Tunare`,
	},
	496: &SpellAnimationType{
		ID:   496,
		Type: "W",
		Name: `Firework of Solusek Ro`,
	},
	497: &SpellAnimationType{
		ID:   497,
		Type: "W",
		Name: `Firework of Innoruuk`,
	},
	498: &SpellAnimationType{
		ID:   498,
		Type: "W",
		Name: `Firework of Bristlebane`,
	},
	499: &SpellAnimationType{
		ID:   499,
		Type: "W",
		Name: `Firework of Erollisi Marr`,
	},
	500: &SpellAnimationType{
		ID:   500,
		Type: "W",
		Name: `Firework of History`,
	},
}
*/
