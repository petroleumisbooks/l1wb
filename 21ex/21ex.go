package main

import "fmt"

/*
Перевод из фаренгейтов в цельсий
*/

// Интерфейс, к которому необходимо произвести адаптацию
type TempSensor interface {
	getCurrentTemp() float64
}

type ClosedTempSensor struct {
	tempInFahrenheit float64
}

func (cts *ClosedTempSensor) getTemperature() float64 {
	return cts.tempInFahrenheit
}

type AdapterSensor struct {
	cts *ClosedTempSensor
}

func (adapterSensor AdapterSensor) getCurrentTemp() float64 {
	return (adapterSensor.cts.tempInFahrenheit - 32.0) * 5.0 / 9.0
}

var cts = ClosedTempSensor{tempInFahrenheit: 132}

func main() {
	var adapterSensor TempSensor = AdapterSensor{cts: &cts}

	fmt.Printf("Current temperature in celsius: %f", adapterSensor.getCurrentTemp())

}
