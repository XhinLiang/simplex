# Simplex

Simplex simplifies JSON objects according to a set of provided rules.

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
  // $ = root
  "remove_properties": [ "test", "debug" ],
  "property_rules": {
    "data": { // rule for $.data(if object) or $.data[0], $.data[1] ...(if array)
      // $ = $.data or $.data[0] or $.data[1] ...
      "remove_properties": [ "data_test", "data_debug" ],
      "property_rules": {
        // nested
      }
    },
    "entity_list": { // rule for $.entity_list(if object) or $.entity_list[0], $.entity_list[1] ....(if array)
      // $ = $.entity_list or $.entity_list[0] or $.entity_list[1] ...
      "remove_properties": [ "entity_test" ],
      "property_rules": {
        "sub_properties": { // rule for $.entity_list[0].sub_properties(if object) or $.entity_list[0].sub_properties[0], $.entity_list[0].sub_properties[1] ...(if array)
          // $ = $.entity_list.sub_properties or $.entity_list[0].sub_properties[0] or $.entity_list[0].sub_properties[1] ...
          "remove_properties": [ "abc", "def" ] // remove $.abc and $.def
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
