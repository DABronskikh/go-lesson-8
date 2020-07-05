package transactions

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

type Transaction struct {
	XMLName string
	Id      string
	From    string
	To      string
	Amount  int64
	Created int64
}

type Service struct {
	mu           sync.Mutex
	Transactions []*Transaction
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(from, to string, amount int64) (string, error) {
	t := &Transaction{
		Id:      "x",
		From:    from,
		To:      to,
		Amount:  amount,
		Created: time.Now().Unix(),
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.Transactions = append(s.Transactions, t)

	return t.Id, nil
}

func (s *Service) ExportCSV(writer io.Writer) error {
	s.mu.Lock()
	if len(s.Transactions) == 0 {
		s.mu.Unlock()
		return nil
	}

	records := [][]string{}
	for _, t := range s.Transactions {
		record := []string{
			t.Id,
			t.From,
			t.To,
			strconv.FormatInt(t.Amount, 10),
			strconv.FormatInt(t.Created, 10),
		}
		records = append(records, record)
	}
	s.mu.Unlock()

	w := csv.NewWriter(writer)
	return w.WriteAll(records)
}

func (s *Service) ImportCSV(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return err
	}

	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Print(err)
		return err
	}

	for _, record := range records {

		t, err := MapRowToTransaction(record)
		if err != nil {
			log.Print(err)
			return err
		}

		s.mu.Lock()
		s.Transactions = append(s.Transactions, t)
		s.mu.Unlock()
	}

	return nil
}

func MapRowToTransaction(records []string) (*Transaction, error) {
	amount, err := strconv.ParseInt(records[3], 10, 64)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	created, err := strconv.ParseInt(records[4], 10, 64)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	tr := &Transaction{
		Id:      records[0],
		From:    records[1],
		To:      records[2],
		Amount:  amount,
		Created: created,
	}

	return tr, nil
}
