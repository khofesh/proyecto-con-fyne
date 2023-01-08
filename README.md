# Proyecto con fyne

project with fyne (golang)

## getting started

install several applications

```shell
sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
```

## tidy

add a simple code in main.go

```golang
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hola Mundo")

	w.SetContent(widget.NewLabel("Hola, mundo!"))

	w.ShowAndRun()
}
```

then, run `go mod tidy`

## packaging

```shell
go install fyne.io/fyne/v2/cmd/fyne@latest

fyne package -appVersion 1.0.0 -name MarkDown -release
```

## sqlite

```shell
sudo dnf install sqlite
sudo dnf install sqlitebrowser
go get github.com/glebarez/go-sqlite
```

clean cache

```shell
go clean -modcache
```

# references

- https://fyne.io/
- https://www.udemy.com/course/building-gui-applications-with-fyne-and-go-golang/
- https://developer.fyne.io/started/
- https://developer.fedoraproject.org/tech/database/sqlite/about.html
