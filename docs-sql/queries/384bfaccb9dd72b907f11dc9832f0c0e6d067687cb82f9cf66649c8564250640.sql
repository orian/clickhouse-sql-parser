CREATE MASKING POLICY mask_high_salaries ON employees
UPDATE salary = 0
WHERE salary > 100000
TO analyst;
