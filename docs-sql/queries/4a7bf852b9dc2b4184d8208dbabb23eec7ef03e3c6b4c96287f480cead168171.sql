SELECT id, sequenceNextNode('backward', 'tail')(dt, page, page = 'Basket', page = 'Basket', page = 'Gift') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift
1970-01-01 09:00:03    1   Exit // point de base, Unmatched with Basket

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home // The result
1970-01-01 09:00:03    2   Gift // Matched with Gift
1970-01-01 09:00:04    2   Basket // point de base, Matched with Basket

1970-01-01 09:00:01    3   Gift
1970-01-01 09:00:02    3   Home // The result
1970-01-01 09:00:03    3   Gift // point de base, Matched with Gift
1970-01-01 09:00:04    3   Basket // point de base, Matched with Basket
