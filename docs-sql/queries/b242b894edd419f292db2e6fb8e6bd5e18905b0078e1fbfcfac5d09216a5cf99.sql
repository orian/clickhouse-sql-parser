SELECT regionToName(number::UInt32, 'en'), regionToCountry(number::UInt32) AS id, regionToName(id, 'en') FROM numbers(13);
