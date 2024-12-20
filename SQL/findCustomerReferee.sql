SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
SET NOCOUNT ON;
-- https://leetcode.com/problems/find-customer-referee/?envType=study-plan-v2&envId=top-sql-50

SELECT name FROM Customer
WHERE NOT referee_id = 2 OR referee_id IS NULL