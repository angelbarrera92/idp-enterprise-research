# Dagger in the enterprise

## Enterprise context

You have many development teams along with the infra/security/networking/cloud teams.
Each team has its codebase, and its way of working, including its tools and processes.
App/Development teams use the tooling provided by the infra/security/networking/cloud teams to build their codebase. The tooling is provided as a service, and the teams are expected to use it. The tooling is inner source, and the teams are not expected to modify it.

Tooling is based on open-source software with some constraints, configuration, and customization.
So your development must understand the tooling, but sometimes they are not able to understand the tooling's constraints as they are not well documented.
For example:
  - Developers have access to a drone CI server, but they can not configure it, they can only use it with just a few agent images, network, and storage configurations.
  - Developers could have access to a kubernetes cluster, but they can not configure it, they can only use it with just a few agent images, network, and storage configurations.

## What to address?

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

I'll go with a heroku/fly.io 2.0 vision. Meaning that the first iteration of the tooling just needs to:

Starting from a simple go web app:

- Build a go API
- Wrap it into an OCI/Artifact image
- Deploy it to a runtime environment

based on a procfile-like file:

```yaml
service:
  type: go # Ideally, automatically detected
  build: go build -o bin/web ./cmd/web
  run: bin/web
security: {} # Default value
deploy:
  type: kubernetes
  environment: playground
```

As you can see, the procfile is very simple, and it's not tied to any specific tooling.
That is where dagger comes in, it's a tool that will read the procfile, and will use the tooling provided by the enterprise to build and deploy the app.

Each team will create its plugins using dagger as the foundation, and the engine/platform-cli will be the one that will use the plugins to build and deploy the app based on the procfile.
Meaning that, on this first iteration, we will have:

- go build plugin
- default security plugin (if not set, it will use the default security policy)
- kubernetes deploy plugin

The next iteration will be to add more plugins, and to make the procfile more flexible, for example:

```yaml
service:
  type:
    name: go # Ideally, automatically detected
    version: 1.16
  build: go build -o bin/web ./cmd/web
  run: bin/web
security: {} # Default value
deploy:
  type: kubernetes
  environment: playground
  replicas: 2
  autoscaling:
    min: 1
    max: 10
    cpu: 80
    memory: 80
```

### Dagger

Even if you are writing code (for the plugins and the platform-cli), you probably won't need to do all by writing code, then dagger will:

- ease the creation of the plugins as it will provide a simple container image API.
  - This will allow you to reuse container images currently in use by the teams in their pipelines.

## Company approach

Renaming the departments to be "platform" and "product" teams is not enough. You need to change the way the teams work together.
Look for the right people to lead the platform team, and make sure that they have the right skills and experience.

These teams need to work together, and they need to be able to communicate with each other, furthermore, they need to be able to understand client (app teams) needs.

## Quick benefits

- Developers can easily run the CI/CD pipeline locally
- The platform team can easily update the CI/CD pipeline, and the developers will get the updates automatically
  - This will reduce the amount of time spent on maintenance
  - This will reduce the amount of time spent on onboarding new teams
- Easy to develop and test the CI/CD pipeline (did you try to write Jenkins tests? or Gitlab CI tests?)
- Easy to evolve by the platform team

## Risks

- Dagger is a new tool, and it's not battle tested
- Teams may not be willing to develop their plugins
  - Need a good documentation
  - Need a good community

## Why not CUE?

CUE is a great tool, but it's not a good fit for this use case.
Providing new technology to the teams is not a good idea, they will not be willing to learn a new language, and they will not be willing to use a new tool.
Instead, using a simple procfile-like file, and a simple language like yaml, will make it easy for the teams to onboard.
In addition, testing go code is easier than testing CUE code (not sure if it's possible to test CUE code at all).

### Quick and dirty POC

- [dagger-chainguard](dagger-chainguard): It could be the foundation of the "go build" plugin.

```bash
$ cd dagger-chainguard
# First, login to the docker registry
# $ docker login angelbarrera92 -p <password>
$ go run main.go go-build --main example/main.go --oci-name angelbarrera92/hello -o hello
...
..
#5 DONE 0.2s

#6 exporting to image
#6 exporting layers
#6 exporting layers 0.3s done
#6 exporting manifest sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d 0.0s done
#6 exporting config sha256:9138258f012b2d4d7ab5b1a506ef1cf3905ec2a16717a4f5781c8a9f3457813e 0.0s done
#6 pushing layers
#6 pushing layers 8.2s done
#6 pushing manifest for docker.io/angelbarrera92/hello:latest@sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d
Published image docker.io/angelbarrera92/hello@sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d
$ docker run --rm docker.io/angelbarrera92/hello@sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d /hello
Unable to find image 'angelbarrera92/hello@sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d' locally
docker.io/angelbarrera92/hello@sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d: Pulling from angelbarrera92/hello
49be0ded2c08: Pull complete 
e57372611676: Pull complete 
Digest: sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d
Status: Downloaded newer image for angelbarrera92/hello@sha256:5b551974220a00f96aaf412dea34d4e721f300fd73354bdf839ef6154487879d
Hello World
```

What it does is build a go app, and wrap it into an OCI image. Both the building and the wrapping are done using the [chainguard container images](https://github.com/chainguard-images).

If the go tech lead decides to use any other build base image, then the chainguard image can be replaced by the new image without impacting the developers.

## References

- https://docs.dagger.io/1220/vs/


> Dagger does not replace your CI: it improves it by adding a portable development layer on top of it.
> Dagger runs on all major CI products. This reduces CI lock-in: you can change CI without rewriting all your pipelines.
> Dagger also runs on your dev machine. This allows dev/CI parity: the same pipelines can be used in CI and development.
