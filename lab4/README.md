# hust-cse-db-lab

a database application run on cmd line using go running on ubuntu

## design

build on go using gorm

## data description

```sql
-- 创建学生表 Student
CREATE TABLE Student (
    Sno CHAR(9) PRIMARY KEY,
    Sname CHAR(20) UNIQUE,
    Ssex CHAR(2),
    Sage SMALLINT,
    Sdept CHAR(20),
    Scholarship CHAR(2)
);

-- 创建课程表 Course
CREATE TABLE Course (
    Cno CHAR(4) PRIMARY KEY,
    Cname CHAR(40),
    Cpno CHAR(4) NULL,
    Ccredit SMALLINT NULL,
    FOREIGN KEY (Cpno) REFERENCES Course(Cno)
);

-- 创建选修表 SC
CREATE TABLE SC (
    Sno CHAR(9),
    Cno CHAR(4),
    Grade SMALLINT,
    PRIMARY KEY (Sno, Cno),
    FOREIGN KEY (Sno) REFERENCES Student(Sno),
    FOREIGN KEY (Cno) REFERENCES Course(Cno)
);
```

## functions

### init

```bash
./hust-cse-db-lab init
```
init database

###  drop

```bash
./hust-cse-db-lab drop
```
drop tables

### addStudent

```bash
./hust-cse-db-lab addStudent -sno 123456789 -sname "张三" -ssex "男" -sage 20 -sdept "计算机" -scholarship "是"
```
add a student. sno and sname are required, others are optional.

### addCourse

```bash
./hust-cse-db-lab addCourse -cno 1234 -cname "数据库" -cpno 1234 -ccredit 4
```
add a course. cno and cname are required, cpno and ccredit are optional.

### alterCourse

```bash
./hust-cse-db-lab alterCourse -cno 1234 -cname "数据库" -cpno 1234 -ccredit 4
```
alter course info. cno is required, others are optional.

### setGrade

```bash
./hust-cse-db-lab setGrade -sno 123456789 -cno 1234 -grade 100
```
set grade. sno, cno and grade are required.

### alterStudent

```bash
./hust-cse-db-lab alterStudent -sno 123456789 -sname "张三" -ssex "男" -sage 20 -sdept "计算机" -scholarship "是"
```
alter student info. sno is required, others are optional.


### queryDepartment

```bash
./hust-cse-db-lab queryDepartment -sdept "计算机"
```
statistics of students in a department, including grade, pass rate, excellent rate, fail count

define excellent as having grade >= 85, pass as having grade >= 60.

### allStat

```bash
./hust-cse-db-lab allStat
```
list all departments' statistics, including grade, pass rate, excellent rate, fail count.

### queryCourse

```bash
./hust-cse-db-lab queryCourse -cno 1234
```
ranked grade info of a course

### queryStudent

```bash
./hust-cse-db-lab queryStudent -sno 123456789
```
basic info of a student

### rankAll

```bash
./hust-cse-db-lab rankAll
```
rank all students by grade, group by department, and list students' name, id, average grade, cources taken as well as grade of each course.

##example

```bash
./hust-cse-db-lab addStudent -sno 200215121 -sname "李勇" -ssex "男" -sage 20 -sdept "CS" -scholarship "否"
./hust-cse-db-lab addStudent -sno 200215122 -sname "刘晨" -ssex "女" -sage 19 -sdept "CS" -scholarship "否"
./hust-cse-db-lab addStudent -sno 200215123 -sname "王敏" -ssex "女" -sage 18 -sdept "MA" -scholarship "否"
./hust-cse-db-lab addStudent -sno 200215125 -sname "张立" -ssex "男" -sage 19 -sdept "IS" -scholarship "否"

./hust-cse-db-lab addCourse -cno 1 -cname "数据库" -cpno 1 -ccredit 4
./hust-cse-db-lab addCourse -cno 2 -cname "数学" -cpno 2 -ccredit 2
./hust-cse-db-lab addCourse -cno 3 -cname "信息系统" -cpno 3 -ccredit 4
./hust-cse-db-lab addCourse -cno 4 -cname "操作系统" -cpno 4 -ccredit 3
./hust-cse-db-lab addCourse -cno 5 -cname "数据结构" -cpno 5 -ccredit 4
./hust-cse-db-lab addCourse -cno 6 -cname "数据处理" -cpno 6 -ccredit 2
./hust-cse-db-lab addCourse -cno 7 -cname "java" -cpno 7 -ccredit 4

./hust-cse-db-lab alterCourse -cno 1 - cpno 5
./hust-cse-db-lab alterCourse -cno 3 - cpno 1
./hust-cse-db-lab alterCourse -cno 4 - cpno 6
./hust-cse-db-lab alterCourse -cno 5 - cpno 7
./hust-cse-db-lab alterCourse -cno 7 - cpno 6

./hust-cse-db-lab setGrade -sno 200215121 -cno 1 -grade 92
./hust-cse-db-lab setGrade -sno 200215121 -cno 2 -grade 85
./hust-cse-db-lab setGrade -sno 200215121 -cno 3 -grade 20
./hust-cse-db-lab setGrade -sno 200215121 -cno 4 -grade 10
./hust-cse-db-lab setGrade -sno 200215122 -cno 2 -grade 100
./hust-cse-db-lab setGrade -sno 200215122 -cno 3 -grade 100
./hust-cse-db-lab setGrade -sno 200215123 -cno 4 -grade 82
./hust-cse-db-lab setGrade -sno 200215123 -cno 5 -grade 90
./hust-cse-db-lab setGrade -sno 200215125 -cno 6 -grade 85
./hust-cse-db-lab setGrade -sno 200215125 -cno 7 -grade 90

./hust-cse-db-lab alterStudent -sno 200215121 -sname "李勇" -ssex "男" -sage 20 -sdept "CS" -scholarship "否"

./hust-cse-db-lab queryDepartment -sdept "CS"

./hust-cse-db-lab queryCourse -cno 3


./hust-cse-db-lab queryStudent -sno 200215121

./hust-cse-db-lab rankAll

./hust-cse-db-lab allStat
```