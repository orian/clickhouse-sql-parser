SELECT arrayMap(x -> hex(x), naiveBayesNgrams('ab', 2, 'codepoint', '0x10FFFE', '0x10FFFF'));
