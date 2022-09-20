package tool

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	. "gitee.com/y2h/fynex/container"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

const jsonStr1 = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
"key":true,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
	 {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func gjsonn() {
	input := widget.NewMultiLineEntry()
	input.SetText(string(pretty.Pretty([]byte(jsonStr1))))
	key := widget.NewMultiLineEntry()
	key.PlaceHolder = "Input json key here"
	key.SetText("name.last")
	value := widget.NewMultiLineEntry()
	value.PlaceHolder = "Input json value here"
	value.SetText("Anderson")
	output := widget.NewMultiLineEntry()
	btnDo := widget.NewButton("JSON Lines", func() {
		input.SetText(`{"name": "Gilbert", "age": 61}
{"name": "Alexa", "age": 34}
{"name": "May", "age": 57}
{"name": "Deloise", "age": 44}`)
		// ..#                   >> 4
		// ..1                   >> {"name": "Alexa", "age": 34}
		// ..3                   >> {"name": "Deloise", "age": 44}
		// ..#.name              >> ["Gilbert","Alexa","May","Deloise"]
		// ..#(name="May").age   >> 57
		value := gjson.Get(input.Text, key.Text)
		output.SetText(value.String())
	})

	btnAppend := widget.NewButton("Get Value", func() {
		value.SetText(
			"name.last" + "\n" +
				"age" + "\n" +
				"children" + "\n" +
				"children.#" + "\n" +
				"children.1" + "\n" +
				"child*.2" + "\n" +
				"c?ildren.0" + "\n" +
				"fav\\.movie" + "\n" +
				"friends.#.first" + "\n" +
				"friends.1.last" + "\n" +
				`friends.#(last=="Murphy").first` + "\n" +
				`friends.#(last=="Murphy")#.first` + "\n" +
				"friends.#(age>45)#.last" + "\n" +
				`friends.#(first%"D*").last` + "\n" +
				`friends.#(first!%"D*").last` + "\n" +
				`friends.#(nets.#(=="fb"))#.first` + "\n",
		)
		// "name.last"          >> "Anderson"
		// "age"                >> 37
		// "children"           >> ["Sara","Alex","Jack"]
		// "children.#"         >> 3
		// "children.1"         >> "Alex"
		// "child*.2"           >> "Jack"
		// "c?ildren.0"         >> "Sara"
		// "fav\.movie"         >> "Deer Hunter"
		// "friends.#.first"    >> ["Dale","Roger","Jane"]
		// "friends.1.last"     >> "Craig"
		// friends.#(last=="Murphy").first    >> "Dale"
		// friends.#(last=="Murphy")#.first   >> ["Dale","Jane"]
		// friends.#(age>45)#.last            >> ["Craig","Murphy"]
		// friends.#(first%"D*").last         >> "Murphy"
		// friends.#(first!%"D*").last        >> "Craig"
		// friends.#(nets.#(=="fb"))#.first   >> ["Dale","Roger"]
		value := gjson.Get(input.Text, key.Text)
		output.SetText(value.String())
	})

	btnGetMany := widget.NewButton("GetMany", func() {
		output.Text=""
		// results := gjson.GetMany(json, "name.first", "name.last", "age")
		value := gjson.GetMany(input.Text, strings.Split(key.Text, ",")...)
		for _, v := range value {
			output.Text = output.Text + "\n" + v.String()
		}
		output.Refresh()
	})

	cc := Border(HBox(btnDo, btnAppend, btnGetMany),
		nil, nil, nil, GridWithCols(3,
			GridWithRows(2, key, value), input, output))
	w := fyne.CurrentApp().NewWindow("https://github.com/tidwall/gjson")

	w.SetContent(cc)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
