export type CohortRole = "mentor" | "member";


export interface CohortDto {
  id: number;
  name: string;
  description: string;
  start_date: string;
  end_date: string;
  role: CohortRole;
  member_count: number;
  joined_at: string;
  created_at: string;
}

/** GET /cohorts/:id & PUT /cohorts/:id — `dto.CohortDetailResponse` */
export interface CohortDetailDto {
  id: number;
  name: string;
  description: string;
  start_date: string;
  end_date: string;
  my_role: CohortRole;
  member_count: number;
  created_by: number;
  created_at: string;
  updated_at: string;
}

/* ────────────────────────────────────────────────────────────
   Request payload — `internal/app/cohort/dto/request.go`
   ──────────────────────────────────────────────────────────── */

/** POST /cohorts — name min 3 max 100, description max 1000 */
export interface CreateCohortDto {
  name: string;
  description: string;
  /** format `YYYY-MM-DD` atau RFC3339 */
  start_date: string;
  /** format `YYYY-MM-DD` atau RFC3339 */
  end_date: string;
}

/** PUT /cohorts/:id — field wajib sama seperti create */
export type UpdateCohortDto = CreateCohortDto;

/* ────────────────────────────────────────────────────────────
   Envelope — `pkg/response.Body`
   ──────────────────────────────────────────────────────────── */

export interface ApiBody<T> {
  success: boolean;
  message?: string;
  data?: T;
  meta?: unknown;
}

export type CohortListResponseDto = ApiBody<CohortDto[]>;
export type CohortDetailResponseDto = ApiBody<CohortDetailDto>;
export type CohortMutationResponseDto = ApiBody<CohortDetailDto | null>;

/* ────────────────────────────────────────────────────────────
   Upload PDF — belum ada route-nya di backend,
   kontrak di bawah mengikuti pola `pkg/storage`
   ──────────────────────────────────────────────────────────── */

export interface UploadCohortPdfDto {
  file: File;
  /** opsional, mis. "silabus" | "modul" */
  document_type?: string;
}

export interface CohortPdfDto {
  id: number;
  cohort_id: number;
  file_name: string;
  file_url: string;
  file_size: number;
  uploaded_at: string;
}

export type UploadCohortPdfResponseDto = ApiBody<CohortPdfDto>;
