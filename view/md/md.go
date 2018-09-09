package md

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/shaowin16/tsdump/model"
	"github.com/shaowin16/tsdump/view"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "md"
)

// View Markdown视图
type View struct {
}

// NewView 返回Markdown视图实例
func NewView() view.Viewer {
	return new(View)
}

// Do 将数据库元数据以Markdown视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) error {
	for i := range items {
		v.renderDB(&items[i], out)
		fmt.Fprintln(out)
		for j := range items[i].Tables {
			fmt.Fprintf(out, "### `%s`\n%s\n\n",
				items[i].Tables[j].Name,
				items[i].Tables[j].Comment,
			)
			v.renderTable(&items[i].Tables[j], out)
			fmt.Fprintln(out)
		}
	}

	return nil
}

func (v *View) renderDB(db *model.DB, out io.Writer) {
	rows := [][]string{[]string{db.Name, db.CharSet, db.Collation}}

	t := tablewriter.NewWriter(out)
	t.SetHeader([]string{"Database", "Character Set", "Collation"})
	t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}

func (v *View) renderTable(table *model.Table, out io.Writer) {
	rows := make([][]string, 0, len(table.Columns))
	for i := range table.Columns {
		rows = append(rows, []string{
			table.Columns[i].Name,
			table.Columns[i].DataType,
			table.Columns[i].Nullable,
			table.Columns[i].Key,
			table.Columns[i].Default,
			// table.Columns[i].CharSet,
			// table.Columns[i].Collation,
			table.Columns[i].Comment,
		})
	}

	t := tablewriter.NewWriter(out)
	// t.SetHeader([]string{"Column", "Data Type", "Nullable", "Key", "Default", "Character Set", "Collation", "Comment"})
	t.SetHeader([]string{"字段名", "数据类型", "是否可为空", "是否为主键", "默认值","备注"})
	t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}
