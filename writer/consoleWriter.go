package writer

import (
	"fmt"
	"karim/microsatellite_analyzer/analyzer/results"
	"strconv"
)

// WriteResultsToConsole writes given results to console
func WriteResultsToConsole(result results.Result) {

	fmt.Println("Printing results to console...")
	fmt.Println("")

	var rows = []string{
		"Replicate amount " + strconv.Itoa(result.ReplicateAmount),
		"Sample amount " + strconv.Itoa(result.SampleAmount),
		"Single sample amount " + strconv.Itoa(result.SingleSampleAmount),
		"",
		"Total amount of alleles " + strconv.Itoa(result.AmountOfAlleles),
		"Amount of alleles for error calculation " + strconv.Itoa(result.AmountOfAllelesForErrorCalculation),
		"Erroneous alleles " + strconv.Itoa(result.AmountOfErroneousAlleles),
		"Error rate " + strconv.FormatFloat((float64(result.AmountOfErroneousAlleles)/float64(result.AmountOfAllelesForErrorCalculation)), 'f', 6, 64),
		"",
		"SAMPLES Allele drop outs Allele drop outs rate Other errors Other errors rate Total error count Total error rate Total alleles",
		"",
	}

	for _, sampleResult := range result.SampleResults {
		if sampleResult.Single {
			rows = append(rows, "Sample "+sampleResult.Name+" (Single)")

		} else {
			sampleErrors := 0
			sampleAlleleDropOuts := 0
			sampleAlleles := 0

			var ambiguousLociResults []results.LociResult

			for _, lociResult := range sampleResult.LociResults {
				sampleErrors += lociResult.AmountOfErroneousAlleles
				sampleAlleleDropOuts += lociResult.AmountOfAlleleDropOuts
				sampleAlleles += lociResult.TotalAmountOfAlleles

				if lociResult.Ambiguous {
					ambiguousLociResults = append(ambiguousLociResults, lociResult)
				}
			}

			rows = append(rows,
				"Sample "+sampleResult.Name+" "+
					strconv.Itoa(sampleAlleleDropOuts)+" "+
					strconv.FormatFloat(float64(sampleAlleleDropOuts)/float64(sampleAlleles), 'f', 6, 64)+" "+
					strconv.Itoa(sampleErrors-sampleAlleleDropOuts)+" "+
					strconv.FormatFloat(float64(sampleErrors-sampleAlleleDropOuts)/float64(sampleAlleles), 'f', 6, 64)+" "+
					strconv.Itoa(sampleErrors)+" "+
					strconv.FormatFloat(float64(sampleErrors)/float64(sampleAlleles), 'f', 6, 64)+" "+
					strconv.Itoa(sampleAlleles))

			for _, ambiguousLociResult := range ambiguousLociResults {
				rows = append(rows, "AMBIGUOUS LOCI RESULT ("+ambiguousLociResult.Name+")")
			}
		}

		rows = append(rows, "")
	}

	rows = append(rows, "Loci")
	rows = append(rows, "")

	for _, lociName := range result.LociNamesInOrder {
		lociResultGroup := result.LociResults[lociName]
		lociErrors := 0
		lociAlleleDropOuts := 0
		lociAlleles := 0

		for _, lociResult := range lociResultGroup {
			lociErrors += lociResult.AmountOfErroneousAlleles
			lociAlleleDropOuts += lociResult.AmountOfAlleleDropOuts
			lociAlleles += lociResult.AmountOfAllelesForErrorCalculation
		}

		rows = append(rows,
			"Loci "+lociName+" "+
				strconv.Itoa(lociAlleleDropOuts)+" "+
				strconv.FormatFloat(float64(lociAlleleDropOuts)/float64(lociAlleles), 'f', 6, 64)+" "+
				strconv.Itoa(lociErrors-lociAlleleDropOuts)+" "+
				strconv.FormatFloat(float64(lociErrors-lociAlleleDropOuts)/float64(lociAlleles), 'f', 6, 64)+" "+
				strconv.Itoa(lociErrors)+" "+
				strconv.FormatFloat(float64(lociErrors)/float64(lociAlleles), 'f', 6, 64)+" "+
				strconv.Itoa(lociAlleles))

		rows = append(rows, "")
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
