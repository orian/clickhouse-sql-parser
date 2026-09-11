SELECT extractTextFromHTML('
<html>
    <head><title>Page Title</title></head>
    <body>
        <p>Hello <b>World</b>!</p>
        <script>alert("test");</script>
        <!-- comment -->
    </body>
</html>
');
