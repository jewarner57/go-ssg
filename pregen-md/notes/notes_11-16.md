## Single Page React Applications

### React Intro

### Single Page Applications
  * web pages that act like multipage static sites but are built from a single html file.
    * links dont open a new html file they just modify the dom
    * ex: twitter's infinite scroll
    * netflix's video preview
    * gmail, infinite scrolling and compose are in the same place

### ES6
  * the latest version of the javascript language

### React
  * A javascript library for building user interfaces
    * Declarative and Component based
      * you write the code that creates your views

  * Components
    * ui elements used like buidling blocks
    * reat component example
    ``` javascript
        function MyComponent(props) {
            return (
            <div>
                <h1>Hello</h1>
            </div>
            )
        }
    ```
    * components should **always** start with an uppercase letter
    * 

  * JSX 
    * an extension to the javascript language that provides xml syntax used to express your ui
    * jsx allows us to write html along side regular javascript
    * single line return statements in jsx dont need parantheses
    * there can only be a single top level node (everything must be wrapped in a single div/tag)
    * must use className instead of class for selectors
    * to run javascript in a jsx block wrap the javascript in { }
      * example
      ``` javascript
      <div>{console.log("Hello")}</div>
      ```
    * 
