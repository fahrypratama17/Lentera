import type { Metadata } from "next";

import CohortContainer from "@/features/member-cohort/container/CohortContainer";

export const metadata: Metadata = {
  title: "Kelas Saya | StudyTrack",
  description: "Semua kelas yang kamu ikuti dan kelola.",
};

export default function CohortPage() {
  return <CohortContainer />;
}
