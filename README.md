# tableprint
A program that reads a CSV input and prints it out as a formatted table

# Usage
```text
Usage: tableprint [--header] [--footer] [--style STYLE] [FILE]

Positional arguments:
  FILE

Options:
  --header, -H           Treats the first row as a header row
  --footer, -F           Treats the last row as a footer row
  --style STYLE, -S STYLE
                         Set the table style: DEFAULT/DOUBLE/LIGHT/ROUNDED/BOLD/COLORED_BRIGHT/COLORED_DARK
  --help, -h             display this help and exit
```

# Examples

```shell
$ echo "Header1,Header2\nCell1,Cell2\nCell3,Cell4\nFooter1,Footer2" | ./tableprint -H -F -S ROUNDED
╭─────────┬─────────╮
│ HEADER1 │ HEADER2 │
├─────────┼─────────┤
│ Cell1   │ Cell2   │
│ Cell3   │ Cell4   │
├─────────┼─────────┤
│ FOOTER1 │ FOOTER2 │
╰─────────┴─────────╯
```
