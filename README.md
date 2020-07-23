# SemanticBumper - A version bumper

SemanticBumper is a CLI that makes version bumping simple, even for the largest of projects.

SemanticBumper will look through files that you specify in your [.bumped](#bumped-schema) file, replacing all text matching the semver specification with your new version.

This project originated from an [open-source-ideas issue](https://github.com/open-source-ideas/open-source-ideas/issues/239). 


## Contents

- [Install](#install)
- [Update](#update)
- [Usage](#usage)
- [Bumped Schema](#bumped-schema)
- [Contributing](#contributing)
- [License](#license)


## Install

To install SemanticBumper, you must have [Go](https://golang.org) installed on your device.
You can install SemanticBumper with the following command:

```bash
go get github.com/TheOtterlord/SemanticBumper
```

### Update

You can update SemanticBumper by calling this command.

```bash
go get -u github.com/TheOtterlord/SemanticBumper
```

Keep track of new releases by [watching](https://docs.github.com/en/github/getting-started-with-github/be-social#watching-a-repository) this repository. 
All new releases are tracked in the [changelog](https://github.com/TheOtterlord/SemanticBumper/blob/master/CHANGELOG.md).


## Usage

To bump files to a new version, you must create a `.bumped` file that follows the [schema](#bumped-schema) provided.
Open the command line and navigate to the directory of your project and associated `.bumped` file.
Run `SemanticBumper myfile.bumped` using your `.bumped` file in place of `myfile.bumped`. 

Example:

```bash
SemanticBumper version.bumped
```


### Init

Calling `SemanticBumper init` will create a `version.bumped` file in the current directory.


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
 - lib/mod.example
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
 - lib/mod.example
```


## Contributing

This project is open-source and accepts contributions from the community. 
If you wish to contribute, please read our [contributing](https://github.com/TheOtterlord/SemanticBumper/blob/master/CONTRIBUTING.md) guide to get started.


## License

This software is distributed under the [MIT](https://choosealicense.com/licenses/mit/) license. 
