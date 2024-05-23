package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	gtranslate "github.com/gilang-as/google-translate"
)

func main() {

	argsWithProg := os.Args

	var myFile = "translation_fr.csv"
	fileHandle, _ := os.Open(myFile)

	var myFile2 = argsWithProg[2]
	fileHandle2, err := os.OpenFile(myFile2, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		panic("Error: " + err.Error())
	}

	r := csv.NewReader(fileHandle)

	w := csv.NewWriter(fileHandle2)

	header := true
	cpt := 1

	var langFrom string
	var langTo string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if header {
			err = w.Write(record)
			if err != nil {
				panic("[" + strconv.Itoa(cpt) + "]  / Error: " + err.Error())
			}
			header = false
			langFrom = record[2]
			langTo = record[3]
		} else {
			//fmt.Println(record[2])
			value := gtranslate.Translate{
				Text: record[2],
				From: langFrom,
				To:   langTo,
			}

			fmt.Println("[" + strconv.Itoa(cpt) + "] to translate: " + value.Text)
			// fmt.Println(value)
			translated, err := gtranslate.Translator(value)
			if err != nil {
				fmt.Println("[" + strconv.Itoa(cpt) + "]  / Error: " + err.Error())
			}
			// fmt.Println("[" + strconv.Itoa(cpt) + "] Traduction: " + translated.Text)
			record[3] = translated.Text
			err = w.Write(record)
			if err != nil {
				panic("[" + strconv.Itoa(cpt) + "]  / Error: " + err.Error())
			}
		}
		cpt++
	}

	w.Flush()
	fileHandle.Close()
	fileHandle2.Close()

}
