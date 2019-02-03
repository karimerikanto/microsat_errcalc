package analyzer

import (
	"errors"
	"karim/microsatellite_analyzer/analyzer/models"
	"karim/microsatellite_analyzer/analyzer/results"
)

//GetResultFromImportData analyzes given import data and returns the result object
func GetResultFromImportData(importData *models.ImportData) (results.Result, error) {
	var result results.Result
	result.LociResults = make(map[string][]results.LociResult)
	sampleArray := importData.Samples

	if len(sampleArray) == 0 {
		return result, nil
	}

	//Validate loci names
	lociNameCounts := make(map[string]bool)
	firstSample := sampleArray[0]
	firstSampleFirstReplica := firstSample.ReplicaArray[0]

	for _, locus := range firstSampleFirstReplica.LocusArray {
		if _, found := lociNameCounts[locus.Name]; found {
			return result, errors.New("Multiple same loci names found. Loci names must be unique")
		}

		lociNameCounts[locus.Name] = true
	}

	result.ReplicateAmount = 0
	result.SingleSampleAmount = 0
	result.AmountOfAllelesForErrorCalculation = 0
	result.SampleAmount = len(sampleArray)

	lociOrdersFilled := false

	for sampleArrayIndex := 0; sampleArrayIndex < len(sampleArray); sampleArrayIndex++ {
		sample := sampleArray[sampleArrayIndex]

		var sampleResult results.SampleResult
		sampleResult.Index = sampleArrayIndex + 1

		result.ReplicateAmount += len(sample.ReplicaArray)
		sampleResult.Single = sample.IsSingle()

		//Fill loci orders
		if !lociOrdersFilled {
			for _, replica := range sample.ReplicaArray {
				for _, locus := range replica.LocusArray {
					result.LociOrder = append(result.LociOrder, locus.Name)
				}

				break
			}

			lociOrdersFilled = true
		}

		//Skip single samples
		if sampleResult.Single {
			result.SingleSampleAmount++

			for _, replica := range sample.ReplicaArray {
				for _, locus := range replica.LocusArray {
					if len(locus.Allele1) > 0 {
						result.AmountOfAlleles++
					}
					if len(locus.Allele2) > 0 {
						result.AmountOfAlleles++
					}

					if len(locus.Allele1) > 0 && len(locus.Allele2) > 0 {
						result.AmountOfLoci++
					}
				}
			}

			result.SampleResults = append(result.SampleResults, sampleResult)

			continue
		}

		// Flip locuses and replicas
		locusArrayLength := len(sample.ReplicaArray[0].LocusArray)
		flippedSampleLocusArray := make([][]models.Locus, locusArrayLength)

		for locusIndex := 0; locusIndex < locusArrayLength; locusIndex++ {
			flippedSampleLocusArray[locusIndex] = make([]models.Locus, len(sample.ReplicaArray))
		}

		for replicaIndex := 0; replicaIndex < len(sample.ReplicaArray); replicaIndex++ {
			replica := sample.ReplicaArray[replicaIndex]

			for locusIndex := 0; locusIndex < len(replica.LocusArray); locusIndex++ {
				flippedSampleLocusArray[locusIndex][replicaIndex] = replica.LocusArray[locusIndex]
			}
		}

		// Create loci results for samples and for the whole result
		for _, loci := range flippedSampleLocusArray {
			if len(loci) > 0 {
				locusName := loci[0].Name
				lociResult := results.CreateLociResult(locusName, loci)
				sampleResult.LociResults = append(sampleResult.LociResults, lociResult)

				result.AmountOfAlleles += lociResult.AmountOfAlleles
				result.AmountOfAllelesForErrorCalculation += lociResult.AmountOfAlleles
				result.AmountOfErroneousAlleles += lociResult.AmountOfErroneousAlleles
				result.AmountOfLoci += len(loci)

				if _, ok := result.LociResults[locusName]; ok {
					result.LociResults[locusName] = append(result.LociResults[locusName], lociResult)
				} else {
					result.LociResults[locusName] = []results.LociResult{lociResult}
				}
			}
		}

		result.SampleResults = append(result.SampleResults, sampleResult)
	}

	return result, nil
}
