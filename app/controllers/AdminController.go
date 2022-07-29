package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentModels "github.com/lambda-platform/lambda/agent/models"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"github.com/lambda-platform/lambda/config"
	krudModels "github.com/lambda-platform/lambda/krud/models"
	puzzleModels "github.com/lambda-platform/lambda/models"
	"github.com/lambda-platform/lambda/utils"
)

type Permissions struct {
	DefaultMenu string      `json:"default_menu"`
	Extra       interface{} `json:"extra"`
	MenuID      int         `json:"menu_id"`
	Permissions interface{} `json:"permissions"`
}

func AdminIndex(UseNotify bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		User := agentUtils.AuthUserObject(c)
		Role := agentModels.Role{}
		DB.DB.Where("id = ?", User["role"]).Find(&Role)
		Permissions_ := Permissions{}
		json.Unmarshal([]byte(Role.Permissions), &Permissions_)
		Menu := puzzleModels.VBSchema{}
		DB.DB.Where("id = ?", Permissions_.MenuID).Find(&Menu)
		MenuSchema := new(interface{})
		json.Unmarshal([]byte(Menu.Schema), &MenuSchema)
		Kruds := []krudModels.Krud{}
		DB.DB.Where("deleted_at IS NULL").Find(&Kruds)

		FirebaseConfig := config.LambdaConfig.Notify.FirebaseConfig
		//csrfToken := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
		csrfToken := ""
		return c.Render("admin", map[string]interface{}{
			"UseNotify":                 UseNotify,
			"title":                     config.LambdaConfig.Title,
			"extraStyles":               config.LambdaConfig.ControlPanel.ExtraStyles,
			"extraScripts":              config.LambdaConfig.ControlPanel.ExtraScripts,
			"primaryColor":              config.LambdaConfig.ControlPanel.PrimaryColor,
			"themeColors":               config.LambdaConfig.ControlPanel.ThemeColors,
			"themeMode":                 config.LambdaConfig.ControlPanel.ThemeMode,
			"favicon":                   config.LambdaConfig.Favicon,
			"logo":                      config.LambdaConfig.Logo,
			"logo_light":                config.LambdaConfig.ControlPanel.LogoLight,
			"logo_dark":                 config.LambdaConfig.ControlPanel.LogoDark,
			"brandBtnUrl":               config.LambdaConfig.ControlPanel.BrandBtnURL,
			"permissions":               Permissions_,
			"menu":                      MenuSchema,
			"cruds":                     Kruds,
			"withCrudLog":               config.LambdaConfig.WithCrudLog,
			"User":                      User,
			"data_form_custom_elements": config.LambdaConfig.DataFormCustomElements,
			"firebase_config":           FirebaseConfig,
			"mix":                       utils.Mix,
			"csrfToken":                 csrfToken,
			"lambda_config":             config.LambdaConfig,
		})
	}
}
