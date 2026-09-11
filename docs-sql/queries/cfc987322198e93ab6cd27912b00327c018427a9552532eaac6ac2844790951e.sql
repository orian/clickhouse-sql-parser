SELECT
    cutURLParameter('http://bigmir.net/?a=b&c=d&e=f#g', 'a') AS url_without_a,
    cutURLParameter('http://bigmir.net/?a=b&c=d&e=f#g', ['c', 'e']) AS url_without_c_and_e;
