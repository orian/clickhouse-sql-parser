SELECT *
FROM places.`datasets.places_os`
WHERE address IS NOT NULL AND postcode IS NOT NULL AND instagram IS NOT NULL
LIMIT 1;
