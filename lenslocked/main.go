package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/csrf"
	"github.com/joncalhoun/lenslocked/controllers"
	"github.com/joncalhoun/lenslocked/migrations"
	"github.com/joncalhoun/lenslocked/models"
	"github.com/joncalhoun/lenslocked/templates"
	"github.com/joncalhoun/lenslocked/views"
)

type SomeType struct {
	Template views.Template
}

func (st SomeType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st.Template.Execute(w, r, nil)
}

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse(filepath)
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "There was an error parsing the template.", http.
// 			StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }

func main() {
	// Setup the database
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// Setup middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: fix this before deploying
		csrf.Secure(false),
	)

	// Setup controllers
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	// Setup our router and routes
	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml",
	))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml", "tailwind.gohtml",
	))))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})
	// 	r.Get("/users/me", usersC.CurrentUser)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// Start the server
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
