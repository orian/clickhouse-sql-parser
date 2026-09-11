SET output_format_json_array_of_rows = 1;
SELECT number FROM numbers(3) FORMAT JSONEachRow;
