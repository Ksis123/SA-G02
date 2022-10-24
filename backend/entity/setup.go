package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("SA-65-G02.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(
		//--------------------SlipList------------------------------------
		&Admin{},
		&Banking{},
		&PaymentStatus{},
		&SlipList{},
		//-----------------StudentList------------------------------------
		&Report{},
		&Status{},
		&StudentList{},
		//-----------------Scholarship------------------------------------
		&Scholarships_status{},
		&Scholarships_type{},
		&Scholarship{},
		//-------------------Donator--------------------------------------
		&Organization{},
		&TypeFund{},
		&Donator{},
	)
	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("555"), 14)

	db.Model(&Admin{}).Create(&Admin{
		Name:     "khem",
		Email:    "khemkhemsiwa555@gmail.com",
		Password: string(password),
	})

	var khem Admin

	db.Raw("SELECT * FROM admins WHERE email = ?", "khemkhemsiwa555@gmail.com").Scan(&khem)

	// -------------------------------------SlipList---------------------------------------
	// --- PaymentStatus Data
	Status1 := PaymentStatus{
		Name: "successful",
	}
	db.Model(&PaymentStatus{}).Create(&Status1)

	Status2 := PaymentStatus{
		Name: "processing",
	}
	db.Model(&PaymentStatus{}).Create(&Status2)

	Status3 := PaymentStatus{
		Name: "Delay",
	}
	db.Model(&PaymentStatus{}).Create(&Status3)

	// --- Banking Data
	bank1 := Banking{
		Name:     "Ace",
		Commerce: "SCB",
		Branch:   "Korat",
	}
	db.Model(&Banking{}).Create(&bank1)

	bank2 := Banking{
		Name:     "Luffy",
		Commerce: "KBANK",
		Branch:   "Korat",
	}
	db.Model(&Banking{}).Create(&bank2)

	// --- StudentList Data
	// s1 := StudentList{
	// 	Report: "201",
	// 	Status: "Pass",
	// }
	// db.Model(&StudentList{}).Create(&s1)

	// s2 := StudentList{
	// 	Report: "202",
	// 	Status: "FAIL",
	// }
	// db.Model(&StudentList{}).Create(&s2)

	// -------------------------------------Scholarship----------------------------------------------
	// Status Data
	Pass := Status{
		Status: "Pass",
	}
	db.Model(&Status{}).Create(&Pass)

	Fail := Status{
		Status: "Fail",
	}
	db.Model(&Status{}).Create(&Fail)

	// Report Data
	R1 := Report{
		Name: "สง่า ไปทุ่ง",
	}
	db.Model(&Report{}).Create(&R1)

	R2 := Report{
		Name: "สยาม ไปไร่",
	}
	db.Model(&Report{}).Create(&R2)

	// -------------------------------------StudentList----------------------------------------------
	// Scholarships_status Data
	status1 := Scholarships_status{
		Status: "ยังเปิดรับอยู่",
	}
	db.Model(&Scholarships_status{}).Create(&status1)

	// Scholarships_type
	type1 := Scholarships_type{
		Type: "ทุนให้เปล่า",
	}
	db.Model(&Scholarships_type{}).Create(&type1)

	// -------------------------------------Donator----------------------------------------------
	// Status Data
	var TypeFunds = []TypeFund{
		{TypeFund: "aum"},
		{TypeFund: "sum"},
		{TypeFund: "bim"},
	}
	db.CreateInBatches(TypeFunds, 3)

	var Organizations = []Organization{
		{Organization: "a"},
		{Organization: "b"},
		{Organization: "c"},
	}
	db.CreateInBatches(Organizations, 3)
	
}
