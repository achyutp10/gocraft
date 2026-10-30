package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math error: square root of -ve no")
	}
	// compute
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}
	// process data
	return nil
}

func main() {

	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result1, err1 := sqrt(-16)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(result1)

	//--------------------

	data := []byte{}
	// if err := process(data); err != nil {
	errr := process(data)
	if errr != nil {
		fmt.Println("Error:", errr)
		return
	}
	fmt.Println("Data processed successfully")

	errr1 := eProcess()
	if errr != nil {
		fmt.Println(errr1)
		return
	}

	errrr := readData()
	if errrr != nil {
		fmt.Println(errrr)
	}
	fmt.Println("Data read successfully")

}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eProcess() error {
	return &myError{"Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}
