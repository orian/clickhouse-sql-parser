SELECT parseDateTimeBestEffortOrNull('23/10/2025 12:12:57') AS valid,
       parseDateTimeBestEffortOrNull('invalid') AS invalid
