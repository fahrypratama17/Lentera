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
