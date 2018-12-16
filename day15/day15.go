package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tile struct {
	x, y     int
	children []*tile
	unit     *unit
}

type units []*unit

type unit struct {
	team   bool
	pos    *tile
	hp, ap int
}

func (u units) Len() int      { return len(u) }
func (u units) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u units) Less(i, j int) bool {
	if u[i].pos.y < u[j].pos.y {
		return true
	} else if u[i].pos.y == u[j].pos.y {
		return u[i].pos.x < u[j].pos.x
	}
	return false
}

func (t tile) coord() string {
	return fmt.Sprintf("%d,%d", t.x, t.y)
}

func getNode(m map[string]*tile, x, y int) *tile {
	e, ok := m[fmt.Sprintf("%d,%d", x, y)]
	if ok {
		return e
	}
	return nil
}

func (u unit) coord() string {
	return u.pos.coord()
}

type meta struct {
	tile   *tile
	action int // 0 up, 1 = left, 2 = right, 3 = down
}

// world's worst breadth first search
func (u unit) closest() (*unit, *tile, int) {
	tovisit := []*tile{u.pos}

	meta := make(map[*tile]*tile)
	meta[u.pos] = nil

	visited := make(map[string]bool)

	z := 0
	for {
		if len(tovisit) == 0 {
			break
		}
		z++
		t := tovisit[0]
		tovisit = tovisit[1:]
		visited[t.coord()] = true

		if t != u.pos && t.unit != nil {
			if t.unit.team != u.team { // unit found
				dist := 0
				unit := t.unit
				for {
					tmp, _ := meta[t]
					if tmp == u.pos {
						return unit, t, dist + 1
					}
					dist++
					t = tmp
				}
				//return nil, t.unit
			}
			continue // same team unit, treat as wall
		}

		for c := range t.children {
			if t.children[c] == nil {
				continue
			}
			_, v := visited[t.children[c].coord()]
			if !v {
				tovisit = append(tovisit, t.children[c])
				noadd(&meta, t.children[c], t)
				visited[t.children[c].coord()] = true
			}
		}
	}
	return nil, nil, -1
}

func noadd(meta *map[*tile]*tile, a, b *tile) {
	_, ok := (*meta)[a]
	if !ok {
		(*meta)[a] = b
	}
}

func main() {
	p1, _ := game(3)
	fmt.Println("Part One:", p1)

	for ap := 4; ; ap++ {
		res, deaths := game(ap)
		//fmt.Println("Elf AP:", ap, "=", res, deaths)
		if deaths == 0 {
			fmt.Println("Part Two:", res)
			break
		}
	}
}

func game(ap int) (int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var units units
	//var tmp []tile
	tmp := make(map[string]*tile, 64)
	var tiles []*tile

	y, x := 0, 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		for x = 0; x < len(scanner.Text()); x++ {
			b := scanner.Text()[x]
			if b == '#' {
				continue
			}

			tile := &tile{x, y, []*tile{nil, nil, nil, nil}, nil}
			tiles = append(tiles, tile)
			tmp[fmt.Sprintf("%d,%d", x, y)] = tile

			if b != '.' {
				team := b == 'G'
				tap := ap
				if team {
					tap = 3
				}
				u := &unit{team, tile, 200, tap}
				units = append(units, u)
				tile.unit = u
			}
		}

		y++
	}

	for i := range tmp {
		tmp[i].children[0] = getNode(tmp, tmp[i].x, tmp[i].y-1)
		tmp[i].children[1] = getNode(tmp, tmp[i].x-1, tmp[i].y)
		tmp[i].children[2] = getNode(tmp, tmp[i].x+1, tmp[i].y)
		tmp[i].children[3] = getNode(tmp, tmp[i].x, tmp[i].y+1)
	}

	//printgame(tmp, x, y)

	i := -1
	for {
		i++
		move(tiles)

		a, b := 0, 0
		for _, c := range units {
			if c.hp > 0 {
				if c.team {
					a++
				} else {
					b++
				}
			}
		}

		if a == 0 || b == 0 {
			break
		}
	}

	sum := 0
	elfdeaths := 0
	for _, c := range units {
		if c.hp > 0 {
			sum += c.hp
		}
		if !c.team && c.hp <= 0 {
			elfdeaths++
		}
	}
	//fmt.Println("Part One:", (i * sum))

	return (i * sum), elfdeaths
}

func printgame(tmp map[string]*tile, sx, sy int) {
	for y := 0; y < sy; y++ {
		hp := ""
		for x := 0; x < sx; x++ {
			t, ok := tmp[fmt.Sprintf("%d,%d", x, y)]
			if ok {
				if t.unit != nil {
					ts := team(t.unit.team)[0:1]
					fmt.Print(ts)
					hp += fmt.Sprintf(" %s(%d)", ts, t.unit.hp)

				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println(hp)
	}
	fmt.Println()
}

func team(b bool) string {
	if b {
		return "Goblin"
	}
	return "Elf"
}

func (u unit) canAttack() (can bool) {
	for _, child := range u.pos.children {
		if child == nil {
			continue
		}
		if child.unit == nil {
			continue
		}
		if child.unit.team == u.team {
			continue
		}
		return true
	}
	return
}

func move(tmp []*tile) {
	actioned := make(map[*unit]bool)
	for _, t := range tmp {
		if t.unit == nil {
			continue
		}
		_, ok := actioned[t.unit]
		if ok {
			continue
		}
		actioned[t.unit] = true

		u := t.unit

		if !u.canAttack() { // move
			c, next, d := u.closest()
			if c != nil && d > 1 {
				//fmt.Printf("\tMove to: Coords(%d, %d)\n", next.x, next.y)
				u.pos.unit = nil
				next.unit = u
				u.pos = next
			}
		}

		//fmt.Printf("Unit Team(%s) Coords(%d, %d) Hp(%d)\n", team(u.team), u.pos.x, u.pos.y, u.hp)

		var target *tile
		for _, child := range u.pos.children {
			if child == nil {
				continue
			}
			if child.unit != nil && child.unit.team != u.team && (target == nil || child.unit.hp < target.unit.hp) {
				target = child
			}
		}

		//fmt.Printf("\tClosest: Unit Team(%s) Coords(%d, %d) Hp(%d)\n", team(c.team), c.pos.x, c.pos.y, c.hp)
		//fmt.Printf("\tAttacking (%d,%d)\n", next.x, next.y)

		if target != nil {
			target.unit.hp -= u.ap
			if target.unit.hp < 0 {
				target.unit = nil
			}
		}
	}
}
