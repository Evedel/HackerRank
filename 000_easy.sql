/* Query the average population for all cities in CITY, rounded down to the nearest integer.*/
SELECT CAST(AVG(population) AS SIGNED)
FROM city

/*Query the sum of the populations for all Japanese cities in CITY. The COUNTRYCODE for Japan is JPN. */
SELECT SUM(population)
FROM city
WHERE countrycode = 'JPN'

/* Query a count of the number of cities in CITY having a Population larger than 100,000. */
select count(*)
from city
where population > 100000

/* 
	We define an employee's total earnings to be their monthly 'salary'*'months' worked,
	and the maximum total earnings to be the maximum total earnings for any employee in the Employee table.
	Write a query to find the maximum total earnings for all employees
	as well as the total number of employees who have maximum total earnings.
	Then print these values as 2 space-separated integers.
*/
select months*salary, COUNT(*)
FROM Employee
GROUP BY months*salary
ORDER BY months*salary DESC
LIMIT 1

/* Query the total population of all cities in CITY where District is California. */
SELECT SUM(population)
FROM city
WHERE district = 'California'

/*
	Query the Western Longitude (LONG_W) for the largest Northern Latitude (LAT_N)
	in STATION that is less than 137.2345.
	Round your answer to 4 decimal places.
*/
SELECT ROUND(LONG_W,4)
FROM station
WHERE LAT_N = (
        SELECT MAX(LAT_N)
        FROM station
        WHERE LAT_N < 137.2345
    )