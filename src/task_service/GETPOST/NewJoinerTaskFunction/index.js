require('dotenv').config();

const db = require('../dboperations');
const sender = require('../sender');
const Task = require('../task');
const Child = require('../child');
const TaskMessage = require('../taskMessage');

module.exports = async function (context, req) {
    context.log('JavaScript HTTP trigger function processed a request.');

    if (req.method == 'POST') {

        if (req.body) {            
            //console.log(req.body)
            let lastInsertedRecord = null

            if (req.body) {            
                lastInsertedRecord = await db.create(req.body)
            }
            
            if (req.body.Task && lastInsertedRecord && lastInsertedRecord.length > 0) {
                req.body.Task.TaskId = lastInsertedRecord[0].id            
                lastInsertedRecord = await db.create(req.body.Task)
            }            

            context.res = {
                headers: { 'Content-Type': 'application/json' },
                statusCode: 201
            }
        }
        else {
            context.res = {
                headers: { 'Content-Type': 'application/json' },
                body: '{ "message": "Empty request body" }',
                statusCode: 400
            }
        }
    }

    else if (req.method == 'GET') {
        try {
            const tasks = await db.getTasks()
            const tasksParentWithoutChild = await db.getTasksParentWithoutChild()
            const result = []
        
            for (var i = 0; i < tasks.length; i++) {
                var childInstance = new Child(
                    tasks[i].ChildId, 
                    tasks[i].Child, 
                    tasks[i].ChildDescription, 
                    tasks[i].ChildEstimatedRequiredHours,
                    tasks[i].ChildStack,
                    tasks[i].ChildMinRole)
        
                var parent = new Task(
                    tasks[i].ParentId, 
                    tasks[i].Parent, 
                    tasks[i].ParentDescription, 
                    tasks[i].ParentEstimatedRequiredHours,
                    tasks[i].ParentStack,
                    tasks[i].ParentMinRole,
                    childInstance)
        
                result.push(parent)
            }
        
            for (var i = 0; i < tasksParentWithoutChild.length; i++) {
                var parent = new Task(
                    tasksParentWithoutChild[i].Id, 
                    tasksParentWithoutChild[i].Name, 
                    tasksParentWithoutChild[i].Description, 
                    tasksParentWithoutChild[i].EstimatedRequiredHours,
                    tasksParentWithoutChild[i].Stack,
                    tasksParentWithoutChild[i].MinRole,
                    null)
        
                result.push(parent)
            }
        
            const statusCode = 200
        
            if (!result.length === 0) {
                statusCode = 204
                result = null
            }
        
            context.res = {
                headers: { 'Content-Type': 'application/json' },
                body: result,
                statusCode: statusCode
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
            body: '{ "message": "Invalid Operation" }',
            statusCode: 404
        }
    }    
}