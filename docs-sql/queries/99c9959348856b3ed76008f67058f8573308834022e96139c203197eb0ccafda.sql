SELECT hex(reinterpret([toUInt8(1), toUInt8(2), toUInt8(255)]::Array(UInt8), 'String')) AS array_of_UInt8_to_string
