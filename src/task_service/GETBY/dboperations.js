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

async function getTask (id) {
    console.log("---------getTasks");

    sql.on('error', err => {
        console.log("DB Error2: " + err); 
    })

    var query = "SELECT t1.Id, " +
                "t1.[Name], " +
                "t1.[Description], " +
                "t1.EstimatedRequiredHours, " +
                "t1.Stack, " +
                "t1.MinRole, " +
                "t2.Id AS ChildId, " +
                "t2.[Name] AS ChildName, " +
                "t2.[Description] AS ChildDescription, " +
                "t2.EstimatedRequiredHours AS ChildEstimatedRequiredHours, " +
                "t2.Stack AS ChildStackName, " +
                "t2.MinRole AS ChildMinRoleName " +
                "FROM Task t1 " +
                "LEFT JOIN Task t2 ON t1.Id = t2.TaskId " +
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

module.exports = {
    getTask:  getTask
}