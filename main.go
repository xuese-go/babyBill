package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	_ "github.com/xuese-go/babyBill/db"
	"github.com/xuese-go/babyBill/service"
	"log"
	"os"
	"time"
)

func init() {
	file, _ := os.Create("sys.log")
	log.SetOutput(file) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

var mw *walk.MainWindow
var tv *walk.TableView
var m = "0.00"
var d = "0.00"

/**
表格相关
*/
type Foo struct {
	Index int
	A     time.Time
	B     float64
	C     string
}

type FooModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Foo
}

func (m *FooModel) RowCount() int {
	return len(m.items)
}

// 指定单元格显示的文本
func (m *FooModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.A

	case 2:
		return item.B

	case 3:
		return item.C
	}

	panic("unexpected col")
}

func NewFooModel(d []*service.Record) *FooModel {
	m := new(FooModel)
	m.items = make([]*Foo, 0)
	//f := &Foo{
	//	Index: 0,
	//	A:     "1",
	//	B:     "2",
	//	C:     "3",
	//}
	//m.items = append(m.items, f)
	for i := range d {
		f := &Foo{
			Index: i,
			A:     d[i].Dates,
			B:     d[i].Money,
			C:     d[i].Matter,
		}
		m.items = append(m.items, f)
	}
	return m
}

func main() {
	var inDE, inDE2 *walk.DateEdit
	var inNUM *walk.NumberEdit
	var inTE *walk.LineEdit
	//表格数据
	model := NewFooModel(make([]*service.Record, 0))

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "宝宝简易记账",
		MinSize:  Size{Width: 600, Height: 400},
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
					DateEdit{
						Date:     Bind("日期"),
						AssignTo: &inDE2,
					},
					LinkLabel{
						Text: "所选月统计金额：",
					},
					LinkLabel{
						Text: m,
					},
					LinkLabel{
						Text: "所选日统计金额：",
					},
					LinkLabel{
						Text: d,
					},
					PushButton{
						Text: "查询",
						OnClicked: func() {
							if data, err := service.Find(""); err != nil {
								dlg("失败")
								log.Panicln(err)
							} else {
								da := NewFooModel(data)
								if err = tv.SetModel(da); err != nil {
									dlg("table赋值失败")
									log.Println(err)
								}
								model.PublishRowsReset()
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
						AssignTo: &tv,
						Columns: []TableViewColumn{
							{Title: "#", Alignment: AlignCenter, Width: 50},
							{Title: "日期", Alignment: AlignCenter, Width: 150, Format: "2006-01-02"},
							{Title: "金额(元)", Alignment: AlignCenter, Width: 100},
							{Title: "事项", Alignment: AlignFar},
						},
						StyleCell: func(style *walk.CellStyle) {
							if style.Row()%2 == 0 {
								style.BackgroundColor = walk.RGB(159, 215, 255)
							} else {
								style.BackgroundColor = walk.RGB(143, 199, 239)
							}
						},
						Model: model,
					},
				},
			},
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
						Text: "提交",
						OnClicked: func() {
							dates := inDE.Date()
							//money := strconv.FormatFloat(inNUM.Value(), 'E', -1, 64)
							money := inNUM.Value()
							matter := inTE.Text()
							if err := service.Save(dates, money, matter); err != nil {
								dlg("失败")
								log.Panicln(err)
							} else {
								dlg("成功")
							}
						},
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Println(err)
	}
}

/**
弹窗
*/
func dlg(str string) {
	var d *walk.Dialog
	var btn1, btn2 *walk.PushButton
	_, err := Dialog{
		AssignTo:      &d,
		DefaultButton: &btn1,
		CancelButton:  &btn2,
		MinSize:       Size{Width: 300, Height: 300},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{
					Columns: 1,
				},
				Children: []Widget{
					Label{
						Text: str,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo:  &btn1,
						Text:      "确定",
						OnClicked: func() { d.Accept() },
					},
					PushButton{
						AssignTo:  &btn2,
						Text:      "取消",
						OnClicked: func() { d.Cancel() },
					},
				},
			},
		},
	}.Run(mw)
	if err != nil {
		log.Println(err)
	}
}
