import type { ReactNode } from "react";
import AppSidebar from "@/shared/components/AppSidebar";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/shared/components/ui/sidebar";

interface DashboardLayoutProps {
  children: ReactNode;
}

const layout = ({ children }: DashboardLayoutProps) => {
  return (
    <SidebarProvider>
      <AppSidebar />

      <SidebarInset className="bg-surface min-w-0">
        <header className="sticky top-0 z-20 flex h-16 shrink-0 items-center justify-between border-b border-[#DCE6ED] bg-white/95 px-5 backdrop-blur-md sm:px-8">
          <div className="flex items-center gap-3">
            <SidebarTrigger className="text-primary-text cursor-pointer" />
            <div className="h-5 w-px bg-[#DCE6ED]" />
            <span className="text-md-semibold text-primary-text">
              Dashboard
            </span>
          </div>
          <span className="text-sm-medium text-secondary-text hidden sm:block">
            StudyTrack Learning Space
          </span>
        </header>

        {children}
      </SidebarInset>
    </SidebarProvider>
  );
};

export default layout;
