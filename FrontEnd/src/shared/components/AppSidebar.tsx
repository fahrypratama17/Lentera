"use client";

import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import {
  BookOpen,
  ClipboardList,
  LayoutDashboard,
  LogOut,
  UserRound,
} from "lucide-react";

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
} from "@/shared/components/ui/sidebar";

interface AppSidebarProps {
  onLogout: () => void | Promise<void>;
}

const navigation = [
  {
    title: "Dashboard",
    href: "/dashboard",
    icon: LayoutDashboard,
  },
  {
    title: "Kelas Saya",
    href: "/cohort",
    icon: BookOpen,
  },
  {
    title: "Tugas",
    href: "/tasks",
    icon: ClipboardList,
  },
  {
    title: "Profil",
    href: "/profile",
    icon: UserRound,
  },
];

const AppSidebar = ({ onLogout }: AppSidebarProps) => {
  const pathname = usePathname();

  const isActiveRoute = (href: string) => {
    return pathname === href || pathname.startsWith(`${href}/`);
  };

  return (
    <Sidebar collapsible="icon" className="border-border border-r">
      <SidebarHeader className="border-border h-16 shrink-0 justify-center border-b px-3 py-0 group-data-[collapsible=icon]:px-1">
        <Link
          href="/dashboard"
          className="flex min-w-0 items-center gap-3 rounded-xl px-2 group-data-[collapsible=icon]:justify-center group-data-[collapsible=icon]:px-0"
        >
          <Image
            src="/logo.webp"
            alt="Logo Lentera"
            width={38}
            height={38}
            className="size-9 shrink-0 object-contain"
          />
          <span className="text-primary truncate text-xl font-bold tracking-tight group-data-[collapsible=icon]:hidden">
            Lentera
          </span>
        </Link>
      </SidebarHeader>

      <SidebarContent className="pt-4">
        <SidebarGroup className="px-3">
          <SidebarGroupLabel className="text-secondary-text mb-2 h-8 px-3 text-sm font-semibold group-data-[collapsible=icon]:hidden">
            Menu Utama
          </SidebarGroupLabel>

          <SidebarGroupContent>
            <SidebarMenu className="gap-1.5 group-data-[collapsible=icon]:gap-4">
              {navigation.map((item) => {
                const Icon = item.icon;
                const isActive = isActiveRoute(item.href);

                return (
                  <SidebarMenuItem key={item.href}>
                    <SidebarMenuButton
                      render={<Link href={item.href} />}
                      tooltip={item.title}
                      isActive={isActive}
                      className="text-primary-text hover:bg-secondary-light hover:text-primary data-[active=true]:bg-secondary-light data-[active=true]:text-primary h-12 rounded-xl px-3 text-base font-medium transition-colors data-[active=true]:font-semibold"
                    >
                      <Icon className="size-5! shrink-0" />
                      <span className="truncate">{item.title}</span>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter className="border-border gap-3 border-t p-3">
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              type="button"
              tooltip="Keluar"
              onClick={() => void onLogout()}
              className="text-destructive hover:bg-destructive/10 hover:text-destructive h-12 cursor-pointer rounded-xl px-3 text-base font-semibold"
            >
              <LogOut className="size-5! shrink-0" />
              <span className="truncate">Keluar</span>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  );
};

export default AppSidebar;
