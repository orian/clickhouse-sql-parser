SELECT
    '<tag>Hello & "World"</tag>' AS original,
    encodeXMLComponent('<tag>Hello & "World"</tag>') AS xml_encoded;
