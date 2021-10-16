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

### Components Description

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
| Task | int | 2 |

**Joiner Task**: `JoinerTask`

| Column Name | Type | Sample |  
|-----------|:-----------:|:-----------|  
| Id | uniqueidentifier | 6BECB7B8-D876-4CFF-A25E-3E66F83DE873 |
| Task | varchar |  |
| Joiner | varchar |  |  

## Reporting Service Description
