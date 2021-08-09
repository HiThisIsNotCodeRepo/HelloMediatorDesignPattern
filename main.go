package main

func main() {
	aStationManager := newStationManager()
	aPassengerTrain := &passengerTrain{
		mediator: aStationManager,
	}
	aGoodsTrain := &goodsTrain{
		mediator: aStationManager,
	}
	aPassengerTrain.requestArrival()
	aGoodsTrain.requestArrival()
	aPassengerTrain.departure()
}
