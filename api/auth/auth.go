package auth

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

const (
	accessTokenCookieName = "access-token"
)

var (
	JwtSecretKey = os.Getenv("SECRET_KEY")
)

type Auth struct {
	User dao.User
}

func (a *Auth) Login(c echo.Context) error {

	userReq := &model.User{}
	err := c.Bind(userReq)
	if err != nil {
		return err
	}

	// Check for user in db
	userDb, err := a.User.GetByEmailAndPassword(userReq.Email, userReq.Password)
	if err != nil {
		return err
	}

	// Throw unauthorized error
	if userReq.Email != userDb.Email || userReq.Password != userDb.Password {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userDb.ID
	claims["role"] = userDb.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// JWTErrorChecker will be executed when user try to access a protected path.
func (a *Auth) JWTErrorChecker(err error, c echo.Context) error {
	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
}

//// TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
//func (a *Auth) TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		// If the user is not authenticated (no user token data in the context), don't do anything.
//		if c.Get("user") == nil {
//			return next(c)
//		}
//		// Gets user token from the context.
//		u := c.Get("user").(*jwt.Token)
//
//		claims := u.Claims.(*jwt.Claims)
//
//		// We ensure that a new token is not issued until enough time has elapsed
//		// In this case, a new token will only be issued if the old token is within
//		// 15 mins of expiry.
//		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
//			// Gets the refresh token from the cookie.
//			rc, err := c.Cookie(refreshTokenCookieName)
//			if err == nil && rc != nil {
//				// Parses token and checks if it valid.
//				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
//					return []byte(GetRefreshJWTSecret()), nil
//				})
//				if err != nil {
//					if err == jwt.ErrSignatureInvalid {
//						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
//					}
//				}
//
//				if tkn != nil && tkn.Valid {
//					// If everything is good, update tokens.
//					_ = GenerateTokensAndSetCookies(&user.User{
//						Name:  claims.Name,
//					}, c)
//				}
//			}
//		}
//
//		return next(c)
//	}
//}
