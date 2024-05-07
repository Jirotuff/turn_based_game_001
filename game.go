package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	name_1 string = "Dario"
	name_2 string = "Pilgrim"
	name_3 string = "Fie"
	name_4 string = "Jessy"
)

var inventory []string
var gold int
var exp_gained int
var gold_gained int
var item_gained []string

var user_input string
var victory bool = false
var display_tutorial bool = true

func main() {

	user_input = ""
	check_victory()

	if display_tutorial {
		Tutorial()
	}

	Dario.show_status()
	Pilgrim.show_status()
	Fie.show_status()
	Jessy.show_status()

	fmt.Println("\ngold: ", gold)
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

			fmt.Println("\n[back]")

			fmt.Scanln(&user_input)

			switch strings.ToLower(user_input) {

			case "back", "b", "ba", "bac":
				clear_screen()
				main()

			}

		case "inv", "i", "in":
			display_inventory()

		case "exit", "ex", "exi":
			quit()

		default:
			clear_screen()
			main()
		}
	}
}

// This block below is for things that need to be checked and or changed

func check_victory() {
	if victory {
		victory = false

		fmt.Println("Victory!")

		exp_gained = rand.Intn(50) + 50

		Dario.exp += exp_gained
		Pilgrim.exp += exp_gained
		Fie.exp += exp_gained
		Jessy.exp += exp_gained

		gold_gained = rand.Intn(30) + 10

		gold = gold + gold_gained

		if rand.Intn(20) == 1 {
			item_gained = append(item_gained, "revival_bead")
		}

		if rand.Intn(20) > 17 {
			item_gained = append(item_gained, "potion")
		}

		inventory = append(inventory, item_gained...)

		fmt.Println("\nLoot:\n\nexp:", exp_gained, "\ngold: ", gold_gained, "\nitems:", item_gained)

		item_gained = nil
		gold_gained = 0
		exp_gained = 0

		Dario.Level_check()
		Fie.Level_check()
		Pilgrim.Level_check()
		Jessy.Level_check()

		Reset_enemy(&Bandit)

		fmt.Println("\nType any key to continue")
		fmt.Scanln(&user_input)
		clear_screen()
		if user_input == (" ") {
			main()
		} else {
			main()
		}

	}
}

// The block below is for different "places"

func combat() {
	fmt.Println("\n\nCombat started!")

	for !victory {
		Dario.Check_player_life()
		Bandit.Check_enemy_life()
		Dario.Player_turn()
		Bandit.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn()
		}
		Bandit.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn()
		}
		Bandit.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn()
		}
		Bandit.Check_enemy_life()

		Bandit.Enemy_turn()
	}
}

func shop() {

	fmt.Println("Welcome to the shop")
	fmt.Println("\nWe have a variety of products available, please take your time choosing")
	fmt.Println("\n- potion")
	fmt.Println("- fire_gem")
	fmt.Println("- revival_bead")
	fmt.Println("\nleave the shop (back)")

	for {

		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "potion", "po", "pot", "poti":
			if gold >= 50 {
				gold -= 50
				fmt.Println("you have bought a potion")
				inventory = append(inventory, "potion")
			} else {
				fmt.Println("You lack gold!")
			}
		case "fire_gem", "fi", "fir", "fire":
			if gold >= 25 {
				gold -= 25
				fmt.Println("you have bought a fire_gem")
				inventory = append(inventory, "fire_gem")
			} else {
				fmt.Println("You lack gold!")
			}
		case "revival_bead", "re", "rev", "revi":
			if gold >= 150 {
				gold -= 150
				fmt.Println("you have bought a revival_bead")
				inventory = append(inventory, "revival_bead")
			} else {
				fmt.Println("You lack gold!")
			}
		case "back", "b", "ba", "bac":
			clear_screen()
			main()

		default:
			fmt.Println("We don't have this item...")
		}

	}
}

func (p *player) display_stats() {
	fmt.Println("\n", p.name, "lv:", p.lv)
	fmt.Println("Exp:", p.exp)
	fmt.Println("\nStrength:", p.strength)
	fmt.Println("Intelligence:", p.intelligence)
	fmt.Println("Agility:", p.agility)
	fmt.Println("Endurance:", p.endurance)
	fmt.Println("Social:", p.social)
}

func display_inventory() {

	fmt.Println(inventory)

	fmt.Println("\nPress Enter to retun to main menu")

	fmt.Scanln(&user_input)

	switch strings.ToLower(user_input) {

	case "back", "b", "ba", "bac":
		clear_screen()
		main()

	default:
		clear_screen()
		main()
	}
}

// This block below is for misc

func contains_string(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}

	}
	return false
}

func remove_item(s []string, item_to_remove string) {
	var temp_inv []string

	for _, item := range s {
		if item != item_to_remove {
			temp_inv = append(temp_inv, item)
		}
	}
	inventory = temp_inv
}

func clear_screen() {
	cmd := exec.Command("clear") // for Unix/Linux
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func quit() {
	os.Exit(0)
}

func Tutorial() {
	display_tutorial = false
	fmt.Println("Welcome to this game...")
	fmt.Println("\nThis is a turn based game, as the player you can type the one of the moves to execute it.")
	fmt.Println("Your goal at this moment is to acquire as much gold as possible")
}

func (p *player) show_status() {
	if p.health > 0 {
		fmt.Println(p.name, ":\nhealth: ", p.health, "skill points: ", p.skill_points)
	} else {
		fmt.Println(p.name, ":\n\033[95m DEAD...\033[0m")
	}
}
