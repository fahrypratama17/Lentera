import {
  BookOpen,
  CalendarDays,
  GraduationCap,
  UsersRound,
} from "lucide-react";

import { toneClasses } from "@/features/dashboard/types/type";
import { formatPeriod } from "../utils/format";
import type {
  CohortDetail,
  MaterialWithProgress,
} from "../types/type";

interface CohortDetailHeaderProps {
  cohort: CohortDetail;
  materials: MaterialWithProgress[];
}

const CohortDetailHeader = ({ cohort, materials }: CohortDetailHeaderProps) => {
  const total = materials.length;
  const completed = materials.filter((m) => m.status === "completed").length;
  const percent = total > 0 ? Math.round((completed / total) * 100) : 0;

  return (
    <div className="flex flex-col gap-4">
      <section
        className={`relative overflow-hidden rounded-2xl p-6 text-white sm:p-8 ${toneClasses[cohort.tone]}`}
      >
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -top-16 -right-12 size-52 rounded-full border-28 border-white/10"
        />
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -right-6 -bottom-20 size-48 rounded-full bg-white/10"
        />

        <div className="relative z-10 flex flex-col gap-4">
          <span className="w-fit rounded-full border border-white/20 bg-white/15 px-3 py-1 text-xs font-semibold backdrop-blur-sm">
            {cohort.role === "mentor" ? "Mentor" : "Member"}
          </span>

          <div>
            <h1 className="text-2xl leading-tight font-bold sm:text-3xl">
              {cohort.title}
            </h1>
            <p className="mt-2 max-w-2xl text-sm leading-relaxed text-white/85">
              {cohort.description}
            </p>
          </div>

          <div className="flex flex-wrap gap-x-6 gap-y-2 text-sm text-white/85">
            <span className="flex items-center gap-2">
              <GraduationCap className="size-4" />
              {cohort.mentorName}
            </span>
            <span className="flex items-center gap-2">
              <CalendarDays className="size-4" />
              {formatPeriod(cohort.startDate, cohort.endDate)}
            </span>
          </div>
        </div>
      </section>

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
        <div className="flex items-center gap-3 rounded-2xl border border-[#DCE6ED] bg-white p-4">
          <div className="bg-secondary-light text-primary flex size-10 items-center justify-center rounded-xl">
            <UsersRound className="size-5" />
          </div>
          <div>
            <p className="text-primary-text text-lg font-bold">
              {cohort.memberCount}
            </p>
            <p className="text-secondary-text text-xs">Anggota</p>
          </div>
        </div>

        <div className="flex items-center gap-3 rounded-2xl border border-[#DCE6ED] bg-white p-4">
          <div className="bg-secondary-light text-primary flex size-10 items-center justify-center rounded-xl">
            <BookOpen className="size-5" />
          </div>
          <div>
            <p className="text-primary-text text-lg font-bold">{total}</p>
            <p className="text-secondary-text text-xs">Materi</p>
          </div>
        </div>

        <div className="flex flex-col justify-center gap-2 rounded-2xl border border-[#DCE6ED] bg-white p-4">
          <div className="flex items-baseline justify-between">
            <p className="text-secondary-text text-xs">Progres kamu</p>
            <p className="text-primary-text text-sm font-bold">
              {completed}/{total} · {percent}%
            </p>
          </div>
          <div
            role="progressbar"
            aria-valuenow={percent}
            aria-valuemin={0}
            aria-valuemax={100}
            className="bg-secondary-light h-2 overflow-hidden rounded-full"
          >
            <div
              className="bg-primary h-full rounded-full transition-all"
              style={{ width: `${percent}%` }}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default CohortDetailHeader;