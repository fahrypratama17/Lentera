import type { Metadata } from "next";
import { Suspense } from "react";

import TaskPageContainer from "@/features/tasks/container/TaskPageContainer";
import TaskSkeleton from "@/features/tasks/components/TaskSkeleton";

export const metadata: Metadata = {
  title: "Tugas | StudyTrack",
  description: "Semua tugas dari kelas yang kamu ikuti dalam satu daftar.",
};

export default function TasksPage() {
  return (
    <Suspense fallback={<TaskSkeleton />}>
      <TaskPageContainer />
    </Suspense>
  );
}
