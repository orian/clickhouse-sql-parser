arthur :) select extractKeyValuePairsWithEscaping('age:a\\x0A\\n\\0') as kv

SELECT extractKeyValuePairsWithEscaping('age:a\\x0A\\n\\0') AS kv

Query id: 44c114f0-5658-4c75-ab87-4574de3a1645

┌─kv────────────────┐
│ {'age':'a\n\n\0'} │
└───────────────────┘
