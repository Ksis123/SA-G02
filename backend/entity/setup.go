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

		//ระบบลงทะเบียนข้อมูลนักศึกษา---------------------------------------------------------------
		&Year{},
		&Faculty{},
		&Advisor{},
		&User{},
		&Student{},

	)
	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("555"), 14)
	
	// -------------------------------------Admin---------------------------------------
	db.Model(&Admin{}).Create(&Admin{
		Name:     "khem",
		Email:    "khemkhemsiwa555@gmail.com",
		Password: string(password),
	})
	var khem Admin
	
	// -------------------------------------User---------------------------------------
	db.Model(&User{}).Create(&User{
		Name:     "nok",
		Email:    "nok@gmail.com",
		Password: string(password),
	})
	var nok User

	db.Raw("SELECT * FROM admins WHERE email = ?", "khemkhemsiwa555@gmail.com").Scan(&khem)
	db.Raw("SELECT * FROM users WHERE email = ?", "nok@gmail.com").Scan(&nok)

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
	
	// -------------------------------------Student----------------------------------------------
	var year = []Year{{Number: "1/2565"}, {Number: "2/2565"}, {Number: "3/2565"}}
	db.CreateInBatches(year, 3)

	var faculties = []Faculty{
		{Name: "Institute of Science (Chemistry)", ThaiName: "สำนักวิชาวิทยาศาสตร์ (สาขาวิชาเคมี)"},
		{Name: "Institute of Science (Mathematics)", ThaiName: "สำนักวิชาวิทยาศาสตร์ (สาขาวิชาคณิตศาสตร์)"},
		{Name: "Institute of Science (Biology)", ThaiName: "สำนักวิชาวิทยาศาสตร์ (สาขาวิชาชิววิทยา)"},
		{Name: "Institute of Science (Physics)", ThaiName: "สำนักวิชาวิทยาศาสตร์ (สาขาวิชาฟิสิกส์)"},
		{Name: "Institute of Science (Sports Science)", ThaiName: "สำนักวิชาวิทยาศาสตร์ (สาขาวิชาวิทยาศาสตร์การกีฬา)"},

		{Name: "Institute of Social Technology (Management Technology)", ThaiName: "สำนักวิชาเทคโนโลยีสังคม (สาขาเทคโนโลยีการจัดการ)"},
		{Name: "Institute of Social Technology (Management)", ThaiName: "สำนักวิชาเทคโนโลยีสังคม (สาขาการจัดการบัณฑิต)"},

		{Name: "Institute of Agricultural Technology (Crop Production Technology)", ThaiName: "สำนักวิชาเทคโนโลยีการเกษตร (สาขาเทคโนโลยีการผลิตพืช)"},
		{Name: "Institute of Agricultural Technology (Animal Production Technology)", ThaiName: "สำนักวิชาเทคโนโลยีการเกษตร (สาขาเทคโนโลยีการผลิตสัตว์)"},
		{Name: "Institute of Agricultural Technology (Animal Technology and Innovation)", ThaiName: "สำนักวิชาเทคโนโลยีการเกษตร (สาขาเทคโนโลยีและนวัตกรรมทางสัตว์)"},
		{Name: "Institute of Agricultural Technology (Food Technology)", ThaiName: "สำนักวิชาเทคโนโลยีการเกษตร (สาขาเทคโนโลยีอาหาร)"},

		{Name: "Institute of Medicine (Doctor Of Dental Surgery)", ThaiName: "สำนักวิชาแพทยศาสตร์ (สาขาแพทยศาสตร์)"},
		{Name: "Institute of Medicine (Medical Science)", ThaiName: "สำนักวิชาแพทยศาสตร์ (สาขาวิทยาศาสตร์การแพทย์)"},

		{Name: "Institute of Engineering (Manufacturing Engineering)", ThaiName: "สำนักวิชาวิศวกรรมศาสตร์ (สาขาวิศวกรรมการผลิต)"},
		{Name: "Institute of Engineering (Agricultural Engineering)", ThaiName: "สำนักวิชาวิศวกรรมศาสตร์ (สาขาวิศวกรรมเกษตร)"},
		{Name: "Institute of Engineering (Computer Engineering)", ThaiName: "สำนักวิชาวิศวกรรมศาสตร์ (สาขาวิศวกรรมคอมพิวเตอร์)"},
		{Name: "Institute of Engineering (Chemical Engineering)", ThaiName: "สำนักวิชาวิศวกรรมศาสตร์ (สาขาวิศวกรรมเคมี)"},
		{Name: "Institute of Engineering (Mechanical Engineering)", ThaiName: "สำนักวิชาวิศวกรรมศาสตร์ (สาขาวิศวกรรมเครื่องกล)"},
		{Name: "Institute of Engineering (Electrical Engineering)", ThaiName: "สำนักวิชาวิศวกรรมศาสตร์ (สาขาวิศวกรรมไฟฟ้า)"},

		{Name: "Institute of Nursing (Bachelor of Nursing Science)", ThaiName: "สำนักวิชาพยาบาลศาสตร์ (สาขาพยาบาลศาสตรบัณฑิต)"},

		{Name: "Institute of Dentistry (Doctor Of Dental Surgery)", ThaiName: "สำนักวิชาทันตแพทยศาสตร์ (สาขาทันตแพทยศาสตรบัณฑิต)"},
		{Name: "Institute of Dentistry (Medical Science)", ThaiName: "สำนักวิชาทันตแพทยศาสตร์ (สาขาวิทยาศาสตร์การแพทย์)"},

		{Name: "Institute of Public Health (Occupational Health and Safety)", ThaiName: "สำนักวิชาสาธารณสุขศาสตร์ (สาขาอาชีวอนามัยและความปลอดภัย)"},
		{Name: "Institute of Public Health (Environmental Health)", ThaiName: "สำนักวิชาสาธารณสุขศาสตร์ (สาขาอนามัยสิ่งแวดล้อม)"},
		{Name: "Institute of Public Health (Nutrition and Dietetics)", ThaiName: "สำนักวิชาสาธารณสุขศาสตร์ (สาขาโภชนวิทยาและการกำหนดอาหาร)"},

		{Name: "Institute of Digital Arts and Science (Information Technology)", ThaiName: "สำนักวิชาศาสตร์และศิลป์ดิจิทัล (สาขาเทคโนโลยีสารสนเทศ)"},
		{Name: "Institute of Digital Arts and Science (Information Science)", ThaiName: "สำนักวิชาศาสตร์และศิลป์ดิจิทัล (สาขาวิทยาการสารสนเทศ)"}}
	db.CreateInBatches(faculties, 9)

	var advisor = []Advisor{
		{Name: "Dr.Piriyakorn Khan-O", ThaiName: "ดร.พิริยกร ขันโอ"},
		{Name: "Dr.Somchai Jaidee", ThaiName: "ดร.สมชาย ใจดี"},
		{Name: "Mr.Sommai Supap", ThaiName: "นายสมหมาย สุภาพ"},
		{Name: "Mr.Pongsakorn Bunyiem", ThaiName: "นายพงศกร บุญเยี่ยม"},
		{Name: "Miss.Busaba  Chuatrakoon", ThaiName: "นางสาวบุษบา ฉั่วตระกูล"}}
	db.CreateInBatches(advisor, 5)

}
