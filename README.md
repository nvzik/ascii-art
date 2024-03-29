

# Ascii-art
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.
What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters.
- This project will handle an input with numbers, letters, spaces, special characters and \n.
# Details
- This project was written in Go and respects the good practices.
# Banner Format
- Each character has a height of 8 lines.
- Characters are separated by a new line \n.
- Here is an example of ' ', '!' and '"'(one dot represents one space) :
``` bash
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......
```

  And we can choose fonts(standard, thinkertoy, shadow) for our output.

- This project will handle an input with numbers, letters, spaces, special characters and \n.
# Details
- This project was written in Go and respects the good practices.
- The usage must respect this format go run . [STRING] [BANNER], any other formats must return the following usage message:
``` bash
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```
Additionally, the program must still be able to run with a single [STRING] argument, with standard font.
## # Usage
``` bash
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $

$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $

$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```
