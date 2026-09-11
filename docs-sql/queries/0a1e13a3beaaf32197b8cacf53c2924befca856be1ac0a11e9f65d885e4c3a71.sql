SET allow_deprecated_error_prone_window_functions = 1;
SELECT
    number,
    runningDifferenceStartingWithFirstValue(number) AS diff
FROM numbers(5);
