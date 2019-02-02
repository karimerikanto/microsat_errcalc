package results

// SampleResult contains the results from one sample
type SampleResult struct {
	Index  int
	Errors int
	Total  int
	Single bool
}
