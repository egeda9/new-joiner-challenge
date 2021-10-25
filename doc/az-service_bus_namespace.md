# Azure Service Bus

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The function applications to manage the event Publish and Subscribe actions for Joiner and Task message handling.

### Service Bus Queues

#### Queue Trigger Functions

Mainly they are functions to support the actions to Publish and Subscribe to a message with the Azure Service Bus Queue as a broker service.

- joinerqueue message out

```json
{
      "type": "serviceBus",
      "direction": "out",
      "connection": "",
      "name": "msg",
      "queueName": "joinerqueue"
}
```

- joinerqueue message in

```json
{
    "name": "newJoinerSubscriberServiceBusTrigger",
    "type": "serviceBusTrigger",
    "direction": "in",
    "queueName": "joinerqueue",
    "connection": ""
}
```

- taskqueue message out

```json
{
      "type": "serviceBus",
      "direction": "out",
      "connection": "",
      "name": "msg",
      "queueName": "taskqueue"
}
```

- taskqueue message in

```json
{
    "name": "taskSubscriberServiceBusTrigger",
    "type": "serviceBusTrigger",
    "direction": "in",
    "queueName": "joinerqueue",
    "connection": ""
}
```