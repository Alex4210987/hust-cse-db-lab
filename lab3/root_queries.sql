-- (9) 用系统管理员登录
USE S_T_U2022xxxxx;
-- 在下一步中将定义触发器，需要管理员权限

-- (10) 对 SC 表建立一个更新触发器
DELIMITER //
CREATE TRIGGER update_scholarship
AFTER UPDATE ON S_T_U2022xxxxx.SC
FOR EACH ROW
BEGIN
    IF NEW.Grade >= 95 THEN
        UPDATE S_T_U2022xxxxx.Student SET Scholarship = '是' WHERE Sno = NEW.Sno;
    ELSE
        IF NOT EXISTS (SELECT * FROM S_T_U2022xxxxx.SC WHERE Sno = NEW.Sno AND Grade >= 95) AND OLD.Grade >= 95 THEN
            UPDATE S_T_U2022xxxxx.Student SET Scholarship = '否' WHERE Sno = NEW.Sno;
        END IF;
    END IF;
END;
//
DELIMITER ;

-- 1）首先把某个学生成绩修改为 98，查询其奖学金
UPDATE S_T_U2022xxxxx.SC SET Grade = 98 WHERE Sno = '200215121' AND Cno = '1';
SELECT Scholarship FROM S_T_U2022xxxxx.Student WHERE Sno = '200215121';

-- 2）再把刚才的成绩修改为 80，再查询其奖学金
UPDATE S_T_U2022xxxxx.SC SET Grade = 80 WHERE Sno = '200215121' AND Cno = '1';
SELECT Scholarship FROM S_T_U2022xxxxx.Student WHERE Sno = '200215121';

-- (11) 删除刚定义的触发器
DROP TRIGGER IF EXISTS update_scholarship;

-- (12) 定义一个存储过程计算 CS 系的课程的平均成绩和最高成绩
DELIMITER //
CREATE PROCEDURE calculate_cs_course_stats()
BEGIN
    SELECT AVG(Grade) AS average_grade, MAX(Grade) AS max_grade
    FROM S_T_U2022xxxxx.SC
    WHERE Cno IN (SELECT Cno FROM S_T_U2022xxxxx.Course WHERE Cpno IS NULL);
END;
//
DELIMITER ;

-- 执行存储过程
CALL calculate_cs_course_stats();

-- (13) 定义一个带学号为参数的查看某个学号的所有课程的成绩
DELIMITER //
CREATE PROCEDURE view_student_courses(sno_param CHAR(9))
BEGIN
    SELECT S_T_U2022xxxxx.Student.Sname, S_T_U2022xxxxx.Student.Sno, S_T_U2022xxxxx.SC.Cno, S_T_U2022xxxxx.Course.Cname, S_T_U2022xxxxx.SC.Grade
    FROM S_T_U2022xxxxx.Student
    JOIN  S_T_U2022xxxxx.SC ON S_T_U2022xxxxx.Student.Sno = S_T_U2022xxxxx.SC.Sno
    JOIN S_T_U2022xxxxx.Course ON S_T_U2022xxxxx.SC.Cno = S_T_U2022xxxxx.Course.Cno
    WHERE S_T_U2022xxxxx.Student.Sno = sno_param;
END;
//
DELIMITER ;

-- 执行存储过程
CALL view_student_courses('200215121');

-- (14) 把上一题改成函数
-- DELIMITER //
-- CREATE FUNCTION get_student_courses(sno_param CHAR(9))
-- RETURNS TABLE (
--     Sname VARCHAR(255),
--     Cno INT,
--     Grade INT
-- )
-- DETERMINISTIC
-- READS SQL DATA
-- BEGIN
--     RETURN (
--         SELECT S_T_U2022xxxxx.Student.Sname, S_T_U2022xxxxx.SC.Cno, S_T_U2022xxxxx.SC.Grade
--         FROM S_T_U2022xxxxx.Student
--         JOIN S_T_U2022xxxxx.SC ON S_T_U2022xxxxx.Student.Sno = sno_param
--     );
-- END;
-- //
-- DELIMITER ;

DELIMITER //

CREATE FUNCTION get_student_courses(sno_param CHAR(9))
RETURNS VARCHAR(2048)
DETERMINISTIC
READS SQL DATA
BEGIN
    DECLARE result VARCHAR(2048);

    SELECT GROUP_CONCAT(CONCAT(S.Sno, ' ', S.Sname, ' ', C.Cno, ' ', C.Cname, ' ', SC.Grade) SEPARATOR '; ')
    INTO result
    FROM Student S
    JOIN SC ON S.Sno = SC.Sno
    JOIN Course C ON SC.Cno = C.Cno
    WHERE S.Sno = sno_param;

    RETURN result;
END //

DELIMITER ;
121');

-- (15) 在 SC 表上定义一个完整性约束，要求成绩在 0-100 之间
-- 修改某个学生的成绩为 120，进行查询
UPDATE S_T_U2022xxxxx.SC SET Grade = 120 WHERE Sno = '200215121' AND Cno = '1';
SELECT * FROM S_T_U2022xxxxx.SC WHERE Sno = '200215121' AND Cno = '1';

-- 修改回来
UPDATE S_T_U2022xxxxx.SC SET Grade = 95 WHERE Sno = '200215121' AND Cno = '1';

-- 定义约束后，再把该学生成绩修改为 120
ALTER TABLE S_T_U2022xxxxx.SC ADD CONSTRAINT check_grade_range CHECK (Grade >= 0 AND Grade <= 100);

-- 进行查询
-- 此时不应该允许将成绩修改为 120
UPDATE S_T_U2022xxxxx.SC SET Grade = 120 WHERE Sno = '200215121' AND Cno = '1';
SELECT * FROM S_T_U2022xxxxx.SC WHERE Sno = '200215121' AND Cno = '1';