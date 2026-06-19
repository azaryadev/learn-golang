// Pointer

package latihan

import "fmt"

type User struct {
	Name string
	Age  int
}

func incrementAgeWithoutPointer(u User) {
	u.Age = u.Age + 1
}

func incrementAgeWithPointer(u *User) {
	u.Age = u.Age + 1
}

type UserRepresentation struct {
	Name  string
	Email *string //Email bisa nil
}

type Application struct {
	Company string
	Status  string
	Notes   *string
}

func (a Application) SetStatusValue(status string) {
	a.Status = status
}

func (a *Application) SetStatusPointer(status string) {
	a.Status = status
}

func appendItem(items []string) {
	items = append(items, "new item")
}

func modifyFirst(items []string) {
	items[0] = "CHANGED"
}

func MateriKeenam() {

	// tanpa pointer
	user := User{Name: "Aldy", Age: 28}
	incrementAgeWithoutPointer(user)
	fmt.Println(user.Age)

	// dengan pointer
	userWithPointer := User{Name: "Aldy", Age: 28}
	incrementAgeWithPointer(&userWithPointer)
	fmt.Println(userWithPointer.Age)

	// representasi "tidak ada value" (nullable)
	// var noEmail *string = nil
	// user1 := UserRepresentation{Name: "Aldy", Email: nil}

	// email := "aldy1612@gmail.com"
	// user2 := User{Name: "Budi", Email: &email}

	// pointer dan method - pointer receiver vs value receiver
	app := Application{Status: "applied"}
	app.SetStatusValue("interview")
	fmt.Println(app.Status) // "applied" - tidak berubah

	app.SetStatusPointer("interview")
	fmt.Println(app.Status) // "inteview" - berubah

	// Slice dan Map
	list := []string{"a", "b"}
	appendItem(list)
	fmt.Println(list) //["a", "b"] - Tidak bertambah (karena append bisa realokasi memory)

	modifyFirst(list)
	fmt.Println(list)

	latihanMandiriPointer()
}

// Function dengan pointer receiver untuk mengubah status
func (a *Application) UpdateStatus(newStatus string) {
	a.Status = newStatus
}

func addNote(app *Application, note string) {
	app.Notes = &note
}

func resetNotes(app *Application) {
	app.Notes = nil
}

func latihanMandiriPointer() {
	// 1. Buat instance tanpa notes (nil)
	app := Application{Company: "Deloitte Digital", Status: "applied"}
	fmt.Println("Initial status:", app.Status)
	fmt.Println("Notes:", app.Notes) // <nil>

	// 2. Update status menggunakan pointer receiver
	app.UpdateStatus("interview")
	fmt.Println("Updated status:", app.Status)

	// 3. Tambahkan notes menggunakan pointer
	addNote(&app, "Recruiter Jason replied positively")
	if app.Notes != nil {
		fmt.Println("Notes:", *app.Notes) // dereference untuk baca value
	}

	// challenge 1
	/*
		Buat function func resetNotes(app *Application)
		yang men-set app.Notes kembali ke nil.
		Panggil function ini setelah langkah 3 di atas,
		lalu cek apakah app.Notes == nil.
	*/
	resetNotes(&app)
	if app.Notes == nil {
		fmt.Println("Notes berhasil di-reset manjadi nil")
	}

	// 4. Bandingkan dengan value receiver (tidak mengubah apa-apa)
	func(a Application) {
		a.Status = "rejected"
	}(app)
	fmt.Println("Status setelah value receiver call:", app.Status) // tetap "interview"

	// challenge 2
	/*
	 Buat dua variabel app1 := Application{...} dan
	 app2 := app1 (kopi biasa, bukan pointer).
	 Ubah app2.Status, lalu print app1.Status
	 dan app2.Status — buktikan bahwa keduanya
	 independen (struct = by value).
	*/

	app1 := Application{Company: "Deloitte Digital", Status: "applied", Notes: nil}
	app2 := app1
	app2.Status = "testing 123"
	fmt.Println("status 1", app1.Status)
	fmt.Println("status 2", app2.Status)

	// challenge 3
	/*
		Sekarang buat app3 := &app1 (pointer),
		ubah app3.Status, lalu print app1.Status — buktikan
		bahwa keduanya terhubung
		(pointer = shared reference).
	*/
	app3 := &app1
	app3.Status = "Offering"
	fmt.Println("status 3", app1.Status)
}
