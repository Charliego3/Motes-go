package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var exts = []string{
	"*.md",
	"*.markdown",
	"*.mdown",
	"*.mkdn",
	"*.mkd",
	"*.mdwn",
	"*.mdtxt",
	"*.mdtext",
	"*.Rmd",
}

func isMd(ext string) bool {
	if len(ext) > 0 && ext[0] != '*' {
		ext = "*" + ext
	}
	for _, e := range exts {
		if strings.EqualFold(ext, e) {
			return true
		}
	}
	return false
}

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

func (a *App) showMessage(typ runtime.DialogType, msgs ...string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          typ,
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
			a.showMessage(runtime.InfoDialog, err.Error())
			return
		}
		err = exec.Command(path).Run()
		if err != nil {
			a.showMessage(runtime.InfoDialog, err.Error())
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
		a.showMessage(runtime.InfoDialog, names...)
		a.showMessage(runtime.InfoDialog, types...)
	})

	root.Append(menu.EditMenu())
	return root
}

func (a *App) Open(path string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Can't open file!",
			Message:       err.Error(),
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		return "", err
	}

	return string(buf), nil
}

// FetchFileTree returns a greeting for the given name
func (a *App) FetchFileTree() ([]FileNode, error) {
	if len(config.Root) == 0 {
		root, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
			DefaultDirectory:     homeDir,
			Title:                "Please choose the note root directory",
			CanCreateDirectories: true,
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Markdown Files (*.md, *.markdown...)",
					Pattern:     "",
				},
			},
		})

		if err != nil {
			_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
				Type:          runtime.ErrorDialog,
				Title:         "Choose root directory has error!!!",
				Message:       err.Error(),
				Buttons:       []string{"OK"},
				DefaultButton: "OK",
			})
			os.Exit(2)
		}

		if len(root) == 0 {
			_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
				Type:          runtime.WarningDialog,
				Title:         "Doesn't choose a root directory!!!",
				Message:       "Then will be exit.",
				Buttons:       []string{"OK"},
				DefaultButton: "OK",
			})
			os.Exit(0)
		}

		a.saveConfigPath(root)
	}

	tree, err := a.buildTree(config.Root)
	if err != nil {
		_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Open directory has error!",
			Message:       err.Error(),
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
	}
	return tree, nil
}

func (a *App) buildTree(path string) ([]FileNode, error) {
	directories, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	tree := make([]FileNode, 0, len(directories))
	for _, entry := range directories {
		name := entry.Name()
		ext := filepath.Ext(name)
		name = strings.TrimSuffix(name, ext)
		if !entry.IsDir() && !isMd(ext) {
			continue
		}

		fpath := filepath.Join(path, entry.Name())
		var children []FileNode
		if entry.IsDir() {
			var err error
			children, err = a.buildTree(fpath)
			if err != nil {
				return nil, err
			}
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		const formatter = "2006-01-02 15:04:05"
		stat := info.Sys().(*syscall.Stat_t)
		createAt := time.Unix(stat.Ctimespec.Sec, 0).Format(formatter)
		updateAt := time.Unix(stat.Mtimespec.Sec, 0).Format(formatter)

		tree = append(tree, FileNode{
			Name:     name,
			Path:     fpath,
			IsDir:    entry.IsDir(),
			CreateAt: createAt,
			UpdateAt: updateAt,
			Children: children,
		})
	}

	if len(tree) > 0 {
		sort.Slice(tree, func(i, j int) bool {
			if tree[i].IsDir && !tree[j].IsDir {
				return true
			} else if !tree[i].IsDir && tree[j].IsDir {
				return false
			}

			return strings.ToLower(tree[i].Name) < strings.ToLower(tree[j].Name)
		})
	}
	return tree, nil
}

type FileNode struct {
	Name     string     `json:"name,omitempty"`
	Path     string     `json:"path,omitempty"`
	IsDir    bool       `json:"isDir,omitempty"`
	CreateAt string     `json:"createAt,omitempty"`
	UpdateAt string     `json:"updateAt,omitempty"`
	Children []FileNode `json:"children,omitempty"`
}

func (a *App) saveConfigPath(root string) {
	if err := saveRootPath(root); err != nil {
		choose, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Save the root directory configuration error!",
			Message:       err.Error() + "\nMaybe will be retry.",
			Buttons:       []string{"Retry", "Cancel"},
			DefaultButton: "Retry",
			CancelButton:  "Cancel",
		})
		if choose == "Retry" {
			a.saveConfigPath(root)
		}
	}
}
