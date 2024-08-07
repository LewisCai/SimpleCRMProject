package lead

import (
	"context"
	"log"

	"github.com/LewisCai/SimpleCRMProject/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Lead struct {
	ID 			primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string `json:"name"`
	Company 	string `json:"company"`
	Email 		string `json:"email"`
	Phone 		string `json:"phone"`

}

// GetLeads handles GET requests to retrieve all leads
func GetLeads(c *fiber.Ctx) error {
	collection := database.Client.Database("simplecrm").Collection("leads")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(context.TODO())

	var leads []Lead
	if err = cursor.All(context.TODO(), &leads); err != nil {
		log.Fatal(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(leads)
}

// GetLead handles GET requests to retrieve a specific lead by ID
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	collection := database.Client.Database("simplecrm").Collection("leads")
	var lead Lead
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&lead)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("Lead not found")
		}
		log.Fatal(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(lead)
}

// NewLead handles POST requests to create a new lead
func NewLead(c *fiber.Ctx) error {
	collection := database.Client.Database("simplecrm").Collection("leads")
	var lead Lead
	if err := c.BodyParser(&lead); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	lead.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), lead)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(lead)
}

// DeleteLead handles DELETE requests to remove a lead by ID
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	collection := database.Client.Database("simplecrm").Collection("leads")
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
		return c.Status(500).SendString(err.Error())
	}

	if res.DeletedCount == 0 {
		return c.Status(404).SendString("Lead not found")
	}

	return c.SendStatus(fiber.StatusNoContent)
}