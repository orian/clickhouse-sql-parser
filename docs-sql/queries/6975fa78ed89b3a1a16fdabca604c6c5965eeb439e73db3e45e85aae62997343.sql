-- While 'bc' (Mainnet) and 'tb' (Testnet) are the only allowed hrp values for the
-- SegWit address format, Bech32 allows any hrp that satisfies the above requirements.
SELECT bech32Encode('abcdefg', unhex('751e76e8199196d454941c45d1b3a323f1433bd6'), 10)
