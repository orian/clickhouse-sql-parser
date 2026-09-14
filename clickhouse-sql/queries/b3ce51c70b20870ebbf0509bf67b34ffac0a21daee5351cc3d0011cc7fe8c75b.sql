
        insert into test_orc_fpd select number as a, cast(number as String) as b from numbers(10000000)
    