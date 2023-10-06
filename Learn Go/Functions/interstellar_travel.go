package main

import (
    f "fmt"
)

func fuelGauge(fuel int32) {
    f.Printf("uu have %d.. umm... liters? units?? of fuel left c:" + "\n", fuel)
}

func calculateFuel(planet string) int32 {
    var fuel int32
    switch planet{
    case "Venus":
        fuel = 300000
    case "Mercury":
        fuel = 500000
    case "Mars":
        fuel = 700000
    default:
        fuel = 0
    }

    return fuel
}

func greetPlanet(planet string) {
    switch planet{
    case "Venus":
        f.Println("welcome toooo VENUSSSS!!")
    case "Mercury":
        f.Println("umm it's uhhh mercury, we have ummm creaters and umm.. an exosphere.. ig? why did u come here?")
    case "Mars":
        f.Println("it's marsss!! :D")
    default:
        f.Println("i got lost :c")
    }
}

func cantFly() {
    f.Println("we can't fly here cuz we don't have enough monies :c")
}

func flyToPlanet(planet string, fuel int32) int32 {
    var fuelRemaining = fuel
    var fuelCost = calculateFuel(planet)
    if fuelRemaining >= fuelCost {
        greetPlanet(planet)
        fuelRemaining -= fuelCost
    } else {
        cantFly()
    }

    return fuelRemaining
}

func main() {
    var fuel int32 = 1000000
    var planetChoice = "Mercury"

    fuel = flyToPlanet(planetChoice, fuel)
    fuelGauge(fuel)
}
