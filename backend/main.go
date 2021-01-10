package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team03/app/controllers"
	_ "github.com/team03/app/docs"
	"github.com/team03/app/ent"
)

type PriceTypes struct {
	PriceType []PriceTypes
}
type PriceType struct {
	Name string
}

type Nurses struct {
	Nurse []Nurse
}
type Nurse struct {
	Name string
}

type Rooms struct {
	Room []Room
}
type Room struct {
	Name string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDentistController(v1, client)
	controllers.NewNurseController(v1, client)
	controllers.NewMedicalfileController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewMedicalCareController(v1, client)
	controllers.NewDentalExpenseController(v1, client)
	controllers.NewPriceTypeController(v1, client)
	controllers.NewAppointmentController(v1, client)
	controllers.NewRoomController(v1, client)

	// Set PriceType Data
	pricetypes := PriceTypes{
		PriceType: []PriceType{
			PriceType{"จ่ายเองโดยตรง"},
			PriceType{"เบิกได้"},
			PriceType{"ญาติผู้ป่วย"},
		},
	}

	for _, pt := range pricetypes.PriceType {

		client.PriceType.
			Create().
			SetName(pt.Name).
			Save(context.Background())
	}

	// Set Room Data
	rooms := Rooms{
		Room: []Room{
			Room{"ห้องหัตถการ 1"},
			Room{"ห้องหัตถการ 2"},
			Room{"ห้องตรวจ R1"},
		},
	}

	for _, r := range rooms.Room {

		client.Room.
			Create().
			SetName(r.Name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
