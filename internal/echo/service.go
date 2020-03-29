package echo

import (
	"context"

	"github.com/hpifu/tpl-go-thrift/api/echoservice"
	"github.com/sirupsen/logrus"
)

type Service struct {
	infoLog   *logrus.Logger
	warnLog   *logrus.Logger
	accessLog *logrus.Logger
}

func NewService() *Service {
	return &Service{
		infoLog:   logrus.New(),
		warnLog:   logrus.New(),
		accessLog: logrus.New(),
	}
}

func (s *Service) SetLogger(infoLog, warnLog, accessLog *logrus.Logger) {
	s.infoLog = infoLog
	s.warnLog = warnLog
	s.accessLog = accessLog
}

func (s *Service) Echo(ctx context.Context, req *echoservice.EchoReq) (*echoservice.EchoRes, error) {
	return &echoservice.EchoRes{
		Message: req.Message,
	}, nil
}
