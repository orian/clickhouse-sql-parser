
        INSERT INTO ts_ids
        SELECT (intDiv(number, 20000), reinterpretAsUUID(sipHash128(intDiv(number, 20000))))::Tuple(UInt64, UUID) AS id
        FROM numbers_mt(40000000)
    