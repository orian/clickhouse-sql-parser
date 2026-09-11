
        insert into test_non_function_column_short_circuit select cast(rand()%2 == 1 as UInt8) as a, cast(rand()%3 == 1 as UInt8) as b, cast(rand()%4 == 1 as UInt8) as c, cast(rand()%5 == 1 as UInt8) as d from numbers(10000000)
    