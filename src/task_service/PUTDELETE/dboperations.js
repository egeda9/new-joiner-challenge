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

async function update (id, task) {
    console.log("---------update");

    sql.on('error', err => {
        console.log("DB Error2: " + err); 
    })

    var query = "UPDATE Task " +
                "SET [Name] = @Name, " +
                "[Description] = @Description," +
                "EstimatedRequiredHours = @EstimatedRequiredHours, " +
                "Stack = @Stack, " +
                "MinRole = @MinRole, " +
                "TaskId = @TaskId, " +
                "UserId = @UserId " +
                "WHERE Id = @Id;"

    try {
    const pool = await sql.connect(config);
    const result_1 = await pool.request()
      .input('Id', sql.Int, id)
      .input('Name', sql.VarChar, task.Name)
      .input('Description', sql.VarChar, task.Description)
      .input('EstimatedRequiredHours', sql.Int, task.EstimatedRequiredHours)
      .input('Stack', sql.VarChar, task.Stack)
      .input('MinRole', sql.VarChar, task.MinRole)
      .input('TaskId', sql.Int, task.TaskId)
      .input('UserId', sql.VarChar, task.UserId)
      .query(query);
    console.log("updateTask:then(result=>");
    sql.close();
    return result_1.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

async function get (id) {
  console.log("---------getTask");

  sql.on('error', err => {
      console.log("DB Error2: " + err); 
  })

  var query = "SELECT * " +
              "FROM Task t1 " +
              "WHERE t1.Id = @Id;"

  try {
    const pool = await sql.connect(config);
    const result_1 = await pool.request()
      .input('Id', sql.Int, id)
      .query(query);
    console.log("getTask:then(result=>");
    sql.close();
    return result_1.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

async function deleteTask (id) {
  console.log("---------deleteTask");

  sql.on('error', err => {
      console.log("DB Error2: " + err); 
  })

  var query = "DELETE " +
              "FROM Task " +
              "WHERE Id = @Id;"

  try {
    const pool = await sql.connect(config);
    const result_1 = await pool.request()
      .input('Id', sql.Int, id)
      .query(query);
    console.log("deleteTask:then(result=>");
    sql.close();
    return result_1.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

async function getParent (taskId) {
  console.log("---------getTask");

  sql.on('error', err => {
      console.log("DB Error2: " + err); 
  })

  var query = "SELECT * " +
              "FROM Task t1 " +
              "WHERE t1.TaskId = @TaskId;"

  try {
    const pool = await sql.connect(config);
    const result_1 = await pool.request()
      .input('TaskId', sql.Int, taskId)
      .query(query);
    console.log("getParent:then(result=>");
    sql.close();
    return result_1.recordset;
  } catch (err_1) {
    console.log("DB Error1: " + err_1);
    sql.close();
    throw err_1;
  }
}

module.exports = {
  update:  update,
  get: get,
  getParent: getParent,
  deleteTask: deleteTask
}