import { Search } from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import { TASK_FILTERS, type TaskToolbarProps } from "../types/type";

const TaskToolbar = ({
  filter,
  cohortId,
  search,
  cohortOptions,
  onFilterChange,
  onCohortChange,
  onSearchChange,
}: TaskToolbarProps) => {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex flex-col gap-3 xl:flex-row xl:items-center xl:justify-between">
        <div
          role="group"
          aria-label="Filter tugas berdasarkan status"
          className="flex w-fit max-w-full flex-wrap items-center gap-1 rounded-xl border border-[#DCE6ED] bg-white p-1"
        >
          {TASK_FILTERS.map((item) => {
            const active = filter === item.value;

            return (
              <Button
                key={item.value}
                type="button"
                variant="ghost"
                size="sm"
                aria-pressed={active}
                onClick={() => onFilterChange(item.value)}
                className={`h-9 cursor-pointer rounded-lg px-3 text-sm font-semibold transition-colors ${
                  active
                    ? "bg-primary hover:bg-secondary text-white hover:text-white"
                    : "text-secondary-text hover:bg-secondary-light hover:text-primary"
                }`}
              >
                {item.label}
              </Button>
            );
          })}
        </div>

        <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
          <label className="sr-only" htmlFor="task-cohort-filter">
            Filter kelas
          </label>

          <select
            id="task-cohort-filter"
            value={cohortId}
            onChange={(event) => onCohortChange(event.target.value)}
            className="text-primary-text focus:border-secondary focus:ring-secondary/15 h-11 w-full rounded-xl border border-[#DCE6ED] bg-white px-4 text-sm outline-none transition focus:ring-4 sm:w-auto sm:min-w-48"
          >
            {cohortOptions.map((option) => (
              <option key={option.value} value={option.value}>
                {option.label}
              </option>
            ))}
          </select>

          <div className="relative w-full sm:max-w-77.5">
            <Search className="text-secondary-text pointer-events-none absolute top-1/2 left-3 size-4 -translate-y-1/2" />

            <Input
              type="search"
              value={search}
              onChange={(event) => onSearchChange(event.target.value)}
              placeholder="Cari materi atau kelas..."
              aria-label="Cari tugas"
              className="text-primary-text placeholder:text-secondary-text/70 focus-visible:border-secondary focus-visible:ring-secondary/15 h-11 w-full rounded-xl border-[#DCE6ED] bg-white pl-10 text-sm shadow-none"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default TaskToolbar;
