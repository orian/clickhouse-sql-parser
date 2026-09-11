SELECT
    parseDateTime32BestEffortOrNull('23/10/2025 12:12:57') AS valid,
    parseDateTime32BestEffortOrNull('invalid date') AS invalid
