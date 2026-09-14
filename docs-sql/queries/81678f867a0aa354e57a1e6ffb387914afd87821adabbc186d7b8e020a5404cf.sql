SELECT
    name,
    toTypeName(name),
    length(name),
    empty(name)
FROM FixedStringTable;
