Formatting bug when commenting fluid call chains

I have code that per `go fmt` should be commented like:

```golang
	b.
		with("alalalalalalal").                    // good explanatory comment
		with("bbababababababababbababab").         // less good explanatory comment
		with("cbcbcbbcbcbcbcbcbcbcbcbcbcbbccbcbc") // fantastic information
```

when I save the file or hit `Reformat Code` I see a flash in the editor of the comments changing to:

```golang
	b.
		with("alalalalalalal"). // good explanatory comment
		with("bbababababababababbababab"). // less good explanatory comment
		with("cbcbcbbcbcbcbcbcbcbcbcbcbcbbccbcbc") // fantastic information
```

and then back to the expected, except the new correctly formatted file is only showing in Goland/IntelliJ and not on disk.

I have [recorded my screen] to try and showcase the behavior while using `git` and `cat` to 
show the current status of the file and to show the difference. This repo has the "incorrectly formatted" version 
committed to make it easier to reproduce and show that the editor and the file system are not in sync.

Inline comments without the fluid calls (i.e. the multiple method calls across multiple lines) don't show a flash when
reformatting and consistently formats the code correctly.

[recorded my screen]: ./go-multi-line-call-comment-formatting.mov
