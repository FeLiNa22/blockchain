package network

type Node struct {
	tp            byte     // transport protocol: 0 = tcp, 1 = udp, 2 = quik
	ip            [16]byte // ip address: supports ipv4 and ipv6
	port          uint16
}

var DEFAULT_PORT uint16 = 24069 // I'm immature
var DEFAULT_TYPE uint8 = 0

/*
	Start the node, setting it up on the blockchain network
	1. Firsts opens a tcp DEFAULT_PORT to listen for messages
	2. Broadcasts to all known nodes that this machine is on the network
	3. Other nodes will vote if the node is safe
	4. If accepted into the network start synchronisation
*/
func start() {
}


/*
	Stops the node from operating
	1. Just close the connection straight away
*/
func stop() {
}


/*
	sends ping request to node
*/
func check_available(node *Node) {

}

/*
	Register listener function that executes
*/



