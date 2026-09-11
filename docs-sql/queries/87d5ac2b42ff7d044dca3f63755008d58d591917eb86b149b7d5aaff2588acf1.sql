SELECT
    player,
    salary,
    team,
    avg(salary) OVER (PARTITION BY team) AS teamAvg,
    salary - teamAvg AS diff
FROM salaries;
