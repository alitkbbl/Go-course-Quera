package main

import (
	"fmt"
	"math"
)

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	return &Car{
		speed:   speed,
		battery: battery,
	}
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	car.battery += minutes / 2
	if car.battery > 100 {
		car.battery = 100
	}
}
func TryFinish(car *Car, distance int) string {
	if car.speed == 0 {
		return ""
	}
	batteryNeeded := float64(distance) / 2.0
	if batteryNeeded > float64(car.battery) {
		car.battery = 0
		return ""
	}
	time := float64(distance) / float64(car.speed)
	car.battery -= int(math.Floor(batteryNeeded))
	return fmt.Sprintf("%.2f hours", time)
}
