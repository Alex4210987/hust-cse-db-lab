package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Student represents the Student model
type Student struct {
	Sno         string `gorm:"primaryKey"`
	Sname       string
	Ssex        string
	Sage        int
	Sdept       string
	Scholarship string
	Average     float64 `gorm:"-"` // Average grade
	Failed      bool    `gorm:"-"`
}

// Course represents the Course model
type Course struct {
	Cno     string `gorm:"primaryKey"`
	Cname   string
	Cpno    string
	Ccredit int
}

// SC represents the SC model
type SC struct {
	Sno   string `gorm:"primaryKey"`
	Cno   string `gorm:"primaryKey"`
	Grade int
}

func initDB() {
	// Create a new database
	db, err = gorm.Open(sqlite.Open("lab4.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	//db = db.Debug()
	// Auto migrate models
	if err := db.AutoMigrate(&Student{}, &Course{}, &SC{}); err != nil {
		fmt.Println("Failed to migrate database:", err)
		os.Exit(1)
	}
}

func dropDB() {
	// Drop database tables
	if err := db.Migrator().DropTable(&Student{}, &Course{}, &SC{}); err != nil {
		fmt.Println("Failed to drop database:", err)
		os.Exit(1)
	}
}

func printHelp() {
	helpMessage := `# hust-cse-db-lab
a database application run on cmd line using go running on ubuntu
## design
build on go using gorm
## data description
- 学生表 Student
- 课程表 Course
- 选修表 SC
## functions
### init
./hust-cse-db-lab init
init database
###  drop
./hust-cse-db-lab drop
drop tables
### addStudent
./hust-cse-db-lab addStudent -sno 123456789 -sname "张三" -ssex "男" -sage 20 -sdept "计算机" -scholarship "是"
add a student. sno and sname are required, others are optional.
### addCourse
./hust-cse-db-lab addCourse -cno 1234 -cname "数据库" -cpno 1234 -ccredit 4
add a course. cno and cname are required, cpno and ccredit are optional.
### alterCourse
./hust-cse-db-lab alterCourse -cno 1234 -cname "数据库" -cpno 1234 -ccredit 4
alter course info. cno is required, others are optional.
### setGrade
./hust-cse-db-lab setGrade -sno 123456789 -cno 1234 -grade 100
set grade. sno, cno and grade are required.
### alterStudent
./hust-cse-db-lab alterStudent -sno 123456789 -sname "张三" -ssex "男" -sage 20 -sdept "计算机" -scholarship "是"
alter student info. sno is required, others are optional.
### queryDepartment
statistics of students in a department, including grade, pass rate, excellent rate, fail count
define excellent as having grade >= 85, pass as having grade >= 60.
### allStat
list all departments' statistics, including grade, pass rate, excellent rate, fail count.
### queryCourse
ranked grade info of a course
### queryStudent
./hust-cse-db-lab queryStudent -sno 123456789
basic info of a student
### rankAll
./hust-cse-db-lab rankAll
rank all students by grade, group by department, and list students' name, id, average grade, cources taken as well as grade of each course.`
	fmt.Println(helpMessage)

}

func main() {
	initDB()

	// Command-line flags
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	dropCmd := flag.NewFlagSet("drop", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("-h", flag.ExitOnError)
	addStudentCmd := flag.NewFlagSet("addStudent", flag.ExitOnError)
	addCourseCmd := flag.NewFlagSet("addCourse", flag.ExitOnError)
	alterCourseCmd := flag.NewFlagSet("alterCourse", flag.ExitOnError)
	setGradeCmd := flag.NewFlagSet("setGrade", flag.ExitOnError)
	alterStudentCmd := flag.NewFlagSet("alterStudent", flag.ExitOnError)
	queryDepartmentCmd := flag.NewFlagSet("queryDepartment", flag.ExitOnError)
	queryCourseCmd := flag.NewFlagSet("queryCourse", flag.ExitOnError)
	queryStudentCmd := flag.NewFlagSet("queryStudent", flag.ExitOnError)
	allStatCmd := flag.NewFlagSet("allStat", flag.ExitOnError)
	rankAllCmd := flag.NewFlagSet("rankAll", flag.ExitOnError)

	// AddStudent command flags
	snoAddStudent := addStudentCmd.String("sno", "", "Student ID")
	snameAddStudent := addStudentCmd.String("sname", "", "Student Name")
	ssexAddStudent := addStudentCmd.String("ssex", "", "Student Gender")
	sageAddStudent := addStudentCmd.Int("sage", 0, "Student Age")
	sdeptAddStudent := addStudentCmd.String("sdept", "", "Student Department")
	scholarshipAddStudent := addStudentCmd.String("scholarship", "", "Student Scholarship")

	// AddCourse command flags
	cnoAddCourse := addCourseCmd.String("cno", "", "Course ID")
	cnameAddCourse := addCourseCmd.String("cname", "", "Course Name")
	cpnoAddCourse := addCourseCmd.String("cpno", "", "Prerequisite Course ID")
	ccreditAddCourse := addCourseCmd.Int("ccredit", 0, "Course Credit")

	// AlterCourse command flags
	cnoAlterCourse := alterCourseCmd.String("cno", "", "Course ID")
	cnameAlterCourse := alterCourseCmd.String("cname", "", "Course Name")
	cpnoAlterCourse := alterCourseCmd.String("cpno", "", "Prerequisite Course ID")
	ccreditAlterCourse := alterCourseCmd.Int("ccredit", 0, "Course Credit")

	// SetGrade command flags
	snoSetGrade := setGradeCmd.String("sno", "", "Student ID")
	cnoSetGrade := setGradeCmd.String("cno", "", "Course ID")
	gradeSetGrade := setGradeCmd.Int("grade", 0, "Grade")

	// AlterStudent command flags
	snoAlterStudent := alterStudentCmd.String("sno", "", "Student ID")
	snameAlterStudent := alterStudentCmd.String("sname", "", "Student Name")
	ssexAlterStudent := alterStudentCmd.String("ssex", "", "Student Gender")
	sageAlterStudent := alterStudentCmd.Int("sage", 0, "Student Age")
	sdeptAlterStudent := alterStudentCmd.String("sdept", "", "Student Department")
	scholarshipAlterStudent := alterStudentCmd.String("scholarship", "", "Student Scholarship")

	// QueryDepartment command flags
	sdeptQueryDepartment := queryDepartmentCmd.String("sdept", "", "Department Name")

	// QueryCourse command flags
	cnoQueryCourse := queryCourseCmd.String("cno", "", "Course ID")

	// QueryStudent command flags
	snoQueryStudent := queryStudentCmd.String("sno", "", "Student ID")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h":
		helpCmd.Parse(os.Args[2:])
		printHelp()

	case "init":
		initCmd.Parse(os.Args[2:])
		initDB()

	case "drop":
		dropCmd.Parse(os.Args[2:])
		dropDB()

	case "addStudent":
		addStudentCmd.Parse(os.Args[2:])
		addStudent(*snoAddStudent, *snameAddStudent, *ssexAddStudent, *sageAddStudent, *sdeptAddStudent, *scholarshipAddStudent)

	case "addCourse":
		addCourseCmd.Parse(os.Args[2:])
		addCourse(*cnoAddCourse, *cnameAddCourse, *cpnoAddCourse, *ccreditAddCourse)

	case "alterCourse":
		alterCourseCmd.Parse(os.Args[2:])
		alterCourse(*cnoAlterCourse, *cnameAlterCourse, *cpnoAlterCourse, *ccreditAlterCourse)

	case "setGrade":
		setGradeCmd.Parse(os.Args[2:])
		setGrade(*snoSetGrade, *cnoSetGrade, *gradeSetGrade)

	case "alterStudent":
		alterStudentCmd.Parse(os.Args[2:])
		alterStudent(*snoAlterStudent, *snameAlterStudent, *ssexAlterStudent, *sageAlterStudent, *sdeptAlterStudent, *scholarshipAlterStudent)

	case "queryDepartment":
		queryDepartmentCmd.Parse(os.Args[2:])
		queryDepartment(*sdeptQueryDepartment)

	case "queryCourse":
		queryCourseCmd.Parse(os.Args[2:])
		queryCourse(*cnoQueryCourse)

	case "queryStudent":
		queryStudentCmd.Parse(os.Args[2:])
		queryStudent(*snoQueryStudent)

	case "allStat":
		allStatCmd.Parse(os.Args[2:])
		allStat()

	case "rankAll":
		rankAllCmd.Parse(os.Args[2:])
		rankAll()

	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

// Functions to handle commands

func addStudent(sno, sname, ssex string, sage int, sdept, scholarship string) {
	student := Student{
		Sno:         sno,
		Sname:       sname,
		Ssex:        ssex,
		Sage:        sage,
		Sdept:       sdept,
		Scholarship: scholarship,
	}

	db.Create(&student)
	fmt.Println("Student added successfully!")
}

func addCourse(cno, cname, cpno string, ccredit int) {
	course := Course{
		Cno:     cno,
		Cname:   cname,
		Cpno:    cpno,
		Ccredit: ccredit,
	}

	db.Create(&course)
	fmt.Println("Course added successfully!")
}

func alterCourse(cno, cname, cpno string, ccredit int) {
	var course Course
	db.Where("Cno = ?", cno).First(&course)

	if cname != "" {
		course.Cname = cname
	}
	if cpno != "" {
		course.Cpno = cpno
	}
	if ccredit != 0 {
		course.Ccredit = ccredit
	}

	db.Save(&course)
	fmt.Println("Course altered successfully!")
}

func setGrade(sno, cno string, grade int) {
	// Check if the student and course exist
	var student Student
	if err := db.Where("Sno = ?", sno).First(&student).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	var course Course
	if err := db.Where("Cno = ?", cno).First(&course).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new SC record with the provided grade
	newSC := SC{
		Sno:   sno,
		Cno:   cno,
		Grade: grade,
	}

	// Use FirstOrCreate to find the existing record or create a new one
	if err := db.Where(&newSC).Assign(SC{Grade: grade}).FirstOrCreate(&newSC).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Grade set successfully!")
}

func alterStudent(sno, sname, ssex string, sage int, sdept, scholarship string) {
	var student Student
	db.Where("Sno = ?", sno).First(&student)

	if sname != "" {
		student.Sname = sname
	}
	if ssex != "" {
		student.Ssex = ssex
	}
	if sage != 0 {
		student.Sage = sage
	}
	if sdept != "" {
		student.Sdept = sdept
	}
	if scholarship != "" {
		student.Scholarship = scholarship
	}

	db.Save(&student)
	fmt.Println("Student altered successfully!")
}

// queryDepartment function
func queryDepartment(sdept string) {
	var students []Student
	db.Where("Sdept = ?", sdept).Find(&students)

	if len(students) == 0 {
		fmt.Println("No students found in the department:", sdept)
		return
	}

	var totalGrade, maxGrade, minGrade, failCount, excellentCount int
	totalStudents := len(students)
	courcesTaken := 0
	failCources := 0

	for _, student := range students {
		var sc []SC
		db.Where("Sno = ?", student.Sno).Find(&sc)

		for _, record := range sc {
			totalGrade += record.Grade
			courcesTaken++

			if record.Grade > maxGrade {
				maxGrade = record.Grade
			}

			if record.Grade < minGrade || minGrade == 0 {
				minGrade = record.Grade
			}

			if record.Grade >= 85 {
				excellentCount++
			}

			if record.Grade < 60 {
				if !student.Failed {
					failCount++
					student.Failed = true
				}
				failCources++
			}
		}
	}

	averageGrade := float64(totalGrade) / float64(courcesTaken)
	passRate := float64(courcesTaken-failCources) / float64(courcesTaken) * 100
	fmt.Println("Department:", sdept)
	fmt.Println("Student Count:", totalStudents)
	fmt.Println("Course Taken:", courcesTaken)
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
	fmt.Printf("Max Grade: %d\n", maxGrade)
	fmt.Printf("Min Grade: %d\n", minGrade)
	fmt.Printf("Pass Rate: %.2f%%\n", passRate)
	fmt.Printf("Excellent Rate: %.2f%%\n", float64(excellentCount)/float64(courcesTaken)*100)
	fmt.Printf("Fail Count: %d\n", failCount)
}

// queryCourse function
func queryCourse(cno string) {
	var course Course
	if err := db.Where("Cno = ?", cno).First(&course).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	if course.Cno == "" {
		fmt.Println("Course not found with ID:", cno)
		return
	}

	fmt.Println("Course ID:", course.Cno)
	fmt.Println("Course Name:", course.Cname)
	fmt.Println("Prerequisite Course ID:", course.Cpno)
	fmt.Println("Course Credit:", course.Ccredit)

	var sc []SC
	db.Where("Cno = ?", cno).Find(&sc)

	if len(sc) == 0 {
		fmt.Println("No students enrolled in this course.")
		return
	}

	fmt.Println("Enrolled Students:")
	fmt.Println("--------------")
	for _, record := range sc {
		var student Student
		db.Where("Sno = ?", record.Sno).First(&student)

		fmt.Printf("Student ID: %s\n", student.Sno)
		fmt.Printf("Student Name: %s\n", student.Sname)
		fmt.Printf("Grade: %d\n", record.Grade)
		fmt.Println("--------------")
	}
}

func queryStudent(sno string) {
	//输入学号，显示该学生的基本信息和选课信息。
	var student Student
	db.Where("Sno = ?", sno).First(&student)
	fmt.Println("Student ID:", student.Sno)
	fmt.Println("Student Name:", student.Sname)

	var sc []SC
	db.Where("Sno = ?", sno).Find(&sc)
	for _, sc := range sc {
		var course Course
		db.Where("Cno = ?", sc.Cno).First(&course)
		fmt.Println("Course ID:", course.Cno)
		fmt.Println("Course Name:", course.Cname)
		fmt.Println("Grade:", sc.Grade)
	}
}

// allStat function
// list all departments' statistics, including grade, pass rate, excellent rate, fail count.
// define excellent as having grade >= 85, pass as having grade >= 60.
func allStat() {
	var departments []string
	if err := db.Table("students").Select("DISTINCT sdept").Pluck("sdept", &departments).Error; err != nil {
		fmt.Println("Error fetching departments:", err)
		return
	}
	for _, dept := range departments {
		queryDepartment(dept)
		fmt.Println("--------------")
	}
}

// rankAll function
// rank all students by grade, group by department, and list students' name, id, average grade, cources taken as well as grade of each course.
// example: cs department 1st: id, name, average grade, courses taken, grade of each course; 2nd: ...; ...; ee department 1st: ...; ...
// average grade should be weighted by course credit.
func rankAll() {
	var departments []string
	if err := db.Table("students").Select("DISTINCT sdept").Pluck("sdept", &departments).Error; err != nil {
		fmt.Println("Error fetching departments:", err)
		return
	}

	for _, dept := range departments {
		fmt.Printf("%s Department\n", dept)
		fmt.Println("----------------------------")

		var students []Student
		db.Where("Sdept = ?", dept).Find(&students)

		// Calculate average grade for each student
		for i := range students {
			var sc []SC
			db.Where("Sno = ?", students[i].Sno).Find(&sc)

			var totalGrade, totalCredit int
			for _, record := range sc {
				var course Course
				db.Where("Cno = ?", record.Cno).First(&course)
				totalGrade += record.Grade * course.Ccredit
				totalCredit += course.Ccredit
			}

			if totalCredit > 0 {
				students[i].Average = float64(totalGrade) / float64(totalCredit)
			}
		}

		// Sort students by average grade in descending order
		sort.Slice(students, func(i, j int) bool {
			return students[i].Average > students[j].Average
		})

		rank := 0
		for _, student := range students {
			rank++
			fmt.Printf("Rank: %d\n", rank)
			fmt.Printf("Student ID: %s\n", student.Sno)
			fmt.Printf("Student Name: %s\n", student.Sname)
			fmt.Printf("Average Grade: %.2f\n", student.Average)

			var sc []SC
			db.Where("Sno = ?", student.Sno).Find(&sc)

			var totalCredit int
			for _, record := range sc {
				var course Course
				db.Where("Cno = ?", record.Cno).First(&course)
				fmt.Printf("Course ID: %s\n", course.Cno)
				fmt.Printf("Course Name: %s\n", course.Cname)
				fmt.Printf("Grade: %d\n", record.Grade)
				totalCredit += course.Ccredit
			}
			fmt.Printf("Courses Taken: %d\n", len(sc))
			fmt.Printf("Credit: %d\n", totalCredit)
			fmt.Println("--------------")
		}
	}
}
