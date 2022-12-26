package updater

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/AdguardTeam/golibs/log"
	"github.com/alpoprosi/curr-change/internal/models"
)

type DB interface {
	SaveBTCUSDT(bu []models.BTCUSDT) error

	SaveBTCFiat(bc []models.BTCFiat) error

	SaveFiat(c []models.Fiat) error
}

// Updater defines methods for updating database or something else.
type Updater interface {
	Run()
	Stop()
}

// CUpdater is the currencies updater. Needs to update currencies (BTC-USDT, fiat, BTC-fiat).
type CUpdater struct {
	db DB

	updTimeFiat    time.Duration
	updTimeBTCUSDT time.Duration

	apiBTC  *url.URL
	apiFiat *url.URL
	client  *http.Client

	done chan bool
}

var _ Updater = (*CUpdater)(nil)

func NewUpdater(db DB, updF, updB time.Duration, apiF, apiB string) (Updater, error) {
	af, err := url.Parse(apiF)
	if err != nil {
		return nil, fmt.Errorf("parsing fiat currencies api url: %w", err)
	}

	ab, err := url.Parse(apiB)
	if err != nil {
		return nil, fmt.Errorf("parsing BTC-USDT api url: %w", err)
	}

	return &CUpdater{
		db: db,

		updTimeFiat:    updF,
		updTimeBTCUSDT: updB,

		apiBTC:  ab,
		apiFiat: af,
		client:  http.DefaultClient,

		done: make(chan bool),
	}, nil
}

func (u *CUpdater) Run() {
	go u.runBTC()
	go u.runFiat()
}

func (u *CUpdater) Stop() {
	u.done <- true
}

func (u *CUpdater) runBTC() {
	t := time.NewTicker(u.updTimeBTCUSDT)
	defer t.Stop()

	for range t.C {
		if err := u.updBTC(); err != nil {
			log.Error("updating btc: %v", err)
		}

		select {
		case <-u.done:
			break
		default:
			continue
		}
	}
}

func (u *CUpdater) runFiat() {
	t := time.NewTicker(u.updTimeFiat)
	defer t.Stop()

	for range t.C {
		if err := u.updFiat(); err != nil {
			log.Error("updating btc: %v", err)
		}

		select {
		case <-u.done:
			break
		default:
			continue
		}
	}
}

type fiatResponse struct {
	ValCurs *valCurs `xml:"ValCurs"`
}

type valCurs struct {
	Date   time.Time `xml:"Date"`
	Valute []valute  `xml:"Valute"`
}

type valute struct {
	ID       string  `xml:"id"`
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Name     string  `xml:"Name"`
	Value    float32 `xml:"Value"`
}

func (u *CUpdater) updFiat() error {
	resp, err := u.client.Get(u.apiFiat.String())
	if err != nil {
		return fmt.Errorf("getting fiat currencies: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading fiat response body: %w", err)
	}

	fr := &fiatResponse{}
	if err = xml.Unmarshal(body, fr); err != nil {
		return fmt.Errorf("unmarshalling fiat response: %w", err)
	}

	_ = u.db.SaveFiat(nil)

	return nil
}

func (u *CUpdater) updBTC() error {
	return nil
}
