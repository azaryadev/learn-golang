package latihan

import (
	"fmt"
)

func ifStatement() {
	// if statement
	age := 20


	if age >= 18 {
		fmt.Println("Dewasa")
	} else if age >= 13 {
		fmt.Println("Remaja")
	} else {
		fmt.Println("Anak-anak")
	}

	// if dengan initialization statement
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// error handling
	// if value, err := someFunction(); err != nil {
	// 	return err
	// } else {
	// 	fmt.Println(value)
	// }
}

func switchCase() {
	// SWITCH
	status := "interview"

	switch status {
	case "applied":
		fmt.Println("Lamaran terkirim")
	case "interview":
		fmt.Println("Sedang Interview")
	case "offer", "accepted":
		fmt.Println()
	default:
		fmt.Println("Status tidak dikenal")
	}
}

func forLoop() {
	// FOR - satu-satunya loop di GO

	// for - klasik
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for - sebagai pengganti while
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for - tanpa kondisi pengganti while (true)
	i := 0

	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func forRange() {

	// Iterasi Slice
	// fruits := []string{"apple", "banana", "mango"}

	// for index, fruit := range fruits {
	// 	fmt.Println(index, fruit)
	// }

	// kalau indeks tidak dibutuhkan - pakai _
	// for _, fruit := range fruits {
	// 	fmt.Println(fruit)
	// }

	// Iterasi Map (mirip object)
	// ages := map[string]int{
	// 	"Aldy": 28,
	// 	"Budi": 30,
	// }

	// for name, age := range ages {
	// 	fmt.Println(name, "is", age, "years old")
	// }

	// map(), Filter(), reduce()
	numbers := []int{1, 2, 3, 4, 5}

	// map
	doubled := []int{}
	for _, n := range numbers {
		doubled = append(doubled, n*2)
		fmt.Println("Map", doubled)
	}

	// filter
	evens := []int{}
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
			fmt.Println("Filter", evens)
		}
	}

	// reduce
	total := 0
	for _, n := range numbers {
		total += n

		fmt.Println("Reduce", total)
	}
}

func latihanMandiri() {
	// 1. if dengan initialization statement
	applications := map[string]string{
		"Deloitte":  "interview",
		"Covena":    "applied",
		"Startup X": "rejected",
	}

	// 2. Iterasi map dan gunakan switch untuk kategorisasi
	for company, status := range applications {
		switch status {
		case "applied":
			fmt.Println(company, "- Masih menunggu respon")
		case "interview":
			fmt.Println(company, "- Sedang proses interview!")
		case "offer":
			fmt.Println(company, "- Selamat, ada offer!")
		case "rejected":
			fmt.Println(company, "- Belum cocok, lanjut ke yang lain")
		default:
			fmt.Println(company, "- Status tidak dikenal")
		}
	}

	// 3. Loop manual untuk menghitung jumlah aplikasi per status
	statusCount := make(map[string]int)
	for _, status := range applications {
		statusCount[status]++
	}

	for status, count := range statusCount {
		fmt.Println(status, ":", count)
	}

	// 4. Pengganti "while" - retry logic sederhana
	attempts := 0
	maxAttempts := 3
	success := false

	for attempts < maxAttempts && !success {
		attempts++
		fmt.Println("Mencoba koneksi ke server... percobaan ke-", attempts)
		if attempts == 3 {
			success = true
			fmt.Println("Berhasil terhubung!")
		}
	}

}

func challenge() {
	scores := []int{85, 90, 65, 70, 95, 60}

	// Slice baru passingScores yang hanya berisi nilai >= 70
	passingScores := []int{}
	for _, score := range scores {
		if score >= 70 {
			passingScores = append(passingScores, score)
		}
	}

	// Hitung rata-rata semua nilai
	total := 0

	for _, score := range scores {
		total += score
	}

	average := float64(total) / float64(len(scores))

	// Hitung berapa banyak nilai yang termasuk "A" (>= 90)
	countA := 0

	for _, score := range scores {
		if score >= 90 {
			countA++
		}
	}

	fmt.Println("Passing Scores:", passingScores)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Jumlah nilai A: %d\n", countA)

	/*
		Buat program yang mensimulasikan pengecekan status lamaran
		kerja setiap "hari" (loop dari hari 1-7).
		Gunakan for dengan label dan break — kalau hari ke-5
		statusnya berubah menjadi "interview",
		langsung break keluar dari loop dan print "Update ditemukan di hari ke-5!"
	*/
	status := "pending"

jobCheck:
	for day := 1; day <= 7; day++ {
		fmt.Printf("Hari ke-%d: Status = %s\n", day, status)

		if day == 5 {
			status = "interview"
			fmt.Println("Update ditemukan di hari ke-5!")
			break jobCheck
		}
	}

	fmt.Println("Pengecekan selesai.")

}

func MateriKeempat() {
	// forRange()

	// latihanMandiri()
	challenge()
}
