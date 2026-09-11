SELECT toJSONString(tuple(1.25, NULL, NaN, +inf, -inf, [])) SETTINGS output_format_json_quote_denormals = 1;
