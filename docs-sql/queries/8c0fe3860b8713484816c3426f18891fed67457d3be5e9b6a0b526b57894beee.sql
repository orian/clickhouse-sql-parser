  SELECT sum(my_float), sum(my_decimal) FROM float_vs_decimal;

  ┌──────sum(my_float)─┬─sum(my_decimal)─┐
  │ 499693.60500000004 │      499693.605 │
  └────────────────────┴─────────────────┘

  SELECT sumKahan(my_float), sumKahan(my_decimal) FROM float_vs_decimal;

  ┌─sumKahan(my_float)─┬─sumKahan(my_decimal)─┐
  │         499693.605 │           499693.605 │
  └────────────────────┴──────────────────────┘
