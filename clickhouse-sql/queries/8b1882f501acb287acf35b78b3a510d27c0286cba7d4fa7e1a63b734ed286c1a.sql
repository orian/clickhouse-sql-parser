
        INSERT INTO with_lonely SELECT number+500000, '2022-10-30', number*10, 1 FROM
        numbers(3000000);
    