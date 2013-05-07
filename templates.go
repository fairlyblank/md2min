package main

var templContent = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="chrome=1">
    <title></title>

    <link rel="stylesheet" href="./markdown.css">
	</head>
	<body>
	<div class="markdown-body">{{.Content}}</div>
	</body>
</html>
`