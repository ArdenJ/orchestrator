package data

import (
	"encoding/json"
	"io"
	"time"
)

// Work defines the data of a catalogue work
type Work struct {
	ID              string        `json:"id"`
	Title           string        `json:"title"`
	Subtitle        string        `json:"subtitle,omitempty"`
	DateCompleted   string        `json:"dateCompleted"`
	Duration        int           `json:"duration"`
	Instruments     []*Instrument `json:"instrumentation"`
	Soloists        []*Instrument `json:"soloists"`
	TotalPerformers int           `json:"totalPerformers"`
	Composer        *Name         `json:"composer,omitempty"`
	Writers         []*Name       `json:"writers,omitempty"`
	Category        string        `json:"category"`
	Description     string        `json:"description"`
	SKU             string        `json:"sku"`
	CreatedAt       string        `json:"-"`
	LastUpdatedAt   string        `json:"-"`
}

// Instrument defines an instrument within a work
type Instrument struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Aux      string `json:"string,omitempty"`
}

// Name defines the first and last names of writers and composers
type Name struct {
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}

// Catalogue is a slice containing instances of Work(s)
type Catalogue []*Work

// ToJSON encodes a []byte from the writer as json
func (c *Catalogue) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}

// FromJSON decodes json from the request body as []byte
func (c *Catalogue) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(c)
}

// GetCatalogue returns all items in the workList as a catalogue
func GetCatalogue() Catalogue {
	return workList
}

// AddWork adds a new work to the workList
func AddWork(w *Work) {
	w.TotalPerformers = len(w.Instruments)
	workList = append(workList, w)
}

var workList = []*Work{
	{
		ID:            "12abc",
		Title:         "Symphony",
		Subtitle:      "for orchestra",
		DateCompleted: time.Now().UTC().String(),
		Duration:      98765432,
		Instruments: []*Instrument{
			{
				Name:     "flute",
				Quantity: 3,
			},
		},
		TotalPerformers: 50,
		Composer:        &Name{Lastname: "Watkins", Firstname: "Huw"},
		Category:        "orchestral",
		Description:     "Symphony for orchestra by Huw Watkins",
		SKU:             "abc123",
		CreatedAt:       time.Now().UTC().String(),
		LastUpdatedAt:   time.Now().UTC().String(),
	},
}
