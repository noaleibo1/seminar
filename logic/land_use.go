package logic

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Cell struct {
	LandUse          LandUse
	IsBicycleStation bool
	Bicycles         int
	DockingStations  int
	I                int
	J                int
	IsAgent          bool
	Altitude         float32
}

type clusters []*cluster

type cluster struct {
	landUse LandUse
	size    int
}

var (
	LandUseMatrix [MatrixSize][MatrixSize]*Cell

	residentialMinCluster = 4
	residentialMaxCluster = 1500

	employmentMinCluster = 1
	employmentMaxCluster = 1000

	counter = 1
)

func InitMatrix() {
	for a, _ := range LandUseMatrix {
		for b, _ := range LandUseMatrix[0] {
			LandUseMatrix[a][b] = &Cell{
				I: a,
				J: b,
			}
		}
	}
}

func GetClusterSizesSortedList() []*cluster {
	var clusters clusters
	currentClusterSize := MaxResidentialClusterSize
	for i := 1; currentClusterSize > MinResidentialClusterSize; i++ {
		size := int(1854.5 * math.Pow(float64(i), -1.0669))
		clusters = append(clusters, &cluster{
			landUse: Residential,
			size:    size,
		})
		currentClusterSize = size
	}

	currentClusterSize = MaxEmploymentClusterSize
	for i := 1; currentClusterSize > MinEmploymentClusterSize; i++ {
		size := int(1854.5 * math.Pow(float64(i), -1.0669))
		clusters = append(clusters, &cluster{
			landUse: Employment,
			size:    size,
		})
		currentClusterSize = size
	}

	sort.Sort(clusters)

	return clusters

	//return []*cluster{
	//	{
	//		landUse: 1,
	//		size:    5,
	//	},
	//	{
	//		landUse: 2,
	//		size:    6,
	//	},
	//	{
	//		landUse: 1,
	//		size:    7,
	//	},
	//	{
	//		landUse: 2,
	//		size:    10,
	//	},
	//	{
	//		landUse: 1,
	//		size:    4,
	//	},
	//}
}

func FillMatrixWithLandUse(clusters []*cluster) {
	for _, cluster := range clusters {
		if cluster.size > (MatrixSize * MatrixSize) {
			fmt.Printf("\n\n!!! cluster.size > (MatrixSize*MatrixSize) !!!\n\n")
			panic(1)
		}
		counter = 0
		freeCell := getRandomCellOfCertainType(LandUseMatrix, Other)
		expand(freeCell.I, freeCell.J, cluster)
	}
}

func expand(i int, j int, c *cluster) {
	if counter > c.size-1 {
		return
	}
	LandUseMatrix[i][j].LandUse = c.landUse
	counter++
	for _, freeNeighbourCell := range getFreeNeighbourCells(i, j) {
		expand(freeNeighbourCell.I, freeNeighbourCell.J, c)
	}
}

func FillMatrixWithAltitude() {
	for _, row := range LandUseMatrix {
		for _, cell := range row {
			fraction := rand.Float32()
			if r := rand.Float32(); r > 0.5 {
				cell.Altitude = float32(rand.Intn(citySteepness)) - fraction
			} else {
				cell.Altitude = float32(rand.Intn(citySteepness)) - fraction + 0
			}
		}
	}
}

func ScatterBicycleStations() {
	for i := 0; i < MatrixSize; i++ {
		for j := 0; j < MatrixSize; j += bicycleStationsGap {
			if (i+j+MatrixSize%bicycleStationsGap+1)%bicycleStationsGap == 1 {
				LandUseMatrix[i][j].IsBicycleStation = true
				LandUseMatrix[i][j].Bicycles = bicyclesPerStation
				LandUseMatrix[i][j].DockingStations = dockingStationsPerStation
			}

		}
	}
}

func getFreeNeighbourCells(i int, j int) []*Cell {
	var cells []*Cell

	//Matrix corners
	if i == MatrixSize-1 && j == MatrixSize-1 {
		if LandUseMatrix[i-1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j])
		}
		if LandUseMatrix[i-1][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j-1])
		}
		if LandUseMatrix[i][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j-1])
		}
		return cells
	}
	if i == MatrixSize-1 && j == 0 {
		if LandUseMatrix[i-1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j])
		}
		if LandUseMatrix[i-1][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j+1])
		}
		if LandUseMatrix[i][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j+1])
		}
		return cells
	}
	if i == 0 && j == MatrixSize-1 {
		if LandUseMatrix[i][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j-1])
		}
		if LandUseMatrix[i+1][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j-1])
		}
		if LandUseMatrix[i+1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j])
		}
		return cells
	}
	if i == 0 && j == 0 {
		if LandUseMatrix[i+1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j])
		}
		if LandUseMatrix[i+1][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j+1])
		}
		if LandUseMatrix[i][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j+1])
		}
		return cells
	}

	if i == MatrixSize-1 {
		if LandUseMatrix[i][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j-1])
		}
		if LandUseMatrix[i-1][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j-1])
		}
		if LandUseMatrix[i-1][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j+1])
		}
		if LandUseMatrix[i][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j+1])
		}
		if LandUseMatrix[i-1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j])
		}
		return cells
	}

	if i == 0 {
		if LandUseMatrix[i][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j-1])
		}
		if LandUseMatrix[i+1][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j-1])
		}
		if LandUseMatrix[i+1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j])
		}
		if LandUseMatrix[i+1][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j+1])
		}
		if LandUseMatrix[i][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j+1])
		}
		return cells
	}

	if j == 0 {
		if LandUseMatrix[i-1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j])
		}
		if LandUseMatrix[i-1][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j+1])
		}
		if LandUseMatrix[i][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j+1])
		}
		if LandUseMatrix[i+1][j+1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j+1])
		}
		if LandUseMatrix[i+1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j])
		}
		return cells
	}

	if j == MatrixSize-1 {
		if LandUseMatrix[i-1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j])
		}
		if LandUseMatrix[i-1][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i-1][j-1])
		}
		if LandUseMatrix[i][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i][j-1])
		}
		if LandUseMatrix[i+1][j-1].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j-1])
		}
		if LandUseMatrix[i+1][j].LandUse == 0 {
			cells = append(cells, LandUseMatrix[i+1][j])
		}
		return cells
	}

	if LandUseMatrix[i-1][j].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i-1][j])
	}
	if LandUseMatrix[i-1][j-1].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i-1][j-1])
	}
	if LandUseMatrix[i][j-1].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i][j-1])
	}
	if LandUseMatrix[i+1][j-1].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i+1][j-1])
	}
	if LandUseMatrix[i+1][j].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i+1][j])
	}
	if LandUseMatrix[i-1][j+1].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i-1][j+1])
	}
	if LandUseMatrix[i][j+1].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i][j+1])
	}
	if LandUseMatrix[i+1][j+1].LandUse == 0 {
		cells = append(cells, LandUseMatrix[i+1][j+1])
	}
	return cells
}

func PrintMatrix(LandUseMatrix [MatrixSize][MatrixSize]*Cell) {
	for _, cells := range LandUseMatrix {
		for _, cell := range cells {
			if cell.IsAgent {
				fmt.Printf("|A|")
			} else {
				fmt.Printf("|%v|", cell.LandUse)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func PrintAltitudeMatrix(LandUseMatrix [MatrixSize][MatrixSize]*Cell) {
	for _, cells := range LandUseMatrix {
		for _, cell := range cells {
			fmt.Printf("|%d|", cell.Altitude)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func PrintBicycleMatrix(LandUseMatrix [MatrixSize][MatrixSize]*Cell) {
	for _, cells := range LandUseMatrix {
		for _, cell := range cells {
			fmt.Printf("|%d|", cell.Bicycles)
		}
		fmt.Printf("\n")
		for i := 0; i < MatrixSize; i++ {
			fmt.Printf("---")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func (c clusters) Less(i, j int) bool {
	return c[i].size < c[j].size
}

func (c clusters) Len() int {
	return len(c)
}

func (c clusters) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
