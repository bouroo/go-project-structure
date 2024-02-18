package infrastructure

import (
	"errors"

	"github.com/gocql/gocql"
)

type ScyllaDBConn struct {
	cluster *gocql.ClusterConfig
	Session *gocql.Session
}

type ScyllaDBOptions struct {
	Hosts []string
}

func NewScyllaDBConn(opts ScyllaDBOptions) (scyllaDBConn *ScyllaDBConn, err error) {
	cluster := gocql.NewCluster(opts.Hosts...)
	if cluster == nil {
		err = errors.New("cluster is nil")
		return
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return
	}
	if session == nil {
		err = errors.New("session is nil")
		return
	}
	return &ScyllaDBConn{Session: session}, nil
}

func (s *ScyllaDBConn) GetSession() (session *gocql.Session, err error) {
	if s.Session == nil || s.Session.Closed() {
		s.Session, err = s.cluster.CreateSession()
	}
	return s.Session, err
}
