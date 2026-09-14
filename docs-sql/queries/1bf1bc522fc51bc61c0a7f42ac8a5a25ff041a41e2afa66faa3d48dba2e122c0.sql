WITH colorOKLCHToSRGB((0.6, 0.12, 40)) AS t
SELECT tuple(toUInt8(t.1), toUInt8(t.2), toUInt8(t.3)) AS RGB;
