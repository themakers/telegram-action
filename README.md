# telegram-action

* TODO: Optimize using binary release?
* https://full-stack.blend.com/how-we-write-github-actions-in-go.html
* https://api.github.com/repos/themakers/telegram-action/releases/latest
```
name: name
author: author
description: description
runs:
  using: node16
  main: invoke-binary.js
```

```javascript
function chooseBinary() {
    // ...
    if (platform === 'linux' && arch === 'x64') {
        return `main-linux-amd64-${VERSION}`
    }
    // ...
}

const binary = chooseBinary()
const mainScript = `${__dirname}/${binary}`
const spawnSyncReturns = childProcess.spawnSync(mainScript, { stdio: 'inherit' })

```


```yaml
---
name: My action
author: My name
description: My description

runs:
  using: docker
  image: Dockerfile





---
steps:
- name: My action
  uses: docker://themakers/telegram-action:latest



```
