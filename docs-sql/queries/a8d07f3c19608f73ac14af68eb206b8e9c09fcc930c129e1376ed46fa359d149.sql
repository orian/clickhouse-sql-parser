  INSERT INTO test_table (id, name, email, age, salary, created_at, is_active, department, score, country)
  SELECT
      generateUUIDv4() as id,
      concat('User_', toString(rand() % 10000)) as name,
      concat('user', toString(rand() % 10000), '@example.com') as email,
      18 + (rand() % 65) as age,
      30000 + (rand() % 100000) as salary,
      now() - toIntervalSecond(rand() % 31536000) as created_at,
      rand() % 2 as is_active,
      arrayElement(['Engineering', 'Marketing', 'Sales', 'HR', 'Finance', 'Operations'], (rand() % 6) + 1) as department,
      rand() / 4294967295.0 * 100 as score,
      arrayElement(['USA', 'UK', 'Germany', 'France', 'Canada', 'Australia', 'Japan', 'Brazil'], (rand() % 8) + 1) as country
  FROM numbers(1000);
