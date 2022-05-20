package cars

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Brand(brand string) (string, error) {

	if brand == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Car brand is %v", brand)
	//message := fmt.Sprintf("Car brand is") // uncomment to break the brand function to check a testing case
	return message, nil
}

func Model(model string) (string, error) {

	if model == "" {
		return "", errors.New("empty model")
	}
	message := fmt.Sprintf("Car model is %v", model)
	return message, nil
}

func Color() (string, error) {

	color := randomColor()

	if color == "" {
		return "", errors.New("empty color")
	}
	message := fmt.Sprintf("Car color is %v", color)
	return message, nil
}

func Milage() (string, error) {

	milage := randomMilage()

	if milage >= 15.0 {
		return "", errors.New("milage exceeds the minimum for a new car")
	}
	message := fmt.Sprintf("Car milage is %v", milage)
	return message, nil
}

func Registration(status bool) string {

	if status == true {
		message := fmt.Sprintf("Car is registered")
		return message
	}

	message := fmt.Sprintf("Car is not registered")
	return message
}

//init is executed automatically right after program startup - after global variables
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomColor() string {

	colors := []string{"blue", "silver", "red", "green", "yellow"}
	return colors[rand.Intn(len(colors))]
}

func randomMilage() float64 {

	milages := []float64{5.2, 8.9, 15.9, 21.4, 4.3, 12.4, 14.9, 7.8}
	return milages[rand.Intn(len(milages))]
}
