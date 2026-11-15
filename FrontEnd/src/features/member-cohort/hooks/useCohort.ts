"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

import { useDashboardController } from "@/features/dashboard/hooks/useDashboardController";
import { buildCohortSlug } from "../utils/slugify";

type CohortItem = ReturnType<
  typeof useDashboardController
>["filteredCohorts"][number];

export const COHORTS_PER_PAGE = 6;

export const useCohortController = () => {
  const router = useRouter();
  const dashboard = useDashboardController();
  const [page, setPage] = useState(1);

  const totalItems = dashboard.filteredCohorts.length;
  const totalPages = Math.max(1, Math.ceil(totalItems / COHORTS_PER_PAGE));

  const currentPage = Math.min(page, totalPages);

  const start = (currentPage - 1) * COHORTS_PER_PAGE;
  const paginatedCohorts = dashboard.filteredCohorts.slice(
    start,
    start + COHORTS_PER_PAGE,
  );

  const setSearch = (value: string) => {
    dashboard.setSearch(value);
    setPage(1);
  };

  const setFilter = (value: Parameters<typeof dashboard.setFilter>[0]) => {
    dashboard.setFilter(value);
    setPage(1);
  };

  const resetFilters = () => {
    dashboard.resetFilters();
    setPage(1);
  };

  const openCohort = (cohort: CohortItem) => {
    router.push(`/cohort/${buildCohortSlug(cohort.title, cohort.id)}/materi`);
    };

  return {
    paginatedCohorts,
    totalItems,
    totalPages,
    currentPage,
    setPage,
    search: dashboard.search,
    setSearch,
    filter: dashboard.filter,
    setFilter,
    resetFilters,
    createCohort: dashboard.createCohort,
    createOpen: dashboard.createOpen,
    openCreateDialog: dashboard.openCreateDialog,
    closeCreateDialog: dashboard.closeCreateDialog,
    openCohort,
  };
};