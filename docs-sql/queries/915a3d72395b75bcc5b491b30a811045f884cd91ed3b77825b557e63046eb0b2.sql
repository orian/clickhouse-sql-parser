Output: sum(number)

Limit (preliminary LIMIT)
│  Limit 1
│  Offset 0
└──Aggregating
   │  Keys: number MOD 4
   │  Aggregates: sum(number)
   │  Skip merging: 0
   └──ReadFromSystemNumbers
         Output: number
