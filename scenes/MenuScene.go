package scenes

import (
	

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

)

type MenuScene struct {
	window fyne.Window
}

func NewMenuScene(fyneWindow fyne.Window) *MenuScene {
	scene := &MenuScene{window: fyneWindow}
	scene.RenderMenu()
	return scene
}

func (s*MenuScene) RenderMenu() {
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background2.png"))
	background.Resize(fyne.NewSize(1920,1080))
	background.Move(fyne.NewPos(-25,0))

	btnStartGame := widget.NewButton("Empezar", s.StartGame)
	btnStartGame.Resize(fyne.NewSize(130,30))
	btnStartGame.Move(fyne.NewPos(400,460))

	s.window.SetContent(container.NewWithoutLayout(background, btnStartGame))
}

func (s *MenuScene) StartGame() {
	NewGameScene(s.window)
	StartVehicleCreation()
	
}
