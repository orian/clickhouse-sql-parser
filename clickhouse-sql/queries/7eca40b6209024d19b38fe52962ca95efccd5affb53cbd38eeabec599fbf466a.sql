
        SELECT max(timeSeriesIdToGroup(id))
        FROM ts_ids
        WHERE id IN (
            SELECT timeSeriesStoreTags((number, reinterpretAsUUID(sipHash128(number)))::Tuple(UInt64, UUID), [('n', toString(number))], '__name__', 'm')
            FROM numbers(2000)
        )
    