## Overview
My first significant code written in the [Go Language](https://golang.org/).

Data is expected to live at https://s3-us-west-2.amazonaws.com/sample-coding-dataset/organization_sample_data.csv. This is hard-coded as a constant

The data consists of the fields id, name, city, state, postal, category with a header

All fields are strings except id, which is int.

Assumptions: The id field is unique.

## Endpoint

~~~~
GET /organizations?category=Greek&city=Washington
~~~~
### Fields
Id: numeric id
Name: string : organization name
City: string : US city name
State: string : US state name
Postal: string : US postal code
Category: string : categorization of org
### Additional query params
Orderby: string: fieldname to order the results by
Direction: string: ASC or DSC

### Dependencies
A simple http router https://github.com/julienschmidt/httprouter

## Known Flaws
* There's not much error handling
* While HTTP query parameters could allow multiple values for the same parameter, this implementation only looks at the first one.
* I'm not 100% sure I have Go's json encoding correct
* The data structures are primitive, and the mapping from the csv table to data is not at all robust
* It's not very idiomatic Go.
