# Write your MySQL query statement below
WITH moreThan5 AS(
    SELECT
    managerId
    FROM Employee
    GROUP BY managerId
    HAVING COUNT(managerId)>=5
)

SELECT
e.name
FROM Employee e
INNER JOIN moreThan5 m
ON e.id = m.managerId