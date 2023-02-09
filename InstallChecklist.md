# Install These

## Angular

You'll need the Angular CLI as well as NPM (Node Package Manager).  
To install NPM, you'll need to install Node.js.

[Link to install NPM, with guide on installing NVM first](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

It's recommended to use NVM to install. Skip if just using a Node installer.  
NVM was meant for Unix systems, but there's a tool called nvm-windows.  
Search and install either version from their corresponding Github repo.  
Run `nvm install Node`. Note that you need to be in cmd admin mode or in Powershell.

You can check your installation worked with `node -v` and `npm -v`. Both should return version numbers.  
If you want to make sure your NPM is up-to-date, run `npm install -g npm`.  

Finally, you can install Angular by runnning `npm install -g @angular/cli`.

[Link to some basic `ng` commands and overview](https://angular.io/cli)

## Go

[Link to install Go](https://go.dev/doc/install)

Install the correct version for your system.

GOPATH does not seem to be absolutely necessary anymore, but this document will be updated if otherwise.

In Go, modules contain packages, which contain source files.  
There should only be one module in the root of the repository.  
A file at the root, **go.mod**, declares the module path (the import path prefix for packages).  
The path also indicates where any `go` commands should look to download.

You also need to install the Go plugin if you're using VS Code. It's just called Go.

Go runs from the terminal using the command `go run filename.go`.

*EXAMPLE: module `github.com/google/go-cmp` has a package in the directory `cmp/`. The import path is `github.com/google/go-cmp/cmp`. [Link to more on this and a first project tutorial](https://go.dev/doc/code)*