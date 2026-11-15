import TaskCard from "./TaskCard";
import type { TaskListProps } from "../types/type";

const TaskList = ({ tasks, onOpen }: TaskListProps) => {
  return (
    <div className="grid grid-cols-1 gap-4 xl:grid-cols-2">
      {tasks.map((task) => (
        <TaskCard key={task.id} task={task} onOpen={onOpen} />
      ))}
    </div>
  );
};

export default TaskList;
