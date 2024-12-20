SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
-- https://leetcode.com/problems/combine-two-tables/?envType=problem-list-v2&envId=database

SELECT firstName, lastName, city, state
FROM Person
LEFT JOIN Address ON Person.personId = Address.personId 