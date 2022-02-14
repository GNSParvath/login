package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admin")

func CreateAdmincontrol(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var admin models.AdminControl
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&admin); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newAdminControl := models.AdminControl{
		Id:                primitive.NewObjectID(),
		CompanyName:       admin.CompanyName,
		TIN:               admin.TIN,
		NumberOfEmployees: admin.NumberOfEmployees,
		Subscription:      admin.Subscription,
		FreeTrail:         admin.FreeTrail,
		Address:           admin.Address,
		ContactNUmber:     admin.ContactNUmber,
	}

	result, err := adminCollection.InsertOne(ctx, newAdminControl)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAdmincontrol(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	adminId := c.Params("adminId")
	var admin models.AdminControl
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(adminId)

	err := adminCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&admin)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": admin}})
}

func EditAdmincontrol(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	adminId := c.Params("adminId")
	var admin models.AdminControl
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(adminId)

	//validate the request body
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&admin); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"CompanyName": admin.CompanyName, "TIN": admin.TIN, "NumberOfEmployees": admin.NumberOfEmployees, "Subscription": admin.Subscription, "FreeTrail": admin.FreeTrail, "Address": admin.Address, "ContactNUmber": admin.ContactNUmber}

	result, err := adminCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated admin control details
	var updatedAdmincontrol models.AdminControl
	if result.MatchedCount == 1 {
		err := adminCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAdmincontrol)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedAdmincontrol}})
}

func DeleteAdmincontrol(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	adminId := c.Params("adminId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(adminId)

	result, err := adminCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Admin with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Admin control details successfully deleted!"}},
	)
}

func GetAlladmindata(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var admindata []models.AdminControl
	defer cancel()

	results, err := adminCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAdmin models.AdminControl
		if err = results.Decode(&singleAdmin); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		admindata = append(admindata, singleAdmin)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": admindata}},
	)
}
