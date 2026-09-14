SELECT id, sequenceNextNode('backward', 'tail')(dt, page, page = 'Basket', page = 'Basket', page = 'Gift') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift
1970-01-01 09:00:03    1   Exit // 基准点，未匹配到 Basket

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home // 结果
1970-01-01 09:00:03    2   Gift // 匹配到 Gift
1970-01-01 09:00:04    2   Basket // 基准点，匹配到 Basket

1970-01-01 09:00:01    3   Gift
1970-01-01 09:00:02    3   Home // 结果
1970-01-01 09:00:03    3   Gift // 基准点, Matched with Gift
1970-01-01 09:00:04    3   Basket // 基准点，匹配到 Basket
