package reportpdf

import (
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/sgaunet/controls/results"
)

func New() *reportPdf {
	r := reportPdf{report: pdf.NewMaroto(consts.Portrait, consts.A4)}
	r.report.SetPageMargins(10, 15, 10)
	return &r
}

func (r *reportPdf) AddTable(title string, results []results.Result) {
	grayColor := getGrayColor()
	var allOk bool
	allOk = true
	redColor := getRedColor()
	greenColor := getGreenColor()

	r.AddSection(title)
	var contents [][]string
	header := []string{"Test", "Result"}

	for _, result := range results {
		if result.Pass {
			contents = append(contents, []string{result.Title, "ok"})
		} else {
			contents = append(contents, []string{result.Title, result.Result})
			allOk = false
		}
	}
	if len(results) == 0 {
		contents = append(contents, []string{"No problems to show", "ok"})
	}

	if allOk {
		r.report.SetBackgroundColor(greenColor)
	} else {
		r.report.SetBackgroundColor(redColor)
	}
	r.report.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{8, 4},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{8, 4},
		},
		Align:                consts.Left,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})
}

func (r *reportPdf) AddTitle(title string) {
	r.report.Row(10, func() {
		r.report.Col(12, func() {
			r.report.Text(title, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  20.0,
			})
		})
	})
}

func (r *reportPdf) AddSection(title string) {
	r.report.Row(10, func() {
		r.report.Col(12, func() {
			r.report.Text(title, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  15.0,
			})
		})
	})
}

func (r *reportPdf) AddSubSection(title string) {
	r.report.Row(10, func() {
		r.report.Col(12, func() {
			r.report.Text(title, props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  10.0,
			})
		})
	})
}

func (r *reportPdf) AddLine() {
	r.report.Line(10)
}

func (r *reportPdf) Export(reportpath string) error {
	return r.report.OutputFileAndClose(reportpath)
}

func (r *reportPdf) AddFooter(version string) {

	blueColor := getBlueColor()

	var host string
	generated := "Generated at: " + time.Now().Format("02-Jan-2006 15:04:05")
	versionStr := "Version : " + version
	hostname, err := os.Hostname()
	if err == nil {
		host = "Generated on: " + hostname
	}

	r.report.RegisterFooter(func() {
		r.report.Row(20, func() {
			r.report.Col(12, func() {
				r.report.Text(generated, props.Text{
					Top:   13,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
				r.report.Text(host, props.Text{
					Top:   16,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
				r.report.Text(versionStr, props.Text{
					Top:   19,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
					Color: blueColor,
				})
			})
		})
	})
}

func (r *reportPdf) AddPAgeBreak() {
	r.report.AddPage()
}
