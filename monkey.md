# Monkey Language Overview
This document shows what Monkey looks like and highlights its core features through examples.

## Variables and Expressions
Monkey supports variable bindings using `let` and basic arithmetic expressions.
```
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

## Arrays and Hash Maps
Monkey includes built-in support for arrays and hash maps (key–value pairs).
```
let myArray = [1, 2, 3, 4, 5];
let sasi = {"name": "SaSi", "age": 28};
```
Elements can be accessed using index or key notation.
```
myArray[0] // => 1
sasi["name"] // => "SaSi"
```

## Functions
Functions are first-class values in Monkey and are defined using the fn keyword.
```
let add = fn(a, b) { return a + b; };
\\ optionally:
let add = fn(a, b) { a + b; };
```
Functions are called using standard call syntax.
```
add(1, 2)
```

## Conditionals and Recursion
Monkey supports conditional expressions with `if` and `else`, which evaluate to values. Below is an example of fibonacci function written in monkey.
```
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
```