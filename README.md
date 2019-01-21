Mooncake
========
**A simplistic validation language for validating complex objects.**

Built with Antlr and Golang.

Getting Started
---------------


Writing Rules
-------------

Validation normally consists of a sequence of rules or statements that verify different parts of object(s).
For now objects in json format are handled.

A simple validation statement has the following syntax:

```
item.name eq nil => [E0123, 'item is null']!!! # if item.name is empty then fatal error
```

Breaking it down we have:

* _Identifier:_ `item.name`
* _Operator:_ `eq`
* _Literal:_ `nil`
* _Implication:_ `=>`
* _Error:_ `[E0123, 'item is null']!!` containing error code, info and severity
* _Comment:_ `# if item.name is empty then severe error`


A group of statements can depend on each or could be independent:

```
~item.name   eq nil  => [E0001, 'info-1']!!!
~item.value  eq nil  => [E0002, 'info-2']!!

 item.info   eq nil  => [E0003, 'info-3']!
```

Here, `~` implies an inorder dependent relation i.e., only if the `item.name` expression doesn't hold
then the following statement `item.value` is evaluated.


`item.info` statement is independent and will be evaluated in any case.

A statement can also have a body block which is evaluated only if the expression doesn't hold:

```
item.name eq nil  => [E0001, 'info-1']!!!
{
    item.value eq nil => [E0002, 'info-2']!!
    {
        item.value.first eq nil => [E0003, 'info-3']!
        {
            item.value.second eq nil => [E0004, 'info-4']!
        }
    }
}
```
This is similar to using `~` to have a dependent relations:

```
~item.name eq nil           => [E0001, 'info-1']!!!
~item.value eq nil          => [E0002, 'info-2']!!
~item.value.first eq nil    => [E0003, 'info-3']!
~item.value.second eq nil   => [E0004, 'info-4']!
```

Blocks are useful when using functions and inline declarations:

```$xslt
 _x @len item.list eq 0  => [E0005, 'info-5']!!
 {
    ~ _x eq 10 => [E0006, 'info-6']!! # some special case
    ~ _x eq 20 => [E0007, 'info-7']!! # another special case
 }
``` 

If there is a need to reference structures then it could be done as:
```
item.name eq ${ctx.item1}  => [E0005, 'info-5']!!
```

Here, `${...}` reflectively gets the passed `struct` and check it against the expression.


Executing Rules
---------------

For evaluation the validation rules are passed to the `mooncake executor` along with an optional `context struct`:

```go
    // read rules
    content, err := ioutil.ReadFile("rules.mck")
    rules := string(content)
    
    // to validate
    json := "sample.json"
    
    // sample context struct
    ctx := struct {
        item1 string
        item2 string
        item3 int
    }{
        "comingFromContextA",
        "comingFromContextB",
        100000,
    }
    
    // pass to executor
    result := executor.Execute(rules, json, ctx)
```

Validation Results
------------------
Currently, the output of validation rules is a structure comprising of all encountered errors
with respect to severity. Example:
```json
{
   "Fatal":[
      {
         "Code":"E0001",
         "Info":"sample is empty"
      }
   ],
   "Severe":[

   ],
   "Warning":[

   ]
}

```


Operators
---------

| Operator                  | Description                                                        |
| :------------------------ | :----------------------------------------------------------------- |
| `eq` or `==`              | left is equal to right                                             |
| `ne` or `!=`              | left is not equal to right                                         |              
| `lt` or `<`               | left is less than right                                            |
| `lte` or `<=`             | left is less or equal to right                                     |
| `gt` or `>`               | left is greater than right                                         |
| `gte` or `>=`             | left is greater than or equal to right                             |
| `in`                      | left exists in right. **incomplete**                               |
| `nin`                     | left does not exists in right. **incomplete**                      |

Functions
---------

Functions are applied on the given object identifier and are later used in the expression.


| Function                  | Description                                                         | Output    |
| :------------------------ | :------------------------------------------------------------------ |-----------|
| @len                      | Gets the length of an array                                         | Integer   |
| @dateTimeLong             | Converts a timestamp to long                                        | Long      |
| @afterCurrentTime         | Checks if the timestamp is after the given timestamp                | Boolean   |
| @min                      | TODO                                                                |           |
| @max                      | TODO                                                                |           |
| @avg                      | TODO                                                                |           |


Errors
------

| Notation                  | Severity                                                           |
| :------------------------ | :----------------------------------------------------------------- |
| `!`                       | Warning                                                            |
| `!!`                      | Severe                                                             |              
| `!!!`                     | Fatal                                                              |

Comments
--------
Inline comments are represented with `#` and can come after a rule statement:

```

~item.name eq nil => [E0001, 'info-1']!! # this is a comment
	
```

Commented out statement looks as follows:
```

# ~item.name eq nil => [E0001, 'info-1']!! # this rule is commented out
	
```


#### TODO:

* Ability to pass custom reference functions in implication
* Ability to have further expression in literal's place
* Ability to configure Error Structure
* Ability to validate Golang structures
* Ability to handle more data types

## License
This project is licensed under the MIT License
