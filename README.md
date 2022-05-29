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

Convert a JSON API response to CSV with [jq](https://stedolan.github.io/jq/) and print it as a table:
```shell
$ curl -s https://api.github.com/repos/github/docs/commits | jq -r '.[] | [.commit.author.name, .sha] | @csv' | tableprint
+-----------------+------------------------------------------+
| Octomerger Bot  | 044bc2bfb95a7d8f86a098da1815e82bcc35a2c6 |
| Octomerger Bot  | ad909b96502aeb691a047fb829e025c82719e707 |
| GitHub Actions  | f1d53f06637143123e420930d5744e0bfb4c9fa2 |
| Octomerger Bot  | df6d94d7ac7729a7c0c0d8e4eddd47bd090bca56 |
| Octomerger Bot  | dfe7239103c08e523d9cd5e530eb7e47542637be |
| GitHub Actions  | 8b180bb92f93b40805a573a3f6791daa2e481bda |
| Octomerger Bot  | d02a07f26c17d5a2524df1a95fc4be0666c60749 |
| Octomerger Bot  | 3ac4cedd662736d55408b65eecfada9aad91ff86 |
| Octomerger Bot  | 7c6a301103ea919e7e7befc502b8cb5728b96334 |
| Octomerger Bot  | 896d856def97a19955bdc465c2d7958eb11c8427 |
| GitHub Actions  | b1ba1acf302ad43084eb0dbd24f903041a295c3d |
| docubot         | ceaaf8e5f204e219e51422faaebad3af42b85f94 |
| docubot         | 60ba2411334cef4b5824ce251086af85ef0793db |
| Octomerger Bot  | 6fefa39fe55372d6dcddef800e367ffa24f2269f |
| Octomerger Bot  | 2d06fe39d13ebb74a9c6e80d1ca7acc119dfaa3f |
| GitHub Actions  | 28dec66086afd856bfab0ad4c5569888b536a944 |
| Octomerger Bot  | 84c0094fdcc45438a692f68120e73dc9de7ea7c6 |
| Octomerger Bot  | aa8194b3a30befd7c15bd804f5fe2cc6d0c1fe52 |
| GitHub Actions  | 1e387af0b28af42961667c990ecce1d790c0f8cf |
| Octomerger Bot  | 4c7374f02e92b7e1a8e8d7b635e5332723ae89ad |
| Octomerger Bot  | ce289ccb655b25d9dccb0acdfc358f3240747fb4 |
| Octomerger Bot  | b1f3a04609dd70941e03e152d72f183fd560feea |
| Peter Bengtsson | c35b134806bd0fbfc5efcc3c4974fcccbeec6947 |
| Octomerger Bot  | 3b1a908b1c13f35582dcad5caf23df6a5f47a1af |
| docubot         | 16ecb94a8ad2b4f4774f916b2129907fa7e8a814 |
| Octomerger Bot  | a05da560348999460c93aa4c285054d9b1c12bf5 |
| Octomerger Bot  | aca6bc83d9d34cb6a4d7534bfd5a3f27b0f52416 |
| docubot         | ed94161e09b35bc8b95664f6178667bc526e877f |
| docubot         | f7eb2d61d12cd9532cae25f1b2ba507ab1b44e9c |
| Octomerger Bot  | dc914926170fcfdf56707f555829ae9c13ed1f43 |
+-----------------+------------------------------------------+```
