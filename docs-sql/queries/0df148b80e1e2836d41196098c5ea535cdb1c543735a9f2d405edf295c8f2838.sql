SELECT parseDateTime64BestEffortOrZero('2025-01-01 01:01:00.123') AS valid,
       parseDateTime64BestEffortOrZero('invalid') AS invalid
