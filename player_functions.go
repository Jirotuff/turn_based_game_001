package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// The player type holds stats and limits
type player struct {
	max_health       int
	max_skill_points int
	name             string
	special          int
	exp              int
	lv               int
	health           int // player health
	skill_points     int // points used to cast magic spells
	strength         int // increases physical damage
	intelligence     int // increases magical damage
	agility          int // increases chance to dodge
	endurance        int // reduces damage taken
	social           int // reduces shop prices
}

// Dario
var Dario = player{
	max_health:       110,
	max_skill_points: 50,
	name:             name_1,
	special:          0,
	exp:              0,
	lv:               1,
	health:           110,
	skill_points:     75,
	strength:         12,
	intelligence:     12,
	agility:          10,
	endurance:        10,
	social:           10,
}

var Pilgrim = player{
	max_health:       120,
	max_skill_points: 50,
	name:             name_2,
	special:          0,
	exp:              0,
	lv:               1,
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
	exp:              0,
	lv:               1,
	health:           90,
	skill_points:     80,
	strength:         10,
	intelligence:     10,
	agility:          14,
	endurance:        8,
	social:           10,
}

var Jessy = player{
	max_health:       80,
	max_skill_points: 90,
	name:             name_4,
	special:          0,
	exp:              0,
	lv:               1,
	health:           100,
	skill_points:     90,
	strength:         8,
	intelligence:     14,
	agility:          12,
	endurance:        10,
	social:           12,
}

// The block below currently only contains the player's turn, might expend upon this later

func (p *player) Player_turn(e *enemy) {
	clear_screen()

	fmt.Println("Friend")
	Dario.show_status()
	Pilgrim.show_status()
	Fie.show_status()
	Jessy.show_status()
	fmt.Println("")
	fmt.Println("\nFoe")
	e.show_status()

	if p.special >= 3 {
		{
			fmt.Println(p.name, "\033[95mfeels a strange power welling up inside... (type 'special' to unleash it)\033[0m")
		}
	}
	fmt.Println(p.name, "'s turn")
	fmt.Println("\nWhat's your move?")
	fmt.Println("\n>> (st)rike\t\t\t> Use your basic weapon\t")
	fmt.Println(">> (h)eal	| 10 SP\t\t> Use an healing spell\t")
	fmt.Println(">> (f)orce | 20 SP\t\t> High citical chance attack")
	fmt.Println(">> (so)ul \t\t\t> Regenerates some SP")
	fmt.Println(">> (It)em \t\t\t> Use an item")

	fmt.Scanln(&user_input)

	switch strings.ToLower(user_input) {

	case "item", "i", "it", "ite":
		p.Use_item()

	case "strike", "st", "str", "stri", "strik":
		p.Player_skill_strike(e)

	case "heal", "h", "he", "hea":
		p.Player_skill_heal()

	case "force", "f", "fo", "for", "forc":
		p.Player_skill_force(e)

	case "soul", "so", "sou":
		p.Player_skill_soul()

	case "kill", "k", "ki", "kil":
		p.Player_skill_kill(e)

	case "special", "sp", "spe", "spec", "speci", "specia":
		if p.special > 2 {
			p.special = 0
			p.Player_skill_special(e)
		} else {
			fmt.Println("You dont have the energy for this move")
		}
	default:
		p.Player_turn(e)

	}
	p.Normalize_stats()

	fmt.Println("press Enter to continue")
	fmt.Scanln(&user_input)
}

// The block below is for storing things that need to be checked (and changed?) in regards to the players

// checks if the player has enough exp to level up, if so increases level and stats
func (p *player) Level_check() {
	if p.exp >= 100 && p.lv < 2 || p.exp >= 500 && p.lv < 3 || p.exp >= 1500 && p.lv < 4 || p.exp >= 3000 && p.lv < 5 || p.exp >= 5000 && p.lv < 6 || p.exp >= 10000 && p.lv < 7 {
		p.lv++
		p.max_health += 10
		p.max_skill_points += 5
		p.health = p.max_health
		p.skill_points = p.max_skill_points
		fmt.Println("\n", p.name, ":\033[92m Level up!!\033[0m")
		fmt.Printf("\nMax HP: %d, Max SP: %d\n", p.max_health, p.max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("(St)rength:", p.strength, "\n(In)telligence: ", p.intelligence, "\n(Ag)ility: ", p.agility, "\n(En)durance: ", p.endurance, "\n(So)cial: ", p.social, "")
		p.level_up()
	}
}
func (p *player) level_up() {
	user_input = ""
	fmt.Scanln(&user_input)

	switch user_input {

	case "strength", "st", "str", "stre":
		p.strength += 2

	case "intelligence", "in", "int", "inte":
		p.intelligence += 2

	case "agility", "ag", "agi", "agil":
		p.agility += 2

	case "endurance", "en", "end", "endu":
		p.endurance += 2

	case "social", "so", "soc", "soci":
		p.social += 2

	default:
		p.level_up()
	}
}

// checks game over
func (p *player) Check_player_life() {
	if Dario.health <= 0 {
		fmt.Println("Your hero has been killed!")
		fmt.Println("\nGold:", gold, "Player level:", Dario.lv)
		fmt.Println("\nPress Enter to quit")

		fmt.Scanln("")
		fmt.Scanf("%s", &user_input)
		if user_input == "exit" {
			os.Exit(0)
		} else {
			os.Exit(0)
		}

	}
}

func (p *player) Normalize_stats() {
	if p.health > p.max_health {
		p.health = p.max_health
	}
	if p.skill_points > p.max_skill_points {
		p.skill_points = p.max_skill_points
	}
	if p.special > 3 {
		p.special = 3
	}
}

// The block below is the place to store all skills that the players can use

func (p *player) Player_skill_kill(e *enemy) {
	damage := rand.Intn(20) + 5 + Dario.strength + 9999
	critical_damage := rand.Intn(20) + 30 + Dario.strength + 9999

	if rand.Intn(11) == 9 { //Critical hit chance
		e.health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

func (p *player) Player_skill_strike(e *enemy) {

	damage := rand.Intn(20) + 5 + p.strength
	critical_damage := rand.Intn(20) + 30 + p.strength

	p.special += 1

	if rand.Intn(11) == 9 { //Critical hit chance
		e.health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

func (p *player) Player_skill_soul() {

	p.skill_points += 25
}

func (p *player) Player_skill_force(e *enemy) {
	damage := rand.Intn(5) + 20 + p.intelligence
	critical_damage := rand.Intn(20) + 30 + p.intelligence

	if p.skill_points >= 20 {

		p.skill_points -= 20

		p.special += 1

		if rand.Intn(3) == 2 { //Critical hit chance

			e.health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			user_input = ""
		} else {
			e.health -= damage
			fmt.Println(damage, "DMG")
			user_input = ""
		}
	} else {
		fmt.Println("You tried to cast force... but you dont have enough SP!")
		user_input = ""
	}
}

func (p *player) Player_skill_heal() {
	if p.skill_points >= 10 {

		p.skill_points = -10

		heal := rand.Intn(20) + 5 + p.intelligence //amount healed

		Dario.health += heal
		if Jessy.health > 0 {
			Jessy.health += heal
		}
		if Pilgrim.health > 0 {
			Pilgrim.health += heal
		}
		if Fie.health > 0 {
			Fie.health += heal
		}
		Dario.Normalize_stats()
		Fie.Normalize_stats()
		Pilgrim.Normalize_stats()
		Jessy.Normalize_stats()

		user_input = ""
	}
}

func (p *player) Player_skill_special(e *enemy) {
	damage := 70 + p.strength
	critical_damage := rand.Intn(20) + 75 + p.strength

	if rand.Intn(11) == 9 { //Critical hit chance
		e.health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

func (p *player) Use_item() {
	fmt.Println(inventory, "\n\nWhat item to use...\n\nType (ba)ck to return")
	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "potion", "p", "po", "pot", "poti":
		if contains_string(inventory, "potion") {
			remove_item(inventory, "potion")
			fmt.Println("You have used a potion...")
			p.health += 30
		} else {
			fmt.Println("You do not have a potion")
		}
	case "fire_gem", "fi", "fir", "fire", "fire_":
		if contains_string(inventory, "fire_gem") {
			remove_item(inventory, "fire_gem")
			fmt.Println(p.name, ": has used a fire_gem...")
			fmt.Println("... DEBUG DEBUG DEBUG ...")
		} else {
			fmt.Println("You do not have a fire_gem")
		}
	case "revival_bead", "re", "rev", "revi", "reviv":
		if contains_string(inventory, "revival_bead") {
			remove_item(inventory, "revival_bead")
			fmt.Println("You have used a revival_bead...")
			if Pilgrim.health < 1 {
				Pilgrim.health = 50
			}
			if Fie.health < 1 {
				Fie.health = 50
			}
			if Jessy.health < 1 {
				Jessy.health = 50
			}
		} else {
			fmt.Println("you do not have a revival bead")
		}
	case "back", "ba", "bac":

	default:
		fmt.Println("Input invalid")
		p.Use_item()
	}
}
