package main

import (
	"fmt"
	"sort"
)

/*
Immune System:
2987 units each with 5418 hit points (immune to slashing; weak to cold, bludgeoning) with an attack that does 17 cold damage at initiative 5
1980 units each with 9978 hit points (immune to cold) with an attack that does 47 cold damage at initiative 19
648 units each with 10733 hit points (immune to radiation, fire, slashing) with an attack that does 143 fire damage at initiative 9
949 units each with 3117 hit points with an attack that does 29 fire damage at initiative 10
5776 units each with 5102 hit points (weak to cold; immune to slashing) with an attack that does 8 radiation damage at initiative 15
1265 units each with 4218 hit points (immune to radiation) with an attack that does 24 radiation damage at initiative 16
3088 units each with 10066 hit points (weak to slashing) with an attack that does 28 slashing damage at initiative 1
498 units each with 1599 hit points (immune to bludgeoning; weak to radiation) with an attack that does 28 bludgeoning damage at initiative 11
3705 units each with 10764 hit points with an attack that does 23 cold damage at initiative 7
3431 units each with 3666 hit points (weak to slashing; immune to bludgeoning) with an attack that does 8 bludgeoning damage at initiative 8

Infection:
2835 units each with 33751 hit points (weak to cold) with an attack that does 21 bludgeoning damage at initiative 13
4808 units each with 32371 hit points (weak to radiation; immune to bludgeoning) with an attack that does 11 cold damage at initiative 14
659 units each with 30577 hit points (weak to fire; immune to radiation) with an attack that does 88 slashing damage at initiative 12
5193 units each with 40730 hit points (immune to radiation, fire, bludgeoning; weak to slashing) with an attack that does 14 cold damage at initiative 20
1209 units each with 44700 hit points (weak to bludgeoning, radiation) with an attack that does 71 fire damage at initiative 18
6206 units each with 51781 hit points (immune to cold) with an attack that does 13 fire damage at initiative 4
602 units each with 22125 hit points (weak to radiation, bludgeoning) with an attack that does 73 cold damage at initiative 3
5519 units each with 37123 hit points (weak to slashing, fire) with an attack that does 12 radiation damage at initiative 2
336 units each with 23329 hit points (weak to fire; immune to cold, bludgeoning, radiation) with an attack that does 134 cold damage at initiative 17
2017 units each with 50511 hit points (immune to bludgeoning) with an attack that does 42 fire damage at initiative 6

Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with
 an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning,
 slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack
 that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire,
 cold) with an attack that does 12 slashing damage at initiative 4

*/

const (
	NONE        = 0
	SLASHING    = 1
	COLD        = 2
	BLUDGEONING = 4
	RADIATION   = 8
	FIRE        = 16
)

type unit struct {
	units, hp, immune, weakness, attack, attackType, initiative int
}
type army []*unit

func (u unit) effectivePower() int {
	return u.units * u.attack
}

func (u unit) isImmune(t int) bool {
	return u.immune&t > 0
}
func (u unit) isWeak(t int) bool {
	return u.weakness&t > 0
}

func (u unit) damage(other *unit) int {
	if other == nil || other.units <= 0 {
		return 0
	}
	if other.isImmune(u.attackType) {
		return 0
	}
	if other.isWeak(u.attackType) {
		return 2 * u.attack * u.units
	}
	return u.attack * u.units
}

func (u *unit) suffer(damage int) {
	tmp := damage / u.hp
	if tmp > u.units {
		tmp = u.units
	}
	//fmt.Printf(" taking %d damage\n", tmp)
	u.units -= damage / u.hp
	if u.units < 0 {
		u.units = 0
	}
}

func (a army) Len() int      { return len(a) }
func (a army) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a army) Less(i, j int) bool {
	if a[i].effectivePower() == a[j].effectivePower() {
		return a[i].initiative > a[j].initiative
	}
	return a[i].effectivePower() > a[j].effectivePower()
}

type byInitiative army

func (b byInitiative) Len() int      { return len(b) }
func (b byInitiative) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byInitiative) Less(i, j int) bool {
	return b[i].initiative > b[j].initiative
}

func (a army) targetMatrix(other army) (matrix map[*unit]*unit) {
	sort.Sort(a)
	sort.Sort(other)

	matrix = make(map[*unit]*unit)

	for i := range a {
		if a[i].units < 1 {
			continue
		}
		var target *unit
		for j := range other {
			if other[j].units <= 0 {
				continue
			}

			// other unit has no match yet
			if _, ok := matrix[other[j]]; !ok {
				dmg := a[i].damage(other[j])
				//fmt.Printf("group %d would deal defending group %d %d damage\n", i, j, dmg)
				if dmg > a[i].damage(target) {
					target = other[j]
				}
			}
		}

		if target != nil {
			matrix[target] = a[i]
		}
	}
	return
}

func mmerge(a map[*unit]*unit, b map[*unit]*unit) (c map[*unit]*unit) {
	c = make(map[*unit]*unit)
	for k, v := range a {
		c[v] = k
	}
	for k, v := range b {
		c[v] = k
	}
	return
}

func attack(a army, matrix map[*unit]*unit) {

	sort.Sort(byInitiative(a))

	for _, u := range a {
		if u.units > 0 && matrix[u] != nil {
			//fmt.Printf("%v attacking %v ", u, matrix[u])
			matrix[u].suffer(u.damage(matrix[u]))
		}
	}
}

func (a army) units() (sum int) {
	for _, v := range a {
		sum += v.units
	}
	return
}

func (a army) hasUnits() bool {
	return a.units() > 0
}

/*
IMMUNE, WEAKNESS
Immune System:
&unit{2987, 5418, SLASHING, COLD|BLUDGEONING, 17, COLD, 5},
&unit{1980, 9978, COLD, NONE, 47, COLD, 19},
&unit{648, 10733, RADIATION|FIRE|SLASHING, NONE, 143, FIRE, 9},
&unit{949, 3117, NONE, NONE, 29, FIRE, 10},
&unit{5776, 5102, SLASHING, COLD, 8, RADIATION, 15},
&unit{1265, 4218, RADIATION, NONE, 24, RADIATION, 16},
&unit{3088, 10066, NONE, SLASHING, 28, SLASHING, 1},
&unit{498, 1599, BLUDGEONING, RADIATION, 28, BLUDGEONING, 11},
&unit{3705, 10764, NONE, NONE, 23, COLD, 7},
&unit{3431, 3666, BLUDGEONING, SLASHING, 8, BLUDGEONING, 8}

Infection:
&unit{2835, 33751, NONE, COLD, 21, BLUDGEONING, 13},
&unit{4808, 32371, BLUDGEONING, RADIATION, 11, COLD, 14},
&unit{659, 30577, RADIATION, FIRE, 88, SLASHING, 12},
&unit{5193, 40730, RADIATION|FIRE|BLUDGEONING, SLASHING, 14, COLD, 20},
&unit{1209, 44700, NONE, BLUDGEONING|RADIATION, 71, FIRE, 18},
&unit{6206, 51781, COLD, NONE, 13, FIRE, 4},
&unit{602, 22125, NONE, RADIATION|BLUDGEONING, 73, COLD, 3},
&unit{5519, 37123, NONE, SLASHING|FIRE, 12, RADIATION, 2},
&unit{336, 23329, COLD|BLUDGEONING|RADIATION, FIRE, 134, COLD, 17},
&unit{2017, 50511, BLUDGEONING, NONE, 42, FIRE, 6}

	immuneSystem := army{
		&unit{17, 5390, NONE, RADIATION | BLUDGEONING, 4507, FIRE, 2},
		&unit{989, 1274, FIRE, BLUDGEONING | SLASHING, 25, SLASHING, 3},
	}
	infection := army{
		&unit{801, 4706, NONE, RADIATION, 116, BLUDGEONING, 1},
		&unit{4485, 2961, RADIATION, FIRE | COLD, 12, SLASHING, 4},
	}

*/

func battle(boost int) (int, int) {

	immuneSystem := army{
		&unit{2987, 5418, SLASHING, COLD | BLUDGEONING, 17, COLD, 5},
		&unit{1980, 9978, COLD, NONE, 47, COLD, 19},
		&unit{648, 10733, RADIATION | FIRE | SLASHING, NONE, 143, FIRE, 9},
		&unit{949, 3117, NONE, NONE, 29, FIRE, 10},
		&unit{5776, 5102, SLASHING, COLD, 8, RADIATION, 15},
		&unit{1265, 4218, RADIATION, NONE, 24, RADIATION, 16},
		&unit{3088, 10066, NONE, SLASHING, 28, SLASHING, 1},
		&unit{498, 1599, BLUDGEONING, RADIATION, 28, BLUDGEONING, 11},
		&unit{3705, 10764, NONE, NONE, 23, COLD, 7},
		&unit{3431, 3666, BLUDGEONING, SLASHING, 8, BLUDGEONING, 8},
	}

	for _, k := range immuneSystem {
		k.attack += boost
	}

	infection := army{
		&unit{2835, 33751, NONE, COLD, 21, BLUDGEONING, 13},
		&unit{4808, 32371, BLUDGEONING, RADIATION, 11, COLD, 14},
		&unit{659, 30577, RADIATION, FIRE, 88, SLASHING, 12},
		&unit{5193, 40730, RADIATION | FIRE | BLUDGEONING, SLASHING, 14, COLD, 20},
		&unit{1209, 44700, NONE, BLUDGEONING | RADIATION, 71, FIRE, 18},
		&unit{6206, 51781, COLD, NONE, 13, FIRE, 4},
		&unit{602, 22125, NONE, RADIATION | BLUDGEONING, 73, COLD, 3},
		&unit{5519, 37123, NONE, SLASHING | FIRE, 12, RADIATION, 2},
		&unit{336, 23329, COLD | BLUDGEONING | RADIATION, FIRE, 134, COLD, 17},
		&unit{2017, 50511, BLUDGEONING, NONE, 42, FIRE, 6},
	}

	i := 0
	for immuneSystem.hasUnits() && infection.hasUnits() {
		matrix1 := infection.targetMatrix(immuneSystem)
		matrix2 := immuneSystem.targetMatrix(infection)
		attack(append(infection, immuneSystem...), mmerge(matrix1, matrix2))
		if i > 50000 {
			return -1, -1
		}
		i++
	}

	return immuneSystem.units(), infection.units()
}

func main() {
	_, p1b := battle(0)
	fmt.Println(p1b)

	for i := 85; ; i++ {
		if z, _ := battle(i); z > 0 {
			fmt.Println(i, z)
			break
		}
	}
}
