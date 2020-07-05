package main

import (
	"fmt"
	"github.com/DABronskikh/go-lesson-8/pkg/transactions"
	"io"
	"log"
	"os"
)

func main() {
	const filenameCSV = "demoFile.csv"
	const filenameJSON = "demoFile.json"
	const filenameXML = "demoFile.xml"

	svc := transactions.NewService()
	for i := 0; i < 20; i++ {
		_, err := svc.Register("001", "002", 1000_00)
		if err != nil {
			log.Print(err)
			return
		}
	}

	// CSV
	if err := demoExportCSV(svc, filenameCSV); err != nil {
		os.Exit(1)
	}

	demoImportCSV := transactions.NewService()
	if err := demoImportCSV.ImportCSV(filenameCSV); err != nil {
		os.Exit(1)
	}

	fmt.Println("demoImportCSV = ", demoImportCSV)

	//JSON
	if err := svc.ExportJSON(filenameJSON); err != nil {
		os.Exit(1)
	}

	demoImportJSON := transactions.NewService()
	if err := demoImportJSON.ImportJSON(filenameJSON); err != nil {
		os.Exit(1)
	}

	fmt.Println("demoImportJSON = ", demoImportJSON)

	// XML
	transactionsXML := &transactions.Transactions{
		Transactions: svc.Transactions,
	}

	if err := transactionsXML.ExportXML(filenameXML); err != nil {
		os.Exit(1)
	}

	demoImportXML := &transactions.Transactions{}
	if err := demoImportXML.ImportXML(filenameXML); err != nil {
		os.Exit(1)
	}

	fmt.Println("demoImportXML = ", demoImportXML)
}

func demoExportCSV(svc *transactions.Service, filename string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Print(err)
		}
	}(file)

	err = svc.ExportCSV(file)
	if err != nil {
		log.Print(err)
		return
	}

	return nil
}
