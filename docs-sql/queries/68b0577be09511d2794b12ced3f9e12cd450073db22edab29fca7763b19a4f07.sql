SET param_id = 42;
SET param_name = 'John';

SELECT * FROM users
WHERE id = {id: UInt32}
AND name = {name: String};
