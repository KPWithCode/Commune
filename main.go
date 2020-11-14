package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//============================MODELS====================================

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// Blog Struct
type Blog struct {
	ID          string  `json:"id"`
	Author      *Author `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}

// Blogs struct
type Blogs struct {
	Blogs []Blog `json:"blogs"`
}

// Author struct
type Author struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//============================DATABASE====================================
// Database instance
var db *sql.DB

// Connect to db function
func Connect() error {
	// Database Settings
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT") // Default port
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

//============================MAIN====================================

func main() {
	// load env
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading env file!")
	}
	// connect
	if err := Connect(); err != nil {
		log.Fatal("Cant connect...", err)
	}


	app := fiber.New()
	// app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// app.Get("/blog", func(c *fiber.Ctx) error {
	// 	// Insert blog into database
	// 	rows, err := db.Query("SELECT id, author, title, description FROM Blogs order by id")
	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}
	// 	defer rows.Close()
	// 	result := Blogs{}

	// 	for rows.Next() {
	// 		blog := Blog{}
	// 		if err := rows.Scan(&blog.ID, &blog.Author, &blog.Title, &blog.Description); err != nil {
	// 			return err // Exit if we get an error
	// 		}

	// 		// Append Employee to Employees
	// 		result.Blogs = append(result.Blogs, blog)
	// 	}
	// 	// Return Employees in JSON format
	// 	return c.JSON(result)
	// })

	// // Add record into postgreSQL
	// app.Post("/blog", func(c *fiber.Ctx) error {
	// 	// New Employee struct
	// 	u := new(Blog)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Insert Employee into database
	// 	res, err := db.Query("INSERT INTO blogs (author, title, description)VALUES (KP, HI, WELCOME)", u.Author, u.Title, u.Description)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return Employee in JSON format
	// 	return c.JSON(u)
	// })

	// Update record into postgreSQL
	// app.Put("/employee", func(c *fiber.Ctx) error {
	// 	// New Employee struct
	// 	u := new(Employee)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Insert Employee into database
	// 	res, err := db.Query("UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5", u.Name, u.Salary, u.Age, u.ID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return Employee in JSON format
	// 	return c.Status(201).JSON(u)
	// })

	// Delete record from postgreSQL
	// app.Delete("/employee", func(c *fiber.Ctx) error {
	// 	// New Employee struct
	// 	u := new(Employee)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Insert Employee into database
	// 	res, err := db.Query("DELETE FROM employees WHERE id = $1", u.ID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return Employee in JSON format
	// 	return c.JSON("Deleted")
	// })

	log.Fatal(app.Listen(":5000"))
}
