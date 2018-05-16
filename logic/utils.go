package logic

import (
	"math/rand"
)

func getRandomCellOfCertainType(landUseMatrix [MatrixSize][MatrixSize]*Cell, landUse LandUse) *Cell {
	i := rand.Intn(MatrixSize)
	j := rand.Intn(MatrixSize)
	for landUseMatrix[i][j].LandUse != landUse {
		i = rand.Intn(MatrixSize)
		j = rand.Intn(MatrixSize)
	}
	return landUseMatrix[i][j]
}
