# Azure Functions

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The function application to manage the Task abstraction. The programming language used in these set of functions is nodejs.

### New Joiner Functions

#### HTTP Trigger

Mainly they are functions to support the CRUD operations over the Task abstraction.

- NewJoinerTaskFunction (GET/POST)
- NewJoinerTaskGetByFunction (GET)
- NewJoinerTaskPutDeleteFunction (PUT/DELETE)