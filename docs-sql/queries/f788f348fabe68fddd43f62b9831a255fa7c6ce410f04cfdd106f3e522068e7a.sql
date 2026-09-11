SELECT formatReadableSize(total_bytes)
FROM system.tables
WHERE name = 'NYPD_Complaint'
