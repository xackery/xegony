package file

import (
	"github.com/xackery/xegony/model"
)

func loadSpellEffectTypeDefault() model.SpellEffectTypes {
	return model.SpellEffectTypes{
		{
			ID:   0,
			Type: 1,
			Name: "Current HP",
		},
		{
			ID:   1,
			Type: 1,
			Name: "Armor Class",
		},
		{
			ID:   2,
			Type: 1,
			Name: "Attack Rating",
		},
		{
			ID:   3,
			Type: 1,
			Name: "Movement Speed",
		},
		{
			ID:   4,
			Type: 1,
			Name: "Strength",
		},
		{
			ID:   5,
			Type: 1,
			Name: "Dexterity",
		},
		{
			ID:   6,
			Type: 1,
			Name: "Agility",
		},
		{
			ID:   7,
			Type: 1,
			Name: "Stamina",
		},
		{
			ID:   8,
			Type: 1,
			Name: "Intelligence",
		},
		{
			ID:   9,
			Type: 1,
			Name: "Wisdom",
		},
		{
			ID:   10,
			Type: 1,
			Name: "Charisma",
		},
		{
			ID:   11,
			Type: 1,
			Name: "Attack Speed",
		},
		{
			ID:   12,
			Type: 0,
			Name: "Invisibility for a Random Duration",
		},
		{
			ID:   13,
			Type: 0,
			Name: "See Invisible",
		},
		{
			ID:   14,
			Type: 0,
			Name: "Water Breathing",
		},
		{
			ID:   15,
			Type: 1,
			Name: "Current Mana",
		},
		{
			ID:   16,
			Type: 0,
			Name: "NPC Frenzy Radius (Not Used)",
		},
		{
			ID:   17,
			Type: 0,
			Name: "NPC Awareness (Not Used)",
		},
		{
			ID:   18,
			Type: 2,
			Name: "Pacify",
		},
		{
			ID:   19,
			Type: 1,
			Name: "Temporary Standing with Target NPC's Faction",
		},
		{
			ID:   20,
			Type: 0,
			Name: "Blindness",
		},
		{
			ID:   21,
			Type: 0,
			Name: "Stun Targets up to Level",
		},
		{
			ID:   22,
			Type: 0,
			Name: "Charm Targets up to Level",
		},
		{
			ID:   23,
			Type: 0,
			Name: "Fear Targets up to Level",
		},
		{
			ID:   24,
			Type: 1,
			Name: "Stamina (No Longer Used)",
		},
		{
			ID:   25,
			Type: 0,
			Name: "Bind Respawn Point",
		},
		{
			ID:   26,
			Type: 0,
			Name: "Return to Respawn Point",
		},
		{
			ID:   27,
			Type: 0,
			Name: "Attempt to Remove Magical Effect",
		},
		{
			ID:   28,
			Type: 2,
			Name: "Invisibility to the Undead for a Random Duration",
		},
		{
			ID:   29,
			Type: 2,
			Name: "Invisibility to Animals for a Random Duration",
		},
		{
			ID:   30,
			Type: 0,
			Name: "NPC Frenzy Radius",
		},
		{
			ID:   31,
			Type: 0,
			Name: "Mesmerize",
		},
		{
			ID:   32,
			Type: 0,
			Name: "Summon Item:",
		},
		{
			ID:   33,
			Type: 0,
			Name: "Summon Pet:",
		},
		{
			ID:   34,
			Type: 0,
			Name: "Confuse (Not Used)",
		},
		{
			ID:   35,
			Type: 1,
			Name: "Disease Counter",
		},
		{
			ID:   36,
			Type: 1,
			Name: "Poison Counter",
		},
		{
			ID:   37,
			Type: 2,
			Name: "Detect Hostile (Not Used)",
		},
		{
			ID:   38,
			Type: 2,
			Name: "Detect Magic (Not Used)",
		},
		{
			ID:   39,
			Type: 2,
			Name: "Detect Poison (Not Used)",
		},
		{
			ID:   40,
			Type: 2,
			Name: "Temporary Invulnerability",
		},
		{
			ID:   41,
			Type: 2,
			Name: "Destroy Target (With No Credit)",
		},
		{
			ID:   42,
			Type: 2,
			Name: "Random Teleport within Spell's Range",
		},
		{
			ID:   43,
			Type: 1,
			Name: "Crippling Blow Chance",
		},
		{
			ID:   44,
			Type: 2,
			Name: "Lycanthropy",
		},
		{
			ID:   45,
			Type: 2,
			Name: "Vampirism (Not Used)",
		},
		{
			ID:   46,
			Type: 1,
			Name: "Resistance to Fire",
		},
		{
			ID:   47,
			Type: 1,
			Name: "Resistance to Cold",
		},
		{
			ID:   48,
			Type: 1,
			Name: "Resistance to Poison",
		},
		{
			ID:   49,
			Type: 1,
			Name: "Resistance to Disease",
		},
		{
			ID:   50,
			Type: 1,
			Name: "Resistance to Magic",
		},
		{
			ID:   51,
			Type: 2,
			Name: "Detect Traps (Not Used)",
		},
		{
			ID:   52,
			Type: 2,
			Name: "Locate Nearest Undead",
		},
		{
			ID:   53,
			Type: 2,
			Name: "Locate Nearest Summoned",
		},
		{
			ID:   54,
			Type: 2,
			Name: "Locate Nearest Animal",
		},
		{
			ID:   55,
			Type: 1,
			Name: "Damage Absorption",
		},
		{
			ID:   56,
			Type: 2,
			Name: "Spin to Face North",
		},
		{
			ID:   57,
			Type: 0,
			Name: "Levitation",
		},
		{
			ID:   58,
			Type: 0,
			Name: "Illusion:",
		},
		{
			ID:   59,
			Type: 1,
			Name: "Damage Shield",
		},
		{
			ID:   60,
			Type: 2,
			Name: "Transfer Item (Not Used)",
		},
		{
			ID:   61,
			Type: 2,
			Name: "Identify Item",
		},
		{
			ID:   62,
			Type: 2,
			Name: "ItemID (Not Used)",
		},
		{
			ID:   63,
			Type: 0,
			Name: "Attempt to Wipe Hate List",
		},
		{
			ID:   64,
			Type: 2,
			Name: "Spinning Stun",
		},
		{
			ID:   65,
			Type: 0,
			Name: "Infravision (Heat Vision)",
		},
		{
			ID:   66,
			Type: 0,
			Name: "Ultravision (Night Vision)",
		},
		{
			ID:   67,
			Type: 0,
			Name: "Summon and Take Control of Seeing Eye",
		},
		{
			ID:   68,
			Type: 2,
			Name: "Destroy Pet and Reclaim Some Mana",
		},
		{
			ID:   69,
			Type: 1,
			Name: "Maximum HP",
		},
		{
			ID:   70,
			Type: 2,
			Name: "Corpse Bomb (Not Used)",
		},
		{
			ID:   71,
			Type: 0,
			Name: "Summon Necromancer Pet:",
		},
		{
			ID:   72,
			Type: 2,
			Name: "Preserve Corpse (Not Used)",
		},
		{
			ID:   73,
			Type: 2,
			Name: "Bind Sight to Target",
		},
		{
			ID:   74,
			Type: 0,
			Name: "Feign Death",
		},
		{
			ID:   75,
			Type: 2,
			Name: "Voice Transfer to Target",
		},
		{
			ID:   76,
			Type: 2,
			Name: "Add Hostile Proximity Alarm",
		},
		{
			ID:   77,
			Type: 2,
			Name: "Locate Nearest Corpse (Optionally of Target)",
		},
		{
			ID:   78,
			Type: 1,
			Name: "Chance to Absorb Magical Attack",
		},
		{
			ID:   79,
			Type: 1,
			Name: "Current HP",
		},
		{
			ID:   80,
			Type: 2,
			Name: "Enchant Light (Not Used)",
		},
		{
			ID:   81,
			Type: 0,
			Name: "Revive with Experience Gain",
		},
		{
			ID:   82,
			Type: 2,
			Name: "Summon Player to Self",
		},
		{
			ID:   83,
			Type: 0,
			Name: "Teleport To:",
		},
		{
			ID:   84,
			Type: 0,
			Name: "Toss Up",
		},
		{
			ID:   85,
			Type: 0,
			Name: "Add Melee Proc:",
		},
		{
			ID:   86,
			Type: 0,
			Name: "NPC Reaction Radius",
		},
		{
			ID:   87,
			Type: 1,
			Name: "Vision Magnification (Adjust Field of View)",
		},
		{
			ID:   88,
			Type: 0,
			Name: "Evacuate To:",
		},
		{
			ID:   89,
			Type: 1,
			Name: "Physical Size",
		},
		{
			ID:   90,
			Type: 2,
			Name: "Cloak (Not Used)",
		},
		{
			ID:   91,
			Type: 0,
			Name: "Summon Corpse",
		},
		{
			ID:   92,
			Type: 1,
			Name: "Hate with Target",
		},
		{
			ID:   93,
			Type: 0,
			Name: "Cancel Adverse Weather",
		},
		{
			ID:   94,
			Type: 2,
			Name: "Limitation: Drop Spell In Combat",
		},
		{
			ID:   95,
			Type: 2,
			Name: "Sacrifice Target Character with Confirmation",
		},
		{
			ID:   96,
			Type: 0,
			Name: "Silence All Spellcasting",
		},
		{
			ID:   97,
			Type: 1,
			Name: "Maximum Mana",
		},
		{
			ID:   98,
			Type: 1,
			Name: "Attack Speed (V2)",
		},
		{
			ID:   99,
			Type: 1,
			Name: "Movement Speed",
		},
		{
			ID:   100,
			Type: 1,
			Name: "Current HP Over Time",
		},
		{
			ID:   101,
			Type: 0,
			Name: "Complete Heal",
		},
		{
			ID:   102,
			Type: 0,
			Name: "Immunity to Fear",
		},
		{
			ID:   103,
			Type: 2,
			Name: "Summon Current Pet to Self",
		},
		{
			ID:   104,
			Type: 0,
			Name: "Translocate",
		},
		{
			ID:   105,
			Type: 0,
			Name: "Prevent Gating to Respawn Point",
		},
		{
			ID:   106,
			Type: 0,
			Name: "Summon Beastlord Pet:",
		},
		{
			ID:   107,
			Type: 1,
			Name: "NPC Level",
		},
		{
			ID:   108,
			Type: 0,
			Name: "Summon Familiar:",
		},
		{
			ID:   109,
			Type: 0,
			Name: "Summon Item Into Bag:",
		},
		{
			ID:   110,
			Type: 1,
			Name: "Archery (Not Used)",
		},
		{
			ID:   111,
			Type: 1,
			Name: "All Magical Resistances",
		},
		{
			ID:   112,
			Type: 1,
			Name: "Effective Spellcasting Level",
		},
		{
			ID:   113,
			Type: 0,
			Name: "Summon Mount:",
		},
		{
			ID:   114,
			Type: 1,
			Name: "Hate Generation Modifier",
		},
		{
			ID:   115,
			Type: 0,
			Name: "Satisfy Hunger and Thirst",
		},
		{
			ID:   116,
			Type: 1,
			Name: "Curse Counter",
		},
		{
			ID:   117,
			Type: 2,
			Name: "Allow Weapons to Hit Magical Targets",
		},
		{
			ID:   118,
			Type: 1,
			Name: "Performance Amplification",
		},
		{
			ID:   119,
			Type: 1,
			Name: "Attack Speed V3",
		},
		{
			ID:   120,
			Type: 1,
			Name: "Healing Modifier",
		},
		{
			ID:   121,
			Type: 1,
			Name: "Reverse (Healing) Damage Shield",
		},
		{
			ID:   122,
			Type: 1,
			Name: "Skill Reduction (Not Used)",
		},
		{
			ID:   123,
			Type: 0,
			Name: "Screech",
		},
		{
			ID:   124,
			Type: 1,
			Name: "Damage Modifier (Focus)",
		},
		{
			ID:   125,
			Type: 1,
			Name: "Healing Modifier (Focus)",
		},
		{
			ID:   126,
			Type: 1,
			Name: "Spell Resistance (Focus)",
		},
		{
			ID:   127,
			Type: 1,
			Name: "Spellcasting Speed (Focus)",
		},
		{
			ID:   128,
			Type: 1,
			Name: "Spell Durations (Focus)",
		},
		{
			ID:   129,
			Type: 1,
			Name: "Spellcasting Range (Focus)",
		},
		{
			ID:   130,
			Type: 1,
			Name: "Spellcasting Hate Modifier (Focus)",
		},
		{
			ID:   131,
			Type: 1,
			Name: "Chance to Conserve Reagents (Focus)",
		},
		{
			ID:   132,
			Type: 1,
			Name: "Spellcasting Mana Cost (Focus)",
		},
		{
			ID:   133,
			Type: 1,
			Name: "Stun Time Modifier (Focus)",
		},
		{
			ID:   134,
			Type: 0,
			Name: "Limit to Maximum Spell Level:",
		},
		{
			ID:   135,
			Type: 0,
			Name: "Limit to Resistance Type:",
		},
		{
			ID:   136,
			Type: 0,
			Name: "Limit to Target Type:",
		},
		{
			ID:   137,
			Type: 0,
			Name: "Limit to Effect:",
		},
		{
			ID:   138,
			Type: 0,
			Name: "Limit to Spell Type:",
		},
		{
			ID:   139,
			Type: 0,
			Name: "Limit to Spell:",
		},
		{
			ID:   140,
			Type: 0,
			Name: "Limit to Minimum Spell Duration:",
		},
		{
			ID:   141,
			Type: 0,
			Name: "Limit to Instant Spells",
		},
		{
			ID:   142,
			Type: 0,
			Name: "Limit to Minimum Spell Level:",
		},
		{
			ID:   143,
			Type: 0,
			Name: "Limit to Minimum Casting Time:",
		},
		{
			ID:   144,
			Type: 0,
			Name: "Limit to Maximum Casting Time:",
		},
		{
			ID:   145,
			Type: 0,
			Name: "Teleport (V2):",
		},
		{
			ID:   146,
			Type: 1,
			Name: "Resistance to Electricity (Not Used)",
		},
		{
			ID:   147,
			Type: 1,
			Name: "Current HP by Percentage of Maximum:",
		},
		{
			ID:   148,
			Type: 0,
			Name: "Stacking - Block",
		},
		{
			ID:   149,
			Type: 0,
			Name: "Stacking - Overwrite",
		},
		{
			ID:   150,
			Type: 1,
			Name: "Chance to Save from Death",
		},
		{
			ID:   151,
			Type: 0,
			Name: "Suspend Pet",
		},
		{
			ID:   152,
			Type: 0,
			Name: "Summon Temporary Pets:",
		},
		{
			ID:   153,
			Type: 0,
			Name: "Balance HP Across Group",
		},
		{
			ID:   154,
			Type: 0,
			Name: "Attempt to Remove Detrimental Effect",
		},
		{
			ID:   155,
			Type: 1,
			Name: "Spell Critical Damage",
		},
		{
			ID:   156,
			Type: 2,
			Name: "Illusion: Target",
		},
		{
			ID:   157,
			Type: 1,
			Name: "Spell Damage Shield",
		},
		{
			ID:   158,
			Type: 1,
			Name: "Chance to Reflect Spells",
		},
		{
			ID:   159,
			Type: 1,
			Name: "All Stats",
		},
		{
			ID:   160,
			Type: 1,
			Name: "Drunkenness (Not Used)",
		},
		{
			ID:   161,
			Type: 1,
			Name: "Spell Damage Mitigation",
		},
		{
			ID:   162,
			Type: 1,
			Name: "Melee Damage Mitigation",
		},
		{
			ID:   163,
			Type: 0,
			Name: "Block All Attacks",
		},
		{
			ID:   164,
			Type: 0,
			Name: "Examine LDoN Chest for Traps",
		},
		{
			ID:   165,
			Type: 0,
			Name: "Disarm LDoN Trap",
		},
		{
			ID:   166,
			Type: 0,
			Name: "Unlock LDoN Chest",
		},
		{
			ID:   167,
			Type: 1,
			Name: "Pet Power",
		},
		{
			ID:   168,
			Type: 1,
			Name: "Melee Damage Mitigation V2",
		},
		{
			ID:   169,
			Type: 1,
			Name: "Critical Hit Chance",
		},
		{
			ID:   170,
			Type: 1,
			Name: "Spell Critical Chance",
		},
		{
			ID:   171,
			Type: 1,
			Name: "Crippling Blow Chance",
		},
		{
			ID:   172,
			Type: 1,
			Name: "Melee Avoidance",
		},
		{
			ID:   173,
			Type: 1,
			Name: "Chance to Riposte",
		},
		{
			ID:   174,
			Type: 1,
			Name: "Chance to Dodge",
		},
		{
			ID:   175,
			Type: 1,
			Name: "Chance to Parry",
		},
		{
			ID:   176,
			Type: 1,
			Name: "Offhand Attack Chance",
		},
		{
			ID:   177,
			Type: 1,
			Name: "Double Attack Chance",
		},
		{
			ID:   178,
			Type: 0,
			Name: "Lifetap from Melee Attacks",
		},
		{
			ID:   179,
			Type: 1,
			Name: "All Instrument Effectiveness",
		},
		{
			ID:   180,
			Type: 1,
			Name: "Chance to Resist Spells",
		},
		{
			ID:   181,
			Type: 1,
			Name: "Chance to Resist Fear",
		},
		{
			ID:   182,
			Type: 1,
			Name: "Attack Speed V4",
		},
		{
			ID:   183,
			Type: 1,
			Name: "Skill Checks",
		},
		{
			ID:   184,
			Type: 1,
			Name: "Hit Chance",
		},
		{
			ID:   185,
			Type: 1,
			Name: "Damage Modifier",
		},
		{
			ID:   186,
			Type: 1,
			Name: "Minimum Damage Modifier",
		},
		{
			ID:   187,
			Type: 0,
			Name: "Balance Mana Across Group",
		},
		{
			ID:   188,
			Type: 1,
			Name: "Chance to Block",
		},
		{
			ID:   189,
			Type: 1,
			Name: "Current Endurance",
		},
		{
			ID:   190,
			Type: 1,
			Name: "Maximum Endurance",
		},
		{
			ID:   191,
			Type: 0,
			Name: "Prevent Melee Attacks",
		},
		{
			ID:   192,
			Type: 1,
			Name: "Hate with Target",
		},
		{
			ID:   193,
			Type: 0,
			Name: "Skill Attack:",
		},
		{
			ID:   194,
			Type: 0,
			Name: "Remove Self from Hate List",
		},
		{
			ID:   195,
			Type: 1,
			Name: "Resistance to Stun",
		},
		{
			ID:   196,
			Type: 1,
			Name: "Chance to Strike Through",
		},
		{
			ID:   197,
			Type: 1,
			Name: "Skill Damage Taken:",
		},
		{
			ID:   198,
			Type: 1,
			Name: "Current Endurance",
		},
		{
			ID:   199,
			Type: 0,
			Name: "Attempt to Taunt the Target",
		},
		{
			ID:   200,
			Type: 1,
			Name: "Chance of Melee Attack Procs",
		},
		{
			ID:   201,
			Type: 1,
			Name: "Chance of Ranged Attack Procs",
		},
		{
			ID:   202,
			Type: 0,
			Name: "Illusion:",
		},
		{
			ID:   203,
			Type: 0,
			Name: "Mass Group Next Buff",
		},
		{
			ID:   204,
			Type: 0,
			Name: "Group Immunity to Fear",
		},
		{
			ID:   205,
			Type: 0,
			Name: "Rampage (Attack Nearby Targets on Hate Lists)",
		},
		{
			ID:   206,
			Type: 0,
			Name: "Area Attempt to Taunt",
		},
		{
			ID:   207,
			Type: 0,
			Name: "Extract Bone Chips from Meat",
		},
		{
			ID:   208,
			Type: 0,
			Name: "Purge Poison (Not Used)",
		},
		{
			ID:   209,
			Type: 0,
			Name: "Attempt to Remove Beneficial Effect",
		},
		{
			ID:   210,
			Type: 0,
			Name: "Pet Shield (Not Used)",
		},
		{
			ID:   211,
			Type: 0,
			Name: "Area Melee Attack",
		},
		{
			ID:   212,
			Type: 1,
			Name: "Spell Critical Chance and Mana Cost",
		},
		{
			ID:   213,
			Type: 1,
			Name: "Pet's Maximum HP",
		},
		{
			ID:   214,
			Type: 1,
			Name: "Maximum HP by Percentage",
		},
		{
			ID:   215,
			Type: 1,
			Name: "Pet Melee Avoidance",
		},
		{
			ID:   216,
			Type: 1,
			Name: "Melee Accuracy",
		},
		{
			ID:   217,
			Type: 0,
			Name: "Chance to Headshot for 32K Damage",
		},
		{
			ID:   218,
			Type: 1,
			Name: "Pet's Chance to Critical Hit",
		},
		{
			ID:   219,
			Type: 0,
			Name: "Critical Spell Hit vs Undead",
		},
		{
			ID:   220,
			Type: 1,
			Name: "Skill Damage Modifier",
		},
		{
			ID:   221,
			Type: 1,
			Name: "Weight Encumbrance",
		},
		{
			ID:   222,
			Type: 1,
			Name: "Chance to Block Attacks from Behind",
		},
		{
			ID:   223,
			Type: 1,
			Name: "Chance to Double Riposte",
		},
		{
			ID:   224,
			Type: 0,
			Name: "Perform Double Riposte",
		},
		{
			ID:   225,
			Type: 0,
			Name: "Perform Double Attack",
		},
		{
			ID:   226,
			Type: 0,
			Name: "Perform Two Handed Bash",
		},
		{
			ID:   227,
			Type: 1,
			Name: "Skill Timer",
		},
		{
			ID:   228,
			Type: 1,
			Name: "Falling Damage (Not Used)",
		},
		{
			ID:   229,
			Type: 1,
			Name: "Chance to Cast Through Interruptions",
		},
		{
			ID:   230,
			Type: 0,
			Name: "Extended Shielding (Not Used)",
		},
		{
			ID:   231,
			Type: 1,
			Name: "Chance for Bashes to Stun",
		},
		{
			ID:   232,
			Type: 1,
			Name: "Chance to Save from Death",
		},
		{
			ID:   233,
			Type: 1,
			Name: "Metabolism",
		},
		{
			ID:   234,
			Type: 1,
			Name: "Apply Poison Time (Not Used)",
		},
		{
			ID:   235,
			Type: 1,
			Name: "Chance to Channel Spells through Interruptions",
		},
		{
			ID:   236,
			Type: 0,
			Name: "Free Pet (Not Used)",
		},
		{
			ID:   237,
			Type: 0,
			Name: "Give Pets Group Buffs",
		},
		{
			ID:   238,
			Type: 0,
			Name: "Make Illusions Persist Across Zones",
		},
		{
			ID:   239,
			Type: 0,
			Name: "Feigned Cast On Chance (Not Used)",
		},
		{
			ID:   240,
			Type: 0,
			Name: "String Unbreakable (Not Used)",
		},
		{
			ID:   241,
			Type: 1,
			Name: "Pet Energy Reclaiming Efficiency",
		},
		{
			ID:   242,
			Type: 1,
			Name: "Chance to Wipe Hate Lists",
		},
		{
			ID:   243,
			Type: 1,
			Name: "Chance to Prevent Charm from Breaking",
		},
		{
			ID:   244,
			Type: 1,
			Name: "Chance to Prevent Root from Breaking",
		},
		{
			ID:   245,
			Type: 0,
			Name: "Trap Circumvention (Not Used)",
		},
		{
			ID:   246,
			Type: 0,
			Name: "Set Breathing Air Supply Level",
		},
		{
			ID:   247,
			Type: 1,
			Name: "Skill Cap:",
		},
		{
			ID:   248,
			Type: 0,
			Name: "Secondary Forte (Not Used)",
		},
		{
			ID:   249,
			Type: 1,
			Name: "Damage Modifier (V3?)",
		},
		{
			ID:   250,
			Type: 1,
			Name: "Chance of Spell Attack Procs",
		},
		{
			ID:   251,
			Type: 0,
			Name: "Consume Projectile:",
		},
		{
			ID:   252,
			Type: 1,
			Name: "Chance to Backstab from the Front",
		},
		{
			ID:   253,
			Type: 1,
			Name: "Front Backstab Minimum Damage",
		},
		{
			ID:   254,
			Type: 0,
			Name: "-",
		},
		{
			ID:   255,
			Type: 1,
			Name: "Shield Duration (Not Used)",
		},
		{
			ID:   256,
			Type: 0,
			Name: "Shroud of Stealth (Not Used)",
		},
		{
			ID:   257,
			Type: 0,
			Name: "Pet Discipline (Not Used)",
		},
		{
			ID:   258,
			Type: 1,
			Name: "Chance for Triple Backstab Damage",
		},
		{
			ID:   259,
			Type: 1,
			Name: "Direct Damage Mitigation",
		},
		{
			ID:   260,
			Type: 1,
			Name: "Singing Effectiveness",
		},
		{
			ID:   261,
			Type: 1,
			Name: "Cap on Singing Effectiveness",
		},
		{
			ID:   262,
			Type: 1,
			Name: "Stat Caps",
		},
		{
			ID:   263,
			Type: 1,
			Name: "Tradeskill Mastery (Not Used)",
		},
		{
			ID:   264,
			Type: 0,
			Name: "Hastened AA Skill (Not Used)",
		},
		{
			ID:   265,
			Type: 0,
			Name: "Immunity to Fizzling Spells",
		},
		{
			ID:   266,
			Type: 1,
			Name: "Chance of Two-Handed Triple Attack",
		},
		{
			ID:   267,
			Type: 0,
			Name: "Pet Discipline 2 (Not Used)",
		},
		{
			ID:   268,
			Type: 0,
			Name: "Reduce Tradeskill Fail (Not Used)",
		},
		{
			ID:   269,
			Type: 1,
			Name: "Maximum Bind Woundable Health",
		},
		{
			ID:   270,
			Type: 1,
			Name: "Range of Bard Songs",
		},
		{
			ID:   271,
			Type: 1,
			Name: "Minimum Movement Speed",
		},
		{
			ID:   272,
			Type: 1,
			Name: "Effective Spellcasting Level V2",
		},
		{
			ID:   273,
			Type: 1,
			Name: "Critical DoT Chance",
		},
		{
			ID:   274,
			Type: 1,
			Name: "Chance to Critical Heal",
		},
		{
			ID:   275,
			Type: 1,
			Name: "Chance to Critical Mend",
		},
		{
			ID:   276,
			Type: 1,
			Name: "Chance of Offhand Attacks",
		},
		{
			ID:   277,
			Type: 1,
			Name: "Chance of Saving from Death",
		},
		{
			ID:   278,
			Type: 1,
			Name: "Chance to Inflict Finishing Blow on Fleeing Enemy",
		},
		{
			ID:   279,
			Type: 1,
			Name: "Chance to Flurry Attacks",
		},
		{
			ID:   280,
			Type: 1,
			Name: "Chance for Pet to Flurry Attacks",
		},
		{
			ID:   281,
			Type: 0,
			Name: "Pet Feign Death (Not Used)",
		},
		{
			ID:   282,
			Type: 1,
			Name: "Bind Wound Effectiveness",
		},
		{
			ID:   283,
			Type: 1,
			Name: "Chance to Perform Double Special Attack",
		},
		{
			ID:   284,
			Type: 0,
			Name: "Set Lay-on-Hands Heal (Not Used)",
		},
		{
			ID:   285,
			Type: 0,
			Name: "Nimble Evasion (Not Used)",
		},
		{
			ID:   286,
			Type: 1,
			Name: "Spell Damage (V3?)",
		},
		{
			ID:   287,
			Type: 1,
			Name: "Song/Spell Duration by 1 Tick (6s)",
		},
		{
			ID:   288,
			Type: 1,
			Name: "Chance to Knockback with Special Attacks",
		},
		{
			ID:   289,
			Type: 0,
			Name: "Cast New Spell when Wearing Off:",
		},
		{
			ID:   290,
			Type: 1,
			Name: "Maximum Movement Speed",
		},
		{
			ID:   291,
			Type: 0,
			Name: "Attempt to Remove Detrimental Effects",
		},
		{
			ID:   292,
			Type: 1,
			Name: "Chance to Strike Through V2",
		},
		{
			ID:   293,
			Type: 1,
			Name: "Resistance to Stuns from the Front",
		},
		{
			ID:   294,
			Type: 1,
			Name: "Critical Spell Chance",
		},
		{
			ID:   295,
			Type: 0,
			Name: "Reduce Timer Special (Not Used)",
		},
		{
			ID:   296,
			Type: 1,
			Name: "Spell Vulnerability Focus",
		},
		{
			ID:   297,
			Type: 1,
			Name: "Incoming Damage",
		},
		{
			ID:   298,
			Type: 1,
			Name: "Physical Size",
		},
		{
			ID:   299,
			Type: 0,
			Name: "Awaken Corpse as Temporary Pet",
		},
		{
			ID:   300,
			Type: 0,
			Name: "Summon Doppelganger",
		},
		{
			ID:   301,
			Type: 1,
			Name: "Archery Damage",
		},
		{
			ID:   302,
			Type: 1,
			Name: "Critical Hit Damage %",
		},
		{
			ID:   303,
			Type: 1,
			Name: "Critical Hit Damage",
		},
		{
			ID:   304,
			Type: 1,
			Name: "Chance to Avoid Offhand Attack Ripostes",
		},
		{
			ID:   305,
			Type: 1,
			Name: "Damage Shield Mitigation",
		},
		{
			ID:   306,
			Type: 0,
			Name: "Army of the Dead (Not Used)",
		},
		{
			ID:   307,
			Type: 0,
			Name: "Appraisal (Not Used)",
		},
		{
			ID:   308,
			Type: 0,
			Name: "Suspend Pet",
		},
		{
			ID:   309,
			Type: 0,
			Name: "Teleport to Caster's Bind Point",
		},
		{
			ID:   310,
			Type: 1,
			Name: "Reuse Timer Reduction",
		},
		{
			ID:   311,
			Type: 0,
			Name: "Limit - Not Applicable to Innate Weapon Procs",
		},
		{
			ID:   312,
			Type: 0,
			Name: "Temporarily Drop to Bottom of Hate List",
		},
		{
			ID:   313,
			Type: 1,
			Name: "Chance to Forage Additional Items",
		},
		{
			ID:   314,
			Type: 0,
			Name: "Invisibility (Fixed Duration)",
		},
		{
			ID:   315,
			Type: 0,
			Name: "Invisibility to Undead (Fixed Duration)",
		},
		{
			ID:   316,
			Type: 0,
			Name: "Invisibility to Animals (Fixed Duration) (Not Used)",
		},
		{
			ID:   317,
			Type: 1,
			Name: "Worn Item HP Regeneration Cap",
		},
		{
			ID:   318,
			Type: 1,
			Name: "Worn Item Mana Regeneration Cap",
		},
		{
			ID:   319,
			Type: 1,
			Name: "Chance for Heal-over-Time to Critical Tick",
		},
		{
			ID:   320,
			Type: 1,
			Name: "Shield Block Chance",
		},
		{
			ID:   321,
			Type: 1,
			Name: "Hate with Target",
		},
		{
			ID:   322,
			Type: 0,
			Name: "Teleport to Home City",
		},
		{
			ID:   323,
			Type: 0,
			Name: "Add Defensive Proc on Hit:",
		},
		{
			ID:   324,
			Type: 0,
			Name: "Convert HP to Mana",
		},
		{
			ID:   325,
			Type: 0,
			Name: "Chance Invs Break to AoE (Not Used)",
		},
		{
			ID:   326,
			Type: 1,
			Name: "Number of Spell Gem Slots",
		},
		{
			ID:   327,
			Type: 1,
			Name: "Maximum Number of Magical Effects",
		},
		{
			ID:   328,
			Type: 1,
			Name: "Damage Taken when Unconscious Before Death",
		},
		{
			ID:   329,
			Type: 1,
			Name: "Damage Absorption with Mana",
		},
		{
			ID:   330,
			Type: 1,
			Name: "Critical Hit Damage with All Skills",
		},
		{
			ID:   331,
			Type: 1,
			Name: "Chance to Salvage Tradeskill Failures",
		},
		{
			ID:   332,
			Type: 0,
			Name: "Summon to Corpse (Not Used)",
		},
		{
			ID:   333,
			Type: 0,
			Name: "Cast New Spell When Rune Fades:",
		},
		{
			ID:   334,
			Type: 1,
			Name: "Current HP V2",
		},
		{
			ID:   335,
			Type: 0,
			Name: "Block Next Spell Focus",
		},
		{
			ID:   336,
			Type: 0,
			Name: "Illusionary Target (Not Used)",
		},
		{
			ID:   337,
			Type: 1,
			Name: "Experience Gain Modifier",
		},
		{
			ID:   338,
			Type: 0,
			Name: "Summon And Resurrect All Corpses",
		},
		{
			ID:   339,
			Type: 0,
			Name: "Trigger Second Spell:",
		},
		{
			ID:   340,
			Type: 0,
			Name: "Add Chance to Trigger Spell:",
		},
		{
			ID:   341,
			Type: 1,
			Name: "Worn Item Attack Rating Cap",
		},
		{
			ID:   342,
			Type: 0,
			Name: "Prevent from Fleeing",
		},
		{
			ID:   343,
			Type: 0,
			Name: "Interrupt Casting",
		},
		{
			ID:   344,
			Type: 1,
			Name: "Chance of Channeling Item Effect through Interruption",
		},
		{
			ID:   345,
			Type: 1,
			Name: "Maximum Permitted Level for Assassinate",
		},
		{
			ID:   346,
			Type: 1,
			Name: "Maximum Permitted Level for HeadShot",
		},
		{
			ID:   347,
			Type: 1,
			Name: "Chance to Perform Double Ranged Attack",
		},
		{
			ID:   348,
			Type: 0,
			Name: "Limit - Minimum Mana Cost:",
		},
		{
			ID:   349,
			Type: 1,
			Name: "Hate Modifier with Shield Equipped",
		},
		{
			ID:   350,
			Type: 1,
			Name: "Current HP by Draining Mana",
		},
		{
			ID:   351,
			Type: 0,
			Name: "Persistent Effect (Not Used)",
		},
		{
			ID:   352,
			Type: 0,
			Name: "Increase Trap Count (Not Used)",
		},
		{
			ID:   353,
			Type: 0,
			Name: "Additional Aura (Not Used)",
		},
		{
			ID:   354,
			Type: 0,
			Name: "Deactivate All Traps (Not Used)",
		},
		{
			ID:   355,
			Type: 0,
			Name: "Learn Trap (Not Used)",
		},
		{
			ID:   356,
			Type: 0,
			Name: "Change Trigger Type (Not Used)",
		},
		{
			ID:   357,
			Type: 0,
			Name: "Silence Casting of Spells Affecting:",
		},
		{
			ID:   358,
			Type: 1,
			Name: "Current Mana",
		},
		{
			ID:   359,
			Type: 0,
			Name: "Passive Sense Trap (Not Used)",
		},
		{
			ID:   360,
			Type: 0,
			Name: "Add Chance to Proc On Kill Shot:",
		},
		{
			ID:   361,
			Type: 0,
			Name: "Add Chance to Cast Spell on Death:",
		},
		{
			ID:   362,
			Type: 1,
			Name: "Potion Belt Slots (Not Used)",
		},
		{
			ID:   363,
			Type: 1,
			Name: "Bandolier Slots (Not Used)",
		},
		{
			ID:   364,
			Type: 1,
			Name: "Chance to Triple Attack",
		},
		{
			ID:   365,
			Type: 0,
			Name: "Add Chance to Proc on Spellcast Kill Shot:",
		},
		{
			ID:   366,
			Type: 1,
			Name: "Damage Modifier with Shield Equipped",
		},
		{
			ID:   367,
			Type: 0,
			Name: "Set Body Type:",
		},
		{
			ID:   368,
			Type: 0,
			Name: "Faction Mod (Not Used)",
		},
		{
			ID:   369,
			Type: 1,
			Name: "Corruption Counter",
		},
		{
			ID:   370,
			Type: 1,
			Name: "Resistance to Corruption",
		},
		{
			ID:   371,
			Type: 1,
			Name: "Attack Speed V4",
		},
		{
			ID:   372,
			Type: 1,
			Name: "Forage Skill (Not Used)",
		},
		{
			ID:   373,
			Type: 0,
			Name: "Cast New Spell on Fade:",
		},
		{
			ID:   374,
			Type: 0,
			Name: "Apply Effect:",
		},
		{
			ID:   375,
			Type: 1,
			Name: "Critical DoT Tick Damage",
		},
		{
			ID:   376,
			Type: 0,
			Name: "Fling (Not Used)",
		},
		{
			ID:   377,
			Type: 0,
			Name: "NPC - Cast New Spell on Fade:",
		},
		{
			ID:   378,
			Type: 0,
			Name: "Add Chance to Resist Spell Effect:",
		},
		{
			ID:   379,
			Type: 0,
			Name: "Directional Short Range Teleport",
		},
		{
			ID:   380,
			Type: 0,
			Name: "Knockdown",
		},
		{
			ID:   381,
			Type: 0,
			Name: "Knock Toward Caster (Not Used)",
		},
		{
			ID:   382,
			Type: 0,
			Name: "Negate Spell Effect:",
		},
		{
			ID:   383,
			Type: 0,
			Name: "Add Chance of Spellcasting to Proc:",
		},
		{
			ID:   384,
			Type: 0,
			Name: "Leap",
		},
		{
			ID:   385,
			Type: 0,
			Name: "Limit - Group Spells Only",
		},
		{
			ID:   386,
			Type: 0,
			Name: "Cast New Spell on Curer when Cured:",
		},
		{
			ID:   387,
			Type: 0,
			Name: "Cast New Spell when Cured:",
		},
		{
			ID:   388,
			Type: 0,
			Name: "Summon Corpse Zone (Not Used)",
		},
		{
			ID:   389,
			Type: 0,
			Name: "Reactivate All Spell Gems in Cooldown",
		},
		{
			ID:   390,
			Type: 0,
			Name: "Timer Lockout (Not Used)",
		},
		{
			ID:   391,
			Type: 1,
			Name: "Melee Vulnerability",
		},
		{
			ID:   392,
			Type: 1,
			Name: "Healing Amount V3",
		},
		{
			ID:   393,
			Type: 1,
			Name: "Incoming Healing",
		},
		{
			ID:   394,
			Type: 1,
			Name: "Incoming Healing",
		},
		{
			ID:   395,
			Type: 1,
			Name: "Chance of Critical Incoming Heals",
		},
		{
			ID:   396,
			Type: 1,
			Name: "Critical Healing Amount",
		},
		{
			ID:   397,
			Type: 1,
			Name: "Pet's Melee Mitigation",
		},
		{
			ID:   398,
			Type: 1,
			Name: "Swarm Pet Duration",
		},
		{
			ID:   399,
			Type: 0,
			Name: "Add Chance to Double-Cast Spells",
		},
		{
			ID:   400,
			Type: 1,
			Name: "Current HP of Group from Own Mana Pool",
		},
		{
			ID:   401,
			Type: 1,
			Name: "Current HP from Own Mana Pool",
		},
		{
			ID:   402,
			Type: 1,
			Name: "Current HP from Own Endurance Pool",
		},
		{
			ID:   403,
			Type: 0,
			Name: "Limit - Spell Class:",
		},
		{
			ID:   404,
			Type: 0,
			Name: "Limit - Spell Subclass:",
		},
		{
			ID:   405,
			Type: 0,
			Name: "Add Chance to Block with Two-Handed Blunt Weapons (Staves)",
		},
		{
			ID:   406,
			Type: 0,
			Name: "Cast New Spell when Fade from NumHits Depleted:",
		},
		{
			ID:   407,
			Type: 0,
			Name: "Cast New Spell when Focus Effect Is Applied:",
		},
		{
			ID:   408,
			Type: 0,
			Name: "Limit - Maximum Percent of HP",
		},
		{
			ID:   409,
			Type: 0,
			Name: "Limit - Maximum Percent of Mana",
		},
		{
			ID:   410,
			Type: 0,
			Name: "Limit - Maximum Percent of Endurance",
		},
		{
			ID:   411,
			Type: 0,
			Name: "Limit to Class:",
		},
		{
			ID:   412,
			Type: 0,
			Name: "Limit to Race:",
		},
		{
			ID:   413,
			Type: 1,
			Name: "Base Power of Skills and Songs",
		},
		{
			ID:   414,
			Type: 0,
			Name: "Limit - Casting Skill:",
		},
		{
			ID:   415,
			Type: 1,
			Name: "FF Item Class (Not Used)",
		},
		{
			ID:   416,
			Type: 1,
			Name: "Armor Class V2",
		},
		{
			ID:   417,
			Type: 1,
			Name: "Mana Regeneration V2",
		},
		{
			ID:   418,
			Type: 1,
			Name: "Skill Damage V2",
		},
		{
			ID:   419,
			Type: 0,
			Name: "Add Chance of Melee Proc:",
		},
		{
			ID:   420,
			Type: 1,
			Name: "NumHits for All Spells",
		},
		{
			ID:   421,
			Type: 1,
			Name: "NumHits for Spell:",
		},
		{
			ID:   422,
			Type: 0,
			Name: "Limit - Minimum NumHits:",
		},
		{
			ID:   423,
			Type: 0,
			Name: "Limit - NumHits Type:",
		},
		{
			ID:   424,
			Type: 1,
			Name: "Gravitation to/from Mob",
		},
		{
			ID:   425,
			Type: 0,
			Name: "Display (Not Used)",
		},
		{
			ID:   426,
			Type: 1,
			Name: "Extended Targets (Not Used)",
		},
		{
			ID:   427,
			Type: 0,
			Name: "Add Chance to Proc with Skill:",
		},
		{
			ID:   428,
			Type: 0,
			Name: "Limit - Skill:",
		},
		{
			ID:   429,
			Type: 0,
			Name: "Add Chance to Proc with Skull Success:",
		},
		{
			ID:   430,
			Type: 0,
			Name: "Post Effect (Not Used)",
		},
		{
			ID:   431,
			Type: 0,
			Name: "Post Effect Data (Not Used)",
		},
		{
			ID:   432,
			Type: 0,
			Name: "Expand Max Active Trophy Ben (Not Used)",
		},
		{
			ID:   433,
			Type: 1,
			Name: "Critical DoT Chance, Decay",
		},
		{
			ID:   434,
			Type: 1,
			Name: "Critical Heal Chance, Decay",
		},
		{
			ID:   435,
			Type: 1,
			Name: "Critical Heal over Time Chance, Decay",
		},
		{
			ID:   436,
			Type: 0,
			Name: "Beneficial Countdown Hold (Not Used)",
		},
		{
			ID:   437,
			Type: 0,
			Name: "Teleport to Anchor (Not Yet Used)",
		},
		{
			ID:   438,
			Type: 0,
			Name: "Translocate to Anchor (Not Yet Used)",
		},
		{
			ID:   439,
			Type: 0,
			Name: "Add Chance to Assassinate Backstab for 32K Damage",
		},
		{
			ID:   440,
			Type: 1,
			Name: "Maximum Permitted Level for Finishing Blow",
		},
		{
			ID:   441,
			Type: 0,
			Name: "Limit - While within Distance from Buff Location",
		},
		{
			ID:   442,
			Type: 0,
			Name: "Trigger New Spell when Target",
		},
		{
			ID:   443,
			Type: 0,
			Name: "Trigger New Spell when Caster",
		},
		{
			ID:   444,
			Type: 0,
			Name: "Lock Aggro on Caster, Lowering Others' Hate by Percentage",
		},
		{
			ID:   445,
			Type: 1,
			Name: "Mercenary Slots (Not Used)",
		},
		{
			ID:   446,
			Type: 0,
			Name: "Limit - Stacking Blocker A",
		},
		{
			ID:   447,
			Type: 0,
			Name: "Limit - Stacking Blocker B",
		},
		{
			ID:   448,
			Type: 0,
			Name: "Limit - Stacking Blocker C",
		},
		{
			ID:   449,
			Type: 0,
			Name: "Limit - Stacking Blocker D",
		},
		{
			ID:   450,
			Type: 1,
			Name: "DoT Damage Mitigation",
		},
		{
			ID:   451,
			Type: 1,
			Name: "Melee Damage Mitigation for Hits Over",
		},
		{
			ID:   452,
			Type: 1,
			Name: "Spell Damage Mitigation for Hits Over",
		},
		{
			ID:   453,
			Type: 0,
			Name: "Trigger on a Melee Hit Over",
		},
		{
			ID:   454,
			Type: 0,
			Name: "Trigger on a Spell Hit Over",
		},
		{
			ID:   455,
			Type: 1,
			Name: "Total Hate Amount",
		},
		{
			ID:   456,
			Type: 1,
			Name: "Total Hate Amount",
		},
		{
			ID:   457,
			Type: 1,
			Name: "HP/Mana/Endurance from Spell Damage",
		},
		{
			ID:   458,
			Type: 1,
			Name: "Faction Gain/Loss Modifier",
		},
		{
			ID:   459,
			Type: 1,
			Name: "Damage Modifier V2 for Skill",
		},
	}
}
