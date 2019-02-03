package results

// Result contains the results from the analyzation
type Result struct {
	SampleResults []SampleResult
	LociResults   map[string][]LociResult
	LociOrder     []string

	ReplicateAmount    int
	SampleAmount       int
	SingleSampleAmount int

	AmountOfLoci                       int
	AmountOfErroneousAlleles           int
	AmountOfAlleles                    int
	AmountOfAllelesForErrorCalculation int
}
