package logic

import (
	"math/rand"
)

type Agent struct {
	AgentType         AgentType
	HomeAddress       *Cell
	EmploymentAddress *Cell
	CurrentPlace      *Cell
}

func CreateHumanAgents(n int) []*Agent {
	var agents []*Agent

	for i := 0; i < n; i++ {
		currentPlaceI := rand.Intn(MatrixSize)
		currentPlaceJ := rand.Intn(MatrixSize)
		agent := Agent{
			AgentType:         Human,
			HomeAddress:       getRandomCellOfCertainType(LandUseMatrix, Residential),
			EmploymentAddress: getRandomCellOfCertainType(LandUseMatrix, Employment),
			CurrentPlace:      LandUseMatrix[currentPlaceI][currentPlaceJ],
		}
		agents = append(agents, &agent)
		LandUseMatrix[currentPlaceI][currentPlaceJ].IsAgent = true
	}

	return agents
}

func MoveHumanAgentsOnce(agents []*Agent, currentTime TimeOfDay) {
	for _, agent := range agents {
		destinationCell := agent.GetNextCell(currentTime)
		agent.ChooseTransportationMethod(destinationCell)
		agent.MoveToCell(destinationCell)
	}
}

func (agent *Agent) MoveToCell(cell *Cell) {
	LandUseMatrix[agent.CurrentPlace.I][agent.CurrentPlace.J].IsAgent = false
	LandUseMatrix[cell.I][cell.J].IsAgent = true
	agent.CurrentPlace = cell
}

func (agent *Agent) GetNextCell(timeOfDay TimeOfDay) *Cell {
	probabilityParameter := rand.Float32()
	if agent.CurrentPlace.LandUse == Residential {
		return agent.getResidentialNextCell(probabilityParameter, timeOfDay)
	}

	if agent.CurrentPlace.LandUse == Employment {
		return agent.getEmploymentNextCell(probabilityParameter, timeOfDay)
	}

	//Agent is in land use Other
	return agent.getOtherNextCell(probabilityParameter, timeOfDay)
}

func (agent *Agent) getResidentialNextCell(probabilityParameter float32, timeOfDay TimeOfDay) *Cell {
	if timeOfDay == WorkHours {
		if probabilityParameter < 0.05 { // 5% of going to Other
			return getRandomCellOfCertainType(LandUseMatrix, Other)
		} else if probabilityParameter < 0.1 { // 10% of staying in the same place
			return agent.CurrentPlace
		} else { // 85% of going to work
			return agent.EmploymentAddress
		}
	}
	if timeOfDay == Afternoon {
		if probabilityParameter < 0.05 { // 5% of going to work
			return agent.EmploymentAddress
		} else if probabilityParameter < 0.15 { // 15% of going to Other
			return getRandomCellOfCertainType(LandUseMatrix, Other)
		} else { // 80% of staying in the same place
			return agent.CurrentPlace
		}
	}
	// Night
	// 0% of going to work
	if probabilityParameter < 0.1 { // 10% of going to Other
		return getRandomCellOfCertainType(LandUseMatrix, Other)
	} else { // 90% of staying in the same place
		return agent.CurrentPlace
	}
}

func (agent *Agent) getEmploymentNextCell(probabilityParameter float32, timeOfDay TimeOfDay) *Cell {
	if timeOfDay == WorkHours {
		if probabilityParameter < 0.05 { // 5% of going to Other
			return getRandomCellOfCertainType(LandUseMatrix, Other)
		} else if probabilityParameter < 0.15 { // 15% of going home
			return agent.HomeAddress
		} else { // 80% of staying in the same place
			return agent.CurrentPlace
		}
	}
	if timeOfDay == Afternoon {
		if probabilityParameter < 0.05 { // 5% of staying in the same place
			return agent.EmploymentAddress
		} else if probabilityParameter < 0.10 { // 10% of going to Other
			return getRandomCellOfCertainType(LandUseMatrix, Other)
		} else { // 85% of going home
			return agent.HomeAddress
		}
	}
	// Night
	// 0% of staying in the same place
	if probabilityParameter < 0.1 { // 10% of going to Other
		return getRandomCellOfCertainType(LandUseMatrix, Other)
	} else { // 90% of going home
		return agent.HomeAddress
	}
}

func (agent *Agent) getOtherNextCell(probabilityParameter float32, timeOfDay TimeOfDay) *Cell {
	if timeOfDay == WorkHours {
		if probabilityParameter < 0.05 { // 5% of staying in the same place
			return agent.CurrentPlace
		} else if probabilityParameter < 0.10 { // 10% of going to home
			return agent.HomeAddress
		} else { // 85% of going to work
			return agent.EmploymentAddress
		}
	}
	if timeOfDay == Afternoon {
		if probabilityParameter < 0.2 { // 20% of going to work
			return agent.EmploymentAddress
		} else if probabilityParameter < 0.30 { // 10% of going to home
			return agent.HomeAddress
		} else { // 50% of staying in the same
			return agent.CurrentPlace
		}
	}
	// Night
	// 0% of going to work
	if probabilityParameter < 0.1 { // 10% of staying in the same place
		return agent.CurrentPlace
	} else { // 90% of going home
		return agent.HomeAddress
	}
}

func (agent *Agent) ChooseTransportationMethod(destinationCell *Cell)  {
	
}
