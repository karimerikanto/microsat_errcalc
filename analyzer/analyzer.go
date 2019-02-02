package analyzer

import (
	"karim/microsatellite_analyzer/analyzer/models"
	"karim/microsatellite_analyzer/analyzer/results"
)

//GetResultFromImportData analyzes given import data and returns the result object
func GetResultFromImportData(importData *models.ImportData) (results.Result, error) {
	var result results.Result

	sampleArray := importData.Samples

	if len(sampleArray) == 0 {
		return result, nil
	}

	result.ReplicateAmount = 0
	result.SingleSampleAmount = 0
	result.SampleAmount = len(sampleArray)

	for sampleArrayIndex := 0; sampleArrayIndex < len(sampleArray); sampleArrayIndex++ {
		sample := sampleArray[sampleArrayIndex]

		var sampleResult results.SampleResult
		sampleResult.Index = sampleArrayIndex + 1

		result.ReplicateAmount += len(sample.ReplicaArray)
		sampleResult.Single = sample.IsSingle()

		if sampleResult.Single {
			result.SingleSampleAmount++
			result.SampleResults = append(result.SampleResults, sampleResult)
			continue
		}

		// Flip locuses and replicas
		/*locusArrayLength := len(sampleArray[0].LocusArray)
		flippedSampleLocusArray := make([][]models.Locus, locusArrayLength)

		for locusIndex := 0; locusIndex < locusArrayLength; locusIndex++ {
			flippedSampleLocusArray[locusIndex] = make([]models.Locus, len(sampleArray))
		}

		for sampleIndex := 0; sampleIndex < len(sampleArray); sampleIndex++ {
			sample := sampleArray[sampleIndex]

			for locusIndex := 0; locusIndex < len(sample.LocusArray); locusIndex++ {
				flippedSampleLocusArray[locusIndex][sampleIndex] = sample.LocusArray[locusIndex]
			}
		}

		// Analyze flipped array
		for locusIndex := 0; locusIndex < len(flippedSampleLocusArray); locusIndex++ {
			flippedSampleLocusArray[locusIndex] = make([]models.Locus, len(sampleArray))
		}*/

		result.SampleResults = append(result.SampleResults, sampleResult)
	}

	return result, nil
}
