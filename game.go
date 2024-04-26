package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

/* Stuff to remember

this is a way to print colored text

{
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 91, "Sample text")
			fmt.Println(colored)
		}
*/

// variables
var display_tutorial bool = true

// player general
var player_name string
var user_input string //player input
var victory bool = false
var gold int = 50
var player_lv int = 1
var player_exp int = 0
var player_max_health int = 100
var player_max_skill_points int = 80
var player_inventory []string
var player_special int

// player stats
var player_health int = 100 //player health
var player_skill_points int = 80
var player_strength int = 10     // increases physical damage
var player_intelligence int = 10 // increases magical damage
var player_agility int = 10      // increases chance to dodge
var player_endurance int = 10    // reduces damage taken
var player_social int = 10       // reduces shop prices

// enemy status
var enemy_input int        //enemy input
var enemy_health int = 100 //enemy health
var enemy_skill_points int = 100
var enemy_max_skill_points int = 80
var enemy_max_health int = 100

// start of program
func main() {

	clear_screen()

	if display_tutorial != false {
		tutorial()
	}

	if victory != false {
		victory = false
		enemy_health = enemy_max_health
		enemy_skill_points = enemy_max_skill_points
		player_exp += rand.Intn(50) + 50
		player_level_up()
		main()
	}

	fmt.Println(player_name, "  Health:", player_health, "SP:", player_skill_points, "Gold:", gold)
	fmt.Println("\nWhat do you want to do?")
	fmt.Println("\nbattle\t\t> finds opponent")
	fmt.Println("shop\t\t> enter the shop")
	fmt.Println("stats\t\t> show player stats")
	fmt.Println("inv\t\t> show player inventory")
	fmt.Println("exit\t\t> exits the game")
	fmt.Println("")

	for {
		fmt.Scanln(&user_input)

		switch user_input {

		case "battle":
			combat()

		case "shop":
			shop()

		case "stats":
			display_stats()

		case "inv":
			display_inventory()

		case "exit":
			quit()

		default:
			main()
		}
	}
}

func save(slot1 any, data interface{}) {

}

// starts the combat encounter
func combat() {
	fmt.Println("\n\nCombat started!")

	for {
		check_player_life()
		check_enemy_life()
		if victory == true {
			gold += rand.Intn(10) + 5
			main()
		}

		player_turn()

		enemy_turn()

	}
}

func player_turn() {
	fmt.Println("")
	if player_special >= 3 {
		{
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 91, "You feel a strange power welling up inside... (type 'special' to unleash it)")
			fmt.Println(colored)
		}
	}
	fmt.Println("\n>", player_name)
	fmt.Println("Health:", player_health, "SP:", player_skill_points)
	fmt.Println("\n> Enemy")
	fmt.Println("Health:", enemy_health)
	fmt.Println("\nWhat's your move?")
	fmt.Println("\n>> strike\t\t\t> Use your basic weapon\t")
	fmt.Println(">> heal\t\t\t\t> Use an healing item\t")
	fmt.Println(">> force | 20 SP\t\t> High citical chance attack")
	fmt.Println(">> soul \t\t\t> Regenerates some SP")
	fmt.Println("")

	fmt.Scanln(&user_input)

	if enemy_health > enemy_max_health {
		enemy_health = enemy_max_health
	}
	if enemy_skill_points > enemy_max_skill_points {
		enemy_skill_points = enemy_max_skill_points
	}

	switch user_input { //gives different options to the player

	case "strike":
		player_skill_strike()

	case "heal":
		player_skill_heal()

	case "force":
		player_skill_force()

	case "soul":
		player_skill_soul()

	case "kill":
		player_skill_kill()

	case "special":
		if player_special > 2 {
			player_special = 0
			player_skill_special()
		} else {
			fmt.Println("You dont have the energy for this move")
		}
	default:
		fmt.Println("Thats a typo! lost your turn XD")
	}

}

// function for enemy turn
func enemy_turn() {

	if player_health > player_max_health {
		player_health = player_max_health
	}
	if player_skill_points > player_max_skill_points {
		player_skill_points = player_max_skill_points
	}

	enemy_input = rand.Intn(3) //gives different options to the enemy

	if enemy_health >= 1 {
		switch enemy_input {

		case 0:
			enemy_skill_strike()

		case 1:
			enemy_skill_heal()

		case 2:
			enemy_skill_force()
		}
	}

}

// displays a tutorial if display_tutorial == true
func tutorial() {
	display_tutorial = false
	fmt.Println("Welcome to this game...")
	fmt.Println("\nThis is a turn based game, as the player you can type the one of the moves to execute it.")
	fmt.Println("Your goal at this moment is to acquire as much gold as possible")
	fmt.Println("\nWhat is you name?")
	fmt.Println("")
	fmt.Scanln(&user_input)
	player_name = user_input

}

// checks if the player is dead
func check_player_life() {
	if player_health <= 0 {
		fmt.Println("You have been killed!")
		fmt.Println("\nGold:", gold, "Player level:", player_lv)
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

// player skill: kill (THIS IS A TEST FEATURE, NOT MEANT FOR FINAL PRODUCT)
func player_skill_kill() {
	damage := rand.Intn(20) + 5 + player_strength + 999
	critical_damage := rand.Intn(20) + 30 + player_strength + 999

	if rand.Intn(11) == 9 { //Critical hit chance
		enemy_health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		enemy_health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

// player skill: strike
func player_skill_strike() {
	damage := rand.Intn(20) + 5 + player_strength
	critical_damage := rand.Intn(20) + 30 + player_strength

	player_special += 1

	if rand.Intn(11) == 9 { //Critical hit chance
		enemy_health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		enemy_health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

// player skill: soul
func player_skill_soul() {
	if true == true {
		player_skill_points += 25
	}
}

// player skill: force
func player_skill_force() {
	damage := rand.Intn(5) + 20 + player_intelligence
	critical_damage := rand.Intn(20) + 30 + player_intelligence

	if player_skill_points >= 20 {

		player_skill_points -= 20

		player_special += 1

		if rand.Intn(3) == 2 { //Critical hit chance

			enemy_health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			user_input = ""
		} else {
			enemy_health -= damage
			fmt.Println(damage, "DMG")
			user_input = ""
		}
	} else {
		fmt.Println("You tried to cast force... but you dont have enough SP!")
		user_input = ""
	}
}

// player skill: heal
func player_skill_heal() {
	heal := rand.Intn(20) + 5 + player_intelligence //amount healed
	player_health += heal
	fmt.Println(heal, "Healed")
	user_input = ""
}

func player_skill_special() {
	damage := 70 + player_strength
	critical_damage := rand.Intn(20) + 75 + player_strength

	if rand.Intn(11) == 9 { //Critical hit chance
		enemy_health -= critical_damage
		fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		fmt.Println("")
		user_input = ""
	} else {
		enemy_health -= damage
		fmt.Println(damage, "DMG")
		user_input = ""
	}
}

// checks if the enemy is dead
func check_enemy_life() {
	if enemy_health <= 0 {
		fmt.Println("Victory!")
		victory = true
	}
}

// enemy skill: strike
func enemy_skill_strike() {
	fmt.Println("Enemy used strike")
	damage := rand.Intn(20) + 5 - player_endurance
	critical_damage := rand.Intn(20) + 30 - player_endurance

	if rand.Intn(100) > player_agility {

		if rand.Intn(11) == 9 { //Critical hit chance
			player_health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		} else {
			player_health -= damage
			fmt.Println(damage, "DMG")
		}
	} else {
		fmt.Println("But it missed!")
	}
}

// enemy skill: heal
func enemy_skill_heal() {
	heal := rand.Intn(20) + 5 //amount healed
	fmt.Println("Enemy has healed")
	enemy_health += heal
	fmt.Println(heal, "Healed")
}

// enemy skill: force
func enemy_skill_force() {

	damage := rand.Intn(10) + 20 - player_endurance
	critical_damage := rand.Intn(20) + 30 - player_endurance

	fmt.Println("Enemy cast force")

	if enemy_skill_points >= 20 {

		enemy_skill_points -= 20

		if rand.Intn(100) >= player_agility {

			if rand.Intn(3) == 1 { //Critical hit chance
				player_health -= critical_damage
				fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			} else {
				player_health -= damage
				fmt.Println(damage, "DMG")
			}
		} else {
			fmt.Println("but it missed")
		}
	} else {
		fmt.Println("but nothing happened...")
		damage = 0
		player_health -= damage
		fmt.Println(damage, "DMG")
		fmt.Scanln()
	}
}

// clear the screen in the CLI
func clear_screen() {
	cmd := exec.Command("clear") // for Unix/Linux
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// checks the exp and increases the player_lv
func player_level_up() {
	if player_exp >= 100 && player_lv < 2 {
		player_lv += 1
		player_max_health += 20
		player_max_skill_points += 5
		player_health = player_max_health
		player_skill_points = player_max_skill_points
		fmt.Println("")
		{
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, "Level up!")
			fmt.Println(colored)
		}
		fmt.Println("\nMax HP:", player_max_health, "Max SP:", player_max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("\nType the first 2 letters...")
		fmt.Println("\nStrength:", player_strength)
		fmt.Println("Intelligence", player_intelligence)
		fmt.Println("Agility:", player_agility)
		fmt.Println("Endurance:", player_endurance)
		fmt.Println("Social:", player_social)
		fmt.Println("")

		fmt.Scanln(&user_input)

		switch user_input {

		case "st":
			player_strength += 2
		case "in":
			player_intelligence += 2
		case "ag":
			player_agility += 2
		case "en":
			player_endurance += 2
		case "so":
			player_social += 2
		}
	}
	if player_exp >= 200 && player_lv < 3 {
		player_lv += 1
		player_max_health += 20
		player_max_skill_points += 5
		player_health = player_max_health
		player_skill_points = player_max_skill_points
		fmt.Println("\nLevel up!!")
		fmt.Println("\nMax HP:", player_max_health, "Max SP:", player_max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("")
		fmt.Scanln(&user_input)

		switch user_input {

		case "st":
			player_strength += 2
		case "in":
			player_intelligence += 2
		case "ag":
			player_agility += 2
		case "en":
			player_endurance += 2
		case "so":
			player_social += 2
		}
	}
	if player_exp >= 500 && player_lv < 4 {
		player_lv += 1
		player_max_health += 20
		player_max_skill_points += 5
		player_health = player_max_health
		player_skill_points = player_max_skill_points
		fmt.Println("\nLevel up!!")
		fmt.Println("\nMax HP:", player_max_health, "Max SP:", player_max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("")
		fmt.Scanln(&user_input)

		switch user_input {

		case "st":
			player_strength += 2
		case "in":
			player_intelligence += 2
		case "ag":
			player_agility += 2
		case "en":
			player_endurance += 2
		case "so":
			player_social += 2
		}
	}
	if player_exp >= 1000 && player_lv < 5 {
		player_lv += 1
		player_max_health += 20
		player_max_skill_points += 5
		player_health = player_max_health
		player_skill_points = player_max_skill_points
		fmt.Println("\nLevel up!!")
		fmt.Println("\nMax HP:", player_max_health, "Max SP:", player_max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Println("")
		fmt.Scanln(&user_input)

		switch user_input {

		case "st":
			player_strength += 2
		case "in":
			player_intelligence += 2
		case "ag":
			player_agility += 2
		case "en":
			player_endurance += 2
		case "so":
			player_social += 2
		}

	}
	if player_exp >= 2000 && player_lv < 6 {
		player_lv += 1
		player_max_health += 20
		player_max_skill_points += 5
		player_health = player_max_health
		player_skill_points = player_max_skill_points
		fmt.Println("\nLevel up!!")
		fmt.Println("\nMax HP:", player_max_health, "Max SP:", player_max_skill_points)
		fmt.Println("\nWhat stat would you like to improve?")
		fmt.Scanln(&user_input)

		switch user_input {

		case "st":
			player_strength += 2
		case "in":
			player_intelligence += 2
		case "ag":
			player_agility += 2
		case "en":
			player_endurance += 2
		case "so":
			player_social += 2
		}
	}
}

// Makes the shop work | not finished yet
func shop() {

	fmt.Println("Welcome to the shop")
	fmt.Println("\nWe have a variety of products available, please take your time choosing")
	fmt.Println("\n- potion")
	fmt.Println("- sword")
	fmt.Println("- shield")
	fmt.Println("\nleave the shop (back)")

	for {

		fmt.Scanln(&user_input)

		switch user_input { //gives different options to the player

		case "potion":
			fmt.Println("you have bought a potion")
			player_inventory = append(player_inventory, "potion")

		case "sword":
			fmt.Println("you have bought a sword")
			player_inventory = append(player_inventory, "sword")

		case "shield":
			fmt.Println("you have bought a shield")
			player_inventory = append(player_inventory, "shield")

		case "back":
			main()

		default:
			fmt.Println("We don't have this item...")
		}

	}
}

func display_stats() {
	fmt.Println("\nPlayer lv:", player_lv)
	fmt.Println("Exp:", player_exp)
	fmt.Println("\nStrength:", player_strength)
	fmt.Println("Intelligence", player_intelligence)
	fmt.Println("Agility:", player_agility)
	fmt.Println("Endurance:", player_endurance)
	fmt.Println("Social:", player_social)
	fmt.Println("\n[back]")

	fmt.Scanln(&user_input)

	switch user_input {

	case "back":
		main()

	default:
		main()
	}
}

// displays the player's inventory | doesn't work yet
func display_inventory() {

	fmt.Println(player_inventory)

	fmt.Println("\n[back]")

	fmt.Scanln(&user_input)

	switch user_input {

	case "back":
		main()

	default:
		main()
	}
}

func quit() {
	os.Exit(0)
}
