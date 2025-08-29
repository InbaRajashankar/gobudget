# gobudget
A command-line finance tracker, analyzer, and plotter. Built using Go and SQLlite

## Command Reference
`grab/g`: get line items
`-d [RANGE]`: date range
`-p [RANGE]`: price range
`-t [VALUE]`: filter by tag
`-n [VALUE]`: number of entries

`grabsum/gs [RANGE]`: get summed counts of line items - can be grouped by a factor over a specific time range.
`-i`: only include income (both by default)
`-e`: only include expense (both by default)
`-t`: group by tag
`-m`: group by month
`-y`: group by year

`grabsumgraph/graph/gsg [RANGE]`: get summed counts of line items - can be grouped by a factor over a specific time range.
`-i`: only include income (both by default)
`-e`: only include expense (both by default)
`-t`: group by tag
`-m`: group by month
`-y`: group by year

enter: enter “enter mode”
`i`: individually
`b`: bulk (csv plaintext)
`c`: bulk csv file - NO HEADER COLUMN

`help/h`: print command reference

`quit/q`: quit program

`setup/s`: creates a new db at the filepath specified in config.json

*Note that `[RANGE]` is a range in the format `start,end` either `start` or `end` can be empty for dates. All dates are M/D/Y.*

