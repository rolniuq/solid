package main

import (
	"fmt"
	"os"
)

type BadReport struct {
	title   string
	content string
}

func (f *BadReport) Generate() string {
	return fmt.Sprintf("Title: %s\nContent: %s", f.title, f.content)
}

func (f *BadReport) Save(fileName string) error {
	report := f.Generate()
	return os.WriteFile(fileName, []byte(report), 0644)
}

func badInit() {
	b := BadReport{
		title:   "quynh",
		content: "quynh",
	}

	data := b.Generate()
	fmt.Println(data)

	err := b.Save("test.txt")
	if err != nil {
		fmt.Errorf("error", err)
	}
}
