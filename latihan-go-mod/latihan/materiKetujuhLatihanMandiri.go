package latihan

import (
	"errors"
	"fmt"
)

// sentinel errors
var (
	ErrApplicationNotFound = errors.New("application not found")
	ErrInvalidStatus       = errors.New("invalid status")
	ErrUnauthorized        = errors.New("user not authorized")
)

type Applications struct {
	ID      string
	Company string
	Status  string
}

// Custom error type
type ValidationErrorMandiri struct {
	Field   string
	Message string
}

func (e *ValidationErrorMandiri) Error() string {
	return fmt.Sprintf("validation error on '%s': %s", e.Field, e.Message)
}

// Simulasi "database"
var applications = map[string]Applications{
	"1": {ID: "1", Company: "Deloitte Digital", Status: "interview"},
	"2": {ID: "2", Company: "Covena.ai", Status: "applied"},
}

func getApplication(id string) (*Applications, error) {
	app, exists := applications[id]
	if !exists {
		return nil, fmt.Errorf("getApplication: %w", ErrApplicationNotFound)
	}
	return &app, nil
}

func deleteApplication(id, userRole string) error {
	if userRole != "admin" {
		return fmt.Errorf("delete application: %w", ErrUnauthorized)
	}

	// cek apakah application ada
	_, exists := applications[id]
	if !exists {
		return fmt.Errorf("delete application: %w", ErrApplicationNotFound)
	}

	delete(applications, id)
	return nil
}

func updateStatus(id, newStatus string) error {
	if !validStatuses[newStatus] {
		return &ValidationErrorMandiri{Field: "status", Message: "must be one of: applied, interview, offer, rejected"}
	}

	_, err := getApplication(id)
	if err != nil {
		return fmt.Errorf("updateStatus: %w", err)
	}

	app := applications[id]
	app.Status = newStatus
	applications[id] = app
	return nil
}

var validStatuses = map[string]bool{
	"applied":   true,
	"interview": true,
	"offer":     true,
	"rejected":  true,
}

func bulkUpdateStatus(updates map[string]string) []error {
	var errs []error

	for id, newStatus := range updates {
		err := updateStatus(id, newStatus)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func MateriKetujuhLatihanMandiri() {

	// 1. Sukses
	err := updateStatus("1", "offer")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Status updated:", applications["1"].Status)
	}

	// 2. Application tidak ditemukan
	err = updateStatus("999", "offer")
	if errors.Is(err, ErrApplicationNotFound) {
		fmt.Println("Handled: application not found, return 404")
	} else if err != nil {
		fmt.Println("Other error:", err)
	}

	// 3. Status tidak valid - gunakan errors.As
	err = updateStatus("2", "ghosted") // status tidak ada dalam validStatuses
	var valErr *ValidationErrorMandiri
	if errors.As(err, &valErr) {
		fmt.Println("Handled: validation error on field", valErr.Field)
		fmt.Println("Message:", valErr.Message)
	}

	// Challenge 1
	/*
	 Tambahkan sentinel error ErrUnauthorized dan function
	 deleteApplication(id, userRole string) error — kalau
	 userRole != "admin", return ErrUnauthorized.
	 Test dengan errors.Is().
	*/
	errDelete := deleteApplication("1", "user")

	if errors.Is(errDelete, ErrUnauthorized) {
		fmt.Println("Handled: unauthorized, return 403")
	} else if errDelete != nil {
		fmt.Println("Other error:", errDelete)
	}

	// Challenge 2
	/*
		Buat function bulkUpdateStatus(updates map[string]string) []error —
		menerima map {id: newStatus}, melakukan updateStatus untuk
		masing-masing, dan mengumpulkan semua error
		(bukan langsung return di error pertama) ke dalam
		slice []error. Print semua error di akhir.
	*/
	updates := map[string]string{
		"1":   "offer",     // valid
		"2":   "ghosted",   // invalid status
		"999": "interview", // application tidak ditemukan
		"3":   "rejected",  // application tidak ditemukan
	}

	errs := bulkUpdateStatus(updates)

	if len(errs) == 0 {
		fmt.Println("Semua status berhasil di-update")
	} else {
		fmt.Println("Terjadi beberapa error:")
		for i, err := range errs {
			fmt.Printf("%d. %v\n", i+1, err)
		}
	}

}
