# Reporting Component

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The function application to manage the reporting information and how this data is decoupled from the main transactional process.

### New Joiners Service Reporting

#### HTTP Trigger (TODO)

Mainly they are functions to support the report generation based on 4 types of reports in CSV format.

- Top 'N' Joiners by 'X' stack (order by amount of task completed).
- Tasks completed and pending by joiner.
- Task completed an uncompleted by 'X' joiner.
- Days left to complete the tasks by joiner (only for joiners with pending tasks, assume 8 hours 1 day).

#### Model stored in Cosmos DB

```json
{        
        "Id": 1,
        "Name": "Run Tests",
        "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "EstimatedRequiredHours": 1,
        "Stack": "--",
        "MinRole": "BA",
        "Task": {    
            "Id": 2,        
            "Name": "Create Code Component",
            "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
            "EstimatedRequiredHours": 2,
            "Stack": ".Net",
            "MinRole": "Developer"
        },
        "User": {
            "Id": 1,
            "Name": "Name 1",
            "Stack": "[\"Python\", \"Scala\",\"Akka\",\"ScalaTest\",\"BlueMix\",\"MqLite\",\"Cloudant\",\"Windows\",\"Spring\",\"Hibernate\",\"Eclipse\",\"Maven\",\"Tomcat\"]",
            "Role": "Developer",
            "Languages": "[\"Romanian\",\"Russian\",\"English\",\"French\"]"
       }
}
```
