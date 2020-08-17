/*
Write a query identifying the type of each record in the 'TRIANGLES' table using its three side lengths.

Output one of the following statements for each record in the table:
	Equilateral: It's a triangle with 3 sides of equal length.
	Isosceles: It's a triangle with 2 sides of equal length.
	Scalene: It's a triangle with 3 sides of differing lengths.
	Not A Triangle: The given values of A, B, and C don't form a triangle.
*/

SELECT
CASE
    WHEN (a+b+c - GREATEST(a,b,c)) <= GREATEST(a,b,c) THEN "Not A Triangle"
    WHEN a=b AND b=c THEN "Equilateral"
    WHEN a=b OR b=c OR a=c THEN "Isosceles"
    ELSE "Scalene"
END
FROM triangles

/*
You are given a table, BST, containing two columns: N and P,
where N represents the value of a node in Binary Tree, and P is the parent of N.

Write a query to find the node type of Binary Tree ordered by the value of the node.

Output one of the following for each node:
	Root: If node is root node.
	Leaf: If node is leaf node.
	Inner: If node is neither root nor leaf node.
*/

SELECT n,
CASE
    WHEN p IS NULL THEN "Root"
    WHEN n IN (
        SELECT p
        FROM BST
    ) THEN "Inner"
    ELSE "Leaf"
END
FROM BST
ORDER BY n


/*
Generate the following two result sets:

Query an alphabetically ordered list of all names in OCCUPATIONS,
    immediately followed by the first letter of each profession as a parenthetical (i.e.: enclosed in parentheses).
    For example: AnActorName(A), ADoctorName(D), AProfessorName(P), and ASingerName(S).

Query the number of ocurrences of each occupation in OCCUPATIONS.
	Sort the occurrences in ascending order, and output them in the following format:
	There are a total of [occupation_count] [occupation]s.
	where [occupation_count] is the number of occurrences of an occupation in OCCUPATIONS
	and [occupation] is the lowercase occupation name.
	If more than one Occupation has the same [occupation_count], they should be ordered alphabetically.
*/
SELECT FinalString
FROM (
    (
        SELECT CONCAT(name,"(",LEFT(occupation,1),")") AS FinalString,
            name as NameOrder, 0 as AmountOrder
        FROM occupations
    )
    UNION
    (
        SELECT CONCAT("There are a total of ",amount," ",LOWER(occupation),"s.") AS FinalString,
            'zzzzz' as NameOrder, amount as AmountOrder
        FROM (
            SELECT COUNT(occupation) AS amount, occupation
            FROM occupations
            GROUP BY occupation
        ) T2
    )
) T1
ORDER BY NameOrder, AmountOrder

/*
Pivot the Occupation column in OCCUPATIONS so that each Name is sorted alphabetically
and displayed underneath its corresponding Occupation.
The output column headers should be Doctor, Professor, Singer, and Actor, respectively.

Note: Print NULL when there are no more names corresponding to an occupation.

Input Format
*/
SET @r1=0, @r2=0, @r3=0, @r4=0;
SELECT MIN(Doctor), MIN(Professor), MIN(Singer), MIN(Actor)
FROM(
  SELECT
    CASE
        WHEN occupation='Doctor' THEN (@r1:=@r1+1)
        WHEN occupation='Professor' THEN (@r2:=@r2+1)
        WHEN occupation='Singer' THEN (@r3:=@r3+1)
        WHEN occupation='Actor' THEN (@r4:=@r4+1)
    END AS RowNumber,
    CASE WHEN occupation='Doctor' THEN name END AS Doctor,
    CASE WHEN occupation='Professor' THEN name END AS Professor,
    CASE WHEN occupation='Singer' THEN name END AS Singer,
    CASE WHEN occupation='Actor' THEN name END AS Actor
  FROM OCCUPATIONS
  ORDER BY name
) Temp
GROUP BY RowNumber

/*
	Consider P1(a,b) and P2(c,d) to be two points on a 2D plane.

	a happens to equal the minimum value in Northern Latitude (LAT_N in STATION).
	b happens to equal the minimum value in Western Longitude (LONG_W in STATION).
	c happens to equal the maximum value in Northern Latitude (LAT_N in STATION).
	d happens to equal the maximum value in Western Longitude (LONG_W in STATION).
	Query the Manhattan Distance between points P1 and P2 and round it to a scale of 4 decimal places.
*/
SELECT ROUND(MAX(LAT_N) - MIN(LAT_N) + MAX(LONG_W) - MIN(LONG_W),4)
FROM station

/*
	Amber's conglomerate corporation just acquired some new companies. Each of the companies follows this hierarchy:
			Founder  =>  Lead Manager  =>  Senior Manager  =>  Manager  =>  Employee
			
	Given the table schemas below, write a query to print the
		company_code, founder name, total number of lead managers, total number of senior managers, total number of managers, and total number of employees.
		Order your output by ascending company_code.
	Note:
		- The tables may contain duplicate records.
		- The company_code is string, so the sorting should not be numeric.
		  For example, if the company_codes are C_1, C_2, and C_10,
		  then the ascending company_codes will be C_1, C_10, and C_2.
*/
SELECT company_code as CCode, founder,
    (SELECT COUNT(DISTINCT(lead_manager_code))
     FROM Lead_Manager
     WHERE company_code = CCode),
    (SELECT COUNT(DISTINCT(senior_manager_code))
     FROM Senior_Manager
     WHERE company_code = CCode),
    (SELECT COUNT(DISTINCT(manager_code))
     FROM Manager
     WHERE company_code = CCode),
    (SELECT COUNT(DISTINCT(employee_code))
     FROM Employee
     WHERE company_code = CCode)
FROM company
ORDER BY CCode

/*
	Consider P1(a,c) and P2(b,d) to be two points on a 2D plane
	where (a,b) are the respective minimum and maximum values of Northern Latitude (LAT_N)
	and (b,d) are the respective minimum and maximum values of Western Longitude (LONG_W) in STATION.

	Query the Euclidean Distance between points P1 and P2 and format your answer to display 4 decimal digits.
*/
SELECT ROUND(SQRT(POWER(MAX(LAT_N) - MIN(LAT_N),2) + POWER(MAX(LONG_W) - MIN(LONG_W),2)),4)
FROM station

/*
	A median is defined as a number separating the higher half of a data set from the lower half.
	Query the median of the Northern Latitudes (LAT_N) from STATION and round your answer to 4 decimal places.
*/
SELECT ROUND(AVG(lat_n),4) as median
FROM (
    SELECT lat_n, @rownum:=@rownum+1 as `row_number`, @total_rows:=@rownum
      FROM station, (SELECT @rownum:=0) r
      ORDER BY lat_n
) T1
WHERE T1.row_number IN ( FLOOR((@total_rows+1)/2), FLOOR((@total_rows+2)/2));