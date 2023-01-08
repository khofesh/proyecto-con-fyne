package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// get the current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)
	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolBar()
	app.Toolbar = toolbar

	priceTabContent := app.pricesTab()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holding", theme.InfoIcon(), canvas.NewText("Holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)
}
