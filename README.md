a - text block-indentation
======

Micro-utility for use in acme to block-indent or un-indent lines of text.  Cribbed from Russ Cox's utility of the same name.  Written in Go.

Usage:

Pipe in data, calling the program as either `a` or `a+` to add one tab character to the front of each linepassed.  Outputs to stdout.

Pipe to `a-` to remove one tab character from the front of each line instead, if there is one.

Provided as public domain software.
