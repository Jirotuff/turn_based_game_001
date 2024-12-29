package main

import (
	"fmt"
	"math/rand"
)

type enemy struct {
	Name             string
	Health           int
	Skill_points     int
	Max_skill_points int
	Max_health       int
}

var Bandit = enemy{
	Name:             "Bandit",
	Health:           200,
	Skill_points:     80,
	Max_skill_points: 80,
	Max_health:       200,
}

var Dark_knight = enemy{
	Name:             "Dark knight",
	Health:           350,
	Skill_points:     200,
	Max_skill_points: 200,
	Max_health:       350,
}

var Golem = enemy{
	Name:             "Stone golem",
	Health:           750,
	Skill_points:     0,
	Max_skill_points: 0,
	Max_health:       750,
}

var Goblin = enemy{
	Name:             "Goblin",
	Health:           80,
	Skill_points:     20,
	Max_skill_points: 20,
	Max_health:       80,
}

var enemy_input int

// The block below currently only contains the enemy's turn, might expend upon this later
func (e *enemy) Enemy_turn() {
	fmt.Println(e.Name, "'s turn")
	fmt.Println("")
	enemy_input = rand.Intn(3)

	if e.Health >= 1 {
		switch enemy_input {

		case 0:
			enemy_input = rand.Intn(3)

			switch enemy_input {

			case 0:
				e.Enemy_skill_strike(&Player)

			case 1:
				e.Enemy_skill_strike(&Pilgrim)

			case 2:
				e.Enemy_skill_strike(&Fie)

			case 3:
				e.Enemy_skill_strike(&Jessy)
			}

		case 1:
			e.Enemy_skill_heal()

		case 2:
			enemy_input = rand.Intn(3)

			switch enemy_input {

			case 0:
				e.Enemy_skill_force(&Player)

			case 1:
				e.Enemy_skill_force(&Pilgrim)

			case 2:
				e.Enemy_skill_force(&Fie)

			case 3:
				e.Enemy_skill_force(&Jessy)
			}

		case 3:
			enemy_input = rand.Intn(3)

			switch enemy_input {

			case 0:
				e.Enemy_skill_smash(&Player)

			case 1:
				e.Enemy_skill_smash(&Pilgrim)

			case 2:
				e.Enemy_skill_smash(&Fie)

			case 3:
				e.Enemy_skill_smash(&Jessy)
			}
		}
		e.Normalize_stats_enemy()
	}
	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln(&user_input)
}

// The block below is for storing things that need to be checked (and changed?) in regards to the enemies

// Reset enemy healt and skill points to max
// this needs to happen because i cant call an instance of an enemy yet
func Reset_enemy(e *enemy) {
	e.Health = e.Max_health
	e.Skill_points = e.Max_skill_points
}

// Check if enemy e Health depleted and declare victory
func (e *enemy) Check_enemy_life() {
	if e.Health <= 0 {
		victory = true

		e.check_victory()
	}
}

// Validate limits on enemy Health
// TODO put these checks inline?
func (e *enemy) Normalize_stats_enemy() {
	if e.Health > e.Max_health {
		e.Health = e.Max_health
	}
	if e.Skill_points > e.Max_skill_points {
		e.Skill_points = e.Max_skill_points
	}
}

// The block below is the place to store all skills that the enemies can use

// Heal an enemy by adding 50 plus a random value between 0 and 50 to Health
func (e *enemy) Enemy_skill_heal() {
	heal := rand.Intn(50) + 50 //amount healed
	fmt.Println(e.Name, "has healed")
	e.Health += heal
	fmt.Println(heal, "Healed")
}

// Enemy e strikes player p and decreases player Health using chance
// TODO could be interesting remodelling this formula and logic
func (e *enemy) Enemy_skill_strike(p *player) {
	fmt.Println(e.Name, "used strike")
	damage := rand.Intn(10) + 50 - p.Endurance
	critical_damage := rand.Intn(20) + 60 - p.Endurance

	if rand.Intn(100) > p.Agility {

		if rand.Intn(11) == 9 { //Critical hit chance
			p.Health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		} else {
			p.Health -= damage
			fmt.Println(damage, "DMG")
		}
	} else {
		fmt.Println("But it missed!")
	}
}

// Enemy e casts force to player p causing damage or critical damage
// TODO better description of the logic
func (e *enemy) Enemy_skill_force(p *player) {

	damage := rand.Intn(10) + 70 - p.Endurance - equipment_magic_defense
	critical_damage := rand.Intn(20) + 80 - p.Endurance - equipment_magic_defense

	fmt.Println(e.Name, "cast force")

	if e.Skill_points >= 20 {

		e.Skill_points -= 20

		if rand.Intn(100) >= p.Agility {

			if rand.Intn(3) == 1 { //Critical hit chance
				p.Health -= critical_damage
				fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
			} else {
				p.Health -= damage
				fmt.Println(damage, "DMG")
			}
		} else {
			fmt.Println("but it missed")
		}
	} else {
		fmt.Println("but nothing happened...")
		damage = 0
		p.Health -= damage
		fmt.Println(damage, "DMG")
		fmt.Scanln()
	}
}

func (e *enemy) Enemy_skill_smash(p *player) {
	fmt.Println(e.Name, "used Smash")
	damage := rand.Intn(10) + 60 - p.Endurance
	critical_damage := rand.Intn(10) + 70 - p.Endurance

	if rand.Intn(100) > p.Agility {

		if rand.Intn(11) == 9 { //Critical hit chance
			p.Health -= critical_damage
			fmt.Println(critical_damage, "DMG / CRITICAL HIT!!")
		} else {
			p.Health -= damage
			fmt.Println(damage, "DMG")
		}
	} else {
		fmt.Println("But it missed!")
	}
}
