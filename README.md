# go-files

A collection of libraries designed to assist in handling my file basic file IO needs
The code is written with Test Driven Development principles in mind, however not all functions have tests as of yet.

## files

The files library is a collection of structs and functions designed to make it simpler to read and write to files.
Currently implemented with UNIX systems in mind. Functionality for Windows is in a work in progress state as I have no need to use it in a Windows environment currently.

NOTE:

- Pretty much all of the file implementations are buffered, and so require an extra signal to actually Write to the file. Usually this is WriteContents.
- The files are compatible with io.ReadWriter, however be careful with using the more complex file structs.

## formatting

Very Basic formatters, being rewritten / phased out.

## table

An implementation of a table. Use as you wish.
Will likely be moved out into a separate library in the near future.
Things to note:

- use \"\" as the default value if the library wants you to set a value and you don't have anything specific in mind, as that is considered to be an empty Cell.
