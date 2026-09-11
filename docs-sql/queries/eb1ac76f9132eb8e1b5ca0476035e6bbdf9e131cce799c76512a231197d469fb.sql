SELECT player, salary,
       ntile(4) OVER (ORDER BY salary DESC, player ASC) AS bucket
FROM salaries;
