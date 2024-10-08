package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func main() {
	itemStyle := lipgloss.NewStyle().MarginRight(1)
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).MarginRight(1)

	t := tree.Root("Groceries").
		Child(
			tree.Root("Fruits").
				Child(
					"Blood Orange",
					"Papaya",
					"Dragonfruit",
					"Yuzu",
				),
			tree.Root("Items").
				Child(
					"Cat Food",
					"Nutella",
					"Powdered Sugar",
				),
			tree.Root("Veggies").
				Child(
					"Leek",
					"Artichoke",
				),
		).ItemStyle(itemStyle).EnumeratorStyle(enumeratorStyle).Enumerator(tree.RoundedEnumerator)

	fmt.Println(t)
}
