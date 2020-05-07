# whoisyou

Take a list of domains and output the hostname and ip.

## Installation

For now the only way to install is to use go get.

```bash
go get github.com/QSoloX/whoisyou
```

## Usage

The basic usuage to save it the output to a file.
```bash
▶ cat domains.txt|whoisyou > output.txt
```
Output without directing the output
```bash
▶ cat domains.txt|whoisyou
https://site.com 127.0.0.1
https://example.com 127.1.1.1
```

## Additional Usage
The output can be used with a python script like below
```python
with open(f"output.txt") as file:
    for line in file:
        # Strip newline chars and split at space, then assign split to vars
        hostname,ip = line.strip()line.split(" ")
        # Do anything with the output
        print(f"{hostname} has the ip of {ip}")
```
```bash
▶ python main.py
https://site.com has the ip of 127.0.0.1
https://example.com has the ip of 127.1.1.1
```


## Known Issues
* There seems to be a output issue when used on windows that makes the file become a utf-16 file.
## License
[MIT](https://choosealicense.com/licenses/mit/)