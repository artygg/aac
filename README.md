# Automated Attendance 

## Indentation
Use one tab per indentation

```
def function():
  if condition:
    do_something()
```

## Maximum Line Length
Limit all lines to a maximum of 79 characters. For docstrings or comments, the limit is 72 characters.

## Blank Lines
Top-level function and class definitions along with method definitions inside a class are surrounded by a single blank line.

```
class className:
    
    def method_one(self):
        pass
    
    def method_two(self):
        pass

```

## Variables and Constants
Use camelCase for variables and constants.
Use descriptive names, except for local variables with a very small scope.

```
number1 = 0
numberBiggerThanNumber1 = 5
```

## Functions
Use camelCase for function names.

```
def sumOfNumbers(number1, number2):
  return number1 + number2
```

## Imports:

Imports should usually be on separate lines.
Import modules at the top of the file, after any module comments and docstrings, and before module globals and constants.
Group imports into three categories: standard library imports, related third-party imports, and local application/library-specific imports. Separate these groups with a blank line.

```
import os
import sys

import requests

from mymodule import myfunction
```

## Whitespace in Expressions and Statements

```
array = [1, 2, 3]
sum1 = 1 + 2
sum2 = (2 + 3) - (3 - 1)
```


