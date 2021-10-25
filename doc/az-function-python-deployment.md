# Azure Functions

<!-- TOC depthFrom:2 orderedList:true -->

## Overview

The function application to manage the Joiner abstraction. The programming language used in this function is python.

### New Joiners Service Publisher

#### Service Bus Trigger

This function is in charge of queueing a message based on a Natural Language processing result.

- NewJoinerReceiverFunction

## Natural Language Processing description

### Corpus

For efficiency and to allow a quick deployment it is used the small english corpus model [en_core_web_sm](https://spacy.io/models) from the Spacy trained pipeline. To gain more accuracy in the future could be used a larger corpus such as **en_core_web_trf**.