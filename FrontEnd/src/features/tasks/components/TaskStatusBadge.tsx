import { TASK_STATUS_CONFIG, type TaskStatus } from "../types/type";

const TaskStatusBadge = ({ status }: { status: TaskStatus }) => {
  const { label, className } = TASK_STATUS_CONFIG[status];

  return (
    <span
      className={`shrink-0 rounded-full px-3 py-1 text-xs font-semibold ${className}`}
    >
      {label}
    </span>
  );
};

export default TaskStatusBadge;
