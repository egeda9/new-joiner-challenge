require('dotenv').config();

const db = require('../dboperations');
const sender = require('../sender');
const joinerservice = require('../joinerservice');
const TaskMessage = require('../taskMessage');

module.exports = async function (context, req) {
    context.log('JavaScript HTTP trigger function processed a request.');   

    var paramId = req.params.id    

    if (paramId > 0) {

        if (req.method == 'PUT') {

            try {            
                var currentTask = await db.get(paramId)
    
                if (currentTask.length > 0) {
                    if (req.body) {
                        
                        var joiner = ''

                        var taskMessage = new TaskMessage(
                            req.body.Id, 
                            req.body.Name,  
                            req.body.Description, 
                            req.body.EstimatedRequiredHours,
                            req.body.Stack,
                            req.body.MinRole
                            )

                        if (req.body.UserId) {
                            
                            joiner = await joinerservice.get(req.body.UserId);

                            if (!joiner) {
                                context.res = {
                                    headers: { 'Content-Type': 'application/json' },
                                    body: '{ "message": "Invalid user assignment" }',
                                    statusCode: 400
                                }
                            }

                            taskMessage.User = joiner
                        }

                        if (!req.body.UserId || joiner) {
                            if (req.body.TaskId && currentTask[0].TaskId != req.body.TaskId) {
                                hasParent = await db.getParent(req.body.TaskId)
        
                                if (hasParent.length == 0) {                         
                                    taskMessage.Task = hasParent
                                    
                                    await db.update(paramId, req.body)
                                    await sender.send(taskMessage)
        
                                    context.res = {   
                                        headers: { 'Content-Type': 'application/json' },     
                                        body: '{ "message": "Successfully updated" }',
                                        statusCode: 200
                                    };                                    
                                }
        
                                else {
                                    context.res = {
                                        headers: { 'Content-Type': 'application/json' },
                                        body: '{ "message": "Can not update parent task. Parent task has a child" }',
                                        statusCode: 400
                                    };
                                }
                            }
                            else {
                                await db.update(paramId, req.body)
                                await sender.send(taskMessage)                               
        
                                context.res = {   
                                    headers: { 'Content-Type': 'application/json' },     
                                    body: '{ "message": "Successfully updated" }',
                                    statusCode: 200
                                };
                            }   
                        }               
                    }
                    else {
                        context.res = {
                            headers: { 'Content-Type': 'application/json' },
                            body: '{ "message": "Invalid request body" }',
                            statusCode: 400
                        }
                    }               
                }
                else {
                    context.res = {
                        headers: { 'Content-Type': 'application/json' },
                        body: '{ "message": "No results" }',
                        statusCode: 204
                    }
                }
            } catch (error) {
                console.error(error);
                
                context.res = {
                    headers: { 'Content-Type': 'application/json' },
                    body: '{ "message": "' + error + '" }',
                    statusCode: 500
                }
            }              
        }
        
        else if (req.method == 'DELETE') {
            try {                
                await db.deleteTask(paramId)

                context.res = {
                    headers: { 'Content-Type': 'application/json' },
                    body: '{ "message": "Successfully deleted" }',
                    statusCode: 500
                }

            } catch (error) {
                console.error(error);

                context.res = {
                    headers: { 'Content-Type': 'application/json' },
                    body: '{ "message": "' + error + '" }',
                    statusCode: 500
                }
            }
        }
    
        else {
            context.res = {
                headers: { 'Content-Type': 'application/json' },
                body: '{ "message": "Not supported operation" }',
                statusCode: 404
            }
        }
    }
    
    else {
        context.res = {
            headers: { 'Content-Type': 'application/json' },
            body: '{ "message": "Invalid route parameter" }',
            statusCode: 404
        }
    }
}