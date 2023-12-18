# hust-cse-db-lab

a database application run on cmd line

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
4）按系统计学生的平均成绩、最好成绩、最差成绩、优秀率、不及格人数。

define excellent as having grade >= 85, pass as having grade >= 60.

### queryCourse

```bash
./hust-cse-db-lab queryCourse -cno 1234
```
5）按系对学生成绩进行排名，同时显示出学生、课程和成绩信息

### queryStudent

```bash
./hust-cse-db-lab queryStudent -sno 123456789
```
6）输入学号，显示该学生的基本信息和选课信息。