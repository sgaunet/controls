package reportmd

import (
	"fmt"
	"time"

	"github.com/atsushinee/go-markdown-generator/doc"
)

func New() *reportMd {
	r := reportMd{report: doc.NewMarkDown()}
	return &r
}

func (r *reportMd) AddTable(title string, table [][]string) {
	var nbColumns int
	nbRows := len(table)
	if nbRows != 0 {
		nbColumns = len(table[0])
	}

	//fmt.Println("Table ", nbRows, nbColumns)
	t := doc.NewTable(nbRows, nbColumns)
	for idx, line := range table {
		//fmt.Println("line=", line)
		for idxRow, value := range line {
			if idx == 0 {
				t.SetTitle(idxRow, line[idxRow])
				//fmt.Println("SetTitle", idxRow, line[idxRow])
			} else {
				//fmt.Println(idxRow, value)
				t.SetContent(idx, idxRow, value)
			}
		}
	}

	r.report.WriteTitle(title, 3)
	r.report.WriteTable(t)
	r.report.WriteLines(2)
}

func (r *reportMd) AddTitle(title string) {
	r.report.WriteTitle(title, doc.LevelTitle).WriteLines(2)
}

func (r *reportMd) AddSection(title string) {
	r.report.WriteLines(2)
	r.report.WriteTitle(title, 2)
	r.report.WriteLines(1)
	// StdOut
	fmt.Printf("%s :\n\n", title)
}

func (r *reportMd) Export(reportpath string) error {
	return r.report.Export(reportpath)
}

func (r *reportMd) AddFooter() {
	r.report.WriteTitle("Infos", 2)
	r.report.WriteLines(1)
	r.report.Write("Generated at :" + time.Now().String() + "<br>")
}

func (r *reportMd) AddPAgeBreak() {
	r.report.Write("<div style = \"display:block; clear:both; page-break-after:always;\"></div>")
}
