package node

import (
	"context"
	"fmt"
	"net/http"

	"github.com/evanchen7/blockchain/database"
)

const (
	DefaultIP      = "127.0.0.1"
	DefaultHTTPort = 8080

	endpointStatus                = "/node/status"
	endpointSync                  = "/node/sync"
	endpointSyncQueryKeyFromBlock = "fromBlock"

	endpointAddPeer             = "/node/peer"
	endpointAddPeerQueryKeyIP   = "ip"
	endpointAddPeerQueryKeyPort = "port"
)

type PeerNode struct {
	IP          string `json:"ip"`
	Port        uint64 `json:"port"`
	IsBootstrap bool   `json:"is_bootstrap"`
	connected   bool
}

type Node struct {
	dataDir string
	ip      string
	port    uint64

	state      *database.State
	knownPeers map[string]PeerNode
}

func (pn PeerNode) TcpAddress() string {
	return fmt.Sprintf("%s:%d", pn.IP, pn.Port)
}

func New(dataDir string, ip string, port uint64, bootstrap PeerNode) *Node {
	knownPeers := make(map[string]PeerNode)
	knownPeers[bootstrap.TcpAddress()] = bootstrap

	return &Node{
		dataDir:    dataDir,
		ip:         ip,
		port:       port,
		knownPeers: knownPeers,
	}
}

func NewPeerNode(ip string, port uint64, isBootStrap bool, connected bool) PeerNode {
	return PeerNode{
		IP:          ip,
		Port:        port,
		IsBootstrap: isBootStrap,
		connected:   connected,
	}
}

func (n *Node) Run(dataDir string) error {
	ctx := context.Background()
	fmt.Println(fmt.Sprintf("Listening on: %s:%d", n.ip, n.port))

	state, err := database.NewStateFromDisk(n.dataDir)
	if err != nil {
		return nil
	}
	defer state.Close()

	n.state = state

	go n.sync(ctx)

	http.HandleFunc("/balances/list", func(w http.ResponseWriter, r *http.Request) {
		listBalancesHandler(w, r, state)
	})

	http.HandleFunc("/tx/add", func(w http.ResponseWriter, r *http.Request) {
		txAddHandler(w, r, state)
	})

	http.HandleFunc(endpointSync, func(w http.ResponseWriter, r *http.Request) {
		syncHandler(w, r, n)
	})

	http.HandleFunc(endpointStatus, func(w http.ResponseWriter, r *http.Request) {
		statusHandler(w, r, n)
	})

	http.HandleFunc(endpointAddPeer, func(w http.ResponseWriter, r *http.Request) {
		addPeerHandler(w, r, n)
	})

	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	return http.ListenAndServe(fmt.Sprintf("%s:%d", n.ip, n.port), nil)
}

func (n *Node) AddPeer(peer PeerNode) {
	n.knownPeers[peer.TcpAddress()] = peer
}

func (n *Node) RemovePeer(peer PeerNode) {
	delete(n.knownPeers, peer.TcpAddress())
}

func (n *Node) IsKnownPeer(peer PeerNode) bool {
	if peer.IP == n.ip && peer.Port == n.port {
		return true
	}

	_, isKnownPeer := n.knownPeers[peer.TcpAddress()]

	return isKnownPeer
}

