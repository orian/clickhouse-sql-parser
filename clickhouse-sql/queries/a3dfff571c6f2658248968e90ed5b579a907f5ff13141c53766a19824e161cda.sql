
        INSERT INTO string_subcolumn_tuple
        SELECT tuple(
            's', 's', 's', 's', 's',
            's', 's', 's', 's', 's',
            's', 's', 's', 's', 's',
            's', 's', 's', 's', 's'
        )
        FROM numbers(10000000)
    