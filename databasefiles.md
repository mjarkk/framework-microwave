# collection files
Inside {collection}.yml are the config files

## Collection configs parts
- **data** = The skeleton of the data
- **premissions** = Who can edit, view, delete what kind of data
- **links** = The data links

## General things
Q: What is `HOST`  
A: It's own document

Q: How do selectors work in `premission` and `links`
A: {a collection OR `HOST`}.{pat to object item} arrays are just seen as objects

## Data
Input:
```yml
data:
  username: required:string:default=someRandomUser
```
Output in collection when creation of user:
```json
{
  "username": "someRandomUser"
}
```

Flags:  
A list of flags that can be used to specify data inside data
- `required` input is require
- `string`, `json`, `int` or `boolean` are requred data types
- `linked` show the input is linked to something else
- `primary` ever array item needs at least one to make linking possible
- `unique` can't be the same in other documents
- `default={some default val}` default val if item is not set
- `min={some int}` set the minimum amound of characters or minimum amound of items in a array (in the case of an array it will start counting by 1)
- `max={some int}` the opposite of `min`
- `regex=/{the regex}/{regex flags}` use regex as matcher
- `reqUppercase` require at least 1 upper case letter
- `reqLowercase` requires at least 1 lower case letter
- `reqSpecial` require at least 1 special character like -, ;, *, ...
- `hash:{hasing algorithm}` hash item (the `max`, `min` effect the string that will get hashed not the final hash)

## Premissions  
Default rules if there are no Premissions:
- View: norules
- Edit: no-one
- Delete: no-one

Ever item can be overwritten by the next item  

Input:
```yml
data:
  foo: required
  bar: required

premissions:
  HOST:
    view: admin
  bar:
    view: norules
```
When requested via browser:
```json
{
  "bar": "bar contents..."
}
```
When requested via browser with admin rights:
```json
{
  "foo": "foo contents...",
  "bar": "bar contents..."
}
```

Select nested item
```yml
premissions:
  some:
    nested:
      item:
        edit: 
          - admin
          - HOST.name
        remove: admin
```

