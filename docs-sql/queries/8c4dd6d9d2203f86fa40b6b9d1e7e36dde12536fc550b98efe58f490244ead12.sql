SELECT toDate('2025-01-01') + toInterval(number, 'day') AS dates
FROM numbers(5)
