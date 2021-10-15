const sql = require('mssql');

const config = {
    port: Number(process.env.DB_PORT),
    server: process.env.DB_SERVER,
    user: process.env.DB_USER,
    password: process.env.DB_PWD,
    database: process.env.DB_NAME,
    options: {
      trustServerCertificate: true    
    }
  }

async function getTasks () {
    console.log("---------getTasks");

    sql.on('error', err => {
        console.log("DB Error2: " + err); 
    })

    var query = "SELECT t2.Id AS ParentId, " +
                "t2.[Name] AS Parent, " +
                "t2.[Description] AS ParentDescription, " +
                "t2.EstimatedRequiredHours AS ParentEstimatedRequiredHours, " +
                "t2.Stack AS ParentStack, " +
                "t2.MinRole AS ParentMinRole, " +
                "t1.Id AS ChildId, " +
                "t1.[Name] AS Child, " +
                "t1.[Description] AS ChildDescription, " +
                "t1.EstimatedRequiredHours AS ChildEstimatedRequiredHours, " +
                "t1.Stack AS ChildStack, " +
                "t1.MinRole AS ChildMinRole " +
                "FROM Task t1 " +
                "INNER JOIN Task t2 ON t1.Id = t2.TaskId ;"

    try {
    const pool = await sql.connect(config);
    const result_1 = await pool.request()
      .query(query);
    console.log("getTasks:then(result=>");
    sql.close();
    return result_1.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

async function getTasksParentWithoutChild () {
  console.log("---------getTasksParentWithoutChild");

  sql.on('error', err => {
      console.log("DB Error2: " + err); 
  })

  var query = "SELECT * " +
              "FROM Task t1 " +
              "WHERE t1.TaskId IS NULL " +
              "AND NOT EXISTS ( " +
              "    SELECT 1  " +
              "  FROM Task t2	" +
              "    WHERE t1.Id = t2.TaskId " +
              ");"

  try {
    const pool = await sql.connect(config);
    const result_1 = await pool.request()
      .query(query);
    console.log("getTasksParentWithoutChild:then(result=>");
    sql.close();
    return result_1.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

async function create (task) {
  console.log("---------create");

  sql.on('error', err => {
      console.log("DB Error2: " + err); 
  })

  var query = "INSERT INTO Task (Name, " +
                                "Description, " + 
                                "EstimatedRequiredHours, " + 
                                "Stack, " + 
                                "MinRole, " +
                                "TaskId) " +
                                "VALUES ( " + 
                                "@Name, " + 
                                "@Description, " +
                                "@EstimatedRequiredHours, " +
                                "@Stack, " +
                                "@MinRole, " + 
                                "@TaskId); " + 
                                "SELECT SCOPE_IDENTITY() AS id;"

  try {

    const pool = await sql.connect(config);
    const result_2 = await pool.request()
      .input('Name', sql.VarChar, task.Name)
      .input('Description', sql.VarChar, task.Description)
      .input('EstimatedRequiredHours', sql.Int, task.EstimatedRequiredHours)
      .input('Stack', sql.VarChar, task.Stack)
      .input('MinRole', sql.VarChar, task.MinRole)
      .input('TaskId', sql.Int, task.TaskId)
      .query(query);
    console.log("create:then(result=>");
    sql.close();
    //console.log(result_2);
    return result_2.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

module.exports = {
  getTasks:  getTasks,
  getTasksParentWithoutChild: getTasksParentWithoutChild,
  create: create
}