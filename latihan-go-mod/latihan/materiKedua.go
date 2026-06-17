package latihan

import "fmt"

const AppName = "Job Application tracker"

func MateriKedua() {
	// 1. Deklarasikan variabel dengan berbagai cara
	var userName string = "Aldy Azarya"
	role := "Frontend Enginner"
	yearsOfExperience := 5
	var isOpenToWork bool = true
	var expectedSalary float64

	// 2. Print semua variabel
	fmt.Println("App:", AppName)
	fmt.Println("Name:", userName)
	fmt.Println("Role:", role)
	fmt.Println("Years of Experience:", yearsOfExperience)
	fmt.Println("Open to Work:", isOpenToWork)
	fmt.Println("Expected salary (zero value):", expectedSalary)

	// 3. Coba konversi type
	var experienceFloat float64 = float64(yearsOfExperience)
	fmt.Println("Experience as float:", experienceFloat)

	// 4. Konstanta group untuk status lamaran
	const (
		StatusApplied   = "applied"
		StatusInterview = "interview"
		StatusOffer     = "offer"
		StatusRejected  = "rejected"
	)
	currentStatus := StatusInterview
	fmt.Println("Application status:", currentStatus)

}
