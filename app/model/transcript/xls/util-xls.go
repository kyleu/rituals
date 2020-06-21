package xls

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func setColumnHeaders(key string, cols []string, f *excelize.File) {
	c := 'A'
	for _, col := range cols {
		f.SetCellValue(key, string(c)+"1", col)
		c++
	}
}

func setData(key string, firstRow int, data [][]interface{}, f *excelize.File) {
	for rowIdx, row := range data {
		for colIdx, col := range row {
			f.SetCellValue(key, fmt.Sprint(string('A'+colIdx), rowIdx+firstRow), col)
		}
	}
}

func setFirstSheetTitle(t string, f *excelize.File) {
	f.SetSheetName(defSheet, t)
}

func setColumnWidths(key string, widths []int, f *excelize.File) {
	for i, w := range widths {
		col := string('A' + i)
		f.SetColWidth(key, col, col, float64(w))
	}
}
