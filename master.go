package modules

import (
	"fmt"
)

const ()

type Master struct {
	nr := 1 
	externalList  [][]int
	currentFloors []int
	directions    []int
	internalList  []int
}

func (m Master) getOptimalExternalList(optimalityChan chan [][]int) {
	//gets new OptimnalExternalList from module
}

func (m Master) sendExternalListToSlaves(masterToCommunicationChan chan [][]int) {
	//sends new external list to communication module 
}

func (m Master) 