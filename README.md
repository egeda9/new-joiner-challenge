# New Joiner Challenge

<!-- TOC depthFrom:2 orderedList:true -->
- [New Joiner Challenge](#new-joiner-challenge)
  - [Project Description](#project-description)
  - [Preconditions](#preconditions)
  - [Architecture Design](#architecture-design)
    - [Technology Stack](#technology-stack)
    - [Components Description](#components-description)
    - [Solution Flow](#solution-flow)
      - [Upload Joiner Profile](#upload-joiner-profile)
      - [Create New Joiner](#create-new-joiner)
      - [Get List of Joiners](#get-list-of-joiners)
      - [Update Joiner](#update-joiner)
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

**Endpoint**: Action POST

```bash
{baseUrl}/api/Joiner
```

**form-data**:

```bash
--form 'file=@"../../Profile_Test.docx"'
```

**Response body (Sample)**:
**Status Code**: OK

```bash
OK
```

- Azure Function 2: Joiner Microservice
  - Stack: Golang

**Endpoint**: Action GET

```bash
{baseUrl}/api/Joiner
```
**Response body (Sample)**:
**Status Code**: OK

```json
[
    {
        "Id": 1,
        "Name": "Name 1",
        "Stack": "[\"Python\", \"Java\", \"Scala\",\"Akka\",\"ScalaTest\",\"BlueMix\",\"MqLite\",\"Cloudant\",\"Windows\",\"Spring\",\"Hibernate\",\"Eclipse\",\"Maven\",\"Tomcat\"]",
        "Role": "Developer",
        "Languages": "[\"Romanian\",\"Russian\",\"English\",\"French\"]"
    },
    {
        "Id": 2,
        "Name": "Name 2",
        "Stack": "[\"Azure SQL Azure\",\"Azure Cosmos DB\",\"Azure Virtual Machines\",\"Bizagi Cloud\",\"BPM Bizagi\"]",
        "Role": "Senior Developer",
        "Languages": "[\"Spanish\",\"English\"]"
    },
    {
        "Id": 3,
        "Name": "Name 3",
        "Stack": "[\"Azure Virtual Machines\",\"BPM Bizagi\"]",
        "Role": "Developer",
        "Languages": "[\"Spanish\",\"English\"]"
    },
    {
        "Id": 4,
        "Name": "Name 4",
        "Stack": "[\"Azure Virtual Machines\",\"Bizagi Cloud\",\"BPM Bizagi\"]",
        "Role": "UX Developer",
        "Languages": "[\"Spanish\",\"English\"]"
    }
]
```

**Endpoint**: Action GET

```bash
{baseUrl}/api/Joiner?id=1
```

**Response body (Sample)**:
**Status Code**: OK

```json
{
        "Id": 1,
        "Name": "Name 1",
        "Stack": "[\"Python\", \"Java\", \"Scala\",\"Akka\",\"ScalaTest\",\"BlueMix\",\"MqLite\",\"Cloudant\",\"Windows\",\"Spring\",\"Hibernate\",\"Eclipse\",\"Maven\",\"Tomcat\"]",
        "Role": "Developer",
        "Languages": "[\"Romanian\",\"Russian\",\"English\",\"French\"]"
}
```

**Endpoint**: Action PUT

```bash
{baseUrl}/api/Joiner?id=1
```

**Request body (Example)**:

```json
{	
	"name": "Name 1",
	"stack": "[\"Python\", \"Scala\",\"Akka\",\"ScalaTest\",\"BlueMix\",\"MqLite\",\"Cloudant\",\"Windows\",\"Spring\",\"Hibernate\",\"Eclipse\",\"Maven\",\"Tomcat\"]",
	"role": "Developer",
	"languages": "[\"Romanian\",\"Russian\",\"English\",\"French\"]"
}
```

**Response body (Sample)**:
**Status Code**: OK

```json
{
  "Id": 1,
  "Name": "Name 1",
  "Stack": "[\"Python\", \"Scala\",\"Akka\",\"ScalaTest\",\"BlueMix\",\"MqLite\",\"Cloudant\",\"Windows\",\"Spring\",\"Hibernate\",\"Eclipse\",\"Maven\",\"Tomcat\"]",
  "Role": "Developer",
  "Languages": "[\"Romanian\",\"Russian\",\"English\",\"French\"]"
}
```

- Azure Function 3: Task Microservice
  - Stack: NodeJS

**Endpoint**: Action GET

```bash
{baseUrl}/api/Task
```
**Response body (Sample)**:
**Status Code**: OK

```json
[
    {
        "Id": 1,
        "Name": "Analyse Requirement",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "EstimatedRequiredHours": 1,
        "Stack": "--",
        "MinRole": "BA",
        "Task": {
            "Id": 2,
            "Name": "Create Code Component",
            "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
            "EstimatedRequiredHours": 3,
            "Stack": ".Net",
            "MinRole": "Developer"
        }
    },
    {
        "Id": 7,
        "Name": "Run Tests",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "EstimatedRequiredHours": 1,
        "Stack": "--",
        "MinRole": "BA",
        "Task": {
            "Id": 8,
            "Name": "Create Code Component",
            "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
            "EstimatedRequiredHours": 2,
            "Stack": ".Net",
            "MinRole": "Developer"
        }
    },
    {
        "Id": 3,
        "Name": "Run Tests",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "EstimatedRequiredHours": 1,
        "Stack": "--",
        "MinRole": "BA",
        "Task": null
    }
]
```

**Endpoint**: Action GET

```bash
{baseUrl}/api/Task/1
```
**Response body (Sample)**:
**Status Code**: OK

```json
{
    "Id": 1,
    "Name": "Analyse Requirement",
    "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
    "EstimatedRequiredHours": 1,
    "Stack": "--",
    "MinRole": "BA",
    "Task": {
        "Id": 2,
        "Name": "Create Code Component",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "EstimatedRequiredHours": 3,
        "Stack": ".Net",
        "MinRole": "Developer"
    }
}
```

**Endpoint**: Action POST

```bash
{baseUrl}/api/Task
```

**Request body (Example)**:

```json
{        
        "Name": "Run Tests",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "EstimatedRequiredHours": 1,
        "Stack": "--",
        "MinRole": "BA",
        "Task": {            
            "Name": "Create Code Component",
            "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
            "EstimatedRequiredHours": 2,
            "Stack": ".Net",
            "MinRole": "Developer"
        }
    }
```

**Response Status Code**: Created

**Endpoint**: Action PUT

```bash
{baseUrl}/api/Task/1
```

**Request body (Example)**:

```json
{        
    "Name": "Create Code Component",
    "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
    "EstimatedRequiredHours": 3,
    "Stack": ".Net",
    "MinRole": "Developer",
    "TaskId": 1,
    "UserId": 1
}
```

*Optional: TaskId and UserId

**Response Status Code**: OK

```json
{
    "message": "Successfully updated"
}
```

**Endpoint**: Action DELETE

```bash
{baseUrl}/api/Task/1
```

**Response Status Code**: OK

```json
{
    "message": "Successfully deleted"
}
```

### Solution Flow

#### Upload Joiner Profile

![Upload Joiner Profile](./doc/img/new_joiners_challenge_upload_joiner_profile.png)

#### Create New Joiner

![Create New Joiner](./doc/img/new_joiners_challenge_create_new_joiner.png)

#### Get List of Joiners

![Get List of Joiners](./doc/img/new_joiners_challenge_get_list_of_joiners.png)

#### Update Joiner

![Update Joiner](./doc/img/new_joiners_challenge_update_joiner.png)

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
