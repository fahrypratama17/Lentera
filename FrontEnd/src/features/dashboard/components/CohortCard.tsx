import { ArrowUpRight, BookOpen, UsersRound } from "lucide-react";

import { Card, CardContent } from "@/shared/components/ui/card";
import { Button } from "@/shared/components/ui/button";

import type { Cohort } from "@/features/dashboard/types/type";

interface CohortCardProps {
  cohort: Cohort;
  onOpen: (cohort: Cohort) => void;
}

const toneClasses: Record<Cohort["tone"], string> = {
  ocean: "bg-[linear-gradient(115deg,#1F506D,#0D6684)]",
  teal: "bg-[linear-gradient(115deg,#0D6684,#328CAC)]",
  slate: "bg-[linear-gradient(115deg,#28465D,#54758B)]",
};

const CohortCard = ({ cohort, onOpen }: CohortCardProps) => {
  const isMentor = cohort.role === "mentor";

  const initials = cohort.mentorName
    .split(" ")
    .slice(0, 2)
    .map((name) => name.charAt(0))
    .join("")
    .toUpperCase();

  return (
    <Card className="group h-full min-w-0 gap-0 overflow-hidden rounded-2xl border border-[#DCE6ED] bg-white py-0 shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-[0_16px_40px_-18px_rgba(31,80,109,0.3)]">
      {/* Colored Header */}
      <div
        className={`relative flex min-h-[155px] flex-col justify-between overflow-hidden p-5 text-white ${toneClasses[cohort.tone]} `}
      >
        {/* Background Circles */}
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -top-14 -right-10 size-44 rounded-full border-[28px] border-white/10"
        />

        <div
          aria-hidden="true"
          className="pointer-events-none absolute -right-4 -bottom-16 size-40 rounded-full bg-white/10"
        />

        <div className="relative z-10 flex items-start justify-between gap-3">
          <span className="rounded-full border border-white/20 bg-white/15 px-3 py-1 text-xs font-semibold text-white backdrop-blur-sm">
            {isMentor ? "Mentor" : "Member"}
          </span>

          <span className="text-xs text-white/75">Cohort</span>
        </div>

        <div className="relative z-10 space-y-1.5">
          <button
            type="button"
            onClick={() => onOpen(cohort)}
            className="line-clamp-2 cursor-pointer text-left text-xl leading-snug font-bold text-white hover:underline"
          >
            {cohort.title}
          </button>

          <p className="truncate text-sm font-medium text-white/85">
            {cohort.mentorName}
          </p>
        </div>
      </div>

      <CardContent className="flex flex-1 flex-col gap-4 px-5 py-5">
        {/* Information */}
        <div className="flex items-center gap-3">
          <div className="bg-secondary-light text-primary flex size-10 shrink-0 items-center justify-center rounded-full text-sm font-bold">
            {initials}
          </div>

          <div className="min-w-0">
            <p className="text-primary-text truncate text-sm font-semibold">
              {cohort.mentorName}
            </p>

            <p className="text-secondary-text text-xs">Mentor Kelas</p>
          </div>
        </div>

        {/* Question / Material Preview */}
        <div className="bg-surface flex min-h-[115px] flex-1 flex-col gap-2 rounded-xl p-4">
          <div className="text-secondary flex items-center gap-2 text-xs font-bold">
            <BookOpen className="size-4" />
            Preview Materi / Tugas
          </div>

          <p className="text-primary-text line-clamp-3 text-sm leading-relaxed">
            {cohort.previewQuestion || "Belum ada materi yang dipublikasikan."}
          </p>
        </div>

        {/* Footer */}
        <div className="flex items-center justify-between gap-2 border-t border-[#DCE6ED] pt-4">
          <div className="text-secondary-text flex flex-wrap items-center gap-3 text-xs">
            <span className="flex items-center gap-1">
              <UsersRound className="size-4" />
              {cohort.memberCount}
            </span>

            <span className="flex items-center gap-1">
              <BookOpen className="size-4" />
              {cohort.materialCount} Materi
            </span>
          </div>

          <Button
            type="button"
            variant="ghost"
            size="sm"
            onClick={() => onOpen(cohort)}
            className="text-primary hover:bg-secondary-light hover:text-secondary cursor-pointer rounded-lg font-bold"
          >
            Lihat
            <ArrowUpRight className="size-4" />
          </Button>
        </div>
      </CardContent>
    </Card>
  );
};

export default CohortCard;
