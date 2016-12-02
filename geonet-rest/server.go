package main

import (
	"github.com/GeoNet/haz/database"
	"github.com/GeoNet/weft"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
)

// For setting Cache-Control and Surrogate-Control headers.
const (
	maxAge10    = "max-age=10"
	maxAge300   = "max-age=300"
	maxAge3600   = "max-age=3600"
	maxAge86400 = "max-age=86400"
)

const (
	V1GeoJSON = "application/vnd.geo+json;version=1"
	V1JSON    = "application/json;version=1"
	V2GeoJSON = "application/vnd.geo+json;version=2"
	V2JSON    = "application/json;version=2"
	JSON      = "application/json"
	protobuf  = "application/x-protobuf"
)

// These are for CAP format and Atom which is not versioned by Accept.
const (
	CAP  = "application/cap+xml"
	Atom = "application/xml"
)

const (
	ErrContent  = "text/plain; charset=utf-8"
	HtmlContent = "text/html; charset=utf-8"
)

const (
	apiS3   = "api.geonet.org.nz"
	sc3mlS3 = "http://seiscompml07.s3-website-ap-southeast-2.amazonaws.com/"
)

var (
	db       database.DB
	client   *http.Client
	s3Client *s3.S3
)

// main connects to the database, sets up request routing, and starts the http server.
func main() {
	var err error
	db, err = database.InitPG()
	if err != nil {
		log.Println("Problem with DB config.")
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("ERROR: problem pinging DB - is it up and contactable? 500s will be served")
	}

	// create an http client to share.
	timeout := time.Duration(5 * time.Second)
	client = &http.Client{
		Timeout: timeout,
	}

	var sess *session.Session
	if sess, err = session.NewSession(); err != nil {
		log.Printf("ERROR creating AWS session 500s will be served %s", err)
	}

	// Credentials can come from environment var (e.g., for dev)
	// or from the role provider when running in EC2.
	// expire the role creds early to force updates.
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&ec2rolecreds.EC2RoleProvider{
				Client: ec2metadata.New(sess),
				ExpiryWindow: time.Second * 30,
			},
		})

	s3Client = s3.New(sess, &aws.Config{Credentials: creds})

	log.Println("starting server")
	http.Handle("/", handler())
	log.Fatal(http.ListenAndServe(":"+os.Getenv("WEB_SERVER_PORT"), nil))
}

// handler creates a mux and wraps it with default handlers.  Separate function to enable testing.
func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	return inbound(mux)
}

func inbound(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO this is a browser cache directive and does not make a lot
		// of sense for an API.
		w.Header().Set("Cache-Control", maxAge10)
		switch r.Method {
		case "GET":
			// Routing is based on Accept query parameters
			// e.g., version=1 in application/json;version=1
			// so caching must Vary based on Accept.
			w.Header().Set("Vary", "Accept")

			h.ServeHTTP(w, r)
		default:
			weft.Write(w, r, &weft.MethodNotAllowed)
			weft.MethodNotAllowed.Count()
			return
		}
	})
}
