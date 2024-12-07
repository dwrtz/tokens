# tokens

`tokens` is a simple command-line tool written in Go that counts the number of tokens in a given text input using [tiktoken-go](https://github.com/pkoukk/tiktoken-go). It supports multiple encodings, allowing you to specify which encoding to use for tokenization.

## Features

- Count the number of tokens for any given text input.
- Accept input from:
  - A specified file: `tokens file.txt`
  - Standard input (stdin): `cat file.txt | ./tokens`
- Display usage help if no input is provided.
- Support multiple encodings, with a default of `cl100k_base`.

## Building

If you have Go installed simply run:
```bash
go build
```


## Usage

- Count tokens from a file:
  ```bash
  ./tokens file.txt
  ```
  
- Count tokens from standard input:
  ```bash
  cat file.txt | ./tokens
  ```

- Specify a different encoding (default is `cl100k_base`):
  ```bash
  cat file.txt | ./tokens -e p50k_base
  ```

Common encodings you might consider are:
- `cl100k_base`
- `p50k_base`
- `r50k_base`

Check the [tiktoken-go repository](https://github.com/pkoukk/tiktoken-go) for details and additional encodings.

If no input is provided (no file argument and no stdin), `tokens` will display a help message.

## Acknowledgments

This project uses the [tiktoken-go](https://github.com/pkoukk/tiktoken-go) library. Special thanks to the maintainers for providing a Go port of OpenAI’s `tiktoken` implementation.

## License

[MIT](LICENSE)
