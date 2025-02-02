# Monkey-Script-WebREPL
A WebREPL Interface for the MonkeyScript interpreter made with Go | HTMX

# Project Extention Ideas
- instead of using Bytes to read chars, use Runes to enable emojis etc as valid tokens
- add postfix operator support
- implement block functions as inline lambdas
- implement incrementers and decrementers
- character escaping \n \t ...
- multiply a string by an integer to repeat it
- turn into a web-based CLI hosted on Azure (go | htmx | sqlite3)

# Supported Data Types
- Integers
- Booleans
- Strings
- Arrays
- Hashmaps
- Functions

# Language Features
- Functions
- Variables
- Functions as Variables
- Closures
- Many Type Arrays
- Anonymous functions

# Examples
- Print Out

    ```
    puts("Welcome to Monkey!");      // Welcome to Monkey! 
    puts(meaning_of_life());        // 42
    ```

- Variables

    ```
    let a = "Hello, World!";
    let b = 1;
    ```

- Strings

    ```
    let hello = "Hello";
    let world = "World";
    let saying = hello + ", " + world + "!";      // Hello, World!
    ```

- Arrays

    ```
    let a = [1, 2, 3, 4];
    let b = push(a, 5);
    b                       // [1, 2, 3, 4, 5]

    // Available builtins
    len(a)                  // 4
    first(a)                // 1
    last(a)                 // 4
    rest(a)                 // [2, 3, 4]
    push(a, 5)              // [1, 2, 3, 4, 5]
    ```

- Hashmaps

    ```
    let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
    people[0]["name"];      // Alice
    ```

- Functions as Variabls

    ```
    let getName = fn(person) { person["name"]; };
    getName(people[0]);     // Alice

    let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } };
    let heythere = makeGreeter("Hey there");
    heythere("Monkey");     // Hey there Monkey!
    ```

Done 😊
