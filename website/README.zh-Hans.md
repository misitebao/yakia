# 网址

这个网站是使用 [Docusaurus 2](https://docusaurus.io/)构建的，一个现代静态网站生成器。

### 安装

```
$ yarn
```

### 地方发展

```
$ yarn start
```

### 构建...

```
$ yarn building
```

此命令将静态内容生成到 `build` 目录中，并且可以使用任何静态内容托管服务。

### 部署

```
$ GIT_USER=<Your GitHub username> USE_SSH=true yarn deploy
```

如果您正在使用 GitHub 页面作为主机， 此命令是构建网站和推送到 `gh-pages` 分支的方便方式。
