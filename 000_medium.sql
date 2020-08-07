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