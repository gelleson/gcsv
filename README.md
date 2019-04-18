# GCSV
[![Go Report Card](https://goreportcard.com/badge/github.com/gelleson/gcsv)](https://goreportcard.com/report/github.com/gelleson/gcsv)

**GCSV** is a yaml based csv file generator.


## Example 

```bash
gcsv generate example.yaml
```

```yaml
# example.yaml
documents:
    - name: example_data
      rows: 1000
      columns:
        - name: id
          type: int
          options: [seq]
        - name: name
          type: string
          options: [first_name]
        - name: last_name
          type: string
          options: [last_name]
        - name: random_date
          type: date
        - name: random_int
          type: int
```

```bash 
out:
CSV is generated
```

|    |             |           |                               |                     | 
|----|-------------|-----------|-------------------------------|---------------------| 
| 0  | similique   | Campbell  | 2009-11-29 00:00:00 +0000 UTC | 5577006791947779410 | 
| 1  | pariatur    | Robertson | 2007-08-14 00:00:00 +0000 UTC | 8674665223082153551 | 
| 2  | corrupti    | Parker    | 2005-09-12 00:00:00 +0000 UTC | 6129484611666145821 | 
| 3  | reiciendis  | Elliott   | 2012-10-16 00:00:00 +0000 UTC | 4037200794235010051 | 
| 4  | earum       | Ruiz      | 2009-02-11 00:00:00 +0000 UTC | 3916589616287113937 | 
| 5  | dolore      | Flores    | 2009-02-26 00:00:00 +0000 UTC | 6334824724549167320 | 
| 6  | et          | Andrews   | 2010-09-10 00:00:00 +0000 UTC | 605394647632969758  | 
| 7  | et          | Dunn      | 2006-07-07 00:00:00 +0000 UTC | 1443635317331776148 | 
| 8  | labore      | Gomez     | 2016-09-27 00:00:00 +0000 UTC | 894385949183117216  | 
| ...  | ... | ..   | ... | ... | 
| 1000 | placeat     | Hawkins   | 2011-04-30 00:00:00 +0000 UTC | 4751997750760398084 |  