package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Define the structure for mapping YAML data
type Acronyms struct {
	Acronyms map[string]string `yaml:"acronyms"`
}

// Function to parse the YAML file
func parseYAML(filename string) (*Acronyms, error) {
	var acronyms Acronyms
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &acronyms)
	if err != nil {
		return nil, err
	}
	return &acronyms, nil
}

// Function to save the updated acronyms back to the YAML file
func saveYAML(filename string, acronyms *Acronyms) error {
	data, err := yaml.Marshal(acronyms)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Function to handle the CLI
func handleCLI(filename string, acronym string, definition string) {
	acronyms, err := parseYAML(filename)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	if definition != "" {
		// Add new acronym
		acronyms.Acronyms[acronym] = definition
		err = saveYAML(filename, acronyms)
		if err != nil {
			log.Fatalf("Error saving to YAML file: %v", err)
		}
		fmt.Printf("Added acronym: %s = %s\n", acronym, definition)
	} else {
		// Search for acronym
		def, exists := acronyms.Acronyms[acronym]
		if exists {
			fmt.Printf("Definition of %s: %s\n", acronym, def)
		} else {
			fmt.Printf("Acronym '%s' not found\n", acronym)
		}
	}
}

// Function to start the Fiber web server
func startHTTPServer(filename string) {
	// Initialize the Fiber app with an HTML engine for template rendering
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files from the "static" directory
	fmt.Printf("Loading statics file ...")
	app.Static("/static", "./static")

	// Render the main page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	// Handle the search request
	app.Get("/search", func(c *fiber.Ctx) error {
		acronym := c.Query("acronym")
		acronyms, err := parseYAML(filename)
		if err != nil {
			return c.Status(500).SendString("Error parsing YAML file")
		}

		// Search for the acronym
		definition, exists := acronyms.Acronyms[acronym]
		if exists {
			return c.Render("index", fiber.Map{
				"Acronym": acronym,
				"Result":  definition,
			})
		} else {
			return c.Render("index", fiber.Map{
				"Acronym": acronym,
			})
		}
	})

	// Handle the add acronym form
	app.Post("/add", func(c *fiber.Ctx) error {
		newAcronym := c.FormValue("newAcronym")
		definition := c.FormValue("definition")
		acronyms, err := parseYAML(filename)
		if err != nil {
			return c.Status(500).SendString("Error parsing YAML file")
		}

		// Add the new acronym
		acronyms.Acronyms[newAcronym] = definition
		err = saveYAML(filename, acronyms)
		if err != nil {
			return c.Status(500).SendString("Error saving YAML file")
		}

		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3000"))
}

func main() {
	// Define the CLI flags
	acronym := flag.String("acronym", "", "The acronym to look up or add")
	definition := flag.String("definition", "", "The definition of the acronym (used when adding an acronym)")
	httpMode := flag.Bool("http", false, "Run the application in HTTP mode (launches a web server)")
	flag.Parse()

	filename := "acronyms.yaml"

	if *httpMode {
		// Launch the web server
		fmt.Println("Starting the web server on http://localhost:3000")
		startHTTPServer(filename)
	} else {
		// Run in CLI mode
		if *acronym == "" {
			log.Fatal("Please provide an acronym using the -acronym flag")
		}
		handleCLI(filename, *acronym, *definition)
	}
}
