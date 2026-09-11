ALTER TABLE uk.uk_price_paid_with_projections_v2
    MATERIALIZE PROJECTION projection_by_year_district_town
SETTINGS mutations_sync = 1
