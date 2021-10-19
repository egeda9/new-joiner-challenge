# New Joiner Challenge

<!-- TOC depthFrom:2 orderedList:true -->
- [New Joiner Challenge](#new-joiner-challenge)
  - [Project Description](#project-description)
  - [Preconditions](#preconditions)
  - [Architecture Design](#architecture-design)
    - [Technology Stack](#technology-stack)
    - [Components Description](#components-description)
    - [Solution Flow](#solution-flow)
    - [Data Model](#data-model)
  - [Reporting Service Description](#reporting-service-description)

## Project Description

This project aims to create, register and manage some of the tasks that a new company joiner should work on.

## Preconditions

- Input file only in PDF or DOCX format.
- Use an Event-Driven Architecture approach.
- Build an approach based on four services with a specific responsibility.
- Each service should be created with a different programming language.

## Architecture Design

![Architecture Design](./doc/img/new_joiners_challenge-architecture-design.png)

### Technology Stack

- [Python](https://www.python.org)
  - [Spacy](https://spacy.io)
- [Golang](https://golang.org)
- [NodeJS](https://nodejs.org/en/)
- [Microsoft SQL](https://www.microsoft.com/en-us/sql-server/)
- [RabbitMQ](https://www.rabbitmq.com)
- [Microsoft Azure](https://azure.microsoft.com/en-us/)
  - [Azure Functions](https://azure.microsoft.com/en-us/services/functions/)
  - [Azure Service Bus](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-messaging-overview)
  - [Azure SQL](https://azure.microsoft.com/en-us/products/azure-sql/database/)
  - [Azure Data Factory](https://docs.microsoft.com/en-us/azure/data-factory/).

### Components Description

- Azure Function 1: File Upload and NLP Microservice
  - Stack: Python
- Azure Function 2: Joiner Microservice
  - Stack: Golang
- Azure Function 3: Task Microservice
  - Stack: NodeJS

### Solution Flow

### Data Model

![Data Model](./doc/img/data_model.png)

**Joiner**: `Joiner`

| Column Name | Type | Sample |  
|-----------|:-----------:|:-----------|  
| Id | int | 1 |
| Name | varchar | Camilo Robles |
| Stack | varchar | ["Python", "Java", "Scala"] |  
| Role | varchar | Developer |
| Languages | varchar | ["Spanish","English"] |
| JoinerMessageAcknowledgementId | int | 300 |

**Joiner Message Acknowledgement**: `JoinerMessageAcknowledgement`

| Column Name | Type | Sample |  
|-----------|:-----------:|:-----------|  
| Id | int | 11 |
| CreatedDate | varchar | 2021-10-11 19:26:12.000 |
| Status | varchar | Complete |  
| Message | varchar | {"DATE": ["2011-2014", "February 2015 |

**Task**: `Task`

| Column Name | Type | Sample |  
|-----------|:-----------:|:-----------|  
| Id | int | 1 |
| Name | varchar | Create Code Component |
| Description | varchar | Lorem ipsum dolor sit |  
| EstimatedRequiredHours | varchar | 2 |
| Stack | varchar | .Net |
| MinRole | varchar | Developer |
| TaskId | int | 2 |
| UserId | int | 1 |

## Reporting Service Description
