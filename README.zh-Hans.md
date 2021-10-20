<p align="center">
  <img src="/logo.png" height="180" />
</p>
<p align="center">
  开源Git存储库模板
</p>
<p align="center">
  <a href="https://github.com/misitebao/standard-repository/blob/main/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/misitebao/standard-repository?style=flat-square"/></a>
  <a href="https://github.com/misitebao/standard-repository"><img alt="GitHub" src="https://img.shields.io/badge/Readme--Style-standard--repository-brightgreen?style=flat-square&color=f83500"/></a>
  <a href="https://github.com/misitebao/standard-repository"><img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/misitebao/standard-repository?style=flat-square"/></a>
  <a href="https://github.com/misitebao"><img alt="GitHub user" src="https://img.shields.io/badge/author-misitebao-brightgreen?style=flat-square"/></a>
  <a href="https://gitter.im/misitebao/standard-repository"><img alt="Gitter" src="https://img.shields.io/gitter/room/misitebao/standard-repository?style=flat-square&color=4ab494"/></a>
</p>

<span id="nav-1"></span>

## 国际化

[English](README.md) | [简体中文](README.zh-Hans.md)

<span id="nav-2"></span>

## 内容目录

<details>
  <summary>点我 打开/关闭 目录列表</summary>

- [国际化](#nav-1)
- [内容目录](#nav-2)
- [项目介绍](#nav-3)
  - [官方网站](#nav-3-1)
  - [背景](#nav-3-2)
- [图形演示](#nav-4)
- [功能特色](#nav-5)
- [架构](#nav-6)
- [新手入门](#nav-7)
- [维护者](#nav-8)
- [贡献者](#nav-9)
- [社区交流](#nav-10)
- [部分用户](#nav-11)
- [发布记录](CHANGE.md)
- [捐赠者](#nav-12)
- [赞助商](#nav-13)
- [特别感谢](#nav-14)
- [版权许可](#nav-15)

</details>

<span id="nav-3"></span>

## 项目介绍

此项目是一个 Github 示例仓库模板，主要内容为 README 的示例模板。

<span id="nav-3-1"></span>

### 官方网站

[官网地址](https://standard-repository.vercel.app)

<span id="nav-3-2"></span>

### 背景

在开源社区寻找所需要的工具和库的时候，发现有很多优秀的代码库，但是没有一个好的 README 文件或者使用教程，导致使用者需要花额外的时间去学习如何使用它，所以本项目提供了一个标准的代码库范本，希望可以帮助到他人。

<span id="nav-4"></span>

## 图形演示

[![单击 Gif 查看完整演示](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/md/template-git-repository-mini.gif)](https://www.bilibili.com/video/BV1d64y1B7pe?share_source=copy_web)

<span id="nav-5"></span>

## 功能特色

- 项目 logo 以及相应数据居中展示
- 提供多语言功能以及示例模板
- README 文件必备的说明
- 内置目录导航功能，解决部分 Markdown 解析引擎不能正确解析导航的问题

<span id="nav-6"></span>

## 架构

```
|—— .gitee                          Gitee 配置文件
| |—— ISSUE_TEMPLATE.md             Gitee Issue 模板
| |—— PULL_REQUEST_TEMPLATE.md      Gitee PR 模板
|—— .github                         Github 配置文件
| |—— ISSUE_TEMPLATE                Github Issue 模板
| | |—— issue-template-bug.md       Github Issue Bug 模板
| | |—— issue-template-feature.md   Github Issue Feature 模板
| |—— workflows                     Github 工作流
| | |—— deploy-for-hugo.yml         Github 工作流 Hugo 示例
| | |—— deploy-for-nodejs.yml       Github 工作流 NodeJS 示例
| |—— pull-request-template.md      Github PR 模板
|—— website                         项目网站
|—— CHANGELOG.md                    发布日志
|—— LICENSE                         许可证
|—— README.md                       英语 README
|—— README.zh-Hans.md               其他语言 README
|—— README.tmpl.md                  README 模板

```

<span id="nav-7"></span>

## 新手入门

[点我](/edit/main/copy-template/README.zh-Hans.tmpl.md) 查看模板内容。

如果您的 README 符合标准存储库，您可以添加一个徽章链接回此规范，以帮助其他人采用此自述文件。

[![符合 standard-repository 的自述文件](https://img.shields.io/badge/Readme--Style-standard--repository-brightgreen?style=flat-square&color=f83500)](https://github.com/misitebao/standard-repository)

以 Markdown 格式添加，请使用以下代码：

```markdown
[![符合 standard-repository 的自述文件](https://img.shields.io/badge/Readme--Style-standard--repository-brightgreen?style=flat-square&color=f83500)](https://github.com/misitebao/standard-repository)
```

以 HTML 格式添加，请使用以下代码：

```html
<a href="https://github.com/misitebao/standard-repository">
  <img
    alt="GitHub"
    src="https://img.shields.io/badge/Readme--Style-standard--repository-brightgreen?style=flat-square&color=f83500"
  />
</a>
```

<span id="nav-8"></span>

## 维护者

感谢这些项目的维护者：

<a href="https://github.com/misitebao"><img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/></a>

<details>
  <summary>点我 打开/关闭 贡献者列表</summary>

- [米司特包](https://github.com/misitebao) - 项目作者，全栈工程师。

</details>

<span id="nav-9"></span>

## 贡献者

感谢所有参与 standard-repository 开发的贡献者。[贡献者列表](https://github.com/misitebao/standard-repository/graphs/contributors)

<span id="nav-10"></span>

## 社区交流

- [Gitter](https://gitter.im/misitebao/standard-repository) - 一个公开的即时聊天工具。

<span id="nav-11"></span>

## 部分用户

- [Tigago](https://github.com/tigateam/tigago) - 一个基于 Go 语言的模块化框架
- [wails-template-vue](https://github.com/misitebao/wails-template-vue) - 基于 Vue 和 Vue-Router 的 Wails 模板

> 如果你的项目使用了此规范，你可以 Fork 此项目并[提交 PR](https://github.com/misitebao/standard-repository/pulls) 将您的项目添加在此列表，以便于让其他人知道您的项目。

<span id="nav-12"></span>

## 捐赠者

<span id="nav-13"></span>

## 赞助商

<span id="nav-14"></span>

## 特别感谢

<span id="nav-15"></span>

## 版权许可

[License MIT](LICENSE)
