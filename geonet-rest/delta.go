package main

import (
	"bytes"
	"github.com/GeoNet/weft"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

// sensor geojson that has a corresponding *.geojson file in S3.
var sensorGeoJSON = map[string]bool{
	"accelerometer": true,
	"barometer": true,
	"broadbandseismometer": true,
	"buildingarraysensor": true,
	"hydrophone": true,
	"microphone": true,
	"pressuresensor": true,
	"shortperiodboreholeseismometer": true,
	"shortperiodseismometer": true,
	"strongmotionsensor": true,
	"gpsantenna": true,
	"metsensor": true,
}

// Proxies sensor *.geojson from S3
func sensorV2(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	if res := weft.CheckQuery(r, []string{"type"}, []string{}); !res.Ok {
		return res
	}

	t := r.URL.Query().Get("type")
	if !sensorGeoJSON[t] {
		return &weft.NotFound
	}

	params := &s3.GetObjectInput{
		Key:    aws.String("delta/" + t +".geojson"),
		Bucket: aws.String("api.geonet.org.nz"),
	}

	res, err := s3Client.GetObject(params)
	if err != nil {
		return weft.InternalServerError(err)
	}
	defer res.Body.Close()

	_, err = b.ReadFrom(res.Body)
	if err != nil {
		return weft.InternalServerError(err)
	}

	h.Set("Content-Type", V2GeoJSON)
	h.Set("Surrogate-Control", maxAge3600)
	return &weft.StatusOK
}
