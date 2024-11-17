package main

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/sarchlab/akita/v3/monitoring"
	"github.com/sarchlab/akita/v3/sim"
	"github.com/sarchlab/zeonica/api"
	"github.com/sarchlab/zeonica/cgra"
	"github.com/sarchlab/zeonica/config"
	"github.com/tebeka/atexit"
)

var inputHeight = 2
var inputWidth = 2 

//go:embed input.cgraasm
var inputKernel string

func inputLayer(driver api.Driver) {
	// Preset input data for testing
	inputData := make([]uint32, inputHeight)
	for i := 0; i < inputHeight; i++ {
		inputData[i] = 3 // Example data, cycling through values 0-255
	}

	// Preset weight and bias data for testing
	weightData := make([]uint32, inputHeight)
	biasData := make([]uint32, inputWidth)
	for i := 0; i < inputHeight; i++ {
		weightData[i] = 2 //Example weight
	}
	for i := 0; i < inputWidth; i++ {
		biasData[i] = 1 // Example bias, set to 1 for simplicity
	}

	for x := 0; x < inputWidth; x++ {
		for y := 0; y < inputHeight; y++ {
			driver.MapProgram(inputKernel, [2]int{x, y})
		}
	}
	fmt.Println("Feeding in weight data...")
	driver.FeedIn(weightData, cgra.West, [2]int{0, inputWidth}, inputWidth, "R")
	driver.Run()

	fmt.Println("Feeding in input data...")
	driver.FeedIn(inputData, cgra.West, [2]int{0, inputHeight}, inputHeight, "B")
	driver.Run()
	
	// Feed in bias data
	fmt.Println("Feeding in bias data...")
	driver.FeedIn(biasData, cgra.North, [2]int{0, inputWidth}, inputWidth, "R")
	driver.Run()

	inputLayerOutput := make([]uint32, inputWidth) // Collect the results from the last row
	driver.FeedIn(biasData, cgra.North, [2]int{0, inputWidth}, inputWidth, "R")
	fmt.Println("Collecting input layer output...")
	driver.Collect(inputLayerOutput, cgra.South, [2]int{0, inputWidth}, inputWidth, "R")
	driver.Run()
	fmt.Println("Input Layer Output:", inputLayerOutput)
}

func main() {
	// Open the log file for writing
	logFile, err := os.OpenFile("cgra_simulation.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	// Redirect stdout and stderr to the log file
	os.Stdout = logFile
	os.Stderr = logFile
	monitor := monitoring.NewMonitor()

	engine := sim.NewSerialEngine()
	monitor.RegisterEngine(engine)

	driver := api.DriverBuilder{}.
		WithEngine(engine).
		WithFreq(1 * sim.GHz).
		Build("Driver")
	monitor.RegisterComponent(driver)

	device := config.DeviceBuilder{}.
		WithEngine(engine).
		WithFreq(1 * sim.GHz).
		WithWidth(inputWidth).
		WithHeight(inputHeight).
		WithMonitor(monitor).
		Build("Device")

	driver.RegisterDevice(device)

	monitor.StartServer()

	// Run the input layer of the MNIST MLP with the driver
	inputLayer(driver)

	// Keep the simulation alive for viewing results
	time.Sleep(10 * time.Hour)
	atexit.Exit(0)
}
