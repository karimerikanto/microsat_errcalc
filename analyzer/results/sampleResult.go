package results

// SampleResult contains the results from one sample
type SampleResult struct {
	Index       int
	Single      bool
	LociResults []LociResult
}
