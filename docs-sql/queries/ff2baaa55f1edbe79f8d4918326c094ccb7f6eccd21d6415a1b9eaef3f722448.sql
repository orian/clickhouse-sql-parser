SELECT
    toUUIDOrZero('550e8400-e29b-41d4-a716-446655440000') AS valid_uuid,
    toUUIDOrZero('invalid-uuid') AS invalid_uuid
