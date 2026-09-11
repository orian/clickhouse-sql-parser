-- Using 'bech32' variant encodes raw data without a witness version byte,
-- compatible with Cosmos SDK, Injective, Osmosis, and other non-SegWit chains.
SELECT bech32Encode('inj', unhex('751e76e8199196d454941c45d1b3a323f1433bd6'), 'bech32')
