# democrablock
[![Go Report Card](https://goreportcard.com/badge/github.com/feditools/democrablock?style=flat-square)](https://goreportcard.com/report/github.com/feditools/democrablock)
[![License](https://img.shields.io/github/license/feditools/democrablock)](https://www.gnu.org/licenses/gpl-3.0.en.html)
[![Release](https://img.shields.io/github/release/feditools/democrablock.svg?style=flat-square)](https://github.com/feditools/democrablock/releases/latest)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/feditools/democrablock)](https://pkg.go.dev/github.com/feditools/democrablock)

A FediBlock list based on voting.

## The Idea
Maintaining a FediBlock list is a lot of work. So lets spread the work around a
bit. Democrablock is a system that allows a group of users known as a 'council' 
to vote on tagging an instance using provided evidence. Voting will be required 
to add a new council member, create a new tag, and apply a tag to an instance.
Decisions will require 50% of current members to vote to become ratified. If the
current number of members is even, 50% + 1 will be required. 