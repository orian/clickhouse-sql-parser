
        INSERT INTO bitmap_cardinality_operands
        SELECT 'small', groupBitmapState(x)
        FROM (SELECT arrayJoin([toUInt32(7), toUInt32(1000000), toUInt32(2999999)]) AS x)
    