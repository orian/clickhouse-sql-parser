
        INSERT INTO t_json_group_by
        SELECT toJSONString(map('a', toString(number % 1000), 'b', toString(number % 7),
                                'c', toString(number % 13),   'd', toString(number % 97),
                                'e', toString(number % 31),   'f', toString(number % 53),
                                'g', toString(number % 11)))::JSON
        FROM numbers(1000000)
    