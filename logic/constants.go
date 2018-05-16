package logic

import "math"

const (
	MatrixSize                   = 10
	bicycleStationsGap           = 5
	NumberOfAgents               = 50
	NumberOfIterations           = 1
)

var (
	totalNumberOfDockingStations = totalNumberOfBicycles * 2
	totalNumberOfBicycles        = float64(NumberOfAgents / 10)
	numberOfBicycleStations      = float64((MatrixSize * MatrixSize) / bicycleStationsGap)
	bicyclesPerStation           = int(math.Ceil(totalNumberOfBicycles / numberOfBicycleStations))
	dockingStationsPerStation    = int(math.Ceil(totalNumberOfDockingStations / numberOfBicycleStations))
)

type AgentType int

const (
	Human AgentType = iota
	BicycleAgent
	DockingStation
)

type TimeOfDay int

const (
	WorkHours TimeOfDay = iota // 09:00-17:00
	Afternoon                  // 17:00-21:00
	Night                      // 21:00-09:00
)

type LandUse int

const (
	Other LandUse = iota
	Residential
	Employment
)

type TransportMethod int

const (
	Bicycle TransportMethod = iota
	Walk
)
