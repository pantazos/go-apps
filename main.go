package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "virtualsports"
const mongoURI = "mongodb+srv://orestispantazos:9gJ3TvZ8HqSOUscY@virtualsports-cluster1.ddk9tdz.mongodb.net/virtualsports?retryWrites=true&w=majority" + dbName

type Employee struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MatchDate        string             `json:"matchDate,omitempty" validate:"required"`
	MatchTime        string             `json:"matchTime,omitempty" validate:"required"`
	Cid              string             `json:"cId,omitempty" validate:"required"`
	MatchID          string             `json:"matchId,omitempty" validate:"required"`
	LeagueID         string             `json:"leagueId,omitempty" validate:"required"`
	HomeTeam         string             `json:"homeTeam,omitempty" validate:"required"`
	AwayTeam         string             `json:"awayTeam,omitempty" validate:"required"`
	HalfTimeScore    string             `json:"halfTimeScore,omitempty" validate:"required"`
	FullTimeScore    string             `json:"fullTimeScore,omitempty" validate:"required"`
	TotalGoals       int32              `json:"totalGoals,omitempty" validate:"required"`
	FirstHalfValue   string             `json:"firstHalfValue,omitempty" validate:"required"`
	SecondHalfValue  string             `json:"secondHalfValue,omitempty" validate:"required"`
	HomeWinOdd       float64            `json:"homeWinOdd,omitempty" validate:"required"`
	DrawOdd          float64            `json:"drawOdd,omitempty" validate:"required"`
	AwayWinOdd       float64            `json:"awayWinOdd,omitempty" validate:"required"`
	Over_2_5_odd     float64            `json:"over_2_5_odd,omitempty" validate:"required"`
	Under_2_5_odd    float64            `json:"under_2_5_odd,omitempty" validate:"required"`
	Over_1_5_odd     float64            `json:"over_1_5_odd,omitempty" validate:"required"`
	Under_1_5_odd    float64            `json:"under_1_5_odd,omitempty" validate:"required"`
	Over_2_5_status  int32              `json:"over_2_5_status,omitempty" validate:"required"`
	Under_2_5_status int32              `json:"under_2_5_status,omitempty" validate:"required"`
	Over_1_5_status  int32              `json:"over_1_5_status,omitempty" validate:"required"`
	Under_1_5_status int32              `json:"under_1_5_status,omitempty" validate:"required"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/match", func(c *fiber.Ctx) error {

		query := bson.D{{}}

		cursor, err := mg.Db.Collection("match").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employees []Employee = make([]Employee, 0)

		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees)
	})

	log.Fatal(app.Listen(":3000"))
}
