rmake
=====

Recursively find parent directory's Makefile, and execute make.

	$ tree /path
	/path
	├── Makefile
	└── to
	    └── [current-directory]

	2 directories, 1 file
	$ cat /path/Makefile
	all:
		echo Hello
	$ rmake
	echo Hello
	Hello

Usage
------

	$ rmake [OPTION]... [MAKE-ARGS]...

	Options:
		--help       show this help message
		--version    print the version

License
--------

MIT License

Author
-------

wara <kusabashira227@gmail.com>
