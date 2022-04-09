## React Pt. 2

### React Virtual Dom
  * the react DOM keeps a representation of the view
  * react DOM only updates when there is a difference to be rendered
    * it only updates the parts that have changed NOT the entire DOM
  
  * When using react you should not use querySelector()
    * elements may be overwritten when the virtual DOM does an update
  * Don't directly change update or edit DOM nodes directly with properties like innterHTML
  * Always make UI changes through a component

### Collections in React
  * A data structure that collects any group of values together
    * Arrays and objects
  * In React a collection might be a list of components
  * React will automatically render all items in an array.
  * Make sure each element in the array has a unique key property:
  ``` javascript
    const headings = [
        <div key="1">1</div>
        <div key="2">2</div>
        <div key="3">3</div>
    ]
    //renders all 3 divs
    render({headings})
  ```
  * The virtual dom needs keys on each item in a list so it can figure out what has changed abd only update that specific item

  * 