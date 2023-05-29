# Simplex

Simplex is a command line interface (CLI) program written in Go that simplifies JSON objects according to a set of provided rules.

## Installation

First, make sure that you have Go installed on your system. If you don't, follow [these instructions](https://golang.org/doc/install) to install it.

To install `simplex`, run the following command:

```bash
go install github.com/xhinliang/simplex@latest
```

## Usage

After installing `simplex`, you can run it with the following command:

```bash
echo '{"Field1":1,"Field2":"data"}' | simplex -c config.json
```

In this command, `config.json` is a JSON file containing your simplification rules and `{"Field1":1,"Field2":"data"}` is the JSON you want to simplify.

If no `-c` option is provided, `simplex` will look for a configuration file in the following locations, in order:

1. `.simplex.json` or `.simplex.jsonc` in the current directory.
2. `~/.simplex.json` or `~/.simplex.jsonc` in the user's home directory.

If no configuration file is found in any of these locations, `simplex` will print an error message and exit.

## Configuration

The configuration file is a JSON file that contains your simplification rules. It must have the following format:

```json
{
	"remove_properties": [ "Test", "Debug" ],
		"property_simplifiers": {
			"Data": {
				"remove_properties": [ "DataTest", "DataDebug" ]
			},
			"EntityList": {
				"property_simplifiers": {
					"SubProperties": {
						"remove_properties": [ "ABC", "DEF" ]
					}
				}
			}
		}
	}
```

In this example, root property named `Test` or `Debug` will be removed from the input JSON.

You can check https://github.com/XhinLiang/gosimplifier/blob/main/simplifier_test.go to explore more examples of config file.

# Contact

If you have any questions, feel free to contact me at xhinliang@gmail.com.
