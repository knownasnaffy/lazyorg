package views

import (
	"fmt"

	"github.com/HubertBel/lazyorg/internal/calendar"
	"github.com/HubertBel/lazyorg/internal/database"
	"github.com/jroimartin/gocui"
)

type NotepadView struct {
	*BaseView

	Database *database.Database
	IsActive bool
	content  string
}

func NewNotepadView(c *calendar.Calendar, db *database.Database) *NotepadView {
	tv := &NotepadView{
		BaseView: NewBaseView("notepad"),
		Database: db,
		IsActive: false,
	}

	content, err := db.GetLatestNote()
	if err == nil {
		tv.content = content
	}

	return tv
}

func (npv *NotepadView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		npv.Name,
		npv.X,
		npv.Y,
		npv.X+npv.W,
		npv.Y+npv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Notepad "
		v.Editable = true
		v.FgColor = gocui.AttrBold
		v.Wrap = true
		v.Clear()
		fmt.Fprint(v, npv.content)
	}

	npv.SaveContent(g)

	return nil
}

func (npv *NotepadView) SaveContent(g *gocui.Gui) error {
	v, err := g.View(npv.Name)
	if err != nil {
		return err
	}

	return npv.Database.SaveNote(v.Buffer())
}

func (npv *NotepadView) ClearContent(g *gocui.Gui) error {
	v, err := g.View(npv.Name)
	if err != nil {
		return err
	}

	v.Clear()
	g.CurrentView().SetCursor(0, 0)
	return npv.SaveContent(g)
}
