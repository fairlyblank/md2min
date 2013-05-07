md2min
======

convert markdown file to minimal html file, using github css.

build
-----

    go build md2min templates.go

usage
-----

    md2min ./example/example.md

will generate html without navigator.

    md2min -nav=h2 ./example/example.md

will generate html with "\<h2\>" title navigator.

