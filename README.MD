# Go-Concat

- Created with go version 1.14.2

## Why I created this app 🤔

In my work I needed to concatenate almost 1 million of documents, one of my co-workers made a java application that processes 4,000 documents (files that average 12kb-45kb) in 8 minutes, this application in Go takes only 4 seconds 😯.

## Environment specs

- Windows 10
- 8 GB RAM
- Intel Core i5 9300H
- M.2 SSD disk


## CLI arguments

- **inputFolder** : Absolute path where the files to be processed are located (documentos unitarios) **Example: C:/go_files/in**
- **outputPath** : Absolute path where the generated files are saved **Example: C:/go_files/out**
- **docsCount** : Number of documents per output file **Example: 1000**
- **outputPrefix** : Output file prefix **Example: PREFIX_**


## Run in dev

Compile and run

```bash
go build cmd/go-concat/main.go && ./main.exe C:/go_files/input_path C:/go_files/output_path 1000 OUTPUT_FILE_PREFIX_ txt
```

## Run compiled app

I rename the compiled file to **go-concat.exe**

```bash
# EXAMPLE executable_name.exe [INPUT_PATH] [OUTPUT_PATH] [DOCS_QUANTITY] [OUTPUT_FILE_PREFIX] [OUTPUT_FILE_EXTENSION]
./go-concat.exe C:/go_files/in C:/go_files/out 10000 OUTPUT_FILE_PREFIX txt
```