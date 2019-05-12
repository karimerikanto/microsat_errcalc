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
	AmountOfLociForErrorCalculation    int
	AmountOfErroneousLoci              int
	AmountOfErroneousAlleles           int
	AmountOfAlleles                    int
	AmountOfAllelesForErrorCalculation int
	AmountOfHeterozygotes              int
	AmountOfHomozygotes                int
	AmountOfAlleleDropouts             int
	AmountOfOtherErrors                int

	ErrorRate                            float64
	AlleleDropoutRatePerHeterozygousLoci float64
	AlleleDropoutRatePerAllLoci          float64
	AlleleDropoutRatePerAllAlleles       float64
	OtherErrorRatePerAllLoci             float64
	OtherErrorRatePerAllAlleles          float64
	LociErrorRate                        float64
}
