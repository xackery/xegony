package cases

/*

//SlotList returns a list of all slots
func (c *ItemRepository) SlotList(item *model.Item) string {

	slots := ""
	if item.Slots&1 == 1 {
		slots += "CHARM "
	}
	if item.Slots&4 == 4 {
		slots += "HEAD "
	}
	if item.Slots&8 == 8 {
		slots += "FACE "
	}
	if item.Slots&18 == 18 {
		slots += "EARS "
	}
	if item.Slots&32 == 32 {
		slots += "NECK "
	}
	if item.Slots&64 == 64 {
		slots += "SHOULDER "
	}
	if item.Slots&128 == 128 {
		slots += "ARMS "
	}
	if item.Slots&256 == 256 {
		slots += "BACK "
	}
	if item.Slots&1536 == 1536 {
		slots += "BRACERS "
	}
	if item.Slots&2048 == 2048 {
		slots += "RANGE "
	}
	if item.Slots&4096 == 4096 {
		slots += "HANDS "
	}
	if item.Slots&8192 == 8192 {
		slots += "PRIMARY "
	}
	if item.Slots&16384 == 16384 {
		slots += "SECONDARY "
	}
	if item.Slots&98304 == 98304 {
		slots += "RINGS "
	}
	if item.Slots&131072 == 131072 {
		slots += "CHEST "
	}
	if item.Slots&262144 == 262144 {
		slots += "LEGS "
	}
	if item.Slots&524288 == 524288 {
		slots += "FEET "
	}
	if item.Slots&1048576 == 1048576 {
		slots += "WAIST "
	}
	if item.Slots&2097152 == 2097152 {
		slots += "AMMO "
	}
	if item.Slots&4194304 == 4194304 {
		slots += "POWERSOURCE "
	}
	if len(slots) > 0 {
		slots = slots[0 : len(slots)-1]
	}
	return slots
}
*/
/*
package cases

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

//SlotsFirstName returns a brief version of slots supported of an item
func (c *ItemRepository) SlotsFirstName() string {
	switch {
	case c.Slots&8192 == 8192:
		if c.Slots > 8192 {
			return "Primary+"
		}
		return "Primary"
	case c.Slots&16384 == 16384:
		if c.Slots > 16384 {
			return "Secondary+"
		}
		return "Secondary"
	case c.Slots&2048 == 2048:
		if c.Slots > 2048 {
			return "Range+"
		}
		return "Range"
	case c.Slots&1 == 1:
		if c.Slots > 1 {
			return "Charm+"
		}
		return "Charm"
	case c.Slots&4 == 4:
		if c.Slots > 4 {
			return "Head+"
		}
		return "Head"
	case c.Slots&8 == 8:
		if c.Slots > 8 {
			return "Face+"
		}
		return "Face"
	case c.Slots&18 == 18:
		if c.Slots > 18 {
			return "Ears+"
		}
		return "Ears"
	case c.Slots&32 == 32:
		if c.Slots > 32 {
			return "Neck+"
		}
		return "Neck"
	case c.Slots&64 == 64:
		if c.Slots > 64 {
			return "Shoulder+"
		}
		return "Shoulder"
	case c.Slots&128 == 128:
		if c.Slots > 128 {
			return "Arms+"
		}
		return "Arms"
	case c.Slots&256 == 256:
		if c.Slots > 256 {
			return "Back+"
		}
		return "Back"
	case c.Slots&1536 == 1536:
		if c.Slots > 1536 {
			return "Bracers+"
		}
		return "Bracers"
	case c.Slots&4096 == 4096:
		if c.Slots > 4096 {
			return "Hands+"
		}
		return "Hands"
	case c.Slots&98304 == 98304:
		if c.Slots > 98304 {
			return "Rings+"
		}
		return "Rings"
	case c.Slots&131072 == 131072:
		if c.Slots > 131072 {
			return "Chest+"
		}
		return "Chest"
	case c.Slots&262144 == 262144:
		if c.Slots > 262144 {
			return "Legs+"
		}
		return "Legs"
	case c.Slots&524288 == 524288:
		if c.Slots > 524288 {
			return "Feet+"
		}
		return "Feet"
	case c.Slots&1048576 == 1048576:
		if c.Slots > 1048576 {
			return "Waist+"
		}
		return "Waist"
	case c.Slots&2097152 == 2097152:
		if c.Slots > 2097152 {
			return "Ammo+"
		}
		return "Ammo"
	case c.Slots&4194304 == 4194304:
		if c.Slots > 4194304 {
			return "Powersource+"
		}
		return "Powersource"
	}
	return "None"
}

//SlotName returns the human readable version of a slot
func (c *ItemRepository) SlotName() string {
	s := c.SlotID
	switch {
	case s == 0:
		return "Charm"
	case s == 1:
		return "Left Ear"
	case s == 2:
		return "Head"
	case s == 3:
		return "Face"
	case s == 4:
		return "Right Ear"
	case s == 5:
		return "Neck"
	case s == 6:
		return "Shoulder"
	case s == 7:
		return "Arms"
	case s == 8:
		return "Back"
	case s == 9:
		return "Left Bracer"
	case s == 10:
		return "Right Bracer"
	case s == 11:
		return "Range"
	case s == 12:
		return "Hands"
	case s == 13:
		return "Primary"
	case s == 14:
		return "Secondary"
	case s == 15:
		return "Left Ring"
	case s == 16:
		return "Right Ring"
	case s == 17:
		return "Chest"
	case s == 18:
		return "Legs"
	case s == 19:
		return "Feet"
	case s == 20:
		return "Waist"
	case s == 21:
		return "Ammo"
	case s == 22:
		return "TopLeft Inventory"
	case s <= 271 && s >= 262:
		return "TopLeft Bag"
	case s == 23:
		return "TopRight Inventory"
	case s <= 281 && s >= 272:
		return "TopRight Bag"
	case s == 24:
		return "TopLeft, One Down Inventory"
	case s <= 291 && s >= 282:
		return "TopLeft, One Down Bag"
	case s == 25:
		return "TopRight, One Down Inventory"
	case s <= 301 && s >= 292:
		return "TopRight, One Down Bag"
	case s == 26:
		return "BottomLeft, Two Up Inventory"
	case s <= 311 && s >= 302:
		return "BottomLeft, Two Up Bag"
	case s == 27:
		return "BottomRight, Two Up Inventory"
	case s <= 321 && s >= 312:
		return "BottomRight, Two Up Bag"
	case s == 28:
		return "BottomLeft Inventory"
	case s <= 331 && s >= 322:
		return "BottomLeft Bag"
	case s == 29:
		return "BottomRight Inventory"
	case s <= 341 && s >= 332:
		return "BottomRight Bag"
	case s == 30:
		return "Cursor"
	case s >= 2000 && s <= 2271:
		return "Bank"
	case s >= 400 && s <= 404:
		return "Tribute"
	case s >= 2500 && s <= 2551:
		return "Shared Bank"
	default:
		return fmt.Sprintf("Unknown (%d)", c.SlotID)
	}
}
*/
