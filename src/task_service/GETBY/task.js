class  Task{
    constructor(Id, Name, Description, EstimatedRequiredHours, Stack, MinRole, Task){
      this.Id = Id;
      this.Name = Name;
      this.Description = Description;
      this.EstimatedRequiredHours = EstimatedRequiredHours;
      this.Stack = Stack;
      this.MinRole = MinRole;
      this.Task = Task;
    }
  }
  
  module.exports = Task;