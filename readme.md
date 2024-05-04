# Adolf

### About

Adolf is a command-line interface (CLI) tool designed to streamline the setup process for Go web applications. Instead of manually configuring your backend project, Adolf automates the process for you.

## Features

- **Quick Setup**: Initialize a basic backend application with pre-configured routing, sample models, and controllers.
- **Database Integration**: Automatically connect your application to a database of your choice.


  
### How does it work?

#### `adolf init`
- Run this command after installing Adolf to bootstrap your backend project.
- During setup, provide a name for your module when prompted.
- After setup completes, execute `go mod tidy` to install required packages.
- Replace `module_name` with your actual module name in all import statements.
- Update the database connection string in `config/app.go` with your credentials.
  

#### `adolf heil`
- Use `adolf heil config.adolf.toml` to configure your backend project with specific database settings.
- Adolf references `config.adolf.toml` to determine the database provider for your project.
- Follow the prompts to provide a module name during setup.
- Replace `module_name` with your actual module name in all import statements.
- After setup, execute `go mod tidy` to install necessary packages.
- Update the database connection string in `config/app.go` with your credentials.
