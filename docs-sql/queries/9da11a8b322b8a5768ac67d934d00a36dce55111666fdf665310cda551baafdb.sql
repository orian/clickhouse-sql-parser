    SELECT * FROM events WHERE COALESCE(regexp_match(msg, '^ERR'), '{}');
