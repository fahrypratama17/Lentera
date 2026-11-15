"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { BookOpen, UsersRound, type LucideIcon } from "lucide-react";

import type { CohortRole } from "../types/type";

interface TabItem {
  label: string;
  segment: string;
  icon: LucideIcon;
  roles: CohortRole[];
}

const tabs: TabItem[] = [
  { label: "Materi", segment: "materi", icon: BookOpen, roles: ["member", "mentor"] },
  { label: "Anggota", segment: "anggota", icon: UsersRound, roles: ["member", "mentor"] },
];

interface CohortTabsProps {
  slug: string;
  role: CohortRole;
}

const CohortTabs = ({ slug, role }: CohortTabsProps) => {
  const pathname = usePathname();
  const visibleTabs = tabs.filter((tab) => tab.roles.includes(role));

  return (
    <nav aria-label="Bagian kelas" className="border-b border-[#DCE6ED]">
      <ul className="-mb-px flex gap-1 overflow-x-auto">
        {visibleTabs.map(({ label, segment, icon: Icon }) => {
          const href = `/cohort/${slug}/${segment}`;
          const active = pathname === href || pathname.startsWith(`${href}/`);

          return (
            <li key={segment}>
              <Link
                href={href}
                aria-current={active ? "page" : undefined}
                className={`flex items-center gap-2 border-b-2 px-4 py-3 text-sm font-semibold whitespace-nowrap transition-colors ${
                  active
                    ? "border-primary text-primary"
                    : "text-secondary-text hover:text-primary border-transparent"
                }`}
              >
                <Icon className="size-4" />
                {label}
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
};

export default CohortTabs;