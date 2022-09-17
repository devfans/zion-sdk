package bridge

import (
	"testing"
	"time"

	br "github.com/polynetwork/bridge-common/chains/bridge"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
)


func TestBrig(t *testing.T) {
	b := br.New("https://bridge.poly.network/v1")
	b.SetExplorer("https://explorer.poly.network/api/v1")

	hash := "0x19eef5345cd0abb70bcfe4db6e3aa0f3cb187014e141ff3564a8f661a81a30c6"

	log.Info("Watch pending for poly bridge tx", "hash", hash)
	// time.Sleep(30 * time.Second) // At least 50s for poly to finish
	log.Info("Watching for poly bridge tx", "hash", hash)

	for {
		res, err := b.FetchTx(&br.CheckTxRequest{Hash: util.LowerHex(hash)})
		if err != nil {
			log.Error("Watch poly tx failure", "hash", hash, "err", err)
			time.Sleep(time.Second)
			continue
		}
		if res.Mchaintx != nil && res.Mchaintx.Txhash != "" {
			log.Info("Poly tx was successful", "poly_hash", res.Mchaintx.Txhash)
			break
		}
		log.Info("Not yet, sleep 10s now")
		time.Sleep(10000 *time.Millisecond)
	}
}