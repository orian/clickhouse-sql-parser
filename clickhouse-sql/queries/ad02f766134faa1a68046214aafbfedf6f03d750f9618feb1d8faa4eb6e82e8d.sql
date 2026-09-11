
    INSERT INTO tab_src_{drop_pct}
    SELECT
      number,
      arrayStringConcat(arrayMap(x -> if(cityHash64(number, x) % 100 < {drop_pct},
        ['a', 'an', 'and', 'are', 'as', 'at', 'be', 'but', 'by', 'for', 'if', 'in', 'into', 'is', 'it', 'no', 'not', 'of', 'on', 'or', 'such', 'that', 'the', 'their', 'then', 'there', 'these', 'they', 'this', 'to', 'was', 'will', 'with'][(cityHash64(number, x, 1) % 33) + 1],
        ['fox', 'cat', 'dog', 'bird', 'tree', 'river', 'stone'][(cityHash64(number, x, 2) % 7) + 1]), range(30)), ' ')
    FROM numbers(3000000);
  