package main

import (
	"fmt"
	"karim/microsatellite_analyzer/analyzer"
	"karim/microsatellite_analyzer/reader"
	"karim/microsatellite_analyzer/writer"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const versionName = "0.1"
const versionNumber = 1

func main() {
	currentDirectory := filepath.Dir(os.Args[0])

	if len(os.Args) < 2 {
		printHelp()
		return
	}
	//TODO: Add more dynamic parameter loop

	fmt.Println("Reading file...")
	fmt.Println("")

	readStartTime := time.Now()
	filepath := os.Args[1]
	importData, err := reader.ReadCsvFileToSampleMatrix(filepath)
	readElapsedTime := time.Since(readStartTime)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Got %v samples from the file in %v seconds\n", len(importData.Samples), strconv.FormatFloat(readElapsedTime.Seconds(), 'f', 3, 64))
	fmt.Println("")
	fmt.Println("Analyzing...")
	fmt.Println("")

	analyzeStartTime := time.Now()
	result, err := analyzer.GetResultFromImportData(importData)
	analyzeElapsedTime := time.Since(analyzeStartTime)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Analyze completed in %v seconds\n", strconv.FormatFloat(analyzeElapsedTime.Seconds(), 'f', 3, 64))

	if len(os.Args) > 2 {
		fmt.Println("")
		outputFileName := os.Args[2]
		writer.WriteResultsToCsvFile(currentDirectory+"\\"+outputFileName, result)

	} else {
		fmt.Println("")
		writer.WriteResultsToConsole(result)
	}
}

func printHelp() {
	fmt.Println("Microsatellite analyzer (version " + versionName + ")")
	fmt.Println("")
	fmt.Println("=== Usage ===")
	fmt.Println("microsatellite_analyzer.exe <source file path> <output file name (optional)>")
	fmt.Println("")
	fmt.Println("=== Notes ===")
	fmt.Println("Source file should be in csv format (';' as a separator). ',' is a forbidden character.")
	fmt.Println("Results are printed to the console by default. Results can be saved as a csv file if the output file is defined.")
	fmt.Println("There is no limitations for columns and rows. You can have your custom text at the beginning of the file and empty lines between the headers and the loci. An empty line between the samples will separate them and create new replicas.")
	fmt.Println("")
	fmt.Println("=== Example file ===")
	fmt.Println("My custom text 1 (optional);")
	fmt.Println("My custom text 2 (optional);")
	fmt.Println("")
	fmt.Println(";Locus1;;Locus2;;Locus3;")
	fmt.Println("Sample1;100;200;100;200;100;200;")
	fmt.Println("Sample2;100;200;100;200;100;200;")
	fmt.Println("")
	fmt.Println("Sample1;100;200;100;200;100;200;")
	fmt.Println("Sample1;100;200;100;200;100;200;")
	fmt.Println("Sample1;100;200;100;200;100;200;")
	fmt.Println("")
	fmt.Println("Sample2;100;200;100;200;100;200;")
	fmt.Println("")
}
