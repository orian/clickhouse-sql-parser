SELECT prime
FROM primes()
WHERE prime >= 1e15
  AND prime % 65537 = 1
LIMIT 1;
