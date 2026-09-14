SELECT
toMonday(toDateTime('2023-04-21 10:20:30')), -- A Friday
toMonday(toDate('2023-04-24'));              -- Already a Monday
