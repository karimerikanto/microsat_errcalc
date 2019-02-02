package main

import (
	"fmt"
	"io/ioutil"
	"karim/microsatellite_analyzer/analyzer"
	"karim/microsatellite_analyzer/analyzer/results"
	"karim/microsatellite_analyzer/reader"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	filepath := os.Args[1]
	importData, err := reader.ReadCsvFileToSampleMatrix(filepath)

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := analyzer.GetResultFromImportData(importData)

	if err != nil {
		fmt.Println(err)
		return
	}

	resultLines := getResultLinesFromResult(result)

	if len(os.Args) > 2 {
		outputFileName := os.Args[2]
		saveResultsToCsvFile(resultLines, currentDirectory+"\\"+outputFileName)

	} else {
		printResultLinesToConsole(resultLines)
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

func printResultLinesToConsole(resultLines [][]string) {
	for _, resultLine := range resultLines {
		fmt.Println(strings.Join(resultLine, " "))
	}
}

func saveResultsToCsvFile(resultLines [][]string, filePath string) {
	if !strings.HasSuffix(filePath, ".csv") {
		filePath += ".csv"
	}

	//TODO: Check the file path for security

	fileContents := ""

	for _, resultLine := range resultLines {
		fileContents += strings.Join(resultLine, ";") + "\n"
	}

	fmt.Println("Writing output file to " + filePath)
	err := ioutil.WriteFile(filePath, []byte(fileContents), 0777)

	if err != nil {
		fmt.Println("Couldn't write the output file. Check the error.")
		fmt.Println(err)
	}
}

func getResultLinesFromResult(result results.Result) [][]string {
	var resultLines = [][]string{
		{
			"Replicate amount", strconv.Itoa(result.ReplicateAmount),
		},
		{
			"Sample amount", strconv.Itoa(result.SampleAmount),
		},
		{
			"Single sample amount", strconv.Itoa(result.SingleSampleAmount),
		},
		{},
	}

	for _, sampleResult := range result.SampleResults {
		if sampleResult.Single {
			resultLines = append(resultLines, []string{
				"Replicate " + strconv.Itoa(sampleResult.Index) + " (Single)",
			})
		} else {
			resultLines = append(resultLines, []string{
				"Replicate " + strconv.Itoa(sampleResult.Index),
			})
		}

		resultLines = append(resultLines, []string{})
	}

	return resultLines
}
