package results

// Result contains the results from the analyzation
type Result struct {
	SampleResults    []SampleResult
	LociResults      map[string][]LociResult
	LociNamesInOrder []string

	ReplicateAmount    int
	SampleAmount       int
	SingleSampleAmount int

	AmountOfLoci                       int
	AmountOfErroneousLoci              int
	AmountOfErroneousAlleles           int
	AmountOfAlleles                    int
	AmountOfAllelesForErrorCalculation int
}
