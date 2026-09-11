SELECT
toBFloat16(toFloat32(42.7)),
toBFloat16(toFloat32('42.7')),
toBFloat16('42.7')
FORMAT Vertical;
