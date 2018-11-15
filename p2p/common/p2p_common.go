package common

import (
	"errors"
	"strconv"
	"strings"

	"github.com/coschain/contentos-go/p2p/depend/core/types"
)

//peer capability
const (
	VERIFY_NODE  = 1 //peer involved in consensus
	SERVICE_NODE = 2 //peer only sync with consensus peer
)

//link and concurrent const
const (
	PER_SEND_LEN        = 1024 * 256 //byte len per conn write
	MAX_BUF_LEN         = 1024 * 256 //the maximum buffer to receive message
	WRITE_DEADLINE      = 5          //deadline of conn write
	REQ_INTERVAL        = 3          //single request max interval in second
	MAX_REQ_RECORD_SIZE = 1000       //the maximum request record size
	MAX_RESP_CACHE_SIZE = 50         //the maximum response cache
)

//msg cmd const
const (
	MSG_CMD_LEN      = 12               //msg type length in byte
	CMD_OFFSET       = 4                //cmd type offet in msg hdr
	CHECKSUM_LEN     = 4                //checksum length in byte
	MSG_HDR_LEN      = 24               //msg hdr length in byte
	MAX_BLK_HDR_CNT  = 500              //hdr count once when sync header
	MAX_INV_HDR_CNT  = 500              //inventory count once when req inv
	MAX_REQ_BLK_ONCE = 16               //req blk count once from one peer when sync blk
	MAX_MSG_LEN      = 30 * 1024 * 1024 //the maximum message length
	MAX_PAYLOAD_LEN  = MAX_MSG_LEN - MSG_HDR_LEN
)

//msg type const
const (
	MAX_ADDR_NODE_CNT = 64 //the maximum peer address from msg
	MAX_INV_BLK_CNT   = 64 //the maximum blk hash cnt of inv msg
)

//info update const
const (
	PROTOCOL_VERSION      = 0     //protocol version
	UPDATE_RATE_PER_BLOCK = 2     //info update rate in one generate block period
	KEEPALIVE_TIMEOUT     = 15    //contact timeout in sec
	DIAL_TIMEOUT          = 6     //connect timeout in sec
	CONN_MONITOR          = 6     //time to retry connect in sec
	CONN_MAX_BACK         = 4000  //max backoff time in micro sec
	MAX_RETRY_COUNT       = 3     //max reconnect time of remote peer
	CHAN_CAPABILITY       = 10000 //channel capability of recv link
	SYNC_BLK_WAIT         = 2     //timespan for blk sync check
)

// The peer state
const (
	INIT        = 0 //initial
	HAND        = 1 //send verion to peer
	HAND_SHAKE  = 2 //haven`t send verion to peer and receive peer`s version
	HAND_SHAKED = 3 //send verion to peer and receive peer`s version
	ESTABLISH   = 4 //receive peer`s verack
	INACTIVITY  = 5 //link broken
)

//cap flag
const (
	HTTP_INFO_FLAG = 0 //peer`s http info bit in cap field
)

//actor const
const (
	ACTOR_TIMEOUT = 5 //actor request timeout in secs
)

//recent contact const
const (
	RECENT_TIMEOUT   = 60
	RECENT_FILE_NAME = "peers.recent"
	RECENT_LIMIT     = 10 //recent contact list limit
)

//PeerAddr represent peer`s net information
type PeerAddr struct {
	Time          int64    //latest timestamp
	Services      uint64   //service type
	IpAddr        [16]byte //ip address
	Port          uint16   //sync port
	ConsensusPort uint16   //consensus port
	ID            uint64   //Unique ID
}

//const channel msg id and type
const (
	VERSION_TYPE           = "version"    //peer`s information
	VERACK_TYPE            = "verack"     //ack msg after version recv
	GetADDR_TYPE           = "getaddr"    //req nbr address from peer
	ADDR_TYPE              = "addr"       //nbr address
	PING_TYPE              = "ping"       //ping  sync height
	PONG_TYPE              = "pong"       //pong  recv nbr height
	GET_DATA_TYPE          = "getdata"    //req data from peer
	BLOCK_TYPE             = "sig_block"      //blk payload
	HASH_TYPE              = "hash"
	REQ_HASH_TYPE          = "req_hash"
	TX_TYPE                = "sig_trx"         //transaction
	NOT_FOUND_TYPE         = "notfound"   //peer can`t find blk according to the hash
	DISCONNECT_TYPE        = "disconnect" //peer disconnect info raise by link
)

type AppendPeerID struct {
	ID uint64 // The peer id
}

type RemovePeerID struct {
	ID uint64 // The peer id
}

type AppendHeaders struct {
	FromID  uint64          // The peer id
	Headers []*types.Header // Headers to be added to the ledger
}

//type AppendBlock struct {
//	FromID    uint64       // The peer id
//	BlockSize uint32       // Block size
//	Block     *types.Block // Block to be added to the ledger
//}

//ParseIPAddr return ip address
func ParseIPAddr(s string) (string, error) {
	i := strings.Index(s, ":")
	if i < 0 {
		return "", errors.New("[p2p]split ip address error")
	}
	return s[:i], nil
}

//ParseIPPort return ip port
func ParseIPPort(s string) (string, error) {
	i := strings.Index(s, ":")
	if i < 0 {
		return "", errors.New("[p2p]split ip port error")
	}
	port, err := strconv.Atoi(s[i+1:])
	if err != nil {
		return "", errors.New("[p2p]parse port error")
	}
	if port <= 0 || port >= 65535 {
		return "", errors.New("[p2p]port out of bound")
	}
	return s[i:], nil
}
