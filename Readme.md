# Getting Started with Golang, Gin, Templ, and Tailwind CSS

In this tutorial, we'll build a Go web application using the **Gin** framework, **Templ** for HTML templating, and **Tailwind CSS** for styling. We'll also add interactivity using **HTMX**.

---

## Prerequisites

1. **Go**: Install Go from [https://golang.org/dl/](https://golang.org/dl/).
2. **Node.js and npm**: Install Node.js from [https://nodejs.org/](https://nodejs.org/).
3. **Templ**: Install the `templ` CLI:
   ```bash
   go install github.com/a-h/templ@latest
   ```

---

## Step 1: Set Up the Project

1. Create a new directory for your project:
   ```bash
   mkdir gin-tailwind-templ
   cd gin-tailwind-templ
   ```

2. Initialize a Go module:
   ```bash
   go mod init gin-tailwind-templ
   ```

3. Install Gin:
   ```bash
   go get -u github.com/gin-gonic/gin
   ```

---

## Step 2: Set Up Tailwind CSS

1. Initialize a `package.json` file:
   ```bash
   npm init -y
   ```

2. Install Tailwind CSS and its dependencies:
   ```bash
   npm install -D tailwindcss postcss autoprefixer
   ```

3. Initialize Tailwind CSS:
   ```bash
   npx tailwindcss init
   ```

4. Create a `tailwind.config.js` file and configure it:
   ```javascript
   // tailwind.config.js
   module.exports = {
     content: ["./templates/**/*.templ"], // Watch .templ files for Tailwind classes
     theme: {
       extend: {},
     },
     plugins: [],
   };
   ```

5. Create a `postcss.config.js` file:
   ```javascript
   // postcss.config.js
   module.exports = {
     plugins: {
       tailwindcss: {},
       autoprefixer: {},
     },
   };
   ```

6. Create a `styles.css` file in a `static/css` directory:
   ```bash
   mkdir -p static/css
   touch static/css/styles.css
   ```

   Add the following to `static/css/styles.css`:
   ```css
   @tailwind base;
   @tailwind components;
   @tailwind utilities;
   ```

7. Add a script to `package.json` to build Tailwind CSS:
   ```json
   "scripts": {
     "build:css": "tailwindcss -i ./static/css/styles.css -o ./static/css/output.css --watch"
   }
   ```

8. Run the Tailwind CSS build process:
   ```bash
   npm run build:css
   ```

---

## Step 3: Create Templ Templates

1. Create a `templates` directory:
   ```bash
   mkdir templates
   ```

2. Create a `home.templ` file:
   ```bash
   touch templates/home.templ
   ```

   Add the following to `templates/home.templ`:
   ```templ
   package templates

   templ HomePage(title string, message string) {
       <!DOCTYPE html>
       <html lang="en">
           <head>
               <meta charset="UTF-8" />
               <meta name="viewport" content="width=device-width, initial-scale=1.0" />
               <title>{ title }</title>
               <link href="/static/css/output.css" rel="stylesheet" />
               <script src="https://unpkg.com/htmx.org@1.9.10"></script>
           </head>
           <body class="bg-gray-100">
               <div class="container mx-auto p-4">
                   <h1 class="text-2xl font-bold text-blue-500">{ title }</h1>
                   <p class="mt-4 text-gray-700">{ message }</p>
                   <button hx-get="/api/data" hx-swap="innerHTML" class="bg-blue-500 text-white p-2 rounded">
                       Click Me
                   </button>
               </div>
           </body>
       </html>
   }
   ```

3. Generate Go code from the `.templ` file:
   ```bash
   templ generate
   ```

   This will create a `home_templ.go` file in the `templates` directory.

---

## Step 4: Set Up the Gin Server

1. Create a `main.go` file:
   ```bash
   touch main.go
   ```

2. Add the following code to `main.go`:
   ```go
   package main

   import (
       "net/http"

       "github.com/gin-gonic/gin"
       "github.com/lordofthemind/gin-tailwind-templ/templates" // Adjust the import path
   )

   func main() {
       r := gin.Default()

       // Serve static files
       r.Static("/static", "./static")

       // Render the home page
       r.GET("/", func(c *gin.Context) {
           templates.HomePage("Gin + Tailwind + Templ", "This is a dynamic message!").Render(c.Request.Context(), c.Writer)
       })

       // Handle HTMX request
       r.GET("/api/data", func(c *gin.Context) {
           c.String(http.StatusOK, "HTMX response!")
       })

       // Start the server
       r.Run(":8080")
   }
   ```

---

## Step 5: Run the Application

1. Start the Tailwind CSS build process (if not already running):
   ```bash
   npm run build:css
   ```

2. Start the Go server:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to `http://localhost:8080`. You should see a page styled with Tailwind CSS and an interactive button powered by HTMX.

---

## Directory Structure

Here’s what your project should look like:

```
.
├── go.mod
├── go.sum
├── main.go
├── package-lock.json
├── package.json
├── postcss.config.js
├── static
│   └── css
│       ├── output.css
│       └── styles.css
├── tailwind.config.js
├── templates
│   ├── home.templ
│   └── home_templ.go
└── tree.txt
```

---

## Step 6: Explanation

1. **Tailwind CSS**:
   - Tailwind is a utility-first CSS framework. We configured it to watch `.templ` files for class names and generate a CSS file (`output.css`).

2. **Templ**:
   - Templ is a Go library for writing HTML templates in a type-safe way. We created a `.templ` file and used `templ generate` to convert it into Go code.

3. **Gin**:
   - Gin is a web framework for Go. We used it to serve static files (CSS) and render the Templ-generated HTML.

4. **HTMX**:
   - HTMX is a lightweight library that allows you to add interactivity to your HTML without writing JavaScript.

5. **Static Files**:
   - The `static` directory contains the Tailwind CSS output. Gin serves this directory at the `/static` route.

---

## Step 7: Extending the Example

1. **Add More Pages**:
   - Create additional `.templ` files and render them in your Gin routes.

2. **Add More Tailwind Styles**:
   - Customize your `tailwind.config.js` to add custom colors, fonts, etc.

3. **Build for Production**:
   - Minify the CSS by adding `NODE_ENV=production` to the Tailwind build script:
     ```json
     "scripts": {
       "build:css": "NODE_ENV=production tailwindcss -i ./static/css/styles.css -o ./static/css/output.css --minify"
     }
     ```

---

## Conclusion

You’ve successfully built a Go web application using Gin, Templ, and Tailwind CSS! This setup provides a modern, type-safe, and efficient way to build web applications in Go. You can now extend this project by adding more pages, passing dynamic data to templates, and enhancing interactivity with HTMX.

