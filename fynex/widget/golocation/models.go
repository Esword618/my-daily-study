package golocation

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// "github.com/gogf/gf/v2/frame/g"
	// "github.com/gogf/gf/v2/util/gutil"
	xwdg "gitee.com/y2h/fynex/widget"
)

func UI() fyne.CanvasObject {
	goloc, err := New()
	if err != nil {
		log.Fatal(err)
	}
	lblMsg := widget.NewEntry()
	cs, err := goloc.AllCountries()
	states, err := goloc.AllStates()
	cities, err := goloc.AllCities()
	if err != nil {
		lblMsg.SetText(err.Error())
	}
	// goloc.database.Close()
	selCity := xwdg.NewSelectGeneric(nil).Col(2).Row(7).
		Do(func(sg *xwdg.SelectGeneric) {
		})

	selState := xwdg.NewSelectGeneric(nil).Col(2).Row(7).
		Do(func(sg *xwdg.SelectGeneric) {
			// for _, v := range states {
			// 	sg.Options = append(sg.Options, v)
			// }
		}).Change(func(sg *xwdg.SelectGeneric) {
		selCity.Options = nil
		for _, v := range cities {
			s := v
			if s.StateId == sg.Selected.(State).Id {
				selCity.Options = append(selCity.Options, s)
			}
		}
		if len(selCity.Options) == 0 {
			selCity.ClearSelected()
			selCity.PlaceHolder = "no data"
			return
		}
		selCity.SetSelectedIndex(0)
	})
	selCountry := xwdg.NewSelectGeneric(nil).Col(2).Row(7).
		Do(func(sg *xwdg.SelectGeneric) {
			for _, v := range cs {
				sg.Options = append(sg.Options, v)
			}
		}).Change(func(sg *xwdg.SelectGeneric) {
		selState.Options = nil
		for _, v := range states {
			s := v
			if s.CountryId == sg.Selected.(Country).Id {
				selState.Options = append(selState.Options, s)
			}
		}
		selState.SetSelectedIndex(0)
	})
	// goloc.GetCountryStates()
	// goloc.GetStateCites()
	return container.NewVBox(lblMsg,
		container.NewGridWithColumns(3,
			selCountry, selState, selCity))
}

//Country - struct that contains the country information
type Country struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Phonecode int    `json:"phonecode"`
}

func (c Country) String() string {
	return c.Name
}

//State - struct for housing the state information
type State struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CountryId int    `json:"country_id"`
}

func (c State) String() string {
	return c.Name
}

//City - Struct for housing the city information
type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StateId int    `json:"state_id"`
}

func (c City) String() string {
	return c.Name
}
