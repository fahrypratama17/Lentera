import type { ReactNode } from "react";
import DashboardShell from "@/shared/layout/DashboardShell";

interface DashboardLayoutProps {
  children: ReactNode;
}

const layout = ({ children }: DashboardLayoutProps) => {
  return <DashboardShell>{children}</DashboardShell>;
};

export default layout;
