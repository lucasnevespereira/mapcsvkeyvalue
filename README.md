# mapcsvkeyvalue

Script to map csv key value data.


## Usage

Replace data/example.csv with your csv file

Set your filenames
```
const (
	csvFileName = "data/example.csv"
	outFileName = "data/output.txt"
)

```

Run `make output`

Your `output.txt` is generated 👍🏻

#### Example CSV

```
123456,10%
12345736,50%
1243563425635,100%
```
#### Example Output

```
"123456": 0.1, 
"12345736": 0.5, 
"1243563425635": 1, 
```

<hr>

#### Additional info
You can edit the output of a single line as you want in `f.WriteString`.
```
	f.WriteString("\"" + data.key + "\": " + data.value + ", \n")
```

You can edit or add your string replacement rules in `rules.go`
```
	{
		character:   "Â",
		replacement: "A",
	},
```
Currently rules are only applied to the `value` but you could modify this for loop to apply them as well to `data.key`
```
    for _, r := range rules {
        data.key = strings.Replace(data.value, r.character, r.replacement, -1)
        data.value = strings.Replace(data.value, r.character, r.replacement, -1)
    }
```

You can run `make clean` to remove your `data/output.txt` file