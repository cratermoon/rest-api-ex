package main

import (
  "errors"
  "log"
  "strconv"
)
// id,name,city,state,postal,category
type Organization struct {
  id int
  name string
  city string
  state string
  postal string
  category string
}

type OrganizationList struct {
  organizations []Organization
}

func (ol *OrganizationList) add(o Organization) {
  ol.organizations = append(ol.organizations, o)
}

func (ol *OrganizationList) count() int {
  return len(ol.organizations)
}
func (ol OrganizationList) Get(id int) (Organization, error) {
  for idx := range ol.organizations {
      if ol.organizations[idx].id == id {
        return ol.organizations[idx], nil
      }
  }
  return Organization{}, errors.New("No such organization")
}

func populate(records [][]string) *OrganizationList {
  organizations := new(OrganizationList)
  for idx := range records {
		r := records[idx]
		id,err := strconv.Atoi(r[0])
    if err != nil {
			log.Fatal(err)
		}
    name := r[1]
    city := r[2]
    state := r[3]
    postal := r[4]
    category := r[5]

		newOrganization := Organization{id,name,city,state,postal,category}
		organizations.add(newOrganization)
	}
  log.Printf("Population complete, %d organizations found\n", organizations.count())
	return organizations
}

func matchName(val string, org Organization) bool {
  return val == org.name
}

func matchCity(val string, org Organization) bool {
  return val == org.city
}

func matchState(val string, org Organization) bool {
  return val == org.state
}

func matchPostal(val string, org Organization) bool {
  return val == org.postal
}

func matchCategory(val string, org Organization) bool {
  return val == org.category
}

func (ol OrganizationList) Search(val string, fn func(string,Organization)(bool)) []Organization {
  matches := []Organization{}
  for idx := range ol.organizations {
    org := ol.organizations[idx]
      if fn(val,org) {
        matches = append(matches, org)
      }
  }
  return matches
}
