package controllers

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/lambda-platform/lambda/DB"
    "io/ioutil"
    "lambda/app/models"
)

func Users(c *fiber.Ctx) error {

    users := []models.Users{}
    DB.DB.Find(&users)

    return c.JSON(users)
}

func ReadIcons(c *fiber.Ctx) error {

    files, err := ioutil.ReadDir("public/assets/icons/duotune/")
    if err != nil {
        fmt.Println(err)
    }

    svgIcons := []map[string]interface{}{}
    for _, file := range files {
        if file.IsDir() {
            svgGroup := map[string]interface{}{}

            svgGroup["title"] = "SVG Duotune " + file.Name()

            iconList := []map[string]string{}
            svgs, err2 := ioutil.ReadDir("public/assets/icons/duotune/" + file.Name())
            if err2 != nil {
                fmt.Println(err2)
            } else {
                for _, svg := range svgs {
                    iconList = append(iconList, map[string]string{
                        "url":  "/assets/icons/duotune/" + file.Name() + "/" + svg.Name(),
                        "name": svg.Name(),
                    })
                }
            }
            svgGroup["svgs"] = iconList

            //fmt.Println(file.Name(), file.IsDir())

            svgIcons = append(svgIcons, svgGroup)
        }

    }
    return c.JSON(svgIcons)
}
