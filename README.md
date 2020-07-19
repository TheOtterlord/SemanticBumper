# SemanticBumper - A version bumper

SemanticBumper is a version bumping CLI program built using [Go](https://golang.org).
The idea for this project was taken from an [OSI issue](https://github.com/open-source-ideas/open-source-ideas/issues/239).

A version bumping program built with Go.


## Install

To install SemanticBumper, you must have installed [Go](https://golang.org).
You can install SemanticBumper with the following command:

```bash
go get github.com/TheOtterlord/SemanticBumper
```


### Update

You can update SemanticBumper by calling this command.

```bash
go get -u github.com/TheOtterlord/SemanticBumper
```


## Usage

To bump files to a new version, you must create a `.bumped` file that follows the [schema](#bumped-schema) provided.
Open the command line and navigate to the directory of your project and associated `.bumped` file.
Run `SemanticBumper myfile.bumped` using your `.bumped` file in place of `myfile.bumped`. 

Example:

```bash
SemanticBumper myfile.bumped
```


### Init

Calling `SemanticBumper init` will create a `version.bumped` file in the current directory.


### Help

You can get information about the CLI by calling `SemanticBumper help`.
It will display information about the options and how to use them.


## Bumped Schema

The `.bumped` files follow a simple schema.

At the top, the `version` field specifies what the **new** version is.

```bumped
version: 0.1.1
```

After that, a `bumps` field can be used to specify all files to be bumped.

```bumped
bumps:
```

These files are written on a new line with a `-` prefix. The files must be relative to the `.bumped` files directory.

```bumped
bumps:
 - README.md
```

Example:

```bumped
version: 1.2.3
bumps:
 - README.md
 - main.example
 - lib.example
```

If you are maintaining an application and library within the smae project and they don't share the same version number, that's ok. 
SemanticBumper supports multiple version bumping groups. 

Example: 

```bumped
version: 1.1.0
bumps: 
 - README.md
 - main.example

version: 1.0.1
 - lib.example
```

TIP: You can also add comments using a `//` prefix.

```bumped
// bump file for my-project

// bump my app
version: 1.1.0
bumps: 
 - README.md
 - main.example

// bump my library
version: 1.0.1
 - lib.example
```


## License

[MIT](https://choosealicense.com/licenses/mit/)
