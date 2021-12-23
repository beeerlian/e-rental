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

func AddTransaction(c *fiber.Ctx) error {
	transactionCollection := config.MI.DB.Collection("transactions")
	customerCollection := config.MI.DB.Collection("customers")
	motorCollection := config.MI.DB.Collection("motors")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	transaction := new(models.Transaction)
	var customer models.Customer
	var motor models.Motor

	if err := c.BodyParser(transaction); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

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

	transaction.Customer = customer
	transaction.Motor = motor

	result, err := transactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to insert transaction",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"result":  result,
		"message": "Transaction added successfully",
	})
}

func GetAllTransactions(c *fiber.Ctx) error {
	transactionCollection := config.MI.DB.Collection("transactions")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var transactions []models.Transaction

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

	total, _ := transactionCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := transactionCollection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Transactions Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var transaction models.Transaction
		cursor.Decode(&transaction)
		transactions = append(transactions, transaction)
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      transactions,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}
