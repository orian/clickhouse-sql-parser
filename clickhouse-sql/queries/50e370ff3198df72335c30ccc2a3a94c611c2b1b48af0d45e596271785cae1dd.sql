
        INSERT INTO with_lonely SELECT number, '2022-11-01', number*10, 0 FROM numbers(500000);
    