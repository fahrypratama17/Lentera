import { Clock } from "lucide-react";

import ProgressStatusBadge from "./ProgressStatusBadge";
import { reviewDecisionConfig } from "../types/type";
import { formatDate } from "../utils/format";
import type { MaterialWithProgress } from "../types/type";

const MaterialItem = ({ material }: { material: MaterialWithProgress }) => {
  const { position, title, deadline, status, reviewDecision, reviewNote } =
    material;

  const isOverdue =
    deadline !== null &&
    new Date(deadline) < new Date() &&
    (status === "not_started" || status === "in_progress");

  return (
    <article className="flex flex-col gap-3 rounded-2xl border border-[#DCE6ED] bg-white p-5 shadow-sm">
      <div className="flex flex-wrap items-start gap-4">
        <div className="bg-secondary-light text-primary flex size-10 shrink-0 items-center justify-center rounded-xl text-sm font-bold">
          {position}
        </div>

        <div className="min-w-0 flex-1">
          <h3 className="text-primary-text leading-snug font-semibold">
            {title}
          </h3>

          {deadline && (
            <p
              className={`mt-1 flex items-center gap-1.5 text-xs ${
                isOverdue ? "text-red-600" : "text-secondary-text"
              }`}
            >
              <Clock className="size-3.5" />
              Tenggat {formatDate(deadline)}
              {isOverdue && " · Terlambat"}
            </p>
          )}
        </div>

        <ProgressStatusBadge status={status} />
      </div>

      {reviewDecision && (
        <div className="bg-surface rounded-xl p-4 text-sm">
          <p
            className={`font-semibold ${reviewDecisionConfig[reviewDecision].className}`}
          >
            Review mentor: {reviewDecisionConfig[reviewDecision].label}
          </p>
          {reviewNote && (
            <p className="text-primary-text mt-1 leading-relaxed">
              {reviewNote}
            </p>
          )}
        </div>
      )}
    </article>
  );
};

export default MaterialItem;