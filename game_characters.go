package main

// player struct
type player struct {
	max_health       int
	max_skill_points int
	name             string
	special          int
	inventory        []string
	exp              int
	lv               int
	gold             int
	health           int // player health
	skill_points     int // points used to cast magic spells
	strength         int // increases physical damage
	intelligence     int // increases magical damage
	agility          int // increases chance to dodge
	endurance        int // reduces damage taken
	social           int // reduces shop prices
}

// players

var Dario = player{
	max_health:       100,
	max_skill_points: 50,
	name:             name_1,
	special:          0,
	inventory:        []string{},
	exp:              0,
	lv:               0,
	gold:             50,
	health:           100,
	skill_points:     75,
	strength:         11,
	intelligence:     11,
	agility:          11,
	endurance:        11,
	social:           11,
}

var Pilgrim = player{
	max_health:       120,
	max_skill_points: 50,
	name:             name_2,
	special:          0,
	inventory:        []string{},
	exp:              0,
	lv:               0,
	gold:             50,
	health:           120,
	skill_points:     70,
	strength:         10,
	intelligence:     8,
	agility:          8,
	endurance:        14,
	social:           10,
}

var Fie = player{
	max_health:       90,
	max_skill_points: 50,
	name:             name_3,
	special:          0,
	inventory:        []string{},
	exp:              0,
	lv:               0,
	gold:             50,
	health:           90,
	skill_points:     80,
	strength:         10,
	intelligence:     10,
	agility:          14,
	endurance:        8,
	social:           10,
}

var Rean = player{
	max_health:       100,
	max_skill_points: 50,
	name:             name_4,
	special:          0,
	inventory:        []string{},
	exp:              0,
	lv:               0,
	gold:             50,
	health:           100,
	skill_points:     75,
	strength:         12,
	intelligence:     10,
	agility:          12,
	endurance:        10,
	social:           12,
}

// enemy struct

type enemy struct {
	input            int //enemy input
	health           int //enemy health
	skill_points     int
	max_skill_points int
	max_health       int
}

//enemies

var enemy_1 = enemy{
	input:            0,
	health:           100,
	skill_points:     100,
	max_skill_points: 80,
	max_health:       100,
}

var enemy_2 = enemy{
	input:            0,
	health:           100,
	skill_points:     100,
	max_skill_points: 80,
	max_health:       100,
}
