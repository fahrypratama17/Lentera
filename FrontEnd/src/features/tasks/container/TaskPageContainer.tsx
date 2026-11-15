import TaskContainer from "./TaskContainer";
import { getMyTasks } from "../services/taskService";

const TaskPageContainer = async () => {
  const tasks = await getMyTasks();

  return <TaskContainer initialTasks={tasks} />;
};

export default TaskPageContainer;
