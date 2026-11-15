import {
  ArrowUpRight,
  BookOpen,
  CalendarDays,
  Link2,
  UserRound,
} from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import { toneClasses } from "@/features/dashboard/types/type";
import { formatDueLabel } from "../utils/format";
import { isOverdue } from "../utils/task";
import TaskStatusBadge from "./TaskStatusBadge";
import { REVIEW_DECISION_CONFIG, type TaskCardProps } from "../types/type";

const TaskCard = ({ task, onOpen }: TaskCardProps) => {
  const overdue = isOverdue(task);
  const canSubmit =
    task.status === "not_started" || task.status === "in_progress";

  return (
    <article className="flex flex-col gap-4 rounded-2xl border border-[#DCE6ED] bg-white p-5 shadow-sm transition-all duration-300 hover:-translate-y-0.5 hover:shadow-[0_16px_40px_-18px_rgba(31,80,109,0.3)]">
      <div className="flex items-start gap-3">
        <div className="bg-secondary-light text-primary flex size-10 shrink-0 items-center justify-center rounded-xl text-sm font-bold">
          {task.position}
        </div>

        <div className="min-w-0 flex-1">
          <span
            className={`inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-[11px] font-semibold text-white ${toneClasses[task.cohortTone]}`}
          >
            <BookOpen className="size-3" />
            <span className="truncate">{task.cohortTitle}</span>
          </span>

          <h3 className="text-primary-text mt-2 leading-snug font-semibold">
            {task.materialTitle}
          </h3>

          <p
            className={`mt-1 flex items-center gap-1.5 text-xs ${
              overdue ? "text-red-600" : "text-secondary-text"
            }`}
          >
            <CalendarDays className="size-3.5" />
            {formatDueLabel(task.deadline)}
          </p>
        </div>

        <TaskStatusBadge status={task.status} />
      </div>

      {task.reviewDecision && (
        <div className="bg-surface rounded-xl p-4 text-sm">
          <p
            className={`font-semibold ${
              REVIEW_DECISION_CONFIG[task.reviewDecision].className
            }`}
          >
            Review mentor: {REVIEW_DECISION_CONFIG[task.reviewDecision].label}
          </p>

          {task.reviewNote && (
            <p className="text-primary-text mt-1 leading-relaxed">
              {task.reviewNote}
            </p>
          )}

          {task.reviewerName && (
            <p className="text-secondary-text mt-2 flex items-center gap-1.5 text-xs">
              <UserRound className="size-3.5" />
              {task.reviewerName}
            </p>
          )}
        </div>
      )}

      {task.submissionUrl && !task.reviewDecision && (
        <p className="text-secondary-text flex items-center gap-1.5 text-xs">
          <Link2 className="size-3.5 shrink-0" />
          <span className="truncate">{task.submissionUrl}</span>
        </p>
      )}

      <div className="flex items-center justify-end border-t border-[#DCE6ED] pt-4">
        <Button
          type="button"
          variant="ghost"
          size="sm"
          onClick={() => onOpen(task)}
          className="text-primary hover:bg-secondary-light hover:text-secondary cursor-pointer rounded-lg font-bold"
        >
          {canSubmit ? "Kerjakan" : "Lihat detail"}
          <ArrowUpRight className="size-4" />
        </Button>
      </div>
    </article>
  );
};

export default TaskCard;
