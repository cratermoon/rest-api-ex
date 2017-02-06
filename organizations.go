package main

import (
  "log"
  "sort"
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
  Organizations map[int][]string `json:"organizations"`
}

func (ol *Organizations) add(id int, org []string) {
  log.Printf("Adding %d %v", id, org)
  ol.Organizations[id] = org
}

func (ol *Organizations) count() int {
  return len(ol.Organizations)
}
func (ol Organizations) Get(id int) Organization {
  if values,ok := ol.Organizations[id]; ok {
    return Organization{id, values[0], values[1], values[2], values[3], values[4]}
  }
  return Organization{0, "NOT FOUND", "","","", ""}
}

func populate(records [][]string) *Organizations {
  organizations := new(Organizations)
  organizations.Organizations = make(map[int][]string)
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


func OrderBy(key string, orgs []Organization) []Organization {
  switch key {
  case "id":
    sort.Sort(ById(orgs))
  case "name":
    sort.Sort(ByName(orgs))
  case "city":
    sort.Sort(ByCity(orgs))
  case "state":
    sort.Sort(ByState(orgs))
  case "postal":
    sort.Sort(ByPostal(orgs))
  case "category":
    sort.Sort(ByCategory(orgs))
  }
  return orgs
}

func ReverseOrderBy(key string, orgs []Organization) []Organization {
  switch key {
  case "id":
    sort.Sort(sort.Reverse(ById(orgs)))
  case "name":
    sort.Sort(sort.Reverse(ByName(orgs)))
  case "city":
    sort.Sort(sort.Reverse(ByCity(orgs)))
  case "state":
    sort.Sort(sort.Reverse(ByState(orgs)))
  case "postal":
    sort.Sort(sort.Reverse(ByPostal(orgs)))
  case "category":
    sort.Sort(sort.Reverse(ByCategory(orgs)))
  }
  return orgs
}

func (ol Organizations) Search(name string, city string, state string, postal string, category string) []Organization {
  matches := []Organization{}
  for id, org := range ol.Organizations {
    if match(name, org[0]) && match(city, org[1]) && match(state, org[2])&& match(postal, org[3])&& match(category, org[4]){
      matches = append(matches, Organization{id, org[0], org[1], org[2], org[3], org[4]})
    }
  }
  return matches
}
