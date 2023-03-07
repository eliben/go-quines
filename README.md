# go-quines

Some quines in the Go programming language

* `quine-str-regular.go`: A quine I came up with after reading some Python code
  that does something similar. It's very simple, without using any special
  ASCII codes for quotes, but it does require a lot of the code on one long
  line.
* `quine-str-rsc.go`: A quine crafted by Russ Cox and show in a
  [2010 blog post](https://research.swtch.com/zip). This one is using the
  Go backtick quote for nicer code formatting. It also demonstrates a common
  technique used in quines to make sure quoting works out, by passing an ASCII
  code for quotes to the printing function.
* `quine-source-embed.go`: A bit of a "cheat". Reading your own source code and
  spitting it out is an obvious way to write quines, but it's so easy that it's
  usually frowned upon. "Morally pure" quines are not supposed to issue file
  reading calls; this one doesn't -- it uses the embedded file support added
  in Go 1.16 to make the toolchain read the file for us.
