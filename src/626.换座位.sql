--
-- @lc app=leetcode.cn id=626 lang=mysql
--
-- [626] 换座位
--
-- https://leetcode-cn.com/problems/exchange-seats/description/
--
-- database
-- Medium (65.38%)
-- Likes:    134
-- Dislikes: 0
-- Total Accepted:    13.1K
-- Total Submissions: 20.1K
-- Testcase Example:  '{"headers": {"seat": ["id","student"]}, "rows": {"seat": [[1,"Abbot"],[2,"Doris"],[3,"Emerson"],[4,"Green"],[5,"Jeames"]]}}'
--
-- 小美是一所中学的信息科技老师，她有一张 seat 座位表，平时用来储存学生名字和与他们相对应的座位 id。
-- 
-- 其中纵列的 id 是连续递增的
-- 
-- 小美想改变相邻俩学生的座位。
-- 
-- 你能不能帮她写一个 SQL query 来输出小美想要的结果呢？
-- 
-- 
-- 
-- 示例：
-- 
-- 
-- +---------+---------+
-- |    id   | student |
-- +---------+---------+
-- |    1    | Abbot   |
-- |    2    | Doris   |
-- |    3    | Emerson |
-- |    4    | Green   |
-- |    5    | Jeames  |
-- +---------+---------+
-- 
-- 
-- 假如数据输入的是上表，则输出结果如下：
-- 
-- 
-- +---------+---------+
-- |    id   | student |
-- +---------+---------+
-- |    1    | Doris   |
-- |    2    | Abbot   |
-- |    3    | Green   |
-- |    4    | Emerson |
-- |    5    | Jeames  |
-- +---------+---------+
-- 
-- 注意：
-- 
-- 如果学生人数是奇数，则不需要改变最后一个同学的座位。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

SELECT id,student from 
(
SELECT student ,
	if((SELECT COUNT(*) FROM seat)%2=1 and (id = (SELECT count(*) from seat)),id, if(id%2=1,id+1,id-1)) as id
FROM seat
) as tmp
ORDER BY id

-- @lc code=end

