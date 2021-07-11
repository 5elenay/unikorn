# Unikorn Documentation - Folder Structure
If you will make a package for unikorn or you will use unikorn, you need these folder structures.

## .unik
Temporary folder that will generator when download a package. This folder will be removed after installation finished.

## unikorn
All of the packages will be in this folder. **DO NOT** touch to this folder.

## unipkg
A file for add all packages and download with one command. Will be usefull when you need to download multiple package.

### Example
`unipkg` file:
```
5elenay unikorn-hello-world
auser arepo abranch
anotheruser anotherrepo
lastuser lastrepo abranch
```
when you run `unikorn add` (don't try this with example unipkg file) it will download all of the packages in this file.

## unikorn.json
Metadata information for package will be in this file.