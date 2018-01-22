package model

import ()

// Aa represents the root objects of Alternate Abilities
// swagger:response
type Aa struct {

	//aa_ability
	ID              int64  `json:"ID" db:"id"`                            // `id` int(10) unsigned NOT NULL,
	Name            string `json:"name" db:"name"`                        // `name` text NOT NULL,
	Category        int64  `json:"category" db:"category"`                // `category` int(10) NOT NULL DEFAULT '-1',
	Classes         int64  `json:"classes" db:"classes"`                  // `classes` int(10) NOT NULL DEFAULT '131070',
	Races           int64  `json:"races" db:"races"`                      // `races` int(10) NOT NULL DEFAULT '65535',
	DrakkinHeritage int64  `json:"drakkinHeritage" db:"drakkin_heritage"` // `drakkin_heritage` int(10) NOT NULL DEFAULT '127',
	Deities         int64  `json:"deities" db:"deities"`                  // `deities` int(10) NOT NULL DEFAULT '131071',
	Status          int64  `json:"status" db:"status"`                    // `status` int(10) NOT NULL DEFAULT '0',
	Type            int64  `json:"type" db:"type"`                        // `type` int(10) NOT NULL DEFAULT '0',
	Charges         int64  `json:"charges" db:"charges"`                  // `charges` int(11) NOT NULL DEFAULT '0',
	GrantOnly       int64  `json:"grantOnly" db:"grant_only"`             // `grant_only` tinyint(4) NOT NULL DEFAULT '0',
	FirstRankID     int64  `json:"firstRankID" db:"first_rank_id"`        // `first_rank_id` int(10) NOT NULL DEFAULT '-1',
	Enabled         int64  `json:"enabled" db:"enabled"`                  // `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
}
