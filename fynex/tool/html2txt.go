package tool

import (
	"io/ioutil"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"

	ctn "fyne.io/fyne/v2/container"
	wdg "fyne.io/fyne/v2/widget"
	"golang.org/x/net/html/charset"
	"jaytaylor.com/html2text"
)

func html2Txt() {
	w := fyne.CurrentApp().NewWindow("Html to txt")

	w.SetContent(delHtmlTagUI())
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}

func delHtmlTagUI() fyne.CanvasObject {
	txtUrl := wdg.NewEntry()
	txtUrl.SetPlaceHolder("input url here")
	txtUrl.SetText("https://www.diyifanwen.com/tool/sanzijing/")
	lead := wdg.NewMultiLineEntry()
	lead.SetPlaceHolder("input html source here")
	lead.Wrapping = fyne.TextWrapWord
	lead.SetText(inputHTML)
	right := wdg.NewMultiLineEntry()
	right.Wrapping = fyne.TextWrapWord
	lbl := wdg.NewLabel("")

	isUrl := wdg.NewCheck("Is Url", nil)
	prettyTables := wdg.NewCheck("PrettyTables", nil)
	btn := wdg.NewButton("Del Html Tag", func() {
		if isUrl.Checked {
			lead.SetText(getHtml(txtUrl.Text))
			lbl.SetText(delHtmlTag(lead.Text, right, prettyTables.Checked))
		} else {
			lbl.SetText(delHtmlTag(lead.Text, right, prettyTables.Checked))
		}

		if right.Text == "" {
			lbl.SetText("Nothing to do , give me work")
		} else {
			lbl.SetText("Success")
		}
	})
	split := ctn.NewHSplit(lead, right)
	split.Offset = 0.5
	return ctn.NewBorder(ctn.NewVBox(txtUrl, ctn.NewHBox(isUrl, prettyTables, btn, lbl)), nil, nil, nil, split)
}

func delHtmlTag(inputHTML string, txt *wdg.Entry, prettyTables bool) string {
	if inputHTML == "" {
		return "nothing to do"
	}

	text, err := html2text.FromString(inputHTML, html2text.Options{PrettyTables: prettyTables})
	errMsg(err)

	txt.SetText(text)
	txt.Refresh()
	return "Success"
}

func errMsg(err error) string {
	if err != nil {
		return "something wrong happen, try anothor url or html string"
	}
	return "Success"
}
func getHtml(urlStr string) string {
	req, err := http.NewRequest("GET", urlStr, nil)
	errMsg(err)

	req.Header = make(http.Header)

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	resp, err := http.DefaultClient.Do(req)
	errMsg(err)

	//@see : https://github.com/golang/net/blob/master/html/charset/charset.go
	reader, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(reader)
	errMsg(err)
	return string(bs)
}

var inputHTML = `
<html>
  <head>
    <title>My Mega Service</title>
    <link rel=\"stylesheet\" href=\"main.css\">
    <style type=\"text/css\">body { color: #fff; }</style>
  </head>

  <body>
    <div class="logo">
      <a href="http://jaytaylor.com/"><img src="/logo-image.jpg" alt="Mega Service"/></a>
    </div>

    <h1>Welcome to your new account on my service!</h1>

    <p>
      Here is some more information:

      <ul>
        <li>Link 1: <a href="https://example.com">Example.com</a></li>
        <li>Link 2: <a href="https://example2.com">Example2.com</a></li>
        <li>Something else</li>
      </ul>
    </p>

    <table>
      <thead>
        <tr><th>Header 1</th><th>Header 2</th></tr>
      </thead>
      <tfoot>
        <tr><td>Footer 1</td><td>Footer 2</td></tr>
      </tfoot>
      <tbody>
        <tr><td>Row 1 Col 1</td><td>Row 1 Col 2</td></tr>
        <tr><td>Row 2 Col 1</td><td>Row 2 Col 2</td></tr>
      </tbody>
    </table>
  </body>
</html>`
