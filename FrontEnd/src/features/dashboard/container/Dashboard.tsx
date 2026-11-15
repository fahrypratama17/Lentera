"use client";

import DashboardBanner from "@/features/dashboard/components/DashboardBanner";
import DashboardHeader from "@/features/dashboard/components/DashboardHeader";
import CohortCard from "@/shared/components/CohortCard";
import CreateCohortDialog from "@/shared/components/CreateCohortDialog";
import CohortPreviewDialog from "@/features/dashboard/components/CohortPreviewDialog";
import { useDashboardController } from "@/features/dashboard/hooks/useDashboardController";
import CohortEmptyState from "@/shared/components/CohortEmptyState";
import Link from "next/link";
import { ArrowRight } from "lucide-react";

const Dashboard = () => {
  const {
    filteredCohorts,
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

        <Link
          href="/cohort"
          className="text-sm-semibold text-secondary hover:text-primary flex cursor-pointer items-center justify-end gap-2 rounded-lg"
        >
          Lihat Semua
          <ArrowRight className="size-4" />
        </Link>

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
