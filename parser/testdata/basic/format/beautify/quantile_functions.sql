-- Origin SQL:
SELECT quantile(0.9)(x), quantiles(0.5, 0.9)(x);


-- Beautify SQL:
SELECT
  quantile(0.9)(x),
  quantiles(0.5, 0.9)(x);
