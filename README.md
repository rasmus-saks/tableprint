# tableprint
A program that reads a CSV input and prints it out as a formatted table

# Usage
```text
Usage: tableprint [--header] [--footer] [--style STYLE] [FILE]

Positional arguments:
  FILE

Options:
  --header, -H           treat the first row as a header row
  --footer, -F           treat the last row as a footer row
  --style STYLE, -S STYLE
                         set the table style: DEFAULT/DOUBLE/LIGHT/ROUNDED/BOLD/COLORED_BRIGHT/COLORED_DARK [default: DEFAULT]
  --help, -h             display this help and exit
```

# Examples

Reading from a piped command
```shell
$ echo "Header1,Header2\nCell1,Cell2\nCell3,Cell4\nFooter1,Footer2" | tableprint -H -F -S ROUNDED
╭─────────┬─────────╮
│ HEADER1 │ HEADER2 │
├─────────┼─────────┤
│ Cell1   │ Cell2   │
│ Cell3   │ Cell4   │
├─────────┼─────────┤
│ FOOTER1 │ FOOTER2 │
╰─────────┴─────────╯
```

Reading from a CSV file.
```shell
$ tableprint -H -F -S DOUBLE input.csv
╔═════════╦═════════╗
║ HEADER1 ║ HEADER2 ║
╠═════════╬═════════╣
║ Cell1   ║ Cell2   ║
║ Cell3   ║ Cell4   ║
╠═════════╬═════════╣
║ FOOTER1 ║ FOOTER2 ║
╚═════════╩═════════╝

```
