SELECT * FROM format(
    'RowBinaryWithNamesAndTypesAndDefaults',
    'x Nullable(UInt32) DEFAULT 42',
    unhex('01' || '0178' || '10' || hex('Nullable(UInt32)') || '01')
);
