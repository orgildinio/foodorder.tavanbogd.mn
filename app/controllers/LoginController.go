package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/agent/models"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"github.com/lambda-platform/lambda/config"
	"io/ioutil"
	appModel "lambda/app/models"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type LoginData struct {
	Token  string      `json:"token"`
	Path   string      `json:"path"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
type User struct {
	Login    string `json:"login" xml:"login" form:"login" query:"login"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}
type UserData struct {
	Id    int64
	Login string
	Role  int64
}
type UserUUIDData struct {
	Id    string
	Login string
	Role  int64
}
type jwtClaims struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Role  int64  `json:"role"`
	jwt.StandardClaims
}
type jwtUUIDClaims struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	Role  int64  `json:"role"`
	jwt.StandardClaims
}
type Permissions struct {
	DefaultMenu string      `json:"default_menu"`
	Extra       interface{} `json:"extra"`
	MenuID      int         `json:"menu_id"`
	Permissions interface{} `json:"permissions"`
}
type LoginUser struct {
	Login    string `json:"login" xml:"login" form:"login" query:"login"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}
type Unauthorized struct {
	Error  string `json:"error"`
	Status bool   `json:"status"`
}

//type LoginRes struct {
//	Token string `json:"token"`
//}

type LoginRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func Login(c *fiber.Ctx) error {
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
			Error:  "Username & password required",
			Status: false,
		})
	}

	foundUser := agentUtils.AuthUserObjectByLogin(u.Login)

	if len(foundUser) == 0 {
		if len(foundUser) == 0 {

			return UbERPMobileLogin(u.Login, u.Password, c)

		}

	}

	//password, err := Hash(u.Password)
	//password_check1 := IsSame(password, foundUser.Password)
	if agentUtils.IsSame(u.Password, foundUser["password"].(string)) {

		var roleID int64 = 0
		var userID int64 = 0
		var userUUID string = ""

		if reflect.TypeOf(foundUser["id"]).String() == "string" {
			if config.Config.SysAdmin.UUID {
				userUUID = foundUser["id"].(string)
			} else {
				i, err := strconv.ParseInt(foundUser["id"].(string), 10, 64)
				if err != nil {
					panic(err)
				}
				userID = i
			}

		} else {
			userID = foundUser["id"].(int64)
		}
		if reflect.TypeOf(foundUser["role"]).String() == "string" {
			i, err := strconv.ParseInt(foundUser["role"].(string), 10, 64)
			if err != nil {
				panic(err)
			}
			roleID = i
		} else {
			roleID = foundUser["role"].(int64)
		}

		if config.Config.SysAdmin.UUID {
			// create jwt token
			token, err := createUUIDJwtToken(UserUUIDData{Id: userUUID, Login: foundUser["login"].(string), Role: roleID})
			if err != nil {
				//log.Println("Error Creating JWT token", err)
				return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
					Error:  "Unauthorized",
					Status: false,
				})
			}

			cookie := new(fiber.Cookie)
			cookie.Name = "token"
			cookie.Path = "/"
			cookie.Value = token
			cookie.Expires = time.Now().Add(time.Hour * time.Duration(config.Config.JWT.Ttl))
			//cookie.HttpOnly = true
			//cookie.Secure = true

			delete(foundUser, "password")

			foundUser["jwt"] = token

			c.Cookie(cookie)

			return c.Status(fiber.StatusOK).JSON(models.LoginData{
				Token:  token,
				Path:   checkRole(roleID),
				Status: true,
				Data:   foundUser,
			})
		} else {
			// create jwt token
			token, err := createJwtToken(UserData{Id: userID, Login: foundUser["login"].(string), Role: roleID})
			if err != nil {
				//log.Println("Error Creating JWT token", err)
				return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
					Error:  "Unauthorized",
					Status: false,
				})
			}

			cookie := new(fiber.Cookie)
			cookie.Name = "token"
			cookie.Path = "/"
			cookie.Value = token
			cookie.Expires = time.Now().Add(time.Hour * time.Duration(config.Config.JWT.Ttl))

			delete(foundUser, "password")

			foundUser["jwt"] = token

			c.Cookie(cookie)

			return c.Status(fiber.StatusOK).JSON(models.LoginData{
				Token:  token,
				Path:   checkRole(roleID),
				Status: true,
				Data:   foundUser,
			})
		}

	}

	return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
		Error:  "Unauthorized",
		Status: false,
	})

}
func createUUIDJwtToken(user UserUUIDData) (string, error) {
	// Set custom claims
	claims := jwt.MapClaims{
		"id":    user.Id,
		"login": user.Login,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * time.Duration(config.Config.JWT.Ttl)).Unix(),
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
func checkRole(role int64) string {
	for _, r := range config.LambdaConfig.RoleRedirects {
		if role == r.RoleID {
			return r.URL
		}
	}

	RoleData := map[string]interface{}{}

	jsonFile, err := os.Open("lambda/role_" + fmt.Sprintf("%v", role) + ".json")

	defer jsonFile.Close()
	if err != nil {
		foundRole := models.Role{}
		DB.DB.Where("id = ?", role).First(&foundRole)
		if foundRole.Permissions != "" {

			Permissions := models.Permissions{}
			json.Unmarshal([]byte(foundRole.Permissions), &Permissions)
			if Permissions.DefaultMenu != "" {
				return config.LambdaConfig.AppURL + Permissions.DefaultMenu
			}
		}
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &RoleData)

	permissonData, err := json.Marshal(RoleData["permissions"])
	if err != nil {
		return "/auth/login"
	}

	permissions := models.Permissions{}
	json.Unmarshal(permissonData, &permissions)
	return config.LambdaConfig.AppURL + permissions.DefaultMenu

}

func createJwtToken(user UserData) (string, error) {
	// Set custom claims
	claims := jwt.MapClaims{
		"id":    user.Id,
		"login": user.Login,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * time.Duration(config.Config.JWT.Ttl)).Unix(),
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func UbERPMobileLogin(Loging string, Password string, c *fiber.Ctx) error {
	url := "http://202.21.112.3:8082/api/auth-service/login"
	method := "POST"

	str := Loging
	splited := strings.Split(str, `\`)

	if len(splited) >= 2 {

		companyID := strings.ToUpper(splited[0])
		userID := splited[1]

		//fmt.Println("{\"password\": \"" + Password + "\",  \"userid\": \"" + userID + "\", \"companyid\": \"" + companyID + "\"}")

		payload := strings.NewReader("{\"password\": \"" + Password + "\",  \"userid\": \"" + userID + "\", \"companyid\": \"" + companyID + "\"}")

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/json-patch+json")
		//req.Header.Add("Content-Type", "text/plain")

		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		loginRES := LoginRes{}
		errParse := json.Unmarshal([]byte(body), &loginRES)

		if errParse != nil || loginRES.AccessToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
				Error:  "User not found",
				Status: false,
			})
		} else {
			fmt.Println("===========")
			return checkUserByToken(loginRES.AccessToken, c)
		}

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
			Error:  "User not found",
			Status: false,
		})
	}
}

type JWTPayLoad struct {
	UserID     string `json:"userID"`
	EmpID      string `json:"EmpID"`
	CompanyID  string `json:"CompanyID"`
	DepID      string `json:"DepID"`
	PositionID string `json:"positionID"`
	Ipaddress  string `json:"ipaddress"`
	Macaddress string `json:"macaddress"`
	Sub        string `json:"sub"`
	Jti        string `json:"jti"`
	Iat        string `json:"iat"`
	Nbf        int    `json:"nbf"`
	Exp        int    `json:"exp"`
	Iss        string `json:"iss"`
	Aud        string `json:"aud"`
}

func checkUserByToken(token string, c *fiber.Ctx) error {
	splited := strings.Split(token, `.`)
	uDec, _ := base64.RawStdEncoding.DecodeString(splited[1])

	payload := JWTPayLoad{}
	if err := json.Unmarshal(uDec, &payload); err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(splited[1])
	//fmt.Println(string(uDec))
	//fmt.Println(payload.EmpID)

	url := "http://202.21.112.3:8082/api/hrms/hr/employee/get?EmployeeID=" + payload.EmpID
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	userInfo := appModel.UserInfo{}
	err := json.Unmarshal([]byte(body), &userInfo)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
			Error:  "User not found",
			Status: false,
		})

	} else {

		//fmt.Println(len(userInfo.Retdata))
		//fmt.Println(len(userInfo.Retdata))
		//fmt.Println(len(userInfo.Retdata))
		if len(userInfo.Retdata) >= 1 {

			return erpKhorooByToken(c, userInfo, true)

		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
				Error:  "User not found",
				Status: false,
			})
		}

	}
	return c.Status(fiber.StatusUnauthorized).JSON(models.Unauthorized{
		Error:  "User not found",
		Status: false,
	})

}

func erpKhorooByToken(c *fiber.Ctx, userInfo appModel.UserInfo, isKhoroo bool) error {

	foundUser := agentUtils.AuthUserObjectByLogin(userInfo.Retdata[0].Regno)
	if len(foundUser) == 0 {

		// role changed
		var role int64 = 3

		// empID added
		newUser := appModel.UserERP{
			Login:          userInfo.Retdata[0].Regno,
			Email:          userInfo.Retdata[0].Email,
			FirstName:      userInfo.Retdata[0].Firstname,
			LastName:       userInfo.Retdata[0].Lastname,
			Phone:          userInfo.Retdata[0].Mobilephone,
			RegisterNumber: userInfo.Retdata[0].Regno,
			Role:           role,
			Birthday:       time.Now(),
			Status:         "1",
			Gender:         "m",
			EmpID:          int(userInfo.Retdata[0].Empid),
			CompanyID:      userInfo.Retdata[0].Companyid,
			CompanyName:    userInfo.Retdata[0].Companyname,
			PosName:        userInfo.Retdata[0].Posname,
			DepName:        userInfo.Retdata[0].Depname,
		}

		DB.DB.Create(&newUser)

		foundUser2 := agentUtils.AuthUserObjectByLogin(newUser.Login)

		token, err := createJwtToken(UserData{Id: newUser.ID, Login: newUser.Login, Role: newUser.Role})
		if err != nil {
			return c.JSON(Unauthorized{
				Error:  "Username & password required",
				Status: false,
			})
		}

		if isKhoroo {
			delete(foundUser2, "password")

			foundUser2["jwt"] = token
			return c.JSON(LoginData{
				Token:  token,
				Path:   "/control#/p/956471dd-88fb-2803-d1ab-302dd63c30c1/0e15ca07-13b5-f9de-aade-4bddd30066fd",
				Status: true,
				Data:   foundUser2,
			})
		}

	} else {
		// create jwt token
		token, err := createJwtToken(UserData{Id: foundUser["id"].(int64), Login: foundUser["login"].(string), Role: foundUser["role"].(int64)})
		if err != nil {
			return c.JSON(Unauthorized{
				Error:  "Username & password required",
				Status: false,
			})
		}

		if isKhoroo {
			delete(foundUser, "password")

			foundUser["jwt"] = token
			return c.JSON(LoginData{
				Token:  token,
				Path:   "",
				Status: true,
				Data:   foundUser,
			})
		}
	}

	return c.JSON(Unauthorized{
		Error:  "Username & password required",
		Status: false,
	})

}
