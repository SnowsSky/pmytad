# Please Make Your Tar A Directory
PMYTAD is a utility to extract tar archives without it spilling everywhere. Automatically creates a folder when none is detected

## Usage
```sh
./pmytad [...tar archives]
```

## Building

### Dependencies
[xz](https://github.com/ulikunitz/xz) - To decompress .xz archives.

```sh
go mod tidy
go build .
```
