package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Constants
const (
	name_1 string = "Dario"
	name_2 string = "Pilgrim"
	name_3 string = "Fie"
	name_4 string = "Jessy"
)

// Variables

var inventory []string
var gold int
var exp_gained int
var gold_earned int

// player struct
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

// players

var Dario = player{
	max_health:       250,
	max_skill_points: 50,
	name:             name_1,
	special:          0,
	exp:              0,
	lv:               1,
	health:           250,
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

// enemy struct

type enemy struct {
	name             string
	health           int //enemy health
	skill_points     int
	max_skill_points int
	max_health       int
}

//enemies

var Bandit = enemy{
	name:             "Bandit",
	health:           300,
	skill_points:     200,
	max_skill_points: 200,
	max_health:       300,
}

// User

var user_input string //player input
var victory bool = false
var display_tutorial bool = true

// Enemy
var enemy_input int //enemy input

// Start of program
func main() {

	user_input = ""
	check_victory()

	if display_tutorial {
		tutorial()
	}

	//fmt.Println(player_name, "  Health:", player_health, "SP:", player_skill_points, "Gold:", gold)

	Dario.show_status()
	Pilgrim.show_status()
	Fie.show_status()
	Jessy.show_status()

	fmt.Println("\nWhat do you want to do?")
	fmt.Println("\nbattle\t\t> finds opponent")
	fmt.Println("shop\t\t> enter the shop")
	fmt.Println("stats\t\t> show player stats")
	fmt.Println("inv\t\t> show player inventory")
	fmt.Println("exit\t\t> exits the game")
	fmt.Println("")

	for {
		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "battle", "b", "ba", "bat":
			combat()

		case "shop", "sh", "sho":
			shop()

		case "stats", "st", "sta", "stat":
			Dario.display_stats()
			Pilgrim.display_stats()
			Fie.display_stats()
			Jessy.display_stats()

			fmt.Scanln(&user_input)

			switch strings.ToLower(user_input) {

			case "back", "b", "ba", "bac":
				main()

			}

		case "inv", "i", "in":
			display_inventory()

		case "exit":
			quit()

		default:
			main()
		}
	}
}

// func save(slot1 any, data interface{}) {}

func check_victory() {
	if victory {
		victory = false

		fmt.Println("Victory!")

		exp_gained = rand.Intn(50) + 50

		Dario.exp += exp_gained
		Pilgrim.exp += exp_gained
		Fie.exp += exp_gained
		Jessy.exp += exp_gained

		gold_earned = rand.Intn(30) + 10

		fmt.Println("\nLoot:\n\nexp:", exp_gained, "\ngold: ", gold_earned)

		gold_earned = 0
		exp_gained = 0

		Dario.level_check()
		Fie.level_check()
		Pilgrim.level_check()
		Jessy.level_check()

		reset_enemy(&Bandit)

		fmt.Println("\nType any key to continue")
		fmt.Scanln(&user_input)
		if user_input == (" ") {
			main()
		} else {
			main()
		}

	}
}

// Starts the combat encounter
func combat() {
	fmt.Println("\n\nCombat started!")

	for !victory {
		Dario.check_player_life()
		Bandit.check_enemy_life()
		Dario.player_turn()
		Bandit.check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.player_turn()
		}
		Bandit.check_enemy_life()
		if Fie.health > 0 {
			Fie.player_turn()
		}
		Bandit.check_enemy_life()
		if Jessy.health > 0 {
			Jessy.player_turn()
		}
		Bandit.check_enemy_life()

		Bandit.enemy_turn()
	}
}

// Function for player turn
func (p *player) player_turn() {
	clear_screen()

	Dario.show_status()
	Pilgrim.show_status()
	Fie.show_status()
	Jessy.show_status()
	fmt.Println("")

	if p.special >= 3 {
		{
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 91, "You feel a strange power welling up inside... (type 'special' to unleash it)")
			fmt.Println(colored)
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

	switch strings.ToLower(user_input) { //gives different options to the player

	case "item", "i", "it", "ite":
		p.use_item()

	case "strike", "st", "str", "stri":
		p.player_skill_strike(&Bandit)

	case "heal", "h", "he", "hea":
		p.player_skill_heal()

	case "force", "f", "fo", "for", "forc":
		p.player_skill_force(&Bandit)

	case "soul", "so", "sou":
		p.player_skill_soul()

	case "kill", "k", "ki", "kil":
		p.player_skill_kill(&Bandit)

	case "special", "sp", "spe", "spec":
		if p.special > 2 {
			p.special = 0
			p.player_skill_special(&Bandit)
		} else {
			fmt.Println("You dont have the energy for this move")
		}
	default:
		fmt.Println("Thats a typo! lost your turn XD")
	}
	p.normalize_stats()

	//insert press to continue...
}

// Function for enemy turn
func (e *enemy) enemy_turn() {
	fmt.Println(e.name, "'s turn")
	fmt.Println("")
	enemy_input = rand.Intn(3) //gives different options to the enemy

	if e.health >= 1 {
		switch enemy_input {

		case 0:
			enemy_input = rand.Intn(3)

			switch enemy_input {

			case 0:
				e.enemy_skill_strike(&Dario)

			case 1:
				e.enemy_skill_strike(&Pilgrim)

			case 2:
				e.enemy_skill_strike(&Fie)

			case 3:
				e.enemy_skill_strike(&Jessy)
			}

		case 1:
			e.enemy_skill_heal()

		case 2:
			enemy_input = rand.Intn(3)

			switch enemy_input {

			case 0:
				e.enemy_skill_force(&Dario)

			case 1:
				e.enemy_skill_force(&Pilgrim)

			case 2:
				e.enemy_skill_force(&Fie)

			case 3:
				e.enemy_skill_force(&Jessy)
			}
		}
		e.normalize_stats_enemy()
	}
}

func reset_enemy(e *enemy) {
	e.health = e.max_health
	e.skill_points = e.max_skill_points
}

// Displays a tutorial if display_tutorial == true
func tutorial() {
	display_tutorial = false
	fmt.Println("Welcome to this game...")
	fmt.Println("\nThis is a turn based game, as the player you can type the one of the moves to execute it.")
	fmt.Println("Your goal at this moment is to acquire as much gold as possible")

}

// Checks if the player is dead
func (p *player) check_player_life() {
	if Dario.health <= 0 {
		fmt.Println("Your hero has been killed!")
		fmt.Println("\nGold:", gold, "Player level:", Dario.lv)
		fmt.Println("\nType anything to quit")

		fmt.Scanln("")
		fmt.Scanf("%s", &user_input)
		if user_input == "exit" {
			os.Exit(0)
		} else {
			os.Exit(0)
		}

	}
}

// Player skill: kill (THIS IS A TEST FEATURE, NOT MEANT FOR FINAL PRODUCT)
func (p *player) player_skill_kill(e *enemy) {
	damage := rand.Intn(20) + 5 + Dario.strength + 999
	critical_damage := rand.Intn(20) + 30 + Dario.strength + 999

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

// Player skill: strike
func (p *player) player_skill_strike(e *enemy) {

	damage := rand.Intn(20) + 5 + Dario.strength
	critical_damage := rand.Intn(20) + 30 + Dario.strength

	Dario.special += 1

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

// Player skill: soul
func (p *player) player_skill_soul() {

	p.skill_points += 25
}

// Player skill: force
func (p *player) player_skill_force(e *enemy) {
	damage := rand.Intn(5) + 20 + Jessy.intelligence
	critical_damage := rand.Intn(20) + 30 + Jessy.intelligence

	if Jessy.skill_points >= 20 {

		Jessy.skill_points -= 20

		Jessy.special += 1

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

// Player skill: heal
func (p *player) player_skill_heal() {
	if p.skill_points >= 10 {

		p.skill_points = -10

		heal := rand.Intn(20) + 5 + Jessy.intelligence //amount healed

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
		if Dario.health > Dario.max_health {
			Dario.health = Dario.max_health
		}
		if Pilgrim.health > Pilgrim.max_health {
			Pilgrim.health = Pilgrim.max_health
		}
		if Fie.health > Fie.max_health {
			Fie.health = Fie.max_health
		}
		if Jessy.health > Jessy.max_health {
			Jessy.health = Jessy.max_health
		}

		user_input = ""
	}
}

// Triggers a special move when the requirement is met
func (p *player) player_skill_special(e *enemy) {
	damage := 70 + Dario.strength
	critical_damage := rand.Intn(20) + 75 + Dario.strength

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

// Checks if the enemy is dead
func (e *enemy) check_enemy_life() {
	if e.health <= 0 {
		victory = true
		check_victory()
	}
}

// Enemy skill: strike
func (e *enemy) enemy_skill_strike(p *player) {
	fmt.Println(e.name, "used strike")
	damage := rand.Intn(10) + 50 - p.endurance
	critical_damage := rand.Intn(20) + 60 - p.endurance

	if rand.Intn(100) > p.agility {

		if rand.Intn(11) == 9 { //Critical hit chance
			p.health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		} else {
			p.health -= damage
			fmt.Println(damage, "DMG")
		}
	} else {
		fmt.Println("But it missed!")
	}
}

// Enemy skill: heal
func (e *enemy) enemy_skill_heal() {
	heal := rand.Intn(50) + 50 //amount healed
	fmt.Println(e.name, "has healed")
	e.health += heal
	fmt.Println(heal, "Healed")
}

// Enemy skill: force
func (e *enemy) enemy_skill_force(p *player) {

	damage := rand.Intn(10) + 70 - p.endurance
	critical_damage := rand.Intn(20) + 80 - p.endurance

	fmt.Println(e.name, "cast force")

	if e.skill_points >= 20 {

		e.skill_points -= 20

		if rand.Intn(100) >= p.agility {

			if rand.Intn(3) == 1 { //Critical hit chance
				p.health -= critical_damage
				fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			} else {
				p.health -= damage
				fmt.Println(damage, "DMG")
			}
		} else {
			fmt.Println("but it missed")
		}
	} else {
		fmt.Println("but nothing happened...")
		damage = 0
		p.health -= damage
		fmt.Println(damage, "DMG")
		fmt.Scanln()
	}
}

// Clear the screen in the CLI
func clear_screen() {
	cmd := exec.Command("clear") // for Unix/Linux
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Checks the exp and increases the player_lv
func (p *player) level_check() {
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
			fmt.Println("Something went wrong")
			p.level_check()
		}
	}
}

// Makes the shop work | not finished yet
func shop() {

	fmt.Println("Welcome to the shop")
	fmt.Println("\nWe have a variety of products available, please take your time choosing")
	fmt.Println("\n- potion")
	fmt.Println("- fire_gem")
	fmt.Println("- revival bead")
	fmt.Println("\nleave the shop (back)")

	for {

		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) { //gives different options to the player

		case "potion", "po", "pot", "poti":
			fmt.Println("you have bought a potion")
			inventory = append(inventory, "potion")

		case "fire_gem", "fi", "fir", "fire":
			fmt.Println("you have bought a fire_gem")
			inventory = append(inventory, "fire_gem")

		case "revival bead", "re", "rev", "revi":
			fmt.Println("you have bought a revival bead")
			inventory = append(inventory, "revival bead")

		case "back", "b", "ba", "bac":
			main()

		default:
			fmt.Println("We don't have this item...")
		}

	}
}

// Displays a player's lv, exp and stats
func (p *player) display_stats() {
	fmt.Println("\n", p.name, "lv:", p.lv)
	fmt.Println("Exp:", p.exp)
	fmt.Println("\nStrength:", p.strength)
	fmt.Println("Intelligence:", p.intelligence)
	fmt.Println("Agility:", p.agility)
	fmt.Println("Endurance:", p.endurance)
	fmt.Println("Social:", p.social)
	fmt.Println("\n[back]")

}

func (p *player) show_status() {
	if p.health > 0 {
		fmt.Println(p.name, ":\nhealth: ", p.health, "skill points: ", p.skill_points)
	} else {
		fmt.Println(p.name, ":\n\033[95m DEAD...\033[0m")
	}
}

// Displays the player's inventory
func display_inventory() {

	fmt.Println(inventory)

	fmt.Println("\nType '(b)ack' to retun to main menu")

	fmt.Scanln(&user_input)

	switch strings.ToLower(user_input) {

	case "(b)ack":
		main()

	default:
		main()
	}
}

func contains_string(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}

	}
	return false
}

func (p *player) use_item() {
	fmt.Println(inventory, "\n\nWhat item to use...\n\nType (ba)ck to return")
	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "potion", "p", "po", "pot", "poti":
		if contains_string(inventory, "potion") {
			fmt.Println("You have used a potion...")
			p.health = +30
		} else {
			fmt.Println("You do not have a potion")
		}
	case "fire_gem", "fi", "fir", "fire", "fire_":
		if contains_string(inventory, "fire_gem") {
			fmt.Println(p.name, ": has used a fire_gem...")
			fmt.Println("... DEBUG DEBUG DEBUG ...")
		}
	case "revival bead", "re", "rev", "revi", "reviv":
		if contains_string(inventory, "revival bead") {
			fmt.Println("You have used a revival bead...")

			if Pilgrim.health < 1 {
				Pilgrim.health = 50
			}
			if Fie.health < 1 {
				Fie.health = 50
			}
			if Jessy.health < 1 {
				Jessy.health = 50
			}
		}
	case "back", "ba", "bac":

	default:
		fmt.Println("Input invalid")
		p.use_item()
	}
}

func (p *player) normalize_stats() {
	if p.health > p.max_health {
		p.health = p.max_health
	}
	if p.skill_points > p.max_skill_points {
		p.skill_points = p.max_skill_points
	}
}

func (e *enemy) normalize_stats_enemy() {
	if e.health > e.max_health {
		e.health = e.max_health
	}
	if e.skill_points > e.max_skill_points {
		e.skill_points = e.max_skill_points
	}
}

/*
func remove_item(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
*/

// Exits the game
func quit() {
	os.Exit(0)
}
