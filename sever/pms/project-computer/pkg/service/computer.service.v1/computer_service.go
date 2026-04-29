package computer_service_v1

import (
	"pms.com/project-computer/internal/dao"
	"pms.com/project-computer/internal/database/tran"
	"pms.com/project-computer/internal/repo"
	"pms.com/project-grpc/computer/action"
)

type ComputerService struct {
	action.UnimplementedComputerServiceServer
	cache repo.Cache

	transaction tran.Transaction
}

func New() *ComputerService {
	return &ComputerService{
		cache: dao.Rc,
	}
}
