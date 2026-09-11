SELECT arrayMap(x -> hex(x), naiveBayesNgrams('xy', 2, 'byte', '0x01', '0xFF'));
