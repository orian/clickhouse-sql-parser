SELECT avg2(toUInt8(3), 1.0) AS result, toTypeName(result) AS type;
-- El tipo devuelto es Float64 ya que el UInt8 debe promoverse a 64 bits para la comparación.
