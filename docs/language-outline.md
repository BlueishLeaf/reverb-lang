# The Reverb Programming Language
This document contains an outline of the Reverb programming language.
If you are interested in viewing the language reference, see the 'language-spec' document in the 'docs' directory.
If want to see some code examples, see the 'code-examples' directory in the 'docs' directory. Note: These examples are
liable to change during development.
## About
Reverb is a dynamically-typed interpreted programming language written in Go. Unlike most conventional languages,
the output of a Reverb program is an audio file, rather than an executable. The goal of Reverb as a language is to allow
one to create music in a nuanced and novel way using code with an easy to learn syntax.
## Features
- Easily-Legible user-friendly syntax.
- Case insensitive.
- First-Class functions.
- Truthiness.
- Familiar data structures (integers, booleans, arrays, hashlists, etc).
- Conditionals and loops.
- Synthesizer.
- A standard library supporting a variety of different instruments and chords.
## Use Cases
- Educational tech (easy to learn and provides novel feedback)
- Assistive tech (provide aural feedback to visually-impaired users)
- General music composition (duh...)
- Game development (dynamic music composition for soundtracks or sfx)
## TODO
- Loops (For and While)
- Improve syntax for if statements (reduce number of 'end's required)
- Improve syntax for functions (call them something other than 'echoes')
- Comments (python-like comments '#')