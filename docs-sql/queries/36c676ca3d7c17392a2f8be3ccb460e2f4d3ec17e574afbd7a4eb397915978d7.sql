-- Range expansion can be beneficial when you need a similar distribution for
-- arguments with wildly different ranges (or cardinality).
-- For example: 'IP Address' (0...FFFFFFFF) and 'Country code' (0...FF).
-- Note: tuple size must be equal to the number of the other arguments.
SELECT hilbertEncode((10, 6), 1024, 16)
