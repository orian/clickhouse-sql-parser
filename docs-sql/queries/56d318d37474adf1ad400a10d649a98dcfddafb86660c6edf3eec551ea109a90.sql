SELECT parseDateTime64BestEffortOrNull('2025-01-01 01:01:00.123') AS valid,
       parseDateTime64BestEffortOrNull('invalid') AS invalid
