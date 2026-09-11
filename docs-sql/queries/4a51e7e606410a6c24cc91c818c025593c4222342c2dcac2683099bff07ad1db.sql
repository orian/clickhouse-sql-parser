SELECT id, sequenceNextNode('forward', 'first_match')(dt, page, page = 'Gift', page = 'Gift', page = 'Home') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift // 基準点
1970-01-01 09:00:03    1   Exit // Home に一致しない

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home
1970-01-01 09:00:03    2   Gift // 基準点
1970-01-01 09:00:04    2   Basket // Home に一致しない

1970-01-01 09:00:01    3   Gift // 基準点
1970-01-01 09:00:02    3   Home // Home に一致
1970-01-01 09:00:03    3   Gift // 結果
1970-01-01 09:00:04    3   Basket
