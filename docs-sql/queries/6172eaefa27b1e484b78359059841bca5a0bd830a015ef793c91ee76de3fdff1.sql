SELECT
    parseDateTime32BestEffortOrZero('23/10/2025 12:12:57') AS valid,
    parseDateTime32BestEffortOrZero('invalid date') AS invalid
