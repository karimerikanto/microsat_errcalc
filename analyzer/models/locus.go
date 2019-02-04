package models

// Locus includes it's locus name and the allele values
type Locus struct {
	Name    string
	Allele1 string
	Allele2 string
}

// IsEmpty returns a state that is both alleles empty
func (locus Locus) IsEmpty() bool {
	return len(locus.Allele1) == 0 && len(locus.Allele2) == 0
}
