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
	
/* Query the average population of all cities in CITY where District is California. */
SELECT AVG(population)
FROM city
WHERE district = 'California'

/* Query the difference between the maximum and minimum populations in CITY. */
SELECT MAX(population) - MIN(population)
FROM city

/*
	Samantha was tasked with calculating the average monthly salaries for all employees in the EMPLOYEES table,
	but did not realize her keyboard's 0 key was broken until after completing the calculation.
	She wants your help finding the difference between her miscalculation (using salaries with any zeroes removed),
	and the actual average salary.

	Write a query calculating the amount of error (i.e.: 'actual' - 'miscalculated' average monthly salaries),
	and round it up to the next integer.
*/
SELECT CEILING(AVG(salary) - AVG(REPLACE(salary,'0','')))
FROM employees

/*
	Query the following two values from the STATION table:
		- The sum of all values in LAT_N rounded to a scale of 2 decimal places.
		- The sum of all values in LONG_W rounded to a scale of 2 decimal places.
*/
SELECT ROUND(SUM(lat_n),2), ROUND(SUM(long_w),2)
FROM station

/*
	Query the sum of Northern Latitudes (LAT_N) from STATION having values greater than 38.7880 and less than 137.2345.
	Truncate your answer to 4 decimal places.
*/
SELECT ROUND(SUM(lat_n),4)
FROM station
WHERE lat_n > 38.7880 AND lat_n < 137.2345

/*
	Query the Western Longitude (LONG_W)
	where the smallest Northern Latitude (LAT_N) in STATION is greater than 38.7780.
	Round your answer to 4 decimal places.
*/
SELECT ROUND(long_w,4)
FROM station
WHERE lat_n > 38.7780
ORDER BY lat_n
LIMIT 1

/*
	Query the greatest value of the Northern Latitudes (LAT_N) from STATION
	that is less than 137.2345.
	Truncate your answer to 4 decimal places.
*/
SELECT ROUND(MAX(lat_n),4)
FROM station
WHERE lat_n < 137.2345

/*
	Query the smallest Northern Latitude (LAT_N) from STATION that is greater than 38.7780.
	Round your answer to 4 decimal places.
*/
SELECT ROUND(MIN(lat_n),4)
FROM station
WHERE lat_n > 38.7780