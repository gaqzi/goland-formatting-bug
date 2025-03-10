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

## Versions

IntelliJ IDEA 2024.3.4.1 (Ultimate Edition)
Build #IU-243.25659.59, built on March 5, 2025
Licensed to Bjorn Andersson
Subscription is active until June 10, 2025.
Runtime version: 21.0.6+8-b631.39 aarch64 (JCEF 122.1.9)
VM: OpenJDK 64-Bit Server VM by JetBrains s.r.o.
Toolkit: sun.lwawt.macosx.LWCToolkit
macOS 14.7.3
GC: G1 Young Generation, G1 Concurrent GC, G1 Old Generation
Memory: 4096M
Cores: 11
Metal Rendering is ON
Registry:
ide.experimental.ui=true
Non-Bundled Plugins:
org.intellij.plugins.hcl (243.25659.42)
com.google.tools.ij.aiplugin (1.11.0-243)
PythonCore (243.24978.46)
name.kropp.intellij.makefile (243.23654.19)
org.jetbrains.plugins.go (243.24978.46)
Pythonid (243.25659.59)
com.intellij.lang.jsgraphql (243.25659.42)
net.ashald.envfile (3.4.2)
com.google.gct.core (24.11.1-233-api-version-223)
Kotlin: 243.25659.59-IJ

