# go-files

## Description

A collection of libraries designed to assist in handling my file basic file IO needs
The code is written with Test Driven Development principles in mind, however not all functions have tests as of yet.

## Problems

I initially made this repository when I felt a lot less comfortable using the golang utils, and I think some of the functionality I have written is still helpful.
Overtime however, I have realised a couple things:

### Enjoyment

Initially I rather enjoyed writing code that I expected to already be somewhere else,
because ultimately I was learning how to use the libraries I was interacting with.

I made this project because I didn't feel comfortable using the base libraries,
and it solved that problem.

Since now I am more comfortable using this, I just want to get my projects done.

#### Testing

Testing is something I put a decent amount of effort into here.
It was good practice for writing unit tests. Testing is slow,
yet incredibly satisfying to see a collection of tests ALL PASS.

### Abstraction

I don't think a lot of the stuff I did was valuable. I still use some of it,
and those sections of the library I think are okay.

### Pattern Recognition

I have started to find some of the exact same functions that I have written,
already existing in libraries I often use. This is a worrying sign,
but also exactly what was intentional at the start.

I understand why it is happening, however my implementation is not better,
and their implementation is simpler to use, so I am gonna switch to using those versions.

#### Example: Filepath

I noticed that `filepath` has a `Join` function, and various functions that split filepaths.
I spent a decent amount of time trying
to solve the same problem, and felt burdened by how I implemented it,
whilst the `filepath` implementation I would rather work with.

<!-- ### Limited Support -->
<!---->
<!-- As one does, one has aspirations to what a codebase would do. -->
<!-- As aspirations go, they are mighty and take a lot of effort to implement, -->
<!-- and easy to forget the next day and wonder: -->
<!---->
<!-- > Why was I trying to do it like this? -->
<!-- > What was I doing? My task, or tangential task that I didn't make note of? -->
<!---->
<!-- And as I have many things I have to work with, that ends up with unfinished features here and there cluttering.  -->
<!-- I should have used branches more actively earlier on in the project.  -->

The code I think is valuable is:

- File (mostly)
- TextFile
- CSV (doesn't fulfill the standard sadly due to not handling newlines with how the ISO specs say)

This also happens to be code I have used on multiple occasions.

Code that I think could be helpful is expanding on stuff like this, such as for sqlite.
That isn't functional however, so it's existance is questionable.

At some point, I need to go through this codebase and clean it up a bit.
It serves value even if all it does it collect and abstracts a bunch of the frequently used libraries I use.

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
