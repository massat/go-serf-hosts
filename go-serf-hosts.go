package goSerfHosts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SerfHosts struct {
	hostsFile string
	entries   []Entry
}

func NewSerfHosts(hostsFile string) *SerfHosts {
	entries := []Entry{}
	return &SerfHosts{
		hostsFile: hostsFile,
		entries:   entries,
	}
}

func (s *SerfHosts) loadHosts() {
	fp, err := os.Open(s.hostsFile)

	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func (s *SerfHosts) addEntry(entry Entry) {
	for _, e := range s.entries {
		if entry.equals(e) {
			return
		}
	}

	s.entries = append(s.entries, entry)
}

func (s *SerfHosts) removeEntry(entry Entry) {
	newEntries := []Entry{}
	for _, e := range s.entries {
		if !entry.equals(e) {
			newEntries = append(newEntries, e)
		}
	}

	s.entries = newEntries
}

func (s *SerfHosts) parseData(data string) payload {
	datas := strings.Split(data, "\t")

	if len(datas) < 2 {
		panic("data must have at least 2 parameters.")
	}

	name := datas[0]
	address := datas[1]
	role := ""
	tags := make(map[string]string)

	// if datas[3] != "" {
	// 	pairs := strings.Split(datas[3], ",")
	// 	for _, pair := range pairs {
	// 		kv := strings.Split(pair, "=")
	// 		tags[kv[0]] = kv[1]
	// 	}
	// }

	return payload{
		name, address, role, tags,
	}
}

func (s *SerfHosts) HandleEvent(event string, data string) {
	payload := s.parseData(data)
	switch event {
	case "member-join":
		s.join(payload)
	case "member-leave":
	case "member-failed":
	case "member-update":
	case "member-reap":
	default:
	}

	defer fmt.Println(s.entries)
}

func (s *SerfHosts) join(payload payload) {
	s.addEntry(newEntry(payload.name, payload.address))
}

func (s *SerfHosts) leave(payload payload) {
	s.removeEntry(newEntry(payload.name, payload.address))
}

type Entry struct {
	Address string
	Alias   string
}

func newEntry(address, alias string) Entry {
	return Entry{address, alias}
}

func (e *Entry) equals(o Entry) bool {
	return e.Address == o.Address && e.Alias == o.Alias
}

type payload struct {
	name    string
	address string
	role    string
	tags    map[string]string
}
