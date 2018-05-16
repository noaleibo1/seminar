package main

import (
	"fmt"
	"github.com/noaleibo1/seminar/logic"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	logic.InitMatrix()

	Clusters := logic.GetClusterSizesSortedList()
	logic.FillMatrixWithLandUse(Clusters)
	fmt.Printf("\n\nLanduse Matrix\n")
	logic.PrintMatrix(logic.LandUseMatrix)
	logic.FillMatrixWithAltitude()
	fmt.Printf("\n\nAltitude Matrix\n")
	logic.PrintAltitudeMatrix(logic.LandUseMatrix)

	agents := logic.CreateHumanAgents(logic.NumberOfAgents)
	fmt.Printf("\n\nMatrix After Agents creation\n")
	logic.PrintMatrix(logic.LandUseMatrix)

	logic.ScatterBicycleStations()
	fmt.Printf("\n\nMatrix Bicycle Stations\n")
	logic.PrintBicycleMatrix(logic.LandUseMatrix)

	for i := 0; i < logic.NumberOfIterations; i++ {
		logic.MoveHumanAgentsOnce(agents, logic.WorkHours)
		fmt.Printf("\n\nLanduse Matrix after iteration %d\n", i)
		logic.PrintMatrix(logic.LandUseMatrix)
	}
}
