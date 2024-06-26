# Contributing

Thanks for thinking about using or contributing to this software ("Project") and its documentation!

- [Policy & Legal Info](#policy)
- [Getting Started](#getting-started)
- [Submitting an Issue](#submitting-an-issue)
- [Submitting Code](#submitting-code)

## Policy

### 1. Introduction

The project maintainer for this Project will only accept contributions using the 
Developer's Certificate of Origin 1.1 located at [developercertificate.org](https://developercertificate.org) 
("DCO"). The DCO is a legally binding statement asserting that you are the 
creator of your contribution, or that you otherwise have the authority to 
distribute the contribution, and that you are intentionally making the contribution 
available under the license associated with the Project ("License").

### 2. Developer Certificate of Origin Process

Before submitting contributing code to this repository for the first time, you'll 
need to sign a Developer Certificate of Origin (DCO) (see below). To agree to the 
DCO, add your name and email address to the CONTRIBUTORS.md file in this Project. 
At a high level, adding your information to this file tells us that you have the 
right to submit the work you're contributing and indicates that you consent 
to our treating the contribution in a way consistent with the license associated 
with this software (as described in LICENSE.md and its documentation ("Project").

### 3. Important Points

Pseudonymous or anonymous contributions are permissible, but you must be reachable 
at the email address provided in the Signed-off-by line.

If your contribution is significant, you are also welcome to add your name and 
copyright date to the source file header.

If you are a U.S. Federal government employee and use a `*.mil` or `*.gov` email 
address, we interpret your Signed-off-by to mean that the contribution was created
in whole or in part by you and that your contribution is not subject to copyright 
protections.

### 4. DCO Text

The full text of the DCO is included below and is available online at [developercertificate.org](https://developercertificate.org):

```txt
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

## Getting Started

go-usajobs is an unofficial Go client library and command-line interface for the
USAJobs.gov API. A USAJobs.gov API Token can be requested from https://developer.usajobs.gov/apirequest/.

This project is written in Go and uses a variety of tools for builds, tests, scans,
and releases. To ease development environment setup, this project uses [devbox](https://www.jetify.com/devbox/docs/installing_devbox/)

After installing `devbox`, clone this git repository, cd into go-usajobs, and
run `devbox shell`. The development environment should be ready after a few seconds!

All development environment tools and dependencies are required to be managed in devbox.

### Code Style

As this project is developed in Go, code must be conformant to gofmt styling. 

Code not conformant to the style standard will be rejected.

## Submitting an Issue

You should feel free on our GitHub repository for anything you find that needs attention. 
That includes content, functionality, design, or anything else!

### Submitting a Bug Report

When submitting a bug report on the website, please be sure to include accurate 
and thorough information about the problem you're observing. Be sure to include:

- Steps to reproduce the problem,
- The CLI Command or Client Function where you observed the problem,
- What you expected to happen,
- What actually happend (or didn't happen), and

## Submitting Code

We appreciate everyone that makes the effort to contribute to this project! To make
a contribution, fork the project, clone your fork locally, create a new branch, make your 
code changes (along with unit tests!), push your local changes back to your fork,
and, assuming everything checks out, submit a PR back to this project.

Your pull request will be reviewed by the team and go through some automated checks 
using a continuous integration and deployment tool.

After review by the team, your PR will either be commented on with a request for 
more information or changes, or it will be merged into the `main` branch and for
the next version release. 


