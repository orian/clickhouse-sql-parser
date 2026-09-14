    SET param_name = 'John Doe';
    SET param_age = 25;
    SET param_salary = 75000.50;

    SELECT name, age, salary FROM users
    WHERE name = {name: String}
      AND age >= {age: UInt8}
      AND salary <= {salary: Float64};
