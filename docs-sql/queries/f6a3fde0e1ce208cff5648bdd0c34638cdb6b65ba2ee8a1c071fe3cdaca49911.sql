arthur :) select extractKeyValuePairs('name:neymar, age:31 team:psg,nationality:brazil') as kv

SELECT extractKeyValuePairs('name:neymar, age:31 team:psg,nationality:brazil') as kv

Query id: f9e0ca6f-3178-4ee2-aa2c-a5517abb9cee

┌─kv──────────────────────────────────────────────────────────────────────┐
│ {'name':'neymar','age':'31','team':'psg','nationality':'brazil'}        │
└─────────────────────────────────────────────────────────────────────────┘
