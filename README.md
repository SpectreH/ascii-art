<h1 align="center">ASCII-ART</h1>

## About The Project
Ascii-art is a program which consists in receiving a string as an argument and outputting the string with different modifications in a graphic representation using ASCII. 

## Installation
```
git clone https://github.com/SpectreH/ascii-art.git
cd ascii-art
```
## Usage
```
go run . [STRING] [BANNER] [OPTION]
```
Available banners:
* <code>standard</code>
* <code>shadow</code>
* <code>thinkertoy</code>

Available options:
* <code>--output</code> - Save ASCII-art into file
* <code>--align</code> - Apply align to ASCII-art
* <code>--color</code> - Set ASCII-art color
* <code>--reverse</code> - Transform ASCII-art to string

## Examples
```
go run . "hello" standard
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
                               
go run . "Hello There!" shadow
                                                                                      
_|    _|          _| _|                _|_|_|_|_| _|                                  
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|       
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| 
                                                                                      
                                                                                      
                                                                                      
go run . "hello" standard --align=center                                                                                   
|                                             _                _    _                                                       |
|                                            | |              | |  | |                                                      |
|                                            | |__      ___   | |  | |    ___                                               |
|                                            |  _ \    / _ \  | |  | |   / _ \                                              |
|                                            | | | |  |  __/  | |  | |  | (_) |                                             |
|                                            |_| |_|   \___|  |_|  |_|   \___/                                              |
|                                                                                                                           |
|                                                                                                                           |

go run . "how are you" shadow --align=justify                                                                            
|                                                                                                                           |
|_|                                                                                                                         |
|_|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|                    _|    _|   _|_|   _|    _| |
|_|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|                  _|    _| _|    _| _|    _| |
|_|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|                        _|    _| _|    _| _|    _| |
|_|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|                    _|_|_|   _|_|     _|_|_| |
|                                                                                                      _|                   |
|                                                                                                  _|_|                     |
|                                                                                                                           |

go run . "hello" standard --output=banner.txt
cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $

go run . --reverse=banner.txt
hello
```
## Additional information

Only standard go packages were in use. In <code>tests</code> folder you can find test files for each flag.

To run test use:
```
bash tests\<filename>.sh
```

## Author

* SpectreH (https://github.com/SpectreH)

## License

This software is released under the MIT License, see LICENSE.
