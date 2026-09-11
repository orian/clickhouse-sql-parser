SELECT toBFloat16OrZero('0x5E'), -- unsupported arguments
       toBFloat16OrZero('12.3'), -- typical use
       toBFloat16OrZero('12.3456789') -- silent loss of precision
