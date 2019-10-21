# The Reverb Programming Language Specification `17/10/19`
## Notation
Most of the specification for Reverb is written in english but lexical elements are written using Extended-Backus-Naur-Form(EBNF).
## Source Code Representation
Reverb source code is Ascii text encoded in UTF-8.
For simplicity, this document will use the unqualified term character to refer to an Ascii code point in the source text.
Each code point is indistinct; for instance, upper and lower case letters are not different characters.
### Characters
```
ascii_letter  = Any ASCII character described as a 'letter' .
letter        = ascii_letter .
decimal_digit = "0" … "9" .
```
## Lexical Elements
### Comments
Comments serve as documentation for Reverb programs. Reverb only supports single line comments denoted by the pound sign '#'.
Comments are ignored by the Reverb interpreter.
### Identifiers
Identifiers in Reverb tag program entities such as variables and types with a name. An identifier is a sequence of one or more letters or digits.
The first character in an identifier must be a letter.

`identifier = letter { letter | decimal_digit } .`
```
foo
bar
varwithdigit69
```
### Keywords
The following words are reserved and may not be used as identifiers. 
```
if          for         begin       return      or
else        hash        end         then        and
echo        var         to          do          
```
### Operators
The following symbols represent all operators in Reverb.
```
+           /           ||          <           !
-           %           ==          >=          ,
*           &&          >           <=          !=
```
### Integer Literals
An integer literal is a series of decimal digits.

`integer_lit = "0" | ( "1" ... "9" ) { decimal_digit }" .`
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