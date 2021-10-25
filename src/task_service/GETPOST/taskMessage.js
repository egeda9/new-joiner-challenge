class  TaskMessage{
    constructor(Id, Name, Description, EstimatedRequiredHours, Stack, MinRole, Task, User){
      this.Id = Id;
      this.Name = Name;
      this.Description = Description;
      this.EstimatedRequiredHours = EstimatedRequiredHours;
      this.Stack = Stack;
      this.MinRole = MinRole;
      this.Task = Task;
      this.User = User;
      this.LastModified = new Date();
    }
  }
  
  module.exports = TaskMessage;