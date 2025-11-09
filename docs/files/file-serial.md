# Basic rules on Serialisation of files and contracts in this library

Rough definitions

- Serialisation: Converting from an internal datastructure to a file storable format
- eg you serialise game state data into JSONs, or the like

- Deserialisation: Converting from a file storable format to an internal datastructure
- eg you deserialise JSONs into game state data

To me, I am using them to go up and down a chain of conversion and manage each struct's datastructure in a way I have control over.
EG:
Files are bytes stored on the hard hard disk
Files are bytes in the File struct stored in memory
TextFiles are string versions of the bytes from the File struct
CSVFiles are tables based off of rows in a text file containing comma separated values

Each time I go down this chain, I am converting from a more basic data format to a more complex, and more accessibly dataformat (for my needs)
EG
The first   round of serialisation, converting bytes on DISK to bytes in MEM      from DISK to MEM
The second  round of serialisation, converting bytes to strings                   from File to TextFile
The third   round of serialisation, converting strings to Records in a Table      from TextFile to CSVFile

And it has it's converse:
The first   round of deserialisation, converting Records in a table to Strings    from CSVFile to TextFile
The second  round of deserialisation, converting Strings to Bytes                 from TextFile to File
The third   round of deserialisation, converting bytes (MEM) to bytes (DISK)      from MEM to DISK

So my definition of serialisation goes more like the following:

- Converting from one storage format and medium to another, for the purpose of being able to manipulate it easier in the new format.

## Specification of the above system

I want to be specific with how I deal with this.
One contract I have is the following:

- Files have a ReadContents and a WriteContents
- ReadContents calls the 'superclass' copy of ReadContents, going up the chain.
- ReadContents then uses the return value to call deserialise, in order to process the data into a useable format

- WriteContents calls serialise with the current buffer
- WriteContents then calls the 'superclass' copy of WriteContents, going down the chain, using the output of the serialise call
