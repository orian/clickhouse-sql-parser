
        INSERT INTO with_lonely SELECT number+70000, '2022-11-01', number*10, 3 FROM
        numbers(500000);
    