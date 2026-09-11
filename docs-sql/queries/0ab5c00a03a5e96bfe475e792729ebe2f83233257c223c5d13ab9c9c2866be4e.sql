SELECT parseDateTime64BestEffortUSOrZero('02/10/2025 12:30:45.123') AS valid_us,
       parseDateTime64BestEffortUSOrZero('invalid') AS invalid
