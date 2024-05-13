package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const ()

var (
	name_2 string = "Pilgrim"
	name_3 string = "Fie"
	name_4 string = "Jessy"

	equipment_phys_offense  int
	equipment_magic_defense int
	get_chest               bool = true
	name_selected           bool = false
	battle_intro            bool = true
	floor_level_key         int  = 1
	current_floor           int
	inventory               []string
	gold                    int
	exp_gained              int
	gold_gained             int
	item_gained             []string
	user_input              string
	victory                 bool = false
	display_tutorial        bool = true
)

func main() {

	if !name_selected {
		name_selected = true
		fmt.Println("What is your name?")
		fmt.Scanln(&user_input)
		Player.name = user_input
	}

	if contains_string(inventory, "bronze_sword") {
		equipment_phys_offense = 10
	}
	if contains_string(inventory, "tin_foil_hat") {
		equipment_magic_defense = 5
	}
	user_input = ""

	if display_tutorial {
		Tutorial()
	}
	battle_intro = true

	Player.show_status()
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
			Player.display_stats()
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

func (e *enemy) check_victory() {
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

		Player.exp += exp_gained
		Pilgrim.exp += exp_gained
		Fie.exp += exp_gained
		Jessy.exp += exp_gained

		gold_gained = rand.Intn(30) + 10

		gold = gold + gold_gained

		if rand.Intn(20) == 1 {
			item_gained = append(item_gained, "revival_bead")
		}
		if rand.Intn(8) == 1 {
			item_gained = append(item_gained, "potion")
		}
		if rand.Intn(30) == 1 {
			item_gained = append(item_gained, "iron")
		}
		if rand.Intn(10) == 1 {
			item_gained = append(item_gained, "copper")
		}
		if rand.Intn(4) == 1 {
			item_gained = append(item_gained, "tin")
		}
		if rand.Intn(5) == 1 {
			item_gained = append(item_gained, "coal")
		}

		inventory = append(inventory, item_gained...)

		fmt.Println("\nLoot:\n\nexp:", exp_gained, "\ngold: ", gold_gained, "\nitems:", item_gained)

		item_gained = nil
		gold_gained = 0
		exp_gained = 0

		Player.Level_check()
		Fie.Level_check()
		Pilgrim.Level_check()
		Jessy.Level_check()

		Reset_enemy(e)

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

// happens after check victory finsishes
func after_combat() {

	if rand.Intn(8) == 1 && get_chest {
		chest()
	}
	get_chest = true

	fmt.Println("Type 'battle' for another enounter or 'entrance' to return to the dungeon entrance")

	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "entrance", "e", "en", "ent", "entr", "entra", "entran", "entranc":
		dungeon()

	case "battle", "b", "ba", "bat", "batt", "battl":
		battle_intro = true
		combat_select()

	default:
		fmt.Println("Is that a typo?")
		after_combat()
	}
}

func combat_select() {

	var combat_select_var int

	switch current_floor {

	case 1:

		combat_select_var = rand.Intn(4)

		switch combat_select_var {

		case 1:
			combat_101()

		case 2:
			combat_102()

		default:
			combat_101()
		}
	case 2:
		combat_select_var = rand.Intn(3)

		switch combat_select_var {

		case 1:
			combat_201()

		case 2:
			combat_202()

		default:
			combat_201()
		}
	}
}

func chest() {
	fmt.Println("\033[96mYou have found a treasure chest\033[0m\n\nWould you like to use a lockpick to open it? [Y/N]\n\n", inventory)

	for {
		fmt.Scanln(&user_input)
		switch strings.ToLower(user_input) {

		case "yes", "y", "ye":
			if contains_string(inventory, "lockpick") {

				if current_floor == 1 {
					remove_item(inventory, "lockpick")
					gold_gained = rand.Intn(150) + 100
					gold += gold_gained
					item_gained = append(item_gained, "potion")
					inventory = append(inventory, item_gained...)
				}

				if current_floor == 2 {
					remove_item(inventory, "lockpick")
					gold_gained = rand.Intn(250) + 150
					gold += gold_gained
					item_gained = append(item_gained, "revival bead")
					inventory = append(inventory, item_gained...)
				}
				fmt.Println("\nLoot:\n\n gold: ", gold_gained, "\nitems: ", item_gained)

				gold_gained = 0
				item_gained = nil
			} else {
				fmt.Println("You don't have any lockpicks...")
			}
		case "no", "n":
			get_chest = false
			after_combat()
		default:
			chest()
		}
	}
}

// The block below is for different "places"

// Floor 1 enemies

func combat_101() {
	fmt.Println("\n\nCombat started!")

	for !victory {
		if battle_intro {
			battle_intro = false
			fmt.Println("\nYou encountered a Bandit!")
			fmt.Println("\npress Enter to continue")
			fmt.Scanln(&user_input)

		}
		Player.Check_player_life()
		Bandit.Check_enemy_life()
		Player.Player_turn(&Bandit)
		Bandit.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn(&Bandit)
			Bandit.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn(&Bandit)
			Bandit.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn(&Bandit)
			Bandit.Check_enemy_life()
		}

		Bandit.Enemy_turn()
	}
}

func combat_102() {
	fmt.Println("\n\nCombat started!")

	for !victory {
		if battle_intro {
			battle_intro = false
			fmt.Println("\nYou encounter a Goblin!")
			fmt.Println("\npress Enter to continue")
			fmt.Scanln(&user_input)

		}
		Player.Check_player_life()
		Goblin.Check_enemy_life()
		Player.Player_turn(&Goblin)
		Goblin.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn(&Goblin)
			Goblin.Check_enemy_life()
		}
		Goblin.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn(&Goblin)
			Goblin.Check_enemy_life()
		}
		Goblin.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn(&Goblin)
			Goblin.Check_enemy_life()
		}

		Goblin.Enemy_turn()
	}
}

// Floor 2 enemies

func combat_201() {
	fmt.Println("\n\nCombat started!")

	if battle_intro {
		battle_intro = false
		fmt.Println("\nYou encountered a Dark knight!")
		fmt.Println("\npress Enter to continue")
		fmt.Scanln(&user_input)
	}
	for !victory {
		Player.Check_player_life()
		Dark_knight.Check_enemy_life()
		Player.Player_turn(&Dark_knight)
		Dark_knight.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn(&Dark_knight)
			Dark_knight.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn(&Dark_knight)
			Dark_knight.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn(&Dark_knight)
			Dark_knight.Check_enemy_life()
		}

		Dark_knight.Enemy_turn()
	}
}

func combat_202() {
	fmt.Println("\n\nCombat started!")

	for !victory {
		if battle_intro {
			battle_intro = false
			fmt.Println("\nYou encounter a stone golem!")
			fmt.Println("\npress Enter to continue")
			fmt.Scanln(&user_input)

		}
		Player.Check_player_life()
		Golem.Check_enemy_life()
		Player.Player_turn(&Golem)
		Golem.Check_enemy_life()
		if Pilgrim.health > 0 {
			Pilgrim.Player_turn(&Golem)
			Golem.Check_enemy_life()
		}
		Golem.Check_enemy_life()
		if Fie.health > 0 {
			Fie.Player_turn(&Golem)
			Golem.Check_enemy_life()
		}
		Golem.Check_enemy_life()
		if Jessy.health > 0 {
			Jessy.Player_turn(&Golem)
			Golem.Check_enemy_life()
		}

		Golem.Enemy_turn()
	}
}

func dungeon() {
	clear_screen()

	fmt.Println("\nWhich floor?\n\nfloor 1...\nfloor 2...\nfloor 3...\n\n[back]")

	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "floor 1", "1":
		current_floor = 1
		combat_select()
	case "floor 2", "2":
		if floor_level_key >= 2 {
			current_floor = 2
			combat_select()
		} else {
			fmt.Println("You don't have the required key!")

		}
	case "floor 3", "3":
		game_finished()

	case "back", "b", "ba", "bac":
		clear_screen()
		main()

	default:
		fmt.Println("Something went wrong")
		dungeon()
	}
}

func shop() {
	clear_screen()

	fmt.Println("Welcome to the shop")
	fmt.Println("\ngold: ", gold)
	fmt.Println("\n- potion\t\t50 gold")
	fmt.Println("- revival_bead\t\t150 gold")
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
	clear_screen()

	fmt.Println("Welcome to the smithy")
	fmt.Println("\nIn here you can craft equipment, materials and items")
	fmt.Println("\n- equipment\n- items\n- materials\n\nleave the smithy (back)")

	fmt.Scanln(&user_input)

	switch strings.ToLower(user_input) {

	case "back", "b", "ba", "bac":
		clear_screen()
		main()

	case "equipment", "equip", "eq", "equ", "equi", "equipm", "equipme", "equipmen", "e":
		smithy_equip()

	case "item", "it", "ite", "i":
		smithy_item()

	case "materials", "material", "materia", "materi", "mater", "mate", "mat", "ma", "m":
		smithy_material()

	default:
		smithy()
	}

}

func smithy_equip() {
	fmt.Println("\n- sword\t\t\tCosts 1 bronze and 50 gold | increases physical offense for the entire party by 10")
	fmt.Println("- tin foil hat\t\tCosts 1 tin and 5 gold | increases magical defense for the entire party by 5")

	for {

		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "sword", "sw", "swo", "swor":
			if contains_string(inventory, "bronze") {
				if gold >= 50 {
					gold -= 50
					remove_item(inventory, "bronze")
					fmt.Println("you've crafted a bronze_sword")
					inventory = append(inventory, "bronze_sword")
				} else {
					fmt.Println("You lack gold!")
				}

			} else {
				fmt.Println("You lack Bronze!")
			}

		case "tin_foil_hat", "tin_foil_ha", "tin_foil_h", "tin_foil", "tin_foi", "tin_fo", "tin_f", "tin", "ti":
			if contains_string(inventory, "tin") {
				if gold >= 5 {
					gold -= 5
					remove_item(inventory, "tin")
					fmt.Println("you've crafted a tin foil hat")
					inventory = append(inventory, "tin_foil_hat")
				} else {
					fmt.Println("You lack gold")
				}
			} else {
				fmt.Println("You lack tin!")
			}

		case "back", "b", "ba", "bac":
			clear_screen()
			smithy()

		case "exit", "ex", "exi":
			clear_screen()
			main()

		default:
			fmt.Println("You cant make this...")
		}

	}
}

func smithy_item() {

	fmt.Println("\n- lockpicks\t\tCosts 1 iron and 20 gold, used to open treasure chests")

	for {
		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "lockpick", "l", "lo", "loc", "lock", "lockp", "lockpi", "lockpic":
			if contains_string(inventory, "iron") {
				if gold >= 20 {
					gold -= 20
					remove_item(inventory, "iron")
					fmt.Println("you've crafted some lockpicks")
					inventory = append(inventory, "lockpick", "lockpick", "lockpick")
				}

			} else {
				fmt.Println("You lack iron!")
			}

		case "back", "b", "ba", "bac":
			clear_screen()
			smithy()

		case "exit", "ex", "exi":
			clear_screen()
			main()

		default:
			fmt.Println("You can't craft this...")
		}
	}
}

func smithy_material() {
	fmt.Println("\n- bronze\t\tCosts 1 copper, 1 tin and 30 gold")

	for {
		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "bronze", "bronz", "bron", "bro", "br":
			if contains_string(inventory, "copper") {
				if contains_string(inventory, "tin") {
					if gold >= 30 {
						gold -= 30
						remove_item(inventory, "copper")
						fmt.Println("you've crafted some bronze")
						inventory = append(inventory, "bronze")
					} else {
						fmt.Println("You lack gold")
					}
				} else {
					fmt.Println("You lack tin")
				}

			} else {
				fmt.Println("You lack copper!")
			}

		case "back", "ba", "bac":
			clear_screen()
			smithy()

		case "exit", "ex", "exi":
			clear_screen()
			main()

		default:
			fmt.Println("You can't craft this...")
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
	fmt.Println("Welcome to this game", Player.name)
	fmt.Println("\nThis is a turn based game, as the player you can type the one of the moves to execute it.")
	fmt.Println("Your goal is to progress to the third floor of the dungeon. defeating enemies on one floor will get you the key to the next")
	fmt.Println("This is a debug version")
	fmt.Println("")
}

func (p *player) show_status() {
	if p.health > 0 {
		fmt.Println(p.name, ":\nhealth: ", p.health, "skill points: ", p.skill_points, "Special points: ", p.special)
	} else {
		fmt.Println(p.name, ":\n\033[95m DEAD...\033[0m")
	}
}

func (e *enemy) show_status() {
	if e.health > 0 {
		fmt.Println(e.name, ":\nhealth: ", e.health, "skill points: ", e.skill_points)
	} else {
		fmt.Println(e.name, ":\n\033[95m DEAD...\033[0m")
	}
}

func game_finished() {
	clear_screen()
	fmt.Println("You have beaten this game!!! congratz!\n\nGame made by Jimmy Roodzant / 2024\n\nPress Enter to exit")
	fmt.Scanln(&user_input)
	os.Exit(0)
}
