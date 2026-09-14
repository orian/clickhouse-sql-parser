SELECT id, sequenceNextNode('forward', 'head')(dt, page, page = 'Home', page = 'Home', page = 'Gift') FROM test_flow GROUP BY id;

                  dt   id   page
 1970-01-01 09:00:01    1   Home // 基准点，与 Home 匹配
 1970-01-01 09:00:02    1   Gift // 与 Gift 匹配
 1970-01-01 09:00:03    1   Exit // 结果

 1970-01-01 09:00:01    2   Home // 基准点，与 Home 匹配
 1970-01-01 09:00:02    2   Home // 与 Gift 不匹配
 1970-01-01 09:00:03    2   Gift
 1970-01-01 09:00:04    2   Basket

 1970-01-01 09:00:01    3   Gift // 基准点，与 Home 不匹配
 1970-01-01 09:00:02    3   Home
 1970-01-01 09:00:03    3   Gift
 1970-01-01 09:00:04    3   Basket
