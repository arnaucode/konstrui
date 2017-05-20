# Timmy [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/timmy)](https://goreportcard.com/report/github.com/arnaucode/timmy)

web templating engine for static websites, written in Go lang


![timmy](https://raw.githubusercontent.com/arnaucode/timmy/master/timmy.png "timmy")


## Example

- Simple project structure example:

```
webInput/
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

<timmy-template html="templates/userTemplate.html" data="templates/userTemplate.json"></timmy-template>

</body>
</html>

```


- Set the template file:

```html
<div class="class1">
    <div class="class2">{{username}}</div>
    <div class="class2">{{description}}</div>
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

- Execute Timmy

```
./timmy
```

- Output:

```html
<!DOCTYPE html>
<html>

<body>
    <h1>My First Heading</h1>
    <p>My first paragraph.</p>
    <div class="class1">
        <div class="class2">Michaela Doe</div>
        <div class="class2">Hi, I'm here to code</div>
        <div class="class2">456456456</div>
    </div>
    <div class="class1">
        <div class="class2">John Doe</div>
        <div class="class2">Hi, I'm here</div>
        <div class="class2">123456789</div>
    </div>
    <div class="class1">
        <div class="class2">Myself</div>
        <div class="class2">How are you</div>
        <div class="class2">no phone</div>
    </div>
</body>

</html>
```
