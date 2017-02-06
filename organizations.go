package main

import (
  "log"
  "strconv"
)
// id,name,city,state,postal,category
type Organization struct {
  Id int `json:"id"`
  Name string `json:"name"`
  City string `json:"city"`
  State string `json:"state"`
  Postal string `json:"postal"`
  Category string `json:"category"`
}

type Organizations struct {
  organizations map[int][]string
}

func (ol *Organizations) add(id int, org []string) {
  log.Printf("Adding %d %v", id, org)
  ol.organizations[id] = org
}

func (ol *Organizations) count() int {
  return len(ol.organizations)
}
func (ol Organizations) Get(id int) Organization {
  if values,ok := ol.organizations[id]; ok {
    return Organization{id, values[0], values[1], values[2], values[3], values[4]}
  }
  return Organization{0, "NOT FOUND", "","","", ""}
}

func populate(records [][]string) *Organizations {
  organizations := new(Organizations)
  organizations.organizations = make(map[int][]string)
  for idx := range records {
		r := records[idx]
		id,err := strconv.Atoi(r[0])
    if err != nil {
			log.Fatal(err)
		}
    org := []string{r[1], r[2], r[3], r[4], r[5]}

		organizations.add(id, org)
	}
  log.Printf("Population complete, %d organizations found\n", organizations.count())
	return organizations
}

func match(criteria string, value string) bool {
  return criteria == "" || criteria == value
}

func (ol Organizations) Search(name string, city string, state string, postal string, category string) []Organization {
  matches := []Organization{}
  for id, org := range ol.organizations {
    if match(name, org[0]) && match(city, org[1]) && match(state, org[2])&& match(postal, org[3])&& match(category, org[4]){
      matches = append(matches, Organization{id, org[0], org[1], org[2], org[3], org[4]})
    }
  }
  return matches
}
