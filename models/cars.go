package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)
const (
    autoSize = 300
    gameWidth    = 350
)

var imags = []string{
    "./assets/vwgreen.png",
    "./assets/vwyellow.png",
    "./assets/fiuum.png",
    "./assets/pinky.png",
}

type Vehicle struct {
    ID        int
    Image     *canvas.Image
    Position fyne.Position
}


func NewVehicle(id int ) *Vehicle {
    rand.Seed(time.Now().UnixNano())
    imagePath := imags[rand.Intn(len(imags))]

    Vehicle := &Vehicle{
        ID: id,
        Image: canvas.NewImageFromURI(storage.NewFileURI(imagePath)),
        Position: fyne.NewPos(300,350), 
       
    }
    Vehicle.Image.Resize(fyne.NewSize(80,80))
    return Vehicle
}
