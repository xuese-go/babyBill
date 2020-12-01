package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/xuese-go/babyBill/service"
	"log"
	"os"
	"strconv"
)

func init() {
	file, _ := os.Create("sys.log")
	log.SetOutput(file) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
func main() {
	var inDE *walk.DateEdit
	var inNUM *walk.NumberEdit
	var inTE *walk.LineEdit

	_, _ = MainWindow{
		Title:   "宝宝简易记账",
		MinSize: Size{600, 400},
		Layout: VBox{
			MarginsZero: true,
			SpacingZero: true,
			Spacing:     0,
		},
		Children: []Widget{
			Composite{
				Layout: Grid{
					Columns: 7,
				},
				Children: []Widget{
					Label{
						Text: "日期:",
					},
					DateEdit{
						Date:     Bind("日期"),
						AssignTo: &inDE,
					},
					Label{
						Text: "金额:",
					},
					NumberEdit{
						Value:    Bind("金额", Range{Min: 0.01, Max: 9999.99}),
						Suffix:   " 元",
						Decimals: 2,
						MinValue: 0.00,
						MaxValue: 9999.99,
						MinSize: Size{
							Width: 200,
						},
						MaxSize: Size{
							Width: 300,
						},
						AssignTo: &inNUM,
					},
					Label{
						Text: "事项:",
					},
					LineEdit{
						Text:      Bind("事项"),
						MaxLength: 100,
						MinSize: Size{
							Width: 300,
						},
						AssignTo: &inTE,
					},
					PushButton{
						Text: "确定",
						OnClicked: func() {
							dates := inDE.Date().Format("2006-01-02 15:04:05")
							money := strconv.FormatFloat(inNUM.Value(), 'E', 2, 64)
							matter := inTE.Text()
							if err := service.Save(dates, money, matter); err != nil {
								log.Panicln(err.Error())
							}
						},
					},
				},
			},
			VSplitter{
				MinSize: Size{
					Height: 390,
				},
				Children: []Widget{
					TableView{
						Columns: []TableViewColumn{
							{Title: "#"},
							{Title: "日期", Format: "2006-01-02 15:04:05", Width: 150},
							{Title: "金额（元）", Alignment: AlignFar},
							{Title: "事项", Alignment: AlignFar},
						},
					},
				},
			},
		},
	}.Run()
}
