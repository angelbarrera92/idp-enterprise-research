# Outcome

After testing a few different technologies during a week, focusing on enterprise aspects like:

- Internal budget
- Internal politics and silos
- Internal resources
- Culture
- 2022-2023: Economic crisis

I'll write what I think could be a good approach to developing an enterprise IDP.

## First phase: Analyze

First, you need to understand the current way of working of the teams in your enterprise, and how fast they can respond to the market.

Imagine you are working in a bank company, a new regulation could be released, and you need to implement it in your codebase.

- How fast can you do it?
- How much time do you need to spend on it?
- How many developers do you need to implement it?
  - Do you need to hire new developers?
  - Are you outsourcing the work?

You probably face thes kinds of problems every day, and you probably have a lot of experience in solving them. But, what about the maintenance of the projects?

**In an event of a security breach, how fast can you fix it?**

Usually, the development is outsourced, and the maintenance is done by different teams. These teams use to support a lot of projects, and they are not always able to fix the problems on time.

> We all know log4j vulnerability is still not fixed in many projects.

As these projects are developed using different technologies *(and different versions)*, fixing a problem in one project in a specific way, doesn't mean that the same problem in another project can be fixed in the same way.
This leads to a lot of duplication of effort, and a lot of wasted time. **Without forgetting the security risks.**

Then, you'll need to understand that these kinds of changes, where developers are used to working in a specific way, are hard to implement.
So there are many ways to enforce it, the two most common are:

- The fast way (top-down): impose it. A CTO - or a manager - can impose a new way of working on the teams.
- The slow way (bottom-up): convince them by providing them with a better way of working, and by showing them the benefits of it.

## Second phase: Plan and design

Then, you'll need to plan the implementation of the new way of working.
This is selecting the right technology and the right way of implementing it.

The three technologies I've tested are:

- [Dagger](dagger.md)
- [acorn](acorn.md)
- [Backstage](backstage.md)

But, this doesn't mean that these are the only technologies that can be used to build an IDP.
In the end, these tools are just tools, and they can be used to build a lot of different things.

The really important thing is to understand the needs of the teams and to understand how the teams are going to use the IDP.
So designing the IDP is not just about selecting the right technology, but also about understanding the needs of the teams.

Sometimes, the teams are not able to tell you what they need, and you'll need to understand it by yourself.
Sometimes, tech leads can help you understand the needs of the teams, but they are not always able to tell you what the teams need.

In addition, we all know that these changes must integrate already existing enterprise services, meaning that if you already have a CI/CD pipeline, you'll need to integrate it with the IDP.
This is the main reason why I've chosen Backstage, because it is the only tool that I've found that can be integrated with existing enterprise services.

Then, you'll need to think about the UX both from the developers' and the platform owners' perspectives.

If you want to onboard new developers, you'll need to provide them with a good UX, and you'll need to provide them with good documentation.

How many requests/tickets do you have to open to get a new developer onboarded?
How many requests/tickets do you have to open to get new repositories created for a new project?
How many requests/tickets do you have to open to get a new domain/cname created for a new project?
...

The important thing is to understand that the UX is not just about the developers, but also about the platform owners.
Platform teams also need to attend to a lot of requests, and they need to be able to do it promptly.

Can you spot the amount of cognitive load that platform teams have to deal with? And the waste of time?


### UX for developers

There was a time when developers were happy just to develop their codebase, and they didn't care about the platform they were using.
They use to bundle their codebase, compile it, and deploy it somewhere. The ops team was responsible for the platform, and the developers were responsible for their codebase.

Then, technologies like Docker and Kubernetes came, and developers started to care about the platform they were using... until they realized that they were spending a lot of time maintaining the platform, and not enough time developing their codebase.

Cognitive load is a real thing, and it is a real problem these days.

If you ask a developer for the preferred way of working, they'll probably tell you that they want to focus on their codebase, and not on the platform.
Something like "heroku" would be perfect for them, but this was not possible in the past, because the platform was not flexible enough.

I would like to ask you a question: **how many developers do you have in your company?**
We all know that there is no just one way of solving a problem, and we all know that there is no just one way of working.

So standardizing is the only way to scale, and optimize the way of working.

**For sure it will be corner cases, it is key to understand that the platform must be flexible enough to support these corner cases without compromising the standardization.**

### Third phase: The team

You'll need to create a new team. You should avoid the common mistake of creating a new team from scratch (hiring new people). It's key to identify the right people in the company and move them to the new team.
The team must be composed of people with different skills, and with different backgrounds.

- Infrastructure
- Developers
  - Frontend
  - Backend
- Security
- Networking
- Business
- Tech writers
...

This team must be flexible in terms of capacity, and it must be able to scale up and down based on the needs of the company.
In addition, a key profile is needed: the product owner of the IDP. This person must be able to understand the needs of the teams and to understand how the teams are going to use the IDP.
Then this person must be responsible for the roadmap of the IDP, and for prioritization of the features.
In the context of a large enterprise, this person must be able to understand the company strategy and to understand how the IDP can help the company to achieve its goals.
Also, this person must have enough power to enforce the new way of working and to enforce the new way of developing software.

Another pillar of the teams would be the enablement team. This team must be responsible for the documentation, and the training of the developers.

## Fourth phase: The implementation

This probably is the less interesting part of the process, if you reach this point, you are already in the right direction.
Just keep in mind that the implementation must be done in a product way, and not in a project way.
Meaning that the implementation must be done in iterations, and the team must be able to deliver value in every iteration.

## Fifth phase: The migration

You will need to consider moving already existing projects to a new way of working. This is not an easy task, and it will take a lot of time.
You'll need to consider the following:
- The migration must be done in iterations, and the team must be able to deliver value in every iteration.
  - Example: By moving an internal legacy project to the IDP, the project will run smoothly, HA, and with a good UX for the developers, so next time you'll need to implement a fix, you'll be able to do it in a timely manner.
- The migration must be done in a way that the teams are not affected by the migration.

## Final words

As i mentioned above, this is just a high-level overview of the process, and it is not a step-by-step guide.
I've tried to explain the process in a way that can be understood by everyone, but I'm sure that there are a lot of things that I've missed.

I prefer to not mention names, but I would like to thank all people that have been helping me during this week.
The people i've been talking with are from different companies in different industries and countries, and they have been really helpful.

In the end, we are all in the same boat, and we all want to build better software, and we all want to build better software faster.

# Bonus: chat.openai.com

[Read what the chat @ openai considers about implementing an IDP here](chat.openai.com.md)

Im glad to see im not far from what an AI thinks about the topic.
