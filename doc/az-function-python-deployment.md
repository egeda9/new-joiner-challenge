# Azure Functions

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The function application to manage the Joiner abstraction. The programming language used in this function is python.

### New Joiners Service Publisher

#### Service Bus Trigger

This function is in charge of queueing a message based on a Natural Language processing result.

- NewJoinerReceiverFunction