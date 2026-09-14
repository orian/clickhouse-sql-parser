arthur :) select extractKeyValuePairs('age:a\\x0A\\n\\0') as kv

SELECT extractKeyValuePairs('age:a\\x0A\\n\\0') AS kv

Query id: e9fd26ee-b41f-4a11-b17f-25af6fd5d356

┌─kv─────────────────────┐
│ {'age':'a\\x0A\\n\\0'} │
└────────────────────────┘
