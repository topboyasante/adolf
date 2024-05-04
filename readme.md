# Adolf

### About
Adolf is a CLI application that bootstraps a web application in go. Instead of setting up your backend project from scratch, you can use adolf to do that.


### Plans
- Initialize a simple backend application with: 
  - routing, 
  - sample models 
  - controllers.
  - All connected to a database
- Add command line flags to choose the DB of your choice
  
### How does it work?
- When you install adolf on your local environment, run `adolf init` to setup your backend project.
- You will be asked to provide a module name during the setup.
- When the setup is complete, run `go mod tidy` to install the neccessary packages used.
- In all import statements, replace `module_name` with the name of your module.