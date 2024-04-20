package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"sort"
)

func main() {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{})
	logger.Info("Starting the application")

	numbers := []float64{3.4, 2.1, 5.8, 3.2, 7.6, 9.0, 4.4}

	mean, err := calculateMean(numbers, logger)
	if err != nil {
		logger.Errorf("Error calculating mean: %v", err)
	} else {
		logger.Infof("Mean of numbers: %.2f", mean)
	}

	median, err := calculateMedian(numbers, logger)
	if err != nil {
		logger.Errorf("Error calculating median: %v", err)
	} else {
		logger.Infof("Median of numbers: %.2f", median)
	}

	stdDev, err := calculateStdDev(numbers, mean, logger)
	if err != nil {
		logger.Errorf("Error calculating standard deviation: %v", err)
	} else {
		logger.Infof("Standard deviation of numbers: %.2f", stdDev)
	}

	logger.Info("Ending the application")
}

func calculateMean(numbers []float64, logger *logrus.Logger) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("input array is empty")
	}
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers)), nil
}

func calculateMedian(numbers []float64, logger *logrus.Logger) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("input array is empty")
	}
	sort.Float64s(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2, nil
	}
	return numbers[mid], nil
}

func calculateStdDev(numbers []float64, mean float64, logger *logrus.Logger) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("input array is empty")
	}
	variance := 0.0
	for _, num := range numbers {
		variance += math.Pow(num-mean, 2)
	}
	variance /= float64(len(numbers))
	return math.Sqrt(variance), nil
}
