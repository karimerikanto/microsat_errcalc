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

	sep := " "

	var rows = []string{
		"Replicate amount" + sep + strconv.Itoa(result.ReplicateAmount),
		"Sample amount" + sep + strconv.Itoa(result.SampleAmount),
		"Single sample amount" + sep + strconv.Itoa(result.SingleSampleAmount),
		"",
		"Total amount of alleles" + sep + strconv.Itoa(result.AmountOfAlleles),
		"Amount of alleles for error calculation" + sep + strconv.Itoa(result.AmountOfAllelesForErrorCalculation),
		"Erroneous alleles" + sep + strconv.Itoa(result.AmountOfErroneousAlleles),
		"Error rate" + sep + strconv.FormatFloat((float64(result.AmountOfErroneousAlleles)/float64(result.AmountOfAllelesForErrorCalculation)), 'f', 6, 64),
		"",
		"Samples" + sep + "Error count" + sep + "Total" + sep + "Rate",
		"",
	}

	for _, sampleResult := range result.SampleResults {
		if sampleResult.Single {
			rows = append(rows, "Sample "+strconv.Itoa(sampleResult.Index)+" (Single)")

		} else {
			sampleErrors := 0
			sampleAlleles := 0

			var ambiguousLociResults []results.LociResult

			for _, lociResult := range sampleResult.LociResults {
				sampleErrors += lociResult.AmountOfErroneousAlleles
				sampleAlleles += lociResult.AmountOfAlleles

				if lociResult.Ambiguous {
					ambiguousLociResults = append(ambiguousLociResults, lociResult)
				}
			}

			rows = append(rows,
				"Sample "+strconv.Itoa(sampleResult.Index)+sep+
					strconv.Itoa(sampleErrors)+sep+
					strconv.Itoa(sampleAlleles)+sep+
					strconv.FormatFloat(float64(sampleErrors)/float64(sampleAlleles), 'f', 6, 64))

			for _, ambiguousLociResult := range ambiguousLociResults {
				rows = append(rows, "AMBIGUOUS LOCI RESULT ("+ambiguousLociResult.Name+")")
			}
		}

		rows = append(rows, "")
	}

	rows = append(rows, "Loci")
	rows = append(rows, "")

	for _, lociName := range result.LociOrder {
		lociResultGroup := result.LociResults[lociName]
		lociErrors := 0
		lociAlleles := 0

		for _, lociResult := range lociResultGroup {
			lociErrors += lociResult.AmountOfErroneousAlleles
			lociAlleles += lociResult.AmountOfAlleles
		}

		rows = append(rows,
			"Loci "+lociName+sep+
				strconv.Itoa(lociErrors)+sep+
				strconv.Itoa(lociAlleles)+sep+
				strconv.FormatFloat(float64(lociErrors)/float64(lociAlleles), 'f', 6, 64))

		rows = append(rows, "")
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
