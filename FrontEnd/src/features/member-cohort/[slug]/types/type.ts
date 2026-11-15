import type { CohortCardProps } from "@/features/dashboard/types/type";

export type CohortItem = CohortCardProps["cohort"];
export type CohortRole = "mentor" | "member";

export type ProgressStatus =
  | "not_started"
  | "in_progress"
  | "submitted"
  | "completed";
export type ReviewDecision = "approved" | "revision";

export type CohortDetail = CohortItem & {
  description: string;
  startDate: string; 
  endDate: string; 
};

export interface MaterialWithProgress {
  id: string | number;
  position: number;
  title: string;
  deadline: string | null;
  status: ProgressStatus;
  reviewDecision: ReviewDecision | null;
  reviewNote: string | null;
}

export interface CohortMember {
  id: string;
  name: string;
  email: string;
  role: CohortRole;
  joinedAt: string;
  isMe: boolean;
}

export const progressStatusConfig: Record<
  ProgressStatus,
  { label: string; className: string }
> = {
  not_started: { label: "Belum dimulai", className: "bg-slate-100 text-slate-600" },
  in_progress: { label: "Sedang dikerjakan", className: "bg-amber-100 text-amber-700" },
  submitted: { label: "Menunggu review", className: "bg-blue-100 text-blue-700" },
  completed: { label: "Selesai", className: "bg-emerald-100 text-emerald-700" },
};

export const reviewDecisionConfig: Record<
  ReviewDecision,
  { label: string; className: string }
> = {
  approved: { label: "Disetujui", className: "text-emerald-700" },
  revision: { label: "Perlu revisi", className: "text-amber-700" },
};