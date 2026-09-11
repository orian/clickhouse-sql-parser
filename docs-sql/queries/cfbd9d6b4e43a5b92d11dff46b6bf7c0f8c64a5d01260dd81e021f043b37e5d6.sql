SELECT midpoint(1, toUInt8(3), 0.5) AS result, toTypeName(result) AS type;
-- El tipo devuelto es Float64 ya que el UInt8 debe promoverse a 64 bits para la comparación.
