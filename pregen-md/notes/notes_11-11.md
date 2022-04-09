## Bundling

### What is bundling?
  * compressing and processing our code into **one** file.

### Minify
  * removes all spaces and line breaks
  * also removes comments

### Uglify 
   * renames elements, obfuscates code and reduces file size
   * might rename a variable to a or b

### Compatibility
   * converts ES6 into older javascript that is compatible with older browser

### Why bundle?
  * makes smaller and fewer files to make websites load faster.

<hr>

### Bundling with Webpack
  * install webpack with npm
  * create a src and a dist folder
  * put your js into the src and your html into dist
  * create your webpack.config.js example:
    ``` javascript
    const path = require('path')

    module.exports = {
      entry: './src/index.js',
      output: {
      filename: 'bundle.js',
        path: path.resolve(__dirname, 'dist')
      }
    }
    ```

  * Add webpack scripts in your package.json
    * Build is for production, develop is for dev
      * use npm run build/develop
    ``` json
    "scripts": {
      "develop": "webpack --mode development --watch",
      "build": "webpack --mode production"
    },
    ```

### _