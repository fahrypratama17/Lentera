import { api } from "@/shared/lib/axios";

import type {
  MaterialListResponseDto,
  MaterialDetailResponseDto,
  ProgressSummaryResponseDto,
} from "./dto";

type Id = string | number;

const endpoints = {
  list: (cohortId: Id) => `/cohorts/${cohortId}/materials`,
  detail: (cohortId: Id, materialId: Id) =>
    `/cohorts/${cohortId}/materials/${materialId}`,
  progress: (cohortId: Id) => `/cohorts/${cohortId}/progress/me`,
};

export async function getMaterialsAction(
  cohortId: Id,
): Promise<MaterialListResponseDto> {
  const { data } = await api.get<MaterialListResponseDto>(
    endpoints.list(cohortId),
  );

  return data;
}

export async function getMaterialDetailAction(
  cohortId: Id,
  materialId: Id,
): Promise<MaterialDetailResponseDto> {
  const { data } = await api.get<MaterialDetailResponseDto>(
    endpoints.detail(cohortId, materialId),
  );

  return data;
}

export async function getProgressSummaryAction(
  cohortId: Id,
): Promise<ProgressSummaryResponseDto> {
  const { data } = await api.get<ProgressSummaryResponseDto>(
    endpoints.progress(cohortId),
  );

  return data;
}