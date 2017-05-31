# Konstrui [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/konstrui)](https://goreportcard.com/report/github.com/arnaucode/konstrui)

web templating engine for static websites, written in Go lang


![konstrui](https://raw.githubusercontent.com/arnaucode/konstrui/master/media/konstrui.png "konstrui")


## Example

- Simple project structure example:

```
webInput/
    konstruiConfig.json
    index.html
    templates/
        userTemplate.html
        userTemplate.json
```

- Set the html file:

```html
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

<konstrui-template html="templates/userTemplate.html" data="templates/userTemplate.json"></konstrui-template>

</body>
</html>

```


- Set the template file:

```html
<div class="class1" id="user[[i]]">
    <div id="username[[i]]" class="class2">{{username}}</div>
    <div id="description[[i]]" class="class2">{{description}}</div>
    <div class="class2">{{phone}}</div>
</div>
```

- Set the template data file:

```json
[{
        "username": "Michaela Doe",
        "description": "Hi, I'm here to code",
        "phone": "456456456"
    },
    {
        "username": "John Doe",
        "description": "Hi, I'm here",
        "phone": "123456789"
    },
    {
        "username": "Myself",
        "description": "How are you",
        "phone": "no phone"
    }
]
```

- Set the configuration file konstruiConfig.json in the webInput directory:

```json
{
    "title": "Web example",
    "author": "arnaucode",
    "github": "github.com/arnaucode",
    "website": "arnaucode.com",
    "files": [
        "index.html",
        "projects.html",
        "app.css"
    ]
}
```


- Execute konstrui

```
./konstrui
```

- Output:

```html
<!DOCTYPE html>
<html>

<body>
    <h1>Heading</h1>
    <p>Paragraph.</p>
    <div class="class1" id="user0">
        <div id="username0" class="class2">Michaela Doe</div>
        <div id="description0" class="class2">Hi, I'm here to code</div>
        <div class="class2">456456456</div>
    </div>
    <div class="class1" id="user1">
        <div id="username1" class="class2">John Doe</div>
        <div id="description1" class="class2">Hi, I'm here</div>
        <div class="class2">123456789</div>
    </div>
    <div class="class1" id="user2">
        <div id="username2" class="class2">Myself</div>
        <div id="description2" class="class2">How are you</div>
        <div class="class2">no phone</div>
    </div>
</body>

</html>
```
