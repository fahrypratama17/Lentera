import { ReactNode, Suspense } from "react";
import DashboardShell from "@/shared/layout/DashboardShell";

interface DashboardLayoutProps {
  children: ReactNode;
}

const layout = ({ children }: DashboardLayoutProps) => {
  return (
    <Suspense fallback={<>{children}</>}>
      <DashboardShell>{children}</DashboardShell>
    </Suspense>
  );
};

export default layout;
