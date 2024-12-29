package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// The player type holds stats and limits
type player struct {
	Max_health       int
	Max_skill_points int
	Name             string
	Special          int
	Exp              int
	Lv               int
	Health           int // player Health
	Skill_points     int // points used to cast magic spells
	Strength         int // increases physical damage
	Intelligence     int // increases magical damage
	Agility          int // increases chance to dodge
	Endurance        int // reduces damage taken
}

// Player
var Player = player{
	Max_health:       110,
	Max_skill_points: 75,
	Name:             "",
	Special:          0,
	Exp:              0,
	Lv:               1,
	Health:           110,
	Skill_points:     75,
	Strength:         12,
	Intelligence:     12,
	Agility:          10,
	Endurance:        10,
}

var Pilgrim = player{
	Max_health:       120,
	Max_skill_points: 70,
	Name:             name_2,
	Special:          0,
	Exp:              0,
	Lv:               1,
	Health:           120,
	Skill_points:     70,
	Strength:         10,
	Intelligence:     8,
	Agility:          8,
	Endurance:        14,
}

var Fie = player{
	Max_health:       90,
	Max_skill_points: 80,
	Name:             name_3,
	Special:          0,
	Exp:              0,
	Lv:               1,
	Health:           90,
	Skill_points:     80,
	Strength:         10,
	Intelligence:     10,
	Agility:          14,
	Endurance:        8,
}

var Jessy = player{
	Max_health:       80,
	Max_skill_points: 90,
	Name:             name_4,
	Special:          0,
	Exp:              0,
	Lv:               1,
	Health:           100,
	Skill_points:     90,
	Strength:         8,
	Intelligence:     14,
	Agility:          12,
	Endurance:        10,
}

// The block below currently only contains the player's turn, might expend upon this later

func (p *player) Player_turn(e *enemy) {
	clear_screen()

	fmt.Println("Friendly party:")
	Player.show_status()
	Pilgrim.show_status()
	Fie.show_status()
	Jessy.show_status()
	fmt.Println("")
	fmt.Println("\nFoe:")
	e.show_status()

	if p.Special >= 3 {
		{
			fmt.Println(p.Name, "\033[95mfeels a strange power welling up inside... (type 'Special' to unleash it)\033[0m")
		}
	}
	fmt.Println(p.Name, "'s turn")
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

	case "Special", "sp", "spe", "spec", "speci", "specia":
		if p.Special > 2 {
			p.Special = 0
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

// checks if the player has enough Exp to level up, if so increases level and stats
func (p *player) Level_check() {
	if p.Exp >= 100 && p.Lv < 2 || p.Exp >= 500 && p.Lv < 3 || p.Exp >= 1500 && p.Lv < 4 || p.Exp >= 3000 && p.Lv < 5 || p.Exp >= 5000 && p.Lv < 6 || p.Exp >= 10000 && p.Lv < 7 {
		p.Lv++
		p.Max_health += 10
		p.Max_skill_points += 5
		p.Health = p.Max_health
		p.Skill_points = p.Max_skill_points
		fmt.Println("\n", p.Name, ":\033[92m Level up!!\033[0m")
		fmt.Printf("\nMax HP: %d, Max SP: %d\n", p.Max_health, p.Max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("(St)rength:", p.Strength, "\n(In)telligence: ", p.Intelligence, "\n(Ag)ility: ", p.Agility, "\n(En)durance: ", p.Endurance)
		p.level_up()
	}
}
func (p *player) level_up() {
	user_input = ""
	fmt.Scanln(&user_input)

	switch strings.ToLower(user_input) {

	case "Strength", "st", "str", "stre":
		p.Strength += 2

	case "Intelligence", "in", "int", "inte":
		p.Intelligence += 2

	case "Agility", "ag", "agi", "agil":
		p.Agility += 2

	case "Endurance", "en", "end", "endu":
		p.Endurance += 2

	default:
		p.level_up()
	}
}

// checks game over
func (p *player) Check_player_life() {
	if Player.Health <= 0 {
		fmt.Println("Your hero has been killed!")
		fmt.Println("\nGold: ", gold, "Player level: ", Player.Lv)
		fmt.Println("\nPress Enter to quit")

		fmt.Scanln("")
		fmt.Scanf("%s", &user_input)
		os.Exit(0)

	}
}

func (p *player) Normalize_stats() {
	if p.Health > p.Max_health {
		p.Health = p.Max_health
	}
	if p.Skill_points > p.Max_skill_points {
		p.Skill_points = p.Max_skill_points
	}
	if p.Special > 3 {
		p.Special = 3
	}
}

// The block below is the place to store all skills that the players can use

func (p *player) Player_skill_kill(e *enemy) {
	damage := rand.Intn(20) + 5 + Player.Strength + 9999
	critical_damage := rand.Intn(20) + 30 + Player.Strength + 9999

	if rand.Intn(11) == 9 { //Critical hit chance
		e.Health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.Health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

func (p *player) Player_skill_strike(e *enemy) {

	damage := rand.Intn(20) + 5 + p.Strength + equipment_phys_offense
	critical_damage := rand.Intn(20) + 30 + p.Strength + equipment_phys_offense

	p.Special += 1

	if rand.Intn(11) == 9 { //Critical hit chance
		e.Health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.Health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

func (p *player) Player_skill_soul() {
	var skill_points_added = rand.Intn(10) + 25

	p.Skill_points += skill_points_added

	fmt.Println("Skill points revitalised: ", skill_points_added)
}

func (p *player) Player_skill_force(e *enemy) {
	damage := rand.Intn(5) + 20 + p.Intelligence
	critical_damage := rand.Intn(20) + 30 + p.Intelligence

	if p.Skill_points >= 20 {

		p.Skill_points -= 20

		p.Special += 1

		if rand.Intn(3) == 2 { //Critical hit chance

			e.Health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			user_input = ""
		} else {
			e.Health -= damage
			fmt.Println(damage, "DMG")
			user_input = ""
		}
	} else {
		fmt.Println("You tried to cast force... but you dont have enough SP!")
		user_input = ""
	}
}

func (p *player) Player_skill_heal() {
	if p.Skill_points >= 10 {

		p.Skill_points -= 10

		heal := rand.Intn(20) + 5 + p.Intelligence //amount healed

		Player.Health += heal

		if Jessy.Health > 0 {
			Jessy.Health += heal
		}
		if Pilgrim.Health > 0 {
			Pilgrim.Health += heal
		}
		if Fie.Health > 0 {
			Fie.Health += heal
		}
		Player.Normalize_stats()
		Fie.Normalize_stats()
		Pilgrim.Normalize_stats()
		Jessy.Normalize_stats()

		fmt.Println("\nParty has healed: ", heal)

		user_input = ""
	}
}

func (p *player) Player_skill_special(e *enemy) {
	damage := 70 + p.Strength + equipment_phys_offense
	critical_damage := rand.Intn(20) + 75 + p.Strength + equipment_phys_offense

	if rand.Intn(11) == 9 { //Critical hit chance
		e.Health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		e.Health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

func (p *player) Use_item() {
	var heal = 30

	fmt.Println(inventory, "\n\nWhat item to use...\n\nType (ba)ck to return")
	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "potion", "p", "po", "pot", "poti":
		if contains_string(inventory, "potion") {
			remove_item(inventory, "potion")
			p.Health += heal
			fmt.Println("You have used a potion, healed: ", heal)
		} else {
			fmt.Println("You do not have a potion")
		}

	case "revival_bead", "re", "rev", "revi", "reviv":
		if contains_string(inventory, "revival_bead") {
			remove_item(inventory, "revival_bead")
			fmt.Println("You have used a revival_bead...")
			if Pilgrim.Health < 1 {
				Pilgrim.Health = 50
			}
			if Fie.Health < 1 {
				Fie.Health = 50
			}
			if Jessy.Health < 1 {
				Jessy.Health = 50
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
