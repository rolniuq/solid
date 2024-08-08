package main

import (
	"fmt"
	"os"
)

type GoodReport struct {
	title   string
	content string
}

func (f *GoodReport) Generate() string {
	return fmt.Sprintf("title: %s\ncontent: %s", f.title, f.content)
}

type GoodReportSaver struct{}

func (f *GoodReportSaver) Save(report, fileName string) error {
	return os.WriteFile(fileName, []byte(report), 0644)
}

func goodInit() {
	g := GoodReport{
		title:   "quynh",
		content: "quynh",
	}

	r := g.Generate()

	s := GoodReportSaver{}
	s.Save(r, "test.txt")
}
