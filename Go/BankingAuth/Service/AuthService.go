package Service

import (
	"BankingAuth/Domain"
	"BankingAuth/Dto"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type AuthService interface {
	Login(Dto.LoginRequest) (*string, error)
	Verify(map[string]string) (bool, error)
}

type DefaultAuthService struct {
	Repo            Domain.AuthorizationRepo
	rolePermissions Domain.RolePermissions
}

func (d DefaultAuthService) Login(request Dto.LoginRequest) (*string, error) {

	login, err := d.Repo.FindBy(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (d DefaultAuthService) Verify(urlParams map[string]string) (bool, error) {
	// convert the string token to JWT struct
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return false, err
	} else {
		/*
		   Checking the validity of the token, this verifies the expiry
		   time and the signature of the token
		*/
		if jwtToken.Valid {
			// type cast the token claims to jwt.MapClaims
			mapClaims := jwtToken.Claims.(jwt.MapClaims)
			// converting the token claims to Claims struct
			if claims, err := Domain.BuildClaimsFromJwtMapClaims(mapClaims); err != nil {
				return false, err
			} else {
				/* if Role if user then check if the account_id and customer_id
				   coming in the URL belongs to the same token
				*/
				if claims.IsUserRole() {
					if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
						return false, nil
					}
				}
				// verify of the role is authorized to use the route
				isAuthorized := d.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
				return isAuthorized, nil
			}
		} else {
			return false, errors.New("invalid token")
		}
	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		log.Println("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func NewAuthService(repo Domain.AuthorizationRepo) DefaultAuthService {
	return DefaultAuthService{
		Repo:            repo,
		rolePermissions: Domain.GetRolePermissions(),
	}
}
