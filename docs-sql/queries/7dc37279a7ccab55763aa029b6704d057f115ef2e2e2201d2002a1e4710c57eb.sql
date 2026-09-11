SELECT COLUMNS('.*_amount|fee|tax') APPLY(avg) APPLY(x -> round(x, 2))
FROM nyc_taxi.trips;
