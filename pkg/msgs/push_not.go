package msgs

import "strings"

// Tipovi push not poruka
const (
	PushNotMsgTipPrivatna  = 1
	PushNotMsgTipBroadcast = 2
	PushNotMsgTipListic    = 3
)

// Poruka za slanje na push notifikacije
type PushNot struct {
	Id         int            `json:"push_not_id"`
	FcmId      string         //token za identifikaciju uredjaja
	FcmTopic   string         //obsolete
	DeviceType int            //0 - android, 1 - ios, 2 - web
	Tip        int            //1 - privatna, 2 - broadcast, 3 - listic
	Tekst      string         //text poruke (za tip=1 ili tip=2)
	Lang       string         //jezik igraca: en ili hr
	Listic     *PushNotListic //podaci o listicu ako je tip=3
}

// Podaci listica za slanje na push notifikacije
type PushNotListic struct {
	Id      string
	Tip     int
	Status  int
	Dobitak float64
	Broj    string
}

// Serializira poruku i pretvara ju u poruku koja se salje na push notifikacije
func (m *PushNot) Serialize() map[string]interface{} {
	d := make(map[string]interface{})
	d["tip"] = m.Tip
	if m.Listic != nil {
		l := m.Listic.Serialize()
		d["listic"] = l
	}
	if m.Tekst != "" {
		articles := strings.Split(m.Tekst, "\n")
		if len(articles) == 2 {
			d["title"] = articles[0]
			d["tekst"] = articles[1]
		} else {
			d["tekst"] = m.Tekst
		}
	}
	return d
}

// Serializira listic koji se salje kao poruka na push notifikacije
func (l *PushNotListic) Serialize() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = l.Id
	m["status"] = l.Status
	m["dobitak"] = l.Dobitak
	m["broj"] = l.Broj
	return m
}

// Kreira novu tekstualnu push notification poruku
func NewPushNotText(id int, tip int, fcmId string, deviceType int, tekst string) *PushNot {
	return &PushNot{Id: id, Tip: tip, FcmId: fcmId, DeviceType: deviceType, Tekst: tekst}
}

// Kreira novu push notification poruku za status listica
func NewPushNotListic(id int, tip int, lTip int, fcmId string, deviceType int, listicId string, status int, dobitak float64, broj string) *PushNot {
	pn := &PushNot{Id: id, Tip: tip, FcmId: fcmId, DeviceType: deviceType}
	if tip == PushNotMsgTipListic {
		pn.Listic = &PushNotListic{Id: listicId, Tip: lTip, Status: status, Dobitak: dobitak, Broj: broj}
	}
	return pn
}

// Da li je poruka za FCM klijent ili FCM topic poruka
func (m *PushNot) IsFcm() bool {
	return m.FcmId != "" || m.FcmTopic != ""
}
