/*
Command cwa is a simple command-line wrapper sufficient for use with most simple
tools/utilities linked against the [CommonWA][1] spec.

Usage is simple:

    cwa [options] <file.wasm> [args]

      -gas int
	    number of instructions the VM can perform (default 4194304)
      -jit-enabled
	    enable jit?
      -main-func string
	    main function to call (because rust is broken) (default "cwa_main")
      -min-pages int
	    number of memory pages to open (default is 2 MB) (default 32)
      -test
	    unit testing?
      -vm-stats
	    dump VM statistics?

Example usage:

    $ cwa shaman.wasm

    $ cwa -min-pages 32 -test -vm-stats tests.wasm

[1]: https://github.com/CommonWA/cwa-spec
*/
package main

//go:generate ./generate_readme.sh
