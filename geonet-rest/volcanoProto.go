package main

import (
	"bytes"
	"database/sql"
	"github.com/GeoNet/haz"
	"github.com/GeoNet/weft"
	"github.com/golang/protobuf/proto"
	"net/http"
	"time"
	"fmt"
)

func valProto(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	if res := weft.CheckQuery(r, []string{}, []string{}); !res.Ok {
		return res
	}

	var err error
	var rows *sql.Rows

	if rows, err = db.Query(`SELECT id, title, alert_level, activity, hazards,
				ST_X(location::geometry), ST_Y(location::geometry)
				FROM haz.volcano JOIN haz.volcanic_alert_level using (alert_level)
				ORDER BY alert_level DESC, title ASC`); err != nil {
		return weft.ServiceUnavailableError(err)
	}

	var vol haz.Volcanoes

	for rows.Next() {
		v := haz.Volcano{Val: &haz.VAL{}}
		if err = rows.Scan(&v.VolcanoID, &v.Title, &v.Val.Level, &v.Val.Activity, &v.Val.Hazards,
			&v.Longitude, &v.Latitude); err != nil {
			return weft.ServiceUnavailableError(err)
		}

		vol.Volcanoes = append(vol.Volcanoes, &v)
	}

	var by []byte

	if by, err = proto.Marshal(&vol); err != nil {
		return weft.ServiceUnavailableError(err)
	}

	b.Write(by)

	h.Set("Content-Type", protobuf)
	return &weft.StatusOK
}

func quakeInVolcanoRegionStatsProto(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {

	if res := weft.CheckQuery(r, []string{}, []string{}); !res.Ok {
		return res
	}
	var q haz.QuakeStats

	var rows *sql.Rows
	var err error
	var volcanoId string

	if volcanoId, err = getVolcanoRegionStats(r); err != nil {
		return weft.BadRequest(err.Error())
	}

	if rows, err = db.Query(quakesPerDayInVolcanoRegionSQL, volcanoId); err != nil {
		return weft.ServiceUnavailableError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t time.Time
		var r haz.Rate

		if err = rows.Scan(&t, &r.Count); err != nil {
			return weft.ServiceUnavailableError(err)
		}

		r.Time = &haz.Timestamp{Sec: t.Unix(), Nsec: int64(t.Nanosecond())}

		q.PerDay = append(q.PerDay, &r)
	}
	rows.Close()

	q.Year = make(map[int32]int32)

	if rows, err = db.Query(fmt.Sprintf(sumMagsInVolcanoRegionSQL, 365), volcanoId); err != nil {
		return weft.ServiceUnavailableError(err)
	}

	defer rows.Close()

	for rows.Next() {
		var k, v int32
		if err = rows.Scan(&k, &v); err != nil {
			return weft.ServiceUnavailableError(err)
		}
		q.Year[k] = v
	}
	rows.Close()

	q.Month = make(map[int32]int32)

	if rows, err = db.Query(fmt.Sprintf(sumMagsInVolcanoRegionSQL, 28), volcanoId); err != nil {
		return weft.ServiceUnavailableError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var k, v int32
		if err = rows.Scan(&k, &v); err != nil {
			return weft.ServiceUnavailableError(err)
		}
		q.Month[k] = v
	}
	rows.Close()

	q.Week = make(map[int32]int32)

	if rows, err = db.Query(fmt.Sprintf(sumMagsInVolcanoRegionSQL, 7), volcanoId); err != nil {
		return weft.ServiceUnavailableError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var k, v int32
		if err = rows.Scan(&k, &v); err != nil {
			return weft.ServiceUnavailableError(err)
		}
		q.Week[k] = v
	}
	rows.Close()

	var by []byte

	if by, err = proto.Marshal(&q); err != nil {
		return weft.ServiceUnavailableError(err)
	}

	b.Write(by)

	h.Set("Content-Type", protobuf)
	h.Set("Surrogate-Control", maxAge300)

	return &weft.StatusOK
}
