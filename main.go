package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type forcedVariant struct {
	fyne.Theme

	variant fyne.ThemeVariant
}

func (f *forcedVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}

func main() {
	a := app.New()

	w := a.NewWindow("Fyne Demo")
	r, _ := fyne.LoadResourceFromPath("logo.png")
	w.SetIcon(r)

	ctrlQ := &desktop.CustomShortcut{
		KeyName:  fyne.KeyQ,
		Modifier: fyne.KeyModifierControl,
	}
	w.Canvas().AddShortcut(ctrlQ, func(shortcut fyne.Shortcut) {
		fmt.Println("'Ctrl+Q' has been typed")
		a.Quit()
	})

	text := canvas.NewText("(Press 'Ctrl+Q' to quit)", color.RGBA{A: 255, R: 255})
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 10
	spacer := canvas.NewText("", color.Transparent)
	spacer.TextSize = 3

	btns := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
		}),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
		}),
		layout.NewSpacer(),
	)

	content := container.NewVBox(
		widget.NewLabel("Hola, Enrique 😀!"),
		text,
		spacer,
		btns,
	)

	hCentered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), content, layout.NewSpacer())
	vCentered := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), hCentered, layout.NewSpacer())

	w.SetContent(vCentered)
	w.Resize(fyne.NewSize(400, 250))
	w.ShowAndRun()
}

/* REFERENCES

COMPILATION COMMANDS:
GOFLAGS="-ldflags=-w -ldflags=-s" go build -o demo .

GOFLAGS="-ldflags=-w -ldflags=-s" fyne package --exe demo

(for Windows)
CGO_ENABLED=1 \
GOOS=windows \
CC=x86_64-w64-mingw32-gcc \
  fyne package \
    --os windows \
    --release \
    --executable dist/demo.exe

PACKAGING FOR DESKTOP:
https://docs.fyne.io/started/
https://docs.fyne.io/started/packaging
https://docs.fyne.io/started/cross-compiling

BUNDLING RESOURCES:
https://docs.fyne.io/extend/bundle.html

GITHUB ACTIONS FOR SETUP GOLANG:
https://github.com/actions/setup-go

GITHUB ACTIONS TO UPLOAD BINARIES TO RELEASE:
https://github.com/svenstaro/upload-release-action

GITHUB ACTIONS: STORING INFORMATION IN VARIABLES:
https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/store-information-in-variables

RELEASE COMMAND:
git tag v1.0.1 && git push origin v1.0.1

FUNCIÓN "ELIMINAR" O "REMOVEITEM" EN LISTAS DE ENLACES #3100:
https://github.com/fyne-io/fyne/issues/3100
https://docs.fyne.io/api/v2.5/data/binding/untypedlist.html
https://pkg.go.dev/fyne.io/fyne/v2@v2.5.1/data/binding#UntypedList

RUNNING WINE IN A DOCKER CONTAINER:
https://hub.docker.com/r/scottyhardy/docker-wine/


./docker-wine --as-root --volume=$(pwd)/dist:/home/wineuser/dist
./docker-wine --volume=$(pwd)/dist:/home/wineuser/dist

ADDITIONAL INFORMATION:
https://leimao.github.io/blog/Docker-Wine/
https://alesnosek.com/blog/2015/07/04/running-wine-within-docker/

https://github.com/orgs/community/discussions/60820

https://github.com/actions/upload-release-asset

https://github.com/actions/checkout

https://stackoverflow.com/questions/27301806/using-after-a-heredoc-in-bash
*/
