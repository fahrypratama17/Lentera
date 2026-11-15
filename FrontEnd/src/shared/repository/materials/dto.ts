export type ProgressStatus =
  | "not_started"
  | "in_progress"
  | "submitted"
  | "completed";

export type ReviewDecision = "approved" | "revision";

export interface MaterialListItemDto {
  id: number;
  position: number;
  title: string;
  deadline: string | null;
  status: ProgressStatus;
  reviewDecision: ReviewDecision | null;
  reviewNote: string | null;
}

export interface MaterialListResponseDto {
  data: MaterialListItemDto[];
  message?: string;
}

export interface MaterialReferenceDto {
  title?: string;
  url: string;
}

export interface MaterialDetailDto extends MaterialListItemDto {
  guide: string;
  references: MaterialReferenceDto[] | null;
  submissionContent: string | null;
  submissionUrl: string | null;
  submittedAt: string | null;
}

export interface MaterialDetailResponseDto {
  data: MaterialDetailDto;
  message?: string;
}

export interface ProgressSummaryDto {
  totalMaterials: number;
  completedMaterials: number;
  percentage: number;
}

export interface ProgressSummaryResponseDto {
  data: ProgressSummaryDto;
  message?: string;
}