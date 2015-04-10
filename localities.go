package msg

// psql -h 127.0.0.1 hazard hazard_w -c "select format('{Name:\`%s\`, Longitude:%s, Latitude:%s, size:"%s"},', name, ST_X(locality_geom),
// ST_Y(locality_geom), size) from qrt.locality where size in (0,1,2) AND ST_X(locality_geom) > 160 AND ST_Y(locality_geom) < -30 order by name;"
//
// Then take out the offshore islands

var localities []Locality

type Locality struct {
	Name                string
	Longitude, Latitude float64
	size                int
}

func init() {
	localities = []Locality{
		{Name: `Akaroa`, Longitude: 172.97, Latitude: -43.82, size: 2},
		{Name: `Alexandra`, Longitude: 169.38, Latitude: -45.25, size: 2},
		{Name: `Amberley`, Longitude: 172.73, Latitude: -43.17, size: 2},
		{Name: `Arthur's Pass`, Longitude: 171.57, Latitude: -42.95, size: 2},
		{Name: `Ashburton`, Longitude: 171.75, Latitude: -43.9, size: 1},
		{Name: `Auckland`, Longitude: 174.77, Latitude: -36.85, size: 0},
		{Name: `Balclutha`, Longitude: 169.73, Latitude: -46.23, size: 2},
		{Name: `Blenheim`, Longitude: 173.95, Latitude: -41.52, size: 1},
		{Name: `Cambridge`, Longitude: 175.47, Latitude: -37.88, size: 1},
		{Name: `Cape Reinga`, Longitude: 172.68, Latitude: -34.43, size: 2},
		{Name: `Castlepoint`, Longitude: 176.22, Latitude: -40.9, size: 2},
		{Name: `Cheviot`, Longitude: 173.27, Latitude: -42.82, size: 2},
		{Name: `Christchurch`, Longitude: 172.63, Latitude: -43.53, size: 0},
		{Name: `Collingwood`, Longitude: 172.68, Latitude: -40.68, size: 2},
		{Name: `Culverden`, Longitude: 172.85, Latitude: -42.78, size: 2},
		{Name: `Dannevirke`, Longitude: 176.1, Latitude: -40.2, size: 2},
		{Name: `Dunedin`, Longitude: 170.5, Latitude: -45.88, size: 1},
		{Name: `Eketahuna`, Longitude: 175.7, Latitude: -40.65, size: 2},
		{Name: `Fairlie`, Longitude: 170.83, Latitude: -44.1, size: 2},
		{Name: `Feilding`, Longitude: 175.57, Latitude: -40.23, size: 1},
		{Name: `French Pass`, Longitude: 173.83, Latitude: -40.93, size: 2},
		{Name: `Geraldine`, Longitude: 171.23, Latitude: -44.1, size: 2},
		{Name: `Gisborne`, Longitude: 178.02, Latitude: -38.67, size: 1},
		{Name: `Gore`, Longitude: 168.93, Latitude: -46.1, size: 1},
		{Name: `Greymouth`, Longitude: 171.2, Latitude: -42.45, size: 1},
		{Name: `Haast`, Longitude: 169.05, Latitude: -43.88, size: 2},
		{Name: `Hamilton`, Longitude: 175.28, Latitude: -37.78, size: 1},
		{Name: `Hanmer Springs`, Longitude: 172.83, Latitude: -42.52, size: 2},
		{Name: `Hastings`, Longitude: 176.85, Latitude: -39.65, size: 1},
		{Name: `Hawera`, Longitude: 174.28, Latitude: -39.58, size: 1},
		{Name: `Hokitika`, Longitude: 170.97, Latitude: -42.72, size: 2},
		{Name: `Hunterville`, Longitude: 175.57, Latitude: -39.93, size: 2},
		{Name: `Invercargill`, Longitude: 168.37, Latitude: -46.42, size: 1},
		{Name: `Kaikoura`, Longitude: 173.68, Latitude: -42.4, size: 2},
		{Name: `Kaitaia`, Longitude: 173.27, Latitude: -35.12, size: 2},
		{Name: `Karamea`, Longitude: 172.12, Latitude: -41.25, size: 2},
		{Name: `Kawhia`, Longitude: 174.82, Latitude: -38.07, size: 2},
		{Name: `Levin`, Longitude: 175.28, Latitude: -40.62, size: 1},
		{Name: `Lumsden`, Longitude: 168.45, Latitude: -45.73, size: 2},
		{Name: `Martinborough`, Longitude: 175.45, Latitude: -41.22, size: 2},
		{Name: `Masterton`, Longitude: 175.65, Latitude: -40.95, size: 1},
		{Name: `Matawai`, Longitude: 177.53, Latitude: -38.35, size: 2},
		{Name: `Methven`, Longitude: 171.65, Latitude: -43.63, size: 2},
		{Name: `Milford Sound`, Longitude: 167.93, Latitude: -44.68, size: 2},
		{Name: `Mokau`, Longitude: 174.62, Latitude: -38.7, size: 2},
		{Name: `Motueka`, Longitude: 173.02, Latitude: -41.12, size: 2},
		{Name: `Mount Cook`, Longitude: 170.1, Latitude: -43.73, size: 2},
		{Name: `Murchison`, Longitude: 172.33, Latitude: -41.8, size: 2},
		{Name: `Murupara`, Longitude: 176.7, Latitude: -38.45, size: 2},
		{Name: `Napier`, Longitude: 176.9, Latitude: -39.5, size: 1},
		{Name: `Nelson`, Longitude: 173.28, Latitude: -41.27, size: 1},
		{Name: `New Plymouth`, Longitude: 174.07, Latitude: -39.07, size: 1},
		{Name: `Oamaru`, Longitude: 170.97, Latitude: -45.1, size: 1},
		{Name: `Ohakune`, Longitude: 175.42, Latitude: -39.42, size: 2},
		{Name: `Opotiki`, Longitude: 177.28, Latitude: -38.02, size: 2},
		{Name: `Opunake`, Longitude: 173.85, Latitude: -39.45, size: 2},
		{Name: `Oxford`, Longitude: 172.2, Latitude: -43.3, size: 2},
		{Name: `Palmerston`, Longitude: 170.72, Latitude: -45.48, size: 2},
		{Name: `Palmerston North`, Longitude: 175.62, Latitude: -40.37, size: 1},
		{Name: `Paraparaumu`, Longitude: 175, Latitude: -40.92, size: 1},
		{Name: `Picton`, Longitude: 174, Latitude: -41.3, size: 2},
		{Name: `Pongaroa`, Longitude: 176.18, Latitude: -40.55, size: 2},
		{Name: `Porangahau`, Longitude: 176.62, Latitude: -40.3, size: 2},
		{Name: `Pukekohe`, Longitude: 174.9, Latitude: -37.2, size: 1},
		{Name: `Queenstown`, Longitude: 168.67, Latitude: -45.03, size: 2},
		{Name: `Ranfurly`, Longitude: 170.1, Latitude: -45.13, size: 2},
		{Name: `Reefton`, Longitude: 171.87, Latitude: -42.12, size: 2},
		{Name: `Rotorua`, Longitude: 176.23, Latitude: -38.13, size: 1},
		{Name: `Roxburgh`, Longitude: 169.32, Latitude: -45.55, size: 2},
		{Name: `Ruatoria`, Longitude: 178.32, Latitude: -37.88, size: 2},
		{Name: `Seddon`, Longitude: 174.07, Latitude: -41.67, size: 2},
		{Name: `St Arnaud`, Longitude: 172.85, Latitude: -41.8, size: 2},
		{Name: `Stratford`, Longitude: 174.28, Latitude: -39.35, size: 2},
		{Name: `Taihape`, Longitude: 175.8, Latitude: -39.68, size: 2},
		{Name: `Taumarunui`, Longitude: 175.27, Latitude: -38.88, size: 2},
		{Name: `Taupo`, Longitude: 176.08, Latitude: -38.7, size: 1},
		{Name: `Tauranga`, Longitude: 176.17, Latitude: -37.68, size: 1},
		{Name: `Te Anau`, Longitude: 167.72, Latitude: -45.42, size: 2},
		{Name: `Te Araroa`, Longitude: 178.37, Latitude: -37.63, size: 2},
		{Name: `Te Aroha`, Longitude: 175.7, Latitude: -37.53, size: 2},
		{Name: `Te Awamutu`, Longitude: 175.33, Latitude: -38.02, size: 1},
		{Name: `Te Kaha`, Longitude: 177.68, Latitude: -37.75, size: 2},
		{Name: `Te Kuiti`, Longitude: 175.17, Latitude: -38.33, size: 2},
		{Name: `Thames`, Longitude: 175.55, Latitude: -37.15, size: 2},
		{Name: `Timaru`, Longitude: 171.25, Latitude: -44.4, size: 1},
		{Name: `Tokomaru Bay`, Longitude: 178.32, Latitude: -38.13, size: 2},
		{Name: `Tokoroa`, Longitude: 175.87, Latitude: -38.23, size: 1},
		{Name: `Tolaga Bay`, Longitude: 178.3, Latitude: -38.37, size: 2},
		{Name: `Tuatapere`, Longitude: 167.68, Latitude: -46.13, size: 2},
		{Name: `Turangi`, Longitude: 175.8, Latitude: -39, size: 2},
		{Name: `Twizel`, Longitude: 170.1, Latitude: -44.27, size: 2},
		{Name: `Waimate`, Longitude: 171.05, Latitude: -44.73, size: 2},
		{Name: `Waipukurau`, Longitude: 176.55, Latitude: -40, size: 2},
		{Name: `Wairoa`, Longitude: 177.42, Latitude: -39.05, size: 2},
		{Name: `Waitangi, Chatham Is.`, Longitude: 183.45, Latitude: -43.95, size: 2},
		{Name: `Wanaka`, Longitude: 169.13, Latitude: -44.7, size: 2},
		{Name: `Waverley`, Longitude: 174.63, Latitude: -39.77, size: 2},
		{Name: `Wellington`, Longitude: 174.77, Latitude: -41.28, size: 0},
		{Name: `Westport`, Longitude: 171.6, Latitude: -41.75, size: 2},
		{Name: `Whakatane`, Longitude: 176.98, Latitude: -37.97, size: 1},
		{Name: `Whangamata`, Longitude: 175.87, Latitude: -37.22, size: 2},
		{Name: `Whanganui`, Longitude: 175.05, Latitude: -39.93, size: 1},
		{Name: `Whangarei`, Longitude: 174.32, Latitude: -35.72, size: 1},
		{Name: `White Island`, Longitude: 177.18, Latitude: -37.52, size: 2},
		{Name: `Whitianga`, Longitude: 175.7, Latitude: -36.82, size: 2},
	}
}

// Compass converts bearing (0-360) to a compass bearing name e.g., south-east.
func Compass(bearing float64) string {
	switch {
	case bearing >= 337.5 && bearing <= 360:
		return "north"
	case bearing >= 0 && bearing <= 22.5:
		return "north"
	case bearing > 22.5 && bearing < 67.5:
		return "north-east"
	case bearing >= 67.5 && bearing <= 112.5:
		return "east"
	case bearing > 112.5 && bearing < 157.5:
		return "south-east"
	case bearing >= 157.5 && bearing <= 202.5:
		return "south"
	case bearing > 202.5 && bearing < 247.5:
		return "south-west"
	case bearing >= 247.5 && bearing <= 292.5:
		return "west"
	case bearing > 292.5 && bearing < 337.5:
		return "north-west"
	default:
		return "north"
	}
}