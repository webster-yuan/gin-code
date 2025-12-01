package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// Create a template function map
var funcMap = template.FuncMap{
	"safe": func(str string) template.HTML {
		return template.HTML(str)
	},
}

func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templateDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err)
	}
	includes, err := filepath.Glob(templateDir + "/includes/*.tmpl")
	if err != nil {
		panic(err)
	}
	pages, err := filepath.Glob(templateDir + "/*.html") // html files in root
	if err != nil {
		panic(err)
	}

	// Add tmpl files in root directory
	tmplPages, err := filepath.Glob(templateDir + "/*.tmpl")
	if err != nil {
		panic(err)
	}
	pages = append(pages, tmplPages...)

	// Register simple pages
	for _, p := range pages {
		// For each template, create it with the funcMap
		tmplName := filepath.Base(p)
		t, err := template.New(tmplName).Funcs(funcMap).ParseFiles(p)
		if err != nil {
			panic(err)
		}
		r.Add(tmplName, t)
	}

	// Generate templates map for layouts and includes
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)

		// For each template with inheritance, create it with the funcMap
		tmplName := filepath.Base(include)
		t, err := template.New(tmplName).Funcs(funcMap).ParseFiles(files...)
		if err != nil {
			panic(err)
		}
		r.Add(tmplName, t)
	}
	return r
}

// SetupTemplates configures template settings
func SetupTemplates(router *gin.Engine, basePath string) {
	// Set the function map for gin's default renderer too
	router.SetFuncMap(funcMap)

	// Use template inheritance - this should be the only template loading method
	templateDir := filepath.Join(basePath, "templates")
	router.HTMLRender = loadTemplates(templateDir)

	// Serve static files
	router.Static("/static", filepath.Join(basePath, "static"))
}

// IndexFunc handles the index page
func IndexFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Now":   time.Now().Format("2006-01-02 15:04:05"),
			"Items": []string{"Gin Framework", "Template Inheritance", "Multitemplate", "Static Files"},
		})
	}
}

// HomeFunc handles the home page
func HomeFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Now": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
}
