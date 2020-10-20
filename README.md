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
    with_header: true
    columns:
      - name: id
        type: seq
        kwargs:
          initial_sequence: 100
      - name: last_name
        type: personal
        kwargs:
          mode: last_name
      - name: random_date
        type: date
        kwargs:
          format: 2006-01-02 03:04:05
          from: 2020
          to: 2021
      - name: random_int
        type: int
```

```bash 
out:
CSV is generated
```

|id |last_name|random_date        |random_int|
|---|---------|-------------------|----------|
|101|Allen    |2021-01-01 01:36:50|50        |
|102|Burton   |2020-05-11 09:32:31|44        |
|103|Meyer    |2020-11-23 09:10:21|56058     |
|104|Wells    |2020-12-16 10:34:11|2         |
|105|Ray      |2020-02-18 03:52:17|9397      |
|106|Jones    |2020-04-24 04:55:20|9189      |
|107|Thompson |2020-04-24 05:09:18|823       |
|108|Cox      |2020-04-01 03:49:08|1         |
|109|Duncan   |2020-12-20 04:46:56|918       |
|110|Howell   |2020-07-23 07:37:29|47685     |

