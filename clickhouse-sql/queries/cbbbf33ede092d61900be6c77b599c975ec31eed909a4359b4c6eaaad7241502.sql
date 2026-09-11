
        select * from test_left all inner join test_right on test_left.k1 = test_right.k1  SETTINGS {settings} format Null
    