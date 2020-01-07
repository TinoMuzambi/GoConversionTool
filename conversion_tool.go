package main

import (
	"bufio"
	"fmt"
	"os"
)

type Length struct {
	toMeter        float64
	toCentimeter   float64
	toMillimeter   float64
	toKilometer    float64
	toNauticalMile float64
	toMile         float64
	toYard         float64
	toFoot         float64
	toInch         float64
	toLightYear    float64
}

type Area struct {
	toSquareMeter      float64
	toSquareCentimeter float64
	toSquareMillimeter float64
	toHectare          float64
	toAcre             float64
	toSquareMile       float64
	toSquareYard       float64
	toSquareFoot       float64
	toSquareInch       float64
}

type Volume struct {
	toCubicMeter      float64
	toCubicCentimeter float64
	toCubicMillimeter float64
	toLiter           float64
	toMilliliter      float64
	toCubicFoot       float64
	toCubicInch       float64
	toCubicYard       float64
	toAcreFloat       float64
}

type Temperature struct {
	toCelsius    float64
	toFahrenheit float64
	toKelvin     float64
}

type Speed struct {
	toMeterPerSecond     float64
	toKilometerPerHour   float64
	toKilometerPerSecond float64
	toKnot               float64
	toMilePerHour        float64
	toFoodPerSecond      float64
	toInchPerSecond      float64
	toMach               float64
}

type Time struct {
	toSecond      float64
	toMinute      float64
	toHour        float64
	toDay         float64
	toWeek        float64
	toYear        float64
	toMillisecond float64
}

type Mass struct {
	toGram     float64
	toKilogram float64
	toTonne    float64
	toPound    float64
	toOunce    float64
	toCarat    float64
	toStone    float64
}

type NumeralSystem struct {
	toDecimal     int
	toBinary      string
	toHexadecimal string
	toOctal       int
}

func main() {
	fmt.Println("====================================")
	fmt.Println("Welcome to my Go Conversion Tool")
	fmt.Println("====================================")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(`
Make a selection of the type of conversion you want by entering the number.
1. Length
2. Area
3. Volume
4. Temperature
5. Speed
6. Time
7. Mass
8. Numeral System
`)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
