package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

func ReadExpensesFromFile() ([]Expense, error) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return []Expense{}, err
	}

	filepath := path.Join(wd, "data.json")

	file, created, err := openFile(filepath)
	if created {
		fmt.Println("There is no history of past expenses")
		return []Expense{}, nil
	}

	defer file.Close()

	bytes, _ := io.ReadAll(file)
	if len(bytes) == 0 {
		fmt.Println("There is no history of past expenses")
		return []Expense{}, nil
	}

	var db Database

	err = json.Unmarshal(bytes, &db)
	if err != nil {
		return nil, err
	}

	return db.Expenses, nil
}

func openFile(filepath string) (f *os.File, created bool, er error) {

	//If file is Created, it fails (err == nil)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err == nil {
		return file, true, nil
	}

	//If file exists err == isExist
	if os.IsExist(err) {
		file, err = os.OpenFile(filepath, os.O_RDWR, 0644)
		return file, false, nil
	}

	return nil, false, err
}
