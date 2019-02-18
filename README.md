# README

A simple tool to remove/add UUID dashes.

## INSTALATION
```
git clone git@github.com:alesr/dashid.git
cd dashid
go install
```

## USAGE

### Removing dashes of a valid UUID

```
ᐅ dashid b0dbb233-0d7e-4fdc-8d34-bfc32262f937
b0dbb2330d7e4fdc8d34bfc32262f937
```

### Adding back the dashes to an otherwise valid UUID

```
ᐅ dashid b0dbb2330d7e4fdc8d34bfc32262f937
b0dbb233-0d7e-4fdc-8d34-bfc32262f937
```
