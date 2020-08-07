/* Query the average population for all cities in CITY, rounded down to the nearest integer.*/
SELECT CAST(AVG(population) AS SIGNED)
FROM city