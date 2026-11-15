import type { Metadata } from "next";
import Dashboard from "@/features/dashboard/container/Dashboard";

export const metadata: Metadata = {
  title: "Dashboard | StudyTrack",
  description: "Kelola cohort dan progres belajar.",
};

const page = () => {
  return <Dashboard />;
};

export default page;
