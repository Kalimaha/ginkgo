package reporter

import (
	"fmt"
	"strings"
)

func PrintLeaves(m map[string]Leaf) {
	for _, leaf1 := range m {
		printLine(leaf1)
		for _, leaf2 := range leaf1.Leaves {
			printLine(leaf2)
			for _, leaf3 := range leaf2.Leaves {
				printLine(leaf3)
				for _, leaf4 := range leaf3.Leaves {
					printLine(leaf4)
				}
			}
		}
	}
}

func printLine(leaf Leaf) {
	if len(leaf.Leaves) == 0 {
		fmt.Println(formatOut(leaf.Description, leaf.Level), icon(leaf.Passed), duration(leaf.Duration))
		if leaf.Passed == -1 {
			fmt.Println(formatOut(fmt.Sprintf("  %s", leaf.Link), leaf.Level))
		}
	} else {
		fmt.Println(formatOut(leaf.Description, leaf.Level))
	}
}

func formatOut(msg string, level int) string {
	blanks := strings.Builder{}
	for i := 0; i < level; i++ {
		blanks.WriteString("  ")
	}
	blanks.WriteString("%s")

	return fmt.Sprintf(blanks.String(), msg)
}

func duration(duration string) string {
	return fmt.Sprintf("[%s]", duration)
}