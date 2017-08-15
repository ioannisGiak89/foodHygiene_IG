package percentages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ioannisGiak89/foodHygiene_IG/web/model"
	"github.com/ioannisGiak89/foodHygiene_IG/web/model/fsaResponse"
)

type Calculator struct{}

// NewCalculator creates a Calculator
func NewCalculator() *Calculator {
	return &Calculator{}
}

// CalculateFhisPercentages calculates the percentages for FHIS scheme type
func (calculator *Calculator) CalculateFhisPercentages(establishments []fsaResponse.Establishment) *model.FhisPercentages {

	ratingCounters := [5]float64{0, 0, 0, 0, 0}
	for _, establishment := range establishments {

		switch strings.TrimSpace(establishment.RatingValue) {
		case "Pass":
			ratingCounters[0] += 1
			break
		case "Pass and Eat Safe":
			ratingCounters[1] += 1
			break
		case "Improvement Required":
			ratingCounters[2] += 1
			break
		case "Awaiting Inspection":
			ratingCounters[3] += 1
			break
		case "Exempt":
			ratingCounters[4] += 1
			break
		default:
			// TODO: Log errors somewhere useful
			fmt.Println("Fhis: Unknown rating value " + establishment.RatingValue)
			continue
		}
	}

	total := len(establishments)
	return &model.FhisPercentages{
		Pass:                calculator.calculatePercentage(ratingCounters[0], total),
		PassAndEatSafe:      calculator.calculatePercentage(ratingCounters[1], total),
		ImprovementRequired: calculator.calculatePercentage(ratingCounters[2], total),
		AwaitingInspection:  calculator.calculatePercentage(ratingCounters[3], total),
		Exempt:              calculator.calculatePercentage(ratingCounters[4], total),
	}
}

// CalculateFhrsPercentages calculates the percentages for FHRS scheme type
func (calculator *Calculator) CalculateFhrsPercentages(establishments []fsaResponse.Establishment) *model.FhrsPercentages {

	ratingCounters := [7]float64{0, 0, 0, 0, 0, 0, 0}
	for _, establishment := range establishments {

		if establishment.RatingValue == "Exempt" {
			ratingCounters[0] += 1
			continue
		} else if establishment.RatingValue == "AwaitingInspection" {
			ratingCounters[6] += 1
			continue
		}

		rating, err := strconv.Atoi(establishment.RatingValue)

		if err != nil {
			// TODO: Log errors somewhere useful
			fmt.Println("Fhrs Unknown rating value " + establishment.RatingValue)
			continue
		}

		ratingCounters[rating] += 1
	}

	total := len(establishments)
	return &model.FhrsPercentages{
		Exempt:             calculator.calculatePercentage(ratingCounters[0], total),
		OneStar:            calculator.calculatePercentage(ratingCounters[1], total),
		TwoStar:            calculator.calculatePercentage(ratingCounters[2], total),
		ThreeStar:          calculator.calculatePercentage(ratingCounters[3], total),
		FourStar:           calculator.calculatePercentage(ratingCounters[4], total),
		FiveStar:           calculator.calculatePercentage(ratingCounters[5], total),
		AwaitingInspection: calculator.calculatePercentage(ratingCounters[6], total),
	}
}

// calculatePercentage does the maths and converts float to string
func (calculator *Calculator) calculatePercentage(ratingCounter float64, totalEstablishments int) string {
	return fmt.Sprintf("%.2f", float32(100.0*ratingCounter/float64(totalEstablishments))) + "%"
}
