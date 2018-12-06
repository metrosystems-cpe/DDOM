package helpers

import (
	"bytes"
	"fmt"

	tm "github.com/buger/goterm"
)

// PanicIfError just panic
func PanicIfError(err error, message string) {
	if err != nil {
		panic(message)
	}
}

// LogError prints in STDOUT the error
func LogError(err error) {
	if err != nil {
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

// BuildPrintableTable builds and returns a printable table of objects
func BuildPrintableTable(colNames []string, rows [][]interface{}) *tm.Table {
	table := tm.NewTable(0, 10, 5, ' ', 0)

	fmt.Fprintf(table, buildRowString(colNames))

	var pattern bytes.Buffer
	for i := 0; i < len(colNames); i++ {
		if i < len(colNames)-1 {
			pattern.WriteString("%s\t")
		} else {
			pattern.WriteString("%s\n")
		}
	}

	for _, row := range rows {
		fmt.Fprintf(table, pattern.String(), row...)
	}

	return table
}
