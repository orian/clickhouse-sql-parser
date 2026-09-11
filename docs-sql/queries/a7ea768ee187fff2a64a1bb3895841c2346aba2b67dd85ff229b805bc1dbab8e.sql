SELECT parseDateTimeBestEffortOrZero('23/10/2025 12:12:57') AS valid,
       parseDateTimeBestEffortOrZero('invalid') AS invalid
