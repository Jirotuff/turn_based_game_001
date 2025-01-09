package main

import (
	"encoding/json"
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
		name_selection()
	}

	if display_tutorial {
		Tutorial()
	}

	check_equipment()

	user_input = ""

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
	fmt.Println("Data management\t> save and load here")
	fmt.Println("Exit\t\t> exits the game")
	fmt.Println("")

	for {
		fmt.Scanln(&user_input)

		switch strings.ToLower(user_input) {

		case "data", "dat", "da":
			data_management()

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

func name_selection() {

	name_selected = true
	fmt.Println("What is your Name?")
	fmt.Scanln(&user_input)
	Player.Name = user_input
}

func check_equipment() {
	if contains_string(inventory, "bronze_sword") {
		equipment_phys_offense = 10
	}
	if contains_string(inventory, "tin_foil_hat") {
		equipment_magic_defense = 5
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

		Player.Exp += exp_gained
		Pilgrim.Exp += exp_gained
		Fie.Exp += exp_gained
		Jessy.Exp += exp_gained

		gold_gained = rand.Intn(30) + 10

		gold += gold_gained

		if rand.Intn(20) == 1 {
			item_gained = append(item_gained, "revival_bead")
		}
		if rand.Intn(8) == 1 {
			item_gained = append(item_gained, "potion")
		}
		if rand.Intn(15) == 1 {
			item_gained = append(item_gained, "iron")
		}
		if rand.Intn(10) == 1 {
			item_gained = append(item_gained, "copper")
		}
		if rand.Intn(4) == 1 {
			item_gained = append(item_gained, "tin")
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
		after_combat()
	}
}

// happens after check victory finsishes
func after_combat() {

	if rand.Intn(8) == 1 && get_chest {
		chest()
	}

	fmt.Println("Type 'battle' for another enounter or 'entrance' to return to the dungeon entrance")

	fmt.Scanln(&user_input)
	switch strings.ToLower(user_input) {

	case "entrance", "e", "en", "ent", "entr", "entra", "entran", "entranc":
		get_chest = true
		dungeon()

	case "battle", "b", "ba", "bat", "batt", "battl":
		get_chest = true
		battle_intro = true
		combat_select()

	default:
		get_chest = false
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
			get_chest = false
			if !contains_string(inventory, "lockpick") {
				fmt.Println("You don't have any lockpicks...")
				continue
			}

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

		case "no", "n":
			get_chest = false
			after_combat()
		default:
			fmt.Println("Was that a typo?")
			continue
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
		if Pilgrim.Health > 0 {
			Pilgrim.Player_turn(&Bandit)
			Bandit.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Fie.Health > 0 {
			Fie.Player_turn(&Bandit)
			Bandit.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Jessy.Health > 0 {
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
		if Pilgrim.Health > 0 {
			Pilgrim.Player_turn(&Goblin)
			Goblin.Check_enemy_life()
		}
		Goblin.Check_enemy_life()
		if Fie.Health > 0 {
			Fie.Player_turn(&Goblin)
			Goblin.Check_enemy_life()
		}
		Goblin.Check_enemy_life()
		if Jessy.Health > 0 {
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
		if Pilgrim.Health > 0 {
			Pilgrim.Player_turn(&Dark_knight)
			Dark_knight.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Fie.Health > 0 {
			Fie.Player_turn(&Dark_knight)
			Dark_knight.Check_enemy_life()
		}
		Bandit.Check_enemy_life()
		if Jessy.Health > 0 {
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
		if Pilgrim.Health > 0 {
			Pilgrim.Player_turn(&Golem)
			Golem.Check_enemy_life()
		}
		Golem.Check_enemy_life()
		if Fie.Health > 0 {
			Fie.Player_turn(&Golem)
			Golem.Check_enemy_life()
		}
		Golem.Check_enemy_life()
		if Jessy.Health > 0 {
			Jessy.Player_turn(&Golem)
			Golem.Check_enemy_life()
		}

		Golem.Enemy_turn()
	}
}

func dungeon() {
	clear_screen()

	fmt.Println("\nWhich floor?\n\nfloor 1...\nfloor 2...\nfloor 3...\n\n[back]")

	for {

		fmt.Scanln(&user_input)
		switch strings.ToLower(user_input) {

		case "floor 1", "1":
			current_floor = 1
			combat_select()
		case "floor 2", "2":
			if floor_level_key < 2 {
				fmt.Println("You don't have the required key!")
				continue
			}

			current_floor = 2
			combat_select()

		case "floor 3", "3":
			if floor_level_key < 3 {
				fmt.Println("You don't have the required key!")
				continue
			}

			game_finished()

		case "back", "b", "ba", "bac":
			clear_screen()
			main()

		default:
			fmt.Println("Something went wrong")
			dungeon()
		}
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
			if gold < 50 {
				fmt.Println("You lack gold!")
				continue
			}

			gold -= 50
			fmt.Println("you have bought a potion")
			inventory = append(inventory, "potion")

		case "revival_bead", "re", "rev", "revi", "reviv", "reviva", "revival_", "revival_b":
			if gold < 150 {
				fmt.Println("You lack gold!")
				continue
			}

			gold -= 150
			fmt.Println("you have bought a revival_bead")
			inventory = append(inventory, "revival_bead")

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
			if !contains_string(inventory, "bronze") {
				fmt.Println("You lack Bronze!")
				continue
			}

			if gold < 50 {
				fmt.Println("You lack gold!")
				continue
			}

			gold -= 50
			remove_item(inventory, "bronze")
			fmt.Println("you've crafted a bronze_sword")
			inventory = append(inventory, "bronze_sword")

		case "tin foil hat", "tin foil ha", "tin foil h", "tin foil", "tin foi", "tin fo", "tin f", "tin", "ti":
			if !contains_string(inventory, "tin") {
				fmt.Println("You lack tin!")
				continue
			}

			if gold < 5 {
				fmt.Println("You lack gold")
				continue
			}

			gold -= 5
			remove_item(inventory, "tin")
			fmt.Println("you've crafted a tin foil hat")
			inventory = append(inventory, "tin_foil_hat")

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
			if !contains_string(inventory, "iron") {
				fmt.Println("You lack iron!")
				continue
			}

			if gold < 20 {
				fmt.Println("You lack gold")
				continue
			}

			gold -= 20
			remove_item(inventory, "iron")
			fmt.Println("you've crafted some lockpicks")
			inventory = append(inventory, "lockpick", "lockpick", "lockpick")

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
			if !contains_string(inventory, "copper") {
				fmt.Println("You lack copper!")
				continue
			}
			if !contains_string(inventory, "tin") {
				fmt.Println("You lack tin")
				continue
			}
			if gold < 30 {
				fmt.Println("You lack gold")
				continue
			}

			gold -= 30
			remove_item(inventory, "copper")
			fmt.Println("you've crafted some bronze")
			inventory = append(inventory, "bronze")

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

	fmt.Println("\n", p.Name, "lv:", p.Lv)
	fmt.Println("Exp:", p.Exp)
	fmt.Println("\nStrength:", p.Strength)
	fmt.Println("Intelligence:", p.Intelligence)
	fmt.Println("Agility:", p.Agility)
	fmt.Println("Endurance:", p.Endurance)
}

func display_inventory() {

	fmt.Println(inventory)

	fmt.Println("\nPress Enter to retun to main menu")

	fmt.Scanln(&user_input)

	clear_screen()
	main()

}

func data_management() {
	clear_screen()

	fmt.Println("In this menu you can save your progress")
	fmt.Println("Type 'save'")

	fmt.Scanln(&user_input)

	switch strings.ToLower(user_input) {

	case "save":
		save_game()
	}
}

func save_game() {

	characters := []player{Player, Pilgrim, Fie, Jessy}

	fmt.Println("Saving game data...")

	character_save, err := json.Marshal(characters)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile("output.json", character_save, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Save completed!")
	fmt.Println("\n***Press enter to continue***")

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
	fmt.Println("Welcome to this game", Player.Name)
	fmt.Println("\nThis is a turn based game, as the player you can type the one of the moves to execute it.")
	fmt.Println("Your goal is to progress to the third floor of the dungeon. defeating enemies on one floor will get you the key to the next")
	fmt.Println("This is a debug version")
	fmt.Println("")
}

func (p *player) show_status() {
	if p.Health > 0 {
		fmt.Println(p.Name, ":\nhealth: ", p.Health, "skill points: ", p.Skill_points, "Special points: ", p.Special)
	} else {
		fmt.Println(p.Name, ":\n\033[95m DEAD...\033[0m")
	}
}

func (e *enemy) show_status() {
	if e.Health > 0 {
		fmt.Println(e.Name, ":\nhealth: ", e.Health, "skill points: ", e.Skill_points)
	} else {
		fmt.Println(e.Name, ":\n\033[95m DEAD...\033[0m")
	}
}

func game_finished() {
	clear_screen()
	fmt.Println("You have beaten this game!!! congratz!\n\nGame made by Jimmy Roodzant / 2024\n\nPress Enter to exit")
	fmt.Scanln(&user_input)
	os.Exit(0)
}
