WITH colorOKLABToSRGB((0.7, 0.1, 0.54)) AS t
SELECT tuple(toUInt8(t.1), toUInt8(t.2), toUInt8(t.3)) AS RGB;
