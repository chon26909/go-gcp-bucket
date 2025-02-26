package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Student ใช้เก็บข้อมูลนักเรียน
type Student struct {
	ID        string
	Title     string // คำนำหน้าชื่อ
	FirstName string
	LastName  string
	NickName  string
	Gender    string
	Class     string // เช่น "1/1", "6/6"
}

// รายชื่อสำหรับ Mock ข้อมูล
var maleFirstNames = []string{"ก้อง", "นัท", "ปอนด์", "ต้น", "เบียร์", "ภูมิ", "บาส", "โอ๊ต", "พีท", "แบงค์"}
var femaleFirstNames = []string{"ฟ้า", "มาย", "น้ำ", "เนย", "แนน", "แพรว", "มิ้น", "อิง", "พลอย", "ดาว"}
var lastNames = []string{"สมใจ", "ใจดี", "พรหมรักษา", "บุญมี", "สิงห์โต", "วงศ์คำ", "แสงจันทร์", "เกษมสุข", "ชื่นใจ", "รุ่งโรจน์"}
var nickNames = []string{"ตี๋", "จ๋า", "บีม", "มิว", "แป้ง", "พีช", "โอ๊ต", "คิม", "แบงค์", "อิ้ง"}

func generateStudents(startID int) []Student {
	var students []Student
	rand.Seed(time.Now().UnixNano())

	studentID := startID // เริ่มต้นจากเลขนักเรียนที่กำหนด

	for grade := 1; grade <= 6; grade++ { // ป.1 - ป.6
		for room := 1; room <= 6; room++ { // 6 ห้องต่อชั้น
			for i := 0; i < 30; i++ { // ห้องละ 30 คน
				isMale := rand.Intn(2) == 0 // สุ่มเพศ: 0 = ชาย, 1 = หญิง
				var firstName, title, gender string

				if isMale {
					firstName = maleFirstNames[rand.Intn(len(maleFirstNames))]
					title = "เด็กชาย"
					gender = "ชาย"
				} else {
					firstName = femaleFirstNames[rand.Intn(len(femaleFirstNames))]
					title = "เด็กหญิง"
					gender = "หญิง"
				}

				student := Student{
					ID:        strconv.Itoa(studentID),
					Title:     title,
					FirstName: firstName,
					LastName:  lastNames[rand.Intn(len(lastNames))],
					NickName:  nickNames[rand.Intn(len(nickNames))],
					Gender:    gender,
					Class:     fmt.Sprintf("%d/%d", grade, room),
				}
				students = append(students, student)
				studentID++
			}
		}
	}

	return students
}

func saveToCSV(students []Student, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// เขียน Header
	writer.Write([]string{"ID", "Title", "First Name", "Last Name", "Nick Name", "Gender", "Class"})

	// เขียนข้อมูลนักเรียน
	for _, student := range students {
		writer.Write([]string{student.ID, student.Title, student.FirstName, student.LastName, student.NickName, student.Gender, student.Class})
	}

	fmt.Println("✅ ไฟล์ถูกสร้าง:", filename)
	return nil
}

func main() {
	students := generateStudents(660001) // เริ่มต้นที่รหัส 660001
	err := saveToCSV(students, "students.csv")
	if err != nil {
		fmt.Println("❌ เกิดข้อผิดพลาด:", err)
	}
}
