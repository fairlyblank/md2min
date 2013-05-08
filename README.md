md2min
======

convert markdown file to minimal html file, using github css.

build
-----

    go build md2min.go templates.go

usage
-----

	Usage: md2min [-nav=h2] name.md
	  name.md: markdown file name
	  -nav="none": navigate level ["none", "h1", "h2", "h3", "h4", "h5", "h6"]

example
-------

	md2min ./example/example.md

will generate html without navigator.

	md2min -nav=h2 ./example/example.md

will generate html with "\<h2\>" title navigator.

