# The Reverb Programming Language Specification `14/10/19`
## Notation
Most of the specification for Reverb is written in english but lexical elements are written using Extended-Backus-Naur-Form(EBNF).
## Source Code Representation
Reverb source code is Ascii text encoded in UTF-8.
For simplicity, this document will use the unqualified term character to refer to an Ascii code point in the source text.
Each code point is indistinct; for instance, upper and lower case letters are not different characters.
### Characters
```
letter        = ascii_letter .
decimal_digit = "0" … "9" .
```
## Lexical Elements
### Comments
### Identifiers
### Keywords
### Operators and Punctuation
### Integer Literals
## Variables
## Data Structures
### Integers
### Booleans
### Arrays
### Dictionaries
## Declarations
### Variables
### Functions
## Expressions
### Index Expression
### Function Literals
### Function Calls
### Operators
#### Operator Precedence
#### Arithmetic Operators
#### Comparison Operators
#### Logical Operators
## Statements
### Return Statements
### If Statements
### For Statements
## Built-In Functions
### Length
### Play Sounds
### Overlay Sounds
### Instruments
#### Piano
#### Sine Wave
## Program Initialization and Execution