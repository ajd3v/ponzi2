// Package iex provides a client to get stock data using the IEX API.
package iex

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// loc is the timezone to use when parsing dates.
var loc = mustLoadLocation("America/New_York")

// Range is the range to specify in the request.
type Range string

// ChartRange values.
const (
	RangeOneDay   Range = "1d"
	RangeTwoYears       = "2y"
)

// GetStocksRequest is the request for GetStocks.
type GetStocksRequest struct {
	Symbols   []string
	Range     Range
	ChartLast int
}

// Stock is the response from calling GetStocks.
type Stock struct {
	Symbol string
	Quote  *Quote
	Chart  []*ChartPoint
}

// Quote is a stock quote.
type Quote struct {
	CompanyName  string
	LatestPrice  float32
	LatestSource string
	LatestTime   string
	LatestUpdate int64
	LatestVolume int
}

// ChartPoint is a single point on the chart.
type ChartPoint struct {
	Date          time.Time
	Open          float32
	High          float32
	Low           float32
	Close         float32
	Volume        int
	Change        float32
	ChangePercent float32
}

// Client is used to make IEX API requests.
type Client struct {
	// dumpAPIResponses dumps API responses into text files.
	dumpAPIResponses bool
}

// NewClient returns a new Client.
func NewClient(dumpAPIResponses bool) *Client {
	return &Client{dumpAPIResponses: dumpAPIResponses}
}

// GetStocks gets a series of trading sessions for a stock symbol.
func (c *Client) GetStocks(ctx context.Context, req *GetStocksRequest) ([]*Stock, error) {
	if len(req.Symbols) == 0 {
		return nil, nil
	}

	if req.Range == "" {
		return nil, errors.New("iex: missing range for chart req")
	}

	if req.ChartLast < 0 {
		return nil, errors.New("iex: last must be greater than or equal to zero")
	}

	u, err := url.Parse("https://api.iextrading.com/1.0/stock/market/batch")
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("symbols", strings.Join(req.Symbols, ","))
	v.Set("types", "quote,chart")
	v.Set("range", string(req.Range))
	v.Set("filter", strings.Join([]string{
		// Keys for quote.
		"companyName",
		"latestPrice",
		"latestSource",
		"latestTime",
		"latestUpdate",
		"latestVolume",

		// Keys for chart.
		"date",
		"minute",
		"open",
		"high",
		"low",
		"close",
		"volume",
		"change",
		"changePercent",
	}, ","))
	if req.ChartLast > 0 {
		v.Set("chartLast", strconv.Itoa(req.ChartLast))
	}
	u.RawQuery = v.Encode()

	httpReq, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(httpReq.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	r := httpResp.Body
	if c.dumpAPIResponses {
		rr, err := dumpResponse(fmt.Sprintf("iex-%s.txt", strings.Join(req.Symbols, "-")), r)
		if err != nil {
			return nil, fmt.Errorf("iex: failed to dump resp: %v", err)
		}
		r = rr
	}

	resp, err := decodeStocks(r)
	if err != nil {
		return nil, fmt.Errorf("iex: failed to decode resp: %v", err)
	}

	return resp, nil
}

func decodeStocks(r io.Reader) ([]*Stock, error) {
	type quote struct {
		CompanyName  string  `json:"companyName"`
		LatestPrice  float64 `json:"latestPrice"`
		LatestSource string  `json:"latestSource"`
		LatestTime   string  `json:"latestTime"`
		LatestUpdate int64   `json:"latestUpdate"`
		LatestVolume int64   `json:"latestVolume"`
	}

	type chartPoint struct {
		Date          string  `json:"date"`
		Minute        string  `json:"minute"`
		Open          float64 `json:"open"`
		High          float64 `json:"high"`
		Low           float64 `json:"low"`
		Close         float64 `json:"close"`
		Volume        float64 `json:"volume"`
		Change        float64 `json:"change"`
		ChangePercent float64 `json:"changePercent"`
	}

	type stock struct {
		Quote *quote        `json:"quote"`
		Chart []*chartPoint `json:"chart"`
	}

	var m map[string]stock
	dec := json.NewDecoder(r)
	if err := dec.Decode(&m); err != nil {
		return nil, fmt.Errorf("json decode failed: %v", err)
	}

	var chs []*Stock

	for s, d := range m {
		ch := &Stock{Symbol: s}

		if d.Quote != nil {
			ch.Quote = &Quote{
				CompanyName:  d.Quote.CompanyName,
				LatestPrice:  float32(d.Quote.LatestPrice),
				LatestSource: d.Quote.LatestSource,
				LatestTime:   d.Quote.LatestTime,
				LatestUpdate: d.Quote.LatestUpdate,
				LatestVolume: int(d.Quote.LatestVolume),
			}
		}

		for _, pt := range d.Chart {
			date, err := parseDateMinute(pt.Date, pt.Minute)
			if err != nil {
				return nil, fmt.Errorf("parsing date (%s) failed: %v", pt.Date, err)
			}

			ch.Chart = append(ch.Chart, &ChartPoint{
				Date:          date,
				Open:          float32(pt.Open),
				High:          float32(pt.High),
				Low:           float32(pt.Low),
				Close:         float32(pt.Close),
				Volume:        int(pt.Volume),
				Change:        float32(pt.Change),
				ChangePercent: float32(pt.ChangePercent),
			})
		}
		sort.Slice(ch.Chart, func(i, j int) bool {
			return ch.Chart[i].Date.Before(ch.Chart[j].Date)
		})
		chs = append(chs, ch)
	}

	return chs, nil
}

func parseDateMinute(date, minute string) (time.Time, error) {
	if minute != "" {
		return time.ParseInLocation("20060102 15:04", date+" "+minute, loc)
	}
	return time.ParseInLocation("2006-01-02", date, loc)
}

func dumpResponse(fileName string, r io.Reader) (io.ReadCloser, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(file, "%s", b)

	return ioutil.NopCloser(bytes.NewBuffer(b)), nil
}

func mustLoadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		log.Fatalf("time.LoadLocation(%s) failed: %v", name, err)
	}
	return loc
}
