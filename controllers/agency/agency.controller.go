package agency

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-fiber-jwt/initializers"
	"github.com/wpcodevo/golang-fiber-jwt/models/Agency"
	"github.com/wpcodevo/golang-fiber-jwt/models/user"
	"github.com/wpcodevo/golang-fiber-jwt/utils"
)

func CreateAgency(c *fiber.Ctx) error {
	var payload *Agency.CreateAgency
	var agency Agency.Agency

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest, ResponseMsg: err.Error()})
	}

	errors := user.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: "Утга зөв эсэхийг шалгана уу", Data: errors})

	}

	agency.Name = payload.Name
	agency.Address = payload.Address
	agency.Phone = payload.Phone
	agency.Email = payload.Email
	agency.Website = payload.Website
	agency.Description = payload.Description
	agency.City = payload.City
	result := initializers.DB.Create(&agency)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: "Алдаа гарлаа", Data: result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseObj{ResponseCode: fiber.StatusOK,
		ResponseMsg: "Амжилттай бүртгэлээ", Data: agency})

}

func GetAgentList(c *fiber.Ctx) error {
	var agents []Agency.Agency
	var request utils.RequestObj
	var conn = initializers.DB

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusOK).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: err.Error()})
	}

	conn = initializers.DB.
		Model(&Agency.Agency{}).
		Scopes(utils.Filter(request.Filter))

	pagination := utils.Pagination{CurrentPageNo: request.PageNo, PerPage: request.PerPage, Sort: request.Sort}
	conn.Debug().
		Scopes(utils.Paginate(agents, &pagination, conn)).
		Find(&agents)

	return c.Status(fiber.StatusOK).JSON(utils.ResponseObj{ResponseCode: fiber.StatusOK,
		ResponseMsg: "Амжилттай", Data: agents, Pagination: &pagination})

}

func GetAgent(c *fiber.Ctx) error {
	id := c.Params("id")
	var agent Agency.Agency
	result := initializers.DB.First(&agent, "id = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: "Алдаа гарлаа", Data: result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseObj{ResponseCode: fiber.StatusOK,
		ResponseMsg: "Амжилттай", Data: agent})

}

func UpdateAgent(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *Agency.UpdateAgency
	var agent Agency.Agency
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest, ResponseMsg: err.Error()})
	}

	errors := user.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: "Утга зөв эсэхийг шалгана уу", Data: errors})

	}

	result := initializers.DB.Model(&agent).Where("id = ?", id).Updates(Agency.Agency{Name: payload.Name,
		Address: payload.Address, Phone: payload.Phone, Email: payload.Email, Website: payload.Website, Description: payload.Description, City: payload.City})
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: "Алдаа гарлаа", Data: result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseObj{ResponseCode: fiber.StatusOK,
		ResponseMsg: "Амжилттай", Data: agent})

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var agent Agency.Agency
	result := initializers.DB.Delete(&agent, "id = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseObj{ResponseCode: fiber.StatusBadRequest,
			ResponseMsg: "Алдаа гарлаа", Data: result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseObj{ResponseCode: fiber.StatusOK,
		ResponseMsg: "Амжилттай", Data: agent})

}