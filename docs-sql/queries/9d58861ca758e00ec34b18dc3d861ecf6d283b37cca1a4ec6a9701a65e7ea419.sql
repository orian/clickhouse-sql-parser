SELECT
    department,
    employee_name AS emp,
    salary,
    max_salary_per_dep,
    total_salary_per_dep,
    round((salary / total_salary_per_dep) * 100, 2) AS `share_per_dep(%)`
FROM
(
    SELECT
        department,
        employee_name,
        salary,
        max(salary) OVER wndw AS max_salary_per_dep,
        sum(salary) OVER wndw AS total_salary_per_dep
    FROM employees
    WINDOW wndw AS (
        PARTITION BY department
        ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
    )
    ORDER BY
        department ASC,
        employee_name ASC
);

┌─department─┬─emp──┬─salary─┬─max_salary_per_dep─┬─total_salary_per_dep─┬─share_per_dep(%)─┐
│ Finance    │ Jean │    505 │                505 │                  915 │            55.19 │
│ Finance    │ Joan │    210 │                505 │                  915 │            22.95 │
│ Finance    │ Jonh │    200 │                505 │                  915 │            21.86 │
│ IT         │ Anna │    300 │                500 │                 1000 │               30 │
│ IT         │ Elen │    500 │                500 │                 1000 │               50 │
│ IT         │ Tim  │    200 │                500 │                 1000 │               20 │
└────────────┴──────┴────────┴────────────────────┴──────────────────────┴──────────────────┘
