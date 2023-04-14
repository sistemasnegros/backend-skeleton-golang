package authService

import (
	authDTO "backend-skeleton-golang/auth/app/dto"
	"os"
	"testing"

	smtpServiceMock "backend-skeleton-golang/test/mocks/commons/app/services/smtp-service"
	usersRepoMock "backend-skeleton-golang/test/mocks/users/infra/mongodb/repo"
	usersDomain "backend-skeleton-golang/users/domain"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	smtp    smtpServiceMock.ISmtpService
	repo    usersRepoMock.IUsers
	service IService
}

func (suite *AuthServiceTestSuite) SetupTest() {

	suite.smtp = smtpServiceMock.ISmtpService{}
	suite.repo = usersRepoMock.IUsers{}
	suite.service = New(&suite.smtp, &suite.repo)

}

func (suite *AuthServiceTestSuite) TestLoginSuccess() {
	os.Setenv("JWT_EXP", "24")
	defer os.Unsetenv("JWT_EXP")

	password := "perrogordo"
	passwordHash := "$2a$14$v/KHEfcZN9Qt20J7gQ6T0.aTer6hFSinPQpGCElJsf7UGEOc14JHi"
	suite.repo.On("FindOne", mock.Anything).Return(&usersDomain.User{Id: "success", Password: passwordHash}, nil)

	code, _ := suite.service.Login(&authDTO.Login{Email: "", Password: password})
	suite.Equal(code, 200)
}

func (suite *AuthServiceTestSuite) TestLoginIvalid() {
	suite.repo.On("FindOne", map[string]interface{}{"email": ""}).Return(&usersDomain.User{}, nil)
	code, _ := suite.service.Login(&authDTO.Login{Email: "", Password: ""})
	suite.Equal(code, 401)
}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
