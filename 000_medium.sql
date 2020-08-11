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