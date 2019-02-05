package results

// SampleResult contains the results from one sample
type SampleResult struct {
	Index       int
	Single      bool
	LociResults []LociResult
}

// GetErrorRate returns error rate from erroronous alleles and total amount of alleles for error calculation
func (sampleResult SampleResult) GetErrorRate() float64 {

	errors := 0
	alleles := 0

	for _, lociResult := range sampleResult.LociResults {
		errors += lociResult.AmountOfErroneousAlleles
		alleles += lociResult.AmountOfAllelesForErrorCalculation
	}

	return float64(errors) / float64(alleles)
}
