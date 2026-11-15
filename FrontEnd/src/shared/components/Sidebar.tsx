"use client";

import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { BookOpen, LayoutDashboard } from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarSeparator,
} from "@/shared/components/ui/sidebar";

const navigation = [
  {
    title: "Dashboard",
    href: "/dashboard",
    icon: LayoutDashboard,
  },
  {
    title: "Kelas Saya",
    href: "/dashboard#classes",
    icon: BookOpen,
  },
];

const AppSidebar = () => {
  const pathname = usePathname();

  return (
    <Sidebar collapsible="icon">
      <SidebarHeader className="p-3">
        <Link
          href="/dashboard"
          className="flex items-center gap-3 rounded-xl px-2 py-2"
        >
          <Image
            src="/logo.webp"
            alt="Logo StudyTrack"
            width={34}
            height={34}
            className="size-8 shrink-0 object-contain"
          />

          <span className="text-lg-bold text-primary truncate group-data-[collapsible=icon]:hidden">
            StudyTrack
          </span>
        </Link>
      </SidebarHeader>

      <SidebarSeparator />

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Menu Utama</SidebarGroupLabel>

          <SidebarGroupContent>
            <SidebarMenu>
              {navigation.map((item) => {
                const Icon = item.icon;
                const isActive =
                  pathname === "/dashboard" && item.href === "/dashboard";

                return (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton
                      render={<Link href={item.href} />}
                      tooltip={item.title}
                      isActive={isActive}
                      className="text-sm-semibold h-10 rounded-lg px-3"
                    >
                      <Icon className="size-4" />
                      <span>{item.title}</span>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter className="p-3">
        <div className="bg-secondary-light rounded-xl p-3 group-data-[collapsible=icon]:hidden">
          <p className="text-sm-bold text-primary">StudyTrack</p>

          <p className="text-secondary-text mt-1 text-xs">
            Belajar dan berkembang bersama.
          </p>
        </div>
      </SidebarFooter>
    </Sidebar>
  );
};

export default AppSidebar;
