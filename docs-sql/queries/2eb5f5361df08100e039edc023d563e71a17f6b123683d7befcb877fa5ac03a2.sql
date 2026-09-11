  taxi=# \timing
  Timing is on.
  taxi=# SELECT round(avg(tip_amount), 2) FROM taxi.trips;
   round
  -------
    1.68
  (1 行)

  Time: 9.438 ms
