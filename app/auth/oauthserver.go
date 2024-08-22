package auth

import (
	"github.com/RangelReale/osin"
	mysql "github.com/felipeweb/osin-mysql"
	"github.com/sltet/garage/app/core"
	"github.com/sltet/garage/app/db"
)

type OauthServer struct {
	*mysql.Storage
}

type OauthServerInterface interface {
	SaveToken(data *osin.AccessData) core.DetailedError
	LoadToken(token string) (*osin.AccessData, core.DetailedError)
	SaveClient(id string, secret string, redirectURI string)
}

func NewOauthServer(em db.EntityManagerInterface) *OauthServer {
	database, err := em.Database().DB()
	if err != nil {
		panic(err)
	}
	store := mysql.New(database, "osin_")
	err = store.CreateSchemas()
	if err != nil {
		panic(err)
	}
	return &OauthServer{store}
}

func (s *OauthServer) SaveToken(data *osin.AccessData) core.DetailedError {
	err := s.SaveAccess(data)
	if err != nil {
		return core.NewSimpleError(err.Error())
	}
	return nil
}

func (s *OauthServer) LoadToken(token string) (*osin.AccessData, core.DetailedError) {
	data, err := s.LoadAccess(token)
	if err != nil {
		return nil, core.NewSimpleError(err.Error())
	}
	return data, nil
}

func (s *OauthServer) SaveClient(id string, secret string, redirectURI string) {
	client := s.CreateClientWithInformation(id, secret, redirectURI, nil)
	_ = s.CreateClient(client)
}
