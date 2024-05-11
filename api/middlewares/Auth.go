package middlewares

import (
	"job-portal-project/api/securities"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Inisialisasi konfigurasi logger
func init() {
	logger.Formatter = &logrus.JSONFormatter{} // Ubah formatter sesuai kebutuhan
	logger.Level = logrus.InfoLevel            // Ubah level log sesuai kebutuhan
}

func JWTAndRBACMiddleware(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate JWT token
		err := securities.GetAuthentication(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check RBAC authorization
		claims := r.Context().Value("claims").(jwt.MapClaims)
		userRole := claims["role"].(string)
		if !contains(allowedRoles, userRole) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}
