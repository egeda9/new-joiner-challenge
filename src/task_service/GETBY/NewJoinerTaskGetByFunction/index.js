require('dotenv').config();

const db = require('../dboperations');
const Task = require('../task');
const Child = require('../child');

module.exports = async function (context, req) {
    context.log('JavaScript HTTP trigger function processed a request.');
    
    var paramId = req.params.id

    try {
        let task = await db.getTask(paramId)
    
        if (task.length > 0) {
            
            var childInstance = null;

            if (task[0].ChildId) {                

                childInstance = new Child(
                    task[0].ChildId, 
                    task[0].ChildName, 
                    task[0].ChildDescription, 
                    task[0].ChildEstimatedRequiredHours,
                    task[0].ChildStackName,
                    task[0].ChildMinRoleName)
            }    
            
            var parent = new Task(
                task[0].Id, 
                task[0].Name, 
                task[0].Description, 
                task[0].EstimatedRequiredHours,
                task[0].Stack,
                task[0].MinRole,
                childInstance)
    
            context.res = {
                headers: { 'Content-Type': 'application/json' },
                body: parent,
                statusCode: 200
            };
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