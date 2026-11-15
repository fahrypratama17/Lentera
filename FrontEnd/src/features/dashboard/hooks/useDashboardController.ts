import { useState } from "react";

import { useDashboardCohorts } from "./useDashboard";
import type { Cohort } from "../types/type";

export function useDashboardController() {
  const {
    filteredCohorts,
    counts,
    search,
    setSearch,
    filter,
    setFilter,
    createCohort,
    resetFilters,
  } = useDashboardCohorts();

  const [createOpen, setCreateOpen] = useState(false);
  const [selectedCohort, setSelectedCohort] = useState<Cohort | null>(null);

  const openCreateDialog = () => setCreateOpen(true);
  const closeCreateDialog = () => setCreateOpen(false);

  const openPreviewDialog = (cohort: Cohort) => {
    setSelectedCohort(cohort);
  };

  const closePreviewDialog = () => {
    setSelectedCohort(null);
  };

  return {
    filteredCohorts,
    counts,
    search,
    setSearch,
    filter,
    setFilter,
    createCohort,
    resetFilters,
    createOpen,
    selectedCohort,
    openCreateDialog,
    closeCreateDialog,
    openPreviewDialog,
    closePreviewDialog,
  };
}
