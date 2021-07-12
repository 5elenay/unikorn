# Unikorn Documentation - Making Your Own Unikorn Package
So... you probably know almost anything about Unikorn now, Let's make our own package!

## Example Repo
Check [this](https://github.com/5elenay/unikorn-hello-world) repository for example.

## What We Need?
- a `src` folder for our package.
- a `metadata.json` file for give some information about our package.

## Metadata
metadata.json must contain these informations:
- `name`: Name of the package, if not it will be the repository name.
    - **Example**: `"helloworld"` 
- `description`: Description for package.
    - **Example**: `"Database client for PostgreSQL."` 
- `tags`: Tags for when we need to find a package.
    - **Example**: `["database", "postgres"]` 
- `pipreq`: If your package uses some packages from PyPi, you can add packages here.
    - **Example**: `["postgrey", "pewn"]` 

## Src Folder
We will add all of the code in `src` folder. Check example in hello-world repo.

## Deploying
Now we need to deploy our package. Its easy! Just open a new github repo, and upload these files to the repo and finished! Now everyone can use your package!