## Publishing and Building

### npm run build
  * creates a build folder with all required files for the complete project
    * build folder can be deployed to github pages

### Steps to deploying on github pages
  * Run: npm run build
  * Run: npm install gh-pages --save-dev
  * Add this to your package.json:
    * "homepage": "https://<username>.github.io/<repo-name>"

  * Also in your package.json:
    * Add this to your list of scripts:
    * "predeploy": "npm run build",
    * "deploy": "gh-pages -d build"

  * Commit and push your work
  * Run: npm run deploy
    * it should automatically create a new branch and add your 