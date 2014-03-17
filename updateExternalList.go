package modules

import (
	"fmt"
	. "math"
	. "time"
)

func Update_external_list(newExternalList chan [][]int, externalList [][]int, currentFloor []int ,currentDir []int, newFloor int, newDir int) {
	load := make([]int, driver.ELEVATORS)

	for i := 0; i < ELEVATORS; i++ {
		if externalList[i][1] == 1 { //if floor not allready is ordered - do nothing?
			//Sleep??
			//Sleep(10 * Nanosecond)
			//check for the floor allready is ordered
		} else {
			//check if
			load[i] += Abs(currentFloor[i] - newFloor)

			if dir[i] != newDir {
				load[i] += 1
			}
		}
	}
	//Checks who has least load
	for i := 0; i < ELEVATORS; i++ {
		temp := 0
		if temp < load[i] { //selects the elevator with least load
			temp = i
		}
	}
	//Assign the order to a floor in the external list
	if i == 0 {
		externalList[newDir][newFloor] = 1
	} else {
		externalList[newDir*i][newFloor*i] = 1
	}
	newExternalList <- externalList

}