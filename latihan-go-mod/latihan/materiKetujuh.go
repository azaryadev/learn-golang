// Pola Error Handling

package latihan

import (
	"errors"
	"fmt"
)

func divideHandling(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Pola Dasar: Check, Handle, Return
// func processApplication(id string) error {
// 	app, err := fetchApplication(id)
// 	if err != nil {
// 		return err
// 	}

// 	err = validateApplication(app)
// 	if err != nil {
// 		return err
// 	}

// 	err = saveApplication(app)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Menambahkan Context ke Error — fmt.Errorf dengan %w
// func fetchApplication(id string) (*Application, error) {
// 	app, err := db.Query(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("fetchApplication: failed to query database: %w", err)
// 	}

// 	return app, nil
// }

// errors.Is() — Mengecek apakah error tertentu ada di "chain"
// var ErrNotFound = errors.New("not found")

func fetchUser(id string) (*User, error) {
	if id == "" {
		return nil, ErrNotFound
	}

	return nil, nil

}

// errors.As() — Mengecek dan "Cast" ke Tipe Error Tertentu
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// Custom Error Types
func validateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Msg: "email is required"}
	}
	return nil
}

// recover — "Catch" untuk panic (Jarang Dipakai Langsung)
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = a / b
	return
}

/*
  defer + recover ini sering dipakai sebagai
  "safety net" di level tertentu
  (misal: middleware HTTP server, supaya satu request yang panic tidak menjatuhkan seluruh server).
  Untuk pemula, cukup tahu ini ada —
  jangan terbiasa pakai panic/recover sebagai
  pengganti error handling normal.
*/

// Pola Umum : Sentinel Errors
var (
	ErrNotFound = errors.New("resource not found")
	// ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

// func getApplication(id string) (*Application, error) {
// 	if id == "" {
// 		return nil, ErrInvalidInput
// 	}

// 	app, exists := database[id]
// 	if !exists {
// 		return nil, ErrNotFound
// 	}
// 	return app, nil
// }

func MateriKetujuh() {

	// result, err := divideHandling(10, 0)
	// if err != nil {
	// 	fmt.Println("Something went wrong:", err)
	// 	return
	// }
	// fmt.Println(result)

	// membuat error sederhana
	// err := errors.New("Something went wrong")
	// fmt.Println(err.Error())
	// fmt.Println(err)

	// Membuat error dengan formatted message
	// name := "aldy"
	// err := fmt.Errorf("user %s not found", name)
	// fmt.Println(err)

	// errors.Is()
	// user, err := fetchUser("")
	// if errors.Is(err, ErrNotFound) {
	// 	fmt.Println("User tidak ditemukan, tampilkan halaman 404")
	// } else if err != nil {
	// 	fmt.Println("Error lain:", err)
	// }
	// fmt.Println(user)

	// errors.As()
	// var valErr *ValidationError
	// if errors.As(err, &valErr) {
	// 	fmt.Println("Field bermasalah:", valErr.Field)
	// }

	// Pola Umum : Sentinel Errors
	// app, err := getApplication("123")
	// switch {
	// case errors.Is(err, ErrNotFound):
	// 	//return 404
	// case errors.Is(err, ErrInvalidInput):
	// 	// return 400
	// case err != nil:
	// 	// return 500
	// default:
	// 	// success
	// }
	/*
	 Ini akan sangat berguna saat membangun
	 HTTP API — sentinel errors membantu menentukan
	 HTTP status code yang tepat berdasarkan
	 jenis error.
	*/
}
