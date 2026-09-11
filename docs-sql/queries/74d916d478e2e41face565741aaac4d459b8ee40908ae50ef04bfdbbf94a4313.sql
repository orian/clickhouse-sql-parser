ALTER TABLE uk.uk_price_paid_with_projections
  (ADD PROJECTION prj_obj_town_price
  (
    SELECT *
    ORDER BY
        town,
        price
  ))
