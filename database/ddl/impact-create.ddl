CREATE SCHEMA impact;

-- impact.intensity_reported is for reported shaking intensity.
-- The numbered geohash columns are used to store the geohash of the location at various lengths.
-- These are then used with group by to aggregate reports at zoom levels represented by the
-- length (spatial extent) of the geohash.  Precomputing them is a preformance enhancement for the select.
CREATE TABLE impact.intensity_reported (
	source TEXT NOT NULL,
	time TIMESTAMP WITH TIME ZONE NOT NULL,
	mmi int NOT NULL CONSTRAINT mmi_check CHECK (mmi >= 1 AND mmi <= 12),
	comment varchar(140) NOT NULL,
	geohash5 varchar(5) NOT NULL,
	geohash6 varchar(6) NOT NULL,
	location GEOGRAPHY(POINT, 4326) NOT NULL,
	UNIQUE (source, time)
);

CREATE INDEX ON impact.intensity_reported (time);
CREATE INDEX ON impact.intensity_reported (geohash5);
CREATE INDEX ON impact.intensity_reported (geohash6);

-- impact.intensity_measured is for measured shaking intensity.
CREATE TABLE impact.intensity_measured (
	source TEXT,
	time TIMESTAMP WITH TIME ZONE NOT NULL,
	mmi int NOT NULL CONSTRAINT mmi_check CHECK (mmi >= 1 AND mmi <= 12),
	location GEOGRAPHY(POINT, 4326) NOT NULL,
	UNIQUE (source)
);

-- tables for peak ground acceleration and velocity.
-- want the maximum values in the last hour for vertical
-- and horizontal for both pga and pgv.  The max horizontal
-- might not occur at the same time as the max vertical value
-- in the hour.  There are usually two horizontal components of
-- ground motion, was just the max of either component for
-- horizontal.

CREATE TABLE impact.pga (
	source TEXT,
	time_h TIMESTAMP WITH TIME ZONE NOT NULL,
	time_v TIMESTAMP WITH TIME ZONE NOT NULL,
	horizontal NUMERIC NOT NULL,
	vertical NUMERIC NOT NULL,
	location GEOGRAPHY(POINT, 4326) NOT NULL,
	UNIQUE (source)
);

CREATE TABLE impact.pgv (
	source TEXT,
	time_h TIMESTAMP WITH TIME ZONE NOT NULL,
	time_v TIMESTAMP WITH TIME ZONE NOT NULL,
	horizontal NUMERIC NOT NULL,
	vertical NUMERIC NOT NULL,
	location GEOGRAPHY(POINT, 4326) NOT NULL,
	UNIQUE (source)
);
