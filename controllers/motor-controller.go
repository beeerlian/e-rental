package controllers

import (
	"context"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/beeerlian/motor-rental/config"
	"github.com/beeerlian/motor-rental/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllMotors(c *fiber.Ctx) error {
	motorCollection := config.MI.DB.Collection("motors")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var motors []models.Motor

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"type": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"police_number": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := motorCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := motorCollection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Motors Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var motor models.Motor
		cursor.Decode(&motor)
		motors = append(motors, motor)
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      motors,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}

func GetMotor(c *fiber.Ctx) error {
	motorCollection := config.MI.DB.Collection("motors")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var motor models.Motor
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := motorCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Motor Not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&motor)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Motor Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    motor,
		"success": true,
	})
}

func AddMotor(c *fiber.Ctx) error {
	motorCollection := config.MI.DB.Collection("motors")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	motor := new(models.Motor)

	if err := c.BodyParser(motor); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	// customerObjId, err := primitive.ObjectIDFromHex(motor.Lecturer)

	// err = findCustomerResult.Decode(&customer)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"success": false,
	// 		"result":  findCustomerResult,
	// 		"message": "Failed to get Customer by Id",
	// 		"error":   err,
	// 	})
	// }

	// customer.MotorsCreated = append(customer.MotorsCreated, *motor)

	result, err := motorCollection.InsertOne(ctx, motor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Motor failed to insert",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "Motor inserted successfully",
	})
}

func UpdateMotor(c *fiber.Ctx) error {
	motorCollection := config.MI.DB.Collection("motors")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	motor := new(models.Motor)

	if err := c.BodyParser(motor); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Motor not found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": motor,
	}
	_, err = motorCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Motor failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Motor updated successfully",
	})
}

func DeleteMotor(c *fiber.Ctx) error {
	motorCollection := config.MI.DB.Collection("motors")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Motor not found",
			"error":   err,
		})
	}
	_, err = motorCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Motor failed to delete",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Motor deleted successfully",
	})
}

func RentMotor(c *fiber.Ctx) error {
	motorCollection := config.MI.DB.Collection("motors")
	customerCollection := config.MI.DB.Collection("customers")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var customer models.Customer
	var motor models.Motor
	var customerActivity models.CustomerActivity

	motorObjId, err := primitive.ObjectIDFromHex(c.Params("motorId"))
	customerObjId, err := primitive.ObjectIDFromHex(c.Params("customerId"))

	findMotorResult := motorCollection.FindOne(ctx, bson.M{"_id": motorObjId})
	findCustomerResult := customerCollection.FindOne(ctx, bson.M{"_id": customerObjId})

	err = findMotorResult.Decode(&motor)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get Motor by Id",
			"error":   err,
		})
	}
	err = findCustomerResult.Decode(&customer)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get Customer by Id",
			"error":   err,
		})
	}
	customerActivity = models.CustomerActivity{MotorId: motor.ID, MotorType: motor.Type, Attende: "no"}

	customer.Activities = append(customer.Activities, customerActivity)

	updateMotor := bson.M{
		"$set": motor,
	}

	updateCustomer := bson.M{
		"$set": customer,
	}

	_, err = motorCollection.UpdateOne(ctx, bson.M{"_id": motorObjId}, updateMotor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to add participant",
			"error":   err.Error(),
		})
	}

	_, err = customerCollection.UpdateOne(ctx, bson.M{"_id": customerObjId}, updateCustomer)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to add customerActivity",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Customer joined successfully",
	})
}
