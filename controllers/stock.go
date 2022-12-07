package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllStock(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	method := "GET"
	tag := "C"
	if c.Query("tag") != "" {
		tag = c.Query("tag")
	}

	url := fmt.Sprintf("%s/stock?tag=%s", configs.API_TRIGGER_URL, tag)
	if c.Query("part_no") != "" {
		url = fmt.Sprintf("%s/stock?tag=%s&part_no=%s", configs.API_TRIGGER_URL, tag, c.Query("part_no"))
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	res, err := client.Do(req)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.OraResponse
	if err = json.Unmarshal(body, &obj); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Get PartName
	var data []models.OraStock
	for _, i := range obj.Data {
		var p models.Part
		if err = db.Where("title=?", i.PartNo).First(&p).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		i.Slug = p.Slug
		i.PartName = p.Description
		data = append(data, i)
	}

	r.Message = obj.Message
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateStock(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowStockByID(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	method := "GET"
	tag := "C"
	if c.Query("tag") != "" {
		tag = c.Query("tag")
	}

	url := fmt.Sprintf("%s/stock/%s?tag=%s", configs.API_TRIGGER_URL, c.Params("id"), tag)
	// fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	res, err := client.Do(req)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.OraResponse
	if err = json.Unmarshal(body, &obj); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Get PartName
	var data []models.OraStock
	for _, i := range obj.Data {
		var p models.Part
		if err = db.Where("title=?", i.PartNo).First(&p).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		i.Slug = p.Slug
		i.PartName = p.Description
		data = append(data, i)
	}

	r.Message = obj.Message
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateStockByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteStockByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GetAllStockByShelve(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	method := "GET"
	tag := "C"
	if c.Query("tag") != "" {
		tag = c.Query("tag")
	}

	url := fmt.Sprintf("%s/shelve/%s?tag=%s", configs.API_TRIGGER_URL, c.Params("shelve_no"), tag)
	// fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	res, err := client.Do(req)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.OraResponse
	if err = json.Unmarshal(body, &obj); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Get PartName
	var data []models.OraStock
	for _, i := range obj.Data {
		var p models.Part
		if err = db.Where("title=?", i.PartNo).First(&p).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		i.Slug = p.Slug
		i.PartName = p.Description
		data = append(data, i)
	}

	r.Message = obj.Message
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GetAllStockBySerialNo(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	url := fmt.Sprintf("%s/serial_no/%s", configs.API_TRIGGER_URL, strings.ToUpper(c.Params("serial_no")))
	// fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	res, err := client.Do(req)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.OraResponse
	if err = json.Unmarshal(body, &obj); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Get PartName
	var data []models.OraStock
	for _, i := range obj.Data {
		var p models.Part
		if err = db.Where("title=?", i.PartNo).First(&p).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		i.Slug = p.Slug
		i.PartName = p.Description
		data = append(data, i)
	}

	r.Message = obj.Message
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}
