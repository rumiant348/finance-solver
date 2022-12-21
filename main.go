package main

import (
	"finance-solver.com/config"
	"finance-solver.com/controllers"
	_ "finance-solver.com/docs"
	"finance-solver.com/middleware"
	"finance-solver.com/models"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
)

//	@title		Finance Solver Api 2
//	@version	1.0
//	@host		localhost:3000
func main() {
	cfg := config.GetConfig()

	services, err := models.NewServices(
		models.WithGorm(cfg.Dialect(), cfg.ConnectionInfo()),
		//models.WithLogMode(!cfg.IsProd()),
		models.WithLogMode(false),
		models.WithUser(cfg.Pepper, cfg.HMACKey),
		models.WithList(),
		//models.WithList(),
		//models.WithExpense(),
	)
	//err = services.AutoMigrate()
	err = services.DestructiveReset()
	if err != nil {
		panic(err)
	}
	defer services.Close()

	// user middleware
	userMw := middleware.User{
		UserService: services.User,
	}
	// auth middleware
	requireUserMw := middleware.RequireUser{}

	usersC := controllers.NewUsers(services.User)
	listsC := controllers.NewLists(services.List)

	r := mux.NewRouter()
	// User routes
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	// List routes
	//r.GET("/lists", e.GetExpenses)
	//r.GET("/lists/:id", e.GetExpensesById)
	r.HandleFunc("/lists", requireUserMw.ApplyFn(listsC.GetAll)).Methods("GET")
	r.HandleFunc("/lists", requireUserMw.ApplyFn(listsC.Create)).Methods("POST")
	//r.POST("/lists", e.PostExpenses)
	//r.DELETE("/lists/:id", e.DeleteExpenseById)
	//
	//
	//// Expense routes
	//r.GET("/expenses", e.GetExpenses)
	//router.GET("/expenses/:id", e.GetExpensesById)
	//router.POST("/expenses", e.PostExpenses)
	//router.DELETE("/expenses/:id", e.DeleteExpenseById)

	// Swagger
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Printf("Starting the server on :%d...\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port),
		handlers.LoggingHandler(os.Stdout, userMw.Apply(r)),
	)
	if err != nil {
		panic(err)
	}
}
