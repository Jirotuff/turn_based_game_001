package main

import (
	"fmt"
	"math/rand"
)

type enemy struct {
	name             string
	health           int
	skill_points     int
	max_skill_points int
	max_health       int
}

var Bandit = enemy{
	name:             "Bandit",
	health:           300,
	skill_points:     200,
	max_skill_points: 200,
	max_health:       300,
}

var enemy_input int

// The block below currently only contains the enemy's turn, might expend upon this later

func (e *enemy) Enemy_turn() {
	fmt.Println(e.name, "'s turn")
	fmt.Println("")
	enemy_input = rand.Intn(3)

	if e.health >= 1 {
		switch enemy_input {

		case 0:
			enemy_input = rand.Intn(3)

			switch enemy_input {

			case 0:
				e.Enemy_skill_strike(&Dario)

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
				e.Enemy_skill_force(&Dario)

			case 1:
				e.Enemy_skill_force(&Pilgrim)

			case 2:
				e.Enemy_skill_force(&Fie)

			case 3:
				e.Enemy_skill_force(&Jessy)
			}
		}
		e.Normalize_stats_enemy()
	}
}

// The block below is for storing things that need to be checked (and changed?) in regards to the enemies

// Reset enemy healt and skill points to max
// TODO is skill supposed to start at max?
func Reset_enemy(e *enemy) {
	e.health = e.max_health
	e.skill_points = e.max_skill_points
}

// Check if enemy e health depleted and declare victory
func (e *enemy) Check_enemy_life() {
	if e.health <= 0 {
		victory = true
		check_victory()
	}
}

// Validate limits on enemy health
// TODO put these checks inline?
func (e *enemy) Normalize_stats_enemy() {
	if e.health > e.max_health {
		e.health = e.max_health
	}
	if e.skill_points > e.max_skill_points {
		e.skill_points = e.max_skill_points
	}
}

// The block below is the place to store all skills that the enemies can use

// Heal an enemy by adding 50 plus a random value between 0 and 50 to health
func (e *enemy) Enemy_skill_heal() {
	heal := rand.Intn(50) + 50 //amount healed
	fmt.Println(e.name, "has healed")
	e.health += heal
	fmt.Println(heal, "Healed")
}

// Enemy e strikes player p and decreases player health using chance
// TODO could be interesting remodelling this formula and logic
func (e *enemy) Enemy_skill_strike(p *player) {
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

// Enemey e applies force to player p causing random damage
// TODO better description of the logic
func (e *enemy) Enemy_skill_force(p *player) {

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
