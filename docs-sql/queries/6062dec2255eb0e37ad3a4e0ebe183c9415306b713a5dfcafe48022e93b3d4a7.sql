-- The type returned is a DateTime64 as the DateTime32 must be promoted to 64 bit for the comparison.
SELECT least(toDateTime32('2025-01-02 12:00:00'), toDateTime64('2025-01-01 12:00:00.000', 3));
