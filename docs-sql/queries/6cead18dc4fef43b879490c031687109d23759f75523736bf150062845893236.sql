CREATE TABLE visits (is_young UInt8, is_female UInt8, is_mobile UInt8) ENGINE = Memory;

-- 80 of the 100 young visitors browse on a mobile device, and only 20 of the 100 older ones do,
-- while the sex of a visitor says nothing about the device.
INSERT INTO visits SELECT 1, number % 2, number < 80 FROM numbers(100);
INSERT INTO visits SELECT 0, number % 2, number < 20 FROM numbers(100);

SELECT round(categoricalInformationValue(is_young, is_mobile)[1], 4) AS iv FROM visits;
