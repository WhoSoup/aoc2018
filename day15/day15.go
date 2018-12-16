package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type tile struct {
	x, y     int
	children []*tile
	unit     *unit
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

type meta struct {
	tile   *tile
	action int // 0 up, 1 = left, 2 = right, 3 = down
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

func (u unit) coord() string {
	return u.pos.coord()
}

// world's worst breadth first search
func (u unit) closest() (*unit, *tile, int) {
	tovisit := []*tile{u.pos}

	meta := make(map[*tile]*tile)
	meta[u.pos] = nil

	queued := make(map[string]bool)

	for {
		if len(tovisit) == 0 {
			break
		}
		t := tovisit[0]
		tovisit = tovisit[1:]
		queued[t.coord()] = true

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
			_, v := queued[t.children[c].coord()]
			if !v {
				tovisit = append(tovisit, t.children[c])
				noadd(&meta, t.children[c], t)
				queued[t.children[c].coord()] = true
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
	for ap := 3; ; ap++ {
		res, deaths := game(ap)
		if ap == 3 {
			fmt.Println("Part One:", res)
		}
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
		if n := getNode(tmp, tmp[i].x, tmp[i].y-1); n != nil {
			tmp[i].children = append(tmp[i].children, n)
		}
		if n := getNode(tmp, tmp[i].x-1, tmp[i].y); n != nil {
			tmp[i].children = append(tmp[i].children, n)
		}
		if n := getNode(tmp, tmp[i].x+1, tmp[i].y); n != nil {
			tmp[i].children = append(tmp[i].children, n)
		}
		if n := getNode(tmp, tmp[i].x, tmp[i].y+1); n != nil {
			tmp[i].children = append(tmp[i].children, n)
		}
	}

	i := -1
	for {
		i++
		move(units)

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

	return (i * sum), elfdeaths
}

func (u unit) canAttack() (can bool) {
	for _, child := range u.pos.children {
		if child != nil && child.unit != nil && child.unit.team != u.team {
			return true
		}
	}
	return
}

func move(units units) {
	sort.Sort(units) // sort units by reading order
	for _, u := range units {
		if u.hp <= 0 {
			continue
		}

		if !u.canAttack() { // move
			c, next, d := u.closest()
			if c != nil && d > 1 {
				u.pos.unit = nil
				next.unit = u
				u.pos = next
			}
		}

		var target *tile
		for _, child := range u.pos.children {
			if child == nil {
				continue
			}
			if child.unit != nil && child.unit.team != u.team && (target == nil || child.unit.hp < target.unit.hp) {
				target = child
			}
		}

		if target != nil {
			target.unit.hp -= u.ap
			if target.unit.hp < 0 {
				target.unit = nil
			}
		}
	}
}
