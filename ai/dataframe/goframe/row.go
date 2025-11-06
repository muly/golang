package main

import (
	"fmt"

	"github.com/kishyassin/goframe"
)

func main() {
	df := goframe.NewDataFrame()
	df.AddColumn(goframe.ConvertToAnyColumn(goframe.NewColumn("name", []string{"Alice", "Bob", "Charlie"})))
	df.AddColumn(goframe.ConvertToAnyColumn(goframe.NewColumn("age", []int{25, 30, 35})))

	// Access a row
	row, _ := df.Row(1)
	fmt.Println("Row 1:", row)

	// Get the first two rows
	head := df.Head(2)
	fmt.Println("Head:", head)

	// Append a new row
	df.AppendRow(df, map[string]any{"name": "Diana", "age": 40})
	fmt.Println("After appending a row:", df)

	// Drop a row
	df.DropRow(1)
	fmt.Println("After dropping a row:", df)
}
