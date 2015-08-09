
# sgfinfo

Steps toward a command-line query tool for SGF Go game files.

This is still rudimentary.

to build:
```
$ go build
```

usage:
```
$ sgfinfo <sgf_file>
```
...which will dump details, in json, about the given game.

Parsing still needs to flesh out some of the darker corners of the SGF spec, but most use-cases are covered including zip-file archives, multiple games per file and gametree variations.

The parser is based on an excellent talk from Rob Pike: ["Lexical Scanning in Go"](https://www.youtube.com/watch?v=HxaD_trXwRE)

## Resources

* [The SGF File Format](http://www.red-bean.com/sgf)
* [Other Game Formats](http://senseis.xmp.net/?FileFormat) (from Sensei's Library)
