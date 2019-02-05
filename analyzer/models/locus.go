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

// IsHomozygot indicates that is this locus a homozygot (both alleles are same and they are not empty and not question marks)
func (locus Locus) IsHomozygot() bool {
	return !locus.IsEmpty() &&
		locus.Allele1 != "?" &&
		locus.Allele2 != "?" &&
		locus.Allele1 == locus.Allele2
}

// IsHeterozygot indicates that is this locus a heterozygot (both alleles are different and they are not empty and not question marks)
func (locus Locus) IsHeterozygot() bool {
	return !locus.IsEmpty() &&
		locus.Allele1 != "?" &&
		locus.Allele2 != "?" &&
		locus.Allele1 != locus.Allele2
}
