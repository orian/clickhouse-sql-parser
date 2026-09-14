
        SELECT count()
        FROM date_preimage_perf
        WHERE seen >= toDateTime('2026-03-02 00:00:00', 'UTC')
          AND seen < toDateTime('2026-03-03 00:00:00', 'UTC')
    