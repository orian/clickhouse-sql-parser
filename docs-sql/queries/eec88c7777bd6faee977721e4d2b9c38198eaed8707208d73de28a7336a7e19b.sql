CREATE TABLE reviews (body String) ENGINE = Memory;
INSERT INTO reviews VALUES ('The package arrived three days late.');
SELECT * FROM reviews WHERE aiFilter(body, 'the customer is angry about shipping')
