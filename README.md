# bib-cli
This is currently just for me but I will add/improve to this as I add to the repository. 

Link to [documentation](https://github.com/harrysharma1/bib-url/blob/main/docs/bibcli.md).

## How to install
Currently you can only build and compile it locally through the git repo. New distributions may come later as the other functionalities are fleshed out.

### Through Git repo
1. Clone repo: 
```bash 
git clone git@github.com:harrysharma1/bib-url.git
```
2. `cd` into `bib-url/` directory
3. Run `go build -o .` to build executable
4. You can then add to `$PATH` 

(This will be cleaned up and user friendly in the future).

## Testing/Coverage
1. Create `coverage.out` file in `bib-cli/` root.
2. Write to `coverage.out` using:
```bash
go test -coverprofile=coverage.out ./...
```

Now to view coverage through CLI/HTML you can either:

```bash
go tool cover -func=coverage.out
```
or
```bash
go tool cover -html=coverage.out
```

## Go Version

1.24.5


