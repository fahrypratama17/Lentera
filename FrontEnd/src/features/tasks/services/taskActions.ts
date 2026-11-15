import { assertCanStart, assertCanSubmit, delay } from "./shared";
import type { SubmitTaskValues, TaskItem } from "../types/type";


export const startTask = async (task: TaskItem): Promise<TaskItem> => {
  assertCanStart(task);
  await delay(250);

  return {
    ...task,
    status: "in_progress",
    startedAt: task.startedAt ?? new Date().toISOString(),
  };
};

export const submitTask = async (
  task: TaskItem,
  values: SubmitTaskValues,
): Promise<TaskItem> => {
  assertCanSubmit(task);

  const content = values.content.trim();
  const url = values.url.trim();

  if (!content && !url) {
    throw new Error("Isi hasil belajar atau tautan pekerjaan terlebih dahulu.");
  }

  await delay(400);

  return {
    ...task,
    status: "submitted",
    submissionContent: content || null,
    submissionUrl: url || null,
    submittedAt: new Date().toISOString(),
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
    reviewedAt: null,
    startedAt: task.startedAt ?? new Date().toISOString(),
  };
};
