
        INSERT INTO json_string_value_formatting
        SELECT
            leftPad(toString(number), 70, 'x'),
            concat(leftPad(toString(number), 58, 'x'), '"a"', unhex('5c'), 'c/d', unhex('0a'), 'e'),
            concat(leftPad(toString(number), 64, 'x'), unhex('e280a8'), 'ab'),
            concat(repeat(concat('"', unhex('5c'), '/', unhex('0a')), 17), leftPad(toString(number), 2, 'x')),
            concat('{{"u":"https://host.example/', toString(number), '/a/b/c","s":"plain value"}}')
        FROM numbers(1700000)
    