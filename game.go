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

var battle_intro bool = true
var floor_level_key int = 1
var current_floor int
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
	battle_intro = true

	Dario.show_status()
	Pilgrim.show_status()
	Fie.show_status()
	Jessy.show_status()

	fmt.Println("\nGold: ", gold)
	fmt.Println("\nWhat do you want to do?")
	fmt.Println("\nDungeon\t\t> enter the dungeon")
	fmt.Println("Shop\t\t> enter the shop")
	fmt.Println("Smithy\t\t> enter the smithy")
	fmt.Println("Stats\t\t> show player stats")
	fmt.Println("Inventory\t> show player inventory")
	fmt.Println("Exit\t\t> exits the game")
	fmt.Println("")

	for {
		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "dungeon", "du", "dun", "dung", "dunge", "dungeo":
			dungeon()

		case "shop", "sh", "sho":
			shop()

		case "smithy", "sm", "smi", "smit", "smith":
			smithy()

		case "stats", "st", "sta", "stat":
			Dario.display_stats()
			Pilgrim.display_stats()
			Fie.display_stats()
			Jessy.display_stats()

			fmt.Println("\nPress enter to continue")

			fmt.Scanln(&user_input)

			switch strings.ToLower(user_input) {

			case "":
				clear_screen()
				main()

			default:
				clear_screen()
				main()
			}
		case "inventory", "i", "in", "inv", "inve", "inven", "invent", "invento", "inventor":
			display_inventory()

		case "exit", "e", "ex", "exi":
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

		if rand.Intn(5) == 1 {

			if floor_level_key < 2 {
				floor_level_key = 2
				fmt.Println("\nYou've found the key to the second floor!")
				inventory = append(inventory, "2F_key")
			}

			if floor_level_key < 3 && current_floor == 2 {
				floor_level_key = 3
				fmt.Println("\nYou've found the key to the third floor!")
				inventory = append(inventory, "3F_key")
			}
		}

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

		if rand.Intn(5) == 1 {
			item_gained = append(item_gained, "iron")
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
			clear_screen()
			after_combat()
		} else {
			clear_screen()
			after_combat()
		}
	}
}

func after_combat() {

	fmt.Println("Type 'battle' for another enounter or 'entrance' to return to the dungeon entrance")

	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "entrance", "e", "en", "ent", "entr", "entra", "entran", "entranc":
		dungeon()

	case "battle", "b", "ba", "bat", "batt", "battl":
		battle_intro = true

		switch current_floor {

		case 1:
			combat_001()
		case 2:
			combat_002()

		}

	default:
		fmt.Println("Is that a typo?")
		after_combat()
	}
}

// The block below is for different "places"

func combat_001() {
	fmt.Println("\n\nCombat started!")

	for !victory {
		if battle_intro {
			battle_intro = false
			fmt.Println("\nYou encountered a Bandit!")
			fmt.Println("\npress Enter to continue")
			fmt.Scanln(&user_input)

		}
		Dario.Check_player_life()
		Bandit.Check_enemy_life()
		Dario.Player_turn()
		Bandit.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn()
			Bandit.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn()
			Bandit.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn()
			Bandit.Check_enemy_life()
		}

		Bandit.Enemy_turn()
	}
}

func combat_002() {
	fmt.Println("\n\nCombat started!")

	if battle_intro {
		battle_intro = false
		fmt.Println("\nYou encountered a Dark knight!")
		fmt.Println("\npress Enter to continue")
		fmt.Scanln(&user_input)
	}
	for !victory {
		Dario.Check_player_life()
		Dark_knight.Check_enemy_life()
		Dario.Player_turn()
		Dark_knight.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn()
			Dark_knight.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn()
			Dark_knight.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn()
			Dark_knight.Check_enemy_life()
		}

		Dark_knight.Enemy_turn()
	}
}

func dungeon() {

	fmt.Println("\nWhich floor?\n\nfloor 1...\nfloor 2...\n\n[back]")

	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "floor 1", "1":
		current_floor = 1
		combat_001()

	case "floor 2", "2":
		if floor_level_key >= 2 {
			current_floor = 2
			combat_002()
		} else {
			fmt.Println("You don't have the required key!")
		}
	case "back", "b", "ba", "bac":
		clear_screen()
		main()

	default:
		fmt.Println("Something went wrong")
		dungeon()
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

		case "potion", "po", "pot", "poti", "potio":
			if gold >= 50 {
				gold -= 50
				fmt.Println("you have bought a potion")
				inventory = append(inventory, "potion")
			} else {
				fmt.Println("You lack gold!")
			}
		case "fire_gem", "fi", "fir", "fire", "fire_", "fire_g", "fire_ge":
			if gold >= 25 {
				gold -= 25
				fmt.Println("you have bought a fire_gem")
				inventory = append(inventory, "fire_gem")
			} else {
				fmt.Println("You lack gold!")
			}
		case "revival_bead", "re", "rev", "revi", "reviv", "reviva", "revival_", "revival_b":
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

func smithy() {
	fmt.Println("Welcome to the smithy")
	fmt.Println("\nIn here you can craft equipment for your party...")
	fmt.Println("\n- sword\t\t\tCosts 1 iron and 50 gold")
	fmt.Println("\nleave the smithy (back)")

	for {

		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "sword", "sw", "swo", "swor":
			if contains_string(inventory, "iron") {
				if gold >= 50 {
					gold -= 50
					remove_item(inventory, "iron")
					fmt.Println("you have crafted a sword")
					inventory = append(inventory, "sword")
				} else {
					fmt.Println("You lack gold!")
				}

			} else {
				fmt.Println("You lack iron!")
			}

		case "back", "b", "ba", "bac":
			clear_screen()
			main()

		default:
			fmt.Println("You cant make this...")
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
	fmt.Println("Your goal is to defeat the demon on the bottom of the dungeon")
}

func (p *player) show_status() {
	if p.health > 0 {
		fmt.Println(p.name, ":\nhealth: ", p.health, "skill points: ", p.skill_points)
	} else {
		fmt.Println(p.name, ":\n\033[95m DEAD...\033[0m")
	}
}
