package results

// Result contains the results from the analyzation
type Result struct {
	SampleResults      []SampleResult
	ReplicateAmount    int
	SampleAmount       int
	SingleSampleAmount int
}
