# Unikorn Documentation - Commands

## help

Shows list of all commands.

### Example(s)

```bash
unikorn help
```

```bash
unikorn help <command>
```

## add

Download & add a package from github.

### Example(s)

```bash
unikorn add
```

```bash
unikorn add <github username> <repo name>
```

```bash
unikorn add <github username> <repo name> <branch>
```

### Options

- `no-confirmation`: Skip confirmation process for command.

## remove

Remove a package from project.

### Example(s)

```bash
unikorn remove
```

```bash
unikorn remove <package name>
```

### Options

- `no-confirmation`: Skip confirmation process for command.

## sync

Sync a package (delete & install latest version.).

### Example(s)

```bash
unikorn sync <package name>
```

### Options

- `no-confirmation`: Skip confirmation process for command.

## find

Find downloaded package(s) from name or tag.

### Example(s)

```bash
unikorn find <package name>
```

```bash
unikorn find <tag>
```

### Options

- `all`: Find all packages.

## list

List downloaded packages.

### Example(s)

```bash
unikorn list
```

## check

Check avaible updates for Unikorn.

### Example(s)

```bash
unikorn check
```

## init

Initialize basic setup for Unikorn.

### Example(s)

```bash
unikorn init
```

### Options

- `no-confirmation`: Skip confirmation process for command.

## version

Check Unikorn version.

### Example(s)

```bash
unikorn version
```
