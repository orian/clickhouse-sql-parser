CREATE HANDLER get_user URL REGEXP '/users/(?P<id>\d+)' AS SELECT * FROM users WHERE id = {id:UInt64};
