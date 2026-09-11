SELECT parseDateTimeBestEffortUSOrZero('02/10/2025') AS valid_us,
       parseDateTimeBestEffortUSOrZero('invalid') AS invalid
