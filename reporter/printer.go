package reporter

import "fmt"

func PrintLeaves(leaves []Leaf) {
	for _, v1 := range leaves {
		printLine(v1)
		for _, v2 := range v1.Leaves {
			printLine(v2)
			for _, v3 := range v2.Leaves {
				printLine(v3)
				for _, v4 := range v3.Leaves {
					printLine(v4)
				}
			}
		}
	}
}

func printLine(leaf Leaf) {
	if len(leaf.Leaves) == 0 {
		fmt.Println(leaf.Description, icon(leaf.Passed))
	} else {
		fmt.Println(leaf.Description)
	}
}
