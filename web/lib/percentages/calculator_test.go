package percentages

import (
	"testing"

	"github.com/eidolon/console/assert"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should return percentages.Calculator struct", func(t *testing.T) {
		result := NewCalculator()
		assert.Equal(t, &Calculator{}, result)
	})
}

func TestCalculator_CalculateFhisPercentages(t *testing.T) {
	calculator := NewCalculator()

	t.Run("should match the values in correctFhisPercentages", func(t *testing.T) {
		result := calculator.CalculateFhisPercentages(testFhisEstablishmentRatingData)
		assert.Equal(t, &correctFhisPercentages, result)
	})

	t.Run("should return Exempt 0", func(t *testing.T) {
		result := calculator.CalculateFhisPercentages(testFhisEstablishmentRatingWithoutExempt)
		assert.Equal(t, &correctFhisPercentagesWithoutExempt, result)
	})

	t.Run("should calculate the rest no matter empty values", func(t *testing.T) {
		result := calculator.CalculateFhisPercentages(testFhisEstablishmentRatingWithEmptyValues)
		assert.Equal(t, &correctFhisPercentagesEmptyValues, result)
	})
}

func TestCalculator_CalculateFhrsPercentages(t *testing.T) {
	calculator := NewCalculator()

	t.Run("should match the values in correctFhrsPercentages", func(t *testing.T) {
		result := calculator.CalculateFhrsPercentages(testFhrsEstablishmentRatingData)
		assert.Equal(t, &correctFhrsPercentages, result)
	})

	t.Run("should return Exempt 0", func(t *testing.T) {
		result := calculator.CalculateFhrsPercentages(testFhrsEstablishmentRatingWithoutExempt)
		assert.Equal(t, &correctFhrsPercentagesWithoutExempt, result)
	})

	t.Run("should calculate the rest no matter empty values", func(t *testing.T) {
		result := calculator.CalculateFhrsPercentages(testFhrsEstablishmentRatingWithEmptyValues)
		assert.Equal(t, &correctFhrsPercentagesEmptyValues, result)
	})
}
