package main

//mock node, no TCP network but just channels

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"strconv"

// 	"singula/node/chain"
// 	"singula/node/node"

// 	"singula.finance/netio"
// )

// func NewNodeMock() (*node.TCPNode, error) {
// 	return &node.TCPNode{
// 		//addr:          addr,
// 		accepting:     false,
// 		ConnectedChan: make(chan net.Conn),
// 		Loglevel:      LOGLEVEL_ON,
// 	}, nil
// }

// func (t *TCPNode) handleConnectionMock(mgr *chain.ChainManager, ntchan netio.Ntchan) {
// 	//tr := 100 * time.Millisecond
// 	//defer ntchan.Conn.Close()
// 	//t.log(fmt.Sprintf("handleConnection"))

// 	//netio.NetConnectorSetup(ntchan)
// 	netio.MockNetConnectorSetupEcho(ntchan)

// 	//go RequestHandlerTelMock(t, ntchan)

// 	//go netio.WriteLoop(ntchan, 100*time.Millisecond)

// }

// //handle requests in telnet style i.e. string encoding
// // func RequestHandlerTelMock(t *TCPNode, ntchan netio.Ntchan) {
// // 	for {
// // 		msg_string := <-ntchan.REQ_in
// // 		t.log(fmt.Sprintf("handle %s ", msg_string))

// // 		reply_msg := "out"
// // 		ntchan.REP_out <- reply_msg
// // 	}
// // }

// //handle new connection
// func (t *TCPNode) HandleConnectTCPMock() {

// 	//TODO! hearbeart, check if peers are alive
// 	//TODO! handshake

// 	for {
// 		newpeerConn := <-t.ConnectedChan
// 		strRemoteAddr := newpeerConn.RemoteAddr().String()
// 		//t.log(fmt.Sprintf("accepted conn %v %v", strRemoteAddr, t.accepting))
// 		//t.log(fmt.Sprintf("new peer %v ", newpeerConn))
// 		// log.Println("> ", t.Peers)
// 		// log.Println("# peers ", len(t.Peers))
// 		Verbose := true
// 		ntchan := netio.ConnNtchan(newpeerConn, "server", strRemoteAddr, Verbose)

// 		p := netio.Peer{Address: strRemoteAddr, NodePort: t.NodePort, NTchan: ntchan}
// 		t.Peers = append(t.Peers, p)

// 		//go t.handleConnection(t.Mgr, ntchan)
// 		go t.handleConnectionMock(t.Mgr, ntchan)

// 		//conn.Close()

// 	}
// }

// func runNodeMock(t *TCPNode) {

// 	//setupLogfile()
// 	log.Println("run node")
// 	//t.log(fmt.Sprintf("run node"))

// 	//t.log(fmt.Sprintf("run node on port: %d", t.Config.NodePort))

// 	// 	//if file exists read the chain

// 	// create block every blocktime sec

// 	// if t.Config.DelgateEnabled {
// 	// 	//go utils.DoEvery(, chain.MakeBlock(mgr, blockTime))

// 	// 	//TODO!
// 	// 	//go chain.MakeBlockLoop(t.Mgr, blocktime)
// 	// }

// 	go t.HandleConnectTCP()

// 	t.RunTCP()
// }

// func runAllMock(config Config) {

// 	log.Println("runNodeAll with config ", config)

// 	node, err := NewNode()
// 	node.Config = config
// 	node.addr = ":" + strconv.Itoa(node.Config.NodePort)
// 	//node.setupLogfile()

// 	//node.log(fmt.Sprintf("PeerAddresses: %v", node.Config.PeerAddresses))

// 	mgr := chain.CreateManager()
// 	node.Mgr = &mgr

// 	//TODO signatures of genesis
// 	node.Mgr.InitAccounts()

// 	//node.initSyncChain(config)

// 	if err != nil {
// 		//node.log(fmt.Sprintf("error creating TCP server"))
// 		return
// 	}

// 	log.Println("run node")
// 	go runNode(node)

// }

// func runNodeWithConfigMock() {

// 	// conffile := "conf.edn"
// 	// log.Println("config file ", conffile)

// 	// if _, err := os.Stat(conffile); os.IsNotExist(err) {
// 	// 	log.Println("config file does not exist. create a file named ", conffile)
// 	// 	return
// 	// }

// 	//config := LoadConfiguration(conffile)
// 	config := getConfig()
// 	log.Println("config ", config)
// 	//log.Println("DelegateName ", config.DelegateName)
// 	//log.Println("CreateGenesis ", config.CreateGenesis)
// 	node, _ := NewNode()
// 	node.Config = *config

// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, os.Interrupt)

// 	go runNode(node) //runAll(config)

// 	<-quit
// 	// log.Println("Got quit signal: shutdown node ...")
// 	// signal.Reset(os.Interrupt)

// 	log.Println("node exiting")

// 	//handle shutdown should never happen, need restart on OS level and error handling

// }

// func main() {

// 	//runNodeWithConfigMock()
// 	fmt.Println("run mock")
// }
