package modules

const (
	FLOORS = 4
)

type Slave struct {
	nr int
	internalList []bool
	externalList [][]int
	currentFloor int //get from driver/IO
	direction    int // get from driver/IO

}

type Master struct {
	nr := 1 
	externalList  [][]int
	currentFloors []int
	directions    []int
	internalList  []int
}

/*
Get orders from the optimalizaton algorithm
*/
func Slave_init() {

	//communicate with driver

	//run externalList

	//internal list

}

func (s Slave) Recive_externalList_from(externalListChan chan [][]int) {
	s.externalList = <-externalListChan
	}
}

func Send_confirmation_to_master_or_comm(message string, outgoingMessageChan chan string) {
	//transfer to the communication module(put on correct tag)
	outgoingOrderChan <- message
}

func (s Slave) Update_current_floor_and_direction(currentFloorChan chan int, directionChan chan int) {
	select {
	case o := <-currenFloorChan:
		s.currentFloor = o
	case o := <-directionChan:
		s.driection = o
	}

}

func (s Slave) Send_slave_to_state(slaveStateChan chan int) { //send next floor to statemachine
	var nextFloor int

	//if heading up, pick up others who is going up, both external and internal

	s.externalList[][]
	//if heading down pick up othes who iss going down, both external and internal

	 slaveStateChan <- nextFloor
}


func (m Master) getOptimalExternalList(optimalityChan chan [][]int) {
	//gets new OptimnalExternalList from module
}

func (m Master) sendExternalListToSlaves(masterToCommunicationChan chan [][]int) {
	//sends new external list to communication module 
}

func (m Master) 