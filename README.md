<p align="center">
  <img src="/logo.gif" height="300" />
</p>
<p align="center">
  Open source Git repository template
</p>
<p align="center">
  <a href="https://github.com/misitebao/standard-repository/blob/main/LICENSE">
    <img alt="GitHub" src="https://img.shields.io/github/license/misitebao/standard-repository?style=flat-square"/>
  </a>
  <a href="https://github.com/misitebao/standard-repository">
    <img alt="GitHub" src="https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat-square.svg"/>
  </a>
  <a href="https://github.com/misitebao/standard-repository">
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/misitebao/standard-repository?style=flat-square"/>
  </a>
  <a href="https://github.com/misitebao/standard-repository/releases">
    <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/misitebao/standard-repository?include_prereleases&sort=semver&style=flat-square">
  </a>
  <a href="https://github.com/misitebao">
    <img alt="GitHub user" src="https://img.shields.io/badge/author-misitebao-brightgreen?style=flat-square"/>
  </a>
  <a href="https://github.com/misitebao/standard-repository/actions/workflows/pre-build.yml">
    <img alt="Pre Build" src="https://img.shields.io/github/workflow/status/misitebao/standard-repository/Pre%20Build%20%7C%20预构建/main?style=flat-square&logo=github"/>
  </a>
  <a href="https://gitter.im/misitebao/standard-repository">
    <img alt="Gitter" src="https://img.shields.io/gitter/room/misitebao/standard-repository?style=flat-square&color=4ab494"/>
  </a>
</p>

<span id="nav-1"></span>

## Internationalization

[English](README.md) | [简体中文](README.zh-Hans.md)

<span id="nav-2"></span>

## Table of Contents

<details>
  <summary>Click me to Open/Close the directory listing</summary>

- [Internationalization](#nav-1)
- [Table of Contents](#nav-2)
- [Introductions](#nav-3)
  - [Official Website](#nav-3-1)
  - [Background](#nav-3-2)
- [Graphic Demo](#nav-4)
- [Features](#nav-5)
- [Architecture](#nav-6)
- [Getting Started](#nav-7)
- [Maintainer](#nav-8)
- [Contributors](#nav-9)
- [Community Exchange](#nav-10)
- [Part Of Users](#nav-11)
- [Release History](CHANGE.md)
- [Donators](#nav-12)
- [Sponsors](#nav-13)
- [Special Thanks](#nav-14)
- [License](#nav-15)

</details>

<span id="nav-3"></span>

## Introductions

This project is a Github sample warehouse template, the main content is the sample template of README file.

<span id="nav-3-1"></span>

### Official Website

[Official Website Link](https://standard-repository.vercel.app)

<span id="nav-3-2"></span>

### Background

When searching for the required tools and libraries in the open source community, I found that there are many excellent code libraries, but there is no good README file or tutorial, which causes users to spend extra time to learn how to use it, so this project provides A standard code base template, hoping to help others.

<span id="nav-4"></span>

## Graphic Demo

[![Click Gif to view the full demo](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/md/template-git-repository-mini.gif)](https://www.youtube.com/embed/bOE3eJ-1eas)

<span id="nav-5"></span>

## Features

- The project logo and corresponding data are displayed in the center
- Provide multi-language functions and sample templates
- README must-have instructions
- Built-in directory navigation function to solve the problem that some Markdown parsing engines cannot parse navigation correctly

<span id="nav-6"></span>

## Architecture

```
|—— .gitee                          Gitee Configuration File
| |—— ISSUE_TEMPLATE.md             Gitee Issue Template
| |—— PULL_REQUEST_TEMPLATE.md      Gitee PR Template
|—— .github                         Github Configuration File
| |—— ISSUE_TEMPLATE                Github Issue Template
| | |—— issue-template-bug.md       Github Issue Bug Template
| | |—— issue-template-feature.md   Github Issue Feature Template
| |—— workflows                     Github Workflows
| | |—— deploy-for-hugo.yml         Github Workflows Hugo Example
| | |—— deploy-for-nodejs.yml       Github Workflows NodeJS Example
| |—— pull-request-template.md      Github PR Template
|—— website                         Project website
|—— CHANGELOG.md                    Release Log
|—— LICENSE                         LICENSE
|—— README.md                       English README
|—— README.zh-Hans.md               Other Language README
|—— README.tmpl.md                  README Template

```

<span id="nav-7"></span>

## Getting Started

[Click me](/copy-template/README.tmpl.md) to view the template file, and then click Edit to copy the template content.

If your README conforms to the standard-repository, you can add a badge to link back to this specification to help others adopt this readme.

[![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat-square.svg)](https://github.com/misitebao/standard-repository)

To add in markdown format, please use the following code:

```markdown
[![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat-square.svg)](https://github.com/misitebao/standard-repository)
```

To add in HTML format, please use the following code:

```html
<a href="https://github.com/misitebao/standard-repository">
  <img
    alt="GitHub"
    src="https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat-square.svg"
  />
</a>
```

Please modify the file name yourself to get the style you want:

| File Nmae               | Style Preview                                                                                                                                   |
| ----------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| badge_flat-square.svg   | ![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat-square.svg)   |
| badge_flat.svg          | ![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat.svg)          |
| badge_for-the-badge.svg | ![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_for-the-badge.svg) |
| badge_plastic.svg       | ![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_plastic.svg)       |
| badge_social.svg        | ![Readme file conforming to standard-repository](https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_social.svg)        |

<span id="nav-8"></span>

## Maintainer

Thanks to the maintainers of these projects:

<a href="https://github.com/misitebao"><img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/></a>

<details>
  <summary>Click me to Open/Close the contributors listing</summary>

- [Misitebao](https://github.com/misitebao) - Project author, full stack engineer.

</details>

<span id="nav-9"></span>

## Contributors

Thank you to all the contributors who participated in the development of standard-repository. [Contributors](https://github.com/misitebao/standard-repository/graphs/contributors)

<a href="https://github.com/crushonyou18"><img src="https://github.com/crushonyou18.png" width="40" height="40" alt="crushonyou18" title="crushonyou18"/></a>

<span id="nav-10"></span>

## Community Exchange

- [Gitter](https://gitter.im/misitebao/standard-repository) - A public instant chat tool.

<span id="nav-11"></span>

## Part Of Users

- [Tigago](https://github.com/tigateam/tigago) - A modular framework based on Go language
- [wails-template-vue](https://github.com/misitebao/wails-template-vue) - A wails template based on Vue and Vue-Router

> If your project uses this specification, you can fork this project and [submit a PR](https://github.com/misitebao/standard-repository/pulls) to add your project to this list, so that others can know your project.

<span id="nav-12"></span>

<!-- ## Donators -->

<span id="nav-13"></span>

<!-- ## Sponsors -->

<span id="nav-14"></span>

<!-- ## Special Thanks -->

<span id="nav-15"></span>

## License

[License MIT](LICENSE)
