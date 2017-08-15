package model

const FHRS = "FHRS"

// FhrsPercentages
type FhrsPercentages struct {
	FiveStar           string
	FourStar           string
	ThreeStar          string
	TwoStar            string
	OneStar            string
	Exempt             string
	AwaitingInspection string
}

// FhisPercentages
type FhisPercentages struct {
	Pass                string
	PassAndEatSafe      string
	ImprovementRequired string
	AwaitingInspection  string
	Exempt              string
}
