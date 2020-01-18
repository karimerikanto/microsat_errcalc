package analyzer

import (
	"errors"
	"karim/microsat_errcalc/analyzer/models"
	"karim/microsat_errcalc/analyzer/results"
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
	result.AmountOfLociForErrorCalculation = 0
	result.AmountOfErroneousLoci = 0
	result.SampleAmount = len(sampleArray)

	lociOrdersFilled := false

	for sampleArrayIndex := 0; sampleArrayIndex < len(sampleArray); sampleArrayIndex++ {
		sample := sampleArray[sampleArrayIndex]

		var sampleResult results.SampleResult
		sampleResult.Name = "(" + sample.GetReplicaNames() + ")"

		result.ReplicateAmount += len(sample.ReplicaArray)
		sampleResult.Single = sample.IsSingle()

		//Fill loci orders
		if !lociOrdersFilled {
			for _, replica := range sample.ReplicaArray {
				for _, locus := range replica.LocusArray {
					result.LociNamesInOrder = append(result.LociNamesInOrder, locus.Name)
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
		flippedSampleLocusArray := getFlippedLocusArrayUsingReplicaArray(sample.ReplicaArray[0].LocusArray, sample.ReplicaArray)

		// Create loci results for samples and for the whole result
		for _, loci := range flippedSampleLocusArray {
			if len(loci) > 0 {
				locusName := loci[0].Name
				lociResult := results.CreateLociResult(locusName, loci)
				sampleResult.LociResults = append(sampleResult.LociResults, lociResult)

				result.AmountOfAlleles += lociResult.TotalAmountOfAlleles
				result.AmountOfAllelesForErrorCalculation += lociResult.AmountOfAllelesForErrorCalculation
				result.AmountOfErroneousAlleles += lociResult.AmountOfErroneousAlleles
				result.AmountOfLoci += lociResult.AmountOfLoci
				result.AmountOfLociForErrorCalculation += lociResult.AmountOfLociForErrorCalculation

				result.AmountOfErroneousLoci += lociResult.AmountOfErroneousLoci

				result.AmountOfHeterozygotes += lociResult.AmountOfHeterozygotes
				result.AmountOfHomozygotes += lociResult.AmountOfHomozygotes
				result.AmountOfAlleleDropouts += lociResult.AmountOfAlleleDropOuts
				result.AmountOfOtherErrors += lociResult.AmountOfErroneousAlleles - lociResult.AmountOfAlleleDropOuts

				if _, ok := result.LociResults[locusName]; ok {
					result.LociResults[locusName] = append(result.LociResults[locusName], lociResult)
				} else {
					result.LociResults[locusName] = []results.LociResult{lociResult}
				}
			}
		}

		result.SampleResults = append(result.SampleResults, sampleResult)

		result.ErrorRate = float64(result.AmountOfErroneousAlleles) / float64(result.AmountOfAllelesForErrorCalculation)
		result.AlleleDropoutRatePerHeterozygousLoci = float64(result.AmountOfAlleleDropouts) / float64(result.AmountOfHeterozygotes)
		result.AlleleDropoutRatePerAllLoci = float64(result.AmountOfAlleleDropouts) / float64(result.AmountOfLociForErrorCalculation)
		result.AlleleDropoutRatePerAllAlleles = float64(result.AmountOfAlleleDropouts) / float64(result.AmountOfAllelesForErrorCalculation)
		result.OtherErrorRatePerAllLoci = float64(result.AmountOfErroneousLoci-result.AmountOfAlleleDropouts) / float64(result.AmountOfLociForErrorCalculation)
		result.OtherErrorRatePerAllAlleles = float64(result.AmountOfOtherErrors) / float64(result.AmountOfAllelesForErrorCalculation)
		result.LociErrorRate = float64(result.AmountOfErroneousLoci) / float64(result.AmountOfLociForErrorCalculation)
	}

	return result, nil
}

func getFlippedLocusArrayUsingReplicaArray(loci []models.Locus, replicas []models.Replica) [][]models.Locus {

	locusArrayLength := len(loci)
	flippedSampleLocusArray := make([][]models.Locus, locusArrayLength)

	for locusIndex := 0; locusIndex < locusArrayLength; locusIndex++ {
		flippedSampleLocusArray[locusIndex] = make([]models.Locus, len(replicas))
	}

	for replicaIndex := 0; replicaIndex < len(replicas); replicaIndex++ {
		replica := replicas[replicaIndex]

		for locusIndex := 0; locusIndex < len(replica.LocusArray); locusIndex++ {
			flippedSampleLocusArray[locusIndex][replicaIndex] = replica.LocusArray[locusIndex]
		}
	}

	return flippedSampleLocusArray
}
