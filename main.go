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

	for i:=0;i<5;i++{
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

		writeToCSV(writer)
	}

}

func writeToCSV(writer *csv.Writer) {


	err := writer.Write([]string{"numberOfWalkDecisions", fmt.Sprintf("%d", logic.NumberOfWalkDecisions)})
	checkError("Cannot write to file", err)

	err = writer.Write([]string{"NumberOfBicycleDecisions", fmt.Sprintf("%d", logic.NumberOfBicycleDecisions)})
	checkError("Cannot write to file", err)

	err = writer.Write([]string{"NumberOfPassesDueToTopography", fmt.Sprintf("%d", logic.NumberOfPassesDueToTopography)})
	checkError("Cannot write to file", err)

	err = writer.Write([]string{"NumberOfPassesDueToRandom", fmt.Sprintf("%d", logic.NumberOfPassesDueToRandom)})
	checkError("Cannot write to file", err)

	err = writer.Write([]string{"NumberOfPassesDueToNoInfra", fmt.Sprintf("%d", logic.NumberOfPassesDueToNoInfra)})
	checkError("Cannot write to file", err)

	err = writer.Write([]string{"\n"})
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		panic(err)
	}
}
