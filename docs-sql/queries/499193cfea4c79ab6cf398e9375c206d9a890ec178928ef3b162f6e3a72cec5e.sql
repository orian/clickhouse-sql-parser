SELECT COLUMNS('.*_amount|fee|tax') APPLY(avg) APPLY(round)
FROM nyc_taxi.trips;
