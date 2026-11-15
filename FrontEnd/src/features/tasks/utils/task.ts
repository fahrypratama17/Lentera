import type { TaskFilter, TaskItem, TaskSummary } from "../types/type";

const DAY = 86_400_000;

export const isOverdue = (task: TaskItem, now: number = Date.now()) =>
  task.deadline !== null &&
  (task.status === "not_started" || task.status === "in_progress") &&
  new Date(task.deadline).getTime() < now;

export const matchesFilter = (
  task: TaskItem,
  filter: TaskFilter,
  now: number = Date.now(),
) => {
  if (filter === "all") return true;
  if (filter === "overdue") return isOverdue(task, now);

  return task.status === filter;
};

export const matchesSearch = (task: TaskItem, keyword: string) => {
  const query = keyword.trim().toLowerCase();
  if (!query) return true;

  return [task.materialTitle, task.cohortTitle]
    .join(" ")
    .toLowerCase()
    .includes(query);
};

export const summarize = (
  tasks: TaskItem[],
  now: number = Date.now(),
): TaskSummary => {
  const total = tasks.length;
  const completed = tasks.filter((task) => task.status === "completed").length;

  return {
    total,
    notStarted: tasks.filter((task) => task.status === "not_started").length,
    inProgress: tasks.filter((task) => task.status === "in_progress").length,
    submitted: tasks.filter((task) => task.status === "submitted").length,
    completed,
    overdue: tasks.filter((task) => isOverdue(task, now)).length,
    completionRate: total === 0 ? 0 : Math.round((completed / total) * 100),
  };
};

export const sortTasks = (tasks: TaskItem[]) =>
  [...tasks].sort((a, b) => {
    if (a.deadline === null && b.deadline === null) {
      return a.position - b.position;
    }
    if (a.deadline === null) return 1;
    if (b.deadline === null) return -1;

    const diff =
      new Date(a.deadline).getTime() - new Date(b.deadline).getTime();

    return diff !== 0 ? diff : a.position - b.position;
  });

export const remainingDays = (
  deadline: string | null,
  now: number = Date.now(),
) => {
  if (deadline === null) return null;
  return Math.ceil((new Date(deadline).getTime() - now) / DAY);
};
