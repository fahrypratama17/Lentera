"use client";

import { useQuery } from "@tanstack/react-query";

import {
  getMaterialsAction,
  getMaterialDetailAction,
  getProgressSummaryAction,
} from "../materials/action";
import type {
  MaterialDetailDto,
  MaterialListItemDto,
  ProgressSummaryDto,
} from "../materials/dto";

type Id = string | number | undefined;

export const materialKeys = {
  all: ["materials"] as const,
  list: (cohortId: Id) => [...materialKeys.all, "list", cohortId] as const,
  detail: (cohortId: Id, materialId: Id) =>
    [...materialKeys.all, "detail", cohortId, materialId] as const,
  progress: (cohortId: Id) =>
    [...materialKeys.all, "progress", cohortId] as const,
};

export function useMaterialsQuery(cohortId: Id) {
  return useQuery<unknown, Error, MaterialListItemDto[]>({
    queryKey: materialKeys.list(cohortId),
    queryFn: () => getMaterialsAction(cohortId!),
    select: (res) => (res as { data: MaterialListItemDto[] }).data,
    enabled: !!cohortId,
  });
}

export function useMaterialDetailQuery(cohortId: Id, materialId: Id) {
  return useQuery<unknown, Error, MaterialDetailDto>({
    queryKey: materialKeys.detail(cohortId, materialId),
    queryFn: () => getMaterialDetailAction(cohortId!, materialId!),
    select: (res) => (res as { data: MaterialDetailDto }).data,
    enabled: !!cohortId && !!materialId,
  });
}

export function useProgressSummaryQuery(cohortId: Id) {
  return useQuery<unknown, Error, ProgressSummaryDto>({
    queryKey: materialKeys.progress(cohortId),
    queryFn: () => getProgressSummaryAction(cohortId!),
    select: (res) => (res as { data: ProgressSummaryDto }).data,
    enabled: !!cohortId,
  });
}