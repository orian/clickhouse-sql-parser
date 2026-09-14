SELECT player, salary, nth_value(player,3) OVER(ORDER BY salary DESC) AS third_highest_salary FROM salaries;
