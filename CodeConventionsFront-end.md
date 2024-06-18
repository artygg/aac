### HTML and CSS Code Conventions

#### HTML Conventions

1. **Doctype Declaration:**
   ```html
   <!DOCTYPE html>
   ```

2. **Character Encoding:**
   ```html
   <meta charset="UTF-8">
   ```

3. **HTML Structure:**
   ```html
   <!DOCTYPE html>
   <html lang="en">
   <head>
       <meta charset="UTF-8">
       <title>Document</title>
       <link rel="stylesheet" href="styles.css">
   </head>
   <body>
       <!-- Content here -->
   </body>
   </html>
   ```

4. **Indentation:**
   - Use 4 spaces for indentation.

5. **Lowercase Tag and Attribute Names:**
   ```html
   <div class="container">
       <img src="image.jpg" alt="Description">
   </div>
   ```

6. **Attribute Quotation:**
   - Always use double quotes for attribute values.
   ```html
   <input type="text" name="username">
   ```

7. **Self-closing Tags:**
   - Include the closing slash in self-closing tags.
   ```html
   <img src="image.jpg" alt="Description" />
   ```

8. **HTML Comments:**
   ```html
   <!-- This is a comment -->
   ```


#### CSS Conventions

1. **Indentation:**
   - Use 4 spaces for indentation.

2. **Selectors:**
   - Keep them short and readable.
   - Avoid overly specific selectors to maintain flexibility.
   ```css
   /* Good */
   .navbar {
       background-color: #333;
   }

   /* Bad */
   div.header .navbar ul li a {
       color: white;
   }
   ```

3. **Block Formatting:**
   ```css
   .class-name {
       property: value;
       property: value;
   }
   ```

4. **Use Hyphens for Class and ID Names:**
   ```css
   .main-container {
       /* styles */
   }
   ```

5. **Avoid Inline Styles:**
   - Write all styles in a separate CSS file or within `<style>` tags in the `<head>` section.
   ```html
   <style>
       .example {
           color: red;
       }
   </style>
   ```

6. **CSS Comments:**
   ```css
   /* This is a comment */
   ```

7. **Use of Hexadecimal, RGB, or HSL for Colors:**
   ```css
   .color-sample {
       color: #ff5733; /* Hexadecimal */
       background-color: rgb(255, 87, 51); /* RGB */
       border-color: hsl(9, 100%, 60%); /* HSL */
   }
   ```

8. **Order of Properties:**
   - Group related properties together.
   - Follow a logical order such as positioning, box model, typography, visual.
   ```css
   .box {
       /* Positioning */
       position: relative;
       top: 10px;
       
       /* Box model */
       margin: 10px;
       padding: 20px;
       border: 1px solid #000;
       
       /* Typography */
       font-family: Arial, sans-serif;
       font-size: 16px;
       
       /* Visual */
       background-color: #fff;
       color: #333;
   }
   ```

9. **Use Shorthand Properties Where Possible:**
   ```css
   .box {
       margin: 10px 20px 30px 40px; /* top, right, bottom, left */
   }
   ```


### JavaScript Code Conventions

1. **Indentation:**
   - Use 2 or 4 spaces for indentation (choose one and be consistent).
   - Do not use tabs.
   ```javascript
   function example() {
       if (true) {
           console.log('Hello, World!');
       }
   }
   ```

2. **Variable Declarations:**
   - Use `const` for constants, `let` for variables that will change.
   - Avoid using `var`.
   ```javascript
   const MAX_USERS = 100;
   let count = 0;
   ```

3. **Semicolons:**
   - Always use semicolons to terminate statements.
   ```javascript
   let x = 5;
   let y = 10;
   console.log(x + y);
   ```

4. **Naming Conventions:**
   - Use camelCase for variables and functions.
   - Use PascalCase for classes and constructor functions.
   ```javascript
   let userName = 'John';
   function getUserName() {
       return userName;
   }
   class User {
       constructor(name) {
           this.name = name;
       }
   }
   ```

5. **Function Declarations:**
   - Prefer function expressions or arrow functions over function declarations.
   ```javascript
   // Function expression
   const add = function(a, b) {
       return a + b;
   };

   // Arrow function
   const multiply = (a, b) => a * b;
   ```

6. **String Usage:**
   - Use single quotes or template literals (backticks) for strings.
   ```javascript
   let greeting = 'Hello, World!';
   let name = 'Alice';
   let message = `Hello, ${name}!`;
   ```

7. **Object and Array Formatting:**
   - Use shorthand syntax for object properties and methods.
   - Include a trailing comma for multi-line objects and arrays.
   ```javascript
   const user = {
       name: 'John',
       age: 30,
       greet() {
           console.log('Hello!');
       },
   };

   const numbers = [
       1,
       2,
       3,
       4,
   ];
   ```

8. **Arrow Functions:**
   - Use arrow functions for anonymous functions and callbacks.
   ```javascript
   const numbers = [1, 2, 3];
   const squares = numbers.map(n => n * n);
   ```

9. **Comments:**
   - Use `//` for single-line comments.
   - Use `/* */` for multi-line comments.
   ```javascript
   // This is a single-line comment

   /*
    * This is a multi-line comment
    * explaining the function below.
    */
   function example() {
       // code
   }
   ```

10. **Error Handling:**
   - Use `try` and `catch` for error handling.
    ```javascript
    try {
        // Code that may throw an error
        let result = riskyFunction();
    } catch (error) {
        console.error('An error occurred:', error);
    }
    ```

11. **Strict Mode:**
   - Use strict mode to catch common coding mistakes.
    ```javascript
    'use strict';
    // Your code here
    ```

12. **Consistent Bracing:**
   - Use consistent bracing style (1TBS - One True Brace Style).
    ```javascript
    if (condition) {
        // code
    } else {
        // code
    }

    function example() {
        // code
    }
    ```

13. **Avoid Global Variables:**
   - Avoid creating global variables. Use modules or IIFE (Immediately Invoked Function Expression) to encapsulate code.
    ```javascript
    (function() {
        // Encapsulated code
        let localVar = 'This is local';
    })();
    ```

### Example Code Block with Conventions

```javascript
'use strict';

// Constants
const MAX_USERS = 100;

// Variables
let userCount = 0;

// Function Expressions
const addUser = (name) => {
    if (userCount < MAX_USERS) {
        userCount++;
        console.log(`User ${name} added. Total users: ${userCount}`);
    } else {
        console.log('Max users reached');
    }
};

// Objects
const user = {
    name: 'Alice',
    age: 25,
    greet() {
        console.log(`Hello, ${this.name}`);
    },
};

// Arrays
const users = ['Alice', 'Bob', 'Charlie'];

// Error Handling
try {
    addUser('John');
    user.greet();
    console.log(users.map(user => user.toUpperCase()));
} catch (error) {
    console.error('An error occurred:', error);
}

// Comments
// Single-line comment

/*
 * Multi-line comment
 * explaining the following code block.
 */

(function() {
    // Encapsulated code
    let localVar = 'This is local';
    console.log(localVar);
})();
```
