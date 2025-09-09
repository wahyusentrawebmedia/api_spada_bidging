package middleware

import (
	"api/spada/internal/model"
	"api/spada/internal/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var jwtCheckURLUser = "/admin/check_auth"

// JWTCheckMiddleware memvalidasi JWT dengan memanggil endpoint eksternal
func JWTCheckMiddlewareUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		urlAkademikAuth := viper.GetString("URL_AKADEMIK_AUTH")
		cc := utils.NewCustomContext(c)

		token := c.Get("Authorization")
		if token == "" {
			return cc.ErrorResponseUnauthorized("Missing Authorization header")
		}
		token = strings.TrimPrefix(token, "Bearer ")

		// Kirim token ke endpoint eksternal untuk validasi
		client := &http.Client{}
		req, err := http.NewRequest("GET", urlAkademikAuth+jwtCheckURLUser, nil)
		if err != nil {
			return cc.ErrorResponseUnauthorized("JWT check request error")
		}
		req.Header.Set("Authorization", "Bearer "+token)

		curlCmd := fmt.Sprintf(
			`curl -X GET "%s%s" -H "Authorization: Bearer %s"`,
			urlAkademikAuth, jwtCheckURLUser, token,
		)
		fmt.Println("[DEBUG] JWTCheckMiddleware - CURL: %s", curlCmd)

		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			if resp != nil && resp.Body != nil {
				bodyBytes, _ := io.ReadAll(resp.Body)
				fmt.Printf("[DEBUG] JWTCheckMiddleware - Response Body: %s\n", string(bodyBytes))
			}
			return cc.ErrorResponseUnauthorized("Invalid token")
		}

		defer resp.Body.Close()
		var jwtResp model.JWTUserCheckResponse
		if err := utils.DecodeJSON(resp.Body, &jwtResp); err != nil {
			return cc.ErrorResponseUnauthorized("Failed to decode JWT check response")
		}

		if jwtResp.Data.IDPerguruanTinggi == 0 {
			return cc.ErrorResponseUnauthorized("User ini tidak terdapat di perguruan tinggi manapun")
		}

		c.Locals("id_perguruan_tinggi", strconv.Itoa(jwtResp.Data.IDPerguruanTinggi))
		c.Locals("username", jwtResp.Data.Username)

		if c.Locals("id_perguruan_tinggi") == "" {
			return cc.ErrorResponseUnauthorized("id_perguruan_tinggi tidak ditemukan di token")
		}

		if c.Locals("username") == "" {
			return cc.ErrorResponseUnauthorized("username tidak ditemukan di token")
		}

		cc.SetLocalsParameter()

		return c.Next()
	}
}
