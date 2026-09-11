
    INSERT INTO tab
    SELECT
        number,
        if(number % 2 = 0,
            concat('alpha-service-', toString(number % 10), '-prod'),
            concat('beta-service-', toString(number % 10), '-dev'))
    FROM numbers(10000000)
  