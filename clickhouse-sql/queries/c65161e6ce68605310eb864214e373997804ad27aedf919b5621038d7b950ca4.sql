
        INSERT INTO with_lonely SELECT number+50000, '2022-11-01', number*10, 1 FROM
        numbers(500000);
    