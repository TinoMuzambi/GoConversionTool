package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	meterVal := num
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
			meterVal = num / l.toLightYear
		default:
			meterVal = 1.0
		}
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

func NewArea(tosquaremeter float64, tosquarecentimeter float64, tosquaremillimeter float64, tohectare float64,
	toacre float64, tosquaremile float64, tosquareyard float64, tosquarefoot float64, tosquareinch float64) (*Area, error) {
	area := &Area{
		toSquareMeter:      tosquaremeter,
		toSquareCentimeter: tosquarecentimeter,
		toSquareMillimeter: tosquaremillimeter,
		toHectare:          tohectare,
		toAcre:             toacre,
		toSquareMile:       tosquaremile,
		toSquareYard:       tosquareyard,
		toSquareFoot:       tosquarefoot,
		toSquareInch:       tosquareinch,
	}
	return area, nil
}

func (a *Area) convertArea(fromUnit string, toUnit string, num float64) (float64, error) {
	sqmeterval := num
	if fromUnit != "sqmeter\n" {
		switch fromUnit {
		case "sqcentimeter\n":
			sqmeterval = num / a.toSquareCentimeter
		case "sqmillimeter\n":
			sqmeterval = num / a.toSquareMillimeter
		case "hectare\n":
			sqmeterval = num / a.toHectare
		case "acre\n":
			sqmeterval = num / a.toAcre
		case "sqmile\n":
			sqmeterval = num / a.toSquareMile
		case "sqyard\n":
			sqmeterval = num / a.toSquareYard
		case "sqfoot\n":
			sqmeterval = num / a.toSquareFoot
		case "sqinch\n":
			sqmeterval = num / a.toSquareInch
		default:
			sqmeterval = 1.0
		}
	}
	switch toUnit {
	case "sqcentimeter\n":
		return sqmeterval * a.toSquareCentimeter, nil
	case "sqmillimeter\n":
		return sqmeterval * a.toSquareMillimeter, nil
	case "hectare\n":
		return sqmeterval * a.toHectare, nil
	case "acre\n":
		return sqmeterval * a.toAcre, nil
	case "sqmile\n":
		return sqmeterval * a.toSquareMile, nil
	case "sqyard\n":
		return sqmeterval * a.toSquareYard, nil
	case "sqfoot\n":
		return sqmeterval * a.toSquareFoot, nil
	case "sqinch\n":
		return sqmeterval * a.toSquareInch, nil
	default:
		return sqmeterval, nil
	}
}

type Volume struct {
	toLiter           float64
	toMilliliter      float64
	toCubicMeter      float64
	toCubicCentimeter float64
	toCubicMillimeter float64
	toCubicFoot       float64
	toCubicInch       float64
	toCubicYard       float64
	toAcreFoot        float64
}

func NewVolume(toliter float64, tomilliliter float64, tocubicmeter float64, tocubiccentimeter float64,
	tocubicmillimeter float64, tocubicfoot float64, tocubicinch float64, tocubicyard float64, toacrefoot float64) (*Volume, error) {
	volume := &Volume{
		toLiter:           toliter,
		toMilliliter:      tomilliliter,
		toCubicMeter:      tocubicmeter,
		toCubicCentimeter: tocubiccentimeter,
		toCubicMillimeter: tocubicmillimeter,
		toCubicFoot:       tocubicfoot,
		toCubicInch:       tocubicinch,
		toCubicYard:       tocubicyard,
		toAcreFoot:        toacrefoot,
	}
	return volume, nil
}

func (v *Volume) convertVolume(fromUnit string, toUnit string, num float64) (float64, error) {
	literVal := num
	if fromUnit != "liter\n" {
		switch fromUnit {
		case "cbmeter\n":
			literVal = num / v.toCubicMeter
		case "cbcentimeter\n":
			literVal = num / v.toCubicCentimeter
		case "cbmillimeter\n":
			literVal = num / v.toCubicMillimeter
		case "milliliter\n":
			literVal = num / v.toMilliliter
		case "cbfoot\n":
			literVal = num / v.toCubicFoot
		case "cbinch\n":
			literVal = num / v.toCubicInch
		case "cbyard\n":
			literVal = num / v.toCubicYard
		case "acrefoot\n":
			literVal = num / v.toAcreFoot
		default:
			literVal = 1.0
		}
	}
	switch toUnit {
	case "cbmeter\n":
		return literVal * v.toCubicMeter, nil
	case "cbcentimeter\n":
		return literVal * v.toCubicCentimeter, nil
	case "cbmillimeter\n":
		return literVal * v.toCubicMillimeter, nil
	case "milliliter\n":
		return literVal * v.toMilliliter, nil
	case "cbfoot\n":
		return literVal * v.toCubicFoot, nil
	case "cbinch\n":
		return literVal * v.toCubicInch, nil
	case "acrefoot\n":
		return literVal * v.toAcreFoot, nil
	default:
		return literVal, nil
	}
}

type Temperature struct {
}

func NewTemperature() (*Temperature, error) {
	temperature := &Temperature{}
	return temperature, nil
}

func (t *Temperature) convertTemperature(fromUnit string, toUnit string, num float64) (float64, error) {
	celsiusVal := num
	switch fromUnit {
	case "fahrenheit\n":
		celsiusVal = (num - 32) * 5 / 9
	case "kelvin\n":
		celsiusVal = num - 273.15
	default:
		celsiusVal = num
	}
	switch toUnit {
	case "fahrenheit\n":
		return (celsiusVal * 9 / 5) + 32, nil
	case "kelvin\n":
		return celsiusVal + 273.15, nil
	default:
		return celsiusVal, nil
	}
}

type Speed struct {
	toMeterPerSecond     float64
	toKilometerPerHour   float64
	toKilometerPerSecond float64
	toKnot               float64
	toMilePerHour        float64
	toFootPerSecond      float64
	toInchPerSecond      float64
	toMach               float64
}

func NewSpeed(tometerspersecond float64, tokilometersperhour float64, tokilometerspersecond float64, toknot float64,
	tomileperhour float64, tofootpersecond float64, toinchpersecond float64, tomach float64) (*Speed, error) {
	speed := &Speed{
		toMeterPerSecond:     tometerspersecond,
		toKilometerPerHour:   tokilometersperhour,
		toKilometerPerSecond: tokilometerspersecond,
		toKnot:               toknot,
		toMilePerHour:        tomileperhour,
		toFootPerSecond:      tofootpersecond,
		toInchPerSecond:      toinchpersecond,
		toMach:               tomach,
	}
	return speed, nil
}

func (s *Speed) convertSpeed(fromUnit string, toUnit string, num float64) (float64, error) {
	mpsVal := num
	if fromUnit != "meterspersecond\n" {
		switch fromUnit {
		case "kilometersperhour\n":
			mpsVal = num / s.toKilometerPerHour
		case "kilometerspersecond\n":
			mpsVal = num / s.toKilometerPerSecond
		case "knot\n":
			mpsVal = num / s.toKnot
		case "mileperhour\n":
			mpsVal = num / s.toMilePerHour
		case "footpersecond\n":
			mpsVal = num / s.toFootPerSecond
		case "inchpersecond\n":
			mpsVal = num / s.toInchPerSecond
		case "mach\n":
			mpsVal = num / s.toMach
		default:
			mpsVal = 1.0
		}
	}
	switch toUnit {
	case "kilometersperhour\n":
		return mpsVal * s.toKilometerPerHour, nil
	case "kilometerspersecond\n":
		return mpsVal * s.toKilometerPerSecond, nil
	case "knot\n":
		return mpsVal * s.toKnot, nil
	case "mileperhour\n":
		return mpsVal * s.toMilePerHour, nil
	case "footpersecond\n":
		return mpsVal * s.toFootPerSecond, nil
	case "inchpersecond\n":
		return mpsVal * s.toInchPerSecond, nil
	case "mach\n":
		return mpsVal * s.toMach, nil
	default:
		return mpsVal, nil
	}
}

type Time struct {
	toSecond      float64
	toMillisecond float64
	toMinute      float64
	toHour        float64
	toDay         float64
	toWeek        float64
	toYear        float64
}

func NewTime(tosecond float64, tomillisecond float64, tominute float64, tohour float64, today float64, toweek float64,
	toyear float64) (*Time, error) {
	time := &Time{
		toSecond:      tosecond,
		toMillisecond: tomillisecond,
		toMinute:      tominute,
		toHour:        tohour,
		toDay:         today,
		toWeek:        toweek,
		toYear:        toyear,
	}
	return time, nil
}

func (t *Time) convertTime(fromUnit string, toUnit string, num float64) (float64, error) {
	secondVal := num
	if fromUnit != "second\n" {
		switch fromUnit {
		case "millisecond\n":
			secondVal = num / t.toMillisecond
		case "minute\n":
			secondVal = num / t.toMinute
		case "hour\n":
			secondVal = num / t.toHour
		case "day\n":
			secondVal = num / t.toDay
		case "week\n":
			secondVal = num / t.toWeek
		case "year":
			secondVal = num / t.toYear
		default:
			secondVal = 1.0
		}
	}
	switch toUnit {
	case "millisecond\n":
		return secondVal * t.toMillisecond, nil
	case "minute\n":
		return secondVal * t.toMinute, nil
	case "hour\n":
		return secondVal * t.toHour, nil
	case "day\n":
		return secondVal * t.toDay, nil
	case "week\n":
		return secondVal * t.toWeek, nil
	case "year":
		return secondVal * t.toYear, nil
	default:
		return secondVal, nil
	}
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

func NewMass(togram float64, tokilogram float64, totonne float64, topound float64, toounce float64, tocarat float64,
	tostone float64) (*Mass, error) {
	mass := &Mass{
		toGram:     togram,
		toKilogram: tokilogram,
		toTonne:    totonne,
		toPound:    topound,
		toOunce:    toounce,
		toCarat:    tocarat,
		toStone:    tostone,
	}
	return mass, nil
}

func (m *Mass) convertMass(fromUnit string, toUnit string, num float64) (float64, error) {
	gramVal := num
	if fromUnit != "gram\n" {
		switch fromUnit {
		case "kilogram\n":
			gramVal = num / m.toKilogram
		case "tonne\n":
			gramVal = num / m.toTonne
		case "pound\n":
			gramVal = num / m.toPound
		case "ounce\n":
			gramVal = num / m.toOunce
		case "stone\n":
			gramVal = num / m.toStone
		default:
			gramVal = 1.0
		}
	}
	switch toUnit {
	case "kilogram\n":
		return gramVal * m.toKilogram, nil
	case "tonne\n":
		return gramVal * m.toTonne, nil
	case "pound\n":
		return gramVal * m.toPound, nil
	case "ounce\n":
		return gramVal * m.toOunce, nil
	case "stone\n":
		return gramVal * m.toStone, nil
	default:
		return gramVal, nil
	}
}

type NumeralSystem struct {
}

func NewNumeralSystem() (*NumeralSystem, error) {
	numeralSystem := &NumeralSystem{}
	return numeralSystem, nil
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func (n *NumeralSystem) convertNumeralSystems(fromUnit string, toUnit string, num string) (string, error) {
	decimalVal := num
	if fromUnit != "decimal\n" {
		switch fromUnit {
		case "binary\n":
			intString, _ := strconv.ParseInt(hexaNumberToInteger(num), 2, 64)
			decimalVal = strconv.FormatInt(intString, 10)
		case "hex\n":
			intString, _ := strconv.ParseInt(hexaNumberToInteger(num), 16, 64)
			decimalVal = strconv.FormatInt(intString, 10)
		case "octal\n":
			intString, _ := strconv.ParseInt(hexaNumberToInteger(num), 8, 64)
			decimalVal = strconv.FormatInt(intString, 10)
		default:
			decimalVal = num
		}
	}
	switch toUnit {
	case "binary\n":
		intString, _ := strconv.Atoi(decimalVal)
		return strconv.FormatInt(int64(intString), 2), nil
	case "hex\n":
		intString, _ := strconv.Atoi(decimalVal)
		return strconv.FormatInt(int64(intString), 16), nil
	case "octal\n":
		intString, _ := strconv.Atoi(decimalVal)
		return strconv.FormatInt(int64(intString), 8), nil
	default:
		return decimalVal, nil
	}
}

func main() {
	length, _ := NewLength(1.00, 100.0, 1000.0, 0.001,
		0.000539956803, 0.000621371192, 1.0936133, 3.2808399, 39.3700787,
		0.00000000000000010570008)
	area, _ := NewArea(1.0, 10000, 1000000, 0.0001,
		0.000247105407, 0.00000038610216, 1.19599005, 10.7639104, 1550.0031)
	volume, _ := NewVolume(1.0, 1000, 0.001, 1000, 1000000,
		0.0353146667, 61.0237441, 0.00130795062, 0.00000081037277)
	temperature, _ := NewTemperature()
	speed, _ := NewSpeed(1.0, 3.6, 0.001, 1.94384449,
		2.23693629, 3.2808399, 39.3700787, 0.0029385836)
	time, _ := NewTime(1.0, 1000, 0.0166666667, 0.000277777778, 0.000011574074,
		0.0000016534392, 0.000000031709792)
	mass, _ := NewMass(1.0, 0.001, 1/1000000, 0.00220462262, 0.0352739619,
		5, 0.000157473044)
	numeralSystem, _ := NewNumeralSystem()

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
	unit, _ := reader.ReadString('\n')
	switch unit {
	case "1\n":
		fmt.Print(`
Choose the unit to convert from:
meter
centimeter
millimeter
kilometer
nautmile
mile
yard
foot
inch
lightyear
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
meter
centimeter
millimeter
kilometer
nautmile
mile
yard
foot
inch
lightyear
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := length.convertLength(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "2\n":
		fmt.Print(`
Choose the unit to convert from:
sqmeter
sqcentimeter
sqmillimeter
hectare
acre
sqmile
sqyard
sqfoot
sqinch
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
sqmeter
sqcentimeter
sqmillimeter
hectare
acre
sqmile
sqyard
sqfoot
sqinch
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := area.convertArea(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "3\n":
		fmt.Print(`
Choose the unit to convert from:
cbmeter
cbcentimeter
cbmillimeter
liter
milliliter
cbyard
cbfoot
cbinch
acrefoot
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
cbmeter
cbcentimeter
cbmillimeter
liter
milliliter
cbyard
cbfoot
cbinch
acrefoot
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := volume.convertVolume(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "4\n":
		fmt.Print(`
Choose the unit to convert from:
celsius
fahrenheit
kelvin
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
celsius
fahrenheit
kelvin
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := temperature.convertTemperature(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "5\n":
		fmt.Print(`
Choose the unit to convert from:
meterpersecond
kilometerperhour
kilometerpersecond
knot
footpersecond
inchpersecond
mach
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
meterpersecond
kilometerperhour
kilometerpersecond
knot
footpersecond
inchpersecond
mach
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := speed.convertSpeed(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "6\n":
		fmt.Print(`
Choose the unit to convert from:
second
millisecond
minute
hour
day
week
year
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
second
millisecond
minute
hour
day
week
year
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := time.convertTime(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "7\n":
		fmt.Print(`
Choose the unit to convert from:
gram
kilogram
tonne
pound
ounce
carat
stone
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
gram
kilogram
tonne
pound
ounce
carat
stone
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum, _ := strconv.ParseFloat(num[:len(num)-1], 64)
		ans, _ := mass.convertMass(fromUnit, toUnit, convNum)
		fmt.Println(ans)
	case "8\n":
		fmt.Print(`
Choose the unit to convert from:
decimal
binary
hex
octal
`)
		fromUnit, _ := reader.ReadString('\n')
		fmt.Print(`
Choose the unit to convert to:
decimal
binary
hex
octal
`)
		toUnit, _ := reader.ReadString('\n')
		fmt.Print("Enter the number:\n")
		num, _ := reader.ReadString('\n')
		convNum := num[:len(num)-1]
		ans, _ := numeralSystem.convertNumeralSystems(fromUnit, toUnit, convNum)
		fmt.Println(strings.ToUpper(ans))
	}
}
