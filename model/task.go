package model

import (
	"fmt"
)

//Task is an everquest entry, grouping the Activities entries
// swagger:response
type Task struct {
	ID           int64  `json:"id" db:"id"`                     //`id` int(11) unsigned NOT NULL DEFAULT '0',
	Duration     int64  `json:"duration" db:"duration"`         //`duration` int(11) unsigned NOT NULL DEFAULT '0',
	Title        string `json:"title" db:"title"`               //`title` varchar(100) NOT NULL DEFAULT '',
	Description  string `json:"description" db:"description"`   //`description` text NOT NULL,
	Reward       string `json:"reward" db:"reward"`             //`reward` varchar(64) NOT NULL DEFAULT '',
	Rewardid     int64  `json:"rewardid" db:"rewardid"`         //`rewardid` int(11) unsigned NOT NULL DEFAULT '0',
	Cashreward   int64  `json:"cashreward" db:"cashreward"`     //`cashreward` int(11) unsigned NOT NULL DEFAULT '0',
	Xpreward     int64  `json:"xpreward" db:"xpreward"`         //`xpreward` int(10) NOT NULL DEFAULT '0',
	Rewardmethod int64  `json:"rewardmethod" db:"rewardmethod"` //`rewardmethod` tinyint(3) unsigned NOT NULL DEFAULT '2',
	Startzone    int64  `json:"startzone" db:"startzone"`       //`startzone` int(11) NOT NULL DEFAULT '0',
	Minlevel     int64  `json:"minlevel" db:"minlevel"`         //`minlevel` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Maxlevel     int64  `json:"maxlevel" db:"maxlevel"`         //`maxlevel` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Repeatable   int64  `json:"repeatable" db:"repeatable"`     //`repeatable` tinyint(1) unsigned NOT NULL DEFAULT '1',
}

//CashRewardName breaks down money into human readable money
func (c *Task) CashRewardName() string {
	return CashName(c.Cashreward)
}

//RewardName shows a list of reward types
func (c *Task) RewardName() string {
	switch c.Rewardmethod {
	case 0:
		return c.Reward
	case 1:
		return fmt.Sprintf("GoalList %d", c.Rewardid)
	case 2:
		return "Perl"
	}
	return "Unknown"
}

//RewardMethodName displays reward method type
func (c *Task) RewardMethodName() string {
	switch c.Rewardmethod {
	case 0:
		return "Item"
	case 1:
		return "Items"
	case 2:
		return "Perl"
	}
	return "Unknown"
}

//RepeatableName returns Yes or No
func (c *Task) RepeatableName() string {
	if c.Repeatable == 0 {
		return "No"
	}
	return "Yes"
}
