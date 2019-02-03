package reader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"karim/microsatellite_analyzer/analyzer/models"
	"os"
)

//Read states
const (
	Headers    = 1
	LocusNames = 2
	Samples    = 3
)

//ReadCsvFileToSampleMatrix reads a csv file and returns a sample array and the file headers
func ReadCsvFileToSampleMatrix(filepath string) (*models.ImportData, error) {
	csvDataMatrix, err := readCsvFile(filepath)

	if err != nil {
		return nil, err
	}

	importData, err := readCsvDataMatrixToImportData(csvDataMatrix)

	if err != nil {
		return nil, err
	}

	return importData, nil
}

func readCsvFile(filepath string) ([][]string, error) {
	var lineMatrix [][]string

	csvFile, err := os.Open(filepath)
	defer csvFile.Close()

	if err != nil {
		fmt.Println("Couldn't read the file. Check the error.")
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';' //TODO: Add automatic recognition for the separator

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			if err != nil {
				fmt.Println("Couldn't read the file. Check the error.")
				return nil, err
			}
		}

		lineMatrix = append(lineMatrix, line)
	}

	return lineMatrix, err
}

func readCsvDataMatrixToImportData(dataLines [][]string) (*models.ImportData, error) {
	state := Headers
	var sampleArray []models.Sample
	var locusNames []string
	var headers [][]string
	var importData = models.NewImportData()
	previousSampleRowWasEmpty := true

	for dataLineIndex, dataLine := range dataLines {
		switch state {
		case Headers:
			if !isHeaderLine(dataLine) {
				state = LocusNames
				goto locusCase
			}

			headers = append(headers, dataLine)
			break
		locusCase:
			fallthrough
		case LocusNames:
			if !isEmpty(dataLine) && !isLocusRow(dataLine) {
				state = Samples
				goto sampleCase
			} else if isLocusRow(dataLine) {
				for _, val := range dataLine {
					if len(val) > 0 {
						locusNames = append(locusNames, val)
					}
				}
			}
			break
		sampleCase:
			fallthrough
		case Samples:
			//TODO: Read sample locuses in different matrix dimension for better performance and less memory usage?
			if isSampleRow(dataLine) {
				sampleName := dataLine[0]
				locusArray, err := readLocusArray(dataLine, locusNames, dataLineIndex)

				if err != nil {
					return importData, err
				}

				importData.AppendLoci(locusArray)

				if previousSampleRowWasEmpty {
					newSample := models.Sample{}
					sampleArray = append(sampleArray, newSample)
				}

				previousSampleRowWasEmpty = false
				replicaArray := &sampleArray[len(sampleArray)-1].ReplicaArray

				replica := models.Replica{Name: sampleName, LocusArray: locusArray}
				*replicaArray = append(*replicaArray, replica)
			} else if len(sampleArray) > 0 && isEmpty(dataLine) {
				previousSampleRowWasEmpty = true
			}
		}
	}

	importData.Samples = sampleArray
	importData.Headers = headers

	return importData, nil
}

func isEmpty(lineDatas []string) bool {
	for _, val := range lineDatas {
		if len(val) > 0 {
			return false
		}
	}

	return true
}

func isHeaderLine(lineDatas []string) bool {
	if isEmpty(lineDatas) || len(lineDatas[0]) == 0 {
		return false
	}

	return true
}

func isLocusRow(lineDatas []string) bool {
	if isEmpty(lineDatas) || len(lineDatas) < 2 || len(lineDatas[0]) > 0 || len(lineDatas[1]) == 0 {
		return false
	}

	return true
}

func isSampleRow(lineDatas []string) bool {
	if isEmpty(lineDatas) || len(lineDatas) < 2 {
		return false
	}

	return true
}

func readLocusArray(lineDatas []string, locusNames []string, lineDataIndex int) ([]models.Locus, error) {
	var locusArray []models.Locus

	for i := 0; i < len(locusNames); i++ {
		locusIndex := (i * 2) + 1

		locus := models.Locus{
			Name: locusNames[i],
		}

		if locusIndex < len(lineDatas) {
			locus.Allele1 = lineDatas[locusIndex]
			locus.Allele2 = lineDatas[locusIndex+1]
		}

		locusArray = append(locusArray, locus)
	}

	return locusArray, nil
}
