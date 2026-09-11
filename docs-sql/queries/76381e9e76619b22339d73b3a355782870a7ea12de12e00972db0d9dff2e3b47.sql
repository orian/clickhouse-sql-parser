SELECT name, code, value
FROM system.errors
WHERE value > 0
ORDER BY code ASC
LIMIT 1

┌─name─────────────┬─code─┬─value─┐
│ CANNOT_OPEN_FILE │   76 │     1 │
└──────────────────┴──────┴───────┘
