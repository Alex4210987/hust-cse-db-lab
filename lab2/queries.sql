-- the db is inited with the same svheema as in lab1

-- 切换到数据库
USE S_T_U2022xxxxx;

-- 查询每门课程及其被选情况
SELECT 
    C.Cno AS 课程号,
    C.Cname AS 课程名称,
    SC.Sno AS 学生学号,
    SC.Grade AS 成绩
FROM Course C
LEFT JOIN SC ON C.Cno = SC.Cno;

-- 查询与“张立”同岁的学生的学号、姓名和年龄
-- 方法1：
SELECT Sno, Sname, Sage
FROM Student
WHERE Sage = (SELECT Sage FROM Student WHERE Sname = '张立');

-- 方法2：
SELECT S2.Sno, S2.Sname, S2.Sage
FROM Student S1, Student S2
WHERE S1.Sname = '张立' AND S1.Sage = S2.Sage;

-- 方法3：
SELECT S.Sno, S.Sname, S.Sage
FROM Student S
JOIN (SELECT Sage FROM Student WHERE Sname = '张立') Z
ON S.Sage = Z.Sage;

-- 查询选修了 3 号课程而且成绩为良好（80~89 分）的所有学生的学号和姓名
SELECT SC.Sno, S.Sname
FROM SC
JOIN Student S ON SC.Sno = S.Sno
WHERE SC.Cno = '3' AND SC.Grade BETWEEN 80 AND 89;

-- 查询学生 200215122 选修的课程号、课程名
SELECT SC.Cno, C.Cname
FROM SC
JOIN Course C ON SC.Cno = C.Cno
WHERE SC.Sno = '200215122';

-- 找出每个学生低于他所选修课程平均成绩 5 分以上的课程号
SELECT SC.Sno, SC.Cno
FROM SC
WHERE SC.Grade < (
    SELECT AVG(Grade) - 5
    FROM SC
    WHERE SC.Sno = SC.Sno
);

-- 查询比所有男生年龄都小的女生的学号、姓名和年龄
SELECT S.Sno, S.Sname, S.Sage
FROM Student S
WHERE S.Ssex = '女' AND S.Sage < ALL (SELECT Sage FROM Student WHERE Ssex = '男');

-- 查询所有选修了 2 号课程的学生姓名及所在系
SELECT S.Sname, S.Sdept
FROM Student S
JOIN SC ON S.Sno = SC.Sno
WHERE SC.Cno = '2';

-- 创建临时表格存放需要更新的学生学号
CREATE TEMPORARY TABLE TempStudents AS
SELECT S.Sno
FROM Student S
JOIN SC ON S.Sno = SC.Sno
WHERE SC.Grade BETWEEN 80 AND 89;

-- 使用 update 语句把平均成绩为良的学生的年龄增加 2 岁，并查询出来
UPDATE Student
SET Sage = Sage + 2
WHERE Sno IN (SELECT Sno FROM TempStudents);

-- 查询更新后的结果
SELECT * FROM Student;

-- 删除临时表格
DROP TEMPORARY TABLE IF EXISTS TempStudents;

-- 使用 insert 语句增加两门课程：C 语言和人工智能，并查询出来
-- 插入两门新课程
INSERT INTO Course (Cno, Cname, Cpno, Ccredit) VALUES
('8', 'C语言', NULL, 3),
('9', '人工智能', NULL, 4);

-- 查询新插入的课程
SELECT * FROM Course;

-- 使用 delete 语句把人工智能课程删除，并查询出来
-- 删除人工智能课程
DELETE FROM Course WHERE Cno = '9';

-- 查询删除后的结果
SELECT * FROM Course;
