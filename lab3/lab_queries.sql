-- the db is inited with the same svheema as in lab1

-- 切换到数据库
USE S_T_U2022xxxxx;

-- (1) 创建 CS 系的视图 CS_View
CREATE VIEW CS_View AS
SELECT * FROM Student WHERE Sdept = 'CS';

-- (2) 在视图 CS_View 上查询 CS 系选修了 1 号课程的学生
SELECT * FROM CS_View WHERE Sno IN (SELECT Sno FROM SC WHERE Cno = '1');

-- (3) 创建 IS 系成绩大于 80 的学生的视图 IS_View
CREATE VIEW IS_View AS
SELECT * FROM Student WHERE Sdept = 'IS' AND (SELECT AVG(Grade) FROM SC WHERE Sno = Student.Sno) > 80;

-- (4) 在视图 IS_View 查询 IS 系成绩大于 80 的学生
SELECT * FROM IS_View;

-- (5) 删除视图 IS_View
DROP VIEW IF EXISTS IS_View;

-- (6) 利用可视化窗口创建 2 个不同的用户 U1 和 U2
DROP USER IF EXISTS 'U1'@'localhost';
CREATE USER 'U1'@'localhost' IDENTIFIED BY '';
DROP USER IF EXISTS 'U2'@'localhost';
CREATE USER 'U2'@'localhost' IDENTIFIED BY '';
FLUSH PRIVILEGES;

-- (6) 给 U1 授予 Student 表的查询和更新的权限
GRANT SELECT, UPDATE ON S_T_U2022xxxxx.Student TO 'U1'@'localhost';

-- (6) 给 U2 对 SC 表授予插入的权限
GRANT INSERT ON S_T_U2022xxxxx.SC TO 'U2'@'localhost';

FLUSH PRIVILEGES;

-- (6) 用 U1 登录
-- 1）查询学生表的信息
SELECT * FROM S_T_U2022xxxxx.Student;

-- 2）把所有学生的年龄增加 1 岁，然后查询
UPDATE S_T_U2022xxxxx.Student SET Sage = Sage + 1;
SELECT * FROM S_T_U2022xxxxx.Student;

-- 3）删除 IS 系的学生
DELETE FROM S_T_U2022xxxxx.Student WHERE Sdept = 'IS';

-- 4）查询 CS 系的选课信息
SELECT Student.Sno, Student.Sname, Course.Cno, Course.Cname 
FROM SC, Course, Student
WHERE Student.Sno = SC.Sno AND SC.Cno = SC.Cno
ORDER BY Student.Sno;

-- (6) 用 U2 登录
-- 1）在 SC 表中插入 1 条记录（‘200215122’，‘1’，75）
INSERT INTO S_T_U2022xxxxx.SC (Sno, Cno, Grade) VALUES ('200215122', '1', 75);

-- 2）查询 SC 表的信息
SELECT * FROM S_T_U2022xxxxx.SC;

-- 3）查询视图 CS_View 的信息
SELECT * FROM CS_View;

-- (7) 用系统管理员登录，收回 U1 的所有权限
REVOKE SELECT, UPDATE ON S_T_U2022xxxxx.Student FROM 'U1'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;