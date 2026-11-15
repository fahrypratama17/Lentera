import {
  AlertTriangle,
  CheckCircle2,
  Clock,
  ListChecks,
  Send,
} from "lucide-react";

import type { TaskSummaryProps } from "../types/type";

const items = [
  { key: "notStarted", label: "Belum mulai", icon: ListChecks },
  { key: "inProgress", label: "Dikerjakan", icon: Clock },
  { key: "submitted", label: "Menunggu review", icon: Send },
  { key: "completed", label: "Selesai", icon: CheckCircle2 },
  { key: "overdue", label: "Terlambat", icon: AlertTriangle },
] as const;

const TaskSummary = ({ summary }: TaskSummaryProps) => {
  return (
    <div className="flex flex-col gap-4">
      <div className="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-5">
        {items.map(({ key, label, icon: Icon }) => (
          <div
            key={key}
            className="flex items-center gap-3 rounded-2xl border border-[#DCE6ED] bg-white p-4"
          >
            <div className="bg-secondary-light text-primary flex size-10 shrink-0 items-center justify-center rounded-xl">
              <Icon className="size-5" />
            </div>
            <div className="min-w-0">
              <p className="text-primary-text text-lg font-bold">
                {summary[key]}
              </p>
              <p className="text-secondary-text truncate text-xs">{label}</p>
            </div>
          </div>
        ))}
      </div>

      <div className="flex flex-col gap-2 rounded-2xl border border-[#DCE6ED] bg-white p-4">
        <div className="flex items-baseline justify-between gap-3">
          <p className="text-secondary-text text-xs">Penyelesaian tugas</p>
          <p className="text-primary-text text-sm font-bold">
            {summary.completed}/{summary.total} · {summary.completionRate}%
          </p>
        </div>

        <div
          role="progressbar"
          aria-valuenow={summary.completionRate}
          aria-valuemin={0}
          aria-valuemax={100}
          aria-label="Penyelesaian tugas"
          className="bg-secondary-light h-2 overflow-hidden rounded-full"
        >
          <div
            className="bg-primary h-full rounded-full transition-all"
            style={{ width: `${summary.completionRate}%` }}
          />
        </div>
      </div>
    </div>
  );
};

export default TaskSummary;
