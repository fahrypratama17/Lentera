"use client";

import type { ReactNode } from "react";
import { usePathname } from "next/navigation";
import AppSidebar from "@/shared/components/AppSidebar";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/shared/components/ui/sidebar";
import { useLogout } from "@/features/auth/hooks/useLogout";

interface DashboardShellProps {
  children: ReactNode;
}

const pageTitles: Record<string, string> = {
  "/dashboard": "Dashboard",
  "/cohort": "Kelas Saya",
  "/tasks": "Tugas",
  "/profile": "Profil",
};

export default function DashboardShell({ children }: DashboardShellProps) {
  const { logout, isPending: logoutPending, error: logoutError } = useLogout();
  const pathname = usePathname();

  const currentPage = Object.keys(pageTitles).find(
    (route) => pathname === route || pathname.startsWith(`${route}/`),
  );

  const title = currentPage ? pageTitles[currentPage] : "Dashboard";

  const handleLogout = async () => {
    // TODO: Implement Logout here
  };

  return (
    <SidebarProvider>
      <AppSidebar
        onLogout={logout}
        logoutPending={logoutPending}
        logoutError={logoutError}
      />

      <SidebarInset className="bg-surface min-w-0">
        <header className="border-border sticky top-0 z-20 flex h-16 shrink-0 items-center justify-between border-b bg-white/95 px-5 backdrop-blur-md sm:px-8">
          <div className="flex items-center gap-3">
            <SidebarTrigger className="text-primary-text cursor-pointer" />

            <div className="bg-border h-5 w-px" />

            <span className="text-primary-text text-base font-semibold">
              {title}
            </span>
          </div>
        </header>

        {children}
      </SidebarInset>
    </SidebarProvider>
  );
}
