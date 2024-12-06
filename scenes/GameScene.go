
package scenes

import (
	"estacionamiento/models"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type GameScene struct {
	window fyne.Window
}

const (
	capacidad    = 20
	numVehiculos = 100
)

var (
	espaciosEstacionamiento = make(chan struct{}, capacidad)
	entrada                  = make(chan struct{}, 1)
	wg                       sync.WaitGroup
	vehicles                 = make([]*models.Vehicle, 0)
	startVehicleCreation = make(chan bool) 
    coordenadasEstacionamiento = []fyne.Position{
        {X: 490, Y: 350},
        {X: 585, Y: 350},
        {X: 686, Y: 350},
        {X: 785, Y: 350},
        {X: 878, Y: 350},
        {X: 970, Y: 350},
        {X: 1070, Y: 350},
        {X: 1165, Y: 350},
        {X: 1258, Y: 350},
        {X: 1355, Y: 350},

        {X: 490, Y: 650},
        {X: 585, Y: 650},
        {X: 686, Y: 650},
        {X: 785, Y: 650},
        {X: 878, Y: 650},
        {X: 970, Y: 650},
        {X: 1070, Y: 650},
        {X: 1165, Y: 650},
        {X: 1258, Y: 650},
        {X: 1355, Y: 650},


       
    }   
)

func StartVehicleCreation() {
	for _, vehicle := range vehicles {
        vehicle.Position = fyne.NewPos(300, 350)
    }
    startVehicleCreation <- true
}

func NewGameScene(fyneWindow fyne.Window) *GameScene {
	sceneGame := &GameScene{window: fyneWindow}
	sceneGame.RenderGame()
	return sceneGame
}


func (s *GameScene) RenderGame() {
    background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.png"))
    background.Resize(fyne.NewSize(1520,1080))
   
    
    vehicleContainer := container.NewWithoutLayout()

    s.window.SetContent(container.NewWithoutLayout(background, vehicleContainer))

    go func() {
        <-startVehicleCreation 
        for i := 0; i < numVehiculos; i++ {
            wg.Add(1)
            vehicle := models.NewVehicle(i)
            vehicles = append(vehicles, vehicle)
            go vehicleLlega(vehicle)
            vehicleContainer.Add(vehicle.Image)
            canvas.Refresh(vehicleContainer) 
             
        }
    }()
}

func (s *GameScene) BackMenu() {
	NewMenuScene(s.window)	
}

func vehicleLlega(vehicle *models.Vehicle) {
    fmt.Printf("El vehículo %d ha llegado.\n", vehicle.ID+1)
    entrada <- struct{}{}
    espaciosEstacionamiento <- struct{}{}
    time.Sleep(1 * time.Second)
    fmt.Printf("El vehículo %d está entrando al estacionamiento.\n", vehicle.ID+1)

    
    coordenadasDisponibles := []int{}
    for i, coordenada := range coordenadasEstacionamiento {
        ocupada := false
        for _, otroVehiculo := range vehicles {
            if otroVehiculo.Position == coordenada {
                ocupada = true
                break
            }
        }
        if !ocupada {
            coordenadasDisponibles = append(coordenadasDisponibles, i)
        }
    }
    if len(coordenadasDisponibles) > 0 {
        randomIndex := rand.Intn(len(coordenadasDisponibles))
        selectedCoordIndex := coordenadasDisponibles[randomIndex]
        vehicle.Position = coordenadasEstacionamiento[selectedCoordIndex]
        vehicle.Image.Move(vehicle.Position)
    } else {
        fmt.Printf("No hay coordenadas de estacionamiento disponibles para el vehículo %d. Espera en la posición inicial.\n", vehicle.ID)
    }
    <-entrada
    fmt.Printf("El vehículo %d está estacionado en la posición %v.\n", vehicle.ID+1, vehicle.Position)
    time.Sleep(time.Duration(1 + rand.Intn(30)) * time.Second)
    <-espaciosEstacionamiento
    fmt.Printf("El vehículo %d está saliendo del estacionamiento.\n", vehicle.ID+1)
    vehicle.Position = fyne.NewPos(200, 650) 
    vehicle.Image.Move(vehicle.Position)
    wg.Done()
}
