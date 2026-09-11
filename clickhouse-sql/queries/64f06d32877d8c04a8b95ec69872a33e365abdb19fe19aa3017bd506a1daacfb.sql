
        insert into test_right
        SELECT
        toString(number % 20),
        toString(number * 10000),
        toString(number * 10000 + 1),
        toString(number * 10000 + 2),
        toString(number * 10000 + 3),
        toString(number * 10000 + 4),
        toString(number * 10000 + 5),
        toString(number * 10000 + 6),
        toString(number * 10000 + 7),
        toString(number * 10000 + 8)
        from system.numbers limit 10000;
    