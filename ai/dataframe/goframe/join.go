package main

import (
	"fmt"
	"log"

	"github.com/kishyassin/goframe"
)

func main() {
	df1 := goframe.NewDataFrame()
	df1.AddColumn(goframe.ConvertToAnyColumn(goframe.NewColumn("id", []int{1, 2, 3})))
	df1.AddColumn(goframe.ConvertToAnyColumn(goframe.NewColumn("value1", []string{"A", "B", "C"})))

	df2 := goframe.NewDataFrame()
	df2.AddColumn(goframe.ConvertToAnyColumn(goframe.NewColumn("id", []int{2, 3, 4})))
	df2.AddColumn(goframe.ConvertToAnyColumn(goframe.NewColumn("value2", []string{"X", "Y", "Z"})))

	// Perform an inner join
	joined, err := df1.InnerJoin(df2, "id")
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(df1)
	fmt.Println(df2)
	fmt.Println(joined)
}