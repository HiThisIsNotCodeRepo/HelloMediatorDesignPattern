# HelloMediatorDesignPattern

> [Source](https://golangbyexample.com/mediator-design-pattern-golang/)

## Procedures for train entering station

1. `train` call `requestArrival` to enter station.
2. `mediator` will check availability of the station ,when station is vacant the `train` will be allowed to enter
   otherwise has to queue.
3. When the occupying train leaves station the `mediator` will notify the first train in the `trainQueue` to enter the
   station by calling `permitArrival()`.
4. `mediator` contains lock means only one of `canLand` and `notifyFree` can be called at a time.
