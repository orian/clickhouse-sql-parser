-- A witness version of 0 will result in a different address string.
SELECT bech32Encode('bc', unhex('751e76e8199196d454941c45d1b3a323f1433bd6'), 0)
