package writer

import (
	"karim/microsatellite_analyzer/analyzer/results"
	"strconv"
	"strings"
)

// WriteResultsToString writes results to string from given file path
func WriteResultsToString(result results.Result) string {
	fileContents := ""

	var rows = []string{
		"Replicate amount\t\t\t\t\t" + strconv.Itoa(result.ReplicateAmount),
		"Sample amount\t\t\t\t\t" + strconv.Itoa(result.SampleAmount),
		"Single sample amount\t\t\t\t" + strconv.Itoa(result.SingleSampleAmount),
		"",
		"Total amount of alleles\t\t\t\t" + strconv.Itoa(result.AmountOfAlleles),
		"Amount of alleles for error calculation\t\t\t" + strconv.Itoa(result.AmountOfAllelesForErrorCalculation),
		"Erroneous alleles\t\t\t\t\t" + strconv.Itoa(result.AmountOfErroneousAlleles),
		"Error rate\t\t\t\t\t\t" + strings.Replace(strconv.FormatFloat(result.ErrorRate, 'f', 6, 64), ".", ",", -1),
		"",
		"Amount of heterozygotes\t\t\t\t" + strconv.Itoa(result.AmountOfHeterozygotes),
		"Amount of homozygotes\t\t\t\t" + strconv.Itoa(result.AmountOfHomozygotes),
		"Allele dropout rate per heterozygous loci\t\t" + strings.Replace(strconv.FormatFloat(result.AlleleDropoutRatePerHeterozygousLoci, 'f', 6, 64), ".", ",", -1),
		"Allele dropout rate per all loci\t\t\t\t" + strings.Replace(strconv.FormatFloat(result.AlleleDropoutRatePerAllLoci, 'f', 6, 64), ".", ",", -1),
		"Allele dropout rate per all alleles\t\t\t" + strings.Replace(strconv.FormatFloat(result.AlleleDropoutRatePerAllAlleles, 'f', 6, 64), ".", ",", -1),
		"Other error rate per all loci\t\t\t\t" + strings.Replace(strconv.FormatFloat(result.OtherErrorRatePerAllLoci, 'f', 6, 64), ".", ",", -1),
		"Other error rate per all alleles\t\t\t\t" + strings.Replace(strconv.FormatFloat(result.OtherErrorRatePerAllAlleles, 'f', 6, 64), ".", ",", -1),
		"",
		"Loci total count\t\t\t\t\t" + strconv.Itoa(result.AmountOfLoci),
		"Loci total count for error calculation\t\t\t" + strconv.Itoa(result.AmountOfLociForErrorCalculation),
		"Loci error count\t\t\t\t\t" + strconv.Itoa(result.AmountOfErroneousLoci),
		"Loci error\t\t\t\t\t\t" + strings.Replace(strconv.FormatFloat(result.LociErrorRate, 'f', 6, 64), ".", ",", -1),
		"",
		"SAVE TO OUTPUT FILE TO GET ALL RESULTS",
	}

	for _, row := range rows {
		fileContents += row + "\r\n"
	}

	return fileContents
}
