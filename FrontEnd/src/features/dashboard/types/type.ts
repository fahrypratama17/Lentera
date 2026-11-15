export type CohortRole = "mentor" | "member";

export type CohortTone = "ocean" | "teal" | "slate";

export interface Cohort {
  id: string;
  title: string;
  description: string;
  mentorName: string;
  period: string;
  role: CohortRole;
  memberCount: number;
  materialCount: number;
  previewQuestion: string;
  tone: CohortTone;
}

export interface CreateCohortValues {
  title: string;
  description: string;
  mentorName: string;
  period: string;
  previewQuestion: string;
}

export type CohortFilter = "all" | CohortRole;

export interface CohortCardProps {
  cohort: Cohort;
  onOpen: (cohort: Cohort) => void;
}

export const toneClasses: Record<Cohort["tone"], string> = {
  ocean: "bg-[linear-gradient(115deg,#1F506D,#0D6684)]",
  teal: "bg-[linear-gradient(115deg,#0D6684,#328CAC)]",
  slate: "bg-[linear-gradient(115deg,#28465D,#54758B)]",
};

export interface DashboardBannerProps {
  onCreate: () => void;
}

export interface CohortToolbarProps {
  filter: CohortFilter;
  onFilterChange: (value: CohortFilter) => void;
  search: string;
  onSearchChange: (value: string) => void;
}

export const filters: {
  label: string;
  value: CohortFilter;
}[] = [
  { label: "Semua", value: "all" },
  { label: "Mentor", value: "mentor" },
  { label: "Member", value: "member" },
];

export interface CohortPreviewDialogProps {
  cohort: Cohort | null;
  onClose: () => void;
}

export interface CohortEmptyStateProps {
  onReset: () => void;
}
