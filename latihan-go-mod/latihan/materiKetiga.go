package latihan

import (
	"errors"
	"fmt"
)

// 1. Function sederhana dengan satu return value
func add(a, b int) int {
	return a + b
}

// 2. Function dengan multiple return values (value + error)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// 3. Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// 4. Higher-order function
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func MateriKetiga() {
	// Test function biasa
	fmt.Println("5 + 3 =", add(5, 3))

	// Test multiple return values - kasus sukses
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Test multiple return values - kasus gagal
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result:", result2)
	}

	// Test variadic function
	fmt.Println("Sum:", sumAll(1, 2, 3, 4, 5))

	// Test higher-order function dengan named function
	multiplyResult := applyOperation(4, 5, func(a, b int) int {
		return a * b
	})
	fmt.Println("4 * 5 =", multiplyResult)

	// Test Challenge 1
	// resultChallenge1, errChallenge1 := validateApplicationStatus("testing")
	// fmt.Println("Result Challenge:", resultChallenge1)
	// fmt.Println("Error Challenge:", errChallenge1)

	// Test Challenge 1
	resultChallenge2, errChallenge2 := getApplicationSummary(10, 3, 1)
	fmt.Println("Result Challenge:", resultChallenge2)
	fmt.Println("Error Challenge:", errChallenge2)

}

// Challenge 1
/**
 * Buat function validateApplicationStatus(status string)
 * (bool, error) — return true, nil kalau status termasuk
 * salah satu dari "applied", "interview", "offer", "rejected";
 * selain itu return false + error yang menjelaskan status tidak valid.
 *
 */

func validateApplicationStatus(status string) (bool, error) {
	switch status {
	case "applied", "interview", "offer", "rejected":
		return true, nil
	default:
		return false, errors.New("status tidak valid")
	}
}

// Challenge 2
/**
 * Buat function getApplicationSummary(total, interviews, offers int) (string, error)
 * Return error kalau interviews > total atau offers > interviews (data tidak masuk akal)
 * Kalau valid, return string summary seperti: "Total: 10, Interview rate: 30%, Offer rate: 10%"
 *
 */

func getApplicationSummary(total, interviews, offers int) (string, error) {
	if interviews > total || offers > interviews {
		return "", errors.New("Data tidak masuk akal")
	}

	interviewRate := (interviews * 100) / total
	offerRate := (offers * 100) / total

	summary := fmt.Sprintf(
		"Total: %d, Interview rate: %d%%, Offer rate: %d%%",
		total,
		interviewRate,
		offerRate,
	)

	return summary, nil
}
