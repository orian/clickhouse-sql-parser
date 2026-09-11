ALTER TABLE uk.uk_price_paid_with_projections
  (ADD PROJECTION prj_gby_county
  (
    SELECT
        county,
        avg(price)
    GROUP BY county
  ))
