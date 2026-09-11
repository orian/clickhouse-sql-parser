SELECT player, salary,
       percent_rank() OVER (ORDER BY salary DESC) AS percent_rank
FROM salaries;
