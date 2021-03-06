// haz-sc3-producer sends SeisComPML file to AWS SNS as Haz JSON messages.
//   * watches a spool directory for SeisComPML files.
//   * reads files and converts them to a Quake.
//   * checks the Quake quality.
//   * converts the Quake to a JSON Haz message and sends it to an AWS SNS topic.
//   * sends a periodic HeartBeat message to the SNS topic.
package main

import (
	"fmt"
	"github.com/GeoNet/haz/msg"
	"github.com/GeoNet/haz/sns"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	sn          sns.SNS
	sc3SpoolDir = os.Getenv("SC3_SPOOL_DIR")
	sc3Site     = os.Getenv("SC3_SITE")
	heartBeatId = os.Getenv("HEARTBEAT_SERVICE_ID")
)

// main kicks off SeisComPML processing and HeartBeat generation.
func main() {
	var err error
	sn, err = sns.Init()
	if err != nil {
		log.Fatalf("ERROR AWS SNS config: %s", err.Error())
	}

	log.Print("starting message listner")

	go heartBeat()
	sc3ml()
}

// sc3 for processing SeisComPML files.
type sc3 struct {
	msg.Quake
	f string
}

// sc3ml polls the spool dir for SeisComPML files and runs Process on them.
func sc3ml() {
	inter := time.Duration(1) * time.Second

	for {
		files, err := ioutil.ReadDir(sc3SpoolDir)
		if err != nil {
			log.Printf("WARN: %s", err.Error())
		}

		for _, fi := range files {
			if strings.HasSuffix(fi.Name(), ".xml") {
				s := &sc3{f: fi.Name()}
				if !msg.Process(s) {
					if err := os.Remove(sc3SpoolDir + "/" + s.f); err != nil {
						log.Println(err)
					}
				}
			}
		}

		time.Sleep(inter)
	}
}

// Process processes SeisComPML files.  Converts them to a msg.Quake, checks the quality, sends them
// to an AWS SNS topic as a msg.Haz encoded as JSON.
func (s *sc3) Process() bool {
	s.Quake = msg.ReadSC3ML07(sc3SpoolDir + "/" + s.f)
	if s.Err() != nil {
		log.Println(s.Err())
		return false
	}

	s.Quake.Site = sc3Site
	s.RxLog()

	if !s.Publish() {
		log.Println(s.Err())
		return false
	}

	h := msg.Haz{Quake: &s.Quake}
	b, err := h.Encode()
	if err != nil {
		s.SetErr(fmt.Errorf("WARN: not sending %s - encoding err %s.", s.f, err.Error()))
		return false
	}

	s.TxLog()

	err = sn.Publish(msg.Raw{Body: string(b)}, 3)
	if err != nil {
		s.SetErr(err)
		return true
	}

	return false
}

// hb for heartbeat messages
type hb struct {
	msg.Haz
}

// heartBeat periodically sends a heartbeat message.
func heartBeat() {
	inter := time.Duration(15) * time.Second

	for {
		h := &hb{msg.Haz{HeartBeat: &msg.HeartBeat{ServiceID: heartBeatId, SentTime: time.Now().UTC()}}}
		msg.Process(h)
		time.Sleep(inter)
	}
}

// Process sends msg.HeartBeat to an AWS SNS topic as a msg.Haz encoded as JSON.
func (h *hb) Process() bool {
	b, err := h.Encode()
	if err != nil {
		h.SetErr(err)
		return false
	}

	h.HeartBeat.TxLog()

	err = sn.Publish(msg.Raw{Body: string(b)}, 0)
	if err != nil {
		h.SetErr(err)
	}

	return false
}
