// Struct

package latihan

import (
	"fmt"
)

// type User struct {
// 	Name     string
// 	Age      int
// 	IsActive bool
// }

func MateriKelima() {

	// Instance Struct
	// user := User{
	// 	Name:     "Aldy Azarya",
	// 	Age:      28,
	// 	IsActive: true,
	// }

	// user := User{"Aldy Azarya", 28 , true}

	// user :=  User {
	// 	Name : "Aldy Azarya",
	// 	// Age dan IsActive otomatis 0 dan false
	// }

	// fmt.Println(user.Name)

	// Struct vs map
	// data := map[string]interface{}{
	// 	"name": "Aldy",
	// 	"age":  28,
	// }

	// Nested Struct - struct dalam struct

	// type Address struct {
	// 	City    string
	// 	Country string
	// }

	// type User struct {
	// 	Name    string
	// 	Age     int
	// 	Address Address
	// }

	// user := User{
	// 	Name: "Aldy Azarya",
	// 	Age:  28,
	// 	Address: Address{
	// 		City:    "Jakarta",
	// 		Country: "Indonesia",
	// 	},
	// }

	// fmt.Println(user.Address.City)

	// Embedded Struct
	// type BaseModel struct {
	// 	ID        int
	// 	CreatedAt string
	// }

	// type User struct {
	// 	BaseModel
	// 	Name string
	// 	Age  int
	// }

	// user := User{
	// 	BaseModel: BaseModel{ID: 1, CreatedAt: "2026-01-01"},
	// 	Name:      "Aldy Zarya",
	// 	Age:       23,
	// }

	// fmt.Println(user.ID)
	// fmt.Println(user.CreatedAt)

	// Struct Tags - Penting untuk JSON
	// type User struct {
	// 	Name  string `json:""name`
	// 	Age   int    `json:""age`
	// 	Email string `json:"email,omitempty`
	// }

	// user := User{Name: "Aldy", Age: 28}

	// jsonData, _ := json.Marshal(user)
	// fmt.Println(string(jsonData))

	// Slice of Struct - Pola yang sering dipakai
	// type Application struct {
	// 	Company string
	// 	Status  string
	// }

	// applications := []Application{
	// 	{Company: "Deloitte", Status: "interview"},
	// 	{Company: "Covena", Status: "applied"},
	// 	{Company: "Startup X", Status: "rejected"},
	// }

	// for _, app := range applications {
	// 	fmt.Println(app.Company, "-", app.Status)
	// }

	// Anonymous Struct - Struct Tanpa Nama Tipe

	// point := struct {
	// 	X int
	// 	Y int
	// }{
	// 	X: 20,
	// 	Y: 21,
	// }

	// fmt.Println(point.X, point.Y)

	// Comparing Structs

	// type Point struct {
	// 	X, Y int
	// }

	// p1 := Point{X: 1, Y: 2}
	// p2 := Point{X: 1, Y: 2}

	// fmt.Println(p1 == p2)

	// latihanMandiriStruct()

	challengeStruct()

}

type Company struct {
	Name     string
	Industry string
}

// type Application struct {
// 	Company  Company // nested struct
// 	Position string
// 	Status   string
// }

func latihanMandiriStruct() {
	// 2. Buat beberapa instance
	// applications := []Application{
	// 	{
	// 		Company:  Company{Name: "Deloitte Digital", Industry: "Consulting"},
	// 		Position: "T&T Consultant - Sitecore Developer",
	// 		Status:   "applied",
	// 	},
	// 	{
	// 		Company:  Company{Name: "Covena.ai", Industry: "Technology"},
	// 		Position: "Frontend Engineer",
	// 		Status:   "interview",
	// 	},
	// }

	// 3. Iterasi dan akses nested field
	// for _, app := range applications {
	// 	fmt.Printf("%s at %s (%s) - Status: %s\n",
	// 		app.Position, app.Company.Name, app.Company.Industry, app.Status)
	// }

	// 4. Update field dari salah satu struct dalam slice
	// applications[1].Status = "offer"
	// fmt.Println("Updated status:", applications[1].Status)

	// 5. Comparing struct
	c1 := Company{Name: "Covena.ai", Industry: "Technology"}
	c2 := Company{Name: "Covena.ai", Industry: "Technology"}
	fmt.Println("c1 == c2:", c1 == c2) // true, dibandingkan by value
}

type Contact struct {
	Email string
	Phone string
}

// type Application struct {
// 	Company string
// 	Status  string
// }

func countByStatus(apps []Application) map[string]int {

	result := make(map[string]int)
	for _, app := range apps {
		result[app.Status]++
	}

	return result

}

func challengeStruct() {

	/*
		Tambahkan struct Contact dengan field Email dan Phone,
		lalu embed struct ini ke dalam Company.
		Akses Email langsung dari instance Company
		tanpa menulis company.Contact.Email.
	*/
	company := Contact{
		Email: "johndoe@gmail.com",
		Phone: "08777878787878",
	}
	fmt.Println(company.Email)

	/*
		Buat function func countByStatus(apps []Application) map[string]int
		yang menerima slice of Application dan mengembalikan
		jumlah aplikasi per status
		(gunakan for...range dan map yang sudah dipelajari sebelumnya).
	*/

	// applications := []Application{
	// 	{Company: "Deloitte", Status: "interview"},
	// 	{Company: "Covena", Status: "applied"},
	// 	{Company: "Startup X", Status: "rejected"},
	// }

	// result := countByStatus(applications)

	// fmt.Println(result)

}
