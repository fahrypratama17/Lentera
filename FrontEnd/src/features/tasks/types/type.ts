import type { CohortRole, CohortTone } from "@/features/dashboard/types/type";

export type TaskStatus =
  | "not_started"
  | "in_progress"
  | "submitted"
  | "completed";

export type ReviewDecision = "approved" | "revision";

export type TaskFilter =
  | "all"
  | "not_started"
  | "in_progress"
  | "submitted"
  | "completed"
  | "overdue";

export interface TaskItem {
  id: string;
  cohortId: string;
  cohortTitle: string;
  cohortTone: CohortTone;
  materialId: string;
  materialTitle: string;
  position: number;
  deadline: string | null;

  status: TaskStatus;
  submissionContent: string | null;
  submissionUrl: string | null;
  submittedAt: string | null;

  reviewDecision: ReviewDecision | null;
  reviewNote: string | null;
  reviewerName: string | null;
  reviewedAt: string | null;

  startedAt: string | null;
  completedAt: string | null;

  role: CohortRole;
}

export interface TaskSummary {
  total: number;
  notStarted: number;
  inProgress: number;
  submitted: number;
  completed: number;
  overdue: number;
  completionRate: number;
}

export interface SubmitTaskValues {
  content: string;
  url: string;
}

export interface TaskFilterOption {
  label: string;
  value: TaskFilter;
}

export interface TaskCohortOption {
  label: string;
  value: string;
}


export interface TaskCardProps {
  task: TaskItem;
  onOpen: (task: TaskItem) => void;
}

export interface TaskListProps {
  tasks: TaskItem[];
  onOpen: (task: TaskItem) => void;
}

export interface TaskToolbarProps {
  filter: TaskFilter;
  cohortId: string;
  search: string;
  cohortOptions: TaskCohortOption[];
  onFilterChange: (value: TaskFilter) => void;
  onCohortChange: (value: string) => void;
  onSearchChange: (value: string) => void;
}

export interface TaskSummaryProps {
  summary: TaskSummary;
}

export interface TaskPaginationProps {
  currentPage: number;
  totalPages: number;
  totalItems: number;
  pageSize: number;
  onPageChange: (page: number) => void;
}

export interface TaskEmptyStateProps {
  onReset: () => void;
}

export interface TaskDialogProps {
  task: TaskItem | null;
  open: boolean;
  submitting: boolean;
  error: string | null;
  onClose: () => void;
  onSubmit: (values: SubmitTaskValues) => void;
  onStart: () => void;
}


export const TASK_FILTERS: TaskFilterOption[] = [
  { label: "Semua", value: "all" },
  { label: "Belum mulai", value: "not_started" },
  { label: "Dikerjakan", value: "in_progress" },
  { label: "Menunggu review", value: "submitted" },
  { label: "Selesai", value: "completed" },
  { label: "Terlambat", value: "overdue" },
];

export const TASK_STATUS_CONFIG: Record<
  TaskStatus,
  { label: string; className: string }
> = {
  not_started: {
    label: "Belum dimulai",
    className: "bg-slate-100 text-slate-600",
  },
  in_progress: {
    label: "Sedang dikerjakan",
    className: "bg-amber-100 text-amber-700",
  },
  submitted: {
    label: "Menunggu review",
    className: "bg-blue-100 text-blue-700",
  },
  completed: {
    label: "Selesai",
    className: "bg-emerald-100 text-emerald-700",
  },
};

export const REVIEW_DECISION_CONFIG: Record<
  ReviewDecision,
  { label: string; className: string }
> = {
  approved: { label: "Disetujui", className: "text-emerald-700" },
  revision: { label: "Perlu revisi", className: "text-amber-700" },
};
