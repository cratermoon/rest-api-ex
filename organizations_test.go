package main

import "testing"

func myAssertTrue(expected string, actual string) bool {
  return expected == actual
}

func TestOrganization(t *testing.T) {
  o := Organization{1,"Happy Place","Portland","OR","97201","Other"}
  if o.id != 1 {
    t.Error("Expected 1, got ", o.id)
  }
}

func TestPopulate(t *testing.T) {
  org1 := []string{"1","Home Place","Portland","OR","97201","Other"}
  org2 := []string{"2","Happy Place","Orlando","FL","32830","Park"}
  orgs := [][]string{org1,org2}
  orgCount := len(orgs)
  if orgCount != 2 {
    t.Error("Expected 2 organizations, got ", orgCount)
  }
  organizations := populate(orgs)
  if organizations == nil {
    t.Error("wut?")
  }
}

func TestGet(t *testing.T) {
  org1 := []string{"1","Home Place","Portland","OR","97201","Other"}
  org2 := []string{"2","Happy Place","Orlando","FL","32830","Park"}
  orgs := [][]string{org1,org2}
  organizations := populate(orgs)
  org,err := organizations.Get(1)
  if err != nil {
    t.Error(err)
  }
  if org.id != 1 || org.name != "Home Place" || org.city != "Portland" ||
    org.state != "OR" || org.postal != "97201" || org.category != "Other" {
      t.Error("Get organization by id 1 failed")
  }

  org,err = organizations.Get(3)
  if err == nil {
    t.Error("Organization with id 3 should not exist")
  }
}

func TestSearch(t *testing.T) {
  org1 := []string{"1","Home Place","Portland","OR","97201","Other"}
  org2 := []string{"2","Happy Place","Orlando","FL","32830","Park"}
  org3 := []string{"3","Work Place","Portland","OR","97202","Non-profit"}
  orgs := [][]string{org1,org2,org3}
  organizations := populate(orgs)
  results := organizations.Search("OR", matchState)
  count := len(results)
  if count != 2 {
    t.Error("There should be 2 organizations matching state OR, got ", count)
  }
  if !myAssertTrue("Home Place", results[0].name) {
    t.Error("The first Organization matched should be Home Place", count)
  }
}

func TestSearchByField(t *testing.T) {
  org1 := []string{"1","Home Place","Portland","OR","97201","Other"}
  org2 := []string{"2","Happy Place","Orlando","FL","32830","Park"}
  org3 := []string{"3","Work Place","Portland","OR","97202","Non-profit"}
  orgs := [][]string{org1,org2,org3}
  organizations := populate(orgs)

  results := organizations.Search("Happy Place", matchName)
  if len(results) != 1 {
    t.Error("Should be only one match for name")
  }

  results = organizations.Search("Protland", matchCity)
  if len(results) != 0 {
    t.Error("There should be no organization matching city name Protland")
  }

  results = organizations.Search("Portland", matchCity)
  count := len(results)
  if count != 2 {
    t.Error("There should be 2 organizations matching city name Portland, got ", count)
  }
}
