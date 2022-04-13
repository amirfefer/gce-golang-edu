package controllers

import (
	"time"
	"strconv"
	computeLib "github.com/create-go-app/fiber-go-template/app/compute"
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
	"math/rand"
)

func GetZones(c *fiber.Ctx) error {
	zones, err := computeLib.ZoneList(c.Params("id"))
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(zones),
		"zones": zones,
	})
}


// GetComputes func gets all exists computes.
// @Description Get all exists compute objects.
// @Summary get all exists compute
// @Tags Computes
// @Accept json
// @Produce json
// @Success 200 {array} models.Compute
// @Router /v1/computes [get]
func GetComputes(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all computes.
	computes, err := db.GetComputes()
	if err != nil {
		// Return, if computes not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"count": 0,
			"computes": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(computes),
		"computes": computes,
	})
}

// CreateCompute func for creates a new compute.
// @Description Create a new compute.
// @Summary create a new compute
// @Tags Compute
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param path body string true "JsonPath"
// @Param projectID body string true "ProjectID"
// @Success 200 {object} models.Compute
// @Router /v1/compute [post]
func CreateCompute(c *fiber.Ctx) error {
	// Create new compute struct
	compute := &models.Compute{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(compute); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"parser": true,
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"db": true,
		})
	}

	// Create a new validator for a Compute model.
	validate := utils.NewValidator()
    rand.Seed(time.Now().UnixNano())
	// Set initialized default data for compute:
	compute.ID = rand.Intn(1000000) + 1
	compute.CreatedAt = time.Now()

	// Validate compute fields.
	if err := validate.Struct(compute); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
			"validation": true,
		})
	}

	// Delete compute by given ID.
	if err := db.CreateCompute(compute); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"compute":  compute,
	})
}

// GetInstancesList func gets all instacnes.
// @Description Get all exists instacnes.
// @Summary get all instacnes
// @Tags GCEInstance
// @Accept json
// @Produce json
// @Success 200 {array} models.GCEInstance
// @Router /v1/computes/:id/list [get]
func GetInstancesList(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get compute.
	computeID, err := strconv.Atoi(c.Params("id"))
	compute, err := db.GetCompute(computeID)
	if err != nil {
		// Return, if computes not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "compute were not found",
			"count": 0,
			"computes": nil,
		})
	}
	client, error := computeLib.CreateClient(&compute)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer computeLib.CloseClient(client)
	instances, err := computeLib.ListInstances(client, &compute)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(instances),
		"instances": instances,
	})
}