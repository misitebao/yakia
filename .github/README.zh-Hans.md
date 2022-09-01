<h1 align="center">Yakia</h1>

<p align="center">
  <a href="https://github.com/misitebao/yakia/blob/main/LICENSE">
    <img alt="GitHub" src="https://img.shields.io/github/license/misitebao/yakia"/>
  </a>
  <a href="https://github.com/misitebao/yakia">
    <img alt="GitHub" src="https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_flat.svg"/>
  </a>
  <a href="https://www.codefactor.io/repository/github/misitebao/yakia">
    <img src="https://www.codefactor.io/repository/github/misitebao/yakia/badge" alt="CodeFactor" />
  </a>
  <a href="https://pkg.go.dev/github.com/misitebao/yakia">
    <img src="https://pkg.go.dev/badge/github.com/misitebao/yakia.svg" alt="Go Reference"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/misitebao/yakia">
    <img src="https://goreportcard.com/badge/github.com/misitebao/yakia" alt="Go Report Card"/>
  </a>
  <a href="https://github.com/misitebao/yakia/actions/workflows/pre-build.yml">
    <img alt="Pre Build" src="https://img.shields.io/github/workflow/status/misitebao/yakia/Pre%20Build%20%7C%20预构建/main?logo=github"/>
  </a>
  <a href="https://app.slack.com/client/T029RQSE6/C03RQSD4KA7">
    <img alt="Slack" src="https://img.shields.io/badge/slack-gophers%2Fyakia%20-blue?logo=slack"/>
  </a>
  <br/>
  <a href="https://pkg.go.dev/github.com/misitebao/yakia/cmd/yakia">
    <img alt="GitHub tag (latest SemVer pre-release)" src="https://img.shields.io/github/v/tag/misitebao/yakia?include_prereleases&label=pkg.go.dev"/>
  </a>
  <a href="https://www.npmjs.com/package/yakia">
    <img alt="npm" src="https://img.shields.io/npm/v/yakia"/>
  </a>
  <a href="https://crates.io/crates/yakia">
    <img alt="Crates.io" src="https://img.shields.io/crates/v/yakia"/>
  </a>
</p>

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md)

</samp>
</strong>
</div>

## 内容目录

<details>
  <summary>点我 打开/关闭 目录列表</summary>

- [国际化](#国际化)
- [内容目录](#内容目录)
- [项目介绍](#项目介绍)
  - [官方网站](#官方网站)
  - [背景](#背景)
- [图形演示](#图形演示)
- [功能](#功能)
- [架构](#架构)
- [快速入门](#快速入门)
- [维护者](#维护者)
- [贡献者](#贡献者)
- [社区交流](#社区交流)
- [部分用户](#部分用户)
- [发布记录](#发布记录)
- [捐赠者](#捐赠者)
- [赞助商](#赞助商)
- [特别感谢](#特别感谢)
- [许可证](#许可证)

</details>

## 项目介绍

此项目是一个 Github 示例仓库模板，主要内容为 README 的示例模板。

### 官方网站

[官网地址](https://yakia.vercel.app)

### 背景

在开源社区寻找所需要的工具和库的时候，发现有很多优秀的代码库，但是没有一个好的 README 文件或者使用教程，导致使用者需要花额外的时间去学习如何使用它，所以本项目提供了一个标准的代码库范本，希望可以帮助到他人。

## 图形演示

[![单击 Gif 查看完整演示](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/template-git-repository-mini.gif)](https://www.bilibili.com/video/BV1d64y1B7pe?share_source=copy_web)

## 功能

- 项目 logo 以及相应数据居中展示
- 提供多语言功能以及示例模板
- README 文件必备的说明
- 内置目录导航功能，解决部分 Markdown 解析引擎不能正确解析导航的问题

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

## 快速入门

如果您的 README 符合标准存储库，您可以添加一个徽章链接回此规范，以帮助其他人采用此自述文件。

[![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_flat-square.svg)](https://github.com/misitebao/yakia)

以 Markdown 格式添加，请使用以下代码：

```markdown
[![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_flat-square.svg)](https://github.com/misitebao/yakia)
```

以 HTML 格式添加，请使用以下代码：

```html
<a href="https://github.com/misitebao/yakia">
  <img
    alt="GitHub"
    src="https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_flat-square.svg"
  />
</a>
```

请自行修改文件名以获取你想要的样式：

| 文件名                  | 样式预览                                                                                             |
| ----------------------- | ---------------------------------------------------------------------------------------------------- |
| badge_flat-square.svg   | ![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_flat-square.svg)   |
| badge_flat.svg          | ![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_flat.svg)          |
| badge_for-the-badge.svg | ![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_for-the-badge.svg) |
| badge_plastic.svg       | ![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_plastic.svg)       |
| badge_social.svg        | ![符合 yakia 的自述文件](https://cdn.jsdelivr.net/gh/misitebao/yakia/assets/badge_social.svg)        |

## 维护者

感谢这些项目的维护者：

<a href="https://github.com/misitebao"><img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/></a>

<details>
  <summary>点我 打开/关闭 维护者列表</summary>

- [米司特包](https://github.com/misitebao) - 项目作者，全栈工程师。

</details>

## 贡献者

感谢所有参与 yakia 开发的贡献者。[贡献者列表](https://github.com/misitebao/yakia/graphs/contributors)

<a href="https://github.com/crushonyou18"><img src="https://github.com/crushonyou18.png" width="40" height="40" alt="crushonyou18" title="crushonyou18"/></a>

## 社区交流

- [Gitter](https://gitter.im/misitebao/yakia) - 一个公开的即时聊天工具。

## 部分用户

- [Tigago](https://github.com/tigateam/tigago) - 一个基于 Go 语言的模块化框架
- [wails-template-vue](https://github.com/misitebao/wails-template-vue) - 基于 Vue 和 Vue-Router 的 Wails 模板

> 如果你的项目使用了此规范，你可以 Fork 此项目并[提交 PR](https://github.com/misitebao/yakia/pulls) 将您的项目添加在此列表，以便于让其他人知道您的项目。

<!-- ## 发布记录 -->

<!-- ## 捐赠者 -->

<!-- ## 赞助商 -->

<!-- ## 特别感谢 -->

## 许可证

[License MIT](../LICENSE)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmisitebao%2Fyakia.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmisitebao%2Fyakia?ref=badge_large)
