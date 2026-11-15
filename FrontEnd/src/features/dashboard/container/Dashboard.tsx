"use client";

import DashboardBanner from "@/features/dashboard/components/DashboardBanner";
import DashboardHeader from "@/features/dashboard/components/DashboardHeader";
import CohortToolbar from "@/features/dashboard/components/CohortToolbar";
import CohortCard from "@/features/dashboard/components/CohortCard";
import CreateCohortDialog from "@/features/dashboard/components/CreateCohortDialog";
import CohortPreviewDialog from "@/features/dashboard/components/CohortPreviewDialog";
import { useDashboardController } from "@/features/dashboard/hooks/useDashboardController";
import CohortEmptyState from "@/features/dashboard/components/CohortEmptyState";

const Dashboard = () => {
  const {
    filteredCohorts,
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
  } = useDashboardController();

  return (
    <div className="flex w-full min-w-0 flex-col gap-9 px-5 py-7 sm:px-8 lg:px-10">
      <DashboardBanner onCreate={openCreateDialog} />

      <section id="classes" className="flex scroll-mt-24 flex-col gap-6">
        <DashboardHeader />

        <CohortToolbar
          filter={filter}
          onFilterChange={setFilter}
          search={search}
          onSearchChange={setSearch}
        />

        {filteredCohorts.length > 0 ? (
          <div className="grid grid-cols-1 gap-5 md:grid-cols-2 xl:grid-cols-3">
            {filteredCohorts.map((cohort) => (
              <CohortCard
                key={cohort.id}
                cohort={cohort}
                onOpen={openPreviewDialog}
              />
            ))}
          </div>
        ) : (
          <CohortEmptyState onReset={resetFilters} />
        )}
      </section>

      <CreateCohortDialog
        open={createOpen}
        onClose={closeCreateDialog}
        onCreate={createCohort}
      />

      <CohortPreviewDialog
        cohort={selectedCohort}
        onClose={closePreviewDialog}
      />
    </div>
  );
};

export default Dashboard;
