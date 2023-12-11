-- 删除已存在的数据库
DROP DATABASE IF EXISTS S_T_U2022xxxxx;

-- 创建新的数据库
CREATE DATABASE S_T_U2022xxxxx;

-- 切换到新创建的数据库
USE S_T_U2022xxxxx;

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


-- 插入示例数据到表 Student
INSERT INTO Student (Sno, Sname, Ssex, Sage, Sdept, Scholarship) VALUES
('200215121', '李勇', '男', 20, 'CS', '否'),
('200215122', '刘晨', '女', 19, 'CS', '否'),
('200215123', '王敏', '女', 18, 'MA', '否'),
('200215125', '张立', '男', 19, 'IS', '否');

-- 插入示例数据到表 Course
INSERT INTO Course (Cno, Cname, Cpno, Ccredit) VALUES
('1', '数据库', NULL, 4),
('2', '数学', NULL, 2),
('3', '信息系统', NULL, 4),
('4', '操作系统', NULL, 3),
('5', '数据结构', NULL, 4),
('6', '数据处理', NULL, 2),
('7', 'java', NULL, 4);

-- 更新 Course 表的 Cpno 字段
UPDATE Course SET Cpno = '5' WHERE Cno = '1';
UPDATE Course SET Cpno = '1' WHERE Cno = '3';
UPDATE Course SET Cpno = '6' WHERE Cno = '4';
UPDATE Course SET Cpno = '7' WHERE Cno = '5';
UPDATE Course SET Cpno = '6' WHERE Cno = '7';

-- 插入示例数据到表 SC
INSERT INTO SC (Sno, Cno, Grade) VALUES
('200215121', '1', 92),
('200215121', '2', 85),
('200215121', '3', 88),
('200215122', '2', 90),
('200215122', '3', 80);

DROP USER IF EXISTS 'lab'@'localhost';

CREATE USER 'lab'@'localhost' IDENTIFIED BY '';

GRANT ALL PRIVILEGES ON S_T_U2022xxxxx.* TO 'lab'@'localhost';

GRANT CREATE USER ON *.* TO 'lab'@'localhost';

GRANT GRANT OPTION ON *.* TO 'lab'@'localhost';

GRANT RELOAD ON *.* TO 'lab'@'localhost';

FLUSH PRIVILEGES;