
        create table if not exists whitespaces
        engine = MergeTree() partition by tuple() order by tuple()
        as
            with 32 - log2(intHash32(number)) + 1 as num_spaces,
                repeat(' ', toUInt32(num_spaces)) as spaces
            select spaces || toString(number) || spaces value
            from numbers_mt(100000000);
    