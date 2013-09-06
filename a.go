/*
 * a
 *
 * Micro-utility for use in acme to block-indent or un-indent
 * lines of text.  Cribbed from Russ Cox's utility of the same name.
 *
 * Public Domain
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addIndent(r *bufio.Reader, w *bufio.Writer) {
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			fmt.Fprintf(w, "\t%s", l)
			break
		}
		fmt.Fprintf(w, "\t%s", l)
	}
}

func delIndent(r *bufio.Reader, w *bufio.Writer) {
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			fmt.Fprintf(w, "%s", strings.TrimPrefix(l, "\t"))
			break
		}
		fmt.Fprintf(w, "%s", strings.TrimPrefix(l, "\t"))
	}
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	if stdin == nil {
		return
	}

	stdout := bufio.NewWriter(os.Stdout)

	exec := strings.TrimLeft(os.Args[0], "./")

	if exec == "a+" || exec == "a" {
		addIndent(stdin, stdout)
	} else if exec == "a-" {
		delIndent(stdin, stdout)
	}

	stdout.Flush()

	return
}
