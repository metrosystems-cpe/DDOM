package helpers

import (
	"bytes"
	"fmt"

	tm "github.com/buger/goterm"
)

func PanicIfError(err error, message string) {
	if err != nil {
		panic(message)
	}
}

func LogError(err error) {
	if err != nil {
		// do things
		fmt.Println(err.Error())
	}
}

func buildRowString(colNames []string) string {
	var buffer bytes.Buffer
	size := len(colNames)
	for idx, name := range colNames {
		buffer.WriteString(name)
		if idx != size-1 {
			buffer.WriteString("\t")
		} else {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

func BuildPrintableTable(colNames []string, rows [][]interface{}) *tm.Table {
	table := tm.NewTable(0, 10, 5, ' ', 0)

	fmt.Fprintf(table, buildRowString(colNames))

	for _, row := range rows {
		fmt.Fprintf(table, "%s\t%s\t%s\n", row...)
	}

	return table
}
