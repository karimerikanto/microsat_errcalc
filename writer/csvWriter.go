package writer

import (
	"fmt"
	"io/ioutil"
	"karim/microsatellite_analyzer/analyzer/results"
	"strconv"
	"strings"
)

// WriteResultsToCsvFile writer results to csv file in given file path
func WriteResultsToCsvFile(filePath string, result results.Result) {
	if !strings.HasSuffix(filePath, ".csv") {
		filePath += ".csv"
	}

	//TODO: Check the file path for security

	fileContents := ""

	var rows = []string{
		"Replicate amount;" + strconv.Itoa(result.ReplicateAmount),
		"Sample amount;" + strconv.Itoa(result.SampleAmount),
		"Single sample amount;" + strconv.Itoa(result.SingleSampleAmount),
		"",
		"Total amount of alleles;" + strconv.Itoa(result.AmountOfAlleles),
		"Amount of alleles for error calculation;" + strconv.Itoa(result.AmountOfAllelesForErrorCalculation),
		"Erroneous alleles;" + strconv.Itoa(result.AmountOfErroneousAlleles),
		"Error rate;" + strings.Replace(strconv.FormatFloat(result.ErrorRate, 'f', 6, 64), ".", ",", -1),
		"",
		"Amount of heterozygotes;" + strconv.Itoa(result.AmountOfHeterozygotes),
		"Amount of homozygotes;" + strconv.Itoa(result.AmountOfHomozygotes),
		"Allele dropout rate per heterozygous loci;" + strings.Replace(strconv.FormatFloat(result.AlleleDropoutRatePerHeterozygousLoci, 'f', 6, 64), ".", ",", -1),
		"Allele dropout rate per all loci;" + strings.Replace(strconv.FormatFloat(result.AlleleDropoutRatePerAllLoci, 'f', 6, 64), ".", ",", -1),
		"Allele dropout rate per all alleles;" + strings.Replace(strconv.FormatFloat(result.AlleleDropoutRatePerAllAlleles, 'f', 6, 64), ".", ",", -1),
		"Other error rate per all loci;" + strings.Replace(strconv.FormatFloat(result.OtherErrorRatePerAllLoci, 'f', 6, 64), ".", ",", -1),
		"Other error rate per all alleles;" + strings.Replace(strconv.FormatFloat(result.OtherErrorRatePerAllAlleles, 'f', 6, 64), ".", ",", -1),
		"",
		"Loci total count;" + strconv.Itoa(result.AmountOfLoci),
		"Loci total count for error calculation;" + strconv.Itoa(result.AmountOfLociForErrorCalculation),
		"Loci error count;" + strconv.Itoa(result.AmountOfErroneousLoci),
		"Loci error;" + strings.Replace(strconv.FormatFloat(result.LociErrorRate, 'f', 6, 64), ".", ",", -1),
		"",
		"SAMPLES;Allele drop outs;Allele drop outs rate;Other errors;Other errors rate;Total error count;Total error rate;Total alleles;Heterozygotes;Homozygotes;Allele dropout rate per heterozygous loci;Allele dropout rate per all loci;Other error rate per all loci;Amount of loci for error calculation;Amount of erroneous loci;",
	}

	for _, sampleResult := range result.SampleResults {
		if sampleResult.Single {
			rows = append(rows, "Sample "+sampleResult.Name+" (Single)")

		} else {
			sampleErrors := 0
			sampleAlleles := 0
			sampleAlleleDropOuts := 0
			sampleHeterozygotes := 0
			sampleHomozygotes := 0
			sampleAmountOfLociForErrorCalculation := 0
			sampleAmountOfErroneousLoci := 0

			var ambiguousLociResults []results.LociResult

			for _, lociResult := range sampleResult.LociResults {
				sampleErrors += lociResult.AmountOfErroneousAlleles
				sampleAlleles += lociResult.TotalAmountOfAlleles
				sampleAlleleDropOuts += lociResult.AmountOfAlleleDropOuts
				sampleHeterozygotes += lociResult.AmountOfHeterozygotes
				sampleHomozygotes += lociResult.AmountOfHomozygotes
				sampleAmountOfLociForErrorCalculation += lociResult.AmountOfLociForErrorCalculation
				sampleAmountOfErroneousLoci += lociResult.AmountOfErroneousLoci

				if lociResult.Ambiguous {
					ambiguousLociResults = append(ambiguousLociResults, lociResult)
				}
			}

			rows = append(rows,
				"Sample "+sampleResult.Name+";"+
					strconv.Itoa(sampleAlleleDropOuts)+";"+
					strings.Replace(strconv.FormatFloat(float64(sampleAlleleDropOuts)/float64(sampleAlleles), 'f', 6, 64), ".", ",", -1)+";"+
					strconv.Itoa(sampleErrors-sampleAlleleDropOuts)+";"+
					strings.Replace(strconv.FormatFloat(float64(sampleErrors-sampleAlleleDropOuts)/float64(sampleAlleles), 'f', 6, 64), ".", ",", -1)+";"+
					strconv.Itoa(sampleErrors)+";"+
					strings.Replace(strconv.FormatFloat(float64(sampleErrors)/float64(sampleAlleles), 'f', 6, 64), ".", ",", -1)+";"+
					strconv.Itoa(sampleAlleles)+";"+
					strconv.Itoa(sampleHeterozygotes)+";"+
					strconv.Itoa(sampleHomozygotes)+";"+
					strings.Replace(strconv.FormatFloat(float64(sampleAlleleDropOuts)/float64(sampleHeterozygotes), 'f', 6, 64), ".", ",", -1)+";"+
					strings.Replace(strconv.FormatFloat(float64(sampleAlleleDropOuts)/float64(sampleAmountOfLociForErrorCalculation), 'f', 6, 64), ".", ",", -1)+";"+
					strings.Replace(strconv.FormatFloat(float64(sampleAmountOfErroneousLoci-sampleAlleleDropOuts)/float64(sampleAmountOfLociForErrorCalculation), 'f', 6, 64), ".", ",", -1)+";"+
					strconv.Itoa(sampleAmountOfLociForErrorCalculation)+";"+
					strconv.Itoa(sampleAmountOfErroneousLoci))

			for _, ambiguousLociResult := range ambiguousLociResults {
				rows = append(rows, "AMBIGUOUS LOCI RESULT ("+ambiguousLociResult.Name+")")
			}
		}
	}

	rows = append(rows, "")
	rows = append(rows, "LOCI")

	for _, lociName := range result.LociNamesInOrder {
		lociResultGroup := result.LociResults[lociName]
		lociErrors := 0
		lociAlleleDropOuts := 0
		lociAlleles := 0
		lociHeterozygotes := 0
		lociHomozygotes := 0
		lociAmountOfLociForErrorCalculation := 0
		lociAmountOfErroneousLoci := 0

		for _, lociResult := range lociResultGroup {
			lociErrors += lociResult.AmountOfErroneousAlleles
			lociAlleles += lociResult.AmountOfAllelesForErrorCalculation
			lociAlleleDropOuts += lociResult.AmountOfAlleleDropOuts
			lociHeterozygotes += lociResult.AmountOfHeterozygotes
			lociHomozygotes += lociResult.AmountOfHomozygotes
			lociAmountOfLociForErrorCalculation += lociResult.AmountOfLociForErrorCalculation
			lociAmountOfErroneousLoci += lociResult.AmountOfErroneousLoci
		}

		rows = append(rows,
			"Loci "+lociName+";"+
				strconv.Itoa(lociAlleleDropOuts)+";"+
				strings.Replace(strconv.FormatFloat(float64(lociAlleleDropOuts)/float64(lociAlleles), 'f', 6, 64), ".", ",", -1)+";"+
				strconv.Itoa(lociErrors-lociAlleleDropOuts)+";"+
				strings.Replace(strconv.FormatFloat(float64(lociErrors-lociAlleleDropOuts)/float64(lociAlleles), 'f', 6, 64), ".", ",", -1)+";"+
				strconv.Itoa(lociErrors)+";"+
				strings.Replace(strconv.FormatFloat(float64(lociErrors)/float64(lociAlleles), 'f', 6, 64), ".", ",", -1)+";"+
				strconv.Itoa(lociAlleles)+";"+
				strconv.Itoa(lociHeterozygotes)+";"+
				strconv.Itoa(lociHomozygotes)+";"+
				strings.Replace(strconv.FormatFloat(float64(lociAlleleDropOuts)/float64(lociHeterozygotes), 'f', 6, 64), ".", ",", -1)+";"+
				strings.Replace(strconv.FormatFloat(float64(lociAlleleDropOuts)/float64(lociAmountOfLociForErrorCalculation), 'f', 6, 64), ".", ",", -1)+";"+
				strings.Replace(strconv.FormatFloat(float64(lociAmountOfErroneousLoci-lociAlleleDropOuts)/float64(lociAmountOfLociForErrorCalculation), 'f', 6, 64), ".", ",", -1)+";"+
				strconv.Itoa(lociAmountOfLociForErrorCalculation)+";"+
				strconv.Itoa(lociAmountOfErroneousLoci))
	}

	for _, row := range rows {
		fileContents += row + "\n"
	}

	fmt.Println("Writing output file to " + filePath)
	err := ioutil.WriteFile(filePath, []byte(fileContents), 0777)

	if err != nil {
		fmt.Println("Couldn't write the output file. Check the error.")
		fmt.Println(err)
	}
}
