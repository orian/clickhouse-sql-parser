-- Consider a list of contacts that may specify multiple ways to contact a customer.

CREATE TABLE aBook
(
    name String,
    mail Nullable(String),
    phone Nullable(String),
    telegram Nullable(UInt32)
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO aBook VALUES ('client 1', NULL, '123-45-67', 123), ('client 2', NULL, NULL, NULL);

-- The mail and phone fields are of type String, but the telegram field is UInt32 so it needs to be converted to String.

-- Get the first available contact method for the customer from the contact list

SELECT name, coalesce(mail, phone, CAST(telegram,'Nullable(String)')) FROM aBook;
