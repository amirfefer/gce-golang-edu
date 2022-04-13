package controllers

import (
	"time"
	"strconv"
	computeLib "github.com/create-go-app/fiber-go-template/app/compute"
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Operate(c *fiber.Ctx) error {
	instanceName := c.Params("name")
	computeId, err := strconv.Atoi(c.Params("id"))
	operation := &models.ComputeOperation{}
		// Check, if received JSON data is valid.
		if err := c.BodyParser(operation); err != nil {
			// Return status 400 and error message.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	compute, err := db.GetCompute(computeId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	client, err := computeLib.CreateClient(&compute)
	defer computeLib.CloseClient(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	instance := &models.GCEInstance{Name: instanceName, Compute: compute}
	switch operation.ComputeOperation {
	case computeLib.Start:
		computeLib.StartInstance(client, instance)
			// Return status 200 OK.
	    return c.JSON(fiber.Map{
		    "error": false,
		    "msg":   nil,
		    "operation":  "started",
		    "instance": instance,
	    })
	case computeLib.Stop:
		computeLib.StopInstance(client, instance)
		return c.JSON(fiber.Map{
		    "error": false,
		    "msg":   nil,
		    "operation":  "stopped",
		    "instance": instance,
	    })
    default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "operation is not supported",
		})
    } 
}

// CreateInstance func for creates a new instance
// @Description Create a new instance.
// @Summary create a new compute
// @Tags Instance
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param path body string true "JsonPath"
// @Param projectID body string true "ProjectID"
// @Success 200 {object} models.Compute
// @Router /v1/compute [post]
func CreateInstance(c *fiber.Ctx) error {
	// Create new compute struct
	instance := &models.GCEInstance{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(instance); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Compute model.
	validate := utils.NewValidator()
	// number, err := strconv.ParseUint(string("90"), 10, 64)
	// Set initialized default data for compute:
	// compute.ID = rand.Intn(1000000) + 1
	instance.CreatedAt = time.Now()
    
	// Validate compute fields.
	if err := validate.Struct(instance); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete compute by given ID.
	if err := db.CreateInstance(instance); err != nil {
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
		"compute":  instance,
	})
}