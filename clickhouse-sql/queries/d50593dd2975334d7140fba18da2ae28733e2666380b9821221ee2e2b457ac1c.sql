
        insert into webpages
            select toString(number % 100), randomString(100),randomString(4000) from numbers(100000)
        settings optimize_on_insert = 0;
    