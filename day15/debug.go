package main

import "fmt"

func printgame(tmp map[string]*tile, sx, sy int) {
	for y := 0; y < sy; y++ {
		hp := ""
		for x := 0; x < sx; x++ {
			t, ok := tmp[fmt.Sprintf("%d,%d", x, y)]
			if ok {
				if t.unit != nil {
					ts := "E"
					if t.unit.team {
						ts = "G"
					}
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
