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
- `string`, `json`, `int`, `byteArr` or `boolean` are requred data
- `linked` show the input is linked to something else
- `primary` ever array item needs at least one to make linking possible (this also makes this vaiable `unique`)
- `unique` can't be the same in other documents
- `default={some default val}` default val if item is not set
- `min={some int}` set the minimum amound of characters or minimum amound of items in a array (in the case of an array it will start counting by 1)
- `max={some int}` the opposite of `min`
- `regex=/{the regex}/` use regex as matcher (This needs to be a regex understandable by golang)
- `reqUppercase` require at least 1 upper case letter
- `reqLowercase` requires at least 1 lower case letter
- `reqSpecial` require at least 1 special character like -, ;, *, ...
- `hashPassword` Hash a password using a secure password hasing algorithm (pbkdf2)
- `hash={hasing algorithm}` hash item using a supported algorithm see [safety/README.md](https://github.com/mjarkk/framework-microwave/blob/master/pkg/safety/README.md) (the `max`, `min` effect the string that will get hashed not the final hash)
- `json` the content is raw json that won't be stringified so it needs to be valid json! it's also not vilterable with graphql
- `json:graphql` the same as `json` but with graphql support
- `json:raw` the same as `json` but the json will get parsed
- `file` you can post a file to this selector, the framework will then save the file in a folder and the file url will be in the file object
- `check={a list of function names}` add a custom check before the data get saved to the database (example: `check=foo,bar,baz`)
- `transformer={a list of function names}` transform the data before it's saved this (example: `transformer=foo,bar,baz`)

NOTE: the flags will be checked and ran after each other from left to right

## Premissions object

### Default rules with no Premissions set

- View: norules
- Edit: no-one
- Delete: no-one

### Rull overwriting

A nested object item can overwrite the rules of it's own object
Example

```yml
HOST:
  view:
    - admin
username:
  view:
    - HOST.username # this can
```

### Demo

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

### Premissions functions, rules, selectors

Every premissions array item has that rule  

```yml
view:
  - admin
  - HOST.name
```

Group access is defined as the group name

```yml
- admin
```

Group spesific access

```yml
- users=someRandomUser
```

Function based premissions

```yml
- func=foo,bar,baz
```

User's data is in object  
This searches if the current user making the request has the premission to view an array

```yml
- HOST.name # HOST.{some}.{path}.{to}.{an}.{item}
```

Data in object matches something

```yml
- HOST.hidden=false
```

Data in object matches regex  
This needs to be a regex understandable by golang

```yml
- HOST.hidden=/^some\s*matcher$/
```

If data doesn't match `(!)...` or `(!1)...`

```yml
friends:
  view:
    - friends & friendsFromFriends & (!)func=blockedUsers
    - admin
```

Ignore all other rules if `(!!)...` or `(!2)...` failes
To still allow a rule over `(!2)` use `(!!!)` or `(!3)`  
The example below shows if someone in the group admin is blacklisted they can still access the data but if someone is allowed by the friendsFromFriends function and are blacklisted they won't beable to view the data  
TODO think about a better `(!), (!!) and (!!!)`

```yml
friends:
  view:
    - (!2)func=blockedUsers
    - friends
    - func=friendsFromFriends
    - (!3)admin
```

Someone needs `foo` and `bar` before they have access

```yml
- HOST.name:HOST.hidden=false
```

### Some other notes

- If something has access to edit an colection item but not has no access to view they are automaticly disalowed to eddit the contents of that item

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
