package tool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"

	. "gitee.com/y2h/fynex/container"
)

const jsonStr = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
"key":true,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
	{"first": "James", "last": "Murphy"},
	{"first": "Roger", "last": "Craig"}
  ]
}`

func sjsonn() {
	input := widget.NewMultiLineEntry()
	input.SetText(string(pretty.Pretty([]byte(jsonStr))))
	key := widget.NewMultiLineEntry()
	key.PlaceHolder = "Input json key here"
	key.SetText("name.last")
	value := widget.NewMultiLineEntry()
	value.PlaceHolder = "Input json value here"
	value.SetText("Anderson")
	output := widget.NewMultiLineEntry()
	btnDo := widget.NewButton("Set or Update", func() {
		// sjson.Set(`{"key":true}`, "key", nil)
		// sjson.Set(`{"key":true}`, "key", false)
		// sjson.Set(`{"key":true}`, "key", 1)
		// sjson.Set(`{"key":true}`, "key", 10.5)
		// sjson.Set(`{"key":true}`, "key", "hello")
		// sjson.Set(`{"key":true}`, "key", []string{"hello", "world"})
		// sjson.Set(`{"key":true}`, "key", map[string]interface{}{"hello":"world"})
		// value, _ := sjson.Set("", "name", "Tom")
		// value, _ := sjson.Set("", "name.last", "Anderson")
		// value, _ := sjson.Set(`{"name":{"last":"Anderson"}}`, "name.first", "Sara")
		// value, _ := sjson.Set(`{"name":{"last":"Anderson"}}`, "name.last", "Smith")
		// value, _ := sjson.Set(`{"friends":["Andy","Carol"]}`, "friends.2", "Sara")
		value, _ := sjson.Set(input.Text, key.Text, value.Text)
		output.SetText(value)
	})

	btnAppend := widget.NewButton("Append", func() {
		// "children.-1"  >> appends a new value to the end of the children array
		// value, _ := sjson.Set(`{"friends":["Andy","Carol"]}`, "friends.4", "Sara")
		value, _ := sjson.Set(input.Text, key.Text+".-1", value.Text)
		output.SetText(value)
	})

	btnDelete := widget.NewButton("Delete", func() {
		// value, _ := sjson.Delete(`{"name":{"first":"Sara","last":"Anderson"}}`, "name.first")
		// value, _ := sjson.Delete(`{"friends":["Andy","Carol"]}`, "friends.1")
		// value, _ := sjson.Delete(`{"friends":["Andy","Carol"]}`, "friends.-1")
		value, _ := sjson.Delete(input.Text, key.Text)
		output.SetText(value)
	})

	cc := Border(HBox(btnDo, btnAppend, btnDelete),
		nil, nil, nil, GridWithCols(3,
			GridWithRows(2, key, value), input, output))
	w := fyne.CurrentApp().NewWindow("https://github.com/tidwall/sjson")

	w.SetContent(cc)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
