package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func NewLength(tometer float64, tocentimeter float64, tomillimeter float64, tokilometer float64, tonauticalmile float64,
	tomile float64, toyard float64, tofoot float64, toinch float64, tolightyear float64) (*Length, error) {
	length := &Length{
		toMeter:        tometer,
		toCentimeter:   tocentimeter,
		toMillimeter:   tomillimeter,
		toKilometer:    tokilometer,
		toNauticalMile: tonauticalmile,
		toMile:         tomile,
		toYard:         toyard,
		toFoot:         tofoot,
		toInch:         toinch,
		toLightYear:    tolightyear,
	}
	return length, nil
}

func (l *Length) convertLength(fromUnit string, toUnit string, num float64) (float64, error) {
	meterVal := 0.0
	if fromUnit != "meter\n" {
		switch fromUnit {
		case "centimeter\n":
			meterVal = num / l.toCentimeter
		case "millimeter\n":
			meterVal = num / l.toMillimeter
		case "kilometer\n":
			meterVal = num / l.toKilometer
		case "nautmile\n":
			meterVal = num / l.toNauticalMile
		case "mile\n":
			meterVal = num / l.toMile
		case "yard\n":
			meterVal = num / l.toYard
		case "foot\n":
			meterVal = num / l.toFoot
		case "inch\n":
			meterVal = num / l.toInch
		case "lightyear\n":
			meterVal = num / l.toInch
		default:
			meterVal = 1.0
		}
		switch toUnit {
		case "centimeter\n":
			return meterVal * l.toCentimeter, nil
		case "millimeter\n":
			return meterVal * l.toMillimeter, nil
		case "kilometer\n":
			return meterVal * l.toKilometer, nil
		case "nautmile\n":
			return meterVal * l.toNauticalMile, nil
		case "mile\n":
			return meterVal * l.toMile, nil
		case "yard\n":
			return meterVal * l.toYard, nil
		case "foot\n":
			return meterVal * l.toFoot, nil
		case "inch\n":
			return meterVal * l.toInch, nil
		case "lightyear\n":
			return meterVal * l.toLightYear, nil
		default:
			return meterVal, nil
		}

	}
	return meterVal, nil
}

//type Area struct {
//	toSquareMeter      float64
//	toSquareCentimeter float64
//	toSquareMillimeter float64
//	toHectare          float64
//	toAcre             float64
//	toSquareMile       float64
//	toSquareYard       float64
//	toSquareFoot       float64
//	toSquareInch       float64
//}
//
//type Volume struct {
//	toCubicMeter      float64
//	toCubicCentimeter float64
//	toCubicMillimeter float64
//	toLiter           float64
//	toMilliliter      float64
//	toCubicFoot       float64
//	toCubicInch       float64
//	toCubicYard       float64
//	toAcreFloat       float64
//}
//
//type Temperature struct {
//	toCelsius    float64
//	toFahrenheit float64
//	toKelvin     float64
//}
//
//type Speed struct {
//	toMeterPerSecond     float64
//	toKilometerPerHour   float64
//	toKilometerPerSecond float64
//	toKnot               float64
//	toMilePerHour        float64
//	toFoodPerSecond      float64
//	toInchPerSecond      float64
//	toMach               float64
//}
//
//type Time struct {
//	toSecond      float64
//	toMinute      float64
//	toHour        float64
//	toDay         float64
//	toWeek        float64
//	toYear        float64
//	toMillisecond float64
//}
//
//type Mass struct {
//	toGram     float64
//	toKilogram float64
//	toTonne    float64
//	toPound    float64
//	toOunce    float64
//	toCarat    float64
//	toStone    float64
//}
//
//type NumeralSystem struct {
//	toDecimal     int
//	toBinary      string
//	toHexadecimal string
//	toOctal       int
//}

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
	length, _ := NewLength(1.00, 100.0, 1000.0, 0.001,
		0.000539956803, 0.000621371192, 1.0936133, 3.2808399, 39.3700787,
		0.00000000000000010570008)
	unit, _ := reader.ReadString('\n')
	switch unit {
	case "1\n":
		fmt.Print("Choose the unit to convert from:")
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print("Choose the unit to convert to:")
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:")
		num, _ := reader.ReadString('\n')
		numStrip := num[:len(num)-1]
		convNum, _ := strconv.ParseFloat(numStrip, 64)
		fmt.Println(length.convertLength(fromUnit, toUnit, convNum))
	}
}
