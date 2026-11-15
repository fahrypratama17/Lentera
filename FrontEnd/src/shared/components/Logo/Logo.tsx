"use client";

import Link from "next/link";
import Image from "next/image";

export function Logo({ onNavigate }: { onNavigate?: () => void }) {
  return (
    <Link
      href="#top"
      onClick={onNavigate}
      aria-label="Lentera - Learning Tracker"
      className="focus-visible:ring-ring flex w-fit shrink-0 items-center gap-3 rounded-full focus-visible:ring-2 focus-visible:outline-none"
    >
      <Image
        src="/logo.webp"
        alt="Logo Lentera"
        width={48}
        height={48}
        className="size-12 object-contain"
      />

      <span className="text-primary text-xl font-extrabold tracking-tight">
        Lentera
      </span>
    </Link>
  );
}
