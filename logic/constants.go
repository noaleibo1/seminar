package logic

import "math"

const (
	MatrixSize                   = 100
	bicycleStationsGap           = 4

	NumberOfAgents               = 10000

	NumberOfIterations           = 90

	MaxResidentialClusterSize = 2000
	MinResidentialClusterSize = 400

	MaxEmploymentClusterSize = 2000
	MinEmploymentClusterSize = 500

	SlopeThreshold = 1.5

	citySteepness = 3
)

var (
	totalNumberOfDockingStations = totalNumberOfBicycles * 2
	totalNumberOfBicycles        = float64(NumberOfAgents / 5)
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

var NumberOfWalkDecisions = 0
var NumberOfBicycleDecisions = 0
var NumberOfPassesDueToTopography = 0
var NumberOfPassesDueToNoInfra = 0
var NumberOfPassesDueToRandom = 0
