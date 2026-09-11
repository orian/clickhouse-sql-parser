
        SELECT
            COUNT()
        FROM
            jit_test
        WHERE
            NOT ignore(a / b + c / d + e / f + g / h + i / j)
        SETTINGS
            compile_expressions = 0;
    