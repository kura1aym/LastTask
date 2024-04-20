package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"math"
	"sort"
	"time"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	requestID := uuid.New().String()
	userID := "user123"

	logger.WithFields(logrus.Fields{
		"timestamp":  time.Now().Format(time.RFC3339),
		"request_id": requestID,
		"user_id":    userID,
	}).Info("Starting the application")

	numbers := []float64{3.4, 2.1, 5.8, 3.2, 7.6, 9.0, 4.4}

	mean, err := calculateMeanWithLog(numbers, logger, requestID, userID)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Errorf("Error calculating mean: %v", err)
	} else {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Infof("Mean of numbers: %.2f", mean)
	}

	median, err := calculateMedianWithLog(numbers, logger, requestID, userID)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Errorf("Error calculating median: %v", err)
	} else {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Infof("Median of numbers: %.2f", median)
	}

	stdDev, err := calculateStdDevWithLog(numbers, mean, logger, requestID, userID)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Errorf("Error calculating standard deviation: %v", err)
	} else {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Infof("Standard deviation of numbers: %.2f", stdDev)
	}

	logger.WithFields(logrus.Fields{
		"timestamp":  time.Now().Format(time.RFC3339),
		"request_id": requestID,
		"user_id":    userID,
	}).Info("Ending the application")
}

func calculateMeanWithLog(numbers []float64, logger *logrus.Logger, requestID, userID string) (float64, error) {
	if len(numbers) == 0 {
		err := fmt.Errorf("input array is empty")
		logError(err, logger, requestID, userID, "Error calculating mean")
		return 0, err
	}
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers)), nil
}

func calculateMedianWithLog(numbers []float64, logger *logrus.Logger, requestID, userID string) (float64, error) {
	if len(numbers) == 0 {
		err := fmt.Errorf("input array is empty")
		logError(err, logger, requestID, userID, "Error calculating median")
		return 0, err
	}
	sort.Float64s(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2, nil
	}
	return numbers[mid], nil
}

func calculateStdDevWithLog(numbers []float64, mean float64, logger *logrus.Logger, requestID, userID string) (float64, error) {
	if len(numbers) == 0 {
		err := fmt.Errorf("input array is empty")
		logError(err, logger, requestID, userID, "Error calculating standard deviation")
		return 0, err
	}
	var variance float64
	for _, num := range numbers {
		variance += math.Pow(num-mean, 2)
	}
	variance /= float64(len(numbers))
	return math.Sqrt(variance), nil
}

func logError(err error, logger *logrus.Logger, requestID, userID, msg string) {
	if err != nil {
		logger.WithFields(logrus.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"request_id": requestID,
			"user_id":    userID,
		}).Errorf("%s: %v", msg, err)
	}
}
