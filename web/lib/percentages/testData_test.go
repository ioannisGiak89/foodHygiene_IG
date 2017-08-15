package percentages

import (
	"github.com/ioannisGiak89/foodHygiene_IG/web/model"
	"github.com/ioannisGiak89/foodHygiene_IG/web/model/fsaResponse"
)

// FHIS Test Data
var testFhisEstablishmentRatingWithEmptyValues = []fsaResponse.Establishment{
	{RatingValue: "Pass"},
	{RatingValue: "Pass"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: ""},
	{RatingValue: "Improvement Required"},
	{RatingValue: "Awaiting Inspection"},
	{RatingValue: ""},
	{RatingValue: ""},
	{RatingValue: "Awaiting Inspection"},
}

var correctFhisPercentagesEmptyValues = model.FhisPercentages{
	Pass:                "20.00%",
	PassAndEatSafe:      "20.00%",
	ImprovementRequired: "10.00%",
	AwaitingInspection:  "20.00%",
	Exempt:              "0.00%",
}

var testFhisEstablishmentRatingData = []fsaResponse.Establishment{
	{RatingValue: "Pass"},
	{RatingValue: "Pass"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Improvement Required"},
	{RatingValue: "Awaiting Inspection"},
	{RatingValue: "Exempt"},
	{RatingValue: "Pass"},
	{RatingValue: "Awaiting Inspection"},
}

var correctFhisPercentages = model.FhisPercentages{
	Pass:                "30.00%",
	PassAndEatSafe:      "30.00%",
	ImprovementRequired: "10.00%",
	AwaitingInspection:  "20.00%",
	Exempt:              "10.00%",
}

var testFhisEstablishmentRatingWithoutExempt = []fsaResponse.Establishment{
	{RatingValue: "Pass"},
	{RatingValue: "Pass"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Pass and Eat Safe"},
	{RatingValue: "Improvement Required"},
	{RatingValue: "Awaiting Inspection"},
	{RatingValue: "Pass"},
	{RatingValue: "Pass"},
	{RatingValue: "Awaiting Inspection"},
}

var correctFhisPercentagesWithoutExempt = model.FhisPercentages{
	Pass:                "40.00%",
	PassAndEatSafe:      "30.00%",
	ImprovementRequired: "10.00%",
	AwaitingInspection:  "20.00%",
	Exempt:              "0.00%",
}

// FHRS Test Data
var testFhrsEstablishmentRatingWithEmptyValues = []fsaResponse.Establishment{
	{RatingValue: "1"},
	{RatingValue: "4"},
	{RatingValue: "Exempt"},
	{RatingValue: "1"},
	{RatingValue: ""},
	{RatingValue: "4"},
	{RatingValue: "AwaitingInspection"},
	{RatingValue: ""},
	{RatingValue: ""},
	{RatingValue: "AwaitingInspection"},
}

var correctFhrsPercentagesEmptyValues = model.FhrsPercentages{
	OneStar:            "20.00%",
	TwoStar:            "0.00%",
	ThreeStar:          "0.00%",
	FourStar:           "20.00%",
	FiveStar:           "0.00%",
	AwaitingInspection: "20.00%",
	Exempt:             "10.00%",
}

var testFhrsEstablishmentRatingData = []fsaResponse.Establishment{
	{RatingValue: "1"},
	{RatingValue: "1"},
	{RatingValue: "2"},
	{RatingValue: "2"},
	{RatingValue: "2"},
	{RatingValue: "3"},
	{RatingValue: "AwaitingInspection"},
	{RatingValue: "Exempt"},
	{RatingValue: "4"},
	{RatingValue: "AwaitingInspection"},
}

var correctFhrsPercentages = model.FhrsPercentages{
	OneStar:            "20.00%",
	TwoStar:            "30.00%",
	ThreeStar:          "10.00%",
	FourStar:           "10.00%",
	FiveStar:           "0.00%",
	AwaitingInspection: "20.00%",
	Exempt:             "10.00%",
}

var testFhrsEstablishmentRatingWithoutExempt = []fsaResponse.Establishment{
	{RatingValue: "1"},
	{RatingValue: "1"},
	{RatingValue: "2"},
	{RatingValue: "2"},
	{RatingValue: "2"},
	{RatingValue: "3"},
	{RatingValue: "AwaitingInspection"},
	{RatingValue: "5"},
	{RatingValue: "4"},
	{RatingValue: "AwaitingInspection"},
}

var correctFhrsPercentagesWithoutExempt = model.FhrsPercentages{
	OneStar:            "20.00%",
	TwoStar:            "30.00%",
	ThreeStar:          "10.00%",
	FourStar:           "10.00%",
	FiveStar:           "10.00%",
	AwaitingInspection: "20.00%",
	Exempt:             "0.00%",
}
