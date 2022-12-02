package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) showMessage(msgs ...string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Debug messages!!!",
		Message:       strings.Join(msgs, "\n"),
		Buttons:       []string{"OK"},
		DefaultButton: "OK",
	})
}

func (a *App) showAbout(cd *menu.CallbackData) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Motes",
		Message:       "Motes is a markdown note book",
		Buttons:       []string{"OK"},
		DefaultButton: "OK",
		CancelButton:  "OK",
	})
}

func (a *App) quit(cd *menu.CallbackData) {
	runtime.Quit(a.ctx)
}

func (a *App) menu() *menu.Menu {
	root := menu.NewMenu()

	motes := root.AddSubmenu("Motes")
	motes.AddText("About Motes", nil, a.showAbout)
	motes.AddText("Open new window", nil, func(cd *menu.CallbackData) {
		path, err := os.Executable()
		if err != nil {
			a.showMessage(err.Error())
			return
		}
		err = exec.Command(path).Run()
		if err != nil {
			a.showMessage(err.Error())
		}
	})
	motes.AddSeparator()
	motes.AddText("Quit", keys.CmdOrCtrl("q"), a.quit)

	fileMenu := root.AddSubmenu("Messages")
	fileMenu.AddText("Show Message", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		var names []string
		var types []string
		for _, item := range root.Items {
			names = append(names, item.Label)
			types = append(types, string(item.Type))
		}
		a.showMessage(names...)
		a.showMessage(types...)
	})

	root.Append(menu.EditMenu())
	return root
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
