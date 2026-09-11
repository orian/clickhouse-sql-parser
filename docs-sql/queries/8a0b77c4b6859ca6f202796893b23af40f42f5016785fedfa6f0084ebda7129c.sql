-- The age of a visitor predicts the device, the sex of a visitor does not.
SELECT arrayMap(x -> round(x, 4), categoricalInformationValue(is_young, is_female, is_mobile)) AS iv
FROM visits;
