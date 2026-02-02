package internal

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

type ReadResult struct {
	Expenses    []Expense
	FileCreated bool
	FileEmpty   bool
}

func ReadExpensesFromFile() (ReadResult, error) {
	wd, err := os.Getwd()
	if err != nil {
		return ReadResult{}, err
	}

	filepath := path.Join(wd, "data.json")

	file, created, err := openFile(filepath)
	if err != nil {
		return ReadResult{}, err
	}
	if created {
		return ReadResult{
			Expenses:    []Expense{},
			FileCreated: true,
		}, nil
	}

	defer file.Close()

	bytes, _ := io.ReadAll(file)
	if len(bytes) == 0 {
		return ReadResult{
			Expenses:  []Expense{},
			FileEmpty: true,
		}, nil
	}

	var expenses []Expense

	err = json.Unmarshal(bytes, &expenses)
	if err != nil {
		return ReadResult{}, err
	}

	return ReadResult{Expenses: expenses}, nil
}

func WriteExpensesToFile(expenses []Expense) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	filepath := path.Join(wd, "data.json")

	b, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}

	tmp := filepath + ".tmp"

	err = os.WriteFile(tmp, b, 0644)
	if err != nil {
		return err
	}

	return os.Rename(tmp, filepath)

}

func openFile(filepath string) (f *os.File, created bool, er error) {

	//If file is Created, it fails (err == nil)
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE|os.O_EXCL, 0644)
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
