SELECT parseDateTimeBestEffortUSOrNull('02/10/2025') AS valid_us,
       parseDateTimeBestEffortUSOrNull('invalid') AS invalid
