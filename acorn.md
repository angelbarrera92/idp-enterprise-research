# Acorn

## Enterprise context

You may find different teams pushing code *(apps)* to the same Kubernetes cluster, and you may find different teams using different tools to build and deploy their code.

Using the same cluster for different teams is a good thing *(multi-tenancy)* because it reduces the number of clusters, and it reduces the number of tools that the teams need to learn *(monitoring, logging, ingress, etc).*

Another different approach is to provide a raw Kubernetes cluster to the teams, and let them build and deploy their code using the tooling of their choice. Providing a raw kubernetes cluster to the teams is not a good idea in the long run, because it will lead to a lot of duplication of effort, and it will lead to a lot of wasted time on setup and maintenance *(kubernetes, monitoring, logging, ingress, etc).*

In addition, teams *(in the enterprise)* are constantly changing, and it's not easy to onboard new members nor teams, and it's not easy to update the tooling used by the teams.

So you would need to provide a generic enough solution that fits most use cases. Otherwise, you will end up with a lot of different tools, and a lot of different processes. This includes the security, the build process, the deployment process, etc.

## What to address?

The main goal is to reduce the number of tools and processes, and to reduce the cognitive load on the teams.
Reducing the cognitive load on the teams is important, it will reduce the time spent on non-productive tasks, and it will reduce the time spent on onboarding new members.

In addition, platform services introduce breaking so frequently, so you will need to update the tooling used by the teams to make it compatible, and it's not easy to do that, because it will require a lot of time and effort.

**These changes must be transparent to the teams, and it must be easy to update the tooling used by the teams.**

- Security:
  - Enforce security policies
    - Both at build time and runtime
  - Code quality
    - Linting
    - Testing
- Onboarding:
  - Make it easy for new developers to onboard
    - Even more important for contractors
- Billing:
  - Reduce the amount of time spent on the initial setup
  - Reduce the amount of time spent on maintenance
- Resiliency:
  - Use best practices while running the code

## Goal

- Reduce duplication of effort
- Reduce wasted time on setup and maintenance
- Reduce the number of tools and processes
- Easy to onboard new teams
- Easy to update the tooling used by the teams
- Reduce the cognitive load on the teams


### Enterprise constraints

You may find some of the following constraints in your enterprise context:

- Internal budget
  - Developing a new tool could be expensive
- Internal politics
  - Silos
- Internal resources
  - You may not have enough internal resources to develop a new tool
- Culture:
  - People may not be willing to use a new tool
- 2022-2023:
  - Economic crisis


## Tech approach

Enterprises will require a uniform way to deploy applications; acorn also provides a way of building applications, but seems like it's not the main focus *(and shoulnd't be)*.

The team *(after having some chats with the team)* is working on [re-architecting the build process](https://github.com/acorn-io/acorn/issues/967), and they are thinking about using dagger as the build tool (which is a good thing).

### Acorn

Acorn is kind of a wrapper around different technologies, helm, kustomize, kpt, etc. 
It's a tool that will read a file called `acornfile`, and will use the tooling to build and deploy the defined app in a kubernetes cluster.
This is packaging the app in a container image, and deploying it to a kubernetes cluster generating the kubernetes manifests on the fly.

The Kubernetes cluster must provide basic services like ingress, storage, etc, nothing really fancy these days.
This is something that the most enterprises already have in place.

## Company approach

A platform team will need to provide an abstraction layer built on top of acorn *(maybe integrated with dagger)*, so that the enterprise can provide their own build process, security, etc.

I can foresee an implementation on top of the App CR defined by acorn. It will be used by the acorn controller to deploy the app in the cluster.
This is a nice interface for the developers, so they don't have to know all the Kubernetes manifests, their structure, their changes, etc.

Then, acorn will be used as a uniform deployment interface, this will reduce the number of tools and processes, and it will reduce the cognitive load on the teams (because they will not need to learn a new tool, **nor interact with the cluster directly**).

## Quick benefits

- Reduce the number of tools and processes
- Reduce the cognitive load on the teams
- Easy to onboard new teams
- Easy to update the tooling used by the teams
- Reduce the amount of time spent on the initial setup
- Reduce the amount of time spent on maintenance

## Risks

Acorn is not a mature project, and it's not a well-known project, so it's not easy to adopt it in the enterprise.
Even if it is backed by top community members, it's not easy to adopt it in the enterprise.

Teams may not be willing to stop interacting with the cluster directly, and they may not be willing to learn a new tool.

I already expressed my doubts around implementing a plugin system that allows the enterprise to provide their own plugins, so that, the enterprise can provide their own build process, security, etc. This is something **the acorn team is not considering**, they are focussing on giving a generic enough solution that fits most of the use cases, **then allow enterprises to customize the resulting manifests with mutating webhooks.**

For enterprises, adopting generic solutions is not always easy, companies/enterprises will need to:

- Understand the solution (acorn)
- Understand how to customize it (if possible)
- Understand how to maintain it (Could it run in HA?)
- Understand how to support it (How to get support?)
- Understand how to upgrade it (How to upgrade?)
  - Is it possible to upgrade without downtime?
- Understand how to integrate it with other solutions - THIS IS THE KEY POINT

## Why not acornfile?

Acornfile is a file format that is used to define the app, and it's not yaml, so it's not easy to onboard new teams.
It is a subset of CUE, so it's not easy to onboard new teams.

## References

- build: [the build process currently is broken](https://github.com/acorn-io/acorn/issues/967), but it doesn't matter, because the acorn team is working on it :)
