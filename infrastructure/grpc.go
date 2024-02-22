package infrastructure

import (
	"context"
	"crypto/tls"
	"errors"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GPRCClientConnConfig struct {
	TargetAddr   string
	TLSConfig    *tls.Config // Optional TLS configuration
	MaxIdleConns int
	DialTimeout  time.Duration
}

type GRPCConnPool struct {
	cfg       *GPRCClientConnConfig
	mu        sync.Mutex
	conns     []*grpc.ClientConn
	idleConns map[string]chan *grpc.ClientConn
	closed    bool
}

func NewGRPCConnectionPool(cfg *GPRCClientConnConfig) (*GRPCConnPool, error) {
	if cfg == nil {
		return nil, errors.New("client connection config required")
	}

	pool := &GRPCConnPool{
		cfg:       cfg,
		conns:     make([]*grpc.ClientConn, 0),
		idleConns: make(map[string]chan *grpc.ClientConn),
		closed:    false,
	}

	return pool, nil
}

func (pool *GRPCConnPool) Get(ctx context.Context) (*grpc.ClientConn, error) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if pool.closed {
		return nil, errors.New("connection pool closed")
	}

	target := pool.cfg.TargetAddr
	idleConnCh, ok := pool.idleConns[target]

	if ok && len(idleConnCh) > 0 {
		select {
		case conn := <-idleConnCh:
			// Return existing idle connection
			return conn, nil
		case <-ctx.Done():
			// Context canceled, return error
			return nil, ctx.Err()
		}
	}

	// Create a new connection if none are idle or context is canceled
	conn, err := pool.dial(ctx)
	if err != nil {
		return nil, err
	}

	// Add the new connection to the pool
	pool.conns = append(pool.conns, conn)
	return conn, nil
}

func (pool *GRPCConnPool) dial(ctx context.Context) (*grpc.ClientConn, error) {
	var dialOpts []grpc.DialOption

	timeoutCtx, cancel := context.WithTimeout(ctx, pool.cfg.DialTimeout)
	defer cancel()

	if pool.cfg.TLSConfig != nil {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(pool.cfg.TLSConfig)))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.DialContext(timeoutCtx, pool.cfg.TargetAddr, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (pool *GRPCConnPool) Put(conn *grpc.ClientConn) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if pool.closed {
		conn.Close()
		return
	}

	target := pool.cfg.TargetAddr
	idleConnCh, ok := pool.idleConns[target]
	if !ok {
		idleConnCh = make(chan *grpc.ClientConn, pool.cfg.MaxIdleConns)
		pool.idleConns[target] = idleConnCh
	}

	select {
	case idleConnCh <- conn:
	default:
		conn.Close()
	}
}

func (pool *GRPCConnPool) CloseAll() {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if pool.closed {
		return
	}

	pool.closed = true
	for _, conn := range pool.conns {
		conn.Close()
	}

	for target, idleConnCh := range pool.idleConns {
		close(idleConnCh)
		delete(pool.idleConns, target)
	}
}
