# microsat_errcalc

## Summary

Microsatellite genotyping error rate calculator program which reads samples from CSV file and outputs analyzed results.

### Usage

Example when ran from command line:
microsat_errcalc.exe <source file path> <output file name (optional)>

Also see Manual.pdf.

### Notes
Source file should be in csv format (';' as a separator). ',' is a forbidden character.
Results are printed to the console by default. Results can be saved as a csv file if the output file is defined.
There is no limitations for columns and rows. You can have your custom text at the beginning of the file and empty lines between the headers and the loci. An empty line between the samples will separate them and create new replicas.

### Example
My custom text 1 (optional);
My custom text 2 (optional);

;Locus1;;Locus2;;Locus3;
Sample1;100;200;100;200;100;200;
Sample2;100;200;100;200;100;200;

Sample1;100;200;100;200;100;200;
Sample1;100;200;100;200;100;200;
Sample1;100;200;100;200;100;200;

Sample2;100;200;100;200;100;200;

## Installation

You need to have Go (https://golang.org/doc/install) installed in your computer.

After installing Go, clone this repository inside your Go workspace. Then run `go build` or `go install` to get executable binary file.

## Running tests

```go test ./...```
