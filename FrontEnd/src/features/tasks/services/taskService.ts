import { cache } from "react";

import { mockCohorts } from "@/features/dashboard/data/data";
import { taskSeeds } from "../data/data";
import { sortTasks } from "../utils/task";
import { delay, isoFromNow } from "./shared";
import type { TaskItem } from "../types/type";


const buildTasks = (): TaskItem[] =>
  taskSeeds.flatMap((seed) => {
    const cohort = mockCohorts.find((item) => item.id === seed.cohortId);
    if (!cohort) return [];

    const hasSubmission =
      seed.status === "submitted" || seed.status === "completed";

    return [
      {
        id: `${seed.cohortId}-p${seed.materialIndex}`,
        cohortId: cohort.id,
        cohortTitle: cohort.title,
        cohortTone: cohort.tone,
        materialId: `${seed.cohortId}-${seed.materialIndex}`,
        materialTitle: seed.materialTitle,
        position: seed.position,
        deadline:
          seed.deadlineOffset === null ? null : isoFromNow(seed.deadlineOffset),
        status: seed.status,
        submissionContent: seed.submissionContent,
        submissionUrl: seed.submissionUrl,
        submittedAt: hasSubmission
          ? isoFromNow(
              seed.deadlineOffset === null ? -1 : seed.deadlineOffset - 1,
            )
          : null,
        reviewDecision: seed.reviewDecision,
        reviewNote: seed.reviewNote,
        reviewerName: seed.reviewerName,
        reviewedAt: seed.reviewDecision ? isoFromNow(-1) : null,
        startedAt: seed.status === "not_started" ? null : isoFromNow(-14),
        completedAt: seed.status === "completed" ? isoFromNow(-3) : null,
        role: cohort.role,
      },
    ];
  });

export const getMyTasks = cache(async (): Promise<TaskItem[]> => {
  await delay();
  return sortTasks(buildTasks());
});
