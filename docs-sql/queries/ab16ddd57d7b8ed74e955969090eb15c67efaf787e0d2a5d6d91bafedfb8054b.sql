SELECT player, salary,
       rank() OVER (ORDER BY salary DESC) AS rank
FROM salaries;
