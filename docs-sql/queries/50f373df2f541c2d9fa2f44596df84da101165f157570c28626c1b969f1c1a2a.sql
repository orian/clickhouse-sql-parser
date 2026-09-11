    SELECT
        mutation_id,
        command,
        is_done
    FROM system.mutations
    WHERE table = 'uk_price_paid'
    ORDER BY create_time DESC
    LIMIT 5;
