package main

// id,name,city,state,postal,category
type ById []Organization
func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

type ByName []Organization
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByCity []Organization
func (a ByCity) Len() int           { return len(a) }
func (a ByCity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCity) Less(i, j int) bool { return a[i].City < a[j].City }

type ByState []Organization
func (a ByState) Len() int           { return len(a) }
func (a ByState) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByState) Less(i, j int) bool { return a[i].State < a[j].State }

type ByPostal []Organization
func (a ByPostal) Len() int           { return len(a) }
func (a ByPostal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPostal) Less(i, j int) bool { return a[i].Postal < a[j].Postal }

type ByCategory []Organization
func (a ByCategory) Len() int           { return len(a) }
func (a ByCategory) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCategory) Less(i, j int) bool { return a[i].Category < a[j].Category }
