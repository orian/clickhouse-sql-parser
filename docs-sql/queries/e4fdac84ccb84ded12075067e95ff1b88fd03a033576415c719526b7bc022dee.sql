SELECT player, salary, 
       first_value(player) OVER (ORDER BY salary DESC) AS highest_paid_player
FROM salaries;
