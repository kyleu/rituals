package pdf

import (
	"github.com/johnfercher/maroto/pkg/consts"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func tr(content func(), height float64, m pdfgen.Maroto) {
	m.Row(height, content)
}

func col(content func(), width uint, m pdfgen.Maroto) {
	m.Col(width, content)
}

func th(content string, width uint, m pdfgen.Maroto) {
	col(func() { m.Text(content, props.Text{Style: consts.Bold}) }, width, m)
}

func td(content string, width uint, m pdfgen.Maroto) {
	col(func() { m.Text(content) }, width, m)
}

func table(cols []string, data [][]string, sizes []uint, m pdfgen.Maroto) {
	m.TableList(cols, data, props.TableList{
		HeaderProp:         props.TableListContent{Style: consts.Bold, GridSizes: sizes},
		ContentProp:        props.TableListContent{GridSizes: sizes},
		HeaderContentSpace: 2,
	})
}

func hr(m pdfgen.Maroto) {
	m.Line(12)
}

func caption(title string, m pdfgen.Maroto) {
	tr(func() { col(func() { m.Text(title, props.Text{Size: 12}) }, 12, m) }, 10, m)
}

func detailRow(k string, v string, m pdfgen.Maroto) {
	tr(func() {
		th(k, 2, m)
		td(v, 10, m)
	}, 8, m)
}
