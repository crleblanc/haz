package main

import (
	"bytes"
	"github.com/GeoNet/weft"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

// Proxies marks.geojson from S3
func marksV2(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	if res := weft.CheckQuery(r, []string{}, []string{}); !res.Ok {
		return res
	}

	params := &s3.GetObjectInput{
		Key:    aws.String("delta/marks.geojson"),
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

// Proxies marks.pb from S3
func marksProto(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	if res := weft.CheckQuery(r, []string{}, []string{}); !res.Ok {
		return res
	}

	params := &s3.GetObjectInput{
		Key:    aws.String("delta/marks.pb"),
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

	h.Set("Content-Type", protobuf)
	h.Set("Surrogate-Control", maxAge3600)
	return &weft.StatusOK
}

// Proxies marks.json from S3
func marksV2JSON(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	if res := weft.CheckQuery(r, []string{}, []string{}); !res.Ok {
		return res
	}

	params := &s3.GetObjectInput{
		Key:    aws.String("delta/marks.json"),
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

	h.Set("Content-Type", V2JSON)
	h.Set("Surrogate-Control", maxAge3600)
	return &weft.StatusOK
}