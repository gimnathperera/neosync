package v1alpha1_jobservice

import (
	"github.com/nucleuscloud/neosync/backend/internal/nucleusdb"
	v1alpha1_connectionservice "github.com/nucleuscloud/neosync/backend/services/mgmt/v1alpha1/connection-service"
	v1alpha1_useraccountservice "github.com/nucleuscloud/neosync/backend/services/mgmt/v1alpha1/user-account-service"
)

type Service struct {
	cfg                *Config
	db                 *nucleusdb.NucleusDb
	userAccountService *v1alpha1_useraccountservice.Service
	connectionService  *v1alpha1_connectionservice.Service
}

type Config struct{}

func New(
	cfg *Config,
	db *nucleusdb.NucleusDb,
	userAccountService *v1alpha1_useraccountservice.Service,
	connectionService *v1alpha1_connectionservice.Service,
) *Service {

	return &Service{
		cfg:                cfg,
		db:                 db,
		userAccountService: userAccountService,
		connectionService:  connectionService,
	}
}