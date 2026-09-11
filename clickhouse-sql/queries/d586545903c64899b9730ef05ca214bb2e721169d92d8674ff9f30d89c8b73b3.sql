
        CREATE TABLE bigint ( u128 UInt128, i128 Int128, u256 UInt256, i256 Int256) ENGINE = Memory
        AS
        SELECT * FROM generateRandom('u128 UInt128, i128 Int128, u256 UInt256, i256 Int256', 42) LIMIT 50000;
    