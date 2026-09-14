SELECT toTimeWithFixedDate('2025-06-15 12:00:00'::DateTime) - toTimeWithFixedDate('2024-05-10 11:00:00'::DateTime) AS result, toTypeName(result)
