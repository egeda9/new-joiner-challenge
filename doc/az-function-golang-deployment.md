# Azure Functions

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The function application to manage the Joiner abstraction. The programming language used in these set of functions is golang.

### New Joiner Functions

#### HTTP Trigger

Mainly they are functions to support the CRUD operations over the Joiner abstraction.

- NewJoinerByFunction (GET)
- NewJoinerFunction (GET)
- NewJoinerUpdateFunction (PUT)