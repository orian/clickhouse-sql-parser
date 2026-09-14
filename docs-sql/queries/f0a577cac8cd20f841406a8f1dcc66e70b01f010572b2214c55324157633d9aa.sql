Output: sum(number)

Union
├──Aggregating
│  │  Keys:
│  │  Aggregates: sum(number)
│  │  Skip merging: 0
│  └──ReadFromSystemNumbers
│        Output: number
└──Sorting (Sorting for ORDER BY)
   │  Sort description: sum(number) ASC
   └──Aggregating
      │  Keys:
      │  Aggregates: sum(number)
      │  Skip merging: 0
      └──ReadFromSystemNumbers
            Output: number
