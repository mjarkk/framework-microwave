# collection files
Inside SomeFolderWithCollections/{collection}.yml are the config/skeleton files for the database.  
These files play a big role in the server because this defines what will be inside the database and what can be requested by who

## Collection configs parts
- **data** = The skeleton of the data
- **premissions** = Who can edit, view, delete what kind of data
- **links** = The data links

## General things
Q: What is `HOST`  
A: It's own document

Q: How do selectors work in `premission` and `links`
A: {a collection OR `HOST`}.{pat to object item} arrays are just seen as objects

Q: what is stored inside the database?
A: The { collection }.yml is just a config file that will be used to check, vind links, etc before data is stored/edited/deleted in the database.

## Data object
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
- `string`, `json`, `int` or `boolean` are requred data 
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
- `hash:{hasing algorithm}` hash item (the `max`, `min` effect the string that will get hashed not the final hash)types
- `json` the content is raw json that won't be stringified so it needs to be valid json! it's also not vilterable with graphql
- `json:graphql` the same as `json` but with graphql support
- `json:raw` the same as `json` but the json will get parsed
- `file` you can post a file to this selector, the framework will then save the file in a folder and the file url will be in the file object
- `check={a list of function names}` add a custom check before the data get saved to the database
- `transformer={a list of function names}` transform the data before it's saved this

NOTE: the flags will be checked and ran after each other from left to right

## Premissions object

### Default rules with no Premissions set
- View: norules
- Edit: no-one
- Delete: no-one

### Rull overwriting
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

Select nested item (because the view is not set it will return to the default value)
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

## links array
a Link is an array with objects  

### Links but there is a copy
Links are not made to litary link data from the origin to a nested object but more to show where the data inside a object/array came from this makes it possible to automaticly copy premissions from the origin and make it possible to automaticly edit all copied object from the origin object/array

### Default rule
```yml
duplicates: true
linksInLinks: false
```

### Settings
Object:
```yml
from: HOST.friends # the object/array to link data
to: users # where it is linked to
duplicates: false # (default=true) allow for duplicates in array? 
linksInLinks: false # (default=false) Allow links in links probebly not a good idea because of the loophole that might be happening
```

### Required
```yml
from: '' # Needs a valid pointer
to: '' # Needs a valid pointer
```

### Example
```yml
links:
  - from: HOST.friends # this object > friends array item
    to: users # use the user collection (the framework will find the right array item automaticly)
    duplicates: false # do not allow duplicates
```

### In witch yaml file do i need to set the link
In the same file as where you placed the link
