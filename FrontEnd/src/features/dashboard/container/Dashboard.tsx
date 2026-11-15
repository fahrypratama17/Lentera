"use client";

import DashboardBanner from "@/features/dashboard/components/DashboardBanner";
import { useDashboardCohorts } from "@/features/dashboard/hooks/useDashboard";
import { useState } from "react";
import { Cohort } from "@/features/dashboard/types/type";

const Dashboard = () => {
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

  return (
    <div>
      <DashboardBanner onCreate={() => setCreateOpen(true)} />
    </div>
  );
};

export default Dashboard;
