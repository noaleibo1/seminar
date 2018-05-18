package main

import (
	"encoding/csv"
	"fmt"
	"github.com/noaleibo1/seminar/logic"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		"iteration",
		"numberOfWalkDecisions",
		"NumberOfBicycleDecisions",
		"NumberOfPassesDueToTopography",
		"NumberOfPassesDueToRandom",
		"NumberOfPassesDueToNoInfra",
	})
	checkError("Cannot write to file", err)

	for i := 0; i < 10; i++ {
		logic.NumberOfWalkDecisions = 0
		logic.NumberOfBicycleDecisions = 0
		logic.NumberOfPassesDueToTopography = 0
		logic.NumberOfPassesDueToNoInfra = 0
		logic.NumberOfPassesDueToRandom = 0
		logic.InitMatrix()

		clusters := logic.GetClusterSizesSortedList()
		logic.FillMatrixWithLandUse(clusters)
		logic.FillMatrixWithAltitude()

		agents := logic.CreateHumanAgents(logic.NumberOfAgents)

		logic.ScatterBicycleStations()

		for i := 0; i < logic.NumberOfIterations; i++ {
			logic.MoveHumanAgentsOnce(agents, logic.TimeOfDay(i))
		}

		writeToCSV(writer, i+1)
	}

}

func writeToCSV(writer *csv.Writer, i int) {
	is := fmt.Sprintf("%d", i)
	a := fmt.Sprintf("%d", logic.NumberOfWalkDecisions)
	b := fmt.Sprintf("%d", logic.NumberOfBicycleDecisions)
	c := fmt.Sprintf("%d", logic.NumberOfPassesDueToTopography)
	d := fmt.Sprintf("%d", logic.NumberOfPassesDueToRandom)
	e := fmt.Sprintf("%d", logic.NumberOfPassesDueToNoInfra)

	err := writer.Write([]string{is, a, b, c, d, e})
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		panic(err)
	}
}
