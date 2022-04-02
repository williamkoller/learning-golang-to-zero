package main

import "fmt"

func dayWeek(number int) string {
	switch number {
	case 1:
		return "sunday"
	case 2:
		return "monday"
	case 3:
		return "tuesday"
	case 4:
		return "wednesday"
	case 5:
		return "thursday"
	case 6:
		return "friday"
	case 7:
		return "saturday"
	default:
		return "number invalid"
	}
}

func dayWeekTwo(number int) string {
	var dayWeek string
	switch {
		case number == 1:
			dayWeek = "sunday"
			fallthrough
		case number == 2:
			dayWeek = "monday"
			fallthrough
		case number == 3:
			dayWeek = "tuesday"
			fallthrough
		case number == 4:
			dayWeek = "wednesday"
			fallthrough
		case number == 5:
			dayWeek = "thursday"
			fallthrough
		case number == 6:
			dayWeek = "friday"
			fallthrough
		case number == 7:
			dayWeek = "saturday"
		default:
			dayWeek = "number invalid"
	}

	return dayWeek
}

func dayWeekThree(number int) string {
	var dayWeek string
	switch {
		case number == 1:
			dayWeek = "sunday"
		case number == 2:
			dayWeek = "monday"
		case number == 3:
			dayWeek = "tuesday"
		case number == 4:
			dayWeek = "wednesday"
		case number == 5:
			dayWeek = "thursday"
		case number == 6:
			dayWeek = "friday"
		case number == 7:
			dayWeek = "saturday"
		default:
			dayWeek = "number invalid"
	}

	return dayWeek
}

func dayWeekFour(number int) string {
	switch {
		case number == 1:
			return "sunday"
		case number == 2:
			return "monday"
		case number == 3:
			return "tuesday"
		case number == 4:
			return "wednesday"
		case number == 5:
			return "thursday"
		case number == 6:
			return "friday"
		case number == 7:
			return "saturday"
		default:
			return "number invalid"
	}
}

func main() {
	day := dayWeek(6)
	dayTwo := dayWeekTwo(1)
	dayThree := dayWeekThree(4)
	dayFour := dayWeekFour(2)

	fmt.Println(day)
	
	fmt.Println("-------")

	fmt.Println(dayTwo)

	fmt.Println("-------")

	fmt.Println(dayThree)

	fmt.Println("-------")

	fmt.Println(dayFour)
}