package main

import (
	"github.com/GeoNet/weft"
	"net/http"
)

var (
	muxV1GeoJSON *http.ServeMux
	muxV1JSON    *http.ServeMux
	muxV2GeoJSON *http.ServeMux
	muxV2JSON    *http.ServeMux
	muxDefault   *http.ServeMux
	muxProto     *http.ServeMux
)

func init() {
	muxV1GeoJSON = http.NewServeMux()
	muxV1GeoJSON.HandleFunc("/quake", weft.MakeHandlerAPI(quakesRegionV1))
	muxV1GeoJSON.HandleFunc("/quake/", weft.MakeHandlerAPI(quakeV1))
	muxV1GeoJSON.HandleFunc("/intensity", weft.MakeHandlerAPI(intensityMeasuredLatestV1))
	muxV1GeoJSON.HandleFunc("/felt/report", weft.MakeHandlerAPI(feltV1))

	muxV1JSON = http.NewServeMux()
	muxV1JSON.HandleFunc("/news/geonet", weft.MakeHandlerAPI(newsV1))

	muxV2GeoJSON = http.NewServeMux()
	muxV2GeoJSON.HandleFunc("/intensity", weft.MakeHandlerAPI(intensityV2))
	muxV2GeoJSON.HandleFunc("/quake", weft.MakeHandlerAPI(quakesV2))
	muxV2GeoJSON.HandleFunc("/quake/", weft.MakeHandlerAPI(quakeV2))
	muxV2GeoJSON.HandleFunc("/quake/history/", weft.MakeHandlerAPI(quakeHistoryV2))
	muxV2GeoJSON.HandleFunc("/volcano/val", weft.MakeHandlerAPI(valV2))
	muxV2GeoJSON.HandleFunc("/volcano/quake/", weft.MakeHandlerAPI(quakesVolcanoRegionV2))
	muxV2GeoJSON.HandleFunc("/volcano/region/", weft.MakeHandlerAPI(volcanoRegionV2))
	muxV2GeoJSON.HandleFunc("/delta/mark", weft.MakeHandlerAPI(marksV2))

	muxV2JSON = http.NewServeMux()
	muxV2JSON.HandleFunc("/news/geonet", weft.MakeHandlerAPI(newsV2))
	muxV2JSON.HandleFunc("/quake/stats", weft.MakeHandlerAPI(quakeStatsV2))
	muxV2JSON.HandleFunc("/delta/mark", weft.MakeHandlerAPI(marksV2JSON))

	// protobufs
	muxProto = http.NewServeMux()
	muxProto.HandleFunc("/quake", weft.MakeHandlerAPI(quakesProto))
	muxProto.HandleFunc("/quake/", weft.MakeHandlerAPI(quakeProto))
	muxProto.HandleFunc("/quake/technical/", weft.MakeHandlerAPI(quakeTechnicalProto))
	muxProto.HandleFunc("/quake/history/", weft.MakeHandlerAPI(quakeHistoryProto))
	muxProto.HandleFunc("/intensity", weft.MakeHandlerAPI(intensityProto))
	muxProto.HandleFunc("/volcano/val", weft.MakeHandlerAPI(valProto))
	muxProto.HandleFunc("/news/geonet", weft.MakeHandlerAPI(newsProto))
	muxProto.HandleFunc("/quake/stats", weft.MakeHandlerAPI(quakeStatsProto))
	muxProto.HandleFunc("/volcano/region/stats/", weft.MakeHandlerAPI(quakeInVolcanoRegionStatsProto))
	muxProto.HandleFunc("/volcano/region/history/", weft.MakeHandlerAPI(volcanoRegionHistoryProto))
	muxProto.HandleFunc("/delta/mark", weft.MakeHandlerAPI(marksProto))

	// muxDefault handles routes with no Accept version.
	// soh routes
	muxDefault = http.NewServeMux()
	muxDefault.HandleFunc("/soh", http.HandlerFunc(soh))
	muxDefault.HandleFunc("/soh/esb", http.HandlerFunc(sohEsb))
	muxDefault.HandleFunc("/soh/up", http.HandlerFunc(up))
	muxDefault.HandleFunc("/soh/impact", http.HandlerFunc(impactSOH))

	muxDefault.HandleFunc("/cap/1.2/GPA1.0/quake/", weft.MakeHandlerAPI(capQuake))
	muxDefault.HandleFunc("/cap/1.2/GPA1.0/feed/atom1.0/quake", weft.MakeHandlerAPI(capQuakeFeed))
	// The 'latest' version of the API for unversioned requests.
	muxDefault.HandleFunc("/quake/", weft.MakeHandlerAPI(quakeV2))
	muxDefault.HandleFunc("/quake", weft.MakeHandlerAPI(quakesV2))
	muxDefault.HandleFunc("/quake/history/", weft.MakeHandlerAPI(quakeHistoryV2))
	muxDefault.HandleFunc("/quake/stats", weft.MakeHandlerAPI(quakeStatsV2))
	muxDefault.HandleFunc("/intensity", weft.MakeHandlerAPI(intensityV2))
	muxDefault.HandleFunc("/news/geonet", weft.MakeHandlerAPI(newsV2))
	muxDefault.HandleFunc("/volcano/val", weft.MakeHandlerAPI(valV2))
	muxDefault.HandleFunc("/volcano/quake/", weft.MakeHandlerAPI(quakesVolcanoRegionV2))
	muxDefault.HandleFunc("/volcano/region/", weft.MakeHandlerAPI(volcanoRegionV2))
	muxDefault.HandleFunc("/quakes/services/all.json", weft.MakeHandlerAPI(quakesWWWall))
	muxDefault.HandleFunc("/quakes/services/felt.json", weft.MakeHandlerAPI(quakesWWWfelt))
	muxDefault.HandleFunc("/quakes/services/quakes/newzealand/", weft.MakeHandlerAPI(quakesWWWnz))
	muxDefault.HandleFunc("/quake/services/quake/", weft.MakeHandlerAPI(quakeWWW))
	muxDefault.HandleFunc("/delta/mark", weft.MakeHandlerAPI(marksV2))

	for _, v := range []*http.ServeMux{muxV1JSON, muxV2JSON, muxV1GeoJSON, muxV2GeoJSON, muxDefault, muxProto} {
		v.HandleFunc("/", weft.MakeHandlerPage(docs))
	}

}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Accept") {
	case protobuf:
		muxProto.ServeHTTP(w, r)
	case V2GeoJSON:
		muxV2GeoJSON.ServeHTTP(w, r)
	case V1GeoJSON:
		muxV1GeoJSON.ServeHTTP(w, r)
	case V1JSON:
		muxV1JSON.ServeHTTP(w, r)
	case V2JSON:
		muxV2JSON.ServeHTTP(w, r)
	default:
		muxDefault.ServeHTTP(w, r)
	}
}
