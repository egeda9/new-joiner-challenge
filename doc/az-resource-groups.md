# Resource Groups

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The logical collection of resources, divided into groups to define application specific responsibilities.

```
└───Ncj.Joiner.Common
│   │   SQL database (Joiner)
│   │   SQL server
|   |   SQL database (Task)
│   │   Service Bus Namespace
│   
└───Ncj.Joiner.FileReceiver
│   │   App Service plan
│   │   Function App
|   |   Application Insights
│   │   Storage account
│   
└───Ncj.Joiner.Subscriber
|   │   App Service plan (4)
|   │   Application Insights (4)
|   |   Function App (4)
|   |   Storage account (4)
│ 
└───Ncj.Task
|   │   App Service plan (3)
|   │   Application Insights (3)
|   |   Function App (3)
|   |   Storage account (3)
│ 
```
