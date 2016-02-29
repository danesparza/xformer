# xformer [![Circle CI](https://circleci.com/gh/danesparza/xformer.svg?style=svg)](https://circleci.com/gh/danesparza/xformer)
Simple data transformer built in Go

### Quick Start
Just grab the [latest release](https://github.com/danesparza/xformer/releases) for your platform.  It's a single binary.

### Command line parameters
If you need help, just run `xformer --help`.

There are a few command line parameters available:

Parameter       | Description
----------      | -----------
inputfile       | The (text) input data file.  Each row of this file will be combined with the template file to create the output  
templatefile    | The template file.  This is a simple text file with a [Go style format placeholder](https://golang.org/pkg/fmt/) where you would like a data row to be inserted
outputfile      | The output file to be created as a result of the transformation  

### Example

**Data** file `mappings.csv`:
```
'Team A','Location 23'
'Team B','Location 4423'
'Team C','Location 823'
'Team Bogus','Location testing'
```

**Template** file `mapinsert.sql`:

```
insert into SomeTable(FieldA, SomeOtherField, TheTeam, TheLocation) values('FieldAValue', 'ThisOtherConstantValue', %v)
```

Running this **command line**:
```
xformer -inputfile="mappings.csv" -templatefile="mapinsert.sql" -outputfile="output.sql"
```

Creates the `output.sql` **result**:
```
insert into SomeTable(FieldA, SomeOtherField, TheTeam, TheLocation) values('FieldAValue', 'ThisOtherConstantValue', 'Team A','Location 23')
insert into SomeTable(FieldA, SomeOtherField, TheTeam, TheLocation) values('FieldAValue', 'ThisOtherConstantValue', 'Team B','Location 4423')
insert into SomeTable(FieldA, SomeOtherField, TheTeam, TheLocation) values('FieldAValue', 'ThisOtherConstantValue', 'Team C','Location 823')
insert into SomeTable(FieldA, SomeOtherField, TheTeam, TheLocation) values('FieldAValue', 'ThisOtherConstantValue', 'Team Bogus','Location testing')
```
### Building 

*To build, make sure you have the latest version of [Go](http://golang.org/) installed.  If you've never used Go before, it's a quick install and [there are installers for multiple platforms](http://golang.org/doc/install), including Windows, Linux and OSX.*

Run the following commands get latest and build.

```bash
go get github.com/danesparza/xformer
go build
```
