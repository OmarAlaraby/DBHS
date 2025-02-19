# DBHS Project Architecture

## Overview
DBHS is structured into multiple packages to ensure modularity and maintainability. Each package is responsible for a specific domain of the application.

---

## Packages and Files

### 1. `main` Package
The **main** package is the entry point of the project, responsible for initializing the server, setting up database connections, and defining global URLs.

#### **Files in `main` package**:
- **server.go**
    - Connects to the PostgreSQL database.
    - Loads configuration settings.
    - Initializes the HTTP server.
    - Starts the application.

- **settings.go**
    - Stores global configuration variables, such as database credentials and server settings.

- **urls.go**
    - Defines global URL routing by aggregating routes from different packages.
    - **Example:**
      ```go
      func defineURLs() {
          accounts.DefineURLs()
          projects.DefineURLs()
          databases.DefineURLs()
      }
      ```
    - Any new package should expose a `DefineURLs()` function to register its routes.

---

### 2. `config` Package
The **config** package contains shared components such as the application instance, HTTP multiplexer, and logging setup.

#### **Files in `config` package**:
- **application.go**
    - Defines the `Application` structs and variables, which holds references to logs and other shared components. And also contains other varialbes and init function
    - **Example:**
      ```go
      package config
  
      import "log"
  
      type Application struct {
          ErrorLog *log.Logger
          InfoLog  *log.Logger
      }
      ```

- **helpers.go**
    - contains the helpers for the application struct

---

### 3. `accounts` Package
Handles user authentication and account management functionalities.

#### **Files in `accounts` package**:
- **urls.go**
    - Registers account-related routes in the global multiplexer.
    - **Example:**
      ```go
      package accounts
  
      import (
          "DBHS/config"
          "net/http"
      )
  
      func DefineURLs() {
          config.Mux.HandleFunc("/signup", SignUpHandler)
          config.Mux.HandleFunc("/signin", SignInHandler)
          // other endpoitns
      }
      ```

- **handlers.go**
    - Contains HTTP handler functions for user authentication.
    - **Example:**
      ```go
      package accounts
  
      import (
          "DBHS/config"
          "net/http"
      )
  
      func SignUpHandler(w http.ResponseWriter, r *http.Request) {
          config.App.InfoLog.Println("Sign up request received")
          w.Write([]byte("Sign up successful"))
      }
  
      func SignInHandler(w http.ResponseWriter, r *http.Request) {
          config.App.InfoLog.Println("Sign in request received")
          w.Write([]byte("Sign in successful"))
      }
      ```

---

### 4. Additional Packages (`projects`, `databases`, etc.)
Each additional feature is added as a separate package following the **accounts** package structure.

#### **Files in each feature package**:
- **models.go**
    - Design the modle in the database
- **urls.go**
    - Registers routes for the package.
- **handlers.go**
    - Implements the business logic for the package.

