/*
 * Compute the number of days between two dates.
 * /* Returns NULL if either argument is NULL */
 */
SELECT
    dateDiff('day', toDate('2024-01-01'), toDate('2024-12-31')) AS days_in_year, -- 365
    dateDiff('day', toDate('2020-01-01'), today()) AS days_since  #! since 2020
    ///////////////////////////////////////////////////////////////////
    # TODO: add hour/minute variants
