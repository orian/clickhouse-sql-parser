SET allow_deprecated_error_prone_window_functions = 1;
SELECT
    number,
    runningDifference(number + 1) AS diff
FROM numbers(100000)
WHERE diff != 1;
