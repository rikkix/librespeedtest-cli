# LibreSpeed command line tool - modified by rikki.moe


Original README.md: [README-raw.md](README-raw.md)

## Additional features
- Support custom User-Agent

## Build

1. Install [Go](https://golang.org/doc/install/source)
2. Get the code
3. Install dependencies
   ```bash
   go mod tidy
   ```
4. Build
    ```bash
    go build -o librespeed-cli
    ```
5. Run
    ```bash
    ./librespeed-cli -h
    ```




## License

`librespeed-cli` is licensed under [GNU Lesser General Public License v3.0](LICENSE)
