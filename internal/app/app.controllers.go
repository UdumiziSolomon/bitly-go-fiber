package app

import (
	"strconv"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetGollies(c *fiber.Ctx) error {
	golies, err := getAllGolies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error getting all goly links" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(golies)
}

func GetGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	goly, err := getGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not retrieve" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func CreateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")   // to accept json
	var goly Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = RandomizeURL(8)
	}
	if err = createGoly(goly); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create goly in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func UpdateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if err = updateGoly(goly); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not update goly in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func DeleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	if err = deleteGoly(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "goly deleted successfully",
	})
}

func Redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")
	goly, err := getGolyByUrl(golyUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not find goly in DB " + err.Error(),
		})
	}

	// update any field
	goly.Clicked += 1
	if err = updateGoly(goly); err != nil {
		log.Printf("Error updating DB, %v\n", err.Error())
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}