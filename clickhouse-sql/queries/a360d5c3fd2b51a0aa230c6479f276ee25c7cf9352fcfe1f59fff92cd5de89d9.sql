
        INSERT INTO with_lonely SELECT number+60000, '2022-11-01', number*10, 2 FROM
        numbers(500000);
    