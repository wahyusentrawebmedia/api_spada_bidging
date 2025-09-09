package middleware

import (
	"api/spada/internal/model"
	"api/spada/internal/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var jwtCheckURL = "/spada/sessions"

// JWTCheckMiddleware memvalidasi JWT dengan memanggil endpoint eksternal
func JWTCheckMiddleware() fiber.Handler {
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
		req, err := http.NewRequest("GET", urlAkademikAuth+jwtCheckURL, nil)
		if err != nil {
			return cc.ErrorResponseUnauthorized("JWT check request error")
		}
		req.Header.Set("Authorization", "Bearer "+token)

		curlCmd := fmt.Sprintf(
			`curl -X GET "%s%s" -H "Authorization: Bearer %s"`,
			urlAkademikAuth, jwtCheckURL, token,
		)
		fmt.Println("[DEBUG] JWTCheckMiddleware - CURL: %s", curlCmd)

		resp, err := client.Do(req)

		if err != nil || resp.StatusCode != http.StatusOK {
			return cc.ErrorResponseUnauthorized("Invalid token")
		}

		defer resp.Body.Close()
		var jwtResp model.JWTCheckResponse
		if err := utils.DecodeJSON(resp.Body, &jwtResp); err != nil {
			// if os.Getenv("DEBUG_HTTP") == "1" {
			fmt.Printf("[DEBUG] Failed to decode JWT check response: %v\n", err)
			// }
			return cc.ErrorResponseUnauthorized("Failed to decode JWT check response")
		}

		if jwtResp.Data.IDPerguruanTinggi == 0 {
			// if os.Getenv("DEBUG_HTTP") == "1" {
			fmt.Println("[DEBUG] User tidak terdapat di perguruan tinggi manapun")
			// }
			return cc.ErrorResponseUnauthorized("User ini tidak terdapat di perguruan tinggi manapun")
		}

		c.Locals("id_perguruan_tinggi", strconv.Itoa(jwtResp.Data.IDPerguruanTinggi))

		if c.Locals("id_perguruan_tinggi") == "" {
			// if os.Getenv("DEBUG_HTTP") == "1" {
			fmt.Println("[DEBUG] id_perguruan_tinggi tidak ditemukan di token")
			// }
			return cc.ErrorResponseUnauthorized("id_perguruan_tinggi tidak ditemukan di token")
		}

		cc.SetLocalsParameter()

		return c.Next()
	}
}
