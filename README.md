md2min
======

convert markdown file to minimal html file, using github css.

install
-------

	go get https://github.com/fairlyblank/md2min

If you want to build standalone execution:

	cd $(GOPATH)/github.com/fairlyblank/md2min/main
	go build -o md2min

usage
-----

As standalone execution:

	Usage: md2min [-nav=h2] name.md
	  name.md: markdown file name
	  -nav="none": navigate level ["none", "h1", "h2", "h3", "h4", "h5", "h6"]

As package, please review [main.go](https://github.com/fairlyblank/md2min/main/main.go).Generally:

### 1. Prepare input byte slice

	bytes, _ := ioutil.ReadAll(filename)

### 2. Prepare output writer

	wr, _ := os.Create(newname)

### 3. Initialize package

Without navigator:

	md := md2min.New("none")

Using "\<h2\>" tag as navigator:

	md := md2min.New("h2")

### 4. Parse

	md.Parse(bytes, wr)

example
-------

	md2min ./example/example.md

will generate html without navigator.

	md2min -nav=h2 ./example/example.md

will generate html with "\<h2\>" title navigator.

