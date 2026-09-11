SELECT parseDateTime64BestEffortUS('02/10/2025 12:30:45.123') AS us_format,
       parseDateTime64BestEffortUS('15/08/2025 10:15:30.456') AS fallback_to_standard
